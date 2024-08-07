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
