package registries

import (
	"errors"

	"github.com/spectrocloud/palette-sdk-go/api/models"
)

func (s *service) ListOCIRegistries() ([]*models.V1OciRegistry, error) {
	return s.client.ListOCIRegistries()
}

func (s *service) GetOciBasicRegistry(uid string) (*models.V1BasicOciRegistry, error) {
	if uid == "" {
		return nil, errUIDRequired
	}
	return s.client.GetOciBasicRegistry(uid)
}

func (s *service) CreateOciBasicRegistry(registry *models.V1BasicOciRegistry) (string, error) {
	if registry == nil {
		return "", errors.New("OCI basic registry is required")
	}
	return s.client.CreateOciBasicRegistry(registry)
}

func (s *service) UpdateOciBasicRegistry(uid string, registry *models.V1BasicOciRegistry) error {
	if uid == "" {
		return errUIDRequired
	}
	return s.client.UpdateOciBasicRegistry(uid, registry)
}

func (s *service) DeleteOciBasicRegistry(uid string) error {
	if uid == "" {
		return errUIDRequired
	}
	return s.client.DeleteOciBasicRegistry(uid)
}

func (s *service) ValidateOciBasicRegistry(body *models.V1BasicOciRegistrySpec) error {
	if body == nil {
		return errors.New("OCI basic registry spec is required")
	}
	return s.client.ValidateOciBasicRegistry(body)
}

func (s *service) GetOciEcrRegistry(uid string) (*models.V1EcrRegistry, error) {
	if uid == "" {
		return nil, errUIDRequired
	}
	return s.client.GetOciEcrRegistry(uid)
}

func (s *service) CreateOciEcrRegistry(registry *models.V1EcrRegistry) (string, error) {
	if registry == nil {
		return "", errors.New("ECR registry is required")
	}
	return s.client.CreateOciEcrRegistry(registry)
}

func (s *service) UpdateOciEcrRegistry(uid string, registry *models.V1EcrRegistry) error {
	if uid == "" {
		return errUIDRequired
	}
	return s.client.UpdateOciEcrRegistry(uid, registry)
}

func (s *service) DeleteOciEcrRegistry(uid string) error {
	if uid == "" {
		return errUIDRequired
	}
	return s.client.DeleteOciEcrRegistry(uid)
}

func (s *service) ValidateOciEcrRegistry(body *models.V1EcrRegistrySpec) error {
	if body == nil {
		return errors.New("ECR registry spec is required")
	}
	return s.client.ValidateOciEcrRegistry(body)
}
