package client

import (
	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
	"github.com/spectrocloud/palette-sdk-go/client/herr"
)

// Cluster

func (h *V1Client) CreateClusterOpenStack(cluster *models.V1SpectroOpenStackClusterEntity, scope string) (string, error) {
	var params *clientV1.V1SpectroClustersOpenStackCreateParams
	switch scope {
	case "project":
		params = clientV1.NewV1SpectroClustersOpenStackCreateParamsWithContext(h.Ctx).WithBody(cluster)
	case "tenant":
		params = clientV1.NewV1SpectroClustersOpenStackCreateParams().WithBody(cluster)
	}

	success, err := h.Client.V1SpectroClustersOpenStackCreate(params)
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

// Machine Pool

func (h *V1Client) CreateMachinePoolOpenStack(CloudConfigId, scope string, machinePool *models.V1OpenStackMachinePoolConfigEntity) error {
	var params *clientV1.V1CloudConfigsOpenStackMachinePoolCreateParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsOpenStackMachinePoolCreateParamsWithContext(h.Ctx).WithConfigUID(CloudConfigId).WithBody(machinePool)
	case "tenant":
		params = clientV1.NewV1CloudConfigsOpenStackMachinePoolCreateParams().WithConfigUID(CloudConfigId).WithBody(machinePool)
	}

	_, err := h.Client.V1CloudConfigsOpenStackMachinePoolCreate(params)
	return err
}

func (h *V1Client) UpdateMachinePoolOpenStack(CloudConfigId, scope string, machinePool *models.V1OpenStackMachinePoolConfigEntity) error {
	var params *clientV1.V1CloudConfigsOpenStackMachinePoolUpdateParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsOpenStackMachinePoolUpdateParamsWithContext(h.Ctx).
			WithConfigUID(CloudConfigId).
			WithMachinePoolName(*machinePool.PoolConfig.Name).
			WithBody(machinePool)
	case "tenant":
		params = clientV1.NewV1CloudConfigsOpenStackMachinePoolUpdateParams().
			WithConfigUID(CloudConfigId).
			WithMachinePoolName(*machinePool.PoolConfig.Name).
			WithBody(machinePool)
	}

	_, err := h.Client.V1CloudConfigsOpenStackMachinePoolUpdate(params)
	return err
}

func (h *V1Client) DeleteMachinePoolOpenStack(CloudConfigId, machinePoolName, scope string) error {
	var params *clientV1.V1CloudConfigsOpenStackMachinePoolDeleteParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsOpenStackMachinePoolDeleteParamsWithContext(h.Ctx).WithConfigUID(CloudConfigId).WithMachinePoolName(machinePoolName)
	case "tenant":
		params = clientV1.NewV1CloudConfigsOpenStackMachinePoolDeleteParams().WithConfigUID(CloudConfigId).WithMachinePoolName(machinePoolName)
	}

	_, err := h.Client.V1CloudConfigsOpenStackMachinePoolDelete(params)
	return err
}

// Cloud Config

func (h *V1Client) GetCloudConfigOpenStack(configUID, scope string) (*models.V1OpenStackCloudConfig, error) {
	var params *clientV1.V1CloudConfigsOpenStackGetParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsOpenStackGetParamsWithContext(h.Ctx).WithConfigUID(configUID)
	case "tenant":
		params = clientV1.NewV1CloudConfigsOpenStackGetParams().WithConfigUID(configUID)
	}

	success, err := h.Client.V1CloudConfigsOpenStackGet(params)

	if herr.IsNotFound(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return success.Payload, nil
}

// Import

func (h *V1Client) ImportClusterOpenStack(meta *models.V1ObjectMetaInputEntity) (string, error) {
	params := clientV1.NewV1SpectroClustersOpenStackImportParamsWithContext(h.Ctx).WithBody(
		&models.V1SpectroOpenStackClusterImportEntity{
			Metadata: meta,
		},
	)
	success, err := h.Client.V1SpectroClustersOpenStackImport(params)
	if err != nil {
		return "", err
	}
	return *success.Payload.UID, nil
}

func (h *V1Client) GetNodeStatusMapOpenStack(configUID, machinePoolName, scope string) (map[string]models.V1CloudMachineStatus, error) {
	var params *clientV1.V1CloudConfigsOpenStackPoolMachinesListParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsOpenStackPoolMachinesListParamsWithContext(h.Ctx).WithConfigUID(configUID).WithMachinePoolName(machinePoolName)
	case "tenant":
		params = clientV1.NewV1CloudConfigsOpenStackPoolMachinesListParams().WithConfigUID(configUID).WithMachinePoolName(machinePoolName)
	}

	mpList, err := h.Client.V1CloudConfigsOpenStackPoolMachinesList(params)
	nMap := map[string]models.V1CloudMachineStatus{}
	if len(mpList.Payload.Items) > 0 {
		for _, node := range mpList.Payload.Items {
			nMap[node.Metadata.UID] = *node.Status
		}
	}
	return nMap, err
}
