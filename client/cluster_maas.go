package client

import (
	"github.com/spectrocloud/hapi/apiutil/transport"
	"github.com/spectrocloud/hapi/models"
	clusterC "github.com/spectrocloud/hapi/spectrocluster/client/v1"
)

// Cluster

func (h *V1Client) CreateClusterMaas(cluster *models.V1SpectroMaasClusterEntity, ClusterContext string) (string, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return "", err
	}

	var params *clusterC.V1SpectroClustersMaasCreateParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1SpectroClustersMaasCreateParamsWithContext(h.Ctx).WithBody(cluster)
	case "tenant":
		params = clusterC.NewV1SpectroClustersMaasCreateParams().WithBody(cluster)
	}

	success, err := client.V1SpectroClustersMaasCreate(params)
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

// Machine Pool

func (h *V1Client) CreateMachinePoolMaas(cloudConfigId, ClusterContext string, machinePool *models.V1MaasMachinePoolConfigEntity) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil
	}

	var params *clusterC.V1CloudConfigsMaasMachinePoolCreateParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsMaasMachinePoolCreateParamsWithContext(h.Ctx).WithConfigUID(cloudConfigId).WithBody(machinePool)
	case "tenant":
		params = clusterC.NewV1CloudConfigsMaasMachinePoolCreateParams().WithConfigUID(cloudConfigId).WithBody(machinePool)
	}

	_, err = client.V1CloudConfigsMaasMachinePoolCreate(params)
	return err
}

func (h *V1Client) UpdateMachinePoolMaas(cloudConfigId, ClusterContext string, machinePool *models.V1MaasMachinePoolConfigEntity) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil
	}

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

	_, err = client.V1CloudConfigsMaasMachinePoolUpdate(params)
	return err
}

func (h *V1Client) DeleteMachinePoolMaas(cloudConfigId, machinePoolName, ClusterContext string) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil
	}

	var params *clusterC.V1CloudConfigsMaasMachinePoolDeleteParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsMaasMachinePoolDeleteParamsWithContext(h.Ctx).WithConfigUID(cloudConfigId).WithMachinePoolName(machinePoolName)
	case "tenant":
		params = clusterC.NewV1CloudConfigsMaasMachinePoolDeleteParams().WithConfigUID(cloudConfigId).WithMachinePoolName(machinePoolName)
	}

	_, err = client.V1CloudConfigsMaasMachinePoolDelete(params)
	return err
}

// Cloud Config

func (h *V1Client) GetCloudConfigMaas(configUID, ClusterContext string) (*models.V1MaasCloudConfig, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}

	var params *clusterC.V1CloudConfigsMaasGetParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsMaasGetParamsWithContext(h.Ctx).WithConfigUID(configUID)
	case "tenant":
		params = clusterC.NewV1CloudConfigsMaasGetParams().WithConfigUID(configUID)
	}

	success, err := client.V1CloudConfigsMaasGet(params)
	if e, ok := err.(*transport.TransportError); ok && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return success.Payload, nil
}

// Import

func (h *V1Client) ImportClusterMaas(meta *models.V1ObjectMetaInputEntity) (string, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return "", err
	}

	params := clusterC.NewV1SpectroClustersMaasImportParamsWithContext(h.Ctx).WithBody(
		&models.V1SpectroMaasClusterImportEntity{
			Metadata: meta,
		},
	)
	success, err := client.V1SpectroClustersMaasImport(params)
	if err != nil {
		return "", err
	}
	return *success.Payload.UID, nil
}

func (h *V1Client) GetMachineListMaas(configUID string, machinePoolName string, ClusterContext string) ([]*models.V1MaasMachine, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}

	var params *clusterC.V1CloudConfigsMaasPoolMachinesListParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsMaasPoolMachinesListParamsWithContext(h.Ctx).WithConfigUID(configUID).WithMachinePoolName(machinePoolName)
	case "tenant":
		params = clusterC.NewV1CloudConfigsMaasPoolMachinesListParams().WithConfigUID(configUID).WithMachinePoolName(machinePoolName)
	}

	mpList, err := client.V1CloudConfigsMaasPoolMachinesList(params)
	return mpList.Payload.Items, err
}

func (h *V1Client) GetMachinesItemsActionsMaas(configUID string, machinePoolName string, ClusterContext string) (map[string]string, error) {
	mpList, err := h.GetMachineListMaas(configUID, machinePoolName, ClusterContext)
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
