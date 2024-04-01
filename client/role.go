package client

import (
	"fmt"

	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
)

func (h *V1Client) GetRole(roleName string) (*models.V1Role, error) {
	params := clientV1.NewV1RolesListParamsWithContext(h.ctx)
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
