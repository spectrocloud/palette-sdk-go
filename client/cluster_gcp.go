package client

import (
	clientv1 "github.com/spectrocloud/palette-sdk-go/api/client/version1"
	"github.com/spectrocloud/palette-sdk-go/api/models"
	"github.com/spectrocloud/palette-sdk-go/client/apiutil"
)

// CreateClusterGcp creates a new GCP IaaS cluster.
func (h *V1Client) CreateClusterGcp(cluster *models.V1SpectroGcpClusterEntity) (string, error) {
	params := clientv1.NewV1SpectroClustersGcpCreateParamsWithContext(h.ctx).
		WithBody(cluster)
	resp, err := h.Client.V1SpectroClustersGcpCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

// CreateMachinePoolGcp creates a new GCP IaaS machine pool.
func (h *V1Client) CreateMachinePoolGcp(cloudConfigUID string, machinePool *models.V1GcpMachinePoolConfigEntity) error {
	params := clientv1.NewV1CloudConfigsGcpMachinePoolCreateParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUID).
		WithBody(machinePool)
	_, err := h.Client.V1CloudConfigsGcpMachinePoolCreate(params)
	return err
}

// UpdateMachinePoolGcp updates an existing GCP IaaS machine pool.
func (h *V1Client) UpdateMachinePoolGcp(cloudConfigUID string, machinePool *models.V1GcpMachinePoolConfigEntity) error {
	params := clientv1.NewV1CloudConfigsGcpMachinePoolUpdateParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUID).
		WithMachinePoolName(*machinePool.PoolConfig.Name).
		WithBody(machinePool)
	_, err := h.Client.V1CloudConfigsGcpMachinePoolUpdate(params)
	return err
}

// DeleteMachinePoolGcp deletes an existing GCP IaaS machine pool.
func (h *V1Client) DeleteMachinePoolGcp(cloudConfigUID, machinePoolName string) error {
	params := clientv1.NewV1CloudConfigsGcpMachinePoolDeleteParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUID).
		WithMachinePoolName(machinePoolName)
	_, err := h.Client.V1CloudConfigsGcpMachinePoolDelete(params)
	return err
}

// GetCloudConfigGcp retrieves an existing GCP IaaS cluster's cloud config.
func (h *V1Client) GetCloudConfigGcp(configUID string) (*models.V1GcpCloudConfig, error) {
	params := clientv1.NewV1CloudConfigsGcpGetParamsWithContext(h.ctx).
		WithConfigUID(configUID)
	resp, err := h.Client.V1CloudConfigsGcpGet(params)
	if apiutil.Is404(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// ImportClusterGcp imports an existing GCP IaaS cluster.
func (h *V1Client) ImportClusterGcp(meta *models.V1ObjectMetaInputEntity) (string, error) {
	params := clientv1.NewV1SpectroClustersGcpImportParamsWithContext(h.ctx).
		WithBody(&models.V1SpectroGcpClusterImportEntity{
			Metadata: meta,
		},
		)
	resp, err := h.Client.V1SpectroClustersGcpImport(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

// PostSpectroClusterGcpImport performs a POST operation to import a GCP cluster.
// This is an abstract method that accepts the full V1SpectroGcpClusterImportEntity model.
// Returns the created cluster UID on success.
func (h *V1Client) PostSpectroClusterGcpImport(entity *models.V1SpectroGcpClusterImportEntity) (string, error) {
	params := clientv1.NewV1SpectroClustersGcpImportParamsWithContext(h.ctx).
		WithBody(entity)
	resp, err := h.Client.V1SpectroClustersGcpImport(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

// GetNodeStatusMapGcp retrieves the status of all nodes in a GCP IaaS machine pool.
func (h *V1Client) GetNodeStatusMapGcp(configUID, machinePoolName string) (map[string]models.V1CloudMachineStatus, error) {
	params := clientv1.NewV1CloudConfigsGcpPoolMachinesListParamsWithContext(h.ctx).
		WithConfigUID(configUID).
		WithMachinePoolName(machinePoolName)
	mpList, err := h.Client.V1CloudConfigsGcpPoolMachinesList(params)
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
