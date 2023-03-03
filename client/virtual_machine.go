package client

import (
	"errors"

	"github.com/spectrocloud/hapi/apiutil/transport"
	"github.com/spectrocloud/hapi/models"
	clusterC "github.com/spectrocloud/hapi/spectrocluster/client/v1"
)

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
		break
	case "tenant":
		params = clusterC.NewV1SpectroClustersVMListParams().WithUID(uid)
		break
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

func (h *V1Client) GetVirtualMachine(uid string, name string, namespace string) (*models.V1ClusterVirtualMachine, error) {
	if h.GetVirtualMachineFn != nil {
		return h.GetVirtualMachineFn(uid)
	}
	vm, err := h.GetVirtualMachineWithoutStatus(uid, name, namespace)
	if err != nil {
		return nil, err
	}

	if vm == nil || vm.Status.Created == false { // TODO: check on what is the correct condition to check for deleted
		return nil, nil
	}

	return vm, nil
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

func (h *V1Client) UpdateVirtualMachine(cluster *models.V1SpectroCluster, body *models.V1SpectroClusterVMUpdateEntity) (*models.V1ClusterVirtualMachine, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}

	uid := cluster.Metadata.UID
	// get cluster scope
	scope := cluster.Metadata.Annotations["scope"]
	var params *clusterC.V1SpectroClustersVMUpdateParams

	// switch cluster scope
	switch scope {
	case "project":
		params = clusterC.NewV1SpectroClustersVMUpdateParamsWithContext(h.Ctx)
		break
	case "tenant":
		params = clusterC.NewV1SpectroClustersVMUpdateParams()
		break
	default:
		return nil, errors.New("invalid cluster scope specified")

	}

	params = params.WithUID(uid).WithBody(body)

	// check if vm exists, return error
	exists, err := h.IsVMExists(cluster, body)
	if err != nil {
		return nil, err
	}
	if exists {
		// cannot update vm as another with same name exists
		return nil, errors.New("cannot update vm as another with same name exists")
	}

	vm, err := client.V1SpectroClustersVMUpdate(params)
	if err != nil {
		return nil, err
	}
	return vm.Payload, nil
}

func (h *V1Client) IsVMExists(cluster *models.V1SpectroCluster, newVM *models.V1SpectroClusterVMUpdateEntity) (bool, error) {
	vms, err := h.GetVirtualMachines(cluster)
	if err != nil {
		return false, err
	}
	// return true if vm exists
	for _, vm := range vms {
		if vm.Metadata.Name == newVM.Yaml && vm.Metadata.Namespace == newVM.Yaml { // TODO: deserialize with name and namespace
			return true, nil
		}
	}

	return false, nil
}

func (h *V1Client) CreateVirtualMachine(uid string, body *models.V1SpectroClusterVMCreateEntity) (*models.V1ClusterVirtualMachine, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}

	// get cluster
	cluster, err := h.GetCluster(uid)
	if err != nil {
		return nil, err
	}
	// get cluster scope
	scope := cluster.Metadata.Annotations["scope"]
	var params *clusterC.V1SpectroClustersVMCreateParams
	switch scope {
	case "project":
		params = clusterC.NewV1SpectroClustersVMCreateParamsWithContext(h.Ctx)
		break
	case "tenant":
		params = clusterC.NewV1SpectroClustersVMCreateParams()
		break
	default:
		return nil, errors.New("invalid cluster scope specified")
	}

	params = params.WithUID(uid).WithBody(body)

	vm, err := client.V1SpectroClustersVMCreate(params)
	if err != nil {
		return nil, err
	}
	return vm.Payload, nil
}

func (h *V1Client) DeleteVirtualMachine(uid string, name string, namespace string) error {
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
		break
	case "tenant":
		params = clusterC.NewV1SpectroClustersVMDeleteParams()
		break
	default:
		return errors.New("invalid cluster scope specified")

	}
	params = params.WithUID(uid).WithVMName(name).WithNamespace(namespace)

	_, err = client.V1SpectroClustersVMDelete(params)
	return err
}
