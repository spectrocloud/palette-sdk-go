package client

import (
	clientv1 "github.com/spectrocloud/palette-sdk-go/api/client/version1"
	"github.com/spectrocloud/palette-sdk-go/api/models"
	"github.com/spectrocloud/palette-sdk-go/client/apiutil"
)

// CreateClusterCloudStack creates a new CloudStack cluster.
func (h *V1Client) CreateClusterCloudStack(cluster *models.V1SpectroCloudStackClusterEntity) (string, error) {
	params := clientv1.NewV1SpectroClustersCloudStackCreateParamsWithContext(h.ctx).
		WithBody(cluster)
	resp, err := h.Client.V1SpectroClustersCloudStackCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

// CreateMachinePoolCloudStack creates a new CloudStack machine pool.
func (h *V1Client) CreateMachinePoolCloudStack(cloudConfigUID string, machinePool *models.V1CloudStackMachinePoolConfigEntity) error {
	params := clientv1.NewV1CloudConfigsCloudStackMachinePoolCreateParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUID).
		WithBody(machinePool)
	_, err := h.Client.V1CloudConfigsCloudStackMachinePoolCreate(params)
	return err
}

// UpdateMachinePoolCloudStack updates an existing CloudStack machine pool.
func (h *V1Client) UpdateMachinePoolCloudStack(cloudConfigUID string, machinePool *models.V1CloudStackMachinePoolConfigEntity) error {
	params := clientv1.NewV1CloudConfigsCloudStackMachinePoolUpdateParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUID).
		WithMachinePoolName(*machinePool.PoolConfig.Name).
		WithBody(machinePool)
	_, err := h.Client.V1CloudConfigsCloudStackMachinePoolUpdate(params)
	return err
}

// DeleteMachinePoolCloudStack deletes an existing CloudStack machine pool.
func (h *V1Client) DeleteMachinePoolCloudStack(cloudConfigUID, machinePoolName string) error {
	params := clientv1.NewV1CloudConfigsCloudStackMachinePoolDeleteParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUID).
		WithMachinePoolName(machinePoolName)
	_, err := h.Client.V1CloudConfigsCloudStackMachinePoolDelete(params)
	return err
}

// GetCloudConfigCloudStack retrieves an existing CloudStack cluster's cloud config.
func (h *V1Client) GetCloudConfigCloudStack(configUID string) (*models.V1CloudStackCloudConfig, error) {
	params := clientv1.NewV1CloudConfigsCloudStackGetParamsWithContext(h.ctx).
		WithConfigUID(configUID)
	resp, err := h.Client.V1CloudConfigsCloudStackGet(params)
	if apiutil.Is404(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// ImportClusterCloudStack imports an existing CloudStack cluster.
func (h *V1Client) ImportClusterCloudStack(meta *models.V1ObjectMetaInputEntity) (string, error) {
	params := clientv1.NewV1SpectroClustersCloudStackImportParamsWithContext(h.ctx).
		WithBody(&models.V1SpectroCloudStackClusterImportEntity{
			Metadata: meta,
		},
		)
	resp, err := h.Client.V1SpectroClustersCloudStackImport(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

// GetNodeStatusMapCloudStack retrieves the status of all nodes in a CloudStack machine pool.
func (h *V1Client) GetNodeStatusMapCloudStack(configUID, machinePoolName string) (map[string]models.V1CloudMachineStatus, error) {
	params := clientv1.NewV1CloudConfigsCloudStackPoolMachinesListParamsWithContext(h.ctx).
		WithConfigUID(configUID).
		WithMachinePoolName(machinePoolName)
	mpList, err := h.Client.V1CloudConfigsCloudStackPoolMachinesList(params)
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
