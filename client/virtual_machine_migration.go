package client

import (
	"errors"
	"fmt"

	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
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

	var params *clientV1.V1SpectroClustersVMMigrateParams
	switch scope {
	case "project":
		params = clientV1.NewV1SpectroClustersVMMigrateParamsWithContext(h.Ctx)
	case "tenant":
		params = clientV1.NewV1SpectroClustersVMMigrateParams()
	default:
		return errors.New("invalid cluster scope specified")
	}
	params = params.WithUID(clusterUid).WithVMName(vmName).WithNamespace(vmNamespace)

	_, err = h.Client.V1SpectroClustersVMMigrate(params)
	if err != nil {
		return err
	}
	return nil
}
