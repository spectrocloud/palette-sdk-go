package access

import (
	"errors"

	"github.com/spectrocloud/palette-sdk-go/api/models"
)

func (s *service) Authenticate(body *models.V1AuthLogin) (*models.V1UserToken, error) {
	if body == nil {
		return nil, errors.New("auth login body is required")
	}
	return s.client.Authenticate(body)
}

func (s *service) AuthRefreshToken(token string) (*models.V1UserToken, error) {
	if token == "" {
		return nil, errors.New("refresh token is required")
	}
	return s.client.AuthRefreshToken(token)
}
