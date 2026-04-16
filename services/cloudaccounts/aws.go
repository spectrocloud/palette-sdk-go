package cloudaccounts

import (
	"errors"

	"github.com/spectrocloud/palette-sdk-go/api/models"
)

func (s *service) GetAwsByName(name, scope string) (*models.V1AwsAccount, error) {
	if name == "" {
		return nil, errors.New("account name is required")
	}
	return s.client.GetCloudAccountAwsByName(name, scope)
}

func (s *service) CreateAws(account *models.V1AwsAccount) (string, error) {
	if account == nil {
		return "", errors.New("AWS account is required")
	}
	return s.client.CreateCloudAccountAws(account)
}

func (s *service) UpdateAws(account *models.V1AwsAccount) error {
	return s.client.UpdateCloudAccountAws(account)
}

func (s *service) DeleteAws(uid string) error {
	if uid == "" {
		return errUIDRequired
	}
	return s.client.DeleteCloudAccountAws(uid)
}

func (s *service) GetAws(uid string) (*models.V1AwsAccount, error) {
	if uid == "" {
		return nil, errUIDRequired
	}
	return s.client.GetCloudAccountAws(uid)
}
