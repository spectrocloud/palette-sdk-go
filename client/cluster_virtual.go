package client

import (
	"errors"

	"github.com/spectrocloud/hapi/apiutil/transport"
	"github.com/spectrocloud/hapi/models"
	clusterC "github.com/spectrocloud/hapi/spectrocluster/client/v1"
)

func (h *V1Client) CreateClusterVirtual(cluster *models.V1SpectroVirtualClusterEntity) (string, error) {
	params := clusterC.NewV1SpectroClustersVirtualCreateParamsWithContext(h.Ctx).WithBody(cluster)
	success, err := h.GetClusterClient().V1SpectroClustersVirtualCreate(params)
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

func (h *V1Client) ResizeClusterVirtual(configUID string, body *models.V1VirtualClusterResize) error {
	params := clusterC.NewV1CloudConfigsVirtualUIDUpdateParamsWithContext(h.Ctx).WithConfigUID(configUID).WithBody(body)
	_, err := h.GetClusterClient().V1CloudConfigsVirtualUIDUpdate(params)
	if err != nil {
		return err
	}
	return nil
}

func (h *V1Client) CreateMachinePoolVirtual(cloudConfigId string, machinePool *models.V1VirtualMachinePoolConfigEntity) error {
	params := clusterC.NewV1CloudConfigsVirtualMachinePoolCreateParamsWithContext(h.Ctx).WithConfigUID(cloudConfigId).WithBody(machinePool)
	_, err := h.GetClusterClient().V1CloudConfigsVirtualMachinePoolCreate(params)
	return err
}

func (h *V1Client) UpdateMachinePoolVirtual(cloudConfigId string, machinePool *models.V1VirtualMachinePoolConfigEntity) error {
	params := clusterC.NewV1CloudConfigsVirtualMachinePoolUpdateParamsWithContext(h.Ctx).
		WithConfigUID(cloudConfigId).
		WithBody(machinePool)
	_, err := h.GetClusterClient().V1CloudConfigsVirtualMachinePoolUpdate(params)
	return err
}

func (h *V1Client) DeleteMachinePoolVirtual(cloudConfigId, machinePoolName string) error {
	params := clusterC.NewV1CloudConfigsVirtualMachinePoolDeleteParamsWithContext(h.Ctx).WithConfigUID(cloudConfigId).WithMachinePoolName(machinePoolName)
	_, err := h.GetClusterClient().V1CloudConfigsVirtualMachinePoolDelete(params)
	return err
}

func (h *V1Client) GetCloudConfigVirtual(configUID string) (*models.V1VirtualCloudConfig, error) {
	params := clusterC.NewV1CloudConfigsVirtualGetParamsWithContext(h.Ctx).WithConfigUID(configUID)
	success, err := h.GetClusterClient().V1CloudConfigsVirtualGet(params)
	var e *transport.TransportError
	if errors.As(err, &e) && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return success.Payload, nil
}

func (h *V1Client) VirtualClusterLifecycleConfigChange(uid string, body *models.V1LifecycleConfigEntity) (string, error) {
	params := clusterC.NewV1SpectroClustersUIDLifecycleConfigUpdateParamsWithContext(h.Ctx).WithUID(uid).WithBody(body)
	_, err := h.GetClusterClient().V1SpectroClustersUIDLifecycleConfigUpdate(params)
	if err != nil {
		return "Fail", err
	}
	return "Success", nil
}

func (h *V1Client) GetNodeStatusMapVirtual(configUID, machinePoolName, ClusterContext string) (map[string]models.V1CloudMachineStatus, error) {
	var params *clusterC.V1CloudConfigsVirtualPoolMachinesListParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsVirtualPoolMachinesListParamsWithContext(h.Ctx).WithConfigUID(configUID).WithMachinePoolName(machinePoolName)
	case "tenant":
		params = clusterC.NewV1CloudConfigsVirtualPoolMachinesListParams().WithConfigUID(configUID).WithMachinePoolName(machinePoolName)
	}

	mpList, err := h.GetClusterClient().V1CloudConfigsVirtualPoolMachinesList(params)
	nMap := map[string]models.V1CloudMachineStatus{}
	if len(mpList.Payload.Items) > 0 {
		for _, node := range mpList.Payload.Items {
			nMap[node.Metadata.UID] = *node.Status
		}
	}
	return nMap, err
}
