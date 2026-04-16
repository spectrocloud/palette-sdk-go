package registries

import (
	"errors"

	"github.com/spectrocloud/palette-sdk-go/api/models"
)

func (s *service) ListHelmRegistries(scope string) ([]*models.V1HelmRegistrySummary, error) {
	return s.client.ListHelmRegistries(scope)
}

func (s *service) GetHelmRegistry(uid string) (*models.V1HelmRegistry, error) {
	if uid == "" {
		return nil, errUIDRequired
	}
	return s.client.GetHelmRegistry(uid)
}

func (s *service) GetHelmRegistryByName(name string) (*models.V1HelmRegistry, error) {
	if name == "" {
		return nil, errors.New("registry name is required")
	}
	return s.client.GetHelmRegistryByName(name)
}

func (s *service) CreateHelmRegistry(registry *models.V1HelmRegistryEntity) (string, error) {
	if registry == nil {
		return "", errors.New("helm registry entity is required")
	}
	return s.client.CreateHelmRegistry(registry)
}

func (s *service) UpdateHelmRegistry(uid string, registry *models.V1HelmRegistry) error {
	if uid == "" {
		return errUIDRequired
	}
	return s.client.UpdateHelmRegistry(uid, registry)
}

func (s *service) DeleteHelmRegistry(uid string) error {
	if uid == "" {
		return errUIDRequired
	}
	return s.client.DeleteHelmRegistry(uid)
}

func (s *service) GetHelmRegistrySyncStatus(uid string) (*models.V1RegistrySyncStatus, error) {
	if uid == "" {
		return nil, errUIDRequired
	}
	return s.client.GetHelmRegistrySyncStatus(uid)
}
