package client

import (
	"fmt"

	clientv1 "github.com/spectrocloud/palette-sdk-go/api/client/version1"
	"github.com/spectrocloud/palette-sdk-go/api/models"
)

// CreateDataVolume creates a new data volume.
func (h *V1Client) CreateDataVolume(uid, name string, body *models.V1VMAddVolumeEntity) (string, error) {
	cluster, err := h.GetCluster(uid)
	if err != nil {
		return "", err
	} else if cluster == nil {
		return "", fmt.Errorf("cluster with uid %s not found", uid)
	}
	params := clientv1.NewV1SpectroClustersVMAddVolumeParamsWithContext(h.ctx).
		WithUID(uid).
		WithBody(body).
		WithVMName(name).
		WithNamespace(body.DataVolumeTemplate.Metadata.Namespace)
	resp, err := h.Client.V1SpectroClustersVMAddVolume(params)
	if err != nil {
		return "", err
	}
	return resp.AuditUID, nil
}

// DeleteDataVolume deletes an existing data volume.
func (h *V1Client) DeleteDataVolume(uid, namespace, name string, body *models.V1VMRemoveVolumeEntity) error {
	cluster, err := h.GetCluster(uid)
	if err != nil {
		return err
	}
	if cluster == nil {
		return fmt.Errorf("cluster with uid %s not found", uid)
	}
	params := clientv1.NewV1SpectroClustersVMRemoveVolumeParamsWithContext(h.ctx).
		WithUID(uid).
		WithVMName(name).
		WithNamespace(namespace).
		WithBody(body)
	_, err = h.Client.V1SpectroClustersVMRemoveVolume(params)
	return err
}
