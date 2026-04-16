package platform

import (
	"errors"

	"github.com/spectrocloud/palette-sdk-go/api/models"
)

func (s *service) GetOrganizationByName(name string) (*models.V1LoginResponse, error) {
	if name == "" {
		return nil, errors.New("organization name is required")
	}
	return s.client.GetOrganizationByName(name)
}

func (s *service) ListOrganizations() ([]*models.V1Organization, error) {
	return s.client.ListOrganizations()
}
