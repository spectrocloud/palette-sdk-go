package cloudaccounts

import (
	"errors"

	"github.com/spectrocloud/palette-sdk-go/api/models"
)

func (s *service) GetCustomByName(cloudType, name, scope string) (*models.V1CustomAccount, error) {
	if cloudType == "" {
		return nil, errors.New("cloud type is required")
	}
	if name == "" {
		return nil, errors.New("account name is required")
	}
	return s.client.GetCloudAccountCustomByName(cloudType, name, scope)
}

func (s *service) CreateCustom(account *models.V1CustomAccountEntity, cloudType string) (string, error) {
	if cloudType == "" {
		return "", errors.New("cloud type is required")
	}
	if account == nil {
		return "", errors.New("custom cloud account is required")
	}
	if err := s.client.ValidateCustomCloudType(cloudType); err != nil {
		return "", err
	}
	return s.client.CreateAccountCustomCloud(account, cloudType)
}

func (s *service) UpdateCustom(uid string, account *models.V1CustomAccountEntity, cloudType string) error {
	if uid == "" {
		return errUIDRequired
	}
	if cloudType == "" {
		return errors.New("cloud type is required")
	}
	return s.client.UpdateAccountCustomCloud(uid, account, cloudType)
}

func (s *service) DeleteCustom(uid, cloudType string) error {
	if uid == "" {
		return errUIDRequired
	}
	if cloudType == "" {
		return errors.New("cloud type is required")
	}
	return s.client.DeleteCloudAccountCustomCloud(uid, cloudType)
}

func (s *service) GetCustom(uid, cloudType string) (*models.V1CustomAccount, error) {
	if uid == "" {
		return nil, errUIDRequired
	}
	if cloudType == "" {
		return nil, errors.New("cloud type is required")
	}
	return s.client.GetCustomCloudAccount(uid, cloudType)
}
