package client

import (
	"log"
	"math/rand"
	"time"

	"github.com/spectrocloud/hapi/models"
	clusterC "github.com/spectrocloud/hapi/spectrocluster/client/v1"
)

func (h *V1Client) UpdateAddonDeployment(client clusterC.ClientService, cluster *models.V1SpectroCluster, body *models.V1SpectroClusterProfiles, newProfile *models.V1ClusterProfile) error {
	uid := cluster.Metadata.UID

	// check if profile id is the same - update, otherwise, delete and create
	is, replaceUID := IsProfileAttachedByName(cluster, newProfile)
	if is {
		body.Profiles[0].ReplaceWithProfile = replaceUID
	}

	resolveNotification := true
	var params *clusterC.V1SpectroClustersPatchProfilesParams
	scope := cluster.Metadata.Annotations["scope"]
	switch scope {
	case "project":
		params = clusterC.NewV1SpectroClustersPatchProfilesParamsWithContext(h.Ctx)
	case "tenant":
		params = clusterC.NewV1SpectroClustersPatchProfilesParams()
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

func (h *V1Client) CreateAddonDeployment(client clusterC.ClientService, cluster *models.V1SpectroCluster, body *models.V1SpectroClusterProfiles) error {
	resolveNotification := false // during initial creation we never need to resolve packs.
	var params *clusterC.V1SpectroClustersPatchProfilesParams
	scope := cluster.Metadata.Annotations["scope"]
	switch scope {
	case "project":
		params = clusterC.NewV1SpectroClustersPatchProfilesParamsWithContext(h.Ctx)
	case "tenant":
		params = clusterC.NewV1SpectroClustersPatchProfilesParams()
	}
	params = params.WithUID(cluster.Metadata.UID).WithBody(body).WithResolveNotification(&resolveNotification)
	err := h.PatchWithRetry(client, params)
	return err
}

func (h *V1Client) PatchWithRetry(client clusterC.ClientService, params *clusterC.V1SpectroClustersPatchProfilesParams) error {
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

func (h *V1Client) ClustersPatchProfiles(client clusterC.ClientService, params *clusterC.V1SpectroClustersPatchProfilesParams) error {
	_, err := client.V1SpectroClustersPatchProfiles(params)
	return err
}

func (h *V1Client) DeleteAddonDeployment(client clusterC.ClientService, uid, scope string, body *models.V1SpectroClusterProfilesDeleteEntity) error {
	var params *clusterC.V1SpectroClustersDeleteProfilesParams
	switch scope {
	case "project":
		params = clusterC.NewV1SpectroClustersDeleteProfilesParamsWithContext(h.Ctx)
	case "tenant":
		params = clusterC.NewV1SpectroClustersDeleteProfilesParams()
	}
	params = params.WithUID(uid).WithBody(body)
	_, err := client.V1SpectroClustersDeleteProfiles(params)
	return err
}
