package client

import (
	"errors"
	"time"

	"github.com/spectrocloud/palette-api-go/apiutil/transport"
	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
)

func (h *V1Client) CreateClusterAks(cluster *models.V1SpectroAzureClusterEntity, scope string) (string, error) {
	var params *clientV1.V1SpectroClustersAksCreateParams
	switch scope {
	case "project":
		params = clientV1.NewV1SpectroClustersAksCreateParamsWithContext(h.Ctx).WithBody(cluster)
	case "tenant":
		params = clientV1.NewV1SpectroClustersAksCreateParams().WithBody(cluster)
	}

	success, err := h.Client.V1SpectroClustersAksCreate(params.WithTimeout(90 * time.Second))
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

func (h *V1Client) CreateMachinePoolAks(CloudConfigId string, machinePool *models.V1AzureMachinePoolConfigEntity, scope string) error {
	var params *clientV1.V1CloudConfigsAksMachinePoolCreateParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsAksMachinePoolCreateParamsWithContext(h.Ctx).WithConfigUID(CloudConfigId).WithBody(machinePool)
	case "tenant":
		params = clientV1.NewV1CloudConfigsAksMachinePoolCreateParams().WithConfigUID(CloudConfigId).WithBody(machinePool)
	}

	_, err := h.Client.V1CloudConfigsAksMachinePoolCreate(params)
	return err
}

func (h *V1Client) UpdateMachinePoolAks(CloudConfigId string, machinePool *models.V1AzureMachinePoolConfigEntity, scope string) error {
	var params *clientV1.V1CloudConfigsAksMachinePoolUpdateParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsAksMachinePoolUpdateParamsWithContext(h.Ctx).
			WithConfigUID(CloudConfigId).
			WithMachinePoolName(*machinePool.PoolConfig.Name).
			WithBody(machinePool)
	case "tenant":
		params = clientV1.NewV1CloudConfigsAksMachinePoolUpdateParams().
			WithConfigUID(CloudConfigId).
			WithMachinePoolName(*machinePool.PoolConfig.Name).
			WithBody(machinePool)
	}

	_, err := h.Client.V1CloudConfigsAksMachinePoolUpdate(params)
	return err
}

func (h *V1Client) DeleteMachinePoolAks(CloudConfigId, machinePoolName, scope string) error {
	var params *clientV1.V1CloudConfigsAksMachinePoolDeleteParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsAksMachinePoolDeleteParamsWithContext(h.Ctx).WithConfigUID(CloudConfigId).WithMachinePoolName(machinePoolName)
	case "tenant":
		params = clientV1.NewV1CloudConfigsAksMachinePoolDeleteParams().WithConfigUID(CloudConfigId).WithMachinePoolName(machinePoolName)
	}

	_, err := h.Client.V1CloudConfigsAksMachinePoolDelete(params)
	return err
}

func (h *V1Client) GetCloudConfigAks(configUID, scope string) (*models.V1AzureCloudConfig, error) {
	var params *clientV1.V1CloudConfigsAksGetParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsAksGetParamsWithContext(h.Ctx).WithConfigUID(configUID)
	case "tenant":
		params = clientV1.NewV1CloudConfigsAksGetParams().WithConfigUID(configUID)
	}

	success, err := h.Client.V1CloudConfigsAksGet(params)
	var e *transport.TransportError
	if errors.As(err, &e) && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return success.Payload, nil
}

func (h *V1Client) GetNodeStatusMapAks(configUID, machinePoolName, scope string) (map[string]models.V1CloudMachineStatus, error) {
	var params *clientV1.V1CloudConfigsAksPoolMachinesListParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsAksPoolMachinesListParamsWithContext(h.Ctx).WithConfigUID(configUID).WithMachinePoolName(machinePoolName)
	case "tenant":
		params = clientV1.NewV1CloudConfigsAksPoolMachinesListParams().WithConfigUID(configUID).WithMachinePoolName(machinePoolName)
	}

	mpList, err := h.Client.V1CloudConfigsAksPoolMachinesList(params)
	nMap := map[string]models.V1CloudMachineStatus{}
	if len(mpList.Payload.Items) > 0 {
		for _, node := range mpList.Payload.Items {
			nMap[node.Metadata.UID] = *node.Status
		}
	}
	return nMap, err
}
