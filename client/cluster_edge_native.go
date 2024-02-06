package client

import (
	"errors"

	"github.com/spectrocloud/palette-api-go/apiutil/transport"
	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
)

func (h *V1Client) CreateClusterEdgeNative(cluster *models.V1SpectroEdgeNativeClusterEntity, scope string) (string, error) {
	var params *clientV1.V1SpectroClustersEdgeNativeCreateParams
	switch scope {
	case "project":
		params = clientV1.NewV1SpectroClustersEdgeNativeCreateParamsWithContext(h.Ctx).WithBody(cluster)
	case "tenant":
		params = clientV1.NewV1SpectroClustersEdgeNativeCreateParams().WithBody(cluster)
	}

	success, err := h.GetClient().V1SpectroClustersEdgeNativeCreate(params)
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

func (h *V1Client) CreateMachinePoolEdgeNative(CloudConfigId, scope string, machinePool *models.V1EdgeNativeMachinePoolConfigEntity) error {
	var params *clientV1.V1CloudConfigsEdgeNativeMachinePoolCreateParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsEdgeNativeMachinePoolCreateParamsWithContext(h.Ctx).WithConfigUID(CloudConfigId).WithBody(machinePool)
	case "tenant":
		params = clientV1.NewV1CloudConfigsEdgeNativeMachinePoolCreateParams().WithConfigUID(CloudConfigId).WithBody(machinePool)
	}

	_, err := h.GetClient().V1CloudConfigsEdgeNativeMachinePoolCreate(params)
	return err
}

func (h *V1Client) UpdateMachinePoolEdgeNative(CloudConfigId, scope string, machinePool *models.V1EdgeNativeMachinePoolConfigEntity) error {
	var params *clientV1.V1CloudConfigsEdgeNativeMachinePoolUpdateParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsEdgeNativeMachinePoolUpdateParamsWithContext(h.Ctx).
			WithConfigUID(CloudConfigId).
			WithMachinePoolName(*machinePool.PoolConfig.Name).
			WithBody(machinePool)
	case "tenant":
		params = clientV1.NewV1CloudConfigsEdgeNativeMachinePoolUpdateParams().
			WithConfigUID(CloudConfigId).
			WithMachinePoolName(*machinePool.PoolConfig.Name).
			WithBody(machinePool)
	}

	_, err := h.GetClient().V1CloudConfigsEdgeNativeMachinePoolUpdate(params)
	return err
}

func (h *V1Client) DeleteMachinePoolEdgeNative(CloudConfigId, machinePoolName, scope string) error {
	var params *clientV1.V1CloudConfigsEdgeNativeMachinePoolDeleteParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsEdgeNativeMachinePoolDeleteParamsWithContext(h.Ctx).WithConfigUID(CloudConfigId).WithMachinePoolName(machinePoolName)
	case "tenant":
		params = clientV1.NewV1CloudConfigsEdgeNativeMachinePoolDeleteParams().WithConfigUID(CloudConfigId).WithMachinePoolName(machinePoolName)
	}

	_, err := h.GetClient().V1CloudConfigsEdgeNativeMachinePoolDelete(params)
	return err
}

func (h *V1Client) GetCloudConfigEdgeNative(configUID, scope string) (*models.V1EdgeNativeCloudConfig, error) {
	var params *clientV1.V1CloudConfigsEdgeNativeGetParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsEdgeNativeGetParamsWithContext(h.Ctx).WithConfigUID(configUID)
	case "tenant":
		params = clientV1.NewV1CloudConfigsEdgeNativeGetParams().WithConfigUID(configUID)
	}

	success, err := h.GetClient().V1CloudConfigsEdgeNativeGet(params)
	var e *transport.TransportError
	if errors.As(err, &e) && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return success.Payload, nil
}

func (h *V1Client) GetNodeStatusMapEdgeNative(configUID, machinePoolName, scope string) (map[string]models.V1CloudMachineStatus, error) {
	var params *clientV1.V1CloudConfigsEdgeNativePoolMachinesListParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsEdgeNativePoolMachinesListParamsWithContext(h.Ctx).WithConfigUID(configUID).WithMachinePoolName(machinePoolName)
	case "tenant":
		params = clientV1.NewV1CloudConfigsEdgeNativePoolMachinesListParams().WithConfigUID(configUID).WithMachinePoolName(machinePoolName)
	}

	mpList, err := h.GetClient().V1CloudConfigsEdgeNativePoolMachinesList(params)
	nMap := map[string]models.V1CloudMachineStatus{}
	if len(mpList.Payload.Items) > 0 {
		for _, node := range mpList.Payload.Items {
			nMap[node.Metadata.UID] = *node.Status
		}
	}
	return nMap, err
}
