package client

import (
	"errors"
	"fmt"

	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
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
	var params *clientV1.V1SpectroClustersVMStartParams
	switch scope {
	case "project":
		params = clientV1.NewV1SpectroClustersVMStartParamsWithContext(h.Ctx)
	case "tenant":
		params = clientV1.NewV1SpectroClustersVMStartParams()
	default:
		return errors.New("invalid cluster scope specified")
	}
	params = params.WithUID(clusterUid).WithVMName(vmName).WithNamespace(vmNamespace)

	_, err = h.GetClient().V1SpectroClustersVMStart(params)
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

	var params *clientV1.V1SpectroClustersVMStopParams
	switch scope {
	case "project":
		params = clientV1.NewV1SpectroClustersVMStopParamsWithContext(h.Ctx)
	case "tenant":
		params = clientV1.NewV1SpectroClustersVMStopParams()
	default:
		return errors.New("invalid cluster scope specified")
	}
	params = params.WithUID(clusterUid).WithVMName(vmName).WithNamespace(vmNamespace)

	_, err = h.GetClient().V1SpectroClustersVMStop(params)
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

	var params *clientV1.V1SpectroClustersVMPauseParams
	switch scope {
	case "project":
		params = clientV1.NewV1SpectroClustersVMPauseParamsWithContext(h.Ctx)
	case "tenant":
		params = clientV1.NewV1SpectroClustersVMPauseParams()
	default:
		return errors.New("invalid cluster scope specified")
	}
	params = params.WithUID(clusterUid).WithVMName(vmName).WithNamespace(vmNamespace)

	_, err = h.GetClient().V1SpectroClustersVMPause(params)
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

	var params *clientV1.V1SpectroClustersVMResumeParams
	switch scope {
	case "project":
		params = clientV1.NewV1SpectroClustersVMResumeParamsWithContext(h.Ctx)
	case "tenant":
		params = clientV1.NewV1SpectroClustersVMResumeParams()
	default:
		return errors.New("invalid cluster scope specified")
	}
	params = params.WithUID(clusterUid).WithVMName(vmName).WithNamespace(vmNamespace)

	_, err = h.GetClient().V1SpectroClustersVMResume(params)
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

	var params *clientV1.V1SpectroClustersVMRestartParams
	switch scope {
	case "project":
		params = clientV1.NewV1SpectroClustersVMRestartParamsWithContext(h.Ctx)
	case "tenant":
		params = clientV1.NewV1SpectroClustersVMRestartParams()
	default:
		return errors.New("invalid cluster scope specified")
	}
	params = params.WithUID(clusterUid).WithVMName(vmName).WithNamespace(vmNamespace)

	_, err = h.GetClient().V1SpectroClustersVMRestart(params)
	if err != nil {
		return err
	}
	return nil
}
