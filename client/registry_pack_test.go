package client

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/spectrocloud/palette-sdk-go/api/models"
)

// PLT-2441: Pack registry lifecycle wrappers for the /v1/registries/pack APIs.

type recordedRequest struct {
	method string
	path   string
	query  string
	body   string
}

func newPackRegistryTestClient(t *testing.T, status int, response string, got *recordedRequest) *V1Client {
	t.Helper()
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		*got = recordedRequest{method: r.Method, path: r.URL.Path, query: r.URL.RawQuery, body: string(body)}
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(status)
		if response != "" {
			_, _ = rw.Write([]byte(response))
		}
	}))
	t.Cleanup(server.Close)
	return New(
		WithPaletteURI(strings.TrimPrefix(server.URL, "http://")),
		WithAPIKey("test-key"),
		WithSchemes([]string{"http"}),
	)
}

func testPackRegistry() *models.V1PackRegistry {
	endpoint := "https://registry.example.com:5000"
	return &models.V1PackRegistry{
		Metadata: &models.V1ObjectMeta{Name: "my-pack-registry"},
		Spec: &models.V1PackRegistrySpec{
			Endpoint: &endpoint,
			Private:  true,
			Auth:     &models.V1RegistryAuth{Type: "basic", Username: "admin", Password: "secret"},
		},
	}
}

func TestCreatePackRegistry(t *testing.T) {
	t.Parallel()
	var got recordedRequest
	c := newPackRegistryTestClient(t, http.StatusCreated, `{"uid":"pack-reg-uid"}`, &got)

	uid, err := c.CreatePackRegistry(testPackRegistry())
	require.NoError(t, err)
	assert.Equal(t, "pack-reg-uid", uid)
	assert.Equal(t, http.MethodPost, got.method)
	assert.Equal(t, "/v1/registries/pack", got.path)

	var sent models.V1PackRegistry
	require.NoError(t, json.Unmarshal([]byte(got.body), &sent))
	assert.Equal(t, "my-pack-registry", sent.Metadata.Name)
	assert.Equal(t, "https://registry.example.com:5000", *sent.Spec.Endpoint)
	assert.True(t, sent.Spec.Private)
	assert.Equal(t, "basic", sent.Spec.Auth.Type)
}

func TestCreatePackRegistryWithoutUID(t *testing.T) {
	t.Parallel()
	var got recordedRequest
	c := newPackRegistryTestClient(t, http.StatusCreated, `{}`, &got)

	_, err := c.CreatePackRegistry(testPackRegistry())
	assert.ErrorContains(t, err, "no UID")
}

func TestGetPackRegistryByUID(t *testing.T) {
	t.Parallel()
	var got recordedRequest
	c := newPackRegistryTestClient(t, http.StatusOK,
		`{"metadata":{"name":"my-pack-registry","uid":"pack-reg-uid"},"spec":{"endpoint":"https://registry.example.com:5000","private":true,"auth":{"type":"noAuth"}}}`, &got)

	registry, err := c.GetPackRegistryByUID("pack-reg-uid")
	require.NoError(t, err)
	assert.Equal(t, http.MethodGet, got.method)
	assert.Equal(t, "/v1/registries/pack/pack-reg-uid", got.path)
	assert.Equal(t, "my-pack-registry", registry.Metadata.Name)
	assert.Equal(t, "https://registry.example.com:5000", *registry.Spec.Endpoint)
	assert.True(t, registry.Spec.Private)
}

func TestUpdatePackRegistry(t *testing.T) {
	t.Parallel()
	var got recordedRequest
	c := newPackRegistryTestClient(t, http.StatusNoContent, "", &got)

	require.NoError(t, c.UpdatePackRegistry("pack-reg-uid", testPackRegistry()))
	assert.Equal(t, http.MethodPut, got.method)
	assert.Equal(t, "/v1/registries/pack/pack-reg-uid", got.path)
	assert.Contains(t, got.body, `"endpoint":"https://registry.example.com:5000"`)
}

func TestDeletePackRegistry(t *testing.T) {
	t.Parallel()
	var got recordedRequest
	c := newPackRegistryTestClient(t, http.StatusNoContent, "", &got)

	require.NoError(t, c.DeletePackRegistry("pack-reg-uid"))
	assert.Equal(t, http.MethodDelete, got.method)
	assert.Equal(t, "/v1/registries/pack/pack-reg-uid", got.path)
}

func TestValidatePackRegistry(t *testing.T) {
	t.Parallel()
	var got recordedRequest
	c := newPackRegistryTestClient(t, http.StatusNoContent, "", &got)

	require.NoError(t, c.ValidatePackRegistry(testPackRegistry().Spec))
	assert.Equal(t, http.MethodPost, got.method)
	assert.Equal(t, "/v1/registries/pack/validate", got.path)
	assert.Contains(t, got.body, `"endpoint":"https://registry.example.com:5000"`)
}

func TestValidatePackRegistryRejected(t *testing.T) {
	t.Parallel()
	var got recordedRequest
	c := newPackRegistryTestClient(t, http.StatusBadRequest,
		`{"code":"InvalidRegistry","message":"unable to reach registry"}`, &got)

	err := c.ValidatePackRegistry(testPackRegistry().Spec)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "unable to reach registry")
}

func TestSyncPackRegistry(t *testing.T) {
	t.Parallel()
	for _, force := range []bool{false, true} {
		var got recordedRequest
		c := newPackRegistryTestClient(t, http.StatusAccepted, "", &got)

		require.NoError(t, c.SyncPackRegistry("pack-reg-uid", force))
		assert.Equal(t, http.MethodPost, got.method)
		assert.Equal(t, "/v1/registries/pack/pack-reg-uid/sync", got.path)
		if force {
			assert.Equal(t, "forceSync=true", got.query)
		} else {
			assert.Equal(t, "forceSync=false", got.query)
		}
	}
}
