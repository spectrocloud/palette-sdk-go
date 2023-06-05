package client

import (
	"fmt"

	"github.com/spectrocloud/hapi/models"
	userC "github.com/spectrocloud/hapi/user/client/v1"
)

func (h *V1Client) GetUser(name string) (*models.V1User, error) {
	client, err := h.GetUserClient()
	if err != nil {
		return nil, err
	}

	params := userC.NewV1UsersListParams()
	users, err := client.V1UsersList(params)
	if err != nil {
		return nil, err
	}

	for _, user := range users.Payload.Items {
		if user.Metadata.Name == name {
			return user, nil
		}
	}

	return nil, fmt.Errorf("user '%s' not found", name)
}
