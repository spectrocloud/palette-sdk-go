package cloudaccounts

import (
	"errors"

	"github.com/spectrocloud/palette-sdk-go/api/models"
)

func (s *service) GetCloudStackByName(name, scope string) (*models.V1CloudStackAccount, error) {
	if name == "" {
		return nil, errors.New("account name is required")
	}
	return s.client.GetCloudAccountApacheCloudStackByName(name, scope)
}

func (s *service) CreateCloudStack(account *models.V1CloudStackAccount) (string, error) {
	if account == nil {
		return "", errors.New("CloudStack account is required")
	}
	return s.client.CreateCloudAccountCloudStack(account)
}

func (s *service) UpdateCloudStack(account *models.V1CloudStackAccount) error {
	return s.client.UpdateCloudAccountCloudStack(account)
}

func (s *service) DeleteCloudStack(uid string) error {
	if uid == "" {
		return errUIDRequired
	}
	return s.client.DeleteCloudAccountCloudStack(uid)
}

func (s *service) GetCloudStack(uid string) (*models.V1CloudStackAccount, error) {
	if uid == "" {
		return nil, errUIDRequired
	}
	return s.client.GetCloudAccountCloudStack(uid)
}
