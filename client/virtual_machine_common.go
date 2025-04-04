package client

import (
	clientv1 "github.com/spectrocloud/palette-sdk-go/api/client/version1"
	"github.com/spectrocloud/palette-sdk-go/api/models"
)

// IsVMExists checks if a virtual machine exists.
func (h *V1Client) IsVMExists(clusterUID, name, namespace string) (bool, error) {
	vm, err := h.GetVirtualMachine(clusterUID, namespace, name)
	if err != nil {
		return false, err
	}
	if vm != nil {
		return true, nil
	}
	return false, nil
}

// GetVirtualMachineWithoutStatus retrieves a virtual machine, regardless of its status.
func (h *V1Client) GetVirtualMachineWithoutStatus(uid, name, namespace string) (*models.V1ClusterVirtualMachine, error) {
	params := clientv1.NewV1SpectroClustersVMGetParamsWithContext(h.ctx).
		WithUID(uid).
		WithVMName(name).
		WithNamespace(namespace)
	resp, err := h.Client.V1SpectroClustersVMGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// GetVirtualMachines retrieves a list of virtual machines for a given cluster.
func (h *V1Client) GetVirtualMachines(cluster *models.V1SpectroCluster) ([]*models.V1ClusterVirtualMachine, error) {
	params := clientv1.NewV1SpectroClustersVMListParams().
		WithContext(ContextForScope(cluster.Metadata.Annotations[Scope], h.projectUID)).
		WithUID(cluster.Metadata.UID)
	resp, err := h.Client.V1SpectroClustersVMList(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload.Items, nil
}
