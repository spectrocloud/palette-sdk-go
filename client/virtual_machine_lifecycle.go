package client

import (
	"errors"
	"fmt"

	clusterC "github.com/spectrocloud/hapi/spectrocluster/client/v1"
)

func (h *V1Client) StartVirtualMachine(scope, clusterUid, vmName, vmNamespace string) error {
	// get cluster
	cluster, err := h.GetCluster(scope, clusterUid)
	if err != nil {
		return err
	}
	if cluster == nil {
		return fmt.Errorf("cluster not found for scope %s and uid %s", scope, clusterUid)
	}

	// get cluster scope
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

	_, err = h.GetClusterClient().V1SpectroClustersVMStart(params)
	if err != nil {
		return err
	}
	return nil
}

func (h *V1Client) StopVirtualMachine(scope, clusterUid, vmName, vmNamespace string) error {
	// get cluster
	cluster, err := h.GetCluster(scope, clusterUid)
	if err != nil {
		return err
	}
	if cluster == nil {
		return fmt.Errorf("cluster not found for scope %s and uid %s", scope, clusterUid)
	}

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

	_, err = h.GetClusterClient().V1SpectroClustersVMStop(params)
	if err != nil {
		return err
	}
	return nil
}

func (h *V1Client) PauseVirtualMachine(scope, clusterUid, vmName, vmNamespace string) error {
	// get cluster
	cluster, err := h.GetCluster(scope, clusterUid)
	if err != nil {
		return err
	}
	if cluster == nil {
		return fmt.Errorf("cluster not found for scope %s and uid %s", scope, clusterUid)
	}

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

	_, err = h.GetClusterClient().V1SpectroClustersVMPause(params)
	if err != nil {
		return err
	}
	return nil
}

func (h *V1Client) ResumeVirtualMachine(scope, clusterUid, vmName, vmNamespace string) error {
	// get cluster
	cluster, err := h.GetCluster(scope, clusterUid)
	if err != nil {
		return err
	}
	if cluster == nil {
		return fmt.Errorf("cluster not found for scope %s and uid %s", scope, clusterUid)
	}

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

	_, err = h.GetClusterClient().V1SpectroClustersVMResume(params)
	if err != nil {
		return err
	}
	return nil
}

func (h *V1Client) RestartVirtualMachine(scope, clusterUid, vmName, vmNamespace string) error {
	// get cluster
	cluster, err := h.GetCluster(scope, clusterUid)
	if err != nil {
		return err
	}
	if cluster == nil {
		return fmt.Errorf("cluster not found for scope %s and uid %s", scope, clusterUid)
	}

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

	_, err = h.GetClusterClient().V1SpectroClustersVMRestart(params)
	if err != nil {
		return err
	}
	return nil
}
