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

func (h *V1Client) CreateMachinePoolVsphere(cloudConfigId string, machinePool *models.V1VsphereMachinePoolConfigEntity) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return err
	}

	params := clusterC.NewV1CloudConfigsVsphereMachinePoolCreateParamsWithContext(h.Ctx).WithConfigUID(cloudConfigId).WithBody(machinePool)
	_, err = client.V1CloudConfigsVsphereMachinePoolCreate(params)
	return err
}

func (h *V1Client) UpdateMachinePoolVsphere(cloudConfigId string, machinePool *models.V1VsphereMachinePoolConfigEntity) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return err
	}

	params := clusterC.NewV1CloudConfigsVsphereMachinePoolUpdateParamsWithContext(h.Ctx).
		WithConfigUID(cloudConfigId).
		WithMachinePoolName(*machinePool.PoolConfig.Name).
		WithBody(machinePool)
	_, err = client.V1CloudConfigsVsphereMachinePoolUpdate(params)
	return err
}

func (h *V1Client) DeleteMachinePoolVsphere(cloudConfigId, machinePoolName string) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return err
	}

	params := clusterC.NewV1CloudConfigsVsphereMachinePoolDeleteParamsWithContext(h.Ctx).WithConfigUID(cloudConfigId).WithMachinePoolName(machinePoolName)
	_, err = client.V1CloudConfigsVsphereMachinePoolDelete(params)
	return err
}

// Cloud Config

func (h *V1Client) GetCloudConfigVsphere(configUID string) (*models.V1VsphereCloudConfig, error) {
	if h.GetCloudConfigVsphereFn != nil {
		return h.GetCloudConfigVsphereFn(configUID)
	}
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}

	params := clusterC.NewV1CloudConfigsVsphereGetParamsWithContext(h.Ctx).WithConfigUID(configUID)
	success, err := client.V1CloudConfigsVsphereGet(params)
	if e, ok := err.(*hapitransport.TransportError); ok && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return success.Payload, nil
}

func (h *V1Client) GetCloudConfigVsphereValues(uid string) (*models.V1VsphereCloudConfig, error) {
	if h.GetCloudConfigVsphereValuesFn != nil {
		return h.GetCloudConfigVsphereValuesFn(uid)
	}
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}

	params := clusterC.NewV1CloudConfigsVsphereGetParamsWithContext(h.Ctx).WithConfigUID(uid)
	cloudConfig, err := client.V1CloudConfigsVsphereGet(params)
	if err != nil {
		return nil, err
	}

	return cloudConfig.Payload, nil
}

func (h *V1Client) UpdateCloudConfigVsphereValues(uid string, clusterConfig *models.V1VsphereCloudClusterConfigEntity) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return err
	}

	params := clusterC.NewV1CloudConfigsVsphereUIDClusterConfigParamsWithContext(h.Ctx).WithConfigUID(uid).WithBody(clusterConfig)
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
