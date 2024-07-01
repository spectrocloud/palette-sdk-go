package client

import (
	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
)

func (h *V1Client) CreateClusterGke(cluster *models.V1SpectroGcpClusterEntity) (string, error) {
	params := clientV1.NewV1SpectroClustersGkeCreateParamsWithContext(h.ctx).WithBody(cluster)
	success, err := h.Client.V1SpectroClustersGkeCreate(params)
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

func (h *V1Client) GetCloudConfigGke(configUID string) (*models.V1GcpCloudConfig, error) {
	params := clientV1.NewV1CloudConfigsGkeGetParamsWithContext(h.ctx).WithConfigUID(configUID)

	success, err := h.Client.V1CloudConfigsGkeGet(params)
	if err != nil {
		return nil, err
	}

	return success.Payload, nil
}

func (h *V1Client) CreateMachinePoolGke(cloudConfigId string, machinePool *models.V1GcpMachinePoolConfigEntity) error {
	params := clientV1.NewV1CloudConfigsGkeMachinePoolCreateParamsWithContext(h.ctx).WithConfigUID(cloudConfigId).WithBody(machinePool)
	_, err := h.Client.V1CloudConfigsGkeMachinePoolCreate(params)
	return err
}

func (h *V1Client) UpdateMachinePoolGke(cloudConfigId string, machinePool *models.V1GcpMachinePoolConfigEntity) error {
	params := clientV1.NewV1CloudConfigsGkeMachinePoolUpdateParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigId).
		WithMachinePoolName(*machinePool.PoolConfig.Name).
		WithBody(machinePool)

	_, err := h.Client.V1CloudConfigsGkeMachinePoolUpdate(params)
	return err
}

func (h *V1Client) DeleteMachinePoolGke(cloudConfigId string, machinePoolName string) error {
	params := clientV1.NewV1CloudConfigsGkeMachinePoolDeleteParamsWithContext(h.ctx).WithConfigUID(cloudConfigId).WithMachinePoolName(machinePoolName)
	_, err := h.Client.V1CloudConfigsGkeMachinePoolDelete(params)
	return err
}

func (h *V1Client) GetNodeStatusMapGke(configUID string, machinePoolName string) (map[string]models.V1CloudMachineStatus, error) {
	params := clientV1.NewV1CloudConfigsGkePoolMachinesListParamsWithContext(h.ctx).WithConfigUID(configUID).WithMachinePoolName(machinePoolName)

	mpList, err := h.Client.V1CloudConfigsGkePoolMachinesList(params)
	nMap := map[string]models.V1CloudMachineStatus{}
	if len(mpList.Payload.Items) > 0 {
		for _, node := range mpList.Payload.Items {
			nMap[node.Metadata.UID] = *node.Status
		}
	}
	return nMap, err
}
