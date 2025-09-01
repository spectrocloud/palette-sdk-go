package client

import (
	clientv1 "github.com/spectrocloud/palette-sdk-go/api/client/version1"
	"github.com/spectrocloud/palette-sdk-go/api/models"
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
		WithContext(ContextForScope(h.baseCtx, cluster.Metadata.Annotations[Scope], h.projectUID)).
		WithUID(uid).
		WithBody(body).
		WithResolveNotification(&resolveNotifications)

	return h.ClustersPatchProfiles(params)
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
		WithContext(ContextForScope(h.baseCtx, cluster.Metadata.Annotations[Scope], h.projectUID)).
		WithUID(cluster.Metadata.UID).
		WithBody(body).
		WithResolveNotification(&resolveNotifications)

	return h.ClustersPatchProfiles(params)
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
