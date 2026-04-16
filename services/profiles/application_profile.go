package profiles

import (
	"errors"

	"github.com/spectrocloud/palette-sdk-go/api/models"
)

func (s *service) GetApplicationProfile(uid string) (*models.V1AppProfile, error) {
	if uid == "" {
		return nil, errUIDRequired
	}
	return s.client.GetApplicationProfile(uid)
}

func (s *service) CreateApplicationProfile(profile *models.V1AppProfileEntity) (string, error) {
	if profile == nil {
		return "", errors.New("application profile entity is required")
	}
	return s.client.CreateApplicationProfile(profile)
}

func (s *service) DeleteApplicationProfile(uid string) error {
	if uid == "" {
		return errUIDRequired
	}
	return s.client.DeleteApplicationProfile(uid)
}
