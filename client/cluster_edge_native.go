package client

import (
	"errors"

	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
	"github.com/spectrocloud/palette-sdk-go/client/apiutil"
)

func (h *V1Client) GetRegistrationToken(tokenName string) (string, error) {
	params := clientV1.NewV1EdgeTokensListParamsWithContext(h.ctx)
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

func (h *V1Client) CreateRegistrationToken(tokenName string, body *models.V1EdgeTokenEntity) (string, error) {
	params := clientV1.NewV1EdgeTokensCreateParamsWithContext(h.ctx).
		WithBody(body)
	_, err := h.Client.V1EdgeTokensCreate(params)
	if err != nil {
		return "", err
	}
	return h.GetRegistrationToken(tokenName)

}

func (h *V1Client) GetEdgeHost(edgeHostId string) (*models.V1EdgeHostDevice, error) {
	params := clientV1.NewV1EdgeHostDevicesUIDGetParamsWithContext(h.ctx).
		WithUID(edgeHostId)
	resp, err := h.Client.V1EdgeHostDevicesUIDGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

func (h *V1Client) ListEdgeHosts() ([]*models.V1EdgeHostsMetadata, error) {
	params := clientV1.NewV1DashboardEdgehostsSearchParamsWithContext(h.ctx)
	resp, err := h.Client.V1DashboardEdgehostsSearch(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload.Items, nil
}

func (h *V1Client) CreateClusterEdgeNative(cluster *models.V1SpectroEdgeNativeClusterEntity) (string, error) {
	params := clientV1.NewV1SpectroClustersEdgeNativeCreateParamsWithContext(h.ctx).
		WithBody(cluster)
	resp, err := h.Client.V1SpectroClustersEdgeNativeCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

func (h *V1Client) CreateMachinePoolEdgeNative(cloudConfigUid string, machinePool *models.V1EdgeNativeMachinePoolConfigEntity) error {
	params := clientV1.NewV1CloudConfigsEdgeNativeMachinePoolCreateParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUid).
		WithBody(machinePool)
	_, err := h.Client.V1CloudConfigsEdgeNativeMachinePoolCreate(params)
	return err
}

func (h *V1Client) UpdateMachinePoolEdgeNative(cloudConfigUid string, machinePool *models.V1EdgeNativeMachinePoolConfigEntity) error {
	params := clientV1.NewV1CloudConfigsEdgeNativeMachinePoolUpdateParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUid).
		WithMachinePoolName(*machinePool.PoolConfig.Name).
		WithBody(machinePool)
	_, err := h.Client.V1CloudConfigsEdgeNativeMachinePoolUpdate(params)
	return err
}

func (h *V1Client) GetMachineEdgeNative(machineUid, machinePoolName, cloudConfigUid string) (*models.V1EdgeNativeMachine, error) {
	params := clientV1.NewV1CloudConfigsEdgeNativePoolMachinesUIDGetParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUid).
		WithMachinePoolName(machinePoolName).
		WithMachineUID(machineUid)
	resp, err := h.Client.V1CloudConfigsEdgeNativePoolMachinesUIDGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

func (h *V1Client) DeleteMachineEdgeNative(clusterUid, edgeHostUid, machinePool, cloudConfigUid string) error {
	params := clientV1.NewV1CloudConfigsEdgeNativePoolMachinesUIDDeleteParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUid).
		WithMachinePoolName(machinePool).
		WithMachineUID(edgeHostUid)
	_, err := h.Client.V1CloudConfigsEdgeNativePoolMachinesUIDDelete(params)
	return err
}

func (h *V1Client) ListMachinesInPoolEdgeNative(machinePoolName, cloudConfigUid string) ([]*models.V1EdgeNativeMachine, error) {
	params := clientV1.NewV1CloudConfigsEdgeNativePoolMachinesListParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUid).
		WithMachinePoolName(machinePoolName)
	resp, err := h.Client.V1CloudConfigsEdgeNativePoolMachinesList(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload.Items, nil
}

func (h *V1Client) DeleteMachinePoolEdgeNative(cloudConfigUid, machinePoolName string) error {
	params := clientV1.NewV1CloudConfigsEdgeNativeMachinePoolDeleteParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUid).
		WithMachinePoolName(machinePoolName)
	_, err := h.Client.V1CloudConfigsEdgeNativeMachinePoolDelete(params)
	return err
}

func (h *V1Client) GetCloudConfigEdgeNative(configUid string) (*models.V1EdgeNativeCloudConfig, error) {
	params := clientV1.NewV1CloudConfigsEdgeNativeGetParamsWithContext(h.ctx).
		WithConfigUID(configUid)
	resp, err := h.Client.V1CloudConfigsEdgeNativeGet(params)
	if err := apiutil.Handle404(err); err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

func (h *V1Client) GetNodeStatusMapEdgeNative(configUid, machinePoolName string) (map[string]models.V1CloudMachineStatus, error) {
	params := clientV1.NewV1CloudConfigsEdgeNativePoolMachinesListParamsWithContext(h.ctx).
		WithConfigUID(configUid).
		WithMachinePoolName(machinePoolName)
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
