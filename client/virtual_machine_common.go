package client

import (
	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
)

func (h *V1Client) IsVMExists(clusterUid, name, namespace string) (bool, error) {
	vm, err := h.GetVirtualMachine(clusterUid, namespace, name)
	if err != nil {
		return false, err
	}
	if vm != nil {
		return true, nil
	}
	return false, nil
}

func (h *V1Client) GetVirtualMachineWithoutStatus(uid, name, namespace string) (*models.V1ClusterVirtualMachine, error) {
	params := clientV1.NewV1SpectroClustersVMGetParamsWithContext(h.ctx).
		WithUID(uid).
		WithVMName(name).
		WithNamespace(namespace)
	resp, err := h.Client.V1SpectroClustersVMGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

func (h *V1Client) GetVirtualMachines(cluster *models.V1SpectroCluster) ([]*models.V1ClusterVirtualMachine, error) {
	params := clientV1.NewV1SpectroClustersVMListParams().
		WithContext(ContextForScope(cluster.Metadata.Annotations[Scope], h.projectUid)).
		WithUID(cluster.Metadata.UID)
	resp, err := h.Client.V1SpectroClustersVMList(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload.Items, nil
}
