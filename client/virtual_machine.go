package client

import (
	"errors"
	"fmt"

	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
)

func (h *V1Client) CreateVirtualMachine(scope, uid string, body *models.V1ClusterVirtualMachine) (*models.V1ClusterVirtualMachine, error) {
	// get cluster
	cluster, err := h.GetCluster(scope, uid)
	if err != nil {
		return nil, err
	}

	// if cluster is nil(deleted or not found), return error
	if cluster == nil {
		return nil, fmt.Errorf("cluster not found for uid %s", uid)
	}

	// get cluster scope
	var params *clientV1.V1SpectroClustersVMCreateParams
	switch scope {
	case "project":
		params = clientV1.NewV1SpectroClustersVMCreateParamsWithContext(h.Ctx)
	case "tenant":
		params = clientV1.NewV1SpectroClustersVMCreateParams()
	default:
		return nil, errors.New("invalid cluster scope specified")
	}

	params = params.WithUID(uid).WithBody(body).WithNamespace(params.Body.Metadata.Namespace)

	vm, err := h.GetClient().V1SpectroClustersVMCreate(params)
	if err != nil {
		return nil, err
	}
	return vm.Payload, nil
}

func (h *V1Client) GetVirtualMachine(scope, uid, namespace, name string) (*models.V1ClusterVirtualMachine, error) {
	vm, err := h.GetVirtualMachineWithoutStatus(scope, uid, name, namespace)
	if err != nil {
		return nil, err
	}
	return vm, nil
}

func (h *V1Client) UpdateVirtualMachine(cluster *models.V1SpectroCluster, vmName string, body *models.V1ClusterVirtualMachine) (*models.V1ClusterVirtualMachine, error) {
	clusterUid := cluster.Metadata.UID
	scope := cluster.Metadata.Annotations["scope"]
	// get cluster scope
	var params *clientV1.V1SpectroClustersVMUpdateParams

	// switch cluster scope
	switch scope {
	case "project":
		params = clientV1.NewV1SpectroClustersVMUpdateParamsWithContext(h.Ctx)
	case "tenant":
		params = clientV1.NewV1SpectroClustersVMUpdateParams()
	default:
		return nil, errors.New("invalid cluster scope specified")
	}

	params = params.WithUID(clusterUid).WithBody(body).WithNamespace(body.Metadata.Namespace).WithVMName(vmName)

	// check if vm exists, return error
	exists, err := h.IsVMExists(scope, clusterUid, vmName, body.Metadata.Namespace)
	if err != nil {
		return nil, err
	}
	if !exists {
		// cannot update vm as another with same name exists
		return nil, errors.New("VM not exists")
	}

	vm, err := h.GetClient().V1SpectroClustersVMUpdate(params)
	if err != nil {
		return nil, err
	}
	return vm.Payload, nil

	//return nil, nil
}

func (h *V1Client) DeleteVirtualMachine(scope, uid, namespace, name string) error {
	// get cluster
	cluster, err := h.GetCluster(scope, uid)
	if err != nil {
		return err
	}
	if cluster == nil {
		return fmt.Errorf("cluster not found for scope %s and uid %s", scope, uid)
	}

	// get cluster scope
	var params *clientV1.V1SpectroClustersVMDeleteParams
	switch scope {
	case "project":
		params = clientV1.NewV1SpectroClustersVMDeleteParamsWithContext(h.Ctx)
	case "tenant":
		params = clientV1.NewV1SpectroClustersVMDeleteParams()
	default:
		return errors.New("invalid cluster scope specified")
	}
	params = params.WithUID(uid).WithVMName(name).WithNamespace(namespace)

	_, err = h.GetClient().V1SpectroClustersVMDelete(params)
	return err
}
