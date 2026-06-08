// Package main demonstrates listing Palette clusters using the SDK.
package main

import (
	"fmt"
	"os"

	"github.com/spectrocloud/palette-sdk-go/api/models"
	"github.com/spectrocloud/palette-sdk-go/client"
)

func main() {

	// Read environment variables

	host := os.Getenv("PALETTE_HOST")
	apiKey := os.Getenv("PALETTE_API_KEY")
	projectUID := os.Getenv("PALETTE_PROJECT_UID")
	trace := os.Getenv("PALETTE_TRACE") == "1" || os.Getenv("PALETTE_TRACE") == "true"
	scope := "tenant"

	if host == "" || apiKey == "" {
		fmt.Println("You must specify the PALETTE_HOST and PALETTE_API_KEY environment variables.")
		os.Exit(1)
	}
	if projectUID != "" {
		scope = "project"
	}

	// Initialize a Palette client

	pc := client.New(
		client.WithPaletteURI(host),
		client.WithAPIKey(apiKey),
	)
	if trace {
		client.WithTransportDebug()(pc)
	}
	if projectUID != "" {
		client.WithScopeProject(projectUID)(pc)
	} else {
		client.WithScopeTenant()(pc)
	}

	// Search for clusters

	if trace {
		fmt.Println("HTTP transport debug logging enabled (PALETTE_TRACE). Logs go to stderr.")
	}
	fmt.Printf("Searching for Palette clusters with %s scope...\n", scope)

	clusters, err := pc.SearchClusterSummaries(&models.V1SearchFilterSpec{}, []*models.V1SearchFilterSortSpec{})
	if err != nil {
		panic(err)
	}

	// Display the results

	if len(clusters) == 0 {
		fmt.Println("\nNo clusters found.")
		return
	}

	fmt.Printf("\nFound %d cluster(s):\n", len(clusters))
	for _, cluster := range clusters {
		fmt.Printf("\nName: %s\n", cluster.Metadata.Name)
		fmt.Printf("  Cloud Type: %s\n", cluster.SpecSummary.CloudConfig.CloudType)
		fmt.Printf("  Project: %s\n", cluster.SpecSummary.ProjectMeta.Name)
	}
}
