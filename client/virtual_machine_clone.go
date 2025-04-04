package client

import (
	"errors"

	clientv1 "github.com/spectrocloud/palette-sdk-go/api/client/version1"
	"github.com/spectrocloud/palette-sdk-go/api/models"
)

// CloneVirtualMachine clones a virtual machine.
func (h *V1Client) CloneVirtualMachine(clusterUID, sourceVMName, cloneVMName, namespace string) error {
	cluster, err := h.GetCluster(clusterUID)
	if err != nil {
		return err
	}
	if cluster == nil {
		return errors.New("cluster not found")
	}
	params := clientv1.NewV1SpectroClustersVMCloneParamsWithContext(h.ctx).
		WithUID(clusterUID).
		WithVMName(sourceVMName).
		WithNamespace(namespace).
		WithBody(&models.V1SpectroClusterVMCloneEntity{
			CloneName: &cloneVMName,
		})
	_, err = h.Client.V1SpectroClustersVMClone(params)
	return err
}
