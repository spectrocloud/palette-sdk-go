package access

import (
	"errors"

	"github.com/spectrocloud/palette-sdk-go/api/models"
)

func (s *service) CreateUser(user *models.V1UserEntity) (string, error) {
	if user == nil {
		return "", errors.New("user entity is required")
	}
	return s.client.CreateUser(user)
}

func (s *service) GetUserByID(uid string) (*models.V1User, error) {
	if uid == "" {
		return nil, errUIDRequired
	}
	return s.client.GetUserByID(uid)
}

func (s *service) GetUserByName(name string) (*models.V1User, error) {
	if name == "" {
		return nil, errors.New("user name is required")
	}
	return s.client.GetUserByName(name)
}

func (s *service) GetUserByEmail(email string) (*models.V1User, error) {
	if email == "" {
		return nil, errors.New("user email is required")
	}
	return s.client.GetUserByEmail(email)
}

func (s *service) GetUsers() (*models.V1Users, error) {
	return s.client.GetUsers()
}

func (s *service) UpdateUser(uid string, user *models.V1UserUpdateEntity) error {
	if uid == "" {
		return errUIDRequired
	}
	return s.client.UpdateUser(uid, user)
}

func (s *service) DeleteUser(uid string) error {
	if uid == "" {
		return errUIDRequired
	}
	return s.client.DeleteUser(uid)
}

func (s *service) AssociateUserProjectRole(userUID string, body *models.V1ProjectRolesPatch) error {
	if userUID == "" {
		return errUIDRequired
	}
	return s.client.AssociateUserProjectRole(userUID, body)
}

func (s *service) AssociateUserTenantRole(userUID string, body *models.V1UserRoleUIDs) error {
	if userUID == "" {
		return errUIDRequired
	}
	return s.client.AssociateUserTenantRole(userUID, body)
}

func (s *service) AssociateUserWorkspaceRole(userUID string, body *models.V1WorkspacesRolesPatch) error {
	if userUID == "" {
		return errUIDRequired
	}
	return s.client.AssociateUserWorkspaceRole(userUID, body)
}
