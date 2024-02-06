package client

import (
	"errors"

	"github.com/spectrocloud/palette-api-go/apiutil/transport"
	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
)

func (h *V1Client) CreateclientV1oxEdge(cluster *models.V1SpectroCoxEdgeClusterEntity, scope string) (string, error) {
	var params *clientV1.V1SpectroClustersCoxEdgeCreateParams
	switch scope {
	case "project":
		params = clientV1.NewV1SpectroClustersCoxEdgeCreateParamsWithContext(h.Ctx).WithBody(cluster)
	case "tenant":
		params = clientV1.NewV1SpectroClustersCoxEdgeCreateParams().WithBody(cluster)
	}

	success, err := h.GetClient().V1SpectroClustersCoxEdgeCreate(params)
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

func (h *V1Client) CreateMachinePoolCoxEdge(CloudConfigId string, machinePool *models.V1CoxEdgeMachinePoolConfigEntity, scope string) error {
	var params *clientV1.V1CloudConfigsCoxEdgeMachinePoolCreateParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsCoxEdgeMachinePoolCreateParamsWithContext(h.Ctx).WithConfigUID(CloudConfigId).WithBody(machinePool)
	case "tenant":
		params = clientV1.NewV1CloudConfigsCoxEdgeMachinePoolCreateParams().WithConfigUID(CloudConfigId).WithBody(machinePool)
	}

	_, err := h.GetClient().V1CloudConfigsCoxEdgeMachinePoolCreate(params)
	return err
}

func (h *V1Client) UpdateMachinePoolCoxEdge(CloudConfigId string, machinePool *models.V1CoxEdgeMachinePoolConfigEntity, scope string) error {
	var params *clientV1.V1CloudConfigsCoxEdgeMachinePoolUpdateParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsCoxEdgeMachinePoolUpdateParamsWithContext(h.Ctx).
			WithConfigUID(CloudConfigId).
			WithMachinePoolName(*machinePool.PoolConfig.Name).
			WithBody(machinePool)
	case "tenant":
		params = clientV1.NewV1CloudConfigsCoxEdgeMachinePoolUpdateParams().
			WithConfigUID(CloudConfigId).
			WithMachinePoolName(*machinePool.PoolConfig.Name).
			WithBody(machinePool)
	}

	_, err := h.GetClient().V1CloudConfigsCoxEdgeMachinePoolUpdate(params)
	return err
}

func (h *V1Client) DeleteMachinePoolCoxEdge(CloudConfigId, machinePoolName, scope string) error {
	var params *clientV1.V1CloudConfigsCoxEdgeMachinePoolDeleteParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsCoxEdgeMachinePoolDeleteParamsWithContext(h.Ctx).WithConfigUID(CloudConfigId).WithMachinePoolName(machinePoolName)
	case "tenant":
		params = clientV1.NewV1CloudConfigsCoxEdgeMachinePoolDeleteParams().WithConfigUID(CloudConfigId).WithMachinePoolName(machinePoolName)
	}

	_, err := h.GetClient().V1CloudConfigsCoxEdgeMachinePoolDelete(params)
	return err
}

func (h *V1Client) GetCloudConfigCoxEdge(configUID, scope string) (*models.V1CoxEdgeCloudConfig, error) {
	var params *clientV1.V1CloudConfigsCoxEdgeGetParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsCoxEdgeGetParamsWithContext(h.Ctx).WithConfigUID(configUID)
	case "tenant":
		params = clientV1.NewV1CloudConfigsCoxEdgeGetParams().WithConfigUID(configUID)
	}

	success, err := h.GetClient().V1CloudConfigsCoxEdgeGet(params)
	var e *transport.TransportError
	if errors.As(err, &e) && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return success.Payload, nil
}

func (h *V1Client) GetNodeStatusMapCoxEdge(configUID, machinePoolName, scope string) (map[string]models.V1CloudMachineStatus, error) {
	var params *clientV1.V1CloudConfigsCoxEdgePoolMachinesListParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsCoxEdgePoolMachinesListParamsWithContext(h.Ctx).WithConfigUID(configUID).WithMachinePoolName(machinePoolName)
	case "tenant":
		params = clientV1.NewV1CloudConfigsCoxEdgePoolMachinesListParams().WithConfigUID(configUID).WithMachinePoolName(machinePoolName)
	}

	mpList, err := h.GetClient().V1CloudConfigsCoxEdgePoolMachinesList(params)
	nMap := map[string]models.V1CloudMachineStatus{}
	if len(mpList.Payload.Items) > 0 {
		for _, node := range mpList.Payload.Items {
			nMap[node.Metadata.UID] = *node.Status
		}
	}
	return nMap, err
}
