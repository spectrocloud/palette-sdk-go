package client

import (
	clientv1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
)

// GetClusterLocationConfig returns a cluster's location config.
func (h *V1Client) GetClusterLocationConfig(uid string) (*models.V1ClusterLocation, error) {
	clusterStatus, err := h.GetClusterWithoutStatus(uid)
	if err != nil {
		return nil, err
	}
	return clusterStatus.Status.Location, nil
}

// UpdateClusterLocationConfig updates a cluster's location config.
func (h *V1Client) UpdateClusterLocationConfig(uid string, config *models.V1SpectroClusterLocationInputEntity) error {
	params := clientv1.NewV1SpectroClustersUIDLocationPutParamsWithContext(h.ctx).
		WithUID(uid).
		WithBody(config)
	_, err := h.Client.V1SpectroClustersUIDLocationPut(params)
	return err
}

// ApplyClusterLocationConfig creates a new cluster location config or updates an existing one.
func (h *V1Client) ApplyClusterLocationConfig(uid string, config *models.V1SpectroClusterLocationInputEntity) error {
	_, err := h.GetClusterLocationConfig(uid)
	if err != nil {
		return err
	}
	return h.UpdateClusterLocationConfig(uid, config) // update method is same as create
}
