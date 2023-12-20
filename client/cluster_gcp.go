package client

import (
	"errors"

	"github.com/spectrocloud/hapi/apiutil/transport"
	"github.com/spectrocloud/hapi/models"
	clusterC "github.com/spectrocloud/hapi/spectrocluster/client/v1"
)

func (h *V1Client) CreateClusterGcp(cluster *models.V1SpectroGcpClusterEntity, ClusterContext string) (string, error) {
	var params *clusterC.V1SpectroClustersGcpCreateParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1SpectroClustersGcpCreateParamsWithContext(h.Ctx).WithBody(cluster)
	case "tenant":
		params = clusterC.NewV1SpectroClustersGcpCreateParams().WithBody(cluster)
	}

	success, err := h.GetClusterClient().V1SpectroClustersGcpCreate(params)
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

func (h *V1Client) CreateMachinePoolGcp(cloudConfigId, ClusterContext string, machinePool *models.V1GcpMachinePoolConfigEntity) error {
	var params *clusterC.V1CloudConfigsGcpMachinePoolCreateParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsGcpMachinePoolCreateParamsWithContext(h.Ctx).WithConfigUID(cloudConfigId).WithBody(machinePool)
	case "tenant":
		params = clusterC.NewV1CloudConfigsGcpMachinePoolCreateParams().WithConfigUID(cloudConfigId).WithBody(machinePool)
	}

	_, err := h.GetClusterClient().V1CloudConfigsGcpMachinePoolCreate(params)
	return err
}

func (h *V1Client) UpdateMachinePoolGcp(cloudConfigId, ClusterContext string, machinePool *models.V1GcpMachinePoolConfigEntity) error {
	var params *clusterC.V1CloudConfigsGcpMachinePoolUpdateParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsGcpMachinePoolUpdateParamsWithContext(h.Ctx).
			WithConfigUID(cloudConfigId).
			WithMachinePoolName(*machinePool.PoolConfig.Name).
			WithBody(machinePool)
	case "tenant":
		params = clusterC.NewV1CloudConfigsGcpMachinePoolUpdateParams().
			WithConfigUID(cloudConfigId).
			WithMachinePoolName(*machinePool.PoolConfig.Name).
			WithBody(machinePool)
	}

	_, err := h.GetClusterClient().V1CloudConfigsGcpMachinePoolUpdate(params)
	return err
}

func (h *V1Client) DeleteMachinePoolGcp(cloudConfigId, machinePoolName, ClusterContext string) error {
	var params *clusterC.V1CloudConfigsGcpMachinePoolDeleteParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsGcpMachinePoolDeleteParamsWithContext(h.Ctx).WithConfigUID(cloudConfigId).WithMachinePoolName(machinePoolName)
	case "tenant":
		params = clusterC.NewV1CloudConfigsGcpMachinePoolDeleteParams().WithConfigUID(cloudConfigId).WithMachinePoolName(machinePoolName)
	}

	_, err := h.GetClusterClient().V1CloudConfigsGcpMachinePoolDelete(params)
	return err
}

func (h *V1Client) GetCloudConfigGcp(configUID, ClusterContext string) (*models.V1GcpCloudConfig, error) {
	var params *clusterC.V1CloudConfigsGcpGetParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsGcpGetParamsWithContext(h.Ctx).WithConfigUID(configUID)
	case "tenant":
		params = clusterC.NewV1CloudConfigsGcpGetParams().WithConfigUID(configUID)
	}

	success, err := h.GetClusterClient().V1CloudConfigsGcpGet(params)
	var e *transport.TransportError
	if errors.As(err, &e) && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return success.Payload, nil
}

func (h *V1Client) ImportClusterGcp(meta *models.V1ObjectMetaInputEntity) (string, error) {
	params := clusterC.NewV1SpectroClustersGcpImportParamsWithContext(h.Ctx).WithBody(
		&models.V1SpectroGcpClusterImportEntity{
			Metadata: meta,
		},
	)
	success, err := h.GetClusterClient().V1SpectroClustersGcpImport(params)
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

func (h *V1Client) GetNodeStatusMapGcp(configUID, machinePoolName, ClusterContext string) (map[string]models.V1CloudMachineStatus, error) {
	var params *clusterC.V1CloudConfigsGcpPoolMachinesListParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsGcpPoolMachinesListParamsWithContext(h.Ctx).WithConfigUID(configUID).WithMachinePoolName(machinePoolName)
	case "tenant":
		params = clusterC.NewV1CloudConfigsGcpPoolMachinesListParams().WithConfigUID(configUID).WithMachinePoolName(machinePoolName)
	}

	mpList, err := h.GetClusterClient().V1CloudConfigsGcpPoolMachinesList(params)
	nMap := map[string]models.V1CloudMachineStatus{}
	if len(mpList.Payload.Items) > 0 {
		for _, node := range mpList.Payload.Items {
			nMap[node.Metadata.UID] = *node.Status
		}
	}
	return nMap, err
}
