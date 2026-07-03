package client

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestResolveCloudStackZoneID(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/v1/cloudaccounts/apache-cloudstack/account-1/properties/zones", r.URL.Path)
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		_, _ = rw.Write([]byte(`{
			"zones": [
				{"id": "zone-other", "name": "other"},
				{"id": "zone-spectro", "name": "spectro"}
			]
		}`))
	}))
	t.Cleanup(server.Close)

	c := New(
		WithPaletteURI(strings.TrimPrefix(server.URL, "http://")),
		WithAPIKey("test-key"),
		WithSchemes([]string{"http"}),
	)

	id, err := c.ResolveCloudStackZoneID("account-1", "spectro")
	require.NoError(t, err)
	assert.Equal(t, "zone-spectro", id)
}

func TestResolveCloudStackNetworkID(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/v1/cloudaccounts/apache-cloudstack/account-1/properties/networks", r.URL.Path)
		assert.Equal(t, "zone-1", r.URL.Query().Get("zone"))
		assert.Equal(t, "project-1", r.URL.Query().Get("projectId"))
		assert.Equal(t, "vpc-1", r.URL.Query().Get("vpcId"))
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		_, _ = rw.Write([]byte(`{
			"networks": [
				{"id": "net-other", "name": "other"},
				{"id": "net-spectro", "name": "spectro"}
			]
		}`))
	}))
	t.Cleanup(server.Close)

	c := New(
		WithPaletteURI(strings.TrimPrefix(server.URL, "http://")),
		WithAPIKey("test-key"),
		WithSchemes([]string{"http"}),
	)

	id, err := c.ResolveCloudStackNetworkID("account-1", "spectro", "zone-1", "project-1", "vpc-1")
	require.NoError(t, err)
	assert.Equal(t, "net-spectro", id)
}

func TestResolveCloudStackNetworkIDNotFound(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, _ *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		_, _ = rw.Write([]byte(`{"networks": [{"id": "net-other", "name": "other"}]}`))
	}))
	t.Cleanup(server.Close)

	c := New(
		WithPaletteURI(strings.TrimPrefix(server.URL, "http://")),
		WithAPIKey("test-key"),
		WithSchemes([]string{"http"}),
	)

	_, err := c.ResolveCloudStackNetworkID("account-1", "spectro", "zone-1", "", "")
	require.Error(t, err)
	assert.Contains(t, err.Error(), `no CloudStack network found with name "spectro"`)
}

func TestResolveCloudStackNetworkIDDuplicate(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, _ *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		_, _ = rw.Write([]byte(`{
			"networks": [
				{"id": "net-1", "name": "spectro"},
				{"id": "net-2", "name": "spectro"}
			]
		}`))
	}))
	t.Cleanup(server.Close)

	c := New(
		WithPaletteURI(strings.TrimPrefix(server.URL, "http://")),
		WithAPIKey("test-key"),
		WithSchemes([]string{"http"}),
	)

	_, err := c.ResolveCloudStackNetworkID("account-1", "spectro", "zone-1", "", "")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "multiple CloudStack networks found")
}
