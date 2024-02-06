package client

import (
	"errors"
	"fmt"

	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
)

func (h *V1Client) CreateDataVolume(scope, uid, name string, body *models.V1VMAddVolumeEntity) (string, error) {
	// get cluster
	cluster, err := h.GetCluster(scope, uid)
	if err != nil {
		return "", err
	}

	// if cluster is nil(deleted or not found), return error
	if cluster == nil {
		return "", fmt.Errorf("cluster not found for uid %s", uid)
	}

	// get cluster scope
	var params *clientV1.V1SpectroClustersVMAddVolumeParams
	switch scope {
	case "project":
		params = clientV1.NewV1SpectroClustersVMAddVolumeParamsWithContext(h.Ctx)
	case "tenant":
		params = clientV1.NewV1SpectroClustersVMAddVolumeParams()
	default:
		return "", errors.New("invalid cluster scope specified")
	}

	params = params.WithUID(uid).WithBody(body).WithVMName(name).WithNamespace(body.DataVolumeTemplate.Metadata.Namespace)

	volume, err := h.GetClient().V1SpectroClustersVMAddVolume(params)
	if err != nil {
		return "", err
	}
	return volume.AuditUID, nil
}

func (h *V1Client) DeleteDataVolume(scope, uid, namespace, name string, body *models.V1VMRemoveVolumeEntity) error {
	// get cluster
	cluster, err := h.GetCluster(scope, uid)
	if err != nil {
		return err
	}
	if cluster == nil {
		return fmt.Errorf("cluster not found for scope %s and uid %s", scope, uid)
	}

	// get cluster scope
	var params *clientV1.V1SpectroClustersVMRemoveVolumeParams
	switch scope {
	case "project":
		params = clientV1.NewV1SpectroClustersVMRemoveVolumeParamsWithContext(h.Ctx)
	case "tenant":
		params = clientV1.NewV1SpectroClustersVMRemoveVolumeParams()
	default:
		return errors.New("invalid cluster scope specified")
	}
	params = params.WithUID(uid).WithVMName(name).WithNamespace(namespace).WithBody(body)

	_, err = h.GetClient().V1SpectroClustersVMRemoveVolume(params)
	return err
}
