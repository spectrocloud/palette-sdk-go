package client

import (
	"fmt"

	clientv1 "github.com/spectrocloud/palette-sdk-go/api/client/v1"
)

// MigrateVirtualMachineNodeToNode migrates a virtual machine from one node to another.
func (h *V1Client) MigrateVirtualMachineNodeToNode(clusterUID, vmName, vmNamespace string) error {
	cluster, err := h.GetCluster(clusterUID)
	if err != nil {
		return err
	}
	if cluster == nil {
		return fmt.Errorf("cluster with uid %s not found", clusterUID)
	}
	params := clientv1.NewV1SpectroClustersVMMigrateParamsWithContext(h.ctx).
		WithUID(clusterUID).
		WithVMName(vmName).
		WithNamespace(vmNamespace)
	_, err = h.Client.V1SpectroClustersVMMigrate(params)
	return err
}
