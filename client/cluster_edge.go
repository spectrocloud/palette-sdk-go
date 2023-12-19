package client

import (
	"errors"

	"github.com/spectrocloud/hapi/apiutil/transport"
	"github.com/spectrocloud/hapi/models"
	clusterC "github.com/spectrocloud/hapi/spectrocluster/client/v1"
)

func (h *V1Client) CreateClusterEdge(cluster *models.V1SpectroEdgeClusterEntity, ClusterContext string) (string, error) {
	var params *clusterC.V1SpectroClustersEdgeCreateParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1SpectroClustersEdgeCreateParamsWithContext(h.Ctx).WithBody(cluster)
	case "tenant":
		params = clusterC.NewV1SpectroClustersEdgeCreateParams().WithBody(cluster)
	}

	success, err := h.GetClusterClient().V1SpectroClustersEdgeCreate(params)
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

func (h *V1Client) CreateMachinePoolEdge(cloudConfigId string, machinePool *models.V1EdgeMachinePoolConfigEntity, ClusterContext string) error {
	var params *clusterC.V1CloudConfigsEdgeMachinePoolCreateParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsEdgeMachinePoolCreateParamsWithContext(h.Ctx).WithConfigUID(cloudConfigId).WithBody(machinePool)
	case "tenant":
		params = clusterC.NewV1CloudConfigsEdgeMachinePoolCreateParams().WithConfigUID(cloudConfigId).WithBody(machinePool)
	}

	_, err := h.GetClusterClient().V1CloudConfigsEdgeMachinePoolCreate(params)
	return err
}

func (h *V1Client) UpdateMachinePoolEdge(cloudConfigId string, machinePool *models.V1EdgeMachinePoolConfigEntity, ClusterContext string) error {
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

	_, err := h.GetClusterClient().V1CloudConfigsEdgeMachinePoolUpdate(params)
	return err
}

func (h *V1Client) DeleteMachinePoolEdge(cloudConfigId, machinePoolName, ClusterContext string) error {
	var params *clusterC.V1CloudConfigsEdgeMachinePoolDeleteParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsEdgeMachinePoolDeleteParamsWithContext(h.Ctx).WithConfigUID(cloudConfigId).WithMachinePoolName(machinePoolName)
	case "tenant":
		params = clusterC.NewV1CloudConfigsEdgeMachinePoolDeleteParams().WithConfigUID(cloudConfigId).WithMachinePoolName(machinePoolName)
	}

	_, err := h.GetClusterClient().V1CloudConfigsEdgeMachinePoolDelete(params)
	return err
}

func (h *V1Client) GetCloudConfigEdge(configUID, ClusterContext string) (*models.V1EdgeCloudConfig, error) {
	var params *clusterC.V1CloudConfigsEdgeGetParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsEdgeGetParamsWithContext(h.Ctx).WithConfigUID(configUID)
	case "tenant":
		params = clusterC.NewV1CloudConfigsEdgeGetParams().WithConfigUID(configUID)
	}

	success, err := h.GetClusterClient().V1CloudConfigsEdgeGet(params)
	var e *transport.TransportError
	if errors.As(err, &e) && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return success.Payload, nil
}

func (h *V1Client) GetNodeStatusMapEdge(configUID, machinePoolName, ClusterContext string) (map[string]models.V1CloudMachineStatus, error) {
	var params *clusterC.V1CloudConfigsEdgePoolMachinesListParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsEdgePoolMachinesListParamsWithContext(h.Ctx).WithConfigUID(configUID).WithMachinePoolName(machinePoolName)
	case "tenant":
		params = clusterC.NewV1CloudConfigsEdgePoolMachinesListParams().WithConfigUID(configUID).WithMachinePoolName(machinePoolName)
	}

	mpList, err := h.GetClusterClient().V1CloudConfigsEdgePoolMachinesList(params)
	nMap := map[string]models.V1CloudMachineStatus{}
	if len(mpList.Payload.Items) > 0 {
		for _, node := range mpList.Payload.Items {
			nMap[node.Metadata.UID] = *node.Status
		}
	}
	return nMap, err
}
