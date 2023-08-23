package client

import (
	hapitransport "github.com/spectrocloud/hapi/apiutil/transport"
	"github.com/spectrocloud/hapi/models"
	clusterC "github.com/spectrocloud/hapi/spectrocluster/client/v1"
)

// Cluster

func (h *V1Client) CreateClusterVsphere(cluster *models.V1SpectroVsphereClusterEntity, ClusterContext string) (string, error) {
	if h.CreateClusterVsphereFn != nil {
		return h.CreateClusterVsphereFn(cluster)
	}
	client, err := h.GetClusterClient()
	if err != nil {
		return "", err
	}

	var params *clusterC.V1SpectroClustersVsphereCreateParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1SpectroClustersVsphereCreateParamsWithContext(h.Ctx).WithBody(cluster)
	case "tenant":
		params = clusterC.NewV1SpectroClustersVsphereCreateParams().WithBody(cluster)
	}

	success, err := client.V1SpectroClustersVsphereCreate(params)
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

// Machine Pool

func (h *V1Client) CreateMachinePoolVsphere(cloudConfigId, ClusterContext string, machinePool *models.V1VsphereMachinePoolConfigEntity) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return err
	}

	var params *clusterC.V1CloudConfigsVsphereMachinePoolCreateParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsVsphereMachinePoolCreateParamsWithContext(h.Ctx).WithConfigUID(cloudConfigId).WithBody(machinePool)
	case "tenant":
		params = clusterC.NewV1CloudConfigsVsphereMachinePoolCreateParams().WithConfigUID(cloudConfigId).WithBody(machinePool)
	}

	_, err = client.V1CloudConfigsVsphereMachinePoolCreate(params)
	return err
}

func (h *V1Client) UpdateMachinePoolVsphere(cloudConfigId, ClusterContext string, machinePool *models.V1VsphereMachinePoolConfigEntity) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return err
	}

	var params *clusterC.V1CloudConfigsVsphereMachinePoolUpdateParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsVsphereMachinePoolUpdateParamsWithContext(h.Ctx).
			WithConfigUID(cloudConfigId).
			WithMachinePoolName(*machinePool.PoolConfig.Name).
			WithBody(machinePool)
	case "tenant":
		params = clusterC.NewV1CloudConfigsVsphereMachinePoolUpdateParams().
			WithConfigUID(cloudConfigId).
			WithMachinePoolName(*machinePool.PoolConfig.Name).
			WithBody(machinePool)
	}

	_, err = client.V1CloudConfigsVsphereMachinePoolUpdate(params)
	return err
}

func (h *V1Client) DeleteMachinePoolVsphere(cloudConfigId, machinePoolName, ClusterContext string) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return err
	}

	var params *clusterC.V1CloudConfigsVsphereMachinePoolDeleteParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsVsphereMachinePoolDeleteParamsWithContext(h.Ctx).WithConfigUID(cloudConfigId).WithMachinePoolName(machinePoolName)
	case "tenant":
		params = clusterC.NewV1CloudConfigsVsphereMachinePoolDeleteParams().WithConfigUID(cloudConfigId).WithMachinePoolName(machinePoolName)
	}

	_, err = client.V1CloudConfigsVsphereMachinePoolDelete(params)
	return err
}

// Cloud Config

func (h *V1Client) GetCloudConfigVsphere(configUID, ClusterContext string) (*models.V1VsphereCloudConfig, error) {
	if h.GetCloudConfigVsphereFn != nil {
		return h.GetCloudConfigVsphereFn(configUID)
	}
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}

	var params *clusterC.V1CloudConfigsVsphereGetParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsVsphereGetParamsWithContext(h.Ctx).WithConfigUID(configUID)
	case "tenant":
		params = clusterC.NewV1CloudConfigsVsphereGetParams().WithConfigUID(configUID)
	}

	success, err := client.V1CloudConfigsVsphereGet(params)
	if e, ok := err.(*hapitransport.TransportError); ok && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return success.Payload, nil
}

func (h *V1Client) GetCloudConfigVsphereValues(uid, ClusterContext string) (*models.V1VsphereCloudConfig, error) {
	if h.GetCloudConfigVsphereValuesFn != nil {
		return h.GetCloudConfigVsphereValuesFn(uid)
	}
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}

	var params *clusterC.V1CloudConfigsVsphereGetParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsVsphereGetParamsWithContext(h.Ctx).WithConfigUID(uid)
	case "tenant":
		params = clusterC.NewV1CloudConfigsVsphereGetParams().WithConfigUID(uid)
	}

	cloudConfig, err := client.V1CloudConfigsVsphereGet(params)
	if err != nil {
		return nil, err
	}

	return cloudConfig.Payload, nil
}

func (h *V1Client) UpdateCloudConfigVsphereValues(uid, ClusterContext string, clusterConfig *models.V1VsphereCloudClusterConfigEntity) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return err
	}

	var params *clusterC.V1CloudConfigsVsphereUIDClusterConfigParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsVsphereUIDClusterConfigParamsWithContext(h.Ctx).WithConfigUID(uid).WithBody(clusterConfig)
	case "tenant":
		params = clusterC.NewV1CloudConfigsVsphereUIDClusterConfigParams().WithConfigUID(uid).WithBody(clusterConfig)
	}

	_, err = client.V1CloudConfigsVsphereUIDClusterConfig(params)

	return err
}

// Import

func (h *V1Client) ImportClusterVsphere(meta *models.V1ObjectMetaInputEntity) (string, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return "", err
	}

	params := clusterC.NewV1SpectroClustersVsphereImportParamsWithContext(h.Ctx).WithBody(
		&models.V1SpectroVsphereClusterImportEntity{
			Metadata: meta,
		},
	)
	success, err := client.V1SpectroClustersVsphereImport(params)
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

func (h *V1Client) GetMachineListVsphere(configUID string, machinePoolName string, ClusterContext string) ([]*models.V1VsphereMachine, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}

	var params *clusterC.V1CloudConfigsVspherePoolMachinesListParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsVspherePoolMachinesListParamsWithContext(h.Ctx).WithConfigUID(configUID).WithMachinePoolName(machinePoolName)
	case "tenant":
		params = clusterC.NewV1CloudConfigsVspherePoolMachinesListParams().WithConfigUID(configUID).WithMachinePoolName(machinePoolName)
	}

	mpList, err := client.V1CloudConfigsVspherePoolMachinesList(params)
	return mpList.Payload.Items, err
}

func (h *V1Client) GetMachinesItemsActionsVsphere(configUID string, machinePoolName string, ClusterContext string) (map[string]string, error) {
	mpList, err := h.GetMachineListVsphere(configUID, machinePoolName, ClusterContext)
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
