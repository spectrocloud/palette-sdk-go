package client

import (
	"github.com/spectrocloud/hapi/models"
	clusterC "github.com/spectrocloud/hapi/spectrocluster/client/v1"

	"github.com/spectrocloud/palette-sdk-go/client/herr"
)

// Cluster

func (h *V1Client) CreateClusterOpenStack(cluster *models.V1SpectroOpenStackClusterEntity, ClusterContext string) (string, error) {
	var params *clusterC.V1SpectroClustersOpenStackCreateParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1SpectroClustersOpenStackCreateParamsWithContext(h.Ctx).WithBody(cluster)
	case "tenant":
		params = clusterC.NewV1SpectroClustersOpenStackCreateParams().WithBody(cluster)
	}

	success, err := h.GetClusterClient().V1SpectroClustersOpenStackCreate(params)
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

// Machine Pool

func (h *V1Client) CreateMachinePoolOpenStack(cloudConfigId, ClusterContext string, machinePool *models.V1OpenStackMachinePoolConfigEntity) error {
	var params *clusterC.V1CloudConfigsOpenStackMachinePoolCreateParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsOpenStackMachinePoolCreateParamsWithContext(h.Ctx).WithConfigUID(cloudConfigId).WithBody(machinePool)
	case "tenant":
		params = clusterC.NewV1CloudConfigsOpenStackMachinePoolCreateParams().WithConfigUID(cloudConfigId).WithBody(machinePool)
	}

	_, err := h.GetClusterClient().V1CloudConfigsOpenStackMachinePoolCreate(params)
	return err
}

func (h *V1Client) UpdateMachinePoolOpenStack(cloudConfigId, ClusterContext string, machinePool *models.V1OpenStackMachinePoolConfigEntity) error {
	var params *clusterC.V1CloudConfigsOpenStackMachinePoolUpdateParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsOpenStackMachinePoolUpdateParamsWithContext(h.Ctx).
			WithConfigUID(cloudConfigId).
			WithMachinePoolName(*machinePool.PoolConfig.Name).
			WithBody(machinePool)
	case "tenant":
		params = clusterC.NewV1CloudConfigsOpenStackMachinePoolUpdateParams().
			WithConfigUID(cloudConfigId).
			WithMachinePoolName(*machinePool.PoolConfig.Name).
			WithBody(machinePool)
	}

	_, err := h.GetClusterClient().V1CloudConfigsOpenStackMachinePoolUpdate(params)
	return err
}

func (h *V1Client) DeleteMachinePoolOpenStack(cloudConfigId, machinePoolName, ClusterContext string) error {
	var params *clusterC.V1CloudConfigsOpenStackMachinePoolDeleteParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsOpenStackMachinePoolDeleteParamsWithContext(h.Ctx).WithConfigUID(cloudConfigId).WithMachinePoolName(machinePoolName)
	case "tenant":
		params = clusterC.NewV1CloudConfigsOpenStackMachinePoolDeleteParams().WithConfigUID(cloudConfigId).WithMachinePoolName(machinePoolName)
	}

	_, err := h.GetClusterClient().V1CloudConfigsOpenStackMachinePoolDelete(params)
	return err
}

// Cloud Config

func (h *V1Client) GetCloudConfigOpenStack(configUID, ClusterContext string) (*models.V1OpenStackCloudConfig, error) {
	var params *clusterC.V1CloudConfigsOpenStackGetParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsOpenStackGetParamsWithContext(h.Ctx).WithConfigUID(configUID)
	case "tenant":
		params = clusterC.NewV1CloudConfigsOpenStackGetParams().WithConfigUID(configUID)
	}

	success, err := h.GetClusterClient().V1CloudConfigsOpenStackGet(params)

	if herr.IsNotFound(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return success.Payload, nil
}

// Import

func (h *V1Client) ImportClusterOpenStack(meta *models.V1ObjectMetaInputEntity) (string, error) {
	params := clusterC.NewV1SpectroClustersOpenStackImportParamsWithContext(h.Ctx).WithBody(
		&models.V1SpectroOpenStackClusterImportEntity{
			Metadata: meta,
		},
	)
	success, err := h.GetClusterClient().V1SpectroClustersOpenStackImport(params)
	if err != nil {
		return "", err
	}
	return *success.Payload.UID, nil
}

func (h *V1Client) GetNodeStatusMapOpenStack(configUID, machinePoolName, ClusterContext string) (map[string]models.V1CloudMachineStatus, error) {
	var params *clusterC.V1CloudConfigsOpenStackPoolMachinesListParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsOpenStackPoolMachinesListParamsWithContext(h.Ctx).WithConfigUID(configUID).WithMachinePoolName(machinePoolName)
	case "tenant":
		params = clusterC.NewV1CloudConfigsOpenStackPoolMachinesListParams().WithConfigUID(configUID).WithMachinePoolName(machinePoolName)
	}

	mpList, err := h.GetClusterClient().V1CloudConfigsOpenStackPoolMachinesList(params)
	nMap := map[string]models.V1CloudMachineStatus{}
	if len(mpList.Payload.Items) > 0 {
		for _, node := range mpList.Payload.Items {
			nMap[node.Metadata.UID] = *node.Status
		}
	}
	return nMap, err
}
