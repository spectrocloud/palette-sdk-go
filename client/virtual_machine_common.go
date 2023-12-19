package client

import (
	"errors"
	"fmt"

	"github.com/spectrocloud/hapi/models"
	clusterC "github.com/spectrocloud/hapi/spectrocluster/client/v1"
)

func (h *V1Client) IsVMExists(scope, clusterUid, vmName, vmNamespace string) (bool, error) {
	vm, err := h.GetVirtualMachine(scope, clusterUid, vmNamespace, vmName)
	if err != nil {
		return false, err
	}
	if vm != nil {
		return true, nil
	}
	return false, nil
}

func (h *V1Client) GetVirtualMachineWithoutStatus(scope, uid, name, namespace string) (*models.V1ClusterVirtualMachine, error) {
	if h.GetVirtualMachineWithoutStatusFn != nil {
		return h.GetVirtualMachineWithoutStatusFn(uid)
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

	success, err := h.GetClusterClient().V1SpectroClustersVMGet(params)
	if err != nil {
		return nil, err
	}

	vm := success.Payload
	return vm, nil
}

func (h *V1Client) GetVirtualMachines(cluster *models.V1SpectroCluster) ([]*models.V1ClusterVirtualMachine, error) {
	var params *clusterC.V1SpectroClustersVMListParams
	switch cluster.Metadata.Annotations["scope"] {
	case "project":
		params = clusterC.NewV1SpectroClustersVMListParamsWithContext(h.Ctx).WithUID(cluster.Metadata.UID)
	case "tenant":
		params = clusterC.NewV1SpectroClustersVMListParams().WithUID(cluster.Metadata.UID)
	default:
		return nil, errors.New("invalid cluster scope specified")
	}

	vms, err := h.GetClusterClient().V1SpectroClustersVMList(params)
	if err != nil {
		return nil, err
	}
	return vms.Payload.Items, nil
}
