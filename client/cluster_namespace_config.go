package client

import (
	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
	"github.com/spectrocloud/palette-sdk-go/client/herr"
)

func (h *V1Client) GetClusterNamespaceConfig(uid, scope string) (*models.V1ClusterNamespaceResources, error) {
	var params *clientV1.V1SpectroClustersUIDConfigNamespacesGetParams
	switch scope {
	case "project":
		params = clientV1.NewV1SpectroClustersUIDConfigNamespacesGetParamsWithContext(h.Ctx).WithUID(uid)
	case "tenant":
		params = clientV1.NewV1SpectroClustersUIDConfigNamespacesGetParams().WithUID(uid)
	}

	success, err := h.GetClient().V1SpectroClustersUIDConfigNamespacesGet(params)
	if err != nil {
		if herr.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}

	return success.Payload, nil
}

// UpdateClusterNamespaceConfig no create for namespaces, there is only update.
func (h *V1Client) UpdateClusterNamespaceConfig(uid, scope string, config *models.V1ClusterNamespaceResourcesUpdateEntity) error {
	var params *clientV1.V1SpectroClustersUIDConfigNamespacesUpdateParams
	switch scope {
	case "project":
		params = clientV1.NewV1SpectroClustersUIDConfigNamespacesUpdateParamsWithContext(h.Ctx).WithUID(uid).WithBody(config)
	case "tenant":
		params = clientV1.NewV1SpectroClustersUIDConfigNamespacesUpdateParams().WithUID(uid).WithBody(config)
	}

	_, err := h.GetClient().V1SpectroClustersUIDConfigNamespacesUpdate(params)
	return err
}

func (h *V1Client) ApplyClusterNamespaceConfig(uid, scope string, config []*models.V1ClusterNamespaceResourceInputEntity) error {
	if _, err := h.GetClusterNamespaceConfig(uid, scope); err != nil {
		return err
	} else {
		return h.UpdateClusterNamespaceConfig(uid, scope, toUpdateNamespace(config)) // update method is same as create
	}
}

func toUpdateNamespace(config []*models.V1ClusterNamespaceResourceInputEntity) *models.V1ClusterNamespaceResourcesUpdateEntity {
	return &models.V1ClusterNamespaceResourcesUpdateEntity{
		Namespaces: config,
	}
}
