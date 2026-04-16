package cloudaccounts

import (
	"errors"

	"github.com/spectrocloud/palette-sdk-go/api/models"
)

func (s *service) GetVsphereByName(name, scope string) (*models.V1VsphereAccount, error) {
	if name == "" {
		return nil, errors.New("account name is required")
	}
	return s.client.GetCloudAccountVsphereByName(name, scope)
}

func (s *service) CreateVsphere(account *models.V1VsphereAccount) (string, error) {
	if account == nil {
		return "", errors.New("vSphere account is required")
	}
	return s.client.CreateCloudAccountVsphere(account)
}

func (s *service) UpdateVsphere(account *models.V1VsphereAccount) error {
	return s.client.UpdateCloudAccountVsphere(account)
}

func (s *service) DeleteVsphere(uid string) error {
	if uid == "" {
		return errUIDRequired
	}
	return s.client.DeleteCloudAccountVsphere(uid)
}

func (s *service) GetVsphere(uid string) (*models.V1VsphereAccount, error) {
	if uid == "" {
		return nil, errUIDRequired
	}
	return s.client.GetCloudAccountVsphere(uid)
}
