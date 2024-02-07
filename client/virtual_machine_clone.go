package client

import (
	"errors"

	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
)

func (h *V1Client) CloneVirtualMachine(clusterUid, sourceVmName, cloneVmName, namespace string) error {
	cluster, err := h.GetCluster(clusterUid)
	if err != nil {
		return err
	}
	if cluster == nil {
		return errors.New("cluster not found")
	}
	params := clientV1.NewV1SpectroClustersVMCloneParamsWithContext(h.ctx).
		WithUID(clusterUid).
		WithVMName(sourceVmName).
		WithNamespace(namespace).
		WithBody(&models.V1SpectroClusterVMCloneEntity{
			CloneName: &cloneVmName,
		})
	_, err = h.Client.V1SpectroClustersVMClone(params)
	return err
}
