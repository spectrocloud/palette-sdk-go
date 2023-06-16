package client

import (
	"errors"

	"github.com/spectrocloud/hapi/models"
	clusterC "github.com/spectrocloud/hapi/spectrocluster/client/v1"
)

func (h *V1Client) CloneVirtualMachine(clusterUid string, cloneVMFromName string, vmName string, vmNamespace string) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return err
	}

	// get cluster
	cluster, err := h.GetCluster(client, clusterUid)

	body := &models.V1SpectroClusterVMCloneEntity{
		CloneName: &vmName,
	}
	if err != nil {
		return err
	}
	// get cluster scope
	scope := cluster.Metadata.Annotations["scope"]
	var params *clusterC.V1SpectroClustersVMCloneParams
	switch scope {
	case "project":
		params = clusterC.NewV1SpectroClustersVMCloneParamsWithContext(h.Ctx)
	case "tenant":
		params = clusterC.NewV1SpectroClustersVMCloneParams()
	default:
		return errors.New("invalid cluster scope specified")
	}
	params = params.WithUID(clusterUid).WithVMName(cloneVMFromName).WithNamespace(vmNamespace).WithBody(body)

	_, err = client.V1SpectroClustersVMClone(params)
	if err != nil {
		return err
	}
	return nil
}
