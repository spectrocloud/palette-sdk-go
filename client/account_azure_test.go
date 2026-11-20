package client

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/spectrocloud/palette-sdk-go/api/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// PLT-2489: validateCloudAccountAzure must call the overlord-scoped MSI
// validate endpoint for identityType=UserAssignedMSI accounts instead of the
// service-principal-only /v1/clouds/azure/account/validate endpoint, since
// UserAssignedMSI accounts carry no clientSecret for that endpoint to check.

func newTestAzureAccount(overlordUID string, identityType *string) *models.V1AzureAccount {
	clientID := "test-client-id"
	tenantID := "test-tenant-id"
	return &models.V1AzureAccount{
		Metadata: &models.V1ObjectMeta{
			Name:        "test-azure-account",
			Annotations: map[string]string{OverlordUID: overlordUID},
		},
		Spec: &models.V1AzureCloudAccount{
			ClientID:     &clientID,
			TenantID:     &tenantID,
			IdentityType: identityType,
		},
	}
}

func pcgRunningPayload() string {
	return `{
		"metadata": {"name": "test-pcg", "uid": "pcg-1"},
		"status": {"state": "Running"}
	}`
}

func TestValidateCloudAccountAzure_UserAssignedMSI_CallsOverlordValidate(t *testing.T) {
	t.Parallel()

	var hitOverlordValidate, hitServicePrincipalValidate bool
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		switch {
		case r.URL.Path == "/v1/overlords/pcg-1":
			rw.WriteHeader(http.StatusOK)
			_, _ = rw.Write([]byte(pcgRunningPayload()))
		case r.URL.Path == "/v1/overlords/azure/pcg-1/account/validate":
			hitOverlordValidate = true
			var body struct {
				Account *models.V1AzureCloudAccount `json:"account"`
			}
			require.NoError(t, json.NewDecoder(r.Body).Decode(&body))
			assert.Equal(t, models.V1AzureCloudAccountIdentityTypeUserAssignedMSI, *body.Account.IdentityType)
			rw.WriteHeader(http.StatusOK)
			_, _ = rw.Write([]byte(`{"managementIdentityResourceId": "resolved-mi-resource-id"}`))
		case r.URL.Path == "/v1/clouds/azure/account/validate":
			hitServicePrincipalValidate = true
			rw.WriteHeader(http.StatusNoContent)
		default:
			t.Errorf("unexpected request path: %s", r.URL.Path)
			rw.WriteHeader(http.StatusNotFound)
		}
	}))
	t.Cleanup(server.Close)

	c := New(
		WithPaletteURI(strings.TrimPrefix(server.URL, "http://")),
		WithAPIKey("test-key"),
		WithSchemes([]string{"http"}),
	)

	identityType := models.V1AzureCloudAccountIdentityTypeUserAssignedMSI
	err := c.validateCloudAccountAzure(newTestAzureAccount("pcg-1", &identityType))
	require.NoError(t, err)
	assert.True(t, hitOverlordValidate, "expected the overlord-scoped MSI validate endpoint to be called")
	assert.False(t, hitServicePrincipalValidate, "the service-principal validate endpoint must not be called for a UserAssignedMSI account")
}

func TestValidateCloudAccountAzure_ServicePrincipal_CallsLegacyValidate(t *testing.T) {
	t.Parallel()

	var hitOverlordValidate, hitServicePrincipalValidate bool
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		switch {
		case r.URL.Path == "/v1/overlords/pcg-1":
			rw.WriteHeader(http.StatusOK)
			_, _ = rw.Write([]byte(pcgRunningPayload()))
		case r.URL.Path == "/v1/overlords/azure/pcg-1/account/validate":
			hitOverlordValidate = true
			rw.WriteHeader(http.StatusOK)
			_, _ = rw.Write([]byte(`{}`))
		case r.URL.Path == "/v1/clouds/azure/account/validate":
			hitServicePrincipalValidate = true
			rw.WriteHeader(http.StatusNoContent)
		default:
			t.Errorf("unexpected request path: %s", r.URL.Path)
			rw.WriteHeader(http.StatusNotFound)
		}
	}))
	t.Cleanup(server.Close)

	c := New(
		WithPaletteURI(strings.TrimPrefix(server.URL, "http://")),
		WithAPIKey("test-key"),
		WithSchemes([]string{"http"}),
	)

	identityType := models.V1AzureCloudAccountIdentityTypeServicePrincipal
	err := c.validateCloudAccountAzure(newTestAzureAccount("pcg-1", &identityType))
	require.NoError(t, err)
	assert.True(t, hitServicePrincipalValidate, "expected the legacy service-principal validate endpoint to be called")
	assert.False(t, hitOverlordValidate, "the overlord-scoped MSI validate endpoint must not be called for a ServicePrincipal account")
}

// TestValidateCloudAccountAzure_NilIdentityType_DefaultsToServicePrincipal
// guards the unset-identityType default (pre-PLT-2489 accounts, and any
// config that omits the new field) against regressing to the MSI path.
func TestValidateCloudAccountAzure_NilIdentityType_DefaultsToServicePrincipal(t *testing.T) {
	t.Parallel()

	var hitServicePrincipalValidate bool
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		switch r.URL.Path {
		case "/v1/overlords/pcg-1":
			rw.WriteHeader(http.StatusOK)
			_, _ = rw.Write([]byte(pcgRunningPayload()))
		case "/v1/clouds/azure/account/validate":
			hitServicePrincipalValidate = true
			rw.WriteHeader(http.StatusNoContent)
		default:
			t.Errorf("unexpected request path: %s", r.URL.Path)
			rw.WriteHeader(http.StatusNotFound)
		}
	}))
	t.Cleanup(server.Close)

	c := New(
		WithPaletteURI(strings.TrimPrefix(server.URL, "http://")),
		WithAPIKey("test-key"),
		WithSchemes([]string{"http"}),
	)

	err := c.validateCloudAccountAzure(newTestAzureAccount("pcg-1", nil))
	require.NoError(t, err)
	assert.True(t, hitServicePrincipalValidate)
}
