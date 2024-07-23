package client

import (
	"errors"

	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
)

func (h *V1Client) CloneVirtualMachine(clusterUID, sourceVMName, cloneVMName, namespace string) error {
	cluster, err := h.GetCluster(clusterUID)
	if err != nil {
		return err
	}
	if cluster == nil {
		return errors.New("cluster not found")
	}
	params := clientV1.NewV1SpectroClustersVMCloneParamsWithContext(h.ctx).
		WithUID(clusterUID).
		WithVMName(sourceVMName).
		WithNamespace(namespace).
		WithBody(&models.V1SpectroClusterVMCloneEntity{
			CloneName: &cloneVMName,
		})
	_, err = h.Client.V1SpectroClustersVMClone(params)
	return err
}
