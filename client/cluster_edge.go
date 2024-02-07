package client

import (
	"errors"

	"github.com/spectrocloud/palette-api-go/apiutil/transport"
	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
)

func (h *V1Client) CreateClusterEdge(cluster *models.V1SpectroEdgeClusterEntity, scope string) (string, error) {
	var params *clientV1.V1SpectroClustersEdgeCreateParams
	switch scope {
	case "project":
		params = clientV1.NewV1SpectroClustersEdgeCreateParamsWithContext(h.Ctx).WithBody(cluster)
	case "tenant":
		params = clientV1.NewV1SpectroClustersEdgeCreateParams().WithBody(cluster)
	}

	success, err := h.Client.V1SpectroClustersEdgeCreate(params)
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

func (h *V1Client) CreateMachinePoolEdge(CloudConfigId string, machinePool *models.V1EdgeMachinePoolConfigEntity, scope string) error {
	var params *clientV1.V1CloudConfigsEdgeMachinePoolCreateParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsEdgeMachinePoolCreateParamsWithContext(h.Ctx).WithConfigUID(CloudConfigId).WithBody(machinePool)
	case "tenant":
		params = clientV1.NewV1CloudConfigsEdgeMachinePoolCreateParams().WithConfigUID(CloudConfigId).WithBody(machinePool)
	}

	_, err := h.Client.V1CloudConfigsEdgeMachinePoolCreate(params)
	return err
}

func (h *V1Client) UpdateMachinePoolEdge(CloudConfigId string, machinePool *models.V1EdgeMachinePoolConfigEntity, scope string) error {
	var params *clientV1.V1CloudConfigsEdgeMachinePoolUpdateParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsEdgeMachinePoolUpdateParamsWithContext(h.Ctx).
			WithConfigUID(CloudConfigId).
			WithMachinePoolName(*machinePool.PoolConfig.Name).
			WithBody(machinePool)
	case "tenant":
		params = clientV1.NewV1CloudConfigsEdgeMachinePoolUpdateParams().
			WithConfigUID(CloudConfigId).
			WithMachinePoolName(*machinePool.PoolConfig.Name).
			WithBody(machinePool)
	}

	_, err := h.Client.V1CloudConfigsEdgeMachinePoolUpdate(params)
	return err
}

func (h *V1Client) DeleteMachinePoolEdge(CloudConfigId, machinePoolName, scope string) error {
	var params *clientV1.V1CloudConfigsEdgeMachinePoolDeleteParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsEdgeMachinePoolDeleteParamsWithContext(h.Ctx).WithConfigUID(CloudConfigId).WithMachinePoolName(machinePoolName)
	case "tenant":
		params = clientV1.NewV1CloudConfigsEdgeMachinePoolDeleteParams().WithConfigUID(CloudConfigId).WithMachinePoolName(machinePoolName)
	}

	_, err := h.Client.V1CloudConfigsEdgeMachinePoolDelete(params)
	return err
}

func (h *V1Client) GetCloudConfigEdge(configUID, scope string) (*models.V1EdgeCloudConfig, error) {
	var params *clientV1.V1CloudConfigsEdgeGetParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsEdgeGetParamsWithContext(h.Ctx).WithConfigUID(configUID)
	case "tenant":
		params = clientV1.NewV1CloudConfigsEdgeGetParams().WithConfigUID(configUID)
	}

	success, err := h.Client.V1CloudConfigsEdgeGet(params)
	var e *transport.TransportError
	if errors.As(err, &e) && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return success.Payload, nil
}

func (h *V1Client) GetNodeStatusMapEdge(configUID, machinePoolName, scope string) (map[string]models.V1CloudMachineStatus, error) {
	var params *clientV1.V1CloudConfigsEdgePoolMachinesListParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsEdgePoolMachinesListParamsWithContext(h.Ctx).WithConfigUID(configUID).WithMachinePoolName(machinePoolName)
	case "tenant":
		params = clientV1.NewV1CloudConfigsEdgePoolMachinesListParams().WithConfigUID(configUID).WithMachinePoolName(machinePoolName)
	}

	mpList, err := h.Client.V1CloudConfigsEdgePoolMachinesList(params)
	nMap := map[string]models.V1CloudMachineStatus{}
	if len(mpList.Payload.Items) > 0 {
		for _, node := range mpList.Payload.Items {
			nMap[node.Metadata.UID] = *node.Status
		}
	}
	return nMap, err
}
