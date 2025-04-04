package client

import (
	clientv1 "github.com/spectrocloud/palette-sdk-go/api/client/version1"
	"github.com/spectrocloud/palette-sdk-go/api/models"
)

// CreateClusterCustomCloud creates a custom cloud cluster.
func (h *V1Client) CreateClusterCustomCloud(cluster *models.V1SpectroCustomClusterEntity, cloudType string) (string, error) {
	params := clientv1.NewV1SpectroClustersCustomCreateParamsWithContext(h.ctx).
		WithCloudType(cloudType).
		WithBody(cluster)
	resp, err := h.Client.V1SpectroClustersCustomCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

// GetCloudConfigCustomCloud retrieves an existing custom cloud cluster's cloud config.
func (h *V1Client) GetCloudConfigCustomCloud(configUID, cloudType string) (*models.V1CustomCloudConfig, error) {
	params := clientv1.NewV1CloudConfigsCustomGetParamsWithContext(h.ctx).
		WithCloudType(cloudType).
		WithConfigUID(configUID)
	resp, err := h.Client.V1CloudConfigsCustomGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// UpdateCloudConfigCustomCloud updates an existing custom cloud cluster's cloud config.
func (h *V1Client) UpdateCloudConfigCustomCloud(updatedConfig *models.V1CustomCloudClusterConfigEntity, configUID, cloudType string) error {
	params := clientv1.NewV1CloudConfigsCustomUIDClusterConfigParamsWithContext(h.ctx).
		WithCloudType(cloudType).
		WithBody(updatedConfig).
		WithConfigUID(configUID)
	_, err := h.Client.V1CloudConfigsCustomUIDClusterConfig(params)
	return err
}

// CreateMachinePoolCustomCloud creates a new custom cloud machine pool.
func (h *V1Client) CreateMachinePoolCustomCloud(mpEntity *models.V1CustomMachinePoolConfigEntity, configUID, cloudType string) error {
	params := clientv1.NewV1CloudConfigsCustomMachinePoolCreateParamsWithContext(h.ctx).
		WithCloudType(cloudType).
		WithBody(mpEntity).
		WithConfigUID(configUID)
	_, err := h.Client.V1CloudConfigsCustomMachinePoolCreate(params)
	return err
}

// UpdateMachinePoolCustomCloud updates an existing custom cloud machine pool.
func (h *V1Client) UpdateMachinePoolCustomCloud(mpEntity *models.V1CustomMachinePoolConfigEntity, machinePoolName, configUID, cloudType string) error {
	params := clientv1.NewV1CloudConfigsCustomMachinePoolUpdateParamsWithContext(h.ctx).
		WithCloudType(cloudType).
		WithBody(mpEntity).
		WithConfigUID(configUID).
		WithMachinePoolName(machinePoolName)
	_, err := h.Client.V1CloudConfigsCustomMachinePoolUpdate(params)
	return err
}

// DeleteMachinePoolCustomCloud deletes an existing custom cloud machine pool.
func (h *V1Client) DeleteMachinePoolCustomCloud(mpName string, configUID, cloudType string) error {
	params := clientv1.NewV1CloudConfigsCustomMachinePoolDeleteParamsWithContext(h.ctx).
		WithCloudType(cloudType).
		WithConfigUID(configUID).
		WithMachinePoolName(mpName)
	_, err := h.Client.V1CloudConfigsCustomMachinePoolDelete(params)
	return err
}
