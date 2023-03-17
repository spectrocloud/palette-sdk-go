package client

import (
	"errors"

	"github.com/spectrocloud/hapi/apiutil/transport"
	"github.com/spectrocloud/hapi/models"
	clusterC "github.com/spectrocloud/hapi/spectrocluster/client/v1"
)

func (h *V1Client) CreateClusterGroup(cluster *models.V1ClusterGroupEntity, scope string) (string, error) {
	if h.CreateClusterGroupFn != nil {
		return h.CreateClusterGroupFn(cluster)
	}

	client, err := h.GetClusterClient()
	if err != nil {
		return "", err
	}

	var params *clusterC.V1ClusterGroupsCreateParams
	switch scope {
	case "project":
		params = clusterC.NewV1ClusterGroupsCreateParamsWithContext(h.Ctx).WithBody(cluster)
	case "tenant":
		params = clusterC.NewV1ClusterGroupsCreateParams().WithBody(cluster)
	default:
		return "", errors.New("invalid scope")
	}
	success, err := client.V1ClusterGroupsCreate(params)
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

func (h *V1Client) DeleteClusterGroup(uid string, scope string) error {
	if h.DeleteClusterGroupFn != nil {
		return h.DeleteClusterGroupFn(uid)
	}
	client, err := h.GetClusterClient()
	if err != nil {
		return err
	}
	var params *clusterC.V1ClusterGroupsUIDDeleteParams
	switch scope {
	case "project":
		params = clusterC.NewV1ClusterGroupsUIDDeleteParamsWithContext(h.Ctx).WithUID(uid)
	case "tenant":
		params = clusterC.NewV1ClusterGroupsUIDDeleteParams().WithUID(uid)
	default:
		return errors.New("invalid scope")
	}
	_, err = client.V1ClusterGroupsUIDDelete(params)
	return err
}

func (h *V1Client) GetClusterGroup(uid string, scope string) (*models.V1ClusterGroup, error) {
	if h.GetClusterGroupFn != nil {
		return h.GetClusterGroupFn(uid)
	}
	group, err := h.GetClusterGroupWithoutStatus(uid, scope)
	if err != nil {
		return nil, err
	}

	if group == nil && group.Status.IsActive {
		return nil, nil
	}

	return group, nil
}

func (h *V1Client) GetClusterGroupWithoutStatus(uid string, scope string) (*models.V1ClusterGroup, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}

	var params *clusterC.V1ClusterGroupsUIDGetParams
	switch scope {
	case "project":
		params = clusterC.NewV1ClusterGroupsUIDGetParamsWithContext(h.Ctx).WithUID(uid)
	case "tenant":
		params = clusterC.NewV1ClusterGroupsUIDGetParams().WithUID(uid)
	default:
		return nil, errors.New("invalid scope")
	}
	success, err := client.V1ClusterGroupsUIDGet(params)
	if e, ok := err.(*transport.TransportError); ok && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	cluster := success.Payload
	return cluster, nil
}

func (h *V1Client) GetClusterGroupByName(name string, clusterGroupContext string) (*models.V1ObjectScopeEntity, error) {
	metadata, err := h.getClusterGroupMetadata(clusterGroupContext)
	if err != nil {
		return nil, err
	}

	for _, groupMeta := range metadata {
		if groupMeta.Name == name && groupMeta.Scope == clusterGroupContext { // tenant or system. keep it to extend to project in future.
			return groupMeta, nil
		}
	}

	return nil, nil
}

func (h *V1Client) GetClusterGroupByNameForProject(name string, clusterGroupContext string) (*models.V1ClusterGroupSummary, error) {
	summaries, err := h.GetClusterGroupSummaries(clusterGroupContext)
	if err != nil {
		return nil, err
	}

	for _, groupSummary := range summaries {
		if groupSummary.Metadata.Name == name && groupSummary.Spec.Scope == clusterGroupContext { // tenant or system. keep it to extend to project in future.
			return groupSummary, nil
		}
	}

	return nil, nil
}

func (h *V1Client) GetClusterGroupMetadata(clusterGroupContext string) ([]*models.V1ObjectScopeEntity, error) {
	metadata, err := h.getClusterGroupMetadata(clusterGroupContext)
	if err != nil {
		return nil, err
	}
	return metadata, nil
}

func (h *V1Client) GetClusterGroupSummaries(clusterGroupContext string) ([]*models.V1ClusterGroupSummary, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}

	var params *clusterC.V1ClusterGroupsHostClusterSummaryParams
	switch clusterGroupContext {
	case "system":
		params = clusterC.NewV1ClusterGroupsHostClusterSummaryParams()
	case "tenant":
		params = clusterC.NewV1ClusterGroupsHostClusterSummaryParamsWithContext(h.Ctx)
	default:
		return nil, errors.New("invalid scope")
	}

	resp, err := client.V1ClusterGroupsHostClusterSummary(params)
	if e, ok := err.(*transport.TransportError); ok && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return resp.Payload.Summaries, nil
}

// Update cluster group metadata by invoking V1ClusterGroupsUIDMetaUpdate hapi api
func (h *V1Client) UpdateClusterGroupMeta(clusterGroup *models.V1ClusterGroupEntity, scope string) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return err
	}
	var params *clusterC.V1ClusterGroupsUIDMetaUpdateParams
	switch scope {
	case "project":
		params = clusterC.NewV1ClusterGroupsUIDMetaUpdateParamsWithContext(h.Ctx).WithUID(clusterGroup.Metadata.UID)
	case "tenant":
		params = clusterC.NewV1ClusterGroupsUIDMetaUpdateParams().WithUID(clusterGroup.Metadata.UID)
	default:
		return errors.New("invalid scope")
	}
	params = params.WithBody(&models.V1ObjectMeta{
		Name:        clusterGroup.Metadata.Name,
		Labels:      clusterGroup.Metadata.Labels,
		Annotations: clusterGroup.Metadata.Annotations,
	})
	_, err = client.V1ClusterGroupsUIDMetaUpdate(params)
	return err
}

// Update cluster group by invoking V1ClusterGroupsUIDHostClusterUpdate hapi api
func (h *V1Client) UpdateClusterGroup(uid string, clusterGroup *models.V1ClusterGroupHostClusterEntity, scope string) error {
	if h.UpdateClusterGroupFn != nil {
		return h.UpdateClusterGroupFn(uid, clusterGroup)
	}
	client, err := h.GetClusterClient()
	if err != nil {
		return err
	}
	var params *clusterC.V1ClusterGroupsUIDHostClusterUpdateParams
	switch scope {
	case "project":
		params = clusterC.NewV1ClusterGroupsUIDHostClusterUpdateParamsWithContext(h.Ctx).WithUID(uid)
	case "tenant":
		params = clusterC.NewV1ClusterGroupsUIDHostClusterUpdateParams().WithUID(uid)
	default:
		return errors.New("invalid scope")
	}
	params = params.WithBody(clusterGroup)
	_, err = client.V1ClusterGroupsUIDHostClusterUpdate(params)
	return err
}

func (h *V1Client) getClusterGroupMetadata(clusterGroupContext string) ([]*models.V1ObjectScopeEntity, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}

	var params *clusterC.V1ClusterGroupsHostClusterMetadataParams
	switch clusterGroupContext {
	case "system":
		params = clusterC.NewV1ClusterGroupsHostClusterMetadataParams()
	case "tenant":
		params = clusterC.NewV1ClusterGroupsHostClusterMetadataParamsWithContext(h.Ctx)
	default:
		return nil, errors.New("invalid scope")
	}

	resp, err := client.V1ClusterGroupsHostClusterMetadata(params)
	if e, ok := err.(*transport.TransportError); ok && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return resp.Payload.Items, nil
}
