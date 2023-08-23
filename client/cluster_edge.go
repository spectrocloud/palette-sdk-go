package client

import (
	"github.com/spectrocloud/hapi/apiutil/transport"
	"github.com/spectrocloud/hapi/models"
	clusterC "github.com/spectrocloud/hapi/spectrocluster/client/v1"
)

func (h *V1Client) CreateClusterEdge(cluster *models.V1SpectroEdgeClusterEntity, ClusterContext string) (string, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return "", err
	}

	var params *clusterC.V1SpectroClustersEdgeCreateParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1SpectroClustersEdgeCreateParamsWithContext(h.Ctx).WithBody(cluster)
	case "tenant":
		params = clusterC.NewV1SpectroClustersEdgeCreateParams().WithBody(cluster)
	}

	success, err := client.V1SpectroClustersEdgeCreate(params)
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

func (h *V1Client) CreateMachinePoolEdge(cloudConfigId string, machinePool *models.V1EdgeMachinePoolConfigEntity, ClusterContext string) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil
	}

	var params *clusterC.V1CloudConfigsEdgeMachinePoolCreateParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsEdgeMachinePoolCreateParamsWithContext(h.Ctx).WithConfigUID(cloudConfigId).WithBody(machinePool)
	case "tenant":
		params = clusterC.NewV1CloudConfigsEdgeMachinePoolCreateParams().WithConfigUID(cloudConfigId).WithBody(machinePool)
	}

	_, err = client.V1CloudConfigsEdgeMachinePoolCreate(params)
	return err
}

func (h *V1Client) UpdateMachinePoolEdge(cloudConfigId string, machinePool *models.V1EdgeMachinePoolConfigEntity, ClusterContext string) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil
	}

	var params *clusterC.V1CloudConfigsEdgeMachinePoolUpdateParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsEdgeMachinePoolUpdateParamsWithContext(h.Ctx).
			WithConfigUID(cloudConfigId).
			WithMachinePoolName(*machinePool.PoolConfig.Name).
			WithBody(machinePool)
	case "tenant":
		params = clusterC.NewV1CloudConfigsEdgeMachinePoolUpdateParams().
			WithConfigUID(cloudConfigId).
			WithMachinePoolName(*machinePool.PoolConfig.Name).
			WithBody(machinePool)
	}

	_, err = client.V1CloudConfigsEdgeMachinePoolUpdate(params)
	return err
}

func (h *V1Client) DeleteMachinePoolEdge(cloudConfigId, machinePoolName, ClusterContext string) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil
	}

	var params *clusterC.V1CloudConfigsEdgeMachinePoolDeleteParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsEdgeMachinePoolDeleteParamsWithContext(h.Ctx).WithConfigUID(cloudConfigId).WithMachinePoolName(machinePoolName)
	case "tenant":
		params = clusterC.NewV1CloudConfigsEdgeMachinePoolDeleteParams().WithConfigUID(cloudConfigId).WithMachinePoolName(machinePoolName)
	}

	_, err = client.V1CloudConfigsEdgeMachinePoolDelete(params)
	return err
}

func (h *V1Client) GetCloudConfigEdge(configUID, ClusterContext string) (*models.V1EdgeCloudConfig, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}

	var params *clusterC.V1CloudConfigsEdgeGetParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsEdgeGetParamsWithContext(h.Ctx).WithConfigUID(configUID)
	case "tenant":
		params = clusterC.NewV1CloudConfigsEdgeGetParams().WithConfigUID(configUID)
	}

	success, err := client.V1CloudConfigsEdgeGet(params)
	if e, ok := err.(*transport.TransportError); ok && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return success.Payload, nil
}

func (h *V1Client) GetMachineListEdge(configUID string, machinePoolName string, ClusterContext string) ([]*models.V1EdgeMachine, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}

	var params *clusterC.V1CloudConfigsEdgePoolMachinesListParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsEdgePoolMachinesListParamsWithContext(h.Ctx).WithConfigUID(configUID).WithMachinePoolName(machinePoolName)
	case "tenant":
		params = clusterC.NewV1CloudConfigsEdgePoolMachinesListParams().WithConfigUID(configUID).WithMachinePoolName(machinePoolName)
	}

	mpList, err := client.V1CloudConfigsEdgePoolMachinesList(params)
	return mpList.Payload.Items, err
}

func (h *V1Client) GetMachinesItemsActionsEdge(configUID string, machinePoolName string, ClusterContext string) (map[string]string, error) {
	mpList, err := h.GetMachineListEdge(configUID, machinePoolName, ClusterContext)
	nMap := map[string]string{}
	if len(mpList) > 0 {
		for _, node := range mpList {
			if node.Status.MaintenanceStatus.Action != "" {
				nMap[node.Metadata.UID] = node.Status.MaintenanceStatus.Action
			}
		}
	}
	return nMap, err
}
