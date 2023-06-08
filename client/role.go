package client

import (
	"fmt"

	"github.com/spectrocloud/hapi/models"
	userC "github.com/spectrocloud/hapi/user/client/v1"
)

func (h *V1Client) GetRole(roleName string) (*models.V1Role, error) {
	client, err := h.GetUserClient()
	if err != nil {
		return nil, err
	}

	params := userC.NewV1RolesListParams()
	roles, err := client.V1RolesList(params)
	if err != nil {
		return nil, err
	}

	for _, role := range roles.Payload.Items {
		if role.Metadata.Name == roleName {
			return role, nil
		}
	}

	return nil, fmt.Errorf("role '%s' not found", roleName)
}