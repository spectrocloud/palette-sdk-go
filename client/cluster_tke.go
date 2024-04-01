package client

import (
	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
	"github.com/spectrocloud/palette-sdk-go/client/apiutil"
)

func (h *V1Client) CreateClusterTke(cluster *models.V1SpectroTencentClusterEntity) (string, error) {
	params := clientV1.NewV1SpectroClustersTkeCreateParamsWithContext(h.ctx).
		WithBody(cluster)
	resp, err := h.Client.V1SpectroClustersTkeCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

func (h *V1Client) CreateMachinePoolTke(cloudConfigUid string, machinePool *models.V1TencentMachinePoolConfigEntity) error {
	params := clientV1.NewV1CloudConfigsTkeMachinePoolCreateParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUid).
		WithBody(machinePool)
	_, err := h.Client.V1CloudConfigsTkeMachinePoolCreate(params)
	return err
}

func (h *V1Client) UpdateMachinePoolTke(cloudConfigUid string, machinePool *models.V1TencentMachinePoolConfigEntity) error {
	params := clientV1.NewV1CloudConfigsTkeMachinePoolUpdateParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUid).
		WithMachinePoolName(*machinePool.PoolConfig.Name).
		WithBody(machinePool)
	_, err := h.Client.V1CloudConfigsTkeMachinePoolUpdate(params)
	return err
}

func (h *V1Client) DeleteMachinePoolTke(cloudConfigUid, machinePoolName string) error {
	params := clientV1.NewV1CloudConfigsTkeMachinePoolDeleteParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUid).
		WithMachinePoolName(machinePoolName)
	_, err := h.Client.V1CloudConfigsTkeMachinePoolDelete(params)
	return err
}

func (h *V1Client) GetCloudConfigTke(configUid string) (*models.V1TencentCloudConfig, error) {
	params := clientV1.NewV1CloudConfigsTkeGetParamsWithContext(h.ctx).
		WithConfigUID(configUid)
	resp, err := h.Client.V1CloudConfigsTkeGet(params)
	if err := apiutil.Handle404(err); err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

func (h *V1Client) GetNodeStatusMapTke(configUid, machinePoolName string) (map[string]models.V1CloudMachineStatus, error) {
	params := clientV1.NewV1CloudConfigsTkePoolMachinesListParamsWithContext(h.ctx).
		WithConfigUID(configUid).
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
