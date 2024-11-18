package client

import (
	"fmt"

	clientv1 "github.com/spectrocloud/palette-sdk-go/api/client/v1"
	"github.com/spectrocloud/palette-sdk-go/api/models"
)

// PermissionScope defines the scope of a permission.
type PermissionScope string

const (
	// PermissionScopeProject represents a project-level scope.
	PermissionScopeProject PermissionScope = "project"

	// PermissionScopeTenant represents a tenant-level scope.
	PermissionScopeTenant PermissionScope = "tenant"

	// PermissionScopeResource represents a resource-level scope.
	PermissionScopeResource PermissionScope = "resource"
)

// GetPermissionByName retrieves an existing permission by name and permissionScope(project, tenant & resource).
func (h *V1Client) GetPermissionByName(permissionName string, permissionScope PermissionScope) (*models.V1Permission, error) {
	// ACL scoped to tenant only
	permScope := string(permissionScope)
	params := clientv1.NewV1PermissionsListParams().WithScope(&permScope)
	resp, err := h.Client.V1PermissionsList(params)
	if err != nil {
		return nil, err
	}

	for _, permission := range resp.Payload {
		if permission.Name == permissionName {
			return permission, nil
		}
	}

	return nil, fmt.Errorf("permission name '%s' not found in scope '%s'", permissionName, permissionScope)
}
