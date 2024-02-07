package client

import (
	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
	"github.com/spectrocloud/palette-sdk-go/client/apiutil"
)

func (h *V1Client) CreateClusterEdge(cluster *models.V1SpectroEdgeClusterEntity) (string, error) {
	params := clientV1.NewV1SpectroClustersEdgeCreateParamsWithContext(h.ctx).
		WithBody(cluster)
	resp, err := h.Client.V1SpectroClustersEdgeCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

func (h *V1Client) CreateMachinePoolEdge(cloudConfigUid string, machinePool *models.V1EdgeMachinePoolConfigEntity) error {
	params := clientV1.NewV1CloudConfigsEdgeMachinePoolCreateParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUid).
		WithBody(machinePool)
	_, err := h.Client.V1CloudConfigsEdgeMachinePoolCreate(params)
	return err
}

func (h *V1Client) UpdateMachinePoolEdge(cloudConfigUid string, machinePool *models.V1EdgeMachinePoolConfigEntity) error {
	params := clientV1.NewV1CloudConfigsEdgeMachinePoolUpdateParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUid).
		WithMachinePoolName(*machinePool.PoolConfig.Name).
		WithBody(machinePool)
	_, err := h.Client.V1CloudConfigsEdgeMachinePoolUpdate(params)
	return err
}

func (h *V1Client) DeleteMachinePoolEdge(cloudConfigUid, machinePoolName string) error {
	params := clientV1.NewV1CloudConfigsEdgeMachinePoolDeleteParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUid).
		WithMachinePoolName(machinePoolName)
	_, err := h.Client.V1CloudConfigsEdgeMachinePoolDelete(params)
	return err
}

func (h *V1Client) GetCloudConfigEdge(configUid string) (*models.V1EdgeCloudConfig, error) {
	params := clientV1.NewV1CloudConfigsEdgeGetParamsWithContext(h.ctx).
		WithConfigUID(configUid)
	resp, err := h.Client.V1CloudConfigsEdgeGet(params)
	if err := apiutil.Handle404(err); err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

func (h *V1Client) GetNodeStatusMapEdge(configUid, machinePoolName string) (map[string]models.V1CloudMachineStatus, error) {
	params := clientV1.NewV1CloudConfigsEdgePoolMachinesListParamsWithContext(h.ctx).
		WithConfigUID(configUid).
		WithMachinePoolName(machinePoolName)
	mpList, err := h.Client.V1CloudConfigsEdgePoolMachinesList(params)
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
