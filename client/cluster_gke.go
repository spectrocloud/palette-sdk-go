package client

import (
	"errors"
	"github.com/spectrocloud/hapi/apiutil/transport"
	"github.com/spectrocloud/hapi/models"
	clusterC "github.com/spectrocloud/hapi/spectrocluster/client/v1"
)

func (h *V1Client) CreateClusterGke(cluster *models.V1SpectroGcpClusterEntity, ClusterContext string) (string, error) {
	var params *clusterC.V1SpectroClustersGkeCreateParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1SpectroClustersGkeCreateParamsWithContext(h.Ctx).WithBody(cluster)
	case "tenant":
		params = clusterC.NewV1SpectroClustersGkeCreateParams().WithBody(cluster)
	}

	success, err := h.GetClusterClient().V1SpectroClustersGkeCreate(params)
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

func (h *V1Client) GetCloudConfigGke(configUID, ClusterContext string) (*models.V1GcpCloudConfig, error) {
	var params *clusterC.V1CloudConfigsGkeGetParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsGkeGetParamsWithContext(h.Ctx).WithConfigUID(configUID)
	case "tenant":
		params = clusterC.NewV1CloudConfigsGkeGetParams().WithConfigUID(configUID)
	}

	success, err := h.GetClusterClient().V1CloudConfigsGkeGet(params)
	var e *transport.TransportError
	if errors.As(err, &e) && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return success.Payload, nil
}

func (h *V1Client) CreateMachinePoolGke(cloudConfigId, ClusterContext string, machinePool *models.V1GcpMachinePoolConfigEntity) error {
	var params *clusterC.V1CloudConfigsGkeMachinePoolCreateParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsGkeMachinePoolCreateParamsWithContext(h.Ctx).WithConfigUID(cloudConfigId).WithBody(machinePool)
	case "tenant":
		params = clusterC.NewV1CloudConfigsGkeMachinePoolCreateParams().WithConfigUID(cloudConfigId).WithBody(machinePool)
	}

	_, err := h.GetClusterClient().V1CloudConfigsGkeMachinePoolCreate(params)
	return err
}

func (h *V1Client) UpdateMachinePoolGke(cloudConfigId, ClusterContext string, machinePool *models.V1GcpMachinePoolConfigEntity) error {
	var params *clusterC.V1CloudConfigsGkeMachinePoolUpdateParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsGkeMachinePoolUpdateParamsWithContext(h.Ctx).
			WithConfigUID(cloudConfigId).
			WithMachinePoolName(*machinePool.PoolConfig.Name).
			WithBody(machinePool)
	case "tenant":
		params = clusterC.NewV1CloudConfigsGkeMachinePoolUpdateParams().
			WithConfigUID(cloudConfigId).
			WithMachinePoolName(*machinePool.PoolConfig.Name).
			WithBody(machinePool)
	}

	_, err := h.GetClusterClient().V1CloudConfigsGkeMachinePoolUpdate(params)
	return err
}

func (h *V1Client) DeleteMachinePoolGke(cloudConfigId, machinePoolName, ClusterContext string) error {
	var params *clusterC.V1CloudConfigsGkeMachinePoolDeleteParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsGkeMachinePoolDeleteParamsWithContext(h.Ctx).WithConfigUID(cloudConfigId).WithMachinePoolName(machinePoolName)
	case "tenant":
		params = clusterC.NewV1CloudConfigsGkeMachinePoolDeleteParams().WithConfigUID(cloudConfigId).WithMachinePoolName(machinePoolName)
	}

	_, err := h.GetClusterClient().V1CloudConfigsGkeMachinePoolDelete(params)
	return err
}
