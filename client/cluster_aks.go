package client

import (
	"time"

	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
	"github.com/spectrocloud/palette-sdk-go/client/apiutil"
)

func (h *V1Client) CreateClusterAks(cluster *models.V1SpectroAzureClusterEntity) (string, error) {
	params := clientV1.NewV1SpectroClustersAksCreateParamsWithContext(h.ctx).
		WithBody(cluster).
		WithTimeout(90 * time.Second)
	resp, err := h.Client.V1SpectroClustersAksCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

func (h *V1Client) CreateMachinePoolAks(cloudConfigUid string, machinePool *models.V1AzureMachinePoolConfigEntity) error {
	params := clientV1.NewV1CloudConfigsAksMachinePoolCreateParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUid).
		WithBody(machinePool)
	_, err := h.Client.V1CloudConfigsAksMachinePoolCreate(params)
	return err
}

func (h *V1Client) UpdateMachinePoolAks(cloudConfigUid string, machinePool *models.V1AzureMachinePoolConfigEntity) error {
	params := clientV1.NewV1CloudConfigsAksMachinePoolUpdateParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUid).
		WithMachinePoolName(*machinePool.PoolConfig.Name).
		WithBody(machinePool)
	_, err := h.Client.V1CloudConfigsAksMachinePoolUpdate(params)
	return err
}

func (h *V1Client) DeleteMachinePoolAks(cloudConfigUid, machinePoolName string) error {
	params := clientV1.NewV1CloudConfigsAksMachinePoolDeleteParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUid).
		WithMachinePoolName(machinePoolName)
	_, err := h.Client.V1CloudConfigsAksMachinePoolDelete(params)
	return err
}

func (h *V1Client) GetCloudConfigAks(configUid string) (*models.V1AzureCloudConfig, error) {
	params := clientV1.NewV1CloudConfigsAksGetParamsWithContext(h.ctx).
		WithConfigUID(configUid)
	resp, err := h.Client.V1CloudConfigsAksGet(params)
	if err := apiutil.Handle404(err); err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

func (h *V1Client) GetNodeStatusMapAks(configUid, machinePoolName string) (map[string]models.V1CloudMachineStatus, error) {
	params := clientV1.NewV1CloudConfigsAksPoolMachinesListParamsWithContext(h.ctx).
		WithConfigUID(configUid).
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
