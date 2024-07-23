package client

import (
	"log"
	"math/rand"
	"time"

	clientv1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
)

// UpdateAddonDeployment updates the addon deployment for a cluster.
func (h *V1Client) UpdateAddonDeployment(cluster *models.V1SpectroCluster, body *models.V1SpectroClusterProfiles, newProfile *models.V1ClusterProfile) error {
	uid := cluster.Metadata.UID

	// check if profile id is the same - update, otherwise, delete and create
	is, replaceUID := IsProfileAttachedByName(cluster, newProfile)
	if is {
		body.Profiles[0].ReplaceWithProfile = replaceUID
	}

	resolveNotifications := true
	params := clientv1.NewV1SpectroClustersPatchProfilesParams().
		WithContext(ContextForScope(cluster.Metadata.Annotations[Scope], h.projectUID)).
		WithUID(uid).
		WithBody(body).
		WithResolveNotification(&resolveNotifications)

	return h.PatchWithRetry(params)
}

// IsProfileAttachedByName checks if a profile is already attached to a cluster.
func IsProfileAttachedByName(cluster *models.V1SpectroCluster, newProfile *models.V1ClusterProfile) (bool, string) {
	for _, profile := range cluster.Spec.ClusterProfileTemplates {
		if profile.Name == newProfile.Metadata.Name {
			return true, profile.UID
		}
	}
	return false, ""
}

// CreateAddonDeployment creates an addon deployment for a cluster.
func (h *V1Client) CreateAddonDeployment(cluster *models.V1SpectroCluster, body *models.V1SpectroClusterProfiles) error {
	resolveNotifications := false // during initial creation we never need to resolve packs
	params := clientv1.NewV1SpectroClustersPatchProfilesParams().
		WithContext(ContextForScope(cluster.Metadata.Annotations[Scope], h.projectUID)).
		WithUID(cluster.Metadata.UID).
		WithBody(body).
		WithResolveNotification(&resolveNotifications)

	return h.PatchWithRetry(params)
}

// PatchWithRetry patches cluster's profiles with with retry logic.
func (h *V1Client) PatchWithRetry(params *clientv1.V1SpectroClustersPatchProfilesParams) error {
	var err error
	rand.NewSource(time.Now().UnixNano())
	for attempt := 0; attempt < h.retryAttempts; attempt++ {
		// small jitter to prevent simultaneous retries. n will be between 0 and number of retries.
		s := rand.Intn(h.retryAttempts) // #nosec G404 - random number is not used for security purposes
		log.Printf("Sleeping %d seconds, retry: %d, cluster:%s, profile:%s, ", s, attempt, params.UID, params.Body.Profiles[0].UID)
		time.Sleep(time.Duration(s) * time.Second)
		err = h.ClustersPatchProfiles(params)
		if err == nil {
			break
		}
	}
	return err
}

// ClustersPatchProfiles patches a cluster's profiles.
func (h *V1Client) ClustersPatchProfiles(params *clientv1.V1SpectroClustersPatchProfilesParams) error {
	_, err := h.Client.V1SpectroClustersPatchProfiles(params)
	return err
}

// DeleteAddonDeployment deletes an addon deployment for a cluster.
func (h *V1Client) DeleteAddonDeployment(uid string, body *models.V1SpectroClusterProfilesDeleteEntity) error {
	params := clientv1.NewV1SpectroClustersDeleteProfilesParamsWithContext(h.ctx).
		WithUID(uid).
		WithBody(body)
	_, err := h.Client.V1SpectroClustersDeleteProfiles(params)
	return err
}
