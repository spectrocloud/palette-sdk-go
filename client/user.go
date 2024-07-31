package client

import (
	"fmt"

	clientv1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
	"github.com/spectrocloud/palette-sdk-go/client/apiutil"
)

// Authenticate authenticates a user.
// Returns a JWT token.
func (h *V1Client) Authenticate(body *models.V1AuthLogin) (*models.V1UserToken, error) {
	params := clientv1.NewV1AuthenticateParams().
		WithBody(body)
	resp, err := h.Client.V1Authenticate(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// GetMe retrieves the authenticated user.
func (h *V1Client) GetUsersInfo() (*models.V1UserInfo, error) {
	params := clientv1.NewV1UsersInfoGetParamsWithContext(h.ctx)
	resp, err := h.Client.V1UsersInfoGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// CRUDL operations on users are all tenant scoped.
// See: hubble/services/svccore/perms/user_acl.go

// GetUsers retrieves all users.
func (h *V1Client) GetUsers() (*models.V1Users, error) {
	// ACL scoped to tenant only
	params := clientv1.NewV1UsersListParams().
		WithLimit(apiutil.Ptr(int64(0)))
	resp, err := h.Client.V1UsersList(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// GetUserByName retrieves an existing user by name.
func (h *V1Client) GetUserByName(name string) (*models.V1User, error) {
	users, err := h.GetUsers()
	if err != nil {
		return nil, err
	}
	for _, user := range users.Items {
		if user.Metadata.Name == name {
			return user, nil
		}
	}
	return nil, fmt.Errorf("user with name '%s' not found", name)
}

// GetUserByEmail retrieves an existing user by email.
func (h *V1Client) GetUserByEmail(email string) (*models.V1User, error) {
	users, err := h.GetUsers()
	if err != nil {
		return nil, err
	}
	for _, user := range users.Items {
		if user.Spec.EmailID == email {
			return user, nil
		}
	}
	return nil, fmt.Errorf("user with email '%s' not found", email)
}

// DeleteUserByUID deletes an existing user by UID.
func (h *V1Client) DeleteUserByUID(uid string) error {
	params := clientv1.NewV1UsersUIDDeleteParams().WithUID(uid)
	_, err := h.Client.V1UsersUIDDelete(params)
	if err != nil {
		return err
	}
	return nil
}

// DeleteUserByName deletes an existing user by name.
func (h *V1Client) DeleteUserByName(name string) error {
	users, err := h.GetUsers()
	if err != nil {
		return err
	}
	for _, user := range users.Items {
		if user.Metadata.Name == name {
			return h.DeleteUserByUID(user.Metadata.UID)
		}
	}
	return fmt.Errorf("user with name '%s' not found", name)
}

// CreateUser creates a new user.
func (h *V1Client) CreateUser(user *models.V1UserEntity) (string, error) {
	param := clientv1.NewV1UsersCreateParams().WithBody(user)
	resp, err := h.Client.V1UsersCreate(param)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}
