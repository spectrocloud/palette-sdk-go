package client

import (
	"errors"
	"fmt"

	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
)

func (h *V1Client) CreateVirtualMachine(uid string, body *models.V1ClusterVirtualMachine) (*models.V1ClusterVirtualMachine, error) {
	cluster, err := h.GetCluster(uid)
	if err != nil {
		return nil, err
	}
	if cluster == nil {
		return nil, fmt.Errorf("cluster with uid %s not found", uid)
	}
	params := clientV1.NewV1SpectroClustersVMCreateParamsWithContext(h.ctx).
		WithUID(uid).
		WithBody(body).
		WithNamespace(body.Metadata.Namespace)
	resp, err := h.Client.V1SpectroClustersVMCreate(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

func (h *V1Client) GetVirtualMachine(uid, namespace, name string) (*models.V1ClusterVirtualMachine, error) {
	vm, err := h.GetVirtualMachineWithoutStatus(uid, name, namespace)
	if err != nil {
		return nil, err
	}
	return vm, nil
}

func (h *V1Client) UpdateVirtualMachine(cluster *models.V1SpectroCluster, vmName string, body *models.V1ClusterVirtualMachine) (*models.V1ClusterVirtualMachine, error) {
	clusterUID := cluster.Metadata.UID
	params := clientV1.NewV1SpectroClustersVMUpdateParams().
		WithContext(ContextForScope(cluster.Metadata.Annotations[Scope], h.projectUID)).
		WithUID(clusterUID).
		WithBody(body).
		WithNamespace(body.Metadata.Namespace).
		WithVMName(vmName)

	// check if vm exists
	exists, err := h.IsVMExists(clusterUID, vmName, body.Metadata.Namespace)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.New("VM does not exist")
	}

	resp, err := h.Client.V1SpectroClustersVMUpdate(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

func (h *V1Client) DeleteVirtualMachine(uid, namespace, name string) error {
	cluster, err := h.GetCluster(uid)
	if err != nil {
		return err
	}
	if cluster == nil {
		return fmt.Errorf("cluster with uid %s not found", uid)
	}
	params := clientV1.NewV1SpectroClustersVMDeleteParamsWithContext(h.ctx).
		WithUID(uid).
		WithVMName(name).
		WithNamespace(namespace)
	_, err = h.Client.V1SpectroClustersVMDelete(params)
	return err
}
