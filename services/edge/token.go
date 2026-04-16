package edge

import (
	"github.com/spectrocloud/palette-sdk-go/api/models"
)

func (s *service) GetTokenValue(name string) (string, error) {
	if name == "" {
		return "", errNameRequired
	}
	return s.client.GetRegistrationToken(name)
}

func (s *service) GetTokenByUID(uid string) (*models.V1EdgeToken, error) {
	if uid == "" {
		return nil, errUIDRequired
	}
	return s.client.GetRegistrationTokenByUID(uid)
}

func (s *service) GetTokenByName(name string) (*models.V1EdgeToken, error) {
	if name == "" {
		return nil, errNameRequired
	}
	return s.client.GetRegistrationTokenByName(name)
}

func (s *service) CreateToken(name string, body *models.V1EdgeTokenEntity) (string, string, error) {
	if name == "" {
		return "", "", errNameRequired
	}
	return s.client.CreateRegistrationToken(name, body)
}

func (s *service) DeleteToken(uid string) error {
	if uid == "" {
		return errUIDRequired
	}
	return s.client.DeleteRegistrationToken(uid)
}
