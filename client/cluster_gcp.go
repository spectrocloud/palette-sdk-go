package client

import (
	"github.com/spectrocloud/hapi/apiutil/transport"
	"github.com/spectrocloud/hapi/models"
	clusterC "github.com/spectrocloud/hapi/spectrocluster/client/v1"
)

func (h *V1Client) CreateClusterGcp(cluster *models.V1SpectroGcpClusterEntity, ClusterContext string) (string, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return "", err
	}

	var params *clusterC.V1SpectroClustersGcpCreateParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1SpectroClustersGcpCreateParamsWithContext(h.Ctx).WithBody(cluster)
	case "tenant":
		params = clusterC.NewV1SpectroClustersGcpCreateParams().WithBody(cluster)
	}

	success, err := client.V1SpectroClustersGcpCreate(params)
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

func (h *V1Client) CreateMachinePoolGcp(cloudConfigId, ClusterContext string, machinePool *models.V1GcpMachinePoolConfigEntity) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil
	}

	var params *clusterC.V1CloudConfigsGcpMachinePoolCreateParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsGcpMachinePoolCreateParamsWithContext(h.Ctx).WithConfigUID(cloudConfigId).WithBody(machinePool)
	case "tenant":
		params = clusterC.NewV1CloudConfigsGcpMachinePoolCreateParams().WithConfigUID(cloudConfigId).WithBody(machinePool)
	}

	_, err = client.V1CloudConfigsGcpMachinePoolCreate(params)
	return err
}

func (h *V1Client) UpdateMachinePoolGcp(cloudConfigId, ClusterContext string, machinePool *models.V1GcpMachinePoolConfigEntity) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil
	}

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

	_, err = client.V1CloudConfigsGcpMachinePoolUpdate(params)
	return err
}

func (h *V1Client) DeleteMachinePoolGcp(cloudConfigId, machinePoolName, ClusterContext string) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil
	}

	var params *clusterC.V1CloudConfigsGcpMachinePoolDeleteParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsGcpMachinePoolDeleteParamsWithContext(h.Ctx).WithConfigUID(cloudConfigId).WithMachinePoolName(machinePoolName)
	case "tenant":
		params = clusterC.NewV1CloudConfigsGcpMachinePoolDeleteParams().WithConfigUID(cloudConfigId).WithMachinePoolName(machinePoolName)
	}

	_, err = client.V1CloudConfigsGcpMachinePoolDelete(params)
	return err
}

func (h *V1Client) GetCloudConfigGcp(configUID, ClusterContext string) (*models.V1GcpCloudConfig, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}

	var params *clusterC.V1CloudConfigsGcpGetParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsGcpGetParamsWithContext(h.Ctx).WithConfigUID(configUID)
	case "tenant":
		params = clusterC.NewV1CloudConfigsGcpGetParams().WithConfigUID(configUID)
	}

	success, err := client.V1CloudConfigsGcpGet(params)
	if e, ok := err.(*transport.TransportError); ok && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return success.Payload, nil
}

func (h *V1Client) ImportClusterGcp(meta *models.V1ObjectMetaInputEntity) (string, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return "", err
	}

	params := clusterC.NewV1SpectroClustersGcpImportParamsWithContext(h.Ctx).WithBody(
		&models.V1SpectroGcpClusterImportEntity{
			Metadata: meta,
		},
	)
	success, err := client.V1SpectroClustersGcpImport(params)
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

func (h *V1Client) GetMachineListGcp(configUID string, machinePoolName string, ClusterContext string) ([]*models.V1GcpMachine, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}

	var params *clusterC.V1CloudConfigsGcpPoolMachinesListParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsGcpPoolMachinesListParamsWithContext(h.Ctx).WithConfigUID(configUID).WithMachinePoolName(machinePoolName)
	case "tenant":
		params = clusterC.NewV1CloudConfigsGcpPoolMachinesListParams().WithConfigUID(configUID).WithMachinePoolName(machinePoolName)
	}

	mpList, err := client.V1CloudConfigsGcpPoolMachinesList(params)
	return mpList.Payload.Items, err
}

func (h *V1Client) GetMachinesItemsActionsGcp(configUID string, machinePoolName string, ClusterContext string) (map[string]string, error) {
	mpList, err := h.GetMachineListGcp(configUID, machinePoolName, ClusterContext)
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
