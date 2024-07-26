package client

import (
	clientv1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
)

// GetMaintenanceStatus defines a function type that retrieves the maintenance status of a machine.
type GetMaintenanceStatus func(string, string, string) (*models.V1MachineMaintenanceStatus, error)

// ToggleMaintenanceOnNode updates maintenance configuration for a node.
func (h *V1Client) ToggleMaintenanceOnNode(nodeMaintenance *models.V1MachineMaintenance, cloudType, configUID, machineName, nodeID string) error {
	params := clientv1.NewV1CloudConfigsMachinePoolsMachineUIDMaintenanceUpdateParamsWithContext(h.ctx).
		WithBody(nodeMaintenance).
		WithCloudType(cloudType).
		WithConfigUID(configUID).
		WithMachinePoolName(machineName).
		WithMachineUID(nodeID)

	_, err := h.Client.V1CloudConfigsMachinePoolsMachineUIDMaintenanceUpdate(params)
	return err
}

// GetNodeMaintenanceStatus retrieves the maintenance status of a specific node.
func (h *V1Client) GetNodeMaintenanceStatus(fn GetMaintenanceStatus, configUid, machineName, nodeId string) (*models.V1MachineMaintenanceStatus, error) {
	return fn(configUid, machineName, nodeId)
}

// GetNodeMaintenanceStatusAws retrieves maintenance status for an AWS IaaS node.
func (h *V1Client) GetNodeMaintenanceStatusAws(configUID, machineName, nodeID string) (*models.V1MachineMaintenanceStatus, error) {
	params := clientv1.NewV1CloudConfigsAwsPoolMachinesUIDGetParamsWithContext(h.ctx).
		WithConfigUID(configUID).
		WithMachinePoolName(machineName).
		WithMachineUID(nodeID)

	resp, err := h.Client.V1CloudConfigsAwsPoolMachinesUIDGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload.Status.MaintenanceStatus, nil
}

// GetNodeMaintenanceStatusMaas retrieves maintenance status for a MAAS node.
func (h *V1Client) GetNodeMaintenanceStatusMaas(configUID, machineName, nodeID string) (*models.V1MachineMaintenanceStatus, error) {
	params := clientv1.NewV1CloudConfigsMaasPoolMachinesUIDGetParamsWithContext(h.ctx).
		WithConfigUID(configUID).
		WithMachinePoolName(machineName).
		WithMachineUID(nodeID)

	resp, err := h.Client.V1CloudConfigsMaasPoolMachinesUIDGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload.Status.MaintenanceStatus, nil
}

// GetNodeMaintenanceStatusAks retrieves maintenance status for an AKS node.
func (h *V1Client) GetNodeMaintenanceStatusAks(configUID, machineName, nodeID string) (*models.V1MachineMaintenanceStatus, error) {
	params := clientv1.NewV1CloudConfigsAksPoolMachinesUIDGetParamsWithContext(h.ctx).
		WithConfigUID(configUID).
		WithMachinePoolName(machineName).
		WithMachineUID(nodeID)

	resp, err := h.Client.V1CloudConfigsAksPoolMachinesUIDGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload.Status.MaintenanceStatus, nil
}

// GetNodeMaintenanceStatusAzure retrieves maintenance status for an Azure IaaS node.
func (h *V1Client) GetNodeMaintenanceStatusAzure(configUID, machineName, nodeID string) (*models.V1MachineMaintenanceStatus, error) {
	params := clientv1.NewV1CloudConfigsAzurePoolMachinesUIDGetParamsWithContext(h.ctx).
		WithConfigUID(configUID).
		WithMachinePoolName(machineName).
		WithMachineUID(nodeID)

	resp, err := h.Client.V1CloudConfigsAzurePoolMachinesUIDGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload.Status.MaintenanceStatus, nil
}

// GetNodeMaintenanceStatusEdgeNative retrieves maintenance status for an edge native node.
func (h *V1Client) GetNodeMaintenanceStatusEdgeNative(configUID, machineName, nodeID string) (*models.V1MachineMaintenanceStatus, error) {
	params := clientv1.NewV1CloudConfigsEdgeNativePoolMachinesUIDGetParamsWithContext(h.ctx).
		WithConfigUID(configUID).
		WithMachinePoolName(machineName).
		WithMachineUID(nodeID)

	resp, err := h.Client.V1CloudConfigsEdgeNativePoolMachinesUIDGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload.Status.MaintenanceStatus, nil
}

// GetNodeMaintenanceStatusEdgeVsphere retrieves maintenance status for a vSphere edge node.
// TODO: edgev1 deprecation
func (h *V1Client) GetNodeMaintenanceStatusEdgeVsphere(configUID, machineName, nodeID string) (*models.V1MachineMaintenanceStatus, error) {
	params := clientv1.NewV1CloudConfigsVspherePoolMachinesUIDGetParamsWithContext(h.ctx).
		WithConfigUID(configUID).
		WithMachinePoolName(machineName).
		WithMachineUID(nodeID)

	resp, err := h.Client.V1CloudConfigsVspherePoolMachinesUIDGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload.Status.MaintenanceStatus, nil
}

// GetNodeMaintenanceStatusEks retrieves maintenance status for an EKS node.
func (h *V1Client) GetNodeMaintenanceStatusEks(configUID, machineName, nodeID string) (*models.V1MachineMaintenanceStatus, error) {
	params := clientv1.NewV1CloudConfigsEksPoolMachinesUIDGetParamsWithContext(h.ctx).
		WithConfigUID(configUID).
		WithMachinePoolName(machineName).
		WithMachineUID(nodeID)

	resp, err := h.Client.V1CloudConfigsEksPoolMachinesUIDGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload.Status.MaintenanceStatus, nil
}

// GetNodeMaintenanceStatusGcp retrieves maintenance status for a GCP IaaS node.
func (h *V1Client) GetNodeMaintenanceStatusGcp(configUID, machineName, nodeID string) (*models.V1MachineMaintenanceStatus, error) {
	params := clientv1.NewV1CloudConfigsGcpPoolMachinesUIDGetParamsWithContext(h.ctx).
		WithConfigUID(configUID).
		WithMachinePoolName(machineName).
		WithMachineUID(nodeID)

	resp, err := h.Client.V1CloudConfigsGcpPoolMachinesUIDGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload.Status.MaintenanceStatus, nil
}

// GetNodeMaintenanceStatusGeneric retrieves maintenance status for a generic node.
func (h *V1Client) GetNodeMaintenanceStatusGeneric(configUID, machineName, nodeID string) (*models.V1MachineMaintenanceStatus, error) {
	params := clientv1.NewV1CloudConfigsGenericPoolMachinesUIDGetParamsWithContext(h.ctx).
		WithConfigUID(configUID).
		WithMachinePoolName(machineName).
		WithMachineUID(nodeID)

	resp, err := h.Client.V1CloudConfigsGenericPoolMachinesUIDGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload.Status.MaintenanceStatus, nil
}

// GetNodeMaintenanceStatusGke retrieves maintenance status for a GKE node.
func (h *V1Client) GetNodeMaintenanceStatusGke(configUID, machineName, nodeID string) (*models.V1MachineMaintenanceStatus, error) {
	params := clientv1.NewV1CloudConfigsGkePoolMachinesUIDGetParamsWithContext(h.ctx).
		WithConfigUID(configUID).
		WithMachinePoolName(machineName).
		WithMachineUID(nodeID)

	resp, err := h.Client.V1CloudConfigsGkePoolMachinesUIDGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload.Status.MaintenanceStatus, nil
}

// GetNodeMaintenanceStatusOpenStack retrieves maintenance status for an OpenStack node.
func (h *V1Client) GetNodeMaintenanceStatusOpenStack(configUID, machineName, nodeID string) (*models.V1MachineMaintenanceStatus, error) {
	params := clientv1.NewV1CloudConfigsOpenStackPoolMachinesUIDGetParamsWithContext(h.ctx).
		WithConfigUID(configUID).
		WithMachinePoolName(machineName).
		WithMachineUID(nodeID)

	resp, err := h.Client.V1CloudConfigsOpenStackPoolMachinesUIDGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload.Status.MaintenanceStatus, nil
}

// GetNodeMaintenanceStatusTke retrieves maintenance status for a TKE node.
func (h *V1Client) GetNodeMaintenanceStatusTke(configUID, machineName, nodeID string) (*models.V1MachineMaintenanceStatus, error) {
	params := clientv1.NewV1CloudConfigsTkePoolMachinesUIDGetParamsWithContext(h.ctx).
		WithConfigUID(configUID).
		WithMachinePoolName(machineName).
		WithMachineUID(nodeID)

	resp, err := h.Client.V1CloudConfigsTkePoolMachinesUIDGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload.Status.MaintenanceStatus, nil
}

// GetNodeVirtualMaintenanceStatusVirtual retrieves maintenance status for a virtual node.
// TODO: deprecate unused virtual cluster functions
func (h *V1Client) GetNodeVirtualMaintenanceStatusVirtual(configUID, machineName, nodeID string) (*models.V1MachineMaintenanceStatus, error) {
	params := clientv1.NewV1CloudConfigsVirtualPoolMachinesUIDGetParamsWithContext(h.ctx).
		WithConfigUID(configUID).
		WithMachinePoolName(machineName).
		WithMachineUID(nodeID)

	resp, err := h.Client.V1CloudConfigsVirtualPoolMachinesUIDGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload.Status.MaintenanceStatus, nil
}

// GetNodeMaintenanceStatusVsphere retrieves maintenance status for a vSphere node.
func (h *V1Client) GetNodeMaintenanceStatusVsphere(configUID, machineName, nodeID string) (*models.V1MachineMaintenanceStatus, error) {
	params := clientv1.NewV1CloudConfigsVspherePoolMachinesUIDGetParamsWithContext(h.ctx).
		WithConfigUID(configUID).
		WithMachinePoolName(machineName).
		WithMachineUID(nodeID)

	resp, err := h.Client.V1CloudConfigsVspherePoolMachinesUIDGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload.Status.MaintenanceStatus, nil
}
