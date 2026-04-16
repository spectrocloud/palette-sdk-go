package access

import (
	"errors"

	"github.com/spectrocloud/palette-sdk-go/api/models"
)

func (s *service) CreateRole(role *models.V1Role) (string, error) {
	if role == nil {
		return "", errors.New("role is required")
	}
	return s.client.CreateRole(role)
}

func (s *service) GetRole(name string) (*models.V1Role, error) {
	if name == "" {
		return nil, errors.New("role name is required")
	}
	return s.client.GetRole(name)
}

func (s *service) GetRoleByID(uid string) (*models.V1Role, error) {
	if uid == "" {
		return nil, errUIDRequired
	}
	return s.client.GetRoleByID(uid)
}

func (s *service) UpdateRole(role *models.V1Role, uid string) error {
	if uid == "" {
		return errUIDRequired
	}
	return s.client.UpdateRole(role, uid)
}

func (s *service) DeleteRole(uid string) error {
	if uid == "" {
		return errUIDRequired
	}
	return s.client.DeleteRole(uid)
}
