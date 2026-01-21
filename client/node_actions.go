package client

import (
	clientv1 "github.com/spectrocloud/palette-sdk-go/api/client/version1"
	"github.com/spectrocloud/palette-sdk-go/api/models"
)

// GetMaintenanceStatus defines a function type that retrieves the maintenance status of a machine.
type GetMaintenanceStatus func(string, string, string) (*models.V1MachineMaintenanceStatus, error)

// GetMachinesList defines a function type that retrieves a list of machines from a machine pool.
// Returns a map where key is the machine name (Metadata.Name) and value is the machine UID (Metadata.UID).
type GetMachinesList func(string, string) (map[string]string, error)

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
func (h *V1Client) GetNodeMaintenanceStatus(fn GetMaintenanceStatus, ConfigUID, machineName, nodeID string) (*models.V1MachineMaintenanceStatus, error) {
	return fn(ConfigUID, machineName, nodeID)
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

// GetNodeMaintenanceStatusCloudStack retrieves maintenance status for a CloudStack node.
func (h *V1Client) GetNodeMaintenanceStatusCloudStack(configUID, machineName, nodeID string) (*models.V1MachineMaintenanceStatus, error) {
	params := clientv1.NewV1CloudConfigsCloudStackPoolMachinesUIDGetParamsWithContext(h.ctx).
		WithConfigUID(configUID).
		WithMachinePoolName(machineName).
		WithMachineUID(nodeID)

	resp, err := h.Client.V1CloudConfigsCloudStackPoolMachinesUIDGet(params)
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

// GetNodeListInEdgeNativeMachinePool retrieves machine in the edge-native machine pool
func (h *V1Client) GetNodeListInEdgeNativeMachinePool(configUID, machinePoolName string) (*models.V1EdgeNativeMachines, error) {
	params := clientv1.NewV1CloudConfigsEdgeNativePoolMachinesListParamsWithContext(h.ctx).
		WithConfigUID(configUID).
		WithMachinePoolName(machinePoolName)
	resp, err := h.Client.V1CloudConfigsEdgeNativePoolMachinesList(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// GetNodeInEdgeNativeMachinePool retrieves specific machine in the edge-native machine pool.
func (h *V1Client) GetNodeInEdgeNativeMachinePool(configUID, machinePoolName, machineID string) (*models.V1EdgeNativeMachine, error) {
	params := clientv1.NewV1CloudConfigsEdgeNativePoolMachinesUIDGetParamsWithContext(h.ctx).
		WithConfigUID(configUID).
		WithMachinePoolName(machinePoolName).WithMachineUID(machineID)
	resp, err := h.Client.V1CloudConfigsEdgeNativePoolMachinesUIDGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// DeleteNodeInEdgeNativeMachinePool delete the specific node in the edge-native machine pool
func (h *V1Client) DeleteNodeInEdgeNativeMachinePool(configUID, machinePoolName, machineID string) error {
	params := clientv1.NewV1CloudConfigsEdgeNativePoolMachinesUIDDeleteParamsWithContext(h.ctx).
		WithConfigUID(configUID).
		WithMachinePoolName(machinePoolName).
		WithMachineUID(machineID)
	_, err := h.Client.V1CloudConfigsEdgeNativePoolMachinesUIDDelete(params)
	if err != nil {
		return err
	}
	return nil
}

// GetMachinesList retrieves a list of machines from a machine pool using the provided function.
func (h *V1Client) GetMachinesList(fn GetMachinesList, configUID, machinePoolName string) (map[string]string, error) {
	return fn(configUID, machinePoolName)
}

// GetMachinesListAws retrieves a list of AWS machines from a machine pool.
// Returns a map where key is the machine name and value is the machine UID.
func (h *V1Client) GetMachinesListAws(configUID, machinePoolName string) (map[string]string, error) {
	params := clientv1.NewV1CloudConfigsAwsPoolMachinesListParamsWithContext(h.ctx).
		WithConfigUID(configUID).
		WithMachinePoolName(machinePoolName)
	mpList, err := h.Client.V1CloudConfigsAwsPoolMachinesList(params)
	if err != nil {
		return nil, err
	}
	machinesMap := make(map[string]string)
	for _, machine := range mpList.Payload.Items {
		if machine.Metadata != nil && machine.Metadata.Name != "" {
			machinesMap[machine.Metadata.Name] = machine.Metadata.UID
		}
	}
	return machinesMap, nil
}

// GetMachinesListAzure retrieves a list of Azure machines from a machine pool.
// Returns a map where key is the machine name and value is the machine UID.
func (h *V1Client) GetMachinesListAzure(configUID, machinePoolName string) (map[string]string, error) {
	params := clientv1.NewV1CloudConfigsAzurePoolMachinesListParamsWithContext(h.ctx).
		WithConfigUID(configUID).
		WithMachinePoolName(machinePoolName)
	mpList, err := h.Client.V1CloudConfigsAzurePoolMachinesList(params)
	if err != nil {
		return nil, err
	}
	machinesMap := make(map[string]string)
	for _, machine := range mpList.Payload.Items {
		if machine.Metadata != nil && machine.Metadata.Name != "" {
			machinesMap[machine.Metadata.Name] = machine.Metadata.UID
		}
	}
	return machinesMap, nil
}

// GetMachinesListGcp retrieves a list of GCP machines from a machine pool.
// Returns a map where key is the machine name and value is the machine UID.
func (h *V1Client) GetMachinesListGcp(configUID, machinePoolName string) (map[string]string, error) {
	params := clientv1.NewV1CloudConfigsGcpPoolMachinesListParamsWithContext(h.ctx).
		WithConfigUID(configUID).
		WithMachinePoolName(machinePoolName)
	mpList, err := h.Client.V1CloudConfigsGcpPoolMachinesList(params)
	if err != nil {
		return nil, err
	}
	machinesMap := make(map[string]string)
	for _, machine := range mpList.Payload.Items {
		if machine.Metadata != nil && machine.Metadata.Name != "" {
			machinesMap[machine.Metadata.Name] = machine.Metadata.UID
		}
	}
	return machinesMap, nil
}

// GetMachinesListVsphere retrieves a list of vSphere machines from a machine pool.
// Returns a map where key is the machine name and value is the machine UID.
func (h *V1Client) GetMachinesListVsphere(configUID, machinePoolName string) (map[string]string, error) {
	params := clientv1.NewV1CloudConfigsVspherePoolMachinesListParamsWithContext(h.ctx).
		WithConfigUID(configUID).
		WithMachinePoolName(machinePoolName)
	mpList, err := h.Client.V1CloudConfigsVspherePoolMachinesList(params)
	if err != nil {
		return nil, err
	}
	machinesMap := make(map[string]string)
	for _, machine := range mpList.Payload.Items {
		if machine.Metadata != nil && machine.Metadata.Name != "" {
			machinesMap[machine.Metadata.Name] = machine.Metadata.UID
		}
	}
	return machinesMap, nil
}

// GetMachinesListGeneric retrieves a list of Generic machines from a machine pool.
// Returns a map where key is the machine name and value is the machine UID.
func (h *V1Client) GetMachinesListGeneric(configUID, machinePoolName string) (map[string]string, error) {
	params := clientv1.NewV1CloudConfigsGenericPoolMachinesListParamsWithContext(h.ctx).
		WithConfigUID(configUID).
		WithMachinePoolName(machinePoolName)
	mpList, err := h.Client.V1CloudConfigsGenericPoolMachinesList(params)
	if err != nil {
		return nil, err
	}
	machinesMap := make(map[string]string)
	for _, machine := range mpList.Payload.Items {
		if machine.Metadata != nil && machine.Metadata.Name != "" {
			machinesMap[machine.Metadata.Name] = machine.Metadata.UID
		}
	}
	return machinesMap, nil
}

// GetMachinesListCloudStack retrieves a list of CloudStack machines from a machine pool.
// Returns a map where key is the machine name and value is the machine UID.
func (h *V1Client) GetMachinesListCloudStack(configUID, machinePoolName string) (map[string]string, error) {
	params := clientv1.NewV1CloudConfigsCloudStackPoolMachinesListParamsWithContext(h.ctx).
		WithConfigUID(configUID).
		WithMachinePoolName(machinePoolName)
	mpList, err := h.Client.V1CloudConfigsCloudStackPoolMachinesList(params)
	if err != nil {
		return nil, err
	}
	machinesMap := make(map[string]string)
	for _, machine := range mpList.Payload.Items {
		if machine.Metadata != nil && machine.Metadata.Name != "" {
			machinesMap[machine.Metadata.Name] = machine.Metadata.UID
		}
	}
	return machinesMap, nil
}
