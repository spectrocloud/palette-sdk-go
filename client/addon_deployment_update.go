package client

import (
	"log"
	"math/rand"
	"time"

	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
)

func (h *V1Client) UpdateAddonDeployment(client clientV1.ClientService, cluster *models.V1SpectroCluster, body *models.V1SpectroClusterProfiles, newProfile *models.V1ClusterProfile) error {
	uid := cluster.Metadata.UID

	// check if profile id is the same - update, otherwise, delete and create
	is, replaceUID := IsProfileAttachedByName(cluster, newProfile)
	if is {
		body.Profiles[0].ReplaceWithProfile = replaceUID
	}

	resolveNotification := true
	var params *clientV1.V1SpectroClustersPatchProfilesParams
	scope := cluster.Metadata.Annotations["scope"]
	switch scope {
	case "project":
		params = clientV1.NewV1SpectroClustersPatchProfilesParamsWithContext(h.Ctx)
	case "tenant":
		params = clientV1.NewV1SpectroClustersPatchProfilesParams()
	}
	params = params.WithUID(uid).WithBody(body).WithResolveNotification(&resolveNotification)
	err := h.PatchWithRetry(client, params)
	return err
}

func IsProfileAttachedByName(cluster *models.V1SpectroCluster, newProfile *models.V1ClusterProfile) (bool, string) {
	for _, profile := range cluster.Spec.ClusterProfileTemplates {
		if profile.Name == newProfile.Metadata.Name {
			return true, profile.UID
		}
	}
	return false, ""
}

func (h *V1Client) CreateAddonDeployment(client clientV1.ClientService, cluster *models.V1SpectroCluster, body *models.V1SpectroClusterProfiles) error {
	resolveNotification := false // during initial creation we never need to resolve packs.
	var params *clientV1.V1SpectroClustersPatchProfilesParams
	scope := cluster.Metadata.Annotations["scope"]
	switch scope {
	case "project":
		params = clientV1.NewV1SpectroClustersPatchProfilesParamsWithContext(h.Ctx)
	case "tenant":
		params = clientV1.NewV1SpectroClustersPatchProfilesParams()
	}
	params = params.WithUID(cluster.Metadata.UID).WithBody(body).WithResolveNotification(&resolveNotification)
	err := h.PatchWithRetry(client, params)
	return err
}

func (h *V1Client) PatchWithRetry(client clientV1.ClientService, params *clientV1.V1SpectroClustersPatchProfilesParams) error {
	var err error

	rand.NewSource(time.Now().UnixNano())
	for attempt := 0; attempt < h.RetryAttempts; attempt++ {
		// small jitter to prevent simultaneous retries
		s := rand.Intn(h.RetryAttempts) // n will be between 0 and number of retries
		log.Printf("Sleeping %d seconds, retry: %d, cluster:%s, profile:%s, ", s, attempt, params.UID, params.Body.Profiles[0].UID)
		time.Sleep(time.Duration(s) * time.Second)
		err = h.ClustersPatchProfiles(client, params)
		if err == nil {
			break
		}
	}
	return err
}

func (h *V1Client) ClustersPatchProfiles(client clientV1.ClientService, params *clientV1.V1SpectroClustersPatchProfilesParams) error {
	_, err := client.V1SpectroClustersPatchProfiles(params)
	return err
}

func (h *V1Client) DeleteAddonDeployment(client clientV1.ClientService, uid, scope string, body *models.V1SpectroClusterProfilesDeleteEntity) error {
	var params *clientV1.V1SpectroClustersDeleteProfilesParams
	switch scope {
	case "project":
		params = clientV1.NewV1SpectroClustersDeleteProfilesParamsWithContext(h.Ctx)
	case "tenant":
		params = clientV1.NewV1SpectroClustersDeleteProfilesParams()
	}
	params = params.WithUID(uid).WithBody(body)
	_, err := client.V1SpectroClustersDeleteProfiles(params)
	return err
}
