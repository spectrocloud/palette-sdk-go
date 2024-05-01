package client

import (
	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
	"github.com/spectrocloud/palette-sdk-go/client/apiutil"
)

func (h *V1Client) CreateClusterLibvirt(cluster *models.V1SpectroLibvirtClusterEntity) (string, error) {
	params := clientV1.NewV1SpectroClustersLibvirtCreateParamsWithContext(h.ctx).
		WithBody(cluster)
	resp, err := h.Client.V1SpectroClustersLibvirtCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

func (h *V1Client) CreateMachinePoolLibvirt(cloudConfigUid string, machinePool *models.V1LibvirtMachinePoolConfigEntity) error {
	params := clientV1.NewV1CloudConfigsLibvirtMachinePoolCreateParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUid).
		WithBody(machinePool)
	_, err := h.Client.V1CloudConfigsLibvirtMachinePoolCreate(params)
	return err
}

func (h *V1Client) UpdateMachinePoolLibvirt(cloudConfigUid string, machinePool *models.V1LibvirtMachinePoolConfigEntity) error {
	params := clientV1.NewV1CloudConfigsLibvirtMachinePoolUpdateParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUid).
		WithMachinePoolName(*machinePool.PoolConfig.Name).
		WithBody(machinePool)
	_, err := h.Client.V1CloudConfigsLibvirtMachinePoolUpdate(params)
	return err
}

func (h *V1Client) DeleteMachinePoolLibvirt(cloudConfigUid, machinePoolName string) error {
	params := clientV1.NewV1CloudConfigsLibvirtMachinePoolDeleteParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUid).
		WithMachinePoolName(machinePoolName)
	_, err := h.Client.V1CloudConfigsLibvirtMachinePoolDelete(params)
	return err
}

func (h *V1Client) GetCloudConfigLibvirt(configUid string) (*models.V1LibvirtCloudConfig, error) {
	params := clientV1.NewV1CloudConfigsLibvirtGetParamsWithContext(h.ctx).
		WithConfigUID(configUid)
	resp, err := h.Client.V1CloudConfigsLibvirtGet(params)
	if apiutil.Is404(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

func (h *V1Client) GetNodeStatusMapLibvirt(configUid, machinePoolName string) (map[string]models.V1CloudMachineStatus, error) {
	params := clientV1.NewV1CloudConfigsLibvirtPoolMachinesListParamsWithContext(h.ctx).
		WithConfigUID(configUid).
		WithMachinePoolName(machinePoolName)
	mpList, err := h.Client.V1CloudConfigsLibvirtPoolMachinesList(params)
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
