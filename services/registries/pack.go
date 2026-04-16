package registries

import (
	"errors"

	"github.com/spectrocloud/palette-sdk-go/api/models"
)

func (s *service) SearchPackRegistries() ([]*models.V1RegistryMetadata, error) {
	return s.client.SearchPackRegistryCommon()
}

func (s *service) GetPackRegistryByName(name string) (*models.V1PackRegistry, error) {
	if name == "" {
		return nil, errors.New("registry name is required")
	}
	return s.client.GetPackRegistryByName(name)
}

func (s *service) ListPackRegistries() ([]*models.V1PackRegistrySummary, error) {
	return s.client.ListPackRegistries()
}

func (s *service) GetPackRegistrySyncStatus(uid string) (*models.V1RegistrySyncStatus, error) {
	if uid == "" {
		return nil, errUIDRequired
	}
	return s.client.GetPackRegistrySyncStatus(uid)
}
