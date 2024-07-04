package client

import (
	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
)

type GetMaintenanceStatus func(string, string, string) (*models.V1MachineMaintenanceStatus, error)

func (h *V1Client) ToggleMaintenanceOnNode(nodeMaintenance *models.V1MachineMaintenance, cloudType, configUid, machineName, nodeId string) error {
	params := clientV1.NewV1CloudConfigsMachinePoolsMachineUIDMaintenanceUpdateParamsWithContext(h.ctx).
		WithBody(nodeMaintenance).
		WithCloudType(cloudType).
		WithConfigUID(configUid).
		WithMachinePoolName(machineName).
		WithMachineUID(nodeId)

	_, err := h.Client.V1CloudConfigsMachinePoolsMachineUIDMaintenanceUpdate(params)
	return err
}

func (h *V1Client) GetNodeMaintenanceStatus(fn GetMaintenanceStatus, configUid, machineName, nodeId string) (*models.V1MachineMaintenanceStatus, error) {
	return fn(configUid, machineName, nodeId)
}

func (h *V1Client) GetNodeMaintenanceStatusAws(configUid, machineName, nodeId string) (*models.V1MachineMaintenanceStatus, error) {
	params := clientV1.NewV1CloudConfigsAwsPoolMachinesUIDGetParamsWithContext(h.ctx).
		WithConfigUID(configUid).
		WithMachinePoolName(machineName).
		WithMachineUID(nodeId)

	resp, err := h.Client.V1CloudConfigsAwsPoolMachinesUIDGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload.Status.MaintenanceStatus, nil
}

func (h *V1Client) GetNodeMaintenanceStatusMaas(configUid, machineName, nodeId string) (*models.V1MachineMaintenanceStatus, error) {
	params := clientV1.NewV1CloudConfigsMaasPoolMachinesUIDGetParamsWithContext(h.ctx).
		WithConfigUID(configUid).
		WithMachinePoolName(machineName).
		WithMachineUID(nodeId)

	resp, err := h.Client.V1CloudConfigsMaasPoolMachinesUIDGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload.Status.MaintenanceStatus, nil
}

func (h *V1Client) GetNodeMaintenanceStatusAks(configUid, machineName, nodeId string) (*models.V1MachineMaintenanceStatus, error) {
	params := clientV1.NewV1CloudConfigsAksPoolMachinesUIDGetParamsWithContext(h.ctx).
		WithConfigUID(configUid).
		WithMachinePoolName(machineName).
		WithMachineUID(nodeId)

	resp, err := h.Client.V1CloudConfigsAksPoolMachinesUIDGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload.Status.MaintenanceStatus, nil
}

func (h *V1Client) GetNodeMaintenanceStatusAzure(configUid, machineName, nodeId string) (*models.V1MachineMaintenanceStatus, error) {
	params := clientV1.NewV1CloudConfigsAzurePoolMachinesUIDGetParamsWithContext(h.ctx).
		WithConfigUID(configUid).
		WithMachinePoolName(machineName).
		WithMachineUID(nodeId)

	resp, err := h.Client.V1CloudConfigsAzurePoolMachinesUIDGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload.Status.MaintenanceStatus, nil
}

//func (h *V1Client) GetNodeMaintenanceStatusCoxEdge(configUid, machineName, nodeId string) (*models.V1MachineMaintenanceStatus, error) {
//	params := clientV1.NewV1CloudConfigsCoxEdgePoolMachinesUIDGetParamsWithContext(h.ctx).
//		WithConfigUID(configUid).
//		WithMachinePoolName(machineName).
//		WithMachineUID(nodeId)
//
//	resp, err := h.Client.V1CloudConfigsCoxEdgePoolMachinesUIDGet(params)
//	if err != nil {
//		return nil, err
//	}
//	return resp.Payload.Status.MaintenanceStatus, nil
//}

func (h *V1Client) GetNodeMaintenanceStatusEdgeNative(configUid, machineName, nodeId string) (*models.V1MachineMaintenanceStatus, error) {
	params := clientV1.NewV1CloudConfigsEdgeNativePoolMachinesUIDGetParamsWithContext(h.ctx).
		WithConfigUID(configUid).
		WithMachinePoolName(machineName).
		WithMachineUID(nodeId)

	resp, err := h.Client.V1CloudConfigsEdgeNativePoolMachinesUIDGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload.Status.MaintenanceStatus, nil
}

//func (h *V1Client) GetNodeMaintenanceStatusEdge(configUid, machineName, nodeId string) (*models.V1MachineMaintenanceStatus, error) {
//	params := clientV1.NewV1CloudConfigsEdgePoolMachinesUIDGetParamsWithContext(h.ctx).
//		WithConfigUID(configUid).
//		WithMachinePoolName(machineName).
//		WithMachineUID(nodeId)
//
//	resp, err := h.Client.V1CloudConfigsEdgePoolMachinesUIDGet(params)
//	if err != nil {
//		return nil, err
//	}
//	return resp.Payload.Status.MaintenanceStatus, nil
//}

func (h *V1Client) GetNodeMaintenanceStatusEdgeVsphere(configUid, machineName, nodeId string) (*models.V1MachineMaintenanceStatus, error) {
	params := clientV1.NewV1CloudConfigsVspherePoolMachinesUIDGetParamsWithContext(h.ctx).
		WithConfigUID(configUid).
		WithMachinePoolName(machineName).
		WithMachineUID(nodeId)

	resp, err := h.Client.V1CloudConfigsVspherePoolMachinesUIDGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload.Status.MaintenanceStatus, nil
}

func (h *V1Client) GetNodeMaintenanceStatusEks(configUid, machineName, nodeId string) (*models.V1MachineMaintenanceStatus, error) {
	params := clientV1.NewV1CloudConfigsEksPoolMachinesUIDGetParamsWithContext(h.ctx).
		WithConfigUID(configUid).
		WithMachinePoolName(machineName).
		WithMachineUID(nodeId)

	resp, err := h.Client.V1CloudConfigsEksPoolMachinesUIDGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload.Status.MaintenanceStatus, nil
}

func (h *V1Client) GetNodeMaintenanceStatusGcp(configUid, machineName, nodeId string) (*models.V1MachineMaintenanceStatus, error) {
	params := clientV1.NewV1CloudConfigsGcpPoolMachinesUIDGetParamsWithContext(h.ctx).
		WithConfigUID(configUid).
		WithMachinePoolName(machineName).
		WithMachineUID(nodeId)

	resp, err := h.Client.V1CloudConfigsGcpPoolMachinesUIDGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload.Status.MaintenanceStatus, nil
}

func (h *V1Client) GetNodeMaintenanceStatusGeneric(configUid, machineName, nodeId string) (*models.V1MachineMaintenanceStatus, error) {
	params := clientV1.NewV1CloudConfigsGenericPoolMachinesUIDGetParamsWithContext(h.ctx).
		WithConfigUID(configUid).
		WithMachinePoolName(machineName).
		WithMachineUID(nodeId)

	resp, err := h.Client.V1CloudConfigsGenericPoolMachinesUIDGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload.Status.MaintenanceStatus, nil
}

func (h *V1Client) GetNodeMaintenanceStatusGke(configUid, machineName, nodeId string) (*models.V1MachineMaintenanceStatus, error) {
	params := clientV1.NewV1CloudConfigsGkePoolMachinesUIDGetParamsWithContext(h.ctx).
		WithConfigUID(configUid).
		WithMachinePoolName(machineName).
		WithMachineUID(nodeId)

	resp, err := h.Client.V1CloudConfigsGkePoolMachinesUIDGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload.Status.MaintenanceStatus, nil
}

func (h *V1Client) GetNodeMaintenanceStatusLibvirt(configUid, machineName, nodeId string) (*models.V1MachineMaintenanceStatus, error) {
	params := clientV1.NewV1CloudConfigsLibvirtPoolMachinesUIDGetParamsWithContext(h.ctx).
		WithConfigUID(configUid).
		WithMachinePoolName(machineName).
		WithMachineUID(nodeId)

	resp, err := h.Client.V1CloudConfigsLibvirtPoolMachinesUIDGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload.Status.MaintenanceStatus, nil
}

func (h *V1Client) GetNodeMaintenanceStatusOpenStack(configUid, machineName, nodeId string) (*models.V1MachineMaintenanceStatus, error) {
	params := clientV1.NewV1CloudConfigsOpenStackPoolMachinesUIDGetParamsWithContext(h.ctx).
		WithConfigUID(configUid).
		WithMachinePoolName(machineName).
		WithMachineUID(nodeId)

	resp, err := h.Client.V1CloudConfigsOpenStackPoolMachinesUIDGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload.Status.MaintenanceStatus, nil
}

func (h *V1Client) GetNodeMaintenanceStatusTke(configUid, machineName, nodeId string) (*models.V1MachineMaintenanceStatus, error) {
	params := clientV1.NewV1CloudConfigsTkePoolMachinesUIDGetParamsWithContext(h.ctx).
		WithConfigUID(configUid).
		WithMachinePoolName(machineName).
		WithMachineUID(nodeId)

	resp, err := h.Client.V1CloudConfigsTkePoolMachinesUIDGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload.Status.MaintenanceStatus, nil
}

func (h *V1Client) GetNodeVirtualMaintenanceStatusVirtual(configUid, machineName, nodeId string) (*models.V1MachineMaintenanceStatus, error) {
	params := clientV1.NewV1CloudConfigsVirtualPoolMachinesUIDGetParamsWithContext(h.ctx).
		WithConfigUID(configUid).
		WithMachinePoolName(machineName).
		WithMachineUID(nodeId)

	resp, err := h.Client.V1CloudConfigsVirtualPoolMachinesUIDGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload.Status.MaintenanceStatus, nil
}

func (h *V1Client) GetNodeMaintenanceStatusVsphere(configUid, machineName, nodeId string) (*models.V1MachineMaintenanceStatus, error) {
	params := clientV1.NewV1CloudConfigsVspherePoolMachinesUIDGetParamsWithContext(h.ctx).
		WithConfigUID(configUid).
		WithMachinePoolName(machineName).
		WithMachineUID(nodeId)

	resp, err := h.Client.V1CloudConfigsVspherePoolMachinesUIDGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload.Status.MaintenanceStatus, nil
}
