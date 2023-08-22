package client

import (
	"github.com/spectrocloud/hapi/models"
	clusterC "github.com/spectrocloud/hapi/spectrocluster/client/v1"
)

func (h *V1Client) ToggleMaintenanceOnNode(nodeMaintenance *models.V1MachineMaintenance, CloudType string, ClusterContext string, ConfigUID string, MachineName string, NodeId string) error {
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

func (h *V1Client) GetNodeMaintenanceStatus(CloudType string, ClusterContext string, ConfigUID string, MachineName string, NodeId string) (*models.V1MachineMaintenanceStatus, error) {
	switch CloudType {

	case "aws":
		nodeMaintenanceStatus, err := h.GetAWSNodeMaintenanceStatus(ClusterContext, ConfigUID, MachineName, NodeId)
		return nodeMaintenanceStatus, err
	case "maas":
		nodeMaintenanceStatus, err := h.GetMAASNodeMaintenanceStatus(ClusterContext, ConfigUID, MachineName, NodeId)
		return nodeMaintenanceStatus, err
	}
	return nil, nil
}

func (h *V1Client) GetAWSNodeMaintenanceStatus(ClusterContext string, ConfigUID string, MachineName string, NodeId string) (*models.V1MachineMaintenanceStatus, error) {
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

func (h *V1Client) GetMAASNodeMaintenanceStatus(ClusterContext string, ConfigUID string, MachineName string, NodeId string) (*models.V1MachineMaintenanceStatus, error) {
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
