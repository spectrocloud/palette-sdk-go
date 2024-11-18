package client

import (
	"fmt"

	clientv1 "github.com/spectrocloud/palette-sdk-go/api/client/v1"
	"github.com/spectrocloud/palette-sdk-go/api/models"
)

// GetPermissionByName retrieves an existing permission by name and permissionScope(project, tenant & resource).
func (h *V1Client) GetPermissionByName(permissionName string, permissionScope string) (*models.V1Permission, error) {
	// ACL scoped to tenant only
	params := clientv1.NewV1PermissionsListParams().WithScope(&permissionScope)
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
