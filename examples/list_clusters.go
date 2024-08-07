package main

import (
	"fmt"
	"os"

	"github.com/spectrocloud/palette-api-go/models"
	"github.com/spectrocloud/palette-sdk-go/client"
)

func main() {
	var host, apiKey, projectUid, scope string

	// Parse command line arguments

	numArgs := len(os.Args)
	if numArgs < 3 || numArgs > 4 {
		fmt.Println("Usage: main <host> <apiKey> [projectUid]")
		os.Exit(1)
	}
	switch numArgs {
	case 3:
		host, apiKey = os.Args[1], os.Args[2]
		scope = "tenant"
	case 4:
		host, apiKey, projectUid = os.Args[1], os.Args[2], os.Args[3]
		scope = "project"
	}

	// Initialize a Palette client

	pc := client.New(
		client.WithPaletteURI(host),
		client.WithAPIKey(apiKey),
	)
	if projectUid != "" {
		client.WithScopeProject(projectUid)(pc)
	} else {
		client.WithScopeTenant()(pc)
	}

	// Search for clusters

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
