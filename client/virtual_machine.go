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
		break
	case "tenant":
		params = clusterC.NewV1SpectroClustersVMUpdateParams()
		break
	default:
		return nil, errors.New("invalid cluster scope specified")

	}

	params = params.WithUID(clusterUid).WithBody(body).WithNamespace(body.Metadata.Namespace).WithVMName(vmName)

	// check if vm exists, return error
	exists, err := h.IsVMExists(clusterUid, vmName, body.Metadata.Namespace)
	if err != nil {
		return nil, err
	}
	if exists == false {
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

	params = params.WithUID(uid).WithBody(body).WithNamespace(params.Body.Metadata.Namespace)

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

func (h *V1Client) RestartVirtualMachine(clusterUid string, vmName string, vmNamespace string) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return err
	}

	// get cluster
	cluster, err := h.GetCluster(clusterUid)
	if err != nil {
		return err
	}
	// get cluster scope
	scope := cluster.Metadata.Annotations["scope"]
	var params *clusterC.V1SpectroClustersVMRestartParams
	switch scope {
	case "project":
		params = clusterC.NewV1SpectroClustersVMRestartParamsWithContext(h.Ctx)
		break
	case "tenant":
		params = clusterC.NewV1SpectroClustersVMRestartParams()
		break
	default:
		return errors.New("invalid cluster scope specified")
	}
	params = params.WithUID(clusterUid).WithVMName(vmName).WithNamespace(vmNamespace)

	_, err = client.V1SpectroClustersVMRestart(params)
	if err != nil {
		return err
	}
	return nil
}

func (h *V1Client) PauseVirtualMachine(clusterUid string, vmName string, vmNamespace string) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return err
	}

	// get cluster
	cluster, err := h.GetCluster(clusterUid)
	if err != nil {
		return err
	}
	// get cluster scope
	scope := cluster.Metadata.Annotations["scope"]
	var params *clusterC.V1SpectroClustersVMPauseParams
	switch scope {
	case "project":
		params = clusterC.NewV1SpectroClustersVMPauseParamsWithContext(h.Ctx)
		break
	case "tenant":
		params = clusterC.NewV1SpectroClustersVMPauseParams()
		break
	default:
		return errors.New("invalid cluster scope specified")
	}
	params = params.WithUID(clusterUid).WithVMName(vmName).WithNamespace(vmNamespace)

	_, err = client.V1SpectroClustersVMPause(params)
	if err != nil {
		return err
	}
	return nil
}

func (h *V1Client) StopVirtualMachine(clusterUid string, vmName string, vmNamespace string) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return err
	}

	// get cluster
	cluster, err := h.GetCluster(clusterUid)
	if err != nil {
		return err
	}
	// get cluster scope
	scope := cluster.Metadata.Annotations["scope"]
	var params *clusterC.V1SpectroClustersVMStopParams
	switch scope {
	case "project":
		params = clusterC.NewV1SpectroClustersVMStopParamsWithContext(h.Ctx)
		break
	case "tenant":
		params = clusterC.NewV1SpectroClustersVMStopParams()
		break
	default:
		return errors.New("invalid cluster scope specified")
	}
	params = params.WithUID(clusterUid).WithVMName(vmName).WithNamespace(vmNamespace)

	_, err = client.V1SpectroClustersVMStop(params)
	if err != nil {
		return err
	}
	return nil
}

func (h *V1Client) StartVirtualMachine(clusterUid string, vmName string, vmNamespace string) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return err
	}

	// get cluster
	cluster, err := h.GetCluster(clusterUid)
	if err != nil {
		return err
	}
	// get cluster scope
	scope := cluster.Metadata.Annotations["scope"]
	var params *clusterC.V1SpectroClustersVMStartParams
	switch scope {
	case "project":
		params = clusterC.NewV1SpectroClustersVMStartParamsWithContext(h.Ctx)
		break
	case "tenant":
		params = clusterC.NewV1SpectroClustersVMStartParams()
		break
	default:
		return errors.New("invalid cluster scope specified")
	}
	params = params.WithUID(clusterUid).WithVMName(vmName).WithNamespace(vmNamespace)

	_, err = client.V1SpectroClustersVMStart(params)
	if err != nil {
		return err
	}
	return nil
}

func (h *V1Client) MigrateVirtualMachineNodeToNode(clusterUid string, vmName string, vmNamespace string) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return err
	}

	// get cluster
	cluster, err := h.GetCluster(clusterUid)
	if err != nil {
		return err
	}
	// get cluster scope
	scope := cluster.Metadata.Annotations["scope"]
	var params *clusterC.V1SpectroClustersVMMigrateParams
	switch scope {
	case "project":
		params = clusterC.NewV1SpectroClustersVMMigrateParamsWithContext(h.Ctx)
		break
	case "tenant":
		params = clusterC.NewV1SpectroClustersVMMigrateParams()
		break
	default:
		return errors.New("invalid cluster scope specified")
	}
	params = params.WithUID(clusterUid).WithVMName(vmName).WithNamespace(vmNamespace)

	_, err = client.V1SpectroClustersVMMigrate(params)
	if err != nil {
		return err
	}
	return nil
}

func (h *V1Client) ResumeVirtualMachine(clusterUid string, vmName string, vmNamespace string) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return err
	}

	// get cluster
	cluster, err := h.GetCluster(clusterUid)
	if err != nil {
		return err
	}
	// get cluster scope
	scope := cluster.Metadata.Annotations["scope"]
	var params *clusterC.V1SpectroClustersVMResumeParams
	switch scope {
	case "project":
		params = clusterC.NewV1SpectroClustersVMResumeParamsWithContext(h.Ctx)
		break
	case "tenant":
		params = clusterC.NewV1SpectroClustersVMResumeParams()
		break
	default:
		return errors.New("invalid cluster scope specified")
	}
	params = params.WithUID(clusterUid).WithVMName(vmName).WithNamespace(vmNamespace)

	_, err = client.V1SpectroClustersVMResume(params)
	if err != nil {
		return err
	}
	return nil
}

func (h *V1Client) CloneVirtualMachine(clusterUid string, cloneVMFromName string, vmName string, vmNamespace string) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return err
	}

	// get cluster
	cluster, err := h.GetCluster(clusterUid)

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
		break
	case "tenant":
		params = clusterC.NewV1SpectroClustersVMCloneParams()
		break
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
