package client

import (
	clientv1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
	"github.com/spectrocloud/palette-sdk-go/client/apiutil"
)

// CreateClusterVirtual creates a new virtual cluster.
func (h *V1Client) CreateClusterVirtual(cluster *models.V1SpectroVirtualClusterEntity) (string, error) {
	params := clientv1.NewV1SpectroClustersVirtualCreateParamsWithContext(h.ctx).
		WithBody(cluster)
	resp, err := h.Client.V1SpectroClustersVirtualCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

// ResizeClusterVirtual resizes a virtual cluster's CPU, RAM, and storage configuration.
func (h *V1Client) ResizeClusterVirtual(configUID string, body *models.V1VirtualClusterResize) error {
	params := clientv1.NewV1CloudConfigsVirtualUIDUpdateParamsWithContext(h.ctx).
		WithConfigUID(configUID).
		WithBody(body)
	_, err := h.Client.V1CloudConfigsVirtualUIDUpdate(params)
	return err
}

// CreateMachinePoolVirtual creates a new virtual machine pool.
// TODO: deprecate unused virtual cluster functions
func (h *V1Client) CreateMachinePoolVirtual(cloudConfigUID string, machinePool *models.V1VirtualMachinePoolConfigEntity) error {
	params := clientv1.NewV1CloudConfigsVirtualMachinePoolCreateParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUID).
		WithBody(machinePool)
	_, err := h.Client.V1CloudConfigsVirtualMachinePoolCreate(params)
	return err
}

// UpdateMachinePoolVirtual updates an existing virtual machine pool.
// TODO: deprecate unused virtual cluster functions
func (h *V1Client) UpdateMachinePoolVirtual(cloudConfigUID string, machinePool *models.V1VirtualMachinePoolConfigEntity) error {
	params := clientv1.NewV1CloudConfigsVirtualMachinePoolUpdateParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUID).
		WithBody(machinePool)
	_, err := h.Client.V1CloudConfigsVirtualMachinePoolUpdate(params)
	return err
}

// DeleteMachinePoolVirtual deletes an existing virtual machine pool.
// TODO: deprecate unused virtual cluster functions
func (h *V1Client) DeleteMachinePoolVirtual(cloudConfigUID, machinePoolName string) error {
	params := clientv1.NewV1CloudConfigsVirtualMachinePoolDeleteParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUID).
		WithMachinePoolName(machinePoolName)
	_, err := h.Client.V1CloudConfigsVirtualMachinePoolDelete(params)
	return err
}

// GetCloudConfigVirtual retrieves an existing virtual cluster's cloud config.
func (h *V1Client) GetCloudConfigVirtual(configUID string) (*models.V1VirtualCloudConfig, error) {
	params := clientv1.NewV1CloudConfigsVirtualGetParamsWithContext(h.ctx).
		WithConfigUID(configUID)
	resp, err := h.Client.V1CloudConfigsVirtualGet(params)
	if apiutil.Is404(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// VirtualClusterLifecycleConfigChange pauses or unpauses an existing virtual cluster.
func (h *V1Client) VirtualClusterLifecycleConfigChange(uid string, body *models.V1LifecycleConfigEntity) (string, error) {
	params := clientv1.NewV1SpectroClustersUIDLifecycleConfigUpdateParamsWithContext(h.ctx).
		WithUID(uid).
		WithBody(body)
	_, err := h.Client.V1SpectroClustersUIDLifecycleConfigUpdate(params)
	if err != nil {
		return "Fail", err
	}
	return "Success", nil
}

// GetNodeStatusMapVirtual retrieves the status of all nodes in a virtual machine pool.
func (h *V1Client) GetNodeStatusMapVirtual(configUID, machinePoolName string) (map[string]models.V1CloudMachineStatus, error) {
	params := clientv1.NewV1CloudConfigsVirtualPoolMachinesListParamsWithContext(h.ctx).
		WithConfigUID(configUID).
		WithMachinePoolName(machinePoolName)
	mpList, err := h.Client.V1CloudConfigsVirtualPoolMachinesList(params)
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
