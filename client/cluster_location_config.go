package client

import (
	"errors"

	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
)

func (h *V1Client) GetClusterLocationConfig(scope, uid string) (*models.V1ClusterLocation, error) {
	if clusterStatus, err := h.GetClusterWithoutStatus(scope, uid); err != nil {
		return nil, err
	} else if clusterStatus != nil && clusterStatus.Status != nil && clusterStatus.Status.Location != nil {
		return clusterStatus.Status.Location, nil
	}

	return nil, errors.New("failed to read cluster location")
}

func (h *V1Client) UpdateClusterLocationConfig(uid, scope string, config *models.V1SpectroClusterLocationInputEntity) error {
	var params *clientV1.V1SpectroClustersUIDLocationPutParams
	switch scope {
	case "project":
		params = clientV1.NewV1SpectroClustersUIDLocationPutParamsWithContext(h.Ctx).WithUID(uid).WithBody(config)
	case "tenant":
		params = clientV1.NewV1SpectroClustersUIDLocationPutParams().WithUID(uid).WithBody(config)
	}

	_, err := h.Client.V1SpectroClustersUIDLocationPut(params)
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
