package client

import (
	"errors"

	"github.com/spectrocloud/hapi/apiutil/transport"
	"github.com/spectrocloud/hapi/models"
	clusterC "github.com/spectrocloud/hapi/spectrocluster/client/v1"
)

func (h *V1Client) CreateClusterEdgeNative(cluster *models.V1SpectroEdgeNativeClusterEntity, ClusterContext string) (string, error) {
	var params *clusterC.V1SpectroClustersEdgeNativeCreateParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1SpectroClustersEdgeNativeCreateParamsWithContext(h.Ctx).WithBody(cluster)
	case "tenant":
		params = clusterC.NewV1SpectroClustersEdgeNativeCreateParams().WithBody(cluster)
	}

	success, err := h.GetClusterClient().V1SpectroClustersEdgeNativeCreate(params)
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

func (h *V1Client) CreateMachinePoolEdgeNative(cloudConfigId, ClusterContext string, machinePool *models.V1EdgeNativeMachinePoolConfigEntity) error {
	var params *clusterC.V1CloudConfigsEdgeNativeMachinePoolCreateParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsEdgeNativeMachinePoolCreateParamsWithContext(h.Ctx).WithConfigUID(cloudConfigId).WithBody(machinePool)
	case "tenant":
		params = clusterC.NewV1CloudConfigsEdgeNativeMachinePoolCreateParams().WithConfigUID(cloudConfigId).WithBody(machinePool)
	}

	_, err := h.GetClusterClient().V1CloudConfigsEdgeNativeMachinePoolCreate(params)
	return err
}

func (h *V1Client) UpdateMachinePoolEdgeNative(cloudConfigId, ClusterContext string, machinePool *models.V1EdgeNativeMachinePoolConfigEntity) error {
	var params *clusterC.V1CloudConfigsEdgeNativeMachinePoolUpdateParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsEdgeNativeMachinePoolUpdateParamsWithContext(h.Ctx).
			WithConfigUID(cloudConfigId).
			WithMachinePoolName(*machinePool.PoolConfig.Name).
			WithBody(machinePool)
	case "tenant":
		params = clusterC.NewV1CloudConfigsEdgeNativeMachinePoolUpdateParams().
			WithConfigUID(cloudConfigId).
			WithMachinePoolName(*machinePool.PoolConfig.Name).
			WithBody(machinePool)
	}

	_, err := h.GetClusterClient().V1CloudConfigsEdgeNativeMachinePoolUpdate(params)
	return err
}

func (h *V1Client) DeleteMachinePoolEdgeNative(cloudConfigId, machinePoolName, ClusterContext string) error {
	var params *clusterC.V1CloudConfigsEdgeNativeMachinePoolDeleteParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsEdgeNativeMachinePoolDeleteParamsWithContext(h.Ctx).WithConfigUID(cloudConfigId).WithMachinePoolName(machinePoolName)
	case "tenant":
		params = clusterC.NewV1CloudConfigsEdgeNativeMachinePoolDeleteParams().WithConfigUID(cloudConfigId).WithMachinePoolName(machinePoolName)
	}

	_, err := h.GetClusterClient().V1CloudConfigsEdgeNativeMachinePoolDelete(params)
	return err
}

func (h *V1Client) GetCloudConfigEdgeNative(configUID, ClusterContext string) (*models.V1EdgeNativeCloudConfig, error) {
	if h.GetCloudConfigEdgeNativeFn != nil {
		return h.GetCloudConfigEdgeNativeFn(configUID, ClusterContext)
	}
	var params *clusterC.V1CloudConfigsEdgeNativeGetParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsEdgeNativeGetParamsWithContext(h.Ctx).WithConfigUID(configUID)
	case "tenant":
		params = clusterC.NewV1CloudConfigsEdgeNativeGetParams().WithConfigUID(configUID)
	}

	success, err := h.GetClusterClient().V1CloudConfigsEdgeNativeGet(params)
	var e *transport.TransportError
	if errors.As(err, &e) && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return success.Payload, nil
}

func (h *V1Client) GetNodeStatusMapEdgeNative(configUID, machinePoolName, ClusterContext string) (map[string]models.V1CloudMachineStatus, error) {
	var params *clusterC.V1CloudConfigsEdgeNativePoolMachinesListParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsEdgeNativePoolMachinesListParamsWithContext(h.Ctx).WithConfigUID(configUID).WithMachinePoolName(machinePoolName)
	case "tenant":
		params = clusterC.NewV1CloudConfigsEdgeNativePoolMachinesListParams().WithConfigUID(configUID).WithMachinePoolName(machinePoolName)
	}

	mpList, err := h.GetClusterClient().V1CloudConfigsEdgeNativePoolMachinesList(params)
	nMap := map[string]models.V1CloudMachineStatus{}
	if len(mpList.Payload.Items) > 0 {
		for _, node := range mpList.Payload.Items {
			nMap[node.Metadata.UID] = *node.Status
		}
	}
	return nMap, err
}
