package client

import (
	clientv1 "github.com/spectrocloud/palette-sdk-go/api/client/v1"
	"github.com/spectrocloud/palette-sdk-go/api/models"
	"github.com/spectrocloud/palette-sdk-go/client/apiutil"
)

// CreateClusterVsphere creates a new vSphere cluster.
func (h *V1Client) CreateClusterVsphere(cluster *models.V1SpectroVsphereClusterEntity) (string, error) {
	params := clientv1.NewV1SpectroClustersVsphereCreateParamsWithContext(h.ctx).
		WithBody(cluster)
	resp, err := h.Client.V1SpectroClustersVsphereCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

// CreateMachinePoolVsphere creates a new vSphere machine pool.
func (h *V1Client) CreateMachinePoolVsphere(cloudConfigUID string, machinePool *models.V1VsphereMachinePoolConfigEntity) error {
	params := clientv1.NewV1CloudConfigsVsphereMachinePoolCreateParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUID).
		WithBody(machinePool)
	_, err := h.Client.V1CloudConfigsVsphereMachinePoolCreate(params)
	return err
}

// UpdateMachinePoolVsphere updates an existing vSphere machine pool.
func (h *V1Client) UpdateMachinePoolVsphere(cloudConfigUID string, machinePool *models.V1VsphereMachinePoolConfigEntity) error {
	params := clientv1.NewV1CloudConfigsVsphereMachinePoolUpdateParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUID).
		WithMachinePoolName(*machinePool.PoolConfig.Name).
		WithBody(machinePool)
	_, err := h.Client.V1CloudConfigsVsphereMachinePoolUpdate(params)
	return err
}

// DeleteMachinePoolVsphere deletes an existing vSphere machine pool.
func (h *V1Client) DeleteMachinePoolVsphere(cloudConfigUID, machinePoolName string) error {
	params := clientv1.NewV1CloudConfigsVsphereMachinePoolDeleteParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUID).
		WithMachinePoolName(machinePoolName)
	_, err := h.Client.V1CloudConfigsVsphereMachinePoolDelete(params)
	return err
}

// GetCloudConfigVsphere retrieves an existing vSphere cluster's cloud config.
func (h *V1Client) GetCloudConfigVsphere(uid string) (*models.V1VsphereCloudConfig, error) {
	params := clientv1.NewV1CloudConfigsVsphereGetParamsWithContext(h.ctx).
		WithConfigUID(uid)
	resp, err := h.Client.V1CloudConfigsVsphereGet(params)
	if apiutil.Is404(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// UpdateCloudConfigVsphere updates an existing vSphere cluster's cloud config.
func (h *V1Client) UpdateCloudConfigVsphere(uid string, cloudConfig *models.V1VsphereCloudClusterConfigEntity) error {
	params := clientv1.NewV1CloudConfigsVsphereUIDClusterConfigParamsWithContext(h.ctx).
		WithConfigUID(uid).
		WithBody(cloudConfig)
	_, err := h.Client.V1CloudConfigsVsphereUIDClusterConfig(params)
	return err
}

// ImportClusterVsphere imports an existing vSphere cluster.
func (h *V1Client) ImportClusterVsphere(meta *models.V1ObjectMetaInputEntity) (string, error) {
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

// GetNodeStatusMapVsphere retrieves the status of all nodes in a vSphere machine pool.
func (h *V1Client) GetNodeStatusMapVsphere(configUID, machinePoolName string) (map[string]models.V1CloudMachineStatus, error) {
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
