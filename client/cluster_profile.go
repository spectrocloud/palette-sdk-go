package client

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"

	clientv1 "github.com/spectrocloud/palette-sdk-go/api/client/v1"
	"github.com/spectrocloud/palette-sdk-go/api/models"
	"github.com/spectrocloud/palette-sdk-go/client/apiutil"
)

// GetClusterProfileSummary returns a summary for an existing cluster profile.
func (h *V1Client) GetClusterProfileSummary(uid string) (*models.V1ClusterProfileSummary, error) {
	params := clientv1.NewV1ClusterProfilesUIDSummaryParamsWithContext(h.ctx).
		WithUID(uid)
	resp, err := h.Client.V1ClusterProfilesUIDSummary(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// GetClusterProfileUID returns the UID of a cluster profile based on the profile name and version.
func (h *V1Client) GetClusterProfileUID(profileName, profileVersion string) (string, error) {
	params := clientv1.NewV1ClusterProfilesMetadataParamsWithContext(h.ctx)
	resp, err := h.Client.V1ClusterProfilesMetadata(params)
	if err != nil {
		return "", err
	}
	for _, profile := range resp.Payload.Items {
		if profile.Metadata.Name == profileName && profile.Spec.Version == profileVersion {
			return profile.Metadata.UID, nil
		}
	}
	return "", fmt.Errorf("cluster profile %s not found", profileName)
}

// ImportClusterProfile imports a cluster profile given the profile content as a string.
func (h *V1Client) ImportClusterProfile(profileContent string) (string, error) {
	tmpFile, err := os.CreateTemp(os.TempDir(), "profile-import")
	if err != nil {
		return "", err
	}
	defer func() {
		if err := os.Remove(tmpFile.Name()); err != nil {
			fmt.Printf("failed to remove tmp file %s: %v", tmpFile.Name(), err)
		}
	}()
	if _, err = tmpFile.Write([]byte(profileContent)); err != nil {
		return "", err
	}
	f, err := os.Open(tmpFile.Name())
	if err != nil {
		return "", err
	}
	params := clientv1.NewV1ClusterProfilesImportFileParamsWithContext(h.ctx).
		WithPublish(apiutil.Ptr(true)).
		WithImportFile(f)
	resp, err := h.Client.V1ClusterProfilesImportFile(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

// UpgradeClusterProfile upgrades a cluster profile with the given profile content.
func (h *V1Client) UpgradeClusterProfile(clusterUID string, body *models.V1SpectroClusterProfiles) error {
	params := clientv1.NewV1SpectroClustersUpdateProfilesParamsWithContext(h.ctx).
		WithUID(clusterUID).
		WithBody(body)
	_, err := h.Client.V1SpectroClustersUpdateProfiles(params)
	return err
}

// AttachAddonsToCluster attaches one or more addon profiles to a cluster.
func (h *V1Client) AttachAddonsToCluster(clusterUID string, profileUIDs []string) error {
	// get existing cluster profile uid list on the cluster
	currentClusterInfo, err := h.GetCluster(clusterUID)
	if err != nil {
		return err
	}

	profileList := make([]*models.V1SpectroClusterProfileEntity, 0, len(currentClusterInfo.Spec.ClusterProfileTemplates))
	packList := make([]*models.V1PackValuesEntity, 0, len(currentClusterInfo.Spec.ClusterProfileTemplates))

	for _, clusterProfile := range currentClusterInfo.Spec.ClusterProfileTemplates {
		for _, value := range clusterProfile.Packs {
			packList = append(packList, &models.V1PackValuesEntity{
				Manifests: nil,
				Name:      value.Name,
				Tag:       value.Tag,
				// Type:      models.NewV1PackType(models.V1PackType(value.Type)),
				Type:   models.V1PackType(value.Type),
				Values: value.Values,
			})
		}
		profileList = append(profileList, &models.V1SpectroClusterProfileEntity{
			PackValues: packList,
			UID:        clusterProfile.UID,
		})
	}

	for _, uid := range profileUIDs {
		packs, err := h.GetPacksByProfile(uid)
		if err != nil {
			return err
		}
		packList = make([]*models.V1PackValuesEntity, 0, len(packs))
		for _, value := range packs {
			packList = append(packList, &models.V1PackValuesEntity{
				Manifests: nil,
				Name:      &value.Spec.Name,
				Tag:       value.Spec.Version,
				Type:      value.Spec.Type,
				Values:    value.Spec.Values,
			})
		}
		profileList = append(profileList, &models.V1SpectroClusterProfileEntity{
			PackValues: packList,
			UID:        uid,
		})
	}
	profiles := &models.V1SpectroClusterProfiles{
		Profiles: profileList,
	}

	return h.UpgradeClusterProfile(clusterUID, profiles)
}

// RemoveAddonsFromCluster removes one or more addon profiles from the cluster.
func (h *V1Client) RemoveAddonsFromCluster(clusterUID string, profileUIDs []string) error {
	// Get existing cluster profile UID list on the cluster
	currentClusterInfo, err := h.GetCluster(clusterUID)
	if err != nil {
		return err
	}

	// Create lists to store updated profiles and packs
	profileList := make([]*models.V1SpectroClusterProfileEntity, 0, len(currentClusterInfo.Spec.ClusterProfileTemplates))

	// Create a map for easier lookup of profileUIDs to be removed
	removalMap := make(map[string]bool, len(profileUIDs))
	for _, uid := range profileUIDs {
		removalMap[uid] = true
	}

	// Iterate through current cluster profiles, excluding those in the profileUIDs list
	for _, clusterProfile := range currentClusterInfo.Spec.ClusterProfileTemplates {
		// If the current profile UID is in the removal map, skip this profile
		if removalMap[clusterProfile.UID] {
			continue
		}

		// Retain the profile and pack values if it is not in the removal list
		packList := make([]*models.V1PackValuesEntity, 0, len(clusterProfile.Packs))
		for _, value := range clusterProfile.Packs {
			packList = append(packList, &models.V1PackValuesEntity{
				Manifests: nil,
				Name:      value.Name,
				Tag:       value.Tag,
				Type:      models.V1PackType(value.Type),
				Values:    value.Values,
			})
		}

		profileList = append(profileList, &models.V1SpectroClusterProfileEntity{
			PackValues: packList,
			UID:        clusterProfile.UID,
		})
	}

	// Prepare the new profiles object for updating the cluster profile
	profiles := &models.V1SpectroClusterProfiles{
		Profiles: profileList,
	}

	// Call the update function to update the cluster profile with the remaining profiles
	return h.UpgradeClusterProfile(clusterUID, profiles)
}

// DeleteClusterProfile deletes an existing cluster profile.
func (h *V1Client) DeleteClusterProfile(uid string) error {
	profile, err := h.GetClusterProfile(uid)
	if err != nil {
		return err
	} else if profile == nil {
		return fmt.Errorf("cluster profile %s not found", uid)
	}
	params := clientv1.NewV1ClusterProfilesDeleteParamsWithContext(h.ctx).
		WithUID(uid)
	_, err = h.Client.V1ClusterProfilesDelete(params)
	return err
}

// GetClusterProfile retrieves an existing cluster profile by UID.
func (h *V1Client) GetClusterProfile(uid string) (*models.V1ClusterProfile, error) {
	params := clientv1.NewV1ClusterProfilesGetParamsWithContext(h.ctx).
		WithUID(uid)
	resp, err := h.Client.V1ClusterProfilesGet(params)
	if apiutil.Is404(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

type values struct {
	HeaderMap map[string]string
}

func (v values) Get(key string) string {
	if v.HeaderMap == nil {
		return ""
	}
	return v.HeaderMap[key]
}

// DownloadClusterProfileUIDSpc retrieves an existing cluster profile by UID.
func (h *V1Client) DownloadClusterProfileUIDSpc(w *io.PipeWriter, profileUid, projectUid string) (string, error) {
	type CustomHeader string

	const CustomHeaders CustomHeader = "CustomHeaders"
	if projectUid != "" {
		h.projectUID = projectUid
		h.ctx = context.WithValue(h.ctx, CustomHeaders, values{
			HeaderMap: map[string]string{
				"ProjectUid": projectUid,
			}})
	}
	params := clientv1.NewV1ClusterProfilesUIDSpcDownloadParamsWithContext(h.ctx).
		WithUID(profileUid)

	resp, err := h.Client.V1ClusterProfilesUIDSpcDownload(params, w)
	if apiutil.Is404(err) {
		return "", nil
	} else if err != nil {
		return "", err
	}
	return resp.ContentDisposition, nil
}

// GetClusterProfiles retrieves all cluster profiles.
func (h *V1Client) GetClusterProfiles() ([]*models.V1ClusterProfileMetadata, error) {
	params := clientv1.NewV1ClusterProfilesMetadataParamsWithContext(h.ctx)
	resp, err := h.Client.V1ClusterProfilesMetadata(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload.Items, nil
}

// PatchClusterProfile patches an existing cluster profile's metadata.
func (h *V1Client) PatchClusterProfile(clusterProfile *models.V1ClusterProfileUpdateEntity, metadata *models.V1ProfileMetaEntity) error {
	params := clientv1.NewV1ClusterProfilesUIDMetadataUpdateParamsWithContext(h.ctx).
		WithUID(clusterProfile.Metadata.UID).
		WithBody(metadata)
	_, err := h.Client.V1ClusterProfilesUIDMetadataUpdate(params)
	return err
}

// UpdateClusterProfile updates an existing cluster profile.
func (h *V1Client) UpdateClusterProfile(clusterProfile *models.V1ClusterProfileUpdateEntity) error {
	params := clientv1.NewV1ClusterProfilesUpdateParamsWithContext(h.ctx).
		WithUID(clusterProfile.Metadata.UID).
		WithBody(clusterProfile)
	_, err := h.Client.V1ClusterProfilesUpdate(params)
	return err
}

// CreateClusterProfile creates a new cluster profile.
func (h *V1Client) CreateClusterProfile(clusterProfile *models.V1ClusterProfileEntity) (string, error) {
	params := clientv1.NewV1ClusterProfilesCreateParamsWithContext(h.ctx).
		WithBody(clusterProfile)
	resp, err := h.Client.V1ClusterProfilesCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

// PublishClusterProfile publishes an existing cluster profile.
// TODO: what does it mean to publish a cluster profile?
func (h *V1Client) PublishClusterProfile(uid string) error {
	params := clientv1.NewV1ClusterProfilesPublishParamsWithContext(h.ctx).
		WithUID(uid)
	_, err := h.Client.V1ClusterProfilesPublish(params)
	return err
}

// GetProfileVariables retrieves all variables for a cluster profile.
func (h *V1Client) GetProfileVariables(uid string) ([]*models.V1Variable, error) {
	params := clientv1.NewV1ClusterProfilesUIDVariablesGetParamsWithContext(h.ctx).WithUID(uid)
	resp, err := h.Client.V1ClusterProfilesUIDVariablesGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload.Variables, nil
}

// UpdateProfileVariables updates variables for a cluster profile.
func (h *V1Client) UpdateProfileVariables(variables *models.V1Variables, uid string) error {
	params := clientv1.NewV1ClusterProfilesUIDVariablesPutParamsWithContext(h.ctx).
		WithUID(uid).
		WithBody(variables)
	_, err := h.Client.V1ClusterProfilesUIDVariablesPut(params)
	return err
}

// ExportClusterProfile exports a cluster profile in the specified format.
// Supported formats are yaml and json.
func (h *V1Client) ExportClusterProfile(uid, format string) (bytes.Buffer, error) {
	var buf bytes.Buffer
	writer := io.Writer(&buf)
	params := clientv1.NewV1ClusterProfilesUIDExportParamsWithContext(h.ctx).
		WithUID(uid).
		WithFormat(&format)
	_, err := h.Client.V1ClusterProfilesUIDExport(params, writer)
	return buf, err
}

// BulkDeleteClusterProfiles deletes multiple cluster profiles.
func (h *V1Client) BulkDeleteClusterProfiles(uids []string) (*models.V1BulkDeleteResponse, error) {
	params := clientv1.NewV1ClusterProfilesBulkDeleteParamsWithContext(h.ctx).
		WithBody(&models.V1BulkDeleteRequest{
			Uids: uids,
		})
	resp, err := h.Client.V1ClusterProfilesBulkDelete(params)
	if err != nil {
		return nil, err
	}
	return resp.GetPayload(), nil
}
