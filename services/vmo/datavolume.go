package vmo

import (
	"errors"

	"github.com/spectrocloud/palette-sdk-go/api/models"
)

func (s *service) CreateDataVolume(clusterUID, name string, body *models.V1VMAddVolumeEntity) (string, error) {
	if clusterUID == "" {
		return "", errClusterUIDRequired
	}
	if name == "" {
		return "", errors.New("data volume name is required")
	}
	return s.client.CreateDataVolume(clusterUID, name, body)
}

func (s *service) DeleteDataVolume(clusterUID, namespace, name string, body *models.V1VMRemoveVolumeEntity) error {
	if clusterUID == "" {
		return errClusterUIDRequired
	}
	if name == "" {
		return errors.New("data volume name is required")
	}
	return s.client.DeleteDataVolume(clusterUID, namespace, name, body)
}
