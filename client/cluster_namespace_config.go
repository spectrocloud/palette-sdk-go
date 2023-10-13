package client

import (
	"github.com/spectrocloud/hapi/models"
	clusterC "github.com/spectrocloud/hapi/spectrocluster/client/v1"
	"github.com/spectrocloud/palette-sdk-go/client/herr"
)

func (h *V1Client) GetClusterNamespaceConfig(uid string, clusterContext string) (*models.V1ClusterNamespaceResources, error) {
	if h.GetClusterNamespaceConfigFn != nil {
		return h.GetClusterNamespaceConfigFn(uid)
	}
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}
	var params *clusterC.V1SpectroClustersUIDConfigNamespacesGetParams
	switch clusterContext {
	case "project":
		params = clusterC.NewV1SpectroClustersUIDConfigNamespacesGetParamsWithContext(h.Ctx).WithUID(uid)
	case "tenant":
		params = clusterC.NewV1SpectroClustersUIDConfigNamespacesGetParams().WithUID(uid)
	}
	success, err := client.V1SpectroClustersUIDConfigNamespacesGet(params)
	if err != nil {
		if herr.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}

	return success.Payload, nil
}

// UpdateClusterNamespaceConfig no create for namespaces, there is only update.
func (h *V1Client) UpdateClusterNamespaceConfig(uid string, clusterContext string, config *models.V1ClusterNamespaceResourcesUpdateEntity) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return err
	}
	var params *clusterC.V1SpectroClustersUIDConfigNamespacesUpdateParams
	switch clusterContext {
	case "project":
		params = clusterC.NewV1SpectroClustersUIDConfigNamespacesUpdateParamsWithContext(h.Ctx).WithUID(uid).WithBody(config)
	case "tenant":
		params = clusterC.NewV1SpectroClustersUIDConfigNamespacesUpdateParams().WithUID(uid).WithBody(config)
	}
	_, err = client.V1SpectroClustersUIDConfigNamespacesUpdate(params)
	return err
}

func (h *V1Client) ApplyClusterNamespaceConfig(uid string, clusterContext string, config []*models.V1ClusterNamespaceResourceInputEntity) error {
	if _, err := h.GetClusterNamespaceConfig(uid, clusterContext); err != nil {
		return err
	} else {
		return h.UpdateClusterNamespaceConfig(uid, clusterContext, toUpdateNamespace(config)) // update method is same as create
	}
}

func toUpdateNamespace(config []*models.V1ClusterNamespaceResourceInputEntity) *models.V1ClusterNamespaceResourcesUpdateEntity {
	return &models.V1ClusterNamespaceResourcesUpdateEntity{
		Namespaces: config,
	}
}
