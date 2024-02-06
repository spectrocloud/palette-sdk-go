package client

import (
	"errors"

	"github.com/spectrocloud/palette-api-go/apiutil/transport"
	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
)

// Cluster

func (h *V1Client) CreateClusterMaas(cluster *models.V1SpectroMaasClusterEntity, scope string) (string, error) {
	var params *clientV1.V1SpectroClustersMaasCreateParams
	switch scope {
	case "project":
		params = clientV1.NewV1SpectroClustersMaasCreateParamsWithContext(h.Ctx).WithBody(cluster)
	case "tenant":
		params = clientV1.NewV1SpectroClustersMaasCreateParams().WithBody(cluster)
	}

	success, err := h.GetClient().V1SpectroClustersMaasCreate(params)
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

// Machine Pool

func (h *V1Client) CreateMachinePoolMaas(CloudConfigId, scope string, machinePool *models.V1MaasMachinePoolConfigEntity) error {
	var params *clientV1.V1CloudConfigsMaasMachinePoolCreateParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsMaasMachinePoolCreateParamsWithContext(h.Ctx).WithConfigUID(CloudConfigId).WithBody(machinePool)
	case "tenant":
		params = clientV1.NewV1CloudConfigsMaasMachinePoolCreateParams().WithConfigUID(CloudConfigId).WithBody(machinePool)
	}

	_, err := h.GetClient().V1CloudConfigsMaasMachinePoolCreate(params)
	return err
}

func (h *V1Client) UpdateMachinePoolMaas(CloudConfigId, scope string, machinePool *models.V1MaasMachinePoolConfigEntity) error {
	var params *clientV1.V1CloudConfigsMaasMachinePoolUpdateParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsMaasMachinePoolUpdateParamsWithContext(h.Ctx).
			WithConfigUID(CloudConfigId).
			WithMachinePoolName(*machinePool.PoolConfig.Name).
			WithBody(machinePool)
	case "tenant":
		params = clientV1.NewV1CloudConfigsMaasMachinePoolUpdateParams().
			WithConfigUID(CloudConfigId).
			WithMachinePoolName(*machinePool.PoolConfig.Name).
			WithBody(machinePool)
	}

	_, err := h.GetClient().V1CloudConfigsMaasMachinePoolUpdate(params)
	return err
}

func (h *V1Client) DeleteMachinePoolMaas(CloudConfigId, machinePoolName, scope string) error {
	var params *clientV1.V1CloudConfigsMaasMachinePoolDeleteParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsMaasMachinePoolDeleteParamsWithContext(h.Ctx).WithConfigUID(CloudConfigId).WithMachinePoolName(machinePoolName)
	case "tenant":
		params = clientV1.NewV1CloudConfigsMaasMachinePoolDeleteParams().WithConfigUID(CloudConfigId).WithMachinePoolName(machinePoolName)
	}

	_, err := h.GetClient().V1CloudConfigsMaasMachinePoolDelete(params)
	return err
}

// Cloud Config

func (h *V1Client) GetCloudConfigMaas(configUID, scope string) (*models.V1MaasCloudConfig, error) {
	var params *clientV1.V1CloudConfigsMaasGetParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsMaasGetParamsWithContext(h.Ctx).WithConfigUID(configUID)
	case "tenant":
		params = clientV1.NewV1CloudConfigsMaasGetParams().WithConfigUID(configUID)
	}

	success, err := h.GetClient().V1CloudConfigsMaasGet(params)
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
	params := clientV1.NewV1SpectroClustersMaasImportParamsWithContext(h.Ctx).WithBody(
		&models.V1SpectroMaasClusterImportEntity{
			Metadata: meta,
		},
	)
	success, err := h.GetClient().V1SpectroClustersMaasImport(params)
	if err != nil {
		return "", err
	}
	return *success.Payload.UID, nil
}

func (h *V1Client) GetNodeStatusMapMaas(configUID, machinePoolName, scope string) (map[string]models.V1CloudMachineStatus, error) {
	var params *clientV1.V1CloudConfigsMaasPoolMachinesListParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsMaasPoolMachinesListParamsWithContext(h.Ctx).WithConfigUID(configUID).WithMachinePoolName(machinePoolName)
	case "tenant":
		params = clientV1.NewV1CloudConfigsMaasPoolMachinesListParams().WithConfigUID(configUID).WithMachinePoolName(machinePoolName)
	}

	mpList, err := h.GetClient().V1CloudConfigsMaasPoolMachinesList(params)
	nMap := map[string]models.V1CloudMachineStatus{}
	if len(mpList.Payload.Items) > 0 {
		for _, node := range mpList.Payload.Items {
			nMap[node.Metadata.UID] = *node.Status
		}
	}
	return nMap, err
}
