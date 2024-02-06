package client

import (
	"errors"

	"github.com/spectrocloud/palette-api-go/apiutil/transport"
	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
)

func (h *V1Client) CreateClusterAzure(cluster *models.V1SpectroAzureClusterEntity, scope string) (string, error) {
	var params *clientV1.V1SpectroClustersAzureCreateParams
	switch scope {
	case "project":
		params = clientV1.NewV1SpectroClustersAzureCreateParamsWithContext(h.Ctx).WithBody(cluster)
	case "tenant":
		params = clientV1.NewV1SpectroClustersAzureCreateParams().WithBody(cluster)
	}

	success, err := h.GetClient().V1SpectroClustersAzureCreate(params)
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

func (h *V1Client) CreateMachinePoolAzure(CloudConfigId, scope string, machinePool *models.V1AzureMachinePoolConfigEntity) error {
	var params *clientV1.V1CloudConfigsAzureMachinePoolCreateParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsAzureMachinePoolCreateParamsWithContext(h.Ctx).WithConfigUID(CloudConfigId).WithBody(machinePool)
	case "tenant":
		params = clientV1.NewV1CloudConfigsAzureMachinePoolCreateParams().WithConfigUID(CloudConfigId).WithBody(machinePool)
	}

	_, err := h.GetClient().V1CloudConfigsAzureMachinePoolCreate(params)
	return err
}

func (h *V1Client) UpdateMachinePoolAzure(CloudConfigId, scope string, machinePool *models.V1AzureMachinePoolConfigEntity) error {
	var params *clientV1.V1CloudConfigsAzureMachinePoolUpdateParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsAzureMachinePoolUpdateParamsWithContext(h.Ctx).
			WithConfigUID(CloudConfigId).
			WithMachinePoolName(*machinePool.PoolConfig.Name).
			WithBody(machinePool)
	case "tenant":
		params = clientV1.NewV1CloudConfigsAzureMachinePoolUpdateParams().
			WithConfigUID(CloudConfigId).
			WithMachinePoolName(*machinePool.PoolConfig.Name).
			WithBody(machinePool)
	}

	_, err := h.GetClient().V1CloudConfigsAzureMachinePoolUpdate(params)
	return err
}

func (h *V1Client) DeleteMachinePoolAzure(CloudConfigId, machinePoolName, scope string) error {
	var params *clientV1.V1CloudConfigsAzureMachinePoolDeleteParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsAzureMachinePoolDeleteParamsWithContext(h.Ctx).WithConfigUID(CloudConfigId).WithMachinePoolName(machinePoolName)
	case "tenant":
		params = clientV1.NewV1CloudConfigsAzureMachinePoolDeleteParams().WithConfigUID(CloudConfigId).WithMachinePoolName(machinePoolName)
	}

	_, err := h.GetClient().V1CloudConfigsAzureMachinePoolDelete(params)
	return err
}

func (h *V1Client) GetCloudConfigAzure(configUID, scope string) (*models.V1AzureCloudConfig, error) {
	var params *clientV1.V1CloudConfigsAzureGetParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsAzureGetParamsWithContext(h.Ctx).WithConfigUID(configUID)
	case "tenant":
		params = clientV1.NewV1CloudConfigsAzureGetParams().WithConfigUID(configUID)
	}

	success, err := h.GetClient().V1CloudConfigsAzureGet(params)

	var e *transport.TransportError
	if errors.As(err, &e) && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return success.Payload, nil
}

func (h *V1Client) ImportClusterAzure(meta *models.V1ObjectMetaInputEntity) (string, error) {
	params := clientV1.NewV1SpectroClustersAzureImportParamsWithContext(h.Ctx).WithBody(
		&models.V1SpectroAzureClusterImportEntity{
			Metadata: meta,
		},
	)
	success, err := h.GetClient().V1SpectroClustersAzureImport(params)
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

func (h *V1Client) GetNodeStatusMapAzure(configUID, machinePoolName, scope string) (map[string]models.V1CloudMachineStatus, error) {
	var params *clientV1.V1CloudConfigsAzurePoolMachinesListParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsAzurePoolMachinesListParamsWithContext(h.Ctx).WithConfigUID(configUID).WithMachinePoolName(machinePoolName)
	case "tenant":
		params = clientV1.NewV1CloudConfigsAzurePoolMachinesListParams().WithConfigUID(configUID).WithMachinePoolName(machinePoolName)
	}

	mpList, err := h.GetClient().V1CloudConfigsAzurePoolMachinesList(params)
	nMap := map[string]models.V1CloudMachineStatus{}
	if len(mpList.Payload.Items) > 0 {
		for _, node := range mpList.Payload.Items {
			nMap[node.Metadata.UID] = *node.Status
		}
	}
	return nMap, err
}
