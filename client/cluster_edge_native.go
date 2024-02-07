package client

import (
	"errors"
	"fmt"

	"github.com/spectrocloud/palette-api-go/apiutil/transport"
	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
)

func (h *V1Client) GetRegistrationToken(scope, tokenName string) (string, error) {
	var params *clientV1.V1EdgeTokensListParams
	switch scope {
	case "project":
		params = clientV1.NewV1EdgeTokensListParamsWithContext(h.Ctx)
	case "tenant":
		params = clientV1.NewV1EdgeTokensListParams()
	default:
		return "", fmt.Errorf("invalid scope: %s", scope)
	}

	resp, err := h.Client.V1EdgeTokensList(params)
	if err != nil {
		return "", err
	}
	tokens := resp.GetPayload()
	if tokens == nil {
		return "", errors.New("failed to list registration tokens")
	}
	for _, token := range tokens.Items {
		if token.Status.IsActive && token.Metadata.Name == tokenName {
			return token.Spec.Token, nil
		}
	}

	return "", nil
}

func (h *V1Client) CreateRegistrationToken(scope, tokenName string, body *models.V1EdgeTokenEntity) (string, error) {
	var params *clientV1.V1EdgeTokensCreateParams
	switch scope {
	case "project":
		params = clientV1.NewV1EdgeTokensCreateParamsWithContext(h.Ctx)
	case "tenant":
		params = clientV1.NewV1EdgeTokensCreateParams()
	default:
		return "", fmt.Errorf("invalid scope: %s", scope)
	}
	params = params.WithBody(body)

	resp, err := h.Client.V1EdgeTokensCreate(params)
	if err != nil {
		return "", err
	}
	uid := resp.GetPayload()
	if uid == nil {
		return "", errors.New("failed to create token")
	}

	return h.GetRegistrationToken(scope, tokenName)

}

func (h *V1Client) GetEdgeHost(scope, edgeHostId string) (*models.V1EdgeHostDevice, error) {
	var params *clientV1.V1EdgeHostDevicesUIDGetParams
	switch scope {
	case "project":
		params = clientV1.NewV1EdgeHostDevicesUIDGetParamsWithContext(h.Ctx)
	case "tenant":
		params = clientV1.NewV1EdgeHostDevicesUIDGetParams()
	default:
		return nil, fmt.Errorf("invalid scope: %s", scope)
	}
	params = params.WithUID(edgeHostId)

	resp, err := h.Client.V1EdgeHostDevicesUIDGet(params)
	if err != nil {
		return nil, err
	}
	return resp.GetPayload(), nil
}

func (h *V1Client) ListEdgeHosts(scope string) ([]*models.V1EdgeHostsMetadata, error) {
	var params *clientV1.V1DashboardEdgehostsSearchParams
	switch scope {
	case "project":
		params = clientV1.NewV1DashboardEdgehostsSearchParamsWithContext(h.Ctx)
	case "tenant":
		params = clientV1.NewV1DashboardEdgehostsSearchParams()
	default:
		return nil, fmt.Errorf("invalid scope: %s", scope)
	}

	resp, err := h.Client.V1DashboardEdgehostsSearch(params)
	if err != nil {
		return nil, err
	}
	return resp.GetPayload().Items, nil
}

func (h *V1Client) CreateClusterEdgeNative(cluster *models.V1SpectroEdgeNativeClusterEntity, scope string) (string, error) {
	var params *clientV1.V1SpectroClustersEdgeNativeCreateParams
	switch scope {
	case "project":
		params = clientV1.NewV1SpectroClustersEdgeNativeCreateParamsWithContext(h.Ctx)
	case "tenant":
		params = clientV1.NewV1SpectroClustersEdgeNativeCreateParams()
	}
	params = params.WithBody(cluster)

	resp, err := h.Client.V1SpectroClustersEdgeNativeCreate(params)
	if err != nil {
		return "", err
	}
	uid := resp.GetPayload()
	if uid == nil {
		return "", errors.New("failed to create cluster")
	}
	return *uid.UID, nil
}

func (h *V1Client) CreateMachinePoolEdgeNative(CloudConfigId, scope string, machinePool *models.V1EdgeNativeMachinePoolConfigEntity) error {
	var params *clientV1.V1CloudConfigsEdgeNativeMachinePoolCreateParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsEdgeNativeMachinePoolCreateParamsWithContext(h.Ctx)
	case "tenant":
		params = clientV1.NewV1CloudConfigsEdgeNativeMachinePoolCreateParams()
	default:
		return fmt.Errorf("invalid scope: %s", scope)
	}
	params = params.WithConfigUID(CloudConfigId).WithBody(machinePool)

	_, err := h.Client.V1CloudConfigsEdgeNativeMachinePoolCreate(params)
	return err
}

func (h *V1Client) UpdateMachinePoolEdgeNative(cloudConfigId, scope string, machinePool *models.V1EdgeNativeMachinePoolConfigEntity) error {
	var params *clientV1.V1CloudConfigsEdgeNativeMachinePoolUpdateParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsEdgeNativeMachinePoolUpdateParamsWithContext(h.Ctx)
	case "tenant":
		params = clientV1.NewV1CloudConfigsEdgeNativeMachinePoolUpdateParams()
	default:
		return fmt.Errorf("invalid scope: %s", scope)
	}
	params = params.WithConfigUID(cloudConfigId).WithMachinePoolName(*machinePool.PoolConfig.Name).WithBody(machinePool)

	_, err := h.Client.V1CloudConfigsEdgeNativeMachinePoolUpdate(params)
	return err
}

func (h *V1Client) GetMachineEdgeNative(scope, machineId, machinePoolName, cloudConfigId string) (*models.V1EdgeNativeMachine, error) {
	var params *clientV1.V1CloudConfigsEdgeNativePoolMachinesUIDGetParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsEdgeNativePoolMachinesUIDGetParamsWithContext(h.Ctx)
	case "tenant":
		params = clientV1.NewV1CloudConfigsEdgeNativePoolMachinesUIDGetParams()
	default:
		return nil, fmt.Errorf("invalid scope: %s", scope)
	}
	params = params.WithConfigUID(cloudConfigId).WithMachinePoolName(machinePoolName).WithMachineUID(machineId)

	resp, err := h.Client.V1CloudConfigsEdgeNativePoolMachinesUIDGet(params)
	if err != nil {
		return nil, err
	}
	return resp.GetPayload(), nil
}

func (h *V1Client) DeleteMachineEdgeNative(scope, clusterId, removeEdgeHostId, machinePool, cloudConfigId string) error {
	var params *clientV1.V1CloudConfigsEdgeNativePoolMachinesUIDDeleteParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsEdgeNativePoolMachinesUIDDeleteParamsWithContext(h.Ctx)
	case "tenant":
		params = clientV1.NewV1CloudConfigsEdgeNativePoolMachinesUIDDeleteParams()
	default:
		return fmt.Errorf("invalid scope: %s", scope)
	}
	params = params.WithConfigUID(cloudConfigId).WithMachinePoolName(machinePool).WithMachineUID(removeEdgeHostId)

	_, err := h.Client.V1CloudConfigsEdgeNativePoolMachinesUIDDelete(params)
	return err
}

func (h *V1Client) ListMachinesInPoolEdgeNative(scope, machinePoolName, cloudConfigId string) ([]*models.V1EdgeNativeMachine, error) {
	var params *clientV1.V1CloudConfigsEdgeNativePoolMachinesListParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsEdgeNativePoolMachinesListParamsWithContext(h.Ctx)
	case "tenant":
		params = clientV1.NewV1CloudConfigsEdgeNativePoolMachinesListParams()
	default:
		return nil, fmt.Errorf("invalid scope: %s", scope)
	}
	params = params.WithConfigUID(cloudConfigId).WithMachinePoolName(machinePoolName)

	resp, err := h.Client.V1CloudConfigsEdgeNativePoolMachinesList(params)
	if err != nil {
		return nil, err
	}

	return resp.GetPayload().Items, nil
}

func (h *V1Client) DeleteMachinePoolEdgeNative(cloudConfigId, machinePoolName, scope string) error {
	var params *clientV1.V1CloudConfigsEdgeNativeMachinePoolDeleteParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsEdgeNativeMachinePoolDeleteParamsWithContext(h.Ctx)
	case "tenant":
		params = clientV1.NewV1CloudConfigsEdgeNativeMachinePoolDeleteParams()
	}
	params = params.WithConfigUID(cloudConfigId).WithMachinePoolName(machinePoolName)

	_, err := h.Client.V1CloudConfigsEdgeNativeMachinePoolDelete(params)
	return err
}

func (h *V1Client) GetCloudConfigEdgeNative(configUID, scope string) (*models.V1EdgeNativeCloudConfig, error) {
	var params *clientV1.V1CloudConfigsEdgeNativeGetParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsEdgeNativeGetParamsWithContext(h.Ctx)
	case "tenant":
		params = clientV1.NewV1CloudConfigsEdgeNativeGetParams()
	default:
		return nil, fmt.Errorf("invalid scope: %s", scope)
	}
	params = params.WithConfigUID(configUID)

	resp, err := h.Client.V1CloudConfigsEdgeNativeGet(params)
	var e *transport.TransportError
	if errors.As(err, &e) && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (h *V1Client) GetNodeStatusMapEdgeNative(configUID, machinePoolName, scope string) (map[string]models.V1CloudMachineStatus, error) {
	var params *clientV1.V1CloudConfigsEdgeNativePoolMachinesListParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsEdgeNativePoolMachinesListParamsWithContext(h.Ctx)
	case "tenant":
		params = clientV1.NewV1CloudConfigsEdgeNativePoolMachinesListParams()
	default:
		return nil, fmt.Errorf("invalid scope: %s", scope)
	}
	params = params.WithConfigUID(configUID).WithMachinePoolName(machinePoolName)

	mpList, err := h.Client.V1CloudConfigsEdgeNativePoolMachinesList(params)
	if err != nil {
		return nil, err
	}

	nMap := map[string]models.V1CloudMachineStatus{}
	if len(mpList.Payload.Items) > 0 {
		for _, node := range mpList.Payload.Items {
			nMap[node.Metadata.UID] = *node.Status
		}
	}
	return nMap, nil
}
