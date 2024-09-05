package client

// CRUDL operations on API keys are all tenant scoped.
// See: hubble/services/service/user/internal/service/apikey/apikey_acl.go

import (
	"fmt"

	clientv1 "github.com/spectrocloud/palette-sdk-go/api/client/v1"
	"github.com/spectrocloud/palette-sdk-go/api/models"
)

// GetAPIKeys retrieves all API Keys.
func (h *V1Client) GetAPIKeys() (*models.V1APIKeys, error) {
	params := clientv1.NewV1APIKeysListParams()
	resp, err := h.Client.V1APIKeysList(params)
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

// DeleteAPIKeyByName deletes an existing API Key by name.
func (h *V1Client) DeleteAPIKeyByName(name string) error {
	keys, err := h.GetAPIKeys()
	if err != nil {
		return err
	}
	for _, key := range keys.Items {
		if key.Metadata.Name == name {
			return h.DeleteAPIKey(key.Metadata.UID)
		}
	}
	return fmt.Errorf("api key with name '%s' not found", name)
}

// DeleteAPIKey deletes an existing API Key by UID.
func (h *V1Client) DeleteAPIKey(uid string) error {
	params := clientv1.NewV1APIKeysUIDDeleteParams().WithUID(uid)
	_, err := h.Client.V1APIKeysUIDDelete(params)
	if err != nil {
		return err
	}
	return nil
}
