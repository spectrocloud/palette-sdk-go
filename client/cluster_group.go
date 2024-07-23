package client

import (
	clientv1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
	"github.com/spectrocloud/palette-sdk-go/client/apiutil"
)

// CreateClusterGroup creates a new cluster group.
func (h *V1Client) CreateClusterGroup(cluster *models.V1ClusterGroupEntity) (string, error) {
	params := clientv1.NewV1ClusterGroupsCreateParamsWithContext(h.ctx).
		WithBody(cluster)
	resp, err := h.Client.V1ClusterGroupsCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

// DeleteClusterGroup deletes an existing cluster group.
func (h *V1Client) DeleteClusterGroup(uid string) error {
	params := clientv1.NewV1ClusterGroupsUIDDeleteParamsWithContext(h.ctx).
		WithUID(uid)
	_, err := h.Client.V1ClusterGroupsUIDDelete(params)
	return err
}

// GetClusterGroup retrieves an existing cluster group whose status is active.
func (h *V1Client) GetClusterGroup(uid string) (*models.V1ClusterGroup, error) {
	group, err := h.GetClusterGroupWithoutStatus(uid)
	if err != nil {
		return nil, err
	}
	if group != nil && group.Status != nil && !group.Status.IsActive {
		return nil, nil
	}
	return group, nil
}

// GetClusterGroupWithoutStatus retrieves an existing cluster group, regardless of status.
func (h *V1Client) GetClusterGroupWithoutStatus(uid string) (*models.V1ClusterGroup, error) {
	params := clientv1.NewV1ClusterGroupsUIDGetParamsWithContext(h.ctx).
		WithUID(uid)
	resp, err := h.Client.V1ClusterGroupsUIDGet(params)
	if apiutil.Is404(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// GetClusterGroupSummaryByName retrieves a summary for an existing cluster group by name.
func (h *V1Client) GetClusterGroupSummaryByName(name string) (*models.V1ClusterGroupSummary, error) {
	summaries, err := h.GetClusterGroupSummaries()
	if err != nil {
		return nil, err
	}
	for _, groupSummary := range summaries {
		if groupSummary.Metadata.Name == name {
			return groupSummary, nil
		}
	}
	return nil, nil
}

// GetClusterGroupScopeMetadata retrieves scope metadata for all cluster groups.
func (h *V1Client) GetClusterGroupScopeMetadata() ([]*models.V1ObjectScopeEntity, error) {
	metadata, err := h.clusterGroupMetadata()
	if err != nil {
		return nil, err
	}
	return metadata, nil
}

// GetClusterGroupScopeMetadataByName retrieves scope metadata for an existing cluster group by name.
func (h *V1Client) GetClusterGroupScopeMetadataByName(name string) (*models.V1ObjectScopeEntity, error) {
	metadata, err := h.clusterGroupMetadata()
	if err != nil {
		return nil, err
	}
	for _, groupMeta := range metadata {
		if groupMeta.Name == name {
			return groupMeta, nil
		}
	}
	return nil, nil
}

// GetClusterGroupSummaries retrieves summaries for all cluster groups.
func (h *V1Client) GetClusterGroupSummaries() ([]*models.V1ClusterGroupSummary, error) {
	params := clientv1.NewV1ClusterGroupsHostClusterSummaryParamsWithContext(h.ctx)
	resp, err := h.Client.V1ClusterGroupsHostClusterSummary(params)
	if apiutil.Is404(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return resp.Payload.Summaries, nil
}

// UpdateClusterGroupMeta updates an existing cluster group's metadata.
func (h *V1Client) UpdateClusterGroupMeta(clusterGroup *models.V1ClusterGroupEntity) error {
	params := clientv1.NewV1ClusterGroupsUIDMetaUpdateParamsWithContext(h.ctx).
		WithUID(clusterGroup.Metadata.UID).
		WithBody(&models.V1ObjectMeta{
			Name:        clusterGroup.Metadata.Name,
			Labels:      clusterGroup.Metadata.Labels,
			Annotations: clusterGroup.Metadata.Annotations,
		})
	_, err := h.Client.V1ClusterGroupsUIDMetaUpdate(params)
	return err
}

// UpdateClusterGroup updates an existing cluster group.
func (h *V1Client) UpdateClusterGroup(uid string, clusterGroup *models.V1ClusterGroupHostClusterEntity) error {
	params := clientv1.NewV1ClusterGroupsUIDHostClusterUpdateParamsWithContext(h.ctx).
		WithUID(uid).
		WithBody(clusterGroup)
	_, err := h.Client.V1ClusterGroupsUIDHostClusterUpdate(params)
	return err
}

// UpdateClusterProfileInClusterGroup updates an existing cluster group's default cluster profiles.
func (h *V1Client) UpdateClusterProfileInClusterGroup(clusterGroupUID string, clusterProfiles *models.V1SpectroClusterProfiles) error {
	params := clientv1.NewV1ClusterGroupsUIDProfilesUpdateParamsWithContext(h.ctx).
		WithUID(clusterGroupUID).
		WithBody(clusterProfiles)
	_, err := h.Client.V1ClusterGroupsUIDProfilesUpdate(params)
	return err
}

func (h *V1Client) clusterGroupMetadata() ([]*models.V1ObjectScopeEntity, error) {
	params := clientv1.NewV1ClusterGroupsHostClusterMetadataParamsWithContext(h.ctx)
	resp, err := h.Client.V1ClusterGroupsHostClusterMetadata(params)
	if apiutil.Is404(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return resp.Payload.Items, nil
}
