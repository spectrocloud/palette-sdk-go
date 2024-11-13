package client

import (
	"errors"
	"fmt"

	clientv1 "github.com/spectrocloud/palette-sdk-go/api/client/v1"
	"github.com/spectrocloud/palette-sdk-go/api/models"
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

// AuthRefreshToken refreshes jwt token.
func (h *V1Client) AuthRefreshToken(token string) (*models.V1UserToken, error) {
	params := clientv1.NewV1AuthRefreshParamsWithContext(h.ctx).WithToken(token)
	resp, err := h.Client.V1AuthRefresh(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// GetUsersInfo retrieves the authenticated user info.
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

// GetUserSummaryByEmail retrieves user by email.
func (h *V1Client) GetUserSummaryByEmail(userEmail string) (*models.V1UserSummary, error) {
	// ACL scoped to tenant only
	filterSummary := &models.V1UsersSummarySpec{
		Filter: &models.V1UsersFilterSpec{
			EmailID: &models.V1FilterString{
				Eq: &userEmail,
			},
		},
	}

	summaryParams := clientv1.NewV1UsersSummaryGetParams().WithBody(filterSummary)
	summaryResponse, err := h.Client.V1UsersSummaryGet(summaryParams)
	if err != nil {
		return nil, err
	}
	if summaryResponse.Payload.Items != nil {
		if len(summaryResponse.Payload.Items) == 1 {
			return summaryResponse.Payload.Items[0], nil
		}
		return nil, errors.New("More than one user found with email: " + userEmail)
	}
	return nil, errors.New("user not found for email: " + userEmail)

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

// GetUserByID retrieves an existing user by ID.
func (h *V1Client) GetUserByID(userUID string) (*models.V1User, error) {
	params := clientv1.NewV1UsersUIDGetParams().WithUID(userUID)
	resp, err := h.Client.V1UsersUIDGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// DeleteUser deletes an existing user by UID.
func (h *V1Client) DeleteUser(uid string) error {
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
			return h.DeleteUser(user.Metadata.UID)
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

// UpdateUser update existing user.
func (h *V1Client) UpdateUser(uid string, user *models.V1UserUpdateEntity) error {
	param := clientv1.NewV1UsersUIDUpdateParams().WithUID(uid).WithBody(user)
	_, err := h.Client.V1UsersUIDUpdate(param)
	if err != nil {
		return err
	}
	return nil
}

// AssociateUserProjectRole associate project role for the user.
func (h *V1Client) AssociateUserProjectRole(userUID string, body *models.V1ProjectRolesPatch) error {
	param := clientv1.NewV1UsersProjectRolesPutParams().WithUID(userUID).WithBody(body)
	_, err := h.Client.V1UsersProjectRolesPut(param)
	if err != nil {
		return err
	}
	return nil
}

// GetUserProjectRole get project role to the user.
func (h *V1Client) GetUserProjectRole(userUID string) (*models.V1ProjectRolesEntity, error) {
	param := clientv1.NewV1UsersProjectRolesParams().WithUID(userUID)
	resp, err := h.Client.V1UsersProjectRoles(param)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// AssociateUserWorkspaceRole associate workspace role to the user.
func (h *V1Client) AssociateUserWorkspaceRole(userUID string, body *models.V1WorkspacesRolesPatch) error {
	param := clientv1.NewV1UsersWorkspaceRolesPutParams().WithUserUID(userUID).WithBody(body)
	_, err := h.Client.V1UsersWorkspaceRolesPut(param)
	if err != nil {
		return err
	}
	return nil
}

// GetUserWorkspaceRole get workspace role for the user.
func (h *V1Client) GetUserWorkspaceRole(userUID string) (*models.V1WorkspaceScopeRoles, error) {
	param := clientv1.NewV1UsersWorkspaceGetRolesParams().WithUserUID(userUID)
	resp, err := h.Client.V1UsersWorkspaceGetRoles(param)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// CreateUserResourceRole associate resource role for the user.
func (h *V1Client) CreateUserResourceRole(userUID string, body *models.V1ResourceRolesUpdateEntity) error {
	param := clientv1.NewV1UsersUIDResourceRolesCreateParams().WithUID(userUID).WithBody(body)
	_, err := h.Client.V1UsersUIDResourceRolesCreate(param)
	if err != nil {
		return err
	}
	return nil
}

// GetUserResourceRoles get resource role for the user.
func (h *V1Client) GetUserResourceRoles(userUID string) ([]*models.V1ResourceRolesEntity, error) {
	param := clientv1.NewV1UsersUIDResourceRolesParams().WithUID(userUID)
	resp, err := h.Client.V1UsersUIDResourceRoles(param)
	if err != nil {
		return nil, err
	}
	return resp.Payload.ResourceRoles, nil
}

// DeleteUserResourceRoles delete resource role for the user.
func (h *V1Client) DeleteUserResourceRoles(userUID string, roleUID string) error {
	param := clientv1.NewV1UsersUIDResourceRolesUIDDeleteParams().WithUID(userUID).WithResourceRoleUID(roleUID)
	_, err := h.Client.V1UsersUIDResourceRolesUIDDelete(param)
	if err != nil {
		return err
	}
	return nil
}

// AssociateUserTenantRole associate tenant role to the user.
func (h *V1Client) AssociateUserTenantRole(userUID string, body *models.V1UserRoleUIDs) error {
	param := clientv1.NewV1UsersUIDRolesUpdateParams().WithUID(userUID).WithBody(body)
	_, err := h.Client.V1UsersUIDRolesUpdate(param)
	if err != nil {
		return err
	}
	return nil
}

// GetUserTenantRole get tenant roles for the user.
func (h *V1Client) GetUserTenantRole(userUID string) (*models.V1UserRolesEntity, error) {
	param := clientv1.NewV1UsersUIDRolesParams().WithUID(userUID)
	resp, err := h.Client.V1UsersUIDRoles(param)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}
