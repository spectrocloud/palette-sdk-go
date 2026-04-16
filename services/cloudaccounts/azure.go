package cloudaccounts

import (
	"errors"

	"github.com/spectrocloud/palette-sdk-go/api/models"
)

func (s *service) GetAzureByName(name, scope string) (*models.V1AzureAccount, error) {
	if name == "" {
		return nil, errors.New("account name is required")
	}
	return s.client.GetCloudAccountAzureByName(name, scope)
}

func (s *service) CreateAzure(account *models.V1AzureAccount) (string, error) {
	if account == nil {
		return "", errors.New("Azure account is required")
	}
	return s.client.CreateCloudAccountAzure(account)
}

func (s *service) UpdateAzure(account *models.V1AzureAccount) error {
	return s.client.UpdateCloudAccountAzure(account)
}

func (s *service) DeleteAzure(uid string) error {
	if uid == "" {
		return errUIDRequired
	}
	return s.client.DeleteCloudAccountAzure(uid)
}

func (s *service) GetAzure(uid string) (*models.V1AzureAccount, error) {
	if uid == "" {
		return nil, errUIDRequired
	}
	return s.client.GetCloudAccountAzure(uid)
}
