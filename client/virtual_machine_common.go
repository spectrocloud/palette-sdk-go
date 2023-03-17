package client

import (
	"errors"

	"github.com/spectrocloud/hapi/apiutil/transport"
	"github.com/spectrocloud/hapi/models"
	clusterC "github.com/spectrocloud/hapi/spectrocluster/client/v1"
)

func (h *V1Client) IsVMExists(clusterUid string, vmName string, vmNamespace string) (bool, error) {
	vm, err := h.GetVirtualMachine(clusterUid, vmName, vmNamespace)
	if err != nil {
		return false, err
	}
	if vm != nil {
		return true, nil
	}
	return false, nil
}

func (h *V1Client) GetVirtualMachineWithoutStatus(uid string, name string, namespace string) (*models.V1ClusterVirtualMachine, error) {
	if h.GetVirtualMachineWithoutStatusFn != nil {
		return h.GetVirtualMachineWithoutStatusFn(uid)
	}
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}

	params := clusterC.NewV1SpectroClustersVMGetParamsWithContext(h.Ctx).WithUID(uid).WithVMName(name).WithNamespace(namespace)
	success, err := client.V1SpectroClustersVMGet(params)
	// handle tenant context here cluster may be a tenant cluster
	if e, ok := err.(*transport.TransportError); ok && e.HttpCode == 404 {
		params := clusterC.NewV1SpectroClustersVMGetParams().WithUID(uid).WithVMName(name).WithNamespace(namespace)
		success, err = client.V1SpectroClustersVMGet(params)
		if e, ok := err.(*transport.TransportError); ok && e.HttpCode == 404 {
			return nil, nil
		} else if err != nil {
			return nil, err
		}
		//return nil, nil
	}
	if e, ok := err.(*transport.TransportError); ok && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	// special check if the cluster is marked deleted
	vm := success.Payload
	return vm, nil
}

func (h *V1Client) GetVirtualMachines(cluster *models.V1SpectroCluster) ([]*models.V1ClusterVirtualMachine, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}

	uid := cluster.Metadata.UID
	// get cluster scope
	scope := cluster.Metadata.Annotations["scope"]
	var params *clusterC.V1SpectroClustersVMListParams

	// switch cluster scope
	switch scope {
	case "project":
		params = clusterC.NewV1SpectroClustersVMListParamsWithContext(h.Ctx)
	case "tenant":
		params = clusterC.NewV1SpectroClustersVMListParams().WithUID(uid)
	default:
		return nil, errors.New("invalid cluster scope specified")
	}

	params = params.WithUID(uid)

	vms, err := client.V1SpectroClustersVMList(params)
	if err != nil {
		return nil, err
	}
	return vms.Payload.Items, nil
}
