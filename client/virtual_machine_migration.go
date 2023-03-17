package client

import (
	"errors"

	clusterC "github.com/spectrocloud/hapi/spectrocluster/client/v1"
)

func (h *V1Client) MigrateVirtualMachineNodeToNode(clusterUid string, vmName string, vmNamespace string) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return err
	}

	// get cluster
	cluster, err := h.GetCluster(clusterUid)
	if err != nil {
		return err
	}
	// get cluster scope
	scope := cluster.Metadata.Annotations["scope"]
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

	_, err = client.V1SpectroClustersVMMigrate(params)
	if err != nil {
		return err
	}
	return nil
}
