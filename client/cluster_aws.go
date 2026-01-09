package client

import (
	clientv1 "github.com/spectrocloud/palette-sdk-go/api/client/version1"
	"github.com/spectrocloud/palette-sdk-go/api/models"
	"github.com/spectrocloud/palette-sdk-go/client/apiutil"
)

// CreateClusterAws creates a new AWS IaaS cluster.
func (h *V1Client) CreateClusterAws(cluster *models.V1SpectroAwsClusterEntity) (string, error) {
	params := clientv1.NewV1SpectroClustersAwsCreateParamsWithContext(h.ctx).
		WithBody(cluster)
	resp, err := h.Client.V1SpectroClustersAwsCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

// CreateMachinePoolAws creates a new AWS IaaS machine pool.
func (h *V1Client) CreateMachinePoolAws(cloudConfigUID string, machinePool *models.V1AwsMachinePoolConfigEntity) error {
	params := clientv1.NewV1CloudConfigsAwsMachinePoolCreateParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUID).
		WithBody(machinePool)
	_, err := h.Client.V1CloudConfigsAwsMachinePoolCreate(params)
	return err
}

// UpdateMachinePoolAws updates an existing AWS IaaS machine pool.
func (h *V1Client) UpdateMachinePoolAws(cloudConfigUID string, machinePool *models.V1AwsMachinePoolConfigEntity) error {
	params := clientv1.NewV1CloudConfigsAwsMachinePoolUpdateParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUID).
		WithMachinePoolName(*machinePool.PoolConfig.Name).
		WithBody(machinePool)
	_, err := h.Client.V1CloudConfigsAwsMachinePoolUpdate(params)
	return err
}

// DeleteMachinePoolAws deletes an existing AWS IaaS machine pool.
func (h *V1Client) DeleteMachinePoolAws(cloudConfigUID, machinePoolName string) error {
	params := clientv1.NewV1CloudConfigsAwsMachinePoolDeleteParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUID).
		WithMachinePoolName(machinePoolName)
	_, err := h.Client.V1CloudConfigsAwsMachinePoolDelete(params)
	return err
}

// GetCloudConfigAws retrieves an existing AWS IaaS cluster's cloud config.
func (h *V1Client) GetCloudConfigAws(configUID string) (*models.V1AwsCloudConfig, error) {
	params := clientv1.NewV1CloudConfigsAwsGetParamsWithContext(h.ctx).
		WithConfigUID(configUID)
	resp, err := h.Client.V1CloudConfigsAwsGet(params)
	if apiutil.Is404(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// ImportClusterAws imports an existing AWS IaaS cluster.
func (h *V1Client) ImportClusterAws(meta *models.V1ObjectMetaInputEntity) (string, error) {
	params := clientv1.NewV1SpectroClustersAwsImportParamsWithContext(h.ctx).
		WithBody(&models.V1SpectroAwsClusterImportEntity{
			Metadata: meta,
		},
		)
	resp, err := h.Client.V1SpectroClustersAwsImport(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

// PostSpectroClusterAwsImport performs a POST operation to import an AWS cluster.
// This is an abstract method that accepts the full V1SpectroAwsClusterImportEntity model.
// Returns the created cluster UID on success.
func (h *V1Client) PostSpectroClusterAwsImport(entity *models.V1SpectroAwsClusterImportEntity) (string, error) {
	params := clientv1.NewV1SpectroClustersAwsImportParamsWithContext(h.ctx).
		WithBody(entity)
	resp, err := h.Client.V1SpectroClustersAwsImport(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

// GetNodeStatusMapAws retrieves the status of all nodes in an AWS IaaS machine pool.
func (h *V1Client) GetNodeStatusMapAws(configUID, machinePoolName string) (map[string]models.V1CloudMachineStatus, error) {
	params := clientv1.NewV1CloudConfigsAwsPoolMachinesListParamsWithContext(h.ctx).
		WithConfigUID(configUID).
		WithMachinePoolName(machinePoolName)
	mpList, err := h.Client.V1CloudConfigsAwsPoolMachinesList(params)
	if err != nil {
		return nil, err
	}
	nMap := map[string]models.V1CloudMachineStatus{}
	if len(mpList.Payload.Items) > 0 {
		for _, node := range mpList.Payload.Items {
			if node.Status.MaintenanceStatus.Action != "" {
				nMap[node.Metadata.UID] = *node.Status
			}
		}
	}
	return nMap, nil
}
