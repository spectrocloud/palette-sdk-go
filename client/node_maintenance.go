package client

import (
	"github.com/spectrocloud/hapi/models"
	clusterC "github.com/spectrocloud/hapi/spectrocluster/client/v1"
)

func (h *V1Client) ToggleMaintenanceOnNode(nodeMaintenance *models.V1MachineMaintenance, ClusterContext string, CloudType string, ConfigUID string, MachineName string, NodeId string) error {
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

func (h *V1Client) GetNodeMaintenanceStatus(ClusterContext string, CloudType string, ConfigUID string, MachineName string, NodeId string) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return err
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
		return err
	}

	return nil
}
