package client

import (
	clientv1 "github.com/spectrocloud/palette-sdk-go/api/client/v1"
	"github.com/spectrocloud/palette-sdk-go/api/models"
)

// CreateClusterGke creates a GKE cluster.
func (h *V1Client) CreateClusterGke(cluster *models.V1SpectroGcpClusterEntity) (string, error) {
	params := clientv1.NewV1SpectroClustersGkeCreateParamsWithContext(h.ctx).WithBody(cluster)
	success, err := h.Client.V1SpectroClustersGkeCreate(params)
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

// GetCloudConfigGke retrieves an existing GKE cluster's cloud config.
func (h *V1Client) GetCloudConfigGke(configUID string) (*models.V1GcpCloudConfig, error) {
	params := clientv1.NewV1CloudConfigsGkeGetParamsWithContext(h.ctx).WithConfigUID(configUID)

	success, err := h.Client.V1CloudConfigsGkeGet(params)
	if err != nil {
		return nil, err
	}

	return success.Payload, nil
}

// CreateMachinePoolGke creates a new GKE machine pool.
func (h *V1Client) CreateMachinePoolGke(cloudConfigID string, machinePool *models.V1GcpMachinePoolConfigEntity) error {
	params := clientv1.NewV1CloudConfigsGkeMachinePoolCreateParamsWithContext(h.ctx).WithConfigUID(cloudConfigID).WithBody(machinePool)
	_, err := h.Client.V1CloudConfigsGkeMachinePoolCreate(params)
	return err
}

// UpdateMachinePoolGke updates an existing GKE machine pool.
func (h *V1Client) UpdateMachinePoolGke(cloudConfigID string, machinePool *models.V1GcpMachinePoolConfigEntity) error {
	params := clientv1.NewV1CloudConfigsGkeMachinePoolUpdateParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigID).
		WithMachinePoolName(*machinePool.PoolConfig.Name).
		WithBody(machinePool)

	_, err := h.Client.V1CloudConfigsGkeMachinePoolUpdate(params)
	return err
}

// DeleteMachinePoolGke deletes an existing GKE machine pool.
func (h *V1Client) DeleteMachinePoolGke(cloudConfigID string, machinePoolName string) error {
	params := clientv1.NewV1CloudConfigsGkeMachinePoolDeleteParamsWithContext(h.ctx).WithConfigUID(cloudConfigID).WithMachinePoolName(machinePoolName)
	_, err := h.Client.V1CloudConfigsGkeMachinePoolDelete(params)
	return err
}

// GetNodeStatusMapGke retrieves the status of all nodes in a GKE machine pool.
func (h *V1Client) GetNodeStatusMapGke(configUID string, machinePoolName string) (map[string]models.V1CloudMachineStatus, error) {
	params := clientv1.NewV1CloudConfigsGkePoolMachinesListParamsWithContext(h.ctx).WithConfigUID(configUID).WithMachinePoolName(machinePoolName)

	mpList, err := h.Client.V1CloudConfigsGkePoolMachinesList(params)
	nMap := map[string]models.V1CloudMachineStatus{}
	if len(mpList.Payload.Items) > 0 {
		for _, node := range mpList.Payload.Items {
			nMap[node.Metadata.UID] = *node.Status
		}
	}
	return nMap, err
}
