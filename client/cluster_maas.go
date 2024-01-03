package client

import (
	"errors"

	"github.com/spectrocloud/hapi/apiutil/transport"
	"github.com/spectrocloud/hapi/models"
	clusterC "github.com/spectrocloud/hapi/spectrocluster/client/v1"
)

// Cluster

func (h *V1Client) CreateClusterMaas(cluster *models.V1SpectroMaasClusterEntity, ClusterContext string) (string, error) {
	var params *clusterC.V1SpectroClustersMaasCreateParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1SpectroClustersMaasCreateParamsWithContext(h.Ctx).WithBody(cluster)
	case "tenant":
		params = clusterC.NewV1SpectroClustersMaasCreateParams().WithBody(cluster)
	}

	success, err := h.GetClusterClient().V1SpectroClustersMaasCreate(params)
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

// Machine Pool

func (h *V1Client) CreateMachinePoolMaas(cloudConfigId, ClusterContext string, machinePool *models.V1MaasMachinePoolConfigEntity) error {
	var params *clusterC.V1CloudConfigsMaasMachinePoolCreateParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsMaasMachinePoolCreateParamsWithContext(h.Ctx).WithConfigUID(cloudConfigId).WithBody(machinePool)
	case "tenant":
		params = clusterC.NewV1CloudConfigsMaasMachinePoolCreateParams().WithConfigUID(cloudConfigId).WithBody(machinePool)
	}

	_, err := h.GetClusterClient().V1CloudConfigsMaasMachinePoolCreate(params)
	return err
}

func (h *V1Client) UpdateMachinePoolMaas(cloudConfigId, ClusterContext string, machinePool *models.V1MaasMachinePoolConfigEntity) error {
	var params *clusterC.V1CloudConfigsMaasMachinePoolUpdateParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsMaasMachinePoolUpdateParamsWithContext(h.Ctx).
			WithConfigUID(cloudConfigId).
			WithMachinePoolName(*machinePool.PoolConfig.Name).
			WithBody(machinePool)
	case "tenant":
		params = clusterC.NewV1CloudConfigsMaasMachinePoolUpdateParams().
			WithConfigUID(cloudConfigId).
			WithMachinePoolName(*machinePool.PoolConfig.Name).
			WithBody(machinePool)
	}

	_, err := h.GetClusterClient().V1CloudConfigsMaasMachinePoolUpdate(params)
	return err
}

func (h *V1Client) DeleteMachinePoolMaas(cloudConfigId, machinePoolName, ClusterContext string) error {
	var params *clusterC.V1CloudConfigsMaasMachinePoolDeleteParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsMaasMachinePoolDeleteParamsWithContext(h.Ctx).WithConfigUID(cloudConfigId).WithMachinePoolName(machinePoolName)
	case "tenant":
		params = clusterC.NewV1CloudConfigsMaasMachinePoolDeleteParams().WithConfigUID(cloudConfigId).WithMachinePoolName(machinePoolName)
	}

	_, err := h.GetClusterClient().V1CloudConfigsMaasMachinePoolDelete(params)
	return err
}

// Cloud Config

func (h *V1Client) GetCloudConfigMaas(configUID, ClusterContext string) (*models.V1MaasCloudConfig, error) {
	var params *clusterC.V1CloudConfigsMaasGetParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsMaasGetParamsWithContext(h.Ctx).WithConfigUID(configUID)
	case "tenant":
		params = clusterC.NewV1CloudConfigsMaasGetParams().WithConfigUID(configUID)
	}

	success, err := h.GetClusterClient().V1CloudConfigsMaasGet(params)
	var e *transport.TransportError
	if errors.As(err, &e) && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return success.Payload, nil
}

// Import

func (h *V1Client) ImportClusterMaas(meta *models.V1ObjectMetaInputEntity) (string, error) {
	params := clusterC.NewV1SpectroClustersMaasImportParamsWithContext(h.Ctx).WithBody(
		&models.V1SpectroMaasClusterImportEntity{
			Metadata: meta,
		},
	)
	success, err := h.GetClusterClient().V1SpectroClustersMaasImport(params)
	if err != nil {
		return "", err
	}
	return *success.Payload.UID, nil
}

func (h *V1Client) GetNodeStatusMapMaas(configUID, machinePoolName, ClusterContext string) (map[string]models.V1CloudMachineStatus, error) {
	var params *clusterC.V1CloudConfigsMaasPoolMachinesListParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsMaasPoolMachinesListParamsWithContext(h.Ctx).WithConfigUID(configUID).WithMachinePoolName(machinePoolName)
	case "tenant":
		params = clusterC.NewV1CloudConfigsMaasPoolMachinesListParams().WithConfigUID(configUID).WithMachinePoolName(machinePoolName)
	}

	mpList, err := h.GetClusterClient().V1CloudConfigsMaasPoolMachinesList(params)
	nMap := map[string]models.V1CloudMachineStatus{}
	if len(mpList.Payload.Items) > 0 {
		for _, node := range mpList.Payload.Items {
			nMap[node.Metadata.UID] = *node.Status
		}
	}
	return nMap, err
}
