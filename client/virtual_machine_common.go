package client

import (
	"errors"
	"fmt"

	"github.com/spectrocloud/hapi/models"
	clusterC "github.com/spectrocloud/hapi/spectrocluster/client/v1"
)

func (h *V1Client) IsVMExists(scope string, clusterUid string, vmName string, vmNamespace string) (bool, error) {
	vm, err := h.GetVirtualMachine(scope, clusterUid, vmNamespace, vmName)
	if err != nil {
		return false, err
	}
	if vm != nil {
		return true, nil
	}
	return false, nil
}

func (h *V1Client) GetVirtualMachineWithoutStatus(scope string, uid string, name string, namespace string) (*models.V1ClusterVirtualMachine, error) {
	if h.GetVirtualMachineWithoutStatusFn != nil {
		return h.GetVirtualMachineWithoutStatusFn(uid)
	}
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}

	var params *clusterC.V1SpectroClustersVMGetParams
	switch scope {
	case "project":
		params = clusterC.NewV1SpectroClustersVMGetParamsWithContext(h.Ctx).WithUID(uid).WithVMName(name).WithNamespace(namespace)
	case "tenant":
		params = clusterC.NewV1SpectroClustersVMGetParams().WithUID(uid).WithVMName(name).WithNamespace(namespace)
	default:
		return nil, fmt.Errorf("invalid scope %s", scope)
	}

	success, err := client.V1SpectroClustersVMGet(params)
	if err != nil {
		return nil, err
	}

	vm := success.Payload
	return vm, nil
}

func (h *V1Client) GetVirtualMachines(cluster *models.V1SpectroCluster) ([]*models.V1ClusterVirtualMachine, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}

	uid := cluster.Metadata.UID
	var params *clusterC.V1SpectroClustersVMListParams

	// switch cluster scope
	switch cluster.Metadata.Annotations["scope"] {
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
