package client

import (
	"errors"

	"github.com/spectrocloud/palette-api-go/apiutil/transport"
	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
)

// Cluster

func (h *V1Client) CreateClusterVsphere(cluster *models.V1SpectroVsphereClusterEntity, scope string) (string, error) {
	var params *clientV1.V1SpectroClustersVsphereCreateParams
	switch scope {
	case "project":
		params = clientV1.NewV1SpectroClustersVsphereCreateParamsWithContext(h.Ctx).WithBody(cluster)
	case "tenant":
		params = clientV1.NewV1SpectroClustersVsphereCreateParams().WithBody(cluster)
	}

	success, err := h.GetClient().V1SpectroClustersVsphereCreate(params)
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

// Machine Pool

func (h *V1Client) CreateMachinePoolVsphere(CloudConfigId, scope string, machinePool *models.V1VsphereMachinePoolConfigEntity) error {
	var params *clientV1.V1CloudConfigsVsphereMachinePoolCreateParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsVsphereMachinePoolCreateParamsWithContext(h.Ctx).WithConfigUID(CloudConfigId).WithBody(machinePool)
	case "tenant":
		params = clientV1.NewV1CloudConfigsVsphereMachinePoolCreateParams().WithConfigUID(CloudConfigId).WithBody(machinePool)
	}

	_, err := h.GetClient().V1CloudConfigsVsphereMachinePoolCreate(params)
	return err
}

func (h *V1Client) UpdateMachinePoolVsphere(CloudConfigId, scope string, machinePool *models.V1VsphereMachinePoolConfigEntity) error {
	var params *clientV1.V1CloudConfigsVsphereMachinePoolUpdateParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsVsphereMachinePoolUpdateParamsWithContext(h.Ctx).
			WithConfigUID(CloudConfigId).
			WithMachinePoolName(*machinePool.PoolConfig.Name).
			WithBody(machinePool)
	case "tenant":
		params = clientV1.NewV1CloudConfigsVsphereMachinePoolUpdateParams().
			WithConfigUID(CloudConfigId).
			WithMachinePoolName(*machinePool.PoolConfig.Name).
			WithBody(machinePool)
	}

	_, err := h.GetClient().V1CloudConfigsVsphereMachinePoolUpdate(params)
	return err
}

func (h *V1Client) DeleteMachinePoolVsphere(CloudConfigId, machinePoolName, scope string) error {
	var params *clientV1.V1CloudConfigsVsphereMachinePoolDeleteParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsVsphereMachinePoolDeleteParamsWithContext(h.Ctx).WithConfigUID(CloudConfigId).WithMachinePoolName(machinePoolName)
	case "tenant":
		params = clientV1.NewV1CloudConfigsVsphereMachinePoolDeleteParams().WithConfigUID(CloudConfigId).WithMachinePoolName(machinePoolName)
	}

	_, err := h.GetClient().V1CloudConfigsVsphereMachinePoolDelete(params)
	return err
}

// Cloud Config

func (h *V1Client) GetCloudConfigVsphere(configUID, scope string) (*models.V1VsphereCloudConfig, error) {
	var params *clientV1.V1CloudConfigsVsphereGetParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsVsphereGetParamsWithContext(h.Ctx).WithConfigUID(configUID)
	case "tenant":
		params = clientV1.NewV1CloudConfigsVsphereGetParams().WithConfigUID(configUID)
	}

	success, err := h.GetClient().V1CloudConfigsVsphereGet(params)
	var e *transport.TransportError
	if errors.As(err, &e) && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return success.Payload, nil
}

func (h *V1Client) GetCloudConfigVsphereValues(uid, scope string) (*models.V1VsphereCloudConfig, error) {
	var params *clientV1.V1CloudConfigsVsphereGetParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsVsphereGetParamsWithContext(h.Ctx).WithConfigUID(uid)
	case "tenant":
		params = clientV1.NewV1CloudConfigsVsphereGetParams().WithConfigUID(uid)
	}

	CloudConfig, err := h.GetClient().V1CloudConfigsVsphereGet(params)
	if err != nil {
		return nil, err
	}

	return CloudConfig.Payload, nil
}

func (h *V1Client) UpdateCloudConfigVsphereValues(uid, scope string, cloudConfig *models.V1VsphereCloudClusterConfigEntity) error {
	var params *clientV1.V1CloudConfigsVsphereUIDClusterConfigParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsVsphereUIDClusterConfigParamsWithContext(h.Ctx).WithConfigUID(uid).WithBody(cloudConfig)
	case "tenant":
		params = clientV1.NewV1CloudConfigsVsphereUIDClusterConfigParams().WithConfigUID(uid).WithBody(cloudConfig)
	}

	_, err := h.GetClient().V1CloudConfigsVsphereUIDClusterConfig(params)
	return err
}

// Import

func (h *V1Client) ImportClusterVsphere(meta *models.V1ObjectMetaInputEntity) (string, error) {
	params := clientV1.NewV1SpectroClustersVsphereImportParamsWithContext(h.Ctx).WithBody(
		&models.V1SpectroVsphereClusterImportEntity{
			Metadata: meta,
		},
	)

	success, err := h.GetClient().V1SpectroClustersVsphereImport(params)
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

func (h *V1Client) GetNodeStatusMapVsphere(configUID, machinePoolName, scope string) (map[string]models.V1CloudMachineStatus, error) {
	var params *clientV1.V1CloudConfigsVspherePoolMachinesListParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsVspherePoolMachinesListParamsWithContext(h.Ctx).WithConfigUID(configUID).WithMachinePoolName(machinePoolName)
	case "tenant":
		params = clientV1.NewV1CloudConfigsVspherePoolMachinesListParams().WithConfigUID(configUID).WithMachinePoolName(machinePoolName)
	}

	mpList, err := h.GetClient().V1CloudConfigsVspherePoolMachinesList(params)
	nMap := map[string]models.V1CloudMachineStatus{}
	if len(mpList.Payload.Items) > 0 {
		for _, node := range mpList.Payload.Items {
			nMap[node.Metadata.UID] = *node.Status
		}
	}
	return nMap, err
}
