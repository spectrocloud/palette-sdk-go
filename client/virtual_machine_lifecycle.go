package client

import (
	"fmt"

	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
)

func (h *V1Client) StartVirtualMachine(clusterUid, vmName, vmNamespace string) error {
	cluster, err := h.GetCluster(clusterUid)
	if err != nil {
		return err
	}
	if cluster == nil {
		return fmt.Errorf("cluster with uid %s not found", clusterUid)
	}
	params := clientV1.NewV1SpectroClustersVMStartParamsWithContext(h.ctx).
		WithUID(clusterUid).
		WithVMName(vmName).
		WithNamespace(vmNamespace)
	_, err = h.Client.V1SpectroClustersVMStart(params)
	return err
}

func (h *V1Client) StopVirtualMachine(clusterUid, vmName, vmNamespace string) error {
	cluster, err := h.GetCluster(clusterUid)
	if err != nil {
		return err
	}
	if cluster == nil {
		return fmt.Errorf("cluster with uid %s not found", clusterUid)
	}
	params := clientV1.NewV1SpectroClustersVMStopParamsWithContext(h.ctx).
		WithUID(clusterUid).
		WithVMName(vmName).
		WithNamespace(vmNamespace)
	_, err = h.Client.V1SpectroClustersVMStop(params)
	return err
}

func (h *V1Client) PauseVirtualMachine(clusterUid, vmName, vmNamespace string) error {
	cluster, err := h.GetCluster(clusterUid)
	if err != nil {
		return err
	}
	if cluster == nil {
		return fmt.Errorf("cluster with uid %s not found", clusterUid)
	}
	params := clientV1.NewV1SpectroClustersVMPauseParamsWithContext(h.ctx).
		WithUID(clusterUid).
		WithVMName(vmName).
		WithNamespace(vmNamespace)
	_, err = h.Client.V1SpectroClustersVMPause(params)
	return err
}

func (h *V1Client) ResumeVirtualMachine(clusterUid, vmName, vmNamespace string) error {
	cluster, err := h.GetCluster(clusterUid)
	if err != nil {
		return err
	}
	if cluster == nil {
		return fmt.Errorf("cluster with uid %s not found", clusterUid)
	}
	params := clientV1.NewV1SpectroClustersVMResumeParamsWithContext(h.ctx).
		WithUID(clusterUid).
		WithVMName(vmName).
		WithNamespace(vmNamespace)
	_, err = h.Client.V1SpectroClustersVMResume(params)
	return err
}

func (h *V1Client) RestartVirtualMachine(clusterUid, vmName, vmNamespace string) error {
	cluster, err := h.GetCluster(clusterUid)
	if err != nil {
		return err
	}
	if cluster == nil {
		return fmt.Errorf("cluster with uid %s not found", clusterUid)
	}
	params := clientV1.NewV1SpectroClustersVMRestartParamsWithContext(h.ctx).
		WithUID(clusterUid).
		WithVMName(vmName).
		WithNamespace(vmNamespace)
	_, err = h.Client.V1SpectroClustersVMRestart(params)
	return err
}
