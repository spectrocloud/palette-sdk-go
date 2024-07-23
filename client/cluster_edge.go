package client

import (
	clientv1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
	"github.com/spectrocloud/palette-sdk-go/client/apiutil"
)

// TODO: edgev1 deprecation

// CreateClusterEdge creates a new edge cluster.
func (h *V1Client) CreateClusterEdge(cluster *models.V1SpectroEdgeClusterEntity) (string, error) {
	params := clientv1.NewV1SpectroClustersEdgeCreateParamsWithContext(h.ctx).
		WithBody(cluster)
	resp, err := h.Client.V1SpectroClustersEdgeCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

// CreateMachinePoolEdge creates a new edge machine pool.
func (h *V1Client) CreateMachinePoolEdge(cloudConfigUID string, machinePool *models.V1EdgeMachinePoolConfigEntity) error {
	params := clientv1.NewV1CloudConfigsEdgeMachinePoolCreateParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUID).
		WithBody(machinePool)
	_, err := h.Client.V1CloudConfigsEdgeMachinePoolCreate(params)
	return err
}

// UpdateMachinePoolEdge updates an existing edge machine pool.
func (h *V1Client) UpdateMachinePoolEdge(cloudConfigUID string, machinePool *models.V1EdgeMachinePoolConfigEntity) error {
	params := clientv1.NewV1CloudConfigsEdgeMachinePoolUpdateParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUID).
		WithMachinePoolName(*machinePool.PoolConfig.Name).
		WithBody(machinePool)
	_, err := h.Client.V1CloudConfigsEdgeMachinePoolUpdate(params)
	return err
}

// DeleteMachinePoolEdge deletes an existing edge machine pool.
func (h *V1Client) DeleteMachinePoolEdge(cloudConfigUID, machinePoolName string) error {
	params := clientv1.NewV1CloudConfigsEdgeMachinePoolDeleteParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUID).
		WithMachinePoolName(machinePoolName)
	_, err := h.Client.V1CloudConfigsEdgeMachinePoolDelete(params)
	return err
}

// GetCloudConfigEdge retrieves an existing edge cluster's cloud config.
func (h *V1Client) GetCloudConfigEdge(configUID string) (*models.V1EdgeCloudConfig, error) {
	params := clientv1.NewV1CloudConfigsEdgeGetParamsWithContext(h.ctx).
		WithConfigUID(configUID)
	resp, err := h.Client.V1CloudConfigsEdgeGet(params)
	if apiutil.Is404(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// GetNodeStatusMapEdge retrieves the status of all nodes in an edge machine pool.
func (h *V1Client) GetNodeStatusMapEdge(configUID, machinePoolName string) (map[string]models.V1CloudMachineStatus, error) {
	params := clientv1.NewV1CloudConfigsEdgePoolMachinesListParamsWithContext(h.ctx).
		WithConfigUID(configUID).
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
