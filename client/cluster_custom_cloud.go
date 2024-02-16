package client

import (
	"fmt"
	"github.com/spectrocloud/hapi/models"
	clusterC "github.com/spectrocloud/hapi/spectrocluster/client/v1"
)

func (h *V1Client) CreateClusterCustomCloud(cluster *models.V1SpectroCustomClusterEntity, cloudType string, clusterContext string) (string, error) {
	if h.CreateClusterCustomCloudFn != nil {
		return h.CreateClusterCustomCloudFn(cluster, cloudType, clusterContext)
	}

	var params *clusterC.V1SpectroClustersCustomCreateParams

	switch clusterContext {
	case "project":
		params = clusterC.NewV1SpectroClustersCustomCreateParamsWithContext(h.Ctx)
	case "tenant":
		params = clusterC.NewV1SpectroClustersCustomCreateParams()
	}
	params = params.WithCloudType(cloudType).WithBody(cluster)
	success, err := h.GetClusterClient().V1SpectroClustersCustomCreate(params)
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

func (h *V1Client) GetCloudConfigCustomCloud(configUID string, cloudType string, clusterContext string) (*models.V1CustomCloudConfig, error) {
	if h.GetCloudConfigCustomCloudFn != nil {
		return h.GetCloudConfigCustomCloudFn(configUID, cloudType, clusterContext)
	}
	var params *clusterC.V1CloudConfigsCustomGetParams
	switch clusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsCustomGetParamsWithContext(h.Ctx)
	case "tenant":
		params = clusterC.NewV1CloudConfigsCustomGetParams()
	default:
		return nil, fmt.Errorf("invalid scope %s", clusterContext)
	}
	params = params.WithCloudType(cloudType).WithConfigUID(configUID)
	success, err := h.GetClusterClient().V1CloudConfigsCustomGet(params)
	if err != nil {
		return nil, err
	}

	// special check if the cluster is marked deleted
	cluster := success.Payload
	return cluster, nil

}

func (h *V1Client) UpdateCloudConfigCustomCloud(updatedConfig *models.V1CustomCloudClusterConfigEntity, configUID string, cloudType string, clusterContext string) error {
	if h.UpdateCloudConfigCustomCloudFn != nil {
		return h.UpdateCloudConfigCustomCloudFn(updatedConfig, configUID, cloudType, clusterContext)
	}

	var params *clusterC.V1CloudConfigsCustomUIDClusterConfigParams
	switch clusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsCustomUIDClusterConfigParamsWithContext(h.Ctx)
	case "tenant":
		params = clusterC.NewV1CloudConfigsCustomUIDClusterConfigParams()
	}
	params = params.WithCloudType(cloudType).WithBody(updatedConfig).WithConfigUID(configUID)
	_, err := h.GetClusterClient().V1CloudConfigsCustomUIDClusterConfig(params)
	if err != nil {
		return err
	}
	return nil
}

func (h *V1Client) CreateMachinePoolCustomCloud(mpEntity *models.V1CustomMachinePoolConfigEntity, configUID string, cloudType string, clusterContext string) error {
	if h.CreateMachinePoolCustomCloudFn != nil {
		return h.CreateMachinePoolCustomCloudFn(mpEntity, configUID, cloudType, clusterContext)
	}

	var params *clusterC.V1CloudConfigsCustomMachinePoolCreateParams
	switch clusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsCustomMachinePoolCreateParamsWithContext(h.Ctx)
	case "tenant":
		params = clusterC.NewV1CloudConfigsCustomMachinePoolCreateParams()
	}
	params = params.WithCloudType(cloudType).WithBody(mpEntity).WithConfigUID(configUID)
	_, err := h.GetClusterClient().V1CloudConfigsCustomMachinePoolCreate(params)
	if err != nil {
		return err
	}
	return nil
}

func (h *V1Client) UpdateMachinePoolCustomCloud(mpEntity *models.V1CustomMachinePoolConfigEntity, machinePoolName, configUID string, cloudType string, clusterContext string) error {
	if h.UpdateMachinePoolCustomCloudFn != nil {
		return h.UpdateMachinePoolCustomCloudFn(mpEntity, machinePoolName, configUID, cloudType, clusterContext)
	}

	var params *clusterC.V1CloudConfigsCustomMachinePoolUpdateParams
	switch clusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsCustomMachinePoolUpdateParamsWithContext(h.Ctx)
	case "tenant":
		params = clusterC.NewV1CloudConfigsCustomMachinePoolUpdateParams()
	}
	params = params.WithCloudType(cloudType).WithBody(mpEntity).WithConfigUID(configUID).WithMachinePoolName(machinePoolName)
	_, err := h.GetClusterClient().V1CloudConfigsCustomMachinePoolUpdate(params)
	if err != nil {
		return err
	}
	return nil
}

func (h *V1Client) DeleteMachinePoolCustomCloud(mpName string, configUID string, cloudType string, clusterContext string) error {
	if h.DeleteMachinePoolCustomCloudFn != nil {
		return h.DeleteMachinePoolCustomCloudFn(mpName, configUID, cloudType, clusterContext)
	}

	var params *clusterC.V1CloudConfigsCustomMachinePoolDeleteParams
	switch clusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsCustomMachinePoolDeleteParamsWithContext(h.Ctx)
	case "tenant":
		params = clusterC.NewV1CloudConfigsCustomMachinePoolDeleteParams()
	}
	params = params.WithCloudType(cloudType).WithConfigUID(configUID).WithMachinePoolName(mpName)
	_, err := h.GetClusterClient().V1CloudConfigsCustomMachinePoolDelete(params)
	if err != nil {
		return err
	}
	return nil
}
