package access

import (
	"errors"
	"time"

	"github.com/spectrocloud/palette-sdk-go/api/models"
)

func (s *service) CreateAPIKey(name string, annotations map[string]string, expiry time.Duration) (string, error) {
	if name == "" {
		return "", errors.New("API key name is required")
	}
	return s.client.CreateAPIKey(name, annotations, expiry)
}

func (s *service) GetAPIKeys() (*models.V1APIKeys, error) {
	return s.client.GetAPIKeys()
}

func (s *service) DeleteAPIKey(uid string) error {
	if uid == "" {
		return errUIDRequired
	}
	return s.client.DeleteAPIKey(uid)
}
