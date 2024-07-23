package client

import (
	clientv1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
	"github.com/spectrocloud/palette-sdk-go/client/apiutil"
)

// TODO: edgev1 deprecation

// CreateClusterEdgeVsphere creates a vSphere edge cluster.
func (h *V1Client) CreateClusterEdgeVsphere(cluster *models.V1SpectroVsphereClusterEntity) (string, error) {
	params := clientv1.NewV1SpectroClustersVsphereCreateParamsWithContext(h.ctx).
		WithBody(cluster)
	resp, err := h.Client.V1SpectroClustersVsphereCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

// CreateMachinePoolEdgeVsphere creates a new vSphere edge machine pool.
func (h *V1Client) CreateMachinePoolEdgeVsphere(cloudConfigUID string, machinePool *models.V1VsphereMachinePoolConfigEntity) error {
	params := clientv1.NewV1CloudConfigsVsphereMachinePoolCreateParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUID).
		WithBody(machinePool)
	_, err := h.Client.V1CloudConfigsVsphereMachinePoolCreate(params)
	return err
}

// UpdateMachinePoolEdgeVsphere updates an existing vSphere edge machine pool.
func (h *V1Client) UpdateMachinePoolEdgeVsphere(cloudConfigUID string, machinePool *models.V1VsphereMachinePoolConfigEntity) error {
	params := clientv1.NewV1CloudConfigsVsphereMachinePoolUpdateParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUID).
		WithMachinePoolName(*machinePool.PoolConfig.Name).
		WithBody(machinePool)
	_, err := h.Client.V1CloudConfigsVsphereMachinePoolUpdate(params)
	return err
}

// DeleteMachinePoolEdgeVsphere deletes an existing vSphere edge machine pool.
func (h *V1Client) DeleteMachinePoolEdgeVsphere(cloudConfigUID, machinePoolName string) error {
	params := clientv1.NewV1CloudConfigsVsphereMachinePoolDeleteParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUID).
		WithMachinePoolName(machinePoolName)
	_, err := h.Client.V1CloudConfigsVsphereMachinePoolDelete(params)
	return err
}

// GetCloudConfigEdgeVsphere retrieves an existing vSphere edge cluster's cloud config.
func (h *V1Client) GetCloudConfigEdgeVsphere(configUID string) (*models.V1VsphereCloudConfig, error) {
	params := clientv1.NewV1CloudConfigsVsphereGetParamsWithContext(h.ctx).
		WithConfigUID(configUID)
	resp, err := h.Client.V1CloudConfigsVsphereGet(params)
	if apiutil.Is404(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// ImportClusterEdgeVsphere imports an existing vSphere edge cluster.
func (h *V1Client) ImportClusterEdgeVsphere(meta *models.V1ObjectMetaInputEntity) (string, error) {
	params := clientv1.NewV1SpectroClustersVsphereImportParamsWithContext(h.ctx).
		WithBody(&models.V1SpectroVsphereClusterImportEntity{
			Metadata: meta,
		},
		)
	resp, err := h.Client.V1SpectroClustersVsphereImport(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

// GetNodeStatusMapEdgeVsphere retrieves the status of all nodes in a vSphere edge machine pool.
func (h *V1Client) GetNodeStatusMapEdgeVsphere(configUID, machinePoolName string) (map[string]models.V1CloudMachineStatus, error) {
	params := clientv1.NewV1CloudConfigsVspherePoolMachinesListParamsWithContext(h.ctx).
		WithConfigUID(configUID).
		WithMachinePoolName(machinePoolName)
	mpList, err := h.Client.V1CloudConfigsVspherePoolMachinesList(params)
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
