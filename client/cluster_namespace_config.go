package client

import (
	clientv1 "github.com/spectrocloud/palette-sdk-go/api/client/v1"
	"github.com/spectrocloud/palette-sdk-go/api/models"
	"github.com/spectrocloud/palette-sdk-go/client/herr"
)

// GetClusterNamespaceConfig gets the namespaced resource configuration for a cluster.
func (h *V1Client) GetClusterNamespaceConfig(uid string) (*models.V1ClusterNamespaceResources, error) {
	params := clientv1.NewV1SpectroClustersUIDConfigNamespacesGetParamsWithContext(h.ctx).
		WithUID(uid)
	resp, err := h.Client.V1SpectroClustersUIDConfigNamespacesGet(params)
	if err != nil {
		if herr.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}
	return resp.Payload, nil
}

// UpdateClusterNamespaceConfig updates an existing namespaced resource configuration for a cluster.
// Create is not supported for namespaced resources.
func (h *V1Client) UpdateClusterNamespaceConfig(uid string, config *models.V1ClusterNamespaceResourcesUpdateEntity) error {
	params := clientv1.NewV1SpectroClustersUIDConfigNamespacesUpdateParamsWithContext(h.ctx).
		WithUID(uid).
		WithBody(config)
	_, err := h.Client.V1SpectroClustersUIDConfigNamespacesUpdate(params)
	return err
}

// ApplyClusterNamespaceConfig creates a new namespaced resource configuration for a cluster or updates its namespaced resource configuration if one exists.
func (h *V1Client) ApplyClusterNamespaceConfig(uid string, config []*models.V1ClusterNamespaceResourceInputEntity) error {
	_, err := h.GetClusterNamespaceConfig(uid)
	if err != nil {
		return err
	}
	return h.UpdateClusterNamespaceConfig(uid, toUpdateNamespace(config)) // update method is same as create
}

func toUpdateNamespace(config []*models.V1ClusterNamespaceResourceInputEntity) *models.V1ClusterNamespaceResourcesUpdateEntity {
	return &models.V1ClusterNamespaceResourcesUpdateEntity{
		Namespaces: config,
	}
}
