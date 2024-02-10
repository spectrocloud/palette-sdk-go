package client

import (
	"context"
	"crypto/tls"
	"net/http"

	openapiclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/spectrocloud/palette-api-go/apiutil/transport"
	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
)

type V1Client struct {
	Client clientV1.ClientService

	ctx                context.Context
	apikey             string
	hubbleUri          string
	projectUid         string
	schemes            []string
	insecureSkipVerify bool
	transportDebug     bool
	retryAttempts      int
}

func New(options ...func(*V1Client)) *V1Client {
	client := &V1Client{
		ctx:           context.Background(),
		retryAttempts: 0,
		schemes:       []string{"https"},
	}
	for _, o := range options {
		o(client)
	}
	client.Client = clientV1.New(client.getTransport(), strfmt.Default)
	return client
}

func WithAPIKey(apiKey string) func(*V1Client) {
	return func(v *V1Client) {
		v.apikey = apiKey
	}
}

func WithHubbleURI(hubbleUri string) func(*V1Client) {
	return func(v *V1Client) {
		v.hubbleUri = hubbleUri
	}
}

func WithInsecureSkipVerify(insecureSkipVerify bool) func(*V1Client) {
	return func(v *V1Client) {
		v.insecureSkipVerify = insecureSkipVerify
	}
}

func WithScopeProject(projectUid string) func(*V1Client) {
	return func(v *V1Client) {
		v.projectUid = projectUid
		v.ctx = ContextForScope("project", projectUid)
	}
}

func WithScopeTenant() func(*V1Client) {
	return func(v *V1Client) {
		v.ctx = ContextForScope("tenant", "")
	}
}

func WithRetries(retries int) func(*V1Client) {
	return func(v *V1Client) {
		v.retryAttempts = retries
	}
}

func WithSchemes(schemes []string) func(*V1Client) {
	return func(v *V1Client) {
		v.schemes = schemes
	}
}

func WithTransportDebug() func(*V1Client) {
	return func(v *V1Client) {
		v.transportDebug = true
	}
}

func ContextForScope(scope, projectUid string) context.Context {
	ctx := context.Background()
	if scope == "project" {
		ctx = context.WithValue(ctx, transport.CUSTOM_HEADERS, transport.Values{
			HeaderMap: map[string]string{
				"ProjectUid": projectUid,
			}},
		)
	}
	return ctx
}

func (h *V1Client) Clone() *V1Client {
	c := New(
		WithAPIKey(h.apikey),
		WithHubbleURI(h.hubbleUri),
		WithInsecureSkipVerify(h.insecureSkipVerify),
		WithRetries(h.retryAttempts),
		WithSchemes(h.schemes),
		WithScopeTenant(),
	)
	if h.projectUid != "" {
		WithScopeProject(h.projectUid)(c)
	}
	if h.transportDebug {
		WithTransportDebug()(c)
	}
	return c
}

func (h *V1Client) getTransport() *transport.Runtime {
	var httpTransport *transport.Runtime
	if h.apikey != "" {
		httpTransport = h.authenticatedTransport()
	} else {
		httpTransport = h.baseTransport()
	}
	return httpTransport
}

func (h *V1Client) authenticatedTransport() *transport.Runtime {
	httpTransport := h.baseTransport()
	httpTransport.DefaultAuthentication = openapiclient.APIKeyAuth(authApiKey, authTokenInput, h.apikey)
	return httpTransport
}

func (h *V1Client) baseTransport() *transport.Runtime {
	httpTransport := transport.NewWithClient(h.hubbleUri, "", h.schemes, h.httpClient())
	httpTransport.RetryAttempts = h.retryAttempts
	httpTransport.Debug = h.transportDebug
	return httpTransport
}

func (h *V1Client) httpClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			TLSClientConfig: &tls.Config{
				Certificates:       []tls.Certificate{},
				InsecureSkipVerify: h.insecureSkipVerify,
			},
		},
	}
}

func (h *V1Client) Validate() error {
	// API key can only be validated by making an API call
	if h.apikey != "" {
		_, err := h.Client.V1ProjectsMetadata(nil)
		if err != nil {
			return err
		}
	}
	return nil
}

func (h *V1Client) ValidateTenantAdmin() error {
	// API key can only be validated by making an API call
	if h.apikey != "" {
		_, err := h.Client.V1UsersList(nil)
		if err != nil {
			return err
		}
	}
	return nil
}
