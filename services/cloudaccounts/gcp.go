package cloudaccounts

import (
	"errors"

	"github.com/spectrocloud/palette-sdk-go/api/models"
)

func (s *service) GetGcpByName(name, scope string) (*models.V1GcpAccount, error) {
	if name == "" {
		return nil, errors.New("account name is required")
	}
	return s.client.GetCloudAccountGcpByName(name, scope)
}

func (s *service) CreateGcp(account *models.V1GcpAccountEntity) (string, error) {
	if account == nil {
		return "", errors.New("GCP account is required")
	}
	return s.client.CreateCloudAccountGcp(account)
}

func (s *service) UpdateGcp(account *models.V1GcpAccountEntity) error {
	return s.client.UpdateCloudAccountGcp(account)
}

func (s *service) DeleteGcp(uid string) error {
	if uid == "" {
		return errUIDRequired
	}
	return s.client.DeleteCloudAccountGcp(uid)
}

func (s *service) GetGcp(uid string) (*models.V1GcpAccount, error) {
	if uid == "" {
		return nil, errUIDRequired
	}
	return s.client.GetCloudAccountGcp(uid)
}
