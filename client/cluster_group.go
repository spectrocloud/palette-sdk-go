package client

import (
	"errors"

	"github.com/spectrocloud/palette-api-go/apiutil/transport"
	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
)

func (h *V1Client) CreateClusterGroup(cluster *models.V1ClusterGroupEntity, scope string) (string, error) {
	var params *clientV1.V1ClusterGroupsCreateParams
	switch scope {
	case "project":
		params = clientV1.NewV1ClusterGroupsCreateParamsWithContext(h.Ctx).WithBody(cluster)
	case "tenant":
		params = clientV1.NewV1ClusterGroupsCreateParams().WithBody(cluster)
	default:
		return "", errors.New("invalid scope " + scope)
	}

	success, err := h.GetClient().V1ClusterGroupsCreate(params)
	if err != nil {
		return "", err
	}
	return *success.Payload.UID, nil
}

func (h *V1Client) DeleteClusterGroup(uid, scope string) error {
	var params *clientV1.V1ClusterGroupsUIDDeleteParams
	switch scope {
	case "project":
		params = clientV1.NewV1ClusterGroupsUIDDeleteParamsWithContext(h.Ctx).WithUID(uid)
	case "tenant":
		params = clientV1.NewV1ClusterGroupsUIDDeleteParams().WithUID(uid)
	default:
		return errors.New("invalid scope " + scope)
	}

	_, err := h.GetClient().V1ClusterGroupsUIDDelete(params)
	return err
}

func (h *V1Client) GetClusterGroup(uid, scope string) (*models.V1ClusterGroup, error) {
	group, err := h.GetClusterGroupWithoutStatus(uid, scope)
	if err != nil {
		return nil, err
	}
	if group != nil && group.Status != nil && !group.Status.IsActive {
		return nil, nil
	}
	return group, nil
}

func (h *V1Client) GetClusterGroupWithoutStatus(uid, scope string) (*models.V1ClusterGroup, error) {
	var params *clientV1.V1ClusterGroupsUIDGetParams
	switch scope {
	case "project":
		params = clientV1.NewV1ClusterGroupsUIDGetParamsWithContext(h.Ctx).WithUID(uid)
	case "tenant":
		params = clientV1.NewV1ClusterGroupsUIDGetParams().WithUID(uid)
	default:
		return nil, errors.New("invalid scope " + scope)
	}

	success, err := h.GetClient().V1ClusterGroupsUIDGet(params)
	var e *transport.TransportError
	if errors.As(err, &e) && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	cluster := success.Payload
	return cluster, nil
}

func (h *V1Client) GetClusterGroupByName(name, clusterGroupContext string) (*models.V1ObjectScopeEntity, error) {
	metadata, err := h.getClusterGroupMetadata(clusterGroupContext)
	if err != nil {
		return nil, err
	}

	for _, groupMeta := range metadata {
		if groupMeta.Name == name && groupMeta.Scope == clusterGroupContext {
			return groupMeta, nil
		}
	}
	return nil, nil
}

func (h *V1Client) GetClusterGroupByNameForProject(name, clusterGroupContext string) (*models.V1ClusterGroupSummary, error) {
	summaries, err := h.GetClusterGroupSummaries(clusterGroupContext)
	if err != nil {
		return nil, err
	}

	for _, groupSummary := range summaries {
		if groupSummary.Metadata.Name == name && groupSummary.Spec.Scope == clusterGroupContext {
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
	var params *clientV1.V1ClusterGroupsHostClusterSummaryParams
	switch clusterGroupContext {
	case "system":
		params = clientV1.NewV1ClusterGroupsHostClusterSummaryParams()
	case "project":
		fallthrough
	case "tenant":
		params = clientV1.NewV1ClusterGroupsHostClusterSummaryParamsWithContext(h.Ctx)
	default:
		return nil, errors.New("invalid scope " + clusterGroupContext)
	}

	resp, err := h.GetClient().V1ClusterGroupsHostClusterSummary(params)
	var e *transport.TransportError
	if errors.As(err, &e) && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return resp.Payload.Summaries, nil
}

// Update cluster group metadata by invoking V1ClusterGroupsUIDMetaUpdate hapi api
func (h *V1Client) UpdateClusterGroupMeta(clusterGroup *models.V1ClusterGroupEntity, scope string) error {
	var params *clientV1.V1ClusterGroupsUIDMetaUpdateParams
	switch scope {
	case "project":
		params = clientV1.NewV1ClusterGroupsUIDMetaUpdateParamsWithContext(h.Ctx).WithUID(clusterGroup.Metadata.UID)
	case "tenant":
		params = clientV1.NewV1ClusterGroupsUIDMetaUpdateParams().WithUID(clusterGroup.Metadata.UID)
	default:
		return errors.New("invalid scope " + scope)
	}

	params = params.WithBody(&models.V1ObjectMeta{
		Name:        clusterGroup.Metadata.Name,
		Labels:      clusterGroup.Metadata.Labels,
		Annotations: clusterGroup.Metadata.Annotations,
	})
	_, err := h.GetClient().V1ClusterGroupsUIDMetaUpdate(params)
	return err
}

// Update cluster group by invoking V1ClusterGroupsUIDHostClusterUpdate hapi api
func (h *V1Client) UpdateClusterGroup(uid string, clusterGroup *models.V1ClusterGroupHostClusterEntity, scope string) error {
	var params *clientV1.V1ClusterGroupsUIDHostClusterUpdateParams
	switch scope {
	case "project":
		params = clientV1.NewV1ClusterGroupsUIDHostClusterUpdateParamsWithContext(h.Ctx).WithUID(uid)
	case "tenant":
		params = clientV1.NewV1ClusterGroupsUIDHostClusterUpdateParams().WithUID(uid)
	default:
		return errors.New("invalid scope " + scope)
	}
	params = params.WithBody(clusterGroup)

	_, err := h.GetClient().V1ClusterGroupsUIDHostClusterUpdate(params)
	return err
}

func (h *V1Client) UpdateClusterProfileInClusterGroup(clusterGroupContext, clusterGroupUid string, clusterProfiles *models.V1SpectroClusterProfiles) error {
	var params *clientV1.V1ClusterGroupsUIDProfilesUpdateParams
	switch clusterGroupContext {
	case "project":
		params = clientV1.NewV1ClusterGroupsUIDProfilesUpdateParamsWithContext(h.Ctx).WithUID(clusterGroupUid)
	case "tenant":
		params = clientV1.NewV1ClusterGroupsUIDProfilesUpdateParams().WithUID(clusterGroupUid)
	default:
		return errors.New("invalid scope " + clusterGroupContext)
	}
	params = params.WithBody(clusterProfiles)

	_, err := h.GetClient().V1ClusterGroupsUIDProfilesUpdate(params)
	return err
}

func (h *V1Client) getClusterGroupMetadata(clusterGroupContext string) ([]*models.V1ObjectScopeEntity, error) {
	var params *clientV1.V1ClusterGroupsHostClusterMetadataParams
	switch clusterGroupContext {
	case "system":
		params = clientV1.NewV1ClusterGroupsHostClusterMetadataParams()
	case "project":
		fallthrough
	case "tenant":
		params = clientV1.NewV1ClusterGroupsHostClusterMetadataParamsWithContext(h.Ctx)
	default:
		return nil, errors.New("invalid scope " + clusterGroupContext)
	}

	resp, err := h.GetClient().V1ClusterGroupsHostClusterMetadata(params)
	var e *transport.TransportError
	if errors.As(err, &e) && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return resp.Payload.Items, nil
}
