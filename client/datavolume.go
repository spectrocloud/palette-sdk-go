package client

import (
	"errors"
	"fmt"

	"github.com/spectrocloud/hapi/models"
	clusterC "github.com/spectrocloud/hapi/spectrocluster/client/v1"
)

func (h *V1Client) CreateDataVolume(uid string, name string, body *models.V1VMAddVolumeEntity) (string, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return "", err
	}

	// get cluster
	cluster, err := h.GetCluster(uid)
	if err != nil {
		return "", err
	}

	// if cluster is nil(deleted or not found), return error
	if cluster == nil {
		return "", fmt.Errorf("cluster not found for uid %s", uid)
	}

	// get cluster scope
	scope := cluster.Metadata.Annotations["scope"]
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

func (h *V1Client) DeleteDataVolume(uid string, namespace string, name string, body *models.V1VMRemoveVolumeEntity) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return err
	}

	// get cluster
	cluster, err := h.GetCluster(uid)
	if err != nil {
		return err
	}
	// get cluster scope
	scope := cluster.Metadata.Annotations["scope"]
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
