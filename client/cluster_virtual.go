package client

import (
	"errors"

	"github.com/spectrocloud/palette-api-go/apiutil/transport"
	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
)

func (h *V1Client) CreateClusterVirtual(cluster *models.V1SpectroVirtualClusterEntity) (string, error) {
	params := clientV1.NewV1SpectroClustersVirtualCreateParamsWithContext(h.Ctx).WithBody(cluster)
	success, err := h.GetClient().V1SpectroClustersVirtualCreate(params)
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

func (h *V1Client) ResizeClusterVirtual(configUID string, body *models.V1VirtualClusterResize) error {
	params := clientV1.NewV1CloudConfigsVirtualUIDUpdateParamsWithContext(h.Ctx).WithConfigUID(configUID).WithBody(body)
	_, err := h.GetClient().V1CloudConfigsVirtualUIDUpdate(params)
	if err != nil {
		return err
	}
	return nil
}

func (h *V1Client) CreateMachinePoolVirtual(CloudConfigId string, machinePool *models.V1VirtualMachinePoolConfigEntity) error {
	params := clientV1.NewV1CloudConfigsVirtualMachinePoolCreateParamsWithContext(h.Ctx).WithConfigUID(CloudConfigId).WithBody(machinePool)
	_, err := h.GetClient().V1CloudConfigsVirtualMachinePoolCreate(params)
	return err
}

func (h *V1Client) UpdateMachinePoolVirtual(CloudConfigId string, machinePool *models.V1VirtualMachinePoolConfigEntity) error {
	params := clientV1.NewV1CloudConfigsVirtualMachinePoolUpdateParamsWithContext(h.Ctx).
		WithConfigUID(CloudConfigId).
		WithBody(machinePool)
	_, err := h.GetClient().V1CloudConfigsVirtualMachinePoolUpdate(params)
	return err
}

func (h *V1Client) DeleteMachinePoolVirtual(CloudConfigId, machinePoolName string) error {
	params := clientV1.NewV1CloudConfigsVirtualMachinePoolDeleteParamsWithContext(h.Ctx).WithConfigUID(CloudConfigId).WithMachinePoolName(machinePoolName)
	_, err := h.GetClient().V1CloudConfigsVirtualMachinePoolDelete(params)
	return err
}

func (h *V1Client) GetCloudConfigVirtual(configUID string) (*models.V1VirtualCloudConfig, error) {
	params := clientV1.NewV1CloudConfigsVirtualGetParamsWithContext(h.Ctx).WithConfigUID(configUID)
	success, err := h.GetClient().V1CloudConfigsVirtualGet(params)
	var e *transport.TransportError
	if errors.As(err, &e) && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return success.Payload, nil
}

func (h *V1Client) VirtualClusterLifecycleConfigChange(uid string, body *models.V1LifecycleConfigEntity) (string, error) {
	params := clientV1.NewV1SpectroClustersUIDLifecycleConfigUpdateParamsWithContext(h.Ctx).WithUID(uid).WithBody(body)
	_, err := h.GetClient().V1SpectroClustersUIDLifecycleConfigUpdate(params)
	if err != nil {
		return "Fail", err
	}
	return "Success", nil
}

func (h *V1Client) GetNodeStatusMapVirtual(configUID, machinePoolName, scope string) (map[string]models.V1CloudMachineStatus, error) {
	var params *clientV1.V1CloudConfigsVirtualPoolMachinesListParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsVirtualPoolMachinesListParamsWithContext(h.Ctx).WithConfigUID(configUID).WithMachinePoolName(machinePoolName)
	case "tenant":
		params = clientV1.NewV1CloudConfigsVirtualPoolMachinesListParams().WithConfigUID(configUID).WithMachinePoolName(machinePoolName)
	}

	mpList, err := h.GetClient().V1CloudConfigsVirtualPoolMachinesList(params)
	nMap := map[string]models.V1CloudMachineStatus{}
	if len(mpList.Payload.Items) > 0 {
		for _, node := range mpList.Payload.Items {
			nMap[node.Metadata.UID] = *node.Status
		}
	}
	return nMap, err
}
