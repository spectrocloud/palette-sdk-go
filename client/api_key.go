package client

// CRUDL operations on API keys are all tenant scoped.
// See: hubble/services/service/user/internal/service/apikey/apikey_acl.go

import (
	"fmt"
	"time"

	clientv1 "github.com/spectrocloud/palette-sdk-go/api/client/v1"
	"github.com/spectrocloud/palette-sdk-go/api/models"
)

// CreateAPIKey creates a new API key for the tenant
func (h *V1Client) CreateAPIKey(name string, annotations map[string]string, expiry time.Duration) (string, error) {
	tenantUID, err := h.GetTenantUID()
	if err != nil {
		return "", fmt.Errorf("failed to get tenant UID: %w", err)
	}
	body := &models.V1APIKeyEntity{
		Metadata: &models.V1ObjectMeta{
			Name:        name,
			Annotations: annotations,
		},
		Spec: &models.V1APIKeySpecEntity{
			UserUID: tenantUID,
			Expiry:  models.V1Time(time.Now().Add(expiry)),
		},
	}
	params := clientv1.NewV1APIKeysCreateParams().WithBody(body)
	resp, err := h.Client.V1APIKeysCreate(params)
	if err != nil {
		return "", fmt.Errorf("failed to create API key: %w", err)
	}
	return resp.Payload.APIKey, nil
}

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
