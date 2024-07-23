package client

import (
	"fmt"

	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
)

func (h *V1Client) MigrateVirtualMachineNodeToNode(clusterUID, vmName, vmNamespace string) error {
	cluster, err := h.GetCluster(clusterUID)
	if err != nil {
		return err
	}
	if cluster == nil {
		return fmt.Errorf("cluster with uid %s not found", clusterUID)
	}
	params := clientV1.NewV1SpectroClustersVMMigrateParamsWithContext(h.ctx).
		WithUID(clusterUID).
		WithVMName(vmName).
		WithNamespace(vmNamespace)
	_, err = h.Client.V1SpectroClustersVMMigrate(params)
	return err
}
