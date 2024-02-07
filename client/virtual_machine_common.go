package client

import (
	"errors"
	"fmt"

	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
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
	var params *clientV1.V1SpectroClustersVMGetParams
	switch scope {
	case "project":
		params = clientV1.NewV1SpectroClustersVMGetParamsWithContext(h.Ctx).WithUID(uid).WithVMName(name).WithNamespace(namespace)
	case "tenant":
		params = clientV1.NewV1SpectroClustersVMGetParams().WithUID(uid).WithVMName(name).WithNamespace(namespace)
	default:
		return nil, fmt.Errorf("invalid scope %s", scope)
	}

	success, err := h.Client.V1SpectroClustersVMGet(params)
	if err != nil {
		return nil, err
	}

	vm := success.Payload
	return vm, nil
}

func (h *V1Client) GetVirtualMachines(cluster *models.V1SpectroCluster) ([]*models.V1ClusterVirtualMachine, error) {
	var params *clientV1.V1SpectroClustersVMListParams
	switch cluster.Metadata.Annotations["scope"] {
	case "project":
		params = clientV1.NewV1SpectroClustersVMListParamsWithContext(h.Ctx).WithUID(cluster.Metadata.UID)
	case "tenant":
		params = clientV1.NewV1SpectroClustersVMListParams().WithUID(cluster.Metadata.UID)
	default:
		return nil, errors.New("invalid cluster scope specified")
	}

	vms, err := h.Client.V1SpectroClustersVMList(params)
	if err != nil {
		return nil, err
	}
	return vms.Payload.Items, nil
}
