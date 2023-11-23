package client

import (
	"errors"

	"github.com/spectrocloud/hapi/apiutil/transport"
	"github.com/spectrocloud/hapi/models"
	clusterC "github.com/spectrocloud/hapi/spectrocluster/client/v1"
)

func (h *V1Client) CreateClusterCoxEdge(cluster *models.V1SpectroCoxEdgeClusterEntity, ClusterContext string) (string, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return "", err
	}

	var params *clusterC.V1SpectroClustersCoxEdgeCreateParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1SpectroClustersCoxEdgeCreateParamsWithContext(h.Ctx).WithBody(cluster)
	case "tenant":
		params = clusterC.NewV1SpectroClustersCoxEdgeCreateParams().WithBody(cluster)
	}

	success, err := client.V1SpectroClustersCoxEdgeCreate(params)
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

func (h *V1Client) CreateMachinePoolCoxEdge(cloudConfigId string, machinePool *models.V1CoxEdgeMachinePoolConfigEntity, ClusterContext string) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil
	}

	var params *clusterC.V1CloudConfigsCoxEdgeMachinePoolCreateParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsCoxEdgeMachinePoolCreateParamsWithContext(h.Ctx).WithConfigUID(cloudConfigId).WithBody(machinePool)
	case "tenant":
		params = clusterC.NewV1CloudConfigsCoxEdgeMachinePoolCreateParams().WithConfigUID(cloudConfigId).WithBody(machinePool)
	}

	_, err = client.V1CloudConfigsCoxEdgeMachinePoolCreate(params)
	return err
}

func (h *V1Client) UpdateMachinePoolCoxEdge(cloudConfigId string, machinePool *models.V1CoxEdgeMachinePoolConfigEntity, ClusterContext string) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil
	}

	var params *clusterC.V1CloudConfigsCoxEdgeMachinePoolUpdateParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsCoxEdgeMachinePoolUpdateParamsWithContext(h.Ctx).
			WithConfigUID(cloudConfigId).
			WithMachinePoolName(*machinePool.PoolConfig.Name).
			WithBody(machinePool)
	case "tenant":
		params = clusterC.NewV1CloudConfigsCoxEdgeMachinePoolUpdateParams().
			WithConfigUID(cloudConfigId).
			WithMachinePoolName(*machinePool.PoolConfig.Name).
			WithBody(machinePool)
	}

	_, err = client.V1CloudConfigsCoxEdgeMachinePoolUpdate(params)
	return err
}

func (h *V1Client) DeleteMachinePoolCoxEdge(cloudConfigId, machinePoolName, ClusterContext string) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil
	}

	var params *clusterC.V1CloudConfigsCoxEdgeMachinePoolDeleteParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsCoxEdgeMachinePoolDeleteParamsWithContext(h.Ctx).WithConfigUID(cloudConfigId).WithMachinePoolName(machinePoolName)
	case "tenant":
		params = clusterC.NewV1CloudConfigsCoxEdgeMachinePoolDeleteParams().WithConfigUID(cloudConfigId).WithMachinePoolName(machinePoolName)
	}

	_, err = client.V1CloudConfigsCoxEdgeMachinePoolDelete(params)
	return err
}

func (h *V1Client) GetCloudConfigCoxEdge(configUID, ClusterContext string) (*models.V1CoxEdgeCloudConfig, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}

	var params *clusterC.V1CloudConfigsCoxEdgeGetParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsCoxEdgeGetParamsWithContext(h.Ctx).WithConfigUID(configUID)
	case "tenant":
		params = clusterC.NewV1CloudConfigsCoxEdgeGetParams().WithConfigUID(configUID)
	}

	success, err := client.V1CloudConfigsCoxEdgeGet(params)
	var e *transport.TransportError
	if errors.As(err, &e) && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return success.Payload, nil
}

func (h *V1Client) GetNodeStatusMapCoxEdge(configUID, machinePoolName, ClusterContext string) (map[string]models.V1CloudMachineStatus, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}

	var params *clusterC.V1CloudConfigsCoxEdgePoolMachinesListParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsCoxEdgePoolMachinesListParamsWithContext(h.Ctx).WithConfigUID(configUID).WithMachinePoolName(machinePoolName)
	case "tenant":
		params = clusterC.NewV1CloudConfigsCoxEdgePoolMachinesListParams().WithConfigUID(configUID).WithMachinePoolName(machinePoolName)
	}

	mpList, err := client.V1CloudConfigsCoxEdgePoolMachinesList(params)
	nMap := map[string]models.V1CloudMachineStatus{}
	if len(mpList.Payload.Items) > 0 {
		for _, node := range mpList.Payload.Items {
			nMap[node.Metadata.UID] = *node.Status
		}
	}
	return nMap, err
}
