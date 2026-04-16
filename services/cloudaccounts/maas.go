package cloudaccounts

import (
	"errors"

	"github.com/spectrocloud/palette-sdk-go/api/models"
)

func (s *service) GetMaasByName(name, scope string) (*models.V1MaasAccount, error) {
	if name == "" {
		return nil, errors.New("account name is required")
	}
	return s.client.GetCloudAccountMaasByName(name, scope)
}

func (s *service) CreateMaas(account *models.V1MaasAccount) (string, error) {
	if account == nil {
		return "", errors.New("MAAS account is required")
	}
	return s.client.CreateCloudAccountMaas(account)
}

func (s *service) UpdateMaas(account *models.V1MaasAccount) error {
	return s.client.UpdateCloudAccountMaas(account)
}

func (s *service) DeleteMaas(uid string) error {
	if uid == "" {
		return errUIDRequired
	}
	return s.client.DeleteCloudAccountMaas(uid)
}

func (s *service) GetMaas(uid string) (*models.V1MaasAccount, error) {
	if uid == "" {
		return nil, errUIDRequired
	}
	return s.client.GetCloudAccountMaas(uid)
}
