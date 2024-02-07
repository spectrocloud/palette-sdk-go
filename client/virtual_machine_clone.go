package client

import (
	"errors"

	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
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

	var params *clientV1.V1SpectroClustersVMCloneParams
	switch scope {
	case "project":
		params = clientV1.NewV1SpectroClustersVMCloneParamsWithContext(h.Ctx)
	case "tenant":
		params = clientV1.NewV1SpectroClustersVMCloneParams()
	default:
		return errors.New("invalid cluster scope specified")
	}
	params = params.WithUID(clusterUid).WithVMName(cloneVMFromName).WithNamespace(vmNamespace).WithBody(body)

	_, err = h.Client.V1SpectroClustersVMClone(params)
	if err != nil {
		return err
	}
	return nil
}
