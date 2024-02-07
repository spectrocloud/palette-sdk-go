package client

import (
	"errors"

	"github.com/spectrocloud/palette-api-go/apiutil/transport"
	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
)

func (h *V1Client) CreateClusterGcp(cluster *models.V1SpectroGcpClusterEntity, scope string) (string, error) {
	var params *clientV1.V1SpectroClustersGcpCreateParams
	switch scope {
	case "project":
		params = clientV1.NewV1SpectroClustersGcpCreateParamsWithContext(h.Ctx).WithBody(cluster)
	case "tenant":
		params = clientV1.NewV1SpectroClustersGcpCreateParams().WithBody(cluster)
	}

	success, err := h.Client.V1SpectroClustersGcpCreate(params)
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

func (h *V1Client) CreateMachinePoolGcp(CloudConfigId, scope string, machinePool *models.V1GcpMachinePoolConfigEntity) error {
	var params *clientV1.V1CloudConfigsGcpMachinePoolCreateParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsGcpMachinePoolCreateParamsWithContext(h.Ctx).WithConfigUID(CloudConfigId).WithBody(machinePool)
	case "tenant":
		params = clientV1.NewV1CloudConfigsGcpMachinePoolCreateParams().WithConfigUID(CloudConfigId).WithBody(machinePool)
	}

	_, err := h.Client.V1CloudConfigsGcpMachinePoolCreate(params)
	return err
}

func (h *V1Client) UpdateMachinePoolGcp(CloudConfigId, scope string, machinePool *models.V1GcpMachinePoolConfigEntity) error {
	var params *clientV1.V1CloudConfigsGcpMachinePoolUpdateParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsGcpMachinePoolUpdateParamsWithContext(h.Ctx).
			WithConfigUID(CloudConfigId).
			WithMachinePoolName(*machinePool.PoolConfig.Name).
			WithBody(machinePool)
	case "tenant":
		params = clientV1.NewV1CloudConfigsGcpMachinePoolUpdateParams().
			WithConfigUID(CloudConfigId).
			WithMachinePoolName(*machinePool.PoolConfig.Name).
			WithBody(machinePool)
	}

	_, err := h.Client.V1CloudConfigsGcpMachinePoolUpdate(params)
	return err
}

func (h *V1Client) DeleteMachinePoolGcp(CloudConfigId, machinePoolName, scope string) error {
	var params *clientV1.V1CloudConfigsGcpMachinePoolDeleteParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsGcpMachinePoolDeleteParamsWithContext(h.Ctx).WithConfigUID(CloudConfigId).WithMachinePoolName(machinePoolName)
	case "tenant":
		params = clientV1.NewV1CloudConfigsGcpMachinePoolDeleteParams().WithConfigUID(CloudConfigId).WithMachinePoolName(machinePoolName)
	}

	_, err := h.Client.V1CloudConfigsGcpMachinePoolDelete(params)
	return err
}

func (h *V1Client) GetCloudConfigGcp(configUID, scope string) (*models.V1GcpCloudConfig, error) {
	var params *clientV1.V1CloudConfigsGcpGetParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsGcpGetParamsWithContext(h.Ctx).WithConfigUID(configUID)
	case "tenant":
		params = clientV1.NewV1CloudConfigsGcpGetParams().WithConfigUID(configUID)
	}

	success, err := h.Client.V1CloudConfigsGcpGet(params)
	var e *transport.TransportError
	if errors.As(err, &e) && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return success.Payload, nil
}

func (h *V1Client) ImportClusterGcp(meta *models.V1ObjectMetaInputEntity) (string, error) {
	params := clientV1.NewV1SpectroClustersGcpImportParamsWithContext(h.Ctx).WithBody(
		&models.V1SpectroGcpClusterImportEntity{
			Metadata: meta,
		},
	)
	success, err := h.Client.V1SpectroClustersGcpImport(params)
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

func (h *V1Client) GetNodeStatusMapGcp(configUID, machinePoolName, scope string) (map[string]models.V1CloudMachineStatus, error) {
	var params *clientV1.V1CloudConfigsGcpPoolMachinesListParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsGcpPoolMachinesListParamsWithContext(h.Ctx).WithConfigUID(configUID).WithMachinePoolName(machinePoolName)
	case "tenant":
		params = clientV1.NewV1CloudConfigsGcpPoolMachinesListParams().WithConfigUID(configUID).WithMachinePoolName(machinePoolName)
	}

	mpList, err := h.Client.V1CloudConfigsGcpPoolMachinesList(params)
	nMap := map[string]models.V1CloudMachineStatus{}
	if len(mpList.Payload.Items) > 0 {
		for _, node := range mpList.Payload.Items {
			nMap[node.Metadata.UID] = *node.Status
		}
	}
	return nMap, err
}
