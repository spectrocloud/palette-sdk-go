package client

import (
	"errors"

	"github.com/spectrocloud/palette-api-go/apiutil/transport"
	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
)

func (h *V1Client) CreateClusterTke(cluster *models.V1SpectroTencentClusterEntity, scope string) (string, error) {
	var params *clientV1.V1SpectroClustersTkeCreateParams
	switch scope {
	case "project":
		params = clientV1.NewV1SpectroClustersTkeCreateParamsWithContext(h.Ctx).WithBody(cluster)
	case "tenant":
		params = clientV1.NewV1SpectroClustersTkeCreateParams().WithBody(cluster)
	}

	success, err := h.GetClient().V1SpectroClustersTkeCreate(params)
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

func (h *V1Client) CreateMachinePoolTke(CloudConfigId, scope string, machinePool *models.V1TencentMachinePoolConfigEntity) error {
	var params *clientV1.V1CloudConfigsTkeMachinePoolCreateParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsTkeMachinePoolCreateParamsWithContext(h.Ctx).WithConfigUID(CloudConfigId).WithBody(machinePool)
	case "tenant":
		params = clientV1.NewV1CloudConfigsTkeMachinePoolCreateParams().WithConfigUID(CloudConfigId).WithBody(machinePool)
	}

	_, err := h.GetClient().V1CloudConfigsTkeMachinePoolCreate(params)
	return err
}

func (h *V1Client) UpdateMachinePoolTke(CloudConfigId, scope string, machinePool *models.V1TencentMachinePoolConfigEntity) error {
	var params *clientV1.V1CloudConfigsTkeMachinePoolUpdateParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsTkeMachinePoolUpdateParamsWithContext(h.Ctx).
			WithConfigUID(CloudConfigId).
			WithMachinePoolName(*machinePool.PoolConfig.Name).
			WithBody(machinePool)
	case "tenant":
		params = clientV1.NewV1CloudConfigsTkeMachinePoolUpdateParams().
			WithConfigUID(CloudConfigId).
			WithMachinePoolName(*machinePool.PoolConfig.Name).
			WithBody(machinePool)
	}

	_, err := h.GetClient().V1CloudConfigsTkeMachinePoolUpdate(params)
	return err
}

func (h *V1Client) DeleteMachinePoolTke(CloudConfigId, machinePoolName, scope string) error {
	var params *clientV1.V1CloudConfigsTkeMachinePoolDeleteParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsTkeMachinePoolDeleteParamsWithContext(h.Ctx).WithConfigUID(CloudConfigId).WithMachinePoolName(machinePoolName)
	case "tenant":
		params = clientV1.NewV1CloudConfigsTkeMachinePoolDeleteParams().WithConfigUID(CloudConfigId).WithMachinePoolName(machinePoolName)
	}

	_, err := h.GetClient().V1CloudConfigsTkeMachinePoolDelete(params)
	return err
}

func (h *V1Client) GetCloudConfigTke(configUID, scope string) (*models.V1TencentCloudConfig, error) {
	var params *clientV1.V1CloudConfigsTkeGetParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsTkeGetParamsWithContext(h.Ctx).WithConfigUID(configUID)
	case "tenant":
		params = clientV1.NewV1CloudConfigsTkeGetParams().WithConfigUID(configUID)
	}

	success, err := h.GetClient().V1CloudConfigsTkeGet(params)
	var e *transport.TransportError
	if errors.As(err, &e) && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return success.Payload, nil
}

func (h *V1Client) GetNodeStatusMapTke(configUID, machinePoolName, scope string) (map[string]models.V1CloudMachineStatus, error) {
	var params *clientV1.V1CloudConfigsTkePoolMachinesListParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsTkePoolMachinesListParamsWithContext(h.Ctx).WithConfigUID(configUID).WithMachinePoolName(machinePoolName)
	case "tenant":
		params = clientV1.NewV1CloudConfigsTkePoolMachinesListParams().WithConfigUID(configUID).WithMachinePoolName(machinePoolName)
	}

	mpList, err := h.GetClient().V1CloudConfigsTkePoolMachinesList(params)
	nMap := map[string]models.V1CloudMachineStatus{}
	if len(mpList.Payload.Items) > 0 {
		for _, node := range mpList.Payload.Items {
			nMap[node.Metadata.UID] = *node.Status
		}
	}
	return nMap, err
}
