package client

import (
	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
)

type GetMaintenanceStatus func(string, string, string, string) (*models.V1MachineMaintenanceStatus, error)

func (h *V1Client) ToggleMaintenanceOnNode(nodeMaintenance *models.V1MachineMaintenance, CloudType, scope, ConfigUID, MachineName, NodeId string) error {
	var params *clientV1.V1CloudConfigsMachinePoolsMachineUIDMaintenanceUpdateParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsMachinePoolsMachineUIDMaintenanceUpdateParamsWithContext(h.Ctx).WithBody(nodeMaintenance)
	case "tenant":
		params = clientV1.NewV1CloudConfigsMachinePoolsMachineUIDMaintenanceUpdateParams().WithBody(nodeMaintenance)
	}
	params.WithCloudType(CloudType)
	params.WithConfigUID(ConfigUID)
	params.WithMachinePoolName(MachineName)
	params.WithMachineUID(NodeId)

	_, err := h.GetClient().V1CloudConfigsMachinePoolsMachineUIDMaintenanceUpdate(params)
	if err != nil {
		return err
	}

	return nil
}

func (h *V1Client) GetNodeValue(nodeId, action string) map[string]interface{} {
	return map[string]interface{}{
		"node_id": nodeId,
		"action":  action,
	}
}

func (h *V1Client) GetNodeMaintenanceStatus(fn GetMaintenanceStatus, scope, ConfigUID, MachineName, NodeId string) (*models.V1MachineMaintenanceStatus, error) {
	return fn(scope, ConfigUID, MachineName, NodeId)
}

func (h *V1Client) GetNodeMaintenanceStatusAws(scope, ConfigUID, MachineName, NodeId string) (*models.V1MachineMaintenanceStatus, error) {
	var params *clientV1.V1CloudConfigsAwsPoolMachinesUIDGetParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsAwsPoolMachinesUIDGetParamsWithContext(h.Ctx)
	case "tenant":
		params = clientV1.NewV1CloudConfigsAwsPoolMachinesUIDGetParams()
	}
	params.WithConfigUID(ConfigUID)
	params.WithMachinePoolName(MachineName)
	params.WithMachineUID(NodeId)

	s, err := h.GetClient().V1CloudConfigsAwsPoolMachinesUIDGet(params)
	print(s)
	if err != nil {
		return nil, err
	}

	return s.Payload.Status.MaintenanceStatus, nil
}

func (h *V1Client) GetNodeMaintenanceStatusMaas(scope, ConfigUID, MachineName, NodeId string) (*models.V1MachineMaintenanceStatus, error) {
	var params *clientV1.V1CloudConfigsMaasPoolMachinesUIDGetParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsMaasPoolMachinesUIDGetParamsWithContext(h.Ctx)
	case "tenant":
		params = clientV1.NewV1CloudConfigsMaasPoolMachinesUIDGetParams()
	}
	params.WithConfigUID(ConfigUID)
	params.WithMachinePoolName(MachineName)
	params.WithMachineUID(NodeId)

	s, err := h.GetClient().V1CloudConfigsMaasPoolMachinesUIDGet(params)
	print(s)
	if err != nil {
		return nil, err
	}

	return s.Payload.Status.MaintenanceStatus, nil
}

func (h *V1Client) GetNodeMaintenanceStatusAks(scope, ConfigUID, MachineName, NodeId string) (*models.V1MachineMaintenanceStatus, error) {
	var params *clientV1.V1CloudConfigsAksPoolMachinesUIDGetParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsAksPoolMachinesUIDGetParamsWithContext(h.Ctx)
	case "tenant":
		params = clientV1.NewV1CloudConfigsAksPoolMachinesUIDGetParams()
	}
	params.WithConfigUID(ConfigUID)
	params.WithMachinePoolName(MachineName)
	params.WithMachineUID(NodeId)

	s, err := h.GetClient().V1CloudConfigsAksPoolMachinesUIDGet(params)
	print(s)
	if err != nil {
		return nil, err
	}

	return s.Payload.Status.MaintenanceStatus, nil
}

func (h *V1Client) GetNodeMaintenanceStatusAzure(scope, ConfigUID, MachineName, NodeId string) (*models.V1MachineMaintenanceStatus, error) {
	var params *clientV1.V1CloudConfigsAzurePoolMachinesUIDGetParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsAzurePoolMachinesUIDGetParamsWithContext(h.Ctx)
	case "tenant":
		params = clientV1.NewV1CloudConfigsAzurePoolMachinesUIDGetParams()
	}
	params.WithConfigUID(ConfigUID)
	params.WithMachinePoolName(MachineName)
	params.WithMachineUID(NodeId)

	s, err := h.GetClient().V1CloudConfigsAzurePoolMachinesUIDGet(params)
	print(s)
	if err != nil {
		return nil, err
	}

	return s.Payload.Status.MaintenanceStatus, nil
}

func (h *V1Client) GetNodeMaintenanceStatusCoxEdge(scope, ConfigUID, MachineName, NodeId string) (*models.V1MachineMaintenanceStatus, error) {
	var params *clientV1.V1CloudConfigsCoxEdgePoolMachinesUIDGetParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsCoxEdgePoolMachinesUIDGetParamsWithContext(h.Ctx)
	case "tenant":
		params = clientV1.NewV1CloudConfigsCoxEdgePoolMachinesUIDGetParams()
	}
	params.WithConfigUID(ConfigUID)
	params.WithMachinePoolName(MachineName)
	params.WithMachineUID(NodeId)

	s, err := h.GetClient().V1CloudConfigsCoxEdgePoolMachinesUIDGet(params)
	print(s)
	if err != nil {
		return nil, err
	}

	return s.Payload.Status.MaintenanceStatus, nil
}

func (h *V1Client) GetNodeMaintenanceStatusEdgeNative(scope, ConfigUID, MachineName, NodeId string) (*models.V1MachineMaintenanceStatus, error) {
	var params *clientV1.V1CloudConfigsEdgeNativePoolMachinesUIDGetParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsEdgeNativePoolMachinesUIDGetParamsWithContext(h.Ctx)
	case "tenant":
		params = clientV1.NewV1CloudConfigsEdgeNativePoolMachinesUIDGetParams()
	}
	params.WithConfigUID(ConfigUID)
	params.WithMachinePoolName(MachineName)
	params.WithMachineUID(NodeId)

	s, err := h.GetClient().V1CloudConfigsEdgeNativePoolMachinesUIDGet(params)
	print(s)
	if err != nil {
		return nil, err
	}

	return s.Payload.Status.MaintenanceStatus, nil
}

func (h *V1Client) GetNodeMaintenanceStatusEdge(scope, ConfigUID, MachineName, NodeId string) (*models.V1MachineMaintenanceStatus, error) {
	var params *clientV1.V1CloudConfigsEdgePoolMachinesUIDGetParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsEdgePoolMachinesUIDGetParamsWithContext(h.Ctx)
	case "tenant":
		params = clientV1.NewV1CloudConfigsEdgePoolMachinesUIDGetParams()
	}
	params.WithConfigUID(ConfigUID)
	params.WithMachinePoolName(MachineName)
	params.WithMachineUID(NodeId)

	s, err := h.GetClient().V1CloudConfigsEdgePoolMachinesUIDGet(params)
	print(s)
	if err != nil {
		return nil, err
	}

	return s.Payload.Status.MaintenanceStatus, nil
}

func (h *V1Client) GetNodeMaintenanceStatusEdgeVsphere(scope, ConfigUID, MachineName, NodeId string) (*models.V1MachineMaintenanceStatus, error) {
	var params *clientV1.V1CloudConfigsVspherePoolMachinesUIDGetParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsVspherePoolMachinesUIDGetParamsWithContext(h.Ctx)
	case "tenant":
		params = clientV1.NewV1CloudConfigsVspherePoolMachinesUIDGetParams()
	}
	params.WithConfigUID(ConfigUID)
	params.WithMachinePoolName(MachineName)
	params.WithMachineUID(NodeId)

	s, err := h.GetClient().V1CloudConfigsVspherePoolMachinesUIDGet(params)
	print(s)
	if err != nil {
		return nil, err
	}

	return s.Payload.Status.MaintenanceStatus, nil
}

func (h *V1Client) GetNodeMaintenanceStatusEks(scope, ConfigUID, MachineName, NodeId string) (*models.V1MachineMaintenanceStatus, error) {
	var params *clientV1.V1CloudConfigsEksPoolMachinesUIDGetParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsEksPoolMachinesUIDGetParamsWithContext(h.Ctx)
	case "tenant":
		params = clientV1.NewV1CloudConfigsEksPoolMachinesUIDGetParams()
	}
	params.WithConfigUID(ConfigUID)
	params.WithMachinePoolName(MachineName)
	params.WithMachineUID(NodeId)

	s, err := h.GetClient().V1CloudConfigsEksPoolMachinesUIDGet(params)
	print(s)
	if err != nil {
		return nil, err
	}

	return s.Payload.Status.MaintenanceStatus, nil
}

func (h *V1Client) GetNodeMaintenanceStatusGcp(scope, ConfigUID, MachineName, NodeId string) (*models.V1MachineMaintenanceStatus, error) {
	var params *clientV1.V1CloudConfigsGcpPoolMachinesUIDGetParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsGcpPoolMachinesUIDGetParamsWithContext(h.Ctx)
	case "tenant":
		params = clientV1.NewV1CloudConfigsGcpPoolMachinesUIDGetParams()
	}
	params.WithConfigUID(ConfigUID)
	params.WithMachinePoolName(MachineName)
	params.WithMachineUID(NodeId)

	s, err := h.GetClient().V1CloudConfigsGcpPoolMachinesUIDGet(params)
	print(s)
	if err != nil {
		return nil, err
	}

	return s.Payload.Status.MaintenanceStatus, nil
}

func (h *V1Client) GetNodeMaintenanceStatusGeneric(scope, ConfigUID, MachineName, NodeId string) (*models.V1MachineMaintenanceStatus, error) {
	var params *clientV1.V1CloudConfigsGenericPoolMachinesUIDGetParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsGenericPoolMachinesUIDGetParamsWithContext(h.Ctx)
	case "tenant":
		params = clientV1.NewV1CloudConfigsGenericPoolMachinesUIDGetParams()
	}
	params.WithConfigUID(ConfigUID)
	params.WithMachinePoolName(MachineName)
	params.WithMachineUID(NodeId)

	s, err := h.GetClient().V1CloudConfigsGenericPoolMachinesUIDGet(params)
	print(s)
	if err != nil {
		return nil, err
	}

	return s.Payload.Status.MaintenanceStatus, nil
}

func (h *V1Client) GetNodeMaintenanceStatusGke(scope, ConfigUID, MachineName, NodeId string) (*models.V1MachineMaintenanceStatus, error) {
	var params *clientV1.V1CloudConfigsGkePoolMachinesUIDGetParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsGkePoolMachinesUIDGetParamsWithContext(h.Ctx)
	case "tenant":
		params = clientV1.NewV1CloudConfigsGkePoolMachinesUIDGetParams()
	}
	params.WithConfigUID(ConfigUID)
	params.WithMachinePoolName(MachineName)
	params.WithMachineUID(NodeId)

	s, err := h.GetClient().V1CloudConfigsGkePoolMachinesUIDGet(params)
	print(s)
	if err != nil {
		return nil, err
	}

	return s.Payload.Status.MaintenanceStatus, nil
}

func (h *V1Client) GetNodeMaintenanceStatusLibvirt(scope, ConfigUID, MachineName, NodeId string) (*models.V1MachineMaintenanceStatus, error) {
	var params *clientV1.V1CloudConfigsLibvirtPoolMachinesUIDGetParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsLibvirtPoolMachinesUIDGetParamsWithContext(h.Ctx)
	case "tenant":
		params = clientV1.NewV1CloudConfigsLibvirtPoolMachinesUIDGetParams()
	}
	params.WithConfigUID(ConfigUID)
	params.WithMachinePoolName(MachineName)
	params.WithMachineUID(NodeId)

	s, err := h.GetClient().V1CloudConfigsLibvirtPoolMachinesUIDGet(params)
	print(s)
	if err != nil {
		return nil, err
	}

	return s.Payload.Status.MaintenanceStatus, nil
}

func (h *V1Client) GetNodeMaintenanceStatusOpenStack(scope, ConfigUID, MachineName, NodeId string) (*models.V1MachineMaintenanceStatus, error) {
	var params *clientV1.V1CloudConfigsOpenStackPoolMachinesUIDGetParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsOpenStackPoolMachinesUIDGetParamsWithContext(h.Ctx)
	case "tenant":
		params = clientV1.NewV1CloudConfigsOpenStackPoolMachinesUIDGetParams()
	}
	params.WithConfigUID(ConfigUID)
	params.WithMachinePoolName(MachineName)
	params.WithMachineUID(NodeId)

	s, err := h.GetClient().V1CloudConfigsOpenStackPoolMachinesUIDGet(params)
	print(s)
	if err != nil {
		return nil, err
	}

	return s.Payload.Status.MaintenanceStatus, nil
}

func (h *V1Client) GetNodeMaintenanceStatusTke(scope, ConfigUID, MachineName, NodeId string) (*models.V1MachineMaintenanceStatus, error) {
	var params *clientV1.V1CloudConfigsTkePoolMachinesUIDGetParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsTkePoolMachinesUIDGetParamsWithContext(h.Ctx)
	case "tenant":
		params = clientV1.NewV1CloudConfigsTkePoolMachinesUIDGetParams()
	}
	params.WithConfigUID(ConfigUID)
	params.WithMachinePoolName(MachineName)
	params.WithMachineUID(NodeId)

	s, err := h.GetClient().V1CloudConfigsTkePoolMachinesUIDGet(params)
	print(s)
	if err != nil {
		return nil, err
	}

	return s.Payload.Status.MaintenanceStatus, nil
}

func (h *V1Client) GetNodeVirtualMaintenanceStatusVirtual(scope, ConfigUID, MachineName, NodeId string) (*models.V1MachineMaintenanceStatus, error) {
	var params *clientV1.V1CloudConfigsVirtualPoolMachinesUIDGetParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsVirtualPoolMachinesUIDGetParamsWithContext(h.Ctx)
	case "tenant":
		params = clientV1.NewV1CloudConfigsVirtualPoolMachinesUIDGetParams()
	}
	params.WithConfigUID(ConfigUID)
	params.WithMachinePoolName(MachineName)
	params.WithMachineUID(NodeId)

	s, err := h.GetClient().V1CloudConfigsVirtualPoolMachinesUIDGet(params)
	print(s)
	if err != nil {
		return nil, err
	}

	return s.Payload.Status.MaintenanceStatus, nil
}

func (h *V1Client) GetNodeMaintenanceStatusVsphere(scope, ConfigUID, MachineName, NodeId string) (*models.V1MachineMaintenanceStatus, error) {
	var params *clientV1.V1CloudConfigsVspherePoolMachinesUIDGetParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsVspherePoolMachinesUIDGetParamsWithContext(h.Ctx)
	case "tenant":
		params = clientV1.NewV1CloudConfigsVspherePoolMachinesUIDGetParams()
	}
	params.WithConfigUID(ConfigUID)
	params.WithMachinePoolName(MachineName)
	params.WithMachineUID(NodeId)

	s, err := h.GetClient().V1CloudConfigsVspherePoolMachinesUIDGet(params)
	print(s)
	if err != nil {
		return nil, err
	}

	return s.Payload.Status.MaintenanceStatus, nil
}
