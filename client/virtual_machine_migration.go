package client

import (
	"errors"
	"fmt"

	clusterC "github.com/spectrocloud/hapi/spectrocluster/client/v1"
)

func (h *V1Client) MigrateVirtualMachineNodeToNode(scope, clusterUid, vmName, vmNamespace string) error {
	// get cluster
	cluster, err := h.GetCluster(scope, clusterUid)
	if err != nil {
		return err
	}
	if cluster == nil {
		return fmt.Errorf("cluster not found for scope %s and uid %s", scope, clusterUid)
	}

	var params *clusterC.V1SpectroClustersVMMigrateParams
	switch scope {
	case "project":
		params = clusterC.NewV1SpectroClustersVMMigrateParamsWithContext(h.Ctx)
	case "tenant":
		params = clusterC.NewV1SpectroClustersVMMigrateParams()
	default:
		return errors.New("invalid cluster scope specified")
	}
	params = params.WithUID(clusterUid).WithVMName(vmName).WithNamespace(vmNamespace)

	_, err = h.GetClusterClient().V1SpectroClustersVMMigrate(params)
	if err != nil {
		return err
	}
	return nil
}
