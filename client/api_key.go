package client

// CRUDL operations on API keys are all tenant scoped.
// See: hubble/services/service/user/internal/service/apikey/apikey_acl.go

import (
	"fmt"
	"github.com/go-openapi/strfmt"
	log "github.com/sirupsen/logrus"
	"time"

	clientv1 "github.com/spectrocloud/palette-sdk-go/api/client/v1"
	"github.com/spectrocloud/palette-sdk-go/api/models"
)

func (h *V1Client) CreateAPIKey(name, TenantUID, emailID, org, password string) (string, error) {
	annotations := map[string]string{
		"Description": "Automation-Test"}
	body := &models.V1APIKeyEntity{
		Metadata: &models.V1ObjectMeta{
			Name:        name,
			Annotations: annotations,
		},
		Spec: &models.V1APIKeySpecEntity{
			UserUID: TenantUID,
			Expiry:  models.V1Time(time.Now().Add(5 * time.Hour)), // Set expiry time
		},
	}
	if _, err := h.AuthLogin(emailID, org, password); err != nil {
		return "", fmt.Errorf("failed to login as sysadmin: %w", err)
	}
	// Check if the JWT is set
	if h.jwt == "" {
		return "", fmt.Errorf("JWT is not set after login")
	}
	log.Infof("JWT before creating API key: %s", h.jwt)
	h.Client = clientv1.New(h.getTransport(), strfmt.Default)

	params := clientv1.NewV1APIKeysCreateParams().WithBody(body)
	resp, err := h.Client.V1APIKeysCreate(params)
	if err != nil {
		return "", fmt.Errorf("failed to create API key: %w", err)
	}
	return resp.Payload.UID, nil
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

func (h *V1Client) AuthLogin(emailID, org, password string) (string, error) {
	loginRequest := &models.V1AuthLogin{
		EmailID:  emailID,
		Org:      org,
		Password: strfmt.Password(password),
	}
	loginParams := clientv1.NewV1AuthenticateParams().WithBody(loginRequest)
	resp, err := h.Client.V1Authenticate(loginParams)
	if err != nil {
		return "", fmt.Errorf("failed to login: %w", err)
	}
	h.jwt = resp.Payload.Authorization
	return h.jwt, nil
}
