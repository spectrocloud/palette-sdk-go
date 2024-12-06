package client

import (
	clientv1 "github.com/spectrocloud/palette-sdk-go/api/client/v1"
	"github.com/spectrocloud/palette-sdk-go/api/models"
)

// EnableClusterHybridSettings enables hybrid mode on a cluster.
func (h *V1Client) EnableClusterHybridSettings(uid string) error {
	params := clientv1.NewV1SpectroClustersUIDHybridSettingsParamsWithContext(h.ctx).
		WithBody(&models.V1ClusterHybridSettingsEntity{
			Enable: true,
		}).WithUID(uid)

	_, err := h.Client.V1SpectroClustersUIDHybridSettings(params)
	return err
}

// PutAwsHybridConfig updates the AWS hybrid config for a cluster.
func (h *V1Client) PutAwsHybridConfig(configUID string, hybridConfig *models.V1AwsCloudHybridConfigEntity) error {
	params := clientv1.NewV1AwsCloudConfigsUIDHybridConfigParamsWithContext(h.ctx).
		WithBody(hybridConfig).
		WithConfigUID(configUID)

	_, err := h.Client.V1AwsCloudConfigsUIDHybridConfig(params)
	return err
}

// CreateAwsHybridMachinePool creates a Hybrid AWS cloud config's Edge-Native machine pool
func (h *V1Client) CreateAwsHybridMachinePool(configUID string, machinePoolConfig *models.V1HybridEdgeNativeMachinePoolConfigEntity) error {
	params := clientv1.NewV1AwsCloudConfigsEdgeNativeUIDMachinePoolCreateParamsWithContext(h.ctx).
		WithBody(machinePoolConfig).
		WithConfigUID(configUID)

	_, err := h.Client.V1AwsCloudConfigsEdgeNativeUIDMachinePoolCreate(params)
	return err
}

// UpdateAwsHybridMachinePool updates the specified Hybrid AWS cluster cloud config's Edge-Native machine pool
func (h *V1Client) UpdateAwsHybridMachinePool(configUID, machinePoolName string, machinePoolConfig *models.V1HybridEdgeNativeMachinePoolConfigUpdateEntity) error {
	params := clientv1.NewV1AwsCloudConfigsEdgeNativeMachinePoolUpdateParamsWithContext(h.ctx).
		WithBody(machinePoolConfig).
		WithMachinePoolName(machinePoolName).
		WithConfigUID(configUID)

	_, err := h.Client.V1AwsCloudConfigsEdgeNativeMachinePoolUpdate(params)
	return err
}

// DeleteAwsHybridMachinePool deletes the specified Edge-Native machine pool of hybrid AWS cluster
func (h *V1Client) DeleteAwsHybridMachinePool(configUID, machinePoolName string) error {
	params := clientv1.NewV1AwsCloudConfigsEdgeNativeMachinePoolDeleteParamsWithContext(h.ctx).
		WithMachinePoolName(machinePoolName).
		WithConfigUID(configUID)

	_, err := h.Client.V1AwsCloudConfigsEdgeNativeMachinePoolDelete(params)
	return err
}
