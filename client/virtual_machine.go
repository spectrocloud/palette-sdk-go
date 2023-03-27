package client

import (
	"errors"
	"fmt"

	"github.com/spectrocloud/hapi/models"
	clusterC "github.com/spectrocloud/hapi/spectrocluster/client/v1"
)

func (h *V1Client) CreateVirtualMachine(uid string, body *models.V1ClusterVirtualMachine) (*models.V1ClusterVirtualMachine, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}

	// get cluster
	cluster, err := h.GetCluster(uid)
	if err != nil {
		return nil, err
	}

	// if cluster is nil(deleted or not found), return error
	if cluster == nil {
		return nil, errors.New(fmt.Sprintf("cluster not found for uid %s", uid))
	}

	// get cluster scope
	scope := cluster.Metadata.Annotations["scope"]
	var params *clusterC.V1SpectroClustersVMCreateParams
	switch scope {
	case "project":
		params = clusterC.NewV1SpectroClustersVMCreateParamsWithContext(h.Ctx)
	case "tenant":
		params = clusterC.NewV1SpectroClustersVMCreateParams()
	default:
		return nil, errors.New("invalid cluster scope specified")
	}

	params = params.WithUID(uid).WithBody(body).WithNamespace(params.Body.Metadata.Namespace)

	vm, err := client.V1SpectroClustersVMCreate(params)
	if err != nil {
		return nil, err
	}
	return vm.Payload, nil
}

func (h *V1Client) GetVirtualMachine(uid string, namespace string, name string) (*models.V1ClusterVirtualMachine, error) {
	if h.GetVirtualMachineFn != nil {
		return h.GetVirtualMachineFn(uid)
	}
	vm, err := h.GetVirtualMachineWithoutStatus(uid, name, namespace)
	if err != nil {
		return nil, err
	}
	return vm, nil
}

func (h *V1Client) UpdateVirtualMachine(cluster *models.V1SpectroCluster, vmName string, body *models.V1ClusterVirtualMachine) (*models.V1ClusterVirtualMachine, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}

	clusterUid := cluster.Metadata.UID
	// get cluster scope
	scope := cluster.Metadata.Annotations["scope"]
	var params *clusterC.V1SpectroClustersVMUpdateParams

	// switch cluster scope
	switch scope {
	case "project":
		params = clusterC.NewV1SpectroClustersVMUpdateParamsWithContext(h.Ctx)
	case "tenant":
		params = clusterC.NewV1SpectroClustersVMUpdateParams()
	default:
		return nil, errors.New("invalid cluster scope specified")
	}

	params = params.WithUID(clusterUid).WithBody(body).WithNamespace(body.Metadata.Namespace).WithVMName(vmName)

	// check if vm exists, return error
	exists, err := h.IsVMExists(clusterUid, vmName, body.Metadata.Namespace)
	if err != nil {
		return nil, err
	}
	if !exists {
		// cannot update vm as another with same name exists
		return nil, errors.New("VM not exists")
	}

	vm, err := client.V1SpectroClustersVMUpdate(params)
	if err != nil {
		return nil, err
	}
	return vm.Payload, nil

	//return nil, nil
}

func (h *V1Client) DeleteVirtualMachine(uid string, namespace string, name string) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return err
	}

	// get cluster
	cluster, err := h.GetCluster(uid)
	if err != nil {
		return err
	}
	// get cluster scope
	scope := cluster.Metadata.Annotations["scope"]
	var params *clusterC.V1SpectroClustersVMDeleteParams
	switch scope {
	case "project":
		params = clusterC.NewV1SpectroClustersVMDeleteParamsWithContext(h.Ctx)
	case "tenant":
		params = clusterC.NewV1SpectroClustersVMDeleteParams()
	default:
		return errors.New("invalid cluster scope specified")
	}
	params = params.WithUID(uid).WithVMName(name).WithNamespace(namespace)

	_, err = client.V1SpectroClustersVMDelete(params)
	return err
}
