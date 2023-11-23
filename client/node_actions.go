package client

import (
	"github.com/spectrocloud/hapi/models"
	clusterC "github.com/spectrocloud/hapi/spectrocluster/client/v1"
)

type GetMaintenanceStatus func(string, string, string, string) (*models.V1MachineMaintenanceStatus, error)

func (h *V1Client) ToggleMaintenanceOnNode(nodeMaintenance *models.V1MachineMaintenance, CloudType, ClusterContext, ConfigUID, MachineName, NodeId string) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return err
	}

	var params *clusterC.V1CloudConfigsMachinePoolsMachineUIDMaintenanceUpdateParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsMachinePoolsMachineUIDMaintenanceUpdateParamsWithContext(h.Ctx).WithBody(nodeMaintenance)
	case "tenant":
		params = clusterC.NewV1CloudConfigsMachinePoolsMachineUIDMaintenanceUpdateParams().WithBody(nodeMaintenance)
	}
	params.WithCloudType(CloudType)
	params.WithConfigUID(ConfigUID)
	params.WithMachinePoolName(MachineName)
	params.WithMachineUID(NodeId)
	_, err = client.V1CloudConfigsMachinePoolsMachineUIDMaintenanceUpdate(params)
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

func (h *V1Client) GetNodeMaintenanceStatus(fn GetMaintenanceStatus, ClusterContext, ConfigUID, MachineName, NodeId string) (*models.V1MachineMaintenanceStatus, error) {
	return fn(ClusterContext, ConfigUID, MachineName, NodeId)
}

func (h *V1Client) GetNodeMaintenanceStatusAws(ClusterContext, ConfigUID, MachineName, NodeId string) (*models.V1MachineMaintenanceStatus, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}

	var params *clusterC.V1CloudConfigsAwsPoolMachinesUIDGetParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsAwsPoolMachinesUIDGetParamsWithContext(h.Ctx)
	case "tenant":
		params = clusterC.NewV1CloudConfigsAwsPoolMachinesUIDGetParams()
	}
	params.WithConfigUID(ConfigUID)
	params.WithMachinePoolName(MachineName)
	params.WithMachineUID(NodeId)
	s, err := client.V1CloudConfigsAwsPoolMachinesUIDGet(params)
	print(s)
	if err != nil {
		return nil, err
	}

	return s.Payload.Status.MaintenanceStatus, nil
}

func (h *V1Client) GetNodeMaintenanceStatusMaas(ClusterContext, ConfigUID, MachineName, NodeId string) (*models.V1MachineMaintenanceStatus, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}

	var params *clusterC.V1CloudConfigsMaasPoolMachinesUIDGetParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsMaasPoolMachinesUIDGetParamsWithContext(h.Ctx)
	case "tenant":
		params = clusterC.NewV1CloudConfigsMaasPoolMachinesUIDGetParams()
	}
	params.WithConfigUID(ConfigUID)
	params.WithMachinePoolName(MachineName)
	params.WithMachineUID(NodeId)
	s, err := client.V1CloudConfigsMaasPoolMachinesUIDGet(params)
	print(s)
	if err != nil {
		return nil, err
	}

	return s.Payload.Status.MaintenanceStatus, nil
}

func (h *V1Client) GetNodeMaintenanceStatusAks(ClusterContext, ConfigUID, MachineName, NodeId string) (*models.V1MachineMaintenanceStatus, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}

	var params *clusterC.V1CloudConfigsAksPoolMachinesUIDGetParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsAksPoolMachinesUIDGetParamsWithContext(h.Ctx)
	case "tenant":
		params = clusterC.NewV1CloudConfigsAksPoolMachinesUIDGetParams()
	}
	params.WithConfigUID(ConfigUID)
	params.WithMachinePoolName(MachineName)
	params.WithMachineUID(NodeId)
	s, err := client.V1CloudConfigsAksPoolMachinesUIDGet(params)
	print(s)
	if err != nil {
		return nil, err
	}

	return s.Payload.Status.MaintenanceStatus, nil
}

func (h *V1Client) GetNodeMaintenanceStatusAzure(ClusterContext, ConfigUID, MachineName, NodeId string) (*models.V1MachineMaintenanceStatus, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}

	var params *clusterC.V1CloudConfigsAzurePoolMachinesUIDGetParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsAzurePoolMachinesUIDGetParamsWithContext(h.Ctx)
	case "tenant":
		params = clusterC.NewV1CloudConfigsAzurePoolMachinesUIDGetParams()
	}
	params.WithConfigUID(ConfigUID)
	params.WithMachinePoolName(MachineName)
	params.WithMachineUID(NodeId)
	s, err := client.V1CloudConfigsAzurePoolMachinesUIDGet(params)
	print(s)
	if err != nil {
		return nil, err
	}

	return s.Payload.Status.MaintenanceStatus, nil
}

func (h *V1Client) GetNodeMaintenanceStatusCoxEdge(ClusterContext, ConfigUID, MachineName, NodeId string) (*models.V1MachineMaintenanceStatus, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}

	var params *clusterC.V1CloudConfigsCoxEdgePoolMachinesUIDGetParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsCoxEdgePoolMachinesUIDGetParamsWithContext(h.Ctx)
	case "tenant":
		params = clusterC.NewV1CloudConfigsCoxEdgePoolMachinesUIDGetParams()
	}
	params.WithConfigUID(ConfigUID)
	params.WithMachinePoolName(MachineName)
	params.WithMachineUID(NodeId)
	s, err := client.V1CloudConfigsCoxEdgePoolMachinesUIDGet(params)
	print(s)
	if err != nil {
		return nil, err
	}

	return s.Payload.Status.MaintenanceStatus, nil
}

func (h *V1Client) GetNodeMaintenanceStatusEdgeNative(ClusterContext, ConfigUID, MachineName, NodeId string) (*models.V1MachineMaintenanceStatus, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}

	var params *clusterC.V1CloudConfigsEdgeNativePoolMachinesUIDGetParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsEdgeNativePoolMachinesUIDGetParamsWithContext(h.Ctx)
	case "tenant":
		params = clusterC.NewV1CloudConfigsEdgeNativePoolMachinesUIDGetParams()
	}
	params.WithConfigUID(ConfigUID)
	params.WithMachinePoolName(MachineName)
	params.WithMachineUID(NodeId)
	s, err := client.V1CloudConfigsEdgeNativePoolMachinesUIDGet(params)
	print(s)
	if err != nil {
		return nil, err
	}

	return s.Payload.Status.MaintenanceStatus, nil
}

func (h *V1Client) GetNodeMaintenanceStatusEdge(ClusterContext, ConfigUID, MachineName, NodeId string) (*models.V1MachineMaintenanceStatus, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}

	var params *clusterC.V1CloudConfigsEdgePoolMachinesUIDGetParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsEdgePoolMachinesUIDGetParamsWithContext(h.Ctx)
	case "tenant":
		params = clusterC.NewV1CloudConfigsEdgePoolMachinesUIDGetParams()
	}
	params.WithConfigUID(ConfigUID)
	params.WithMachinePoolName(MachineName)
	params.WithMachineUID(NodeId)
	s, err := client.V1CloudConfigsEdgePoolMachinesUIDGet(params)
	print(s)
	if err != nil {
		return nil, err
	}

	return s.Payload.Status.MaintenanceStatus, nil
}

func (h *V1Client) GetNodeMaintenanceStatusEdgeVsphere(ClusterContext, ConfigUID, MachineName, NodeId string) (*models.V1MachineMaintenanceStatus, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}

	var params *clusterC.V1CloudConfigsVspherePoolMachinesUIDGetParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsVspherePoolMachinesUIDGetParamsWithContext(h.Ctx)
	case "tenant":
		params = clusterC.NewV1CloudConfigsVspherePoolMachinesUIDGetParams()
	}
	params.WithConfigUID(ConfigUID)
	params.WithMachinePoolName(MachineName)
	params.WithMachineUID(NodeId)
	s, err := client.V1CloudConfigsVspherePoolMachinesUIDGet(params)
	print(s)
	if err != nil {
		return nil, err
	}

	return s.Payload.Status.MaintenanceStatus, nil
}

func (h *V1Client) GetNodeMaintenanceStatusEks(ClusterContext, ConfigUID, MachineName, NodeId string) (*models.V1MachineMaintenanceStatus, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}

	var params *clusterC.V1CloudConfigsEksPoolMachinesUIDGetParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsEksPoolMachinesUIDGetParamsWithContext(h.Ctx)
	case "tenant":
		params = clusterC.NewV1CloudConfigsEksPoolMachinesUIDGetParams()
	}
	params.WithConfigUID(ConfigUID)
	params.WithMachinePoolName(MachineName)
	params.WithMachineUID(NodeId)
	s, err := client.V1CloudConfigsEksPoolMachinesUIDGet(params)
	print(s)
	if err != nil {
		return nil, err
	}

	return s.Payload.Status.MaintenanceStatus, nil
}

func (h *V1Client) GetNodeMaintenanceStatusGcp(ClusterContext, ConfigUID, MachineName, NodeId string) (*models.V1MachineMaintenanceStatus, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}

	var params *clusterC.V1CloudConfigsGcpPoolMachinesUIDGetParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsGcpPoolMachinesUIDGetParamsWithContext(h.Ctx)
	case "tenant":
		params = clusterC.NewV1CloudConfigsGcpPoolMachinesUIDGetParams()
	}
	params.WithConfigUID(ConfigUID)
	params.WithMachinePoolName(MachineName)
	params.WithMachineUID(NodeId)
	s, err := client.V1CloudConfigsGcpPoolMachinesUIDGet(params)
	print(s)
	if err != nil {
		return nil, err
	}

	return s.Payload.Status.MaintenanceStatus, nil
}

func (h *V1Client) GetNodeMaintenanceStatusGeneric(ClusterContext, ConfigUID, MachineName, NodeId string) (*models.V1MachineMaintenanceStatus, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}

	var params *clusterC.V1CloudConfigsGenericPoolMachinesUIDGetParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsGenericPoolMachinesUIDGetParamsWithContext(h.Ctx)
	case "tenant":
		params = clusterC.NewV1CloudConfigsGenericPoolMachinesUIDGetParams()
	}
	params.WithConfigUID(ConfigUID)
	params.WithMachinePoolName(MachineName)
	params.WithMachineUID(NodeId)
	s, err := client.V1CloudConfigsGenericPoolMachinesUIDGet(params)
	print(s)
	if err != nil {
		return nil, err
	}

	return s.Payload.Status.MaintenanceStatus, nil
}

func (h *V1Client) GetNodeMaintenanceStatusGke(ClusterContext, ConfigUID, MachineName, NodeId string) (*models.V1MachineMaintenanceStatus, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}

	var params *clusterC.V1CloudConfigsGkePoolMachinesUIDGetParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsGkePoolMachinesUIDGetParamsWithContext(h.Ctx)
	case "tenant":
		params = clusterC.NewV1CloudConfigsGkePoolMachinesUIDGetParams()
	}
	params.WithConfigUID(ConfigUID)
	params.WithMachinePoolName(MachineName)
	params.WithMachineUID(NodeId)
	s, err := client.V1CloudConfigsGkePoolMachinesUIDGet(params)
	print(s)
	if err != nil {
		return nil, err
	}

	return s.Payload.Status.MaintenanceStatus, nil
}

func (h *V1Client) GetNodeMaintenanceStatusLibvirt(ClusterContext, ConfigUID, MachineName, NodeId string) (*models.V1MachineMaintenanceStatus, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}

	var params *clusterC.V1CloudConfigsLibvirtPoolMachinesUIDGetParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsLibvirtPoolMachinesUIDGetParamsWithContext(h.Ctx)
	case "tenant":
		params = clusterC.NewV1CloudConfigsLibvirtPoolMachinesUIDGetParams()
	}
	params.WithConfigUID(ConfigUID)
	params.WithMachinePoolName(MachineName)
	params.WithMachineUID(NodeId)
	s, err := client.V1CloudConfigsLibvirtPoolMachinesUIDGet(params)
	print(s)
	if err != nil {
		return nil, err
	}

	return s.Payload.Status.MaintenanceStatus, nil
}

func (h *V1Client) GetNodeMaintenanceStatusOpenStack(ClusterContext, ConfigUID, MachineName, NodeId string) (*models.V1MachineMaintenanceStatus, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}

	var params *clusterC.V1CloudConfigsOpenStackPoolMachinesUIDGetParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsOpenStackPoolMachinesUIDGetParamsWithContext(h.Ctx)
	case "tenant":
		params = clusterC.NewV1CloudConfigsOpenStackPoolMachinesUIDGetParams()
	}
	params.WithConfigUID(ConfigUID)
	params.WithMachinePoolName(MachineName)
	params.WithMachineUID(NodeId)
	s, err := client.V1CloudConfigsOpenStackPoolMachinesUIDGet(params)
	print(s)
	if err != nil {
		return nil, err
	}

	return s.Payload.Status.MaintenanceStatus, nil
}

func (h *V1Client) GetNodeMaintenanceStatusTke(ClusterContext, ConfigUID, MachineName, NodeId string) (*models.V1MachineMaintenanceStatus, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}

	var params *clusterC.V1CloudConfigsTkePoolMachinesUIDGetParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsTkePoolMachinesUIDGetParamsWithContext(h.Ctx)
	case "tenant":
		params = clusterC.NewV1CloudConfigsTkePoolMachinesUIDGetParams()
	}
	params.WithConfigUID(ConfigUID)
	params.WithMachinePoolName(MachineName)
	params.WithMachineUID(NodeId)
	s, err := client.V1CloudConfigsTkePoolMachinesUIDGet(params)
	print(s)
	if err != nil {
		return nil, err
	}

	return s.Payload.Status.MaintenanceStatus, nil
}

func (h *V1Client) GetNodeVirtualMaintenanceStatusVirtual(ClusterContext, ConfigUID, MachineName, NodeId string) (*models.V1MachineMaintenanceStatus, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}

	var params *clusterC.V1CloudConfigsVirtualPoolMachinesUIDGetParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsVirtualPoolMachinesUIDGetParamsWithContext(h.Ctx)
	case "tenant":
		params = clusterC.NewV1CloudConfigsVirtualPoolMachinesUIDGetParams()
	}
	params.WithConfigUID(ConfigUID)
	params.WithMachinePoolName(MachineName)
	params.WithMachineUID(NodeId)
	s, err := client.V1CloudConfigsVirtualPoolMachinesUIDGet(params)
	print(s)
	if err != nil {
		return nil, err
	}

	return s.Payload.Status.MaintenanceStatus, nil
}

func (h *V1Client) GetNodeMaintenanceStatusVsphere(ClusterContext, ConfigUID, MachineName, NodeId string) (*models.V1MachineMaintenanceStatus, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}

	var params *clusterC.V1CloudConfigsVspherePoolMachinesUIDGetParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsVspherePoolMachinesUIDGetParamsWithContext(h.Ctx)
	case "tenant":
		params = clusterC.NewV1CloudConfigsVspherePoolMachinesUIDGetParams()
	}
	params.WithConfigUID(ConfigUID)
	params.WithMachinePoolName(MachineName)
	params.WithMachineUID(NodeId)
	s, err := client.V1CloudConfigsVspherePoolMachinesUIDGet(params)
	print(s)
	if err != nil {
		return nil, err
	}

	return s.Payload.Status.MaintenanceStatus, nil
}
