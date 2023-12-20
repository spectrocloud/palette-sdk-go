package client

import (
	"errors"

	"github.com/spectrocloud/hapi/apiutil/transport"
	"github.com/spectrocloud/hapi/models"
	clusterC "github.com/spectrocloud/hapi/spectrocluster/client/v1"
)

func (h *V1Client) CreateClusterAzure(cluster *models.V1SpectroAzureClusterEntity, ClusterContext string) (string, error) {
	var params *clusterC.V1SpectroClustersAzureCreateParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1SpectroClustersAzureCreateParamsWithContext(h.Ctx).WithBody(cluster)
	case "tenant":
		params = clusterC.NewV1SpectroClustersAzureCreateParams().WithBody(cluster)
	}

	success, err := h.GetClusterClient().V1SpectroClustersAzureCreate(params)
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

func (h *V1Client) CreateMachinePoolAzure(cloudConfigId, ClusterContext string, machinePool *models.V1AzureMachinePoolConfigEntity) error {
	var params *clusterC.V1CloudConfigsAzureMachinePoolCreateParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsAzureMachinePoolCreateParamsWithContext(h.Ctx).WithConfigUID(cloudConfigId).WithBody(machinePool)
	case "tenant":
		params = clusterC.NewV1CloudConfigsAzureMachinePoolCreateParams().WithConfigUID(cloudConfigId).WithBody(machinePool)
	}

	_, err := h.GetClusterClient().V1CloudConfigsAzureMachinePoolCreate(params)
	return err
}

func (h *V1Client) UpdateMachinePoolAzure(cloudConfigId, ClusterContext string, machinePool *models.V1AzureMachinePoolConfigEntity) error {
	var params *clusterC.V1CloudConfigsAzureMachinePoolUpdateParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsAzureMachinePoolUpdateParamsWithContext(h.Ctx).
			WithConfigUID(cloudConfigId).
			WithMachinePoolName(*machinePool.PoolConfig.Name).
			WithBody(machinePool)
	case "tenant":
		params = clusterC.NewV1CloudConfigsAzureMachinePoolUpdateParamsWithContext(h.Ctx).
			WithConfigUID(cloudConfigId).
			WithMachinePoolName(*machinePool.PoolConfig.Name).
			WithBody(machinePool)
	}

	_, err := h.GetClusterClient().V1CloudConfigsAzureMachinePoolUpdate(params)
	return err
}

func (h *V1Client) DeleteMachinePoolAzure(cloudConfigId, machinePoolName, ClusterContext string) error {
	var params *clusterC.V1CloudConfigsAzureMachinePoolDeleteParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsAzureMachinePoolDeleteParamsWithContext(h.Ctx).WithConfigUID(cloudConfigId).WithMachinePoolName(machinePoolName)
	case "tenant":
		params = clusterC.NewV1CloudConfigsAzureMachinePoolDeleteParams().WithConfigUID(cloudConfigId).WithMachinePoolName(machinePoolName)
	}

	_, err := h.GetClusterClient().V1CloudConfigsAzureMachinePoolDelete(params)
	return err
}

func (h *V1Client) GetCloudConfigAzure(configUID, ClusterContext string) (*models.V1AzureCloudConfig, error) {
	var params *clusterC.V1CloudConfigsAzureGetParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsAzureGetParamsWithContext(h.Ctx).WithConfigUID(configUID)
	case "tenant":
		params = clusterC.NewV1CloudConfigsAzureGetParams().WithConfigUID(configUID)
	}

	success, err := h.GetClusterClient().V1CloudConfigsAzureGet(params)

	var e *transport.TransportError
	if errors.As(err, &e) && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return success.Payload, nil
}

func (h *V1Client) ImportClusterAzure(meta *models.V1ObjectMetaInputEntity) (string, error) {
	params := clusterC.NewV1SpectroClustersAzureImportParamsWithContext(h.Ctx).WithBody(
		&models.V1SpectroAzureClusterImportEntity{
			Metadata: meta,
		},
	)
	success, err := h.GetClusterClient().V1SpectroClustersAzureImport(params)
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

func (h *V1Client) GetNodeStatusMapAzure(configUID, machinePoolName, ClusterContext string) (map[string]models.V1CloudMachineStatus, error) {
	var params *clusterC.V1CloudConfigsAzurePoolMachinesListParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsAzurePoolMachinesListParamsWithContext(h.Ctx).WithConfigUID(configUID).WithMachinePoolName(machinePoolName)
	case "tenant":
		params = clusterC.NewV1CloudConfigsAzurePoolMachinesListParams().WithConfigUID(configUID).WithMachinePoolName(machinePoolName)
	}

	mpList, err := h.GetClusterClient().V1CloudConfigsAzurePoolMachinesList(params)
	nMap := map[string]models.V1CloudMachineStatus{}
	if len(mpList.Payload.Items) > 0 {
		for _, node := range mpList.Payload.Items {
			nMap[node.Metadata.UID] = *node.Status
		}
	}
	return nMap, err
}
