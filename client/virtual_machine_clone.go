package client

import (
	"errors"

	"github.com/spectrocloud/hapi/models"
	clusterC "github.com/spectrocloud/hapi/spectrocluster/client/v1"
)

func (h *V1Client) CloneVirtualMachine(scope, clusterUid, cloneVMFromName, vmName, vmNamespace string) error {
	cluster, err := h.GetCluster(scope, clusterUid)
	if err != nil {
		return err
	}
	if cluster == nil {
		return errors.New("cluster not found")
	}

	body := &models.V1SpectroClusterVMCloneEntity{
		CloneName: &vmName,
	}

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

	_, err = h.GetClusterClient().V1SpectroClustersVMClone(params)
	if err != nil {
		return err
	}
	return nil
}
