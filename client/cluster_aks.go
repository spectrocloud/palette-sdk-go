package client

import (
	"time"

	clientv1 "github.com/spectrocloud/palette-sdk-go/api/client/v1"
	"github.com/spectrocloud/palette-sdk-go/api/models"
	"github.com/spectrocloud/palette-sdk-go/client/apiutil"
)

// CreateClusterAks creates a new AKS cluster.
func (h *V1Client) CreateClusterAks(cluster *models.V1SpectroAzureClusterEntity) (string, error) {
	params := clientv1.NewV1SpectroClustersAksCreateParamsWithContext(h.ctx).
		WithBody(cluster).
		WithTimeout(90 * time.Second)
	resp, err := h.Client.V1SpectroClustersAksCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

// CreateMachinePoolAks creates a new AKS machine pool.
func (h *V1Client) CreateMachinePoolAks(cloudConfigUID string, machinePool *models.V1AzureMachinePoolConfigEntity) error {
	params := clientv1.NewV1CloudConfigsAksMachinePoolCreateParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUID).
		WithBody(machinePool)
	_, err := h.Client.V1CloudConfigsAksMachinePoolCreate(params)
	return err
}

// UpdateMachinePoolAks updates an existing AKS machine pool.
func (h *V1Client) UpdateMachinePoolAks(cloudConfigUID string, machinePool *models.V1AzureMachinePoolConfigEntity) error {
	params := clientv1.NewV1CloudConfigsAksMachinePoolUpdateParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUID).
		WithMachinePoolName(*machinePool.PoolConfig.Name).
		WithBody(machinePool)
	_, err := h.Client.V1CloudConfigsAksMachinePoolUpdate(params)
	return err
}

// DeleteMachinePoolAks deletes an existing AKS machine pool.
func (h *V1Client) DeleteMachinePoolAks(cloudConfigUID, machinePoolName string) error {
	params := clientv1.NewV1CloudConfigsAksMachinePoolDeleteParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUID).
		WithMachinePoolName(machinePoolName)
	_, err := h.Client.V1CloudConfigsAksMachinePoolDelete(params)
	return err
}

// GetCloudConfigAks retrieves an existing AKS cluster's cloud config.
func (h *V1Client) GetCloudConfigAks(configUID string) (*models.V1AzureCloudConfig, error) {
	params := clientv1.NewV1CloudConfigsAksGetParamsWithContext(h.ctx).
		WithConfigUID(configUID)
	resp, err := h.Client.V1CloudConfigsAksGet(params)
	if apiutil.Is404(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// GetNodeStatusMapAks retrieves the status of all nodes in an AKS machine pool.
func (h *V1Client) GetNodeStatusMapAks(configUID, machinePoolName string) (map[string]models.V1CloudMachineStatus, error) {
	params := clientv1.NewV1CloudConfigsAksPoolMachinesListParamsWithContext(h.ctx).
		WithConfigUID(configUID).
		WithMachinePoolName(machinePoolName)
	mpList, err := h.Client.V1CloudConfigsAksPoolMachinesList(params)
	if err != nil {
		return nil, err
	}
	nMap := map[string]models.V1CloudMachineStatus{}
	if len(mpList.Payload.Items) > 0 {
		for _, node := range mpList.Payload.Items {
			if node.Status.MaintenanceStatus.Action != "" {
				nMap[node.Metadata.UID] = *node.Status
			}
		}
	}
	return nMap, nil
}
