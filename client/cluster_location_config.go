package client

import (
	"errors"

	"github.com/spectrocloud/hapi/models"
	clusterC "github.com/spectrocloud/hapi/spectrocluster/client/v1"
)

func (h *V1Client) GetClusterLocationConfig(scope, uid string) (*models.V1ClusterLocation, error) {
	if clusterStatus, err := h.GetClusterWithoutStatus(scope, uid); err != nil {
		return nil, err
	} else if clusterStatus != nil && clusterStatus.Status != nil && clusterStatus.Status.Location != nil {
		return clusterStatus.Status.Location, nil
	}

	return nil, errors.New("failed to read cluster location")
}

func (h *V1Client) UpdateClusterLocationConfig(uid string, clusterContext string, config *models.V1SpectroClusterLocationInputEntity) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return err
	}
	var params *clusterC.V1SpectroClustersUIDLocationPutParams
	switch clusterContext {
	case "project":
		params = clusterC.NewV1SpectroClustersUIDLocationPutParamsWithContext(h.Ctx).WithUID(uid).WithBody(config)
	case "tenant":
		params = clusterC.NewV1SpectroClustersUIDLocationPutParams().WithUID(uid).WithBody(config)
	}
	_, err = client.V1SpectroClustersUIDLocationPut(params)
	return err
}

func (h *V1Client) ApplyClusterLocationConfig(scope, uid string, config *models.V1SpectroClusterLocationInputEntity) error {
	if curentConfig, err := h.GetClusterLocationConfig(scope, uid); err != nil {
		return err
	} else if curentConfig == nil {
		return h.UpdateClusterLocationConfig(uid, scope, config)
	} else {
		return h.UpdateClusterLocationConfig(uid, scope, config)
	}
}
