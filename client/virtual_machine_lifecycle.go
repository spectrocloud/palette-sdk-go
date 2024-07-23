package client

import (
	"fmt"

	clientv1 "github.com/spectrocloud/palette-api-go/client/v1"
)

// StartVirtualMachine starts a virtual machine.
func (h *V1Client) StartVirtualMachine(clusterUID, vmName, vmNamespace string) error {
	cluster, err := h.GetCluster(clusterUID)
	if err != nil {
		return err
	}
	if cluster == nil {
		return fmt.Errorf("cluster with uid %s not found", clusterUID)
	}
	params := clientv1.NewV1SpectroClustersVMStartParamsWithContext(h.ctx).
		WithUID(clusterUID).
		WithVMName(vmName).
		WithNamespace(vmNamespace)
	_, err = h.Client.V1SpectroClustersVMStart(params)
	return err
}

// StopVirtualMachine stops a virtual machine.
func (h *V1Client) StopVirtualMachine(clusterUID, vmName, vmNamespace string) error {
	cluster, err := h.GetCluster(clusterUID)
	if err != nil {
		return err
	}
	if cluster == nil {
		return fmt.Errorf("cluster with uid %s not found", clusterUID)
	}
	params := clientv1.NewV1SpectroClustersVMStopParamsWithContext(h.ctx).
		WithUID(clusterUID).
		WithVMName(vmName).
		WithNamespace(vmNamespace)
	_, err = h.Client.V1SpectroClustersVMStop(params)
	return err
}

// PauseVirtualMachine pauses a virtual machine.
func (h *V1Client) PauseVirtualMachine(clusterUID, vmName, vmNamespace string) error {
	cluster, err := h.GetCluster(clusterUID)
	if err != nil {
		return err
	}
	if cluster == nil {
		return fmt.Errorf("cluster with uid %s not found", clusterUID)
	}
	params := clientv1.NewV1SpectroClustersVMPauseParamsWithContext(h.ctx).
		WithUID(clusterUID).
		WithVMName(vmName).
		WithNamespace(vmNamespace)
	_, err = h.Client.V1SpectroClustersVMPause(params)
	return err
}

// ResumeVirtualMachine resumes a virtual machine.
func (h *V1Client) ResumeVirtualMachine(clusterUID, vmName, vmNamespace string) error {
	cluster, err := h.GetCluster(clusterUID)
	if err != nil {
		return err
	}
	if cluster == nil {
		return fmt.Errorf("cluster with uid %s not found", clusterUID)
	}
	params := clientv1.NewV1SpectroClustersVMResumeParamsWithContext(h.ctx).
		WithUID(clusterUID).
		WithVMName(vmName).
		WithNamespace(vmNamespace)
	_, err = h.Client.V1SpectroClustersVMResume(params)
	return err
}

// RestartVirtualMachine restarts a virtual machine.
func (h *V1Client) RestartVirtualMachine(clusterUID, vmName, vmNamespace string) error {
	cluster, err := h.GetCluster(clusterUID)
	if err != nil {
		return err
	}
	if cluster == nil {
		return fmt.Errorf("cluster with uid %s not found", clusterUID)
	}
	params := clientv1.NewV1SpectroClustersVMRestartParamsWithContext(h.ctx).
		WithUID(clusterUID).
		WithVMName(vmName).
		WithNamespace(vmNamespace)
	_, err = h.Client.V1SpectroClustersVMRestart(params)
	return err
}
