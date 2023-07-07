package client

import (
	"errors"
	"fmt"

	"github.com/spectrocloud/hapi/models"
	clusterC "github.com/spectrocloud/hapi/spectrocluster/client/v1"
)

func (h *V1Client) CreateDataVolume(scope, uid, name string, body *models.V1VMAddVolumeEntity) (string, error) {
	if h.CreateDataVolumeFn != nil {
		return h.CreateDataVolumeFn(uid, name, body)
	}
	client, err := h.GetClusterClient()
	if err != nil {
		return "", err
	}

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
	var params *clusterC.V1SpectroClustersVMAddVolumeParams
	switch scope {
	case "project":
		params = clusterC.NewV1SpectroClustersVMAddVolumeParamsWithContext(h.Ctx)
	case "tenant":
		params = clusterC.NewV1SpectroClustersVMAddVolumeParams()
	default:
		return "", errors.New("invalid cluster scope specified")
	}

	params = params.WithUID(uid).WithBody(body).WithVMName(name).WithNamespace(body.DataVolumeTemplate.Metadata.Namespace)

	volume, err := client.V1SpectroClustersVMAddVolume(params)
	if err != nil {
		return "", err
	}
	return volume.AuditUID, nil
}

func (h *V1Client) DeleteDataVolume(scope, uid, namespace, name string, body *models.V1VMRemoveVolumeEntity) error {
	if h.DeleteDataVolumeFn != nil {
		return h.DeleteDataVolumeFn(uid, namespace, name, body)
	}
	client, err := h.GetClusterClient()
	if err != nil {
		return err
	}

	// get cluster
	cluster, err := h.GetCluster(scope, uid)
	if err != nil {
		return err
	}
	if cluster == nil {
		return fmt.Errorf("cluster not found for scope %s and uid %s", scope, uid)
	}

	// get cluster scope
	var params *clusterC.V1SpectroClustersVMRemoveVolumeParams
	switch scope {
	case "project":
		params = clusterC.NewV1SpectroClustersVMRemoveVolumeParamsWithContext(h.Ctx)
	case "tenant":
		params = clusterC.NewV1SpectroClustersVMRemoveVolumeParams()
	default:
		return errors.New("invalid cluster scope specified")
	}
	params = params.WithUID(uid).WithVMName(name).WithNamespace(namespace).WithBody(body)

	_, err = client.V1SpectroClustersVMRemoveVolume(params)
	return err
}
