// Package main demonstrates clearing cluster profile tags via PatchClusterProfile (PLT-2269).
//
// Prerequisites:
//   - A cluster profile with at least one tag/label, OR set PALETTE_SEED_TAG=1 to add a
//     temporary tag before clearing (use on a test profile only).
//
// Usage:
//
//	export PALETTE_HOST="your.palette.host"
//	export PALETTE_API_KEY="your-api-key"
//	export PALETTE_CLUSTER_PROFILE_UID="profile-uid"
//	# optional: export PALETTE_PROJECT_UID="project-uid"
//	# optional: export PALETTE_TRACE=1          # log PATCH body on stderr
//	# optional: export PALETTE_SEED_TAG=1       # add a tag first if profile has none
//	# optional: export PALETTE_RESTORE_TAGS=1   # restore original tags after the test
//
//	go run ./examples/cluster_profile_clear_tags/
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spectrocloud/palette-sdk-go/api/models"
	"github.com/spectrocloud/palette-sdk-go/client"
	"github.com/spectrocloud/palette-sdk-go/client/apiutil"
)

const seedTagKey = "plt-2269-example"

func main() {
	host := os.Getenv("PALETTE_HOST")
	apiKey := os.Getenv("PALETTE_API_KEY")
	projectUID := os.Getenv("PALETTE_PROJECT_UID")
	profileUID := os.Getenv("PALETTE_CLUSTER_PROFILE_UID")
	profileName := os.Getenv("PALETTE_CLUSTER_PROFILE_NAME")
	profileVersion := os.Getenv("PALETTE_CLUSTER_PROFILE_VERSION")
	trace := os.Getenv("PALETTE_TRACE") == "1" || strings.EqualFold(os.Getenv("PALETTE_TRACE"), "true")
	seedTag := os.Getenv("PALETTE_SEED_TAG") == "1" || strings.EqualFold(os.Getenv("PALETTE_SEED_TAG"), "true")
	restoreTags := os.Getenv("PALETTE_RESTORE_TAGS") == "1" || strings.EqualFold(os.Getenv("PALETTE_RESTORE_TAGS"), "true")

	if host == "" || apiKey == "" {
		fmt.Println("Set PALETTE_HOST and PALETTE_API_KEY.")
		os.Exit(1)
	}
	if profileUID == "" && (profileName == "" || profileVersion == "") {
		fmt.Println("Set PALETTE_CLUSTER_PROFILE_UID or PALETTE_CLUSTER_PROFILE_NAME + PALETTE_CLUSTER_PROFILE_VERSION.")
		os.Exit(1)
	}

	pc := client.New(
		client.WithPaletteURI(host),
		client.WithAPIKey(apiKey),
	)
	if trace {
		client.WithTransportDebug()(pc)
		fmt.Println("HTTP transport debug enabled (PALETTE_TRACE). PATCH request body is logged to stderr.")
	}
	if projectUID != "" {
		client.WithScopeProject(projectUID)(pc)
	} else {
		client.WithScopeTenant()(pc)
	}

	if profileUID == "" {
		var err error
		profileUID, err = pc.GetClusterProfileUID(profileName, profileVersion)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Resolved profile %q version %q -> UID %s\n", profileName, profileVersion, profileUID)
	}

	profile, err := pc.GetClusterProfile(profileUID)
	if err != nil {
		panic(err)
	}
	if profile == nil || profile.Metadata == nil {
		fmt.Println("Cluster profile not found.")
		os.Exit(1)
	}

	originalLabels := copyLabels(profile.Metadata.Labels)
	fmt.Printf("\nProfile: %s (UID: %s, version: %s)\n", profile.Metadata.Name, profile.Metadata.UID, profileVersionOrSpec(profile))
	printLabels("Labels before patch", profile.Metadata.Labels)

	if len(profile.Metadata.Labels) == 0 {
		if !seedTag {
			fmt.Println("\nProfile has no tags. Set PALETTE_SEED_TAG=1 to add a temporary tag before clearing, or pick a profile that already has tags.")
			os.Exit(1)
		}
		fmt.Println("\nSeeding temporary tag (PALETTE_SEED_TAG=1)...")
		if err := patchProfileMetadata(pc, profile, map[string]string{seedTagKey: "spectro__tag"}); err != nil {
			panic(err)
		}
		profile, err = pc.GetClusterProfile(profileUID)
		if err != nil {
			panic(err)
		}
		printLabels("Labels after seed", profile.Metadata.Labels)
		if len(profile.Metadata.Labels) == 0 {
			fmt.Println("FAIL: seed tag patch did not add a label.")
			os.Exit(1)
		}
	}

	clearBody := buildPatchMetadata(profile, map[string]string{})
	patchJSON, err := apiutil.MarshalProfileMetaPatchBody(clearBody)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nPatch JSON (note explicit empty labels object):\n%s\n", prettyJSON(patchJSON))

	fmt.Println("\nPatching metadata to clear all tags...")
	if err := pc.PatchClusterProfile(
		&models.V1ClusterProfileUpdateEntity{Metadata: profile.Metadata},
		clearBody,
	); err != nil {
		panic(err)
	}

	after, err := pc.GetClusterProfile(profileUID)
	if err != nil {
		panic(err)
	}
	printLabels("\nLabels after clear patch", after.Metadata.Labels)

	if len(after.Metadata.Labels) > 0 {
		fmt.Println("\nFAIL: tags were not cleared. PLT-2269 fix may not be applied or API did not accept empty labels.")
		os.Exit(1)
	}
	fmt.Println("\nPASS: cluster profile has zero tags after patch.")

	if restoreTags {
		fmt.Println("\nRestoring original tags (PALETTE_RESTORE_TAGS=1)...")
		if err := patchProfileMetadata(pc, after, originalLabels); err != nil {
			panic(err)
		}
		restored, err := pc.GetClusterProfile(profileUID)
		if err != nil {
			panic(err)
		}
		printLabels("Labels after restore", restored.Metadata.Labels)
	}
}

func patchProfileMetadata(pc *client.V1Client, profile *models.V1ClusterProfile, labels map[string]string) error {
	return pc.PatchClusterProfile(
		&models.V1ClusterProfileUpdateEntity{Metadata: profile.Metadata},
		buildPatchMetadata(profile, labels),
	)
}

func buildPatchMetadata(profile *models.V1ClusterProfile, labels map[string]string) *models.V1ProfileMetaEntity {
	description := ""
	if profile.Metadata != nil && profile.Metadata.Annotations != nil {
		description = profile.Metadata.Annotations["description"]
	}
	version := profileVersionOrSpec(profile)
	return &models.V1ProfileMetaEntity{
		Metadata: &models.V1ObjectMetaInputEntity{
			Name: profile.Metadata.Name,
			Annotations: map[string]string{
				"description": description,
			},
			Labels: labels,
		},
		Spec: &models.V1ClusterProfileSpecEntity{
			Version: version,
		},
	}
}

func profileVersionOrSpec(profile *models.V1ClusterProfile) string {
	if profile.Spec != nil && profile.Spec.Version != "" {
		return profile.Spec.Version
	}
	if profile.Spec != nil && len(profile.Spec.Versions) > 0 && profile.Spec.Versions[0] != nil {
		return profile.Spec.Versions[0].Version
	}
	return ""
}

func copyLabels(in map[string]string) map[string]string {
	if len(in) == 0 {
		return nil
	}
	out := make(map[string]string, len(in))
	for k, v := range in {
		out[k] = v
	}
	return out
}

func printLabels(title string, labels map[string]string) {
	fmt.Println(title + ":")
	if len(labels) == 0 {
		fmt.Println("  (none)")
		return
	}
	for k, v := range labels {
		if v == "spectro__tag" {
			fmt.Printf("  %s\n", k)
		} else {
			fmt.Printf("  %s: %s\n", k, v)
		}
	}
}

func prettyJSON(raw []byte) string {
	var v interface{}
	if err := json.Unmarshal(raw, &v); err != nil {
		return string(raw)
	}
	out, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return string(raw)
	}
	return string(out)
}
