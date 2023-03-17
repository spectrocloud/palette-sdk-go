package client

import (
	"errors"

	clusterC "github.com/spectrocloud/hapi/spectrocluster/client/v1"
)

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
	case "tenant":
		params = clusterC.NewV1SpectroClustersVMStartParams()
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
	case "tenant":
		params = clusterC.NewV1SpectroClustersVMStopParams()
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
	case "tenant":
		params = clusterC.NewV1SpectroClustersVMPauseParams()
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
	case "tenant":
		params = clusterC.NewV1SpectroClustersVMResumeParams()
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
	case "tenant":
		params = clusterC.NewV1SpectroClustersVMRestartParams()
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
