package client

import (
	"fmt"

	clientv1 "github.com/spectrocloud/palette-sdk-go/api/client/v1"
	"github.com/spectrocloud/palette-sdk-go/api/models"
)

// CRUDL operations on roles are all tenant scoped.
// See: hubble/services/service/user/internal/service/role/role_acl.go

// GetRole retrieves an existing role by name.
func (h *V1Client) GetRole(roleName string) (*models.V1Role, error) {
	// ACL scoped to tenant only
	params := clientv1.NewV1RolesListParams()
	resp, err := h.Client.V1RolesList(params)
	if err != nil {
		return nil, err
	}

	for _, role := range resp.Payload.Items {
		if role.Metadata.Name == roleName {
			return role, nil
		}
	}

	return nil, fmt.Errorf("role '%s' not found", roleName)
}

// GetRoleByID retrieves an existing role by ID.
func (h *V1Client) GetRoleByID(roleUID string) (*models.V1Role, error) {
	// ACL scoped to tenant only
	params := clientv1.NewV1RolesUIDGetParams().WithUID(roleUID)
	resp, err := h.Client.V1RolesUIDGet(params)
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

// CreateRole create new role.
func (h *V1Client) CreateRole(role *models.V1Role) (string, error) {
	// ACL scoped to tenant only
	params := clientv1.NewV1RolesCreateParams().WithBody(role)
	resp, err := h.Client.V1RolesCreate(params)
	if err != nil {
		return "", err
	}

	return *resp.Payload.UID, nil
}

// UpdateRole Update existing role with ID
func (h *V1Client) UpdateRole(role *models.V1Role, roleUID string) error {
	// ACL scoped to tenant only
	params := clientv1.NewV1RolesUIDUpdateParams().WithBody(role).WithUID(roleUID)
	_, err := h.Client.V1RolesUIDUpdate(params)
	return err
}

// DeleteRole Delete existing role with ID
func (h *V1Client) DeleteRole(roleUID string) error {
	// ACL scoped to tenant only
	params := clientv1.NewV1RolesUIDDeleteParams().WithUID(roleUID)
	_, err := h.Client.V1RolesUIDDelete(params)
	return err
}
