package client

import (
	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
	"github.com/spectrocloud/palette-sdk-go/client/apiutil"
)

func (h *V1Client) CreateClusterVirtual(cluster *models.V1SpectroVirtualClusterEntity) (string, error) {
	params := clientV1.NewV1SpectroClustersVirtualCreateParamsWithContext(h.ctx).
		WithBody(cluster)
	resp, err := h.Client.V1SpectroClustersVirtualCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

func (h *V1Client) ResizeClusterVirtual(configUid string, body *models.V1VirtualClusterResize) error {
	params := clientV1.NewV1CloudConfigsVirtualUIDUpdateParamsWithContext(h.ctx).
		WithConfigUID(configUid).
		WithBody(body)
	_, err := h.Client.V1CloudConfigsVirtualUIDUpdate(params)
	return err
}

func (h *V1Client) CreateMachinePoolVirtual(cloudConfigUid string, machinePool *models.V1VirtualMachinePoolConfigEntity) error {
	params := clientV1.NewV1CloudConfigsVirtualMachinePoolCreateParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUid).
		WithBody(machinePool)
	_, err := h.Client.V1CloudConfigsVirtualMachinePoolCreate(params)
	return err
}

func (h *V1Client) UpdateMachinePoolVirtual(cloudConfigUid string, machinePool *models.V1VirtualMachinePoolConfigEntity) error {
	params := clientV1.NewV1CloudConfigsVirtualMachinePoolUpdateParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUid).
		WithBody(machinePool)
	_, err := h.Client.V1CloudConfigsVirtualMachinePoolUpdate(params)
	return err
}

func (h *V1Client) DeleteMachinePoolVirtual(cloudConfigUid, machinePoolName string) error {
	params := clientV1.NewV1CloudConfigsVirtualMachinePoolDeleteParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUid).
		WithMachinePoolName(machinePoolName)
	_, err := h.Client.V1CloudConfigsVirtualMachinePoolDelete(params)
	return err
}

func (h *V1Client) GetCloudConfigVirtual(configUid string) (*models.V1VirtualCloudConfig, error) {
	params := clientV1.NewV1CloudConfigsVirtualGetParamsWithContext(h.ctx).
		WithConfigUID(configUid)
	resp, err := h.Client.V1CloudConfigsVirtualGet(params)
	if apiutil.Is404(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

func (h *V1Client) VirtualClusterLifecycleConfigChange(uid string, body *models.V1LifecycleConfigEntity) (string, error) {
	params := clientV1.NewV1SpectroClustersUIDLifecycleConfigUpdateParamsWithContext(h.ctx).
		WithUID(uid).
		WithBody(body)
	_, err := h.Client.V1SpectroClustersUIDLifecycleConfigUpdate(params)
	if err != nil {
		return "Fail", err
	}
	return "Success", nil
}

func (h *V1Client) GetNodeStatusMapVirtual(configUid, machinePoolName string) (map[string]models.V1CloudMachineStatus, error) {
	params := clientV1.NewV1CloudConfigsVirtualPoolMachinesListParamsWithContext(h.ctx).
		WithConfigUID(configUid).
		WithMachinePoolName(machinePoolName)
	mpList, err := h.Client.V1CloudConfigsVirtualPoolMachinesList(params)
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
