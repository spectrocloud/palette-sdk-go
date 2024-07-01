package client

import (
	"fmt"

	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
)

func (h *V1Client) CreateDataVolume(uid, name string, body *models.V1VMAddVolumeEntity) (string, error) {
	cluster, err := h.GetCluster(uid)
	if err != nil {
		return "", err
	} else if cluster == nil {
		return "", fmt.Errorf("cluster with uid %s not found", uid)
	}
	params := clientV1.NewV1SpectroClustersVMAddVolumeParamsWithContext(h.ctx).
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

func (h *V1Client) DeleteDataVolume(uid, namespace, name string, body *models.V1VMRemoveVolumeEntity) error {
	cluster, err := h.GetCluster(uid)
	if err != nil {
		return err
	}
	if cluster == nil {
		return fmt.Errorf("cluster with uid %s not found", uid)
	}
	params := clientV1.NewV1SpectroClustersVMRemoveVolumeParamsWithContext(h.ctx).
		WithUID(uid).
		WithVMName(name).
		WithNamespace(namespace).
		WithBody(body)
	_, err = h.Client.V1SpectroClustersVMRemoveVolume(params)
	return err
}
