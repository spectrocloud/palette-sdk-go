package client

import (
	"fmt"

	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
)

func (h *V1Client) GetUsers() (*models.V1Users, error) {
	limit := int64(0)
	params := clientV1.NewV1UsersListParams().WithLimit(&limit)
	users, err := h.GetClient().V1UsersList(params)
	if err != nil {
		return nil, err
	}

	return users.Payload, nil
}

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
