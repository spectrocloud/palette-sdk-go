package client

import (
	clientv1 "github.com/spectrocloud/palette-sdk-go/api/client/v1"
	"github.com/spectrocloud/palette-sdk-go/api/models"
	"github.com/spectrocloud/palette-sdk-go/client/apiutil"
)

// CreateClusterTke creates a new TKE cluster.
func (h *V1Client) CreateClusterTke(cluster *models.V1SpectroTencentClusterEntity) (string, error) {
	params := clientv1.NewV1SpectroClustersTkeCreateParamsWithContext(h.ctx).
		WithBody(cluster)
	resp, err := h.Client.V1SpectroClustersTkeCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

// CreateMachinePoolTke creates a new TKE machine pool.
func (h *V1Client) CreateMachinePoolTke(cloudConfigUID string, machinePool *models.V1TencentMachinePoolConfigEntity) error {
	params := clientv1.NewV1CloudConfigsTkeMachinePoolCreateParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUID).
		WithBody(machinePool)
	_, err := h.Client.V1CloudConfigsTkeMachinePoolCreate(params)
	return err
}

// UpdateMachinePoolTke updates an existing TKE machine pool.
func (h *V1Client) UpdateMachinePoolTke(cloudConfigUID string, machinePool *models.V1TencentMachinePoolConfigEntity) error {
	params := clientv1.NewV1CloudConfigsTkeMachinePoolUpdateParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUID).
		WithMachinePoolName(*machinePool.PoolConfig.Name).
		WithBody(machinePool)
	_, err := h.Client.V1CloudConfigsTkeMachinePoolUpdate(params)
	return err
}

// DeleteMachinePoolTke deletes an existing TKE machine pool.
func (h *V1Client) DeleteMachinePoolTke(cloudConfigUID, machinePoolName string) error {
	params := clientv1.NewV1CloudConfigsTkeMachinePoolDeleteParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUID).
		WithMachinePoolName(machinePoolName)
	_, err := h.Client.V1CloudConfigsTkeMachinePoolDelete(params)
	return err
}

// GetCloudConfigTke retrieves an existing TKE cluster's cloud config.
func (h *V1Client) GetCloudConfigTke(configUID string) (*models.V1TencentCloudConfig, error) {
	params := clientv1.NewV1CloudConfigsTkeGetParamsWithContext(h.ctx).
		WithConfigUID(configUID)
	resp, err := h.Client.V1CloudConfigsTkeGet(params)
	if apiutil.Is404(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// GetNodeStatusMapTke retrieves the status of all nodes in a TKE machine pool.
func (h *V1Client) GetNodeStatusMapTke(configUID, machinePoolName string) (map[string]models.V1CloudMachineStatus, error) {
	params := clientv1.NewV1CloudConfigsTkePoolMachinesListParamsWithContext(h.ctx).
		WithConfigUID(configUID).
		WithMachinePoolName(machinePoolName)
	mpList, err := h.Client.V1CloudConfigsTkePoolMachinesList(params)
	if err != nil {
		return nil, err
	}
	nMap := map[string]models.V1CloudMachineStatus{}
	if len(mpList.Payload.Items) > 0 {
		for _, node := range mpList.Payload.Items {
			nMap[node.Metadata.UID] = *node.Status
		}
	}
	return nMap, nil
}
