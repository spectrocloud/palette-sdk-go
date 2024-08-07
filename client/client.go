// Package client provides a client for the Spectro Cloud API.
package client

import (
	"context"
	"crypto/tls"
	"net/http"

	openapiclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/spectrocloud/palette-sdk-go/api/apiutil/transport"
	clientv1 "github.com/spectrocloud/palette-sdk-go/api/client/v1"
	"github.com/spectrocloud/palette-sdk-go/api/models"
)

// V1Client is a client for the Spectro Cloud API.
type V1Client struct {
	Client clientv1.ClientService

	ctx                context.Context
	apikey             string
	jwt                string
	username           string
	password           string
	paletteURI         string
	projectUID         string
	schemes            []string
	insecureSkipVerify bool
	transportDebug     bool
	retryAttempts      int
}

// New creates a new V1Client.
func New(options ...func(*V1Client)) *V1Client {
	client := &V1Client{
		ctx:           context.Background(),
		retryAttempts: 0,
		schemes:       []string{"https"},
	}
	for _, o := range options {
		o(client)
	}
	client.Client = clientv1.New(client.getTransport(), strfmt.Default)
	return client
}

// WithAPIKey sets the API key for the client.
func WithAPIKey(apiKey string) func(*V1Client) {
	return func(v *V1Client) {
		v.apikey = apiKey
	}
}

// WithJWT sets the JWT for the client.
func WithJWT(jwt string) func(*V1Client) {
	return func(v *V1Client) {
		v.jwt = jwt
	}
}

// WithUsername sets the username for the client.
func WithUsername(username string) func(*V1Client) {
	return func(v *V1Client) {
		v.username = username
	}
}

// WithPassword sets the password for the client.
func WithPassword(password string) func(*V1Client) {
	return func(v *V1Client) {
		v.password = password
	}
}

// WithPaletteURI sets the Palette URI for the client.
func WithPaletteURI(paletteURI string) func(*V1Client) {
	return func(v *V1Client) {
		v.paletteURI = paletteURI
	}
}

// WithInsecureSkipVerify sets insecureSkipVerify for the client.
func WithInsecureSkipVerify(insecureSkipVerify bool) func(*V1Client) {
	return func(v *V1Client) {
		v.insecureSkipVerify = insecureSkipVerify
	}
}

// WithScopeProject sets the project UID for the client.
func WithScopeProject(projectUID string) func(*V1Client) {
	return func(v *V1Client) {
		v.projectUID = projectUID
		v.ctx = ContextForScope("project", projectUID)
	}
}

// WithScopeTenant sets the tenant scope for the client.
func WithScopeTenant() func(*V1Client) {
	return func(v *V1Client) {
		v.ctx = ContextForScope("tenant", "")
	}
}

// WithRetries sets the number of retries for the client.
func WithRetries(retries int) func(*V1Client) {
	return func(v *V1Client) {
		v.retryAttempts = retries
	}
}

// WithSchemes sets the schemes for the client.
func WithSchemes(schemes []string) func(*V1Client) {
	return func(v *V1Client) {
		v.schemes = schemes
	}
}

// WithTransportDebug sets the client's HTTP transport debug flag.
func WithTransportDebug() func(*V1Client) {
	return func(v *V1Client) {
		v.transportDebug = true
	}
}

// ContextForScope returns a context with the given scope and optional project UID.
func ContextForScope(scope, projectUID string) context.Context {
	ctx := context.Background()
	if scope == "project" {
		ctx = context.WithValue(ctx, transport.CUSTOM_HEADERS, transport.Values{
			HeaderMap: map[string]string{
				"ProjectUid": projectUID,
			}},
		)
	}
	return ctx
}

// Clone creates a new V1Client with the same configuration as the original.
func (h *V1Client) Clone() *V1Client {
	opts := []func(*V1Client){
		WithPaletteURI(h.paletteURI),
		WithInsecureSkipVerify(h.insecureSkipVerify),
		WithRetries(h.retryAttempts),
		WithSchemes(h.schemes),
		WithScopeTenant(),
	}
	if h.apikey != "" {
		opts = append(opts, WithAPIKey(h.apikey))
	}
	if h.jwt != "" {
		opts = append(opts, WithJWT(h.jwt))
	}
	if h.username != "" && h.password != "" {
		opts = append(opts, WithUsername(h.username), WithPassword(h.password))
	}
	if h.projectUID != "" {
		opts = append(opts, WithScopeProject(h.projectUID))
	}
	if h.transportDebug {
		opts = append(opts, WithTransportDebug())
	}
	return New(opts...)
}

func (h *V1Client) getTransport() *transport.Runtime {
	if h.username != "" && h.password != "" {
		if err := h.handleBasicAuth(); err != nil {
			return nil
		}
	}
	var t *transport.Runtime
	if h.apikey != "" {
		t = h.apiKeyTransport()
	} else if h.jwt != "" {
		t = h.jwtTransport()
	} else {
		t = h.baseTransport()
	}
	return t
}

func (h *V1Client) apiKeyTransport() *transport.Runtime {
	httpTransport := h.baseTransport()
	httpTransport.DefaultAuthentication = openapiclient.APIKeyAuth(authAPIKey, authTokenInput, h.apikey)
	return httpTransport
}

func (h *V1Client) jwtTransport() *transport.Runtime {
	httpTransport := h.baseTransport()
	httpTransport.DefaultAuthentication = openapiclient.APIKeyAuth(authJwt, authTokenInput, h.jwt)
	return httpTransport
}

func (h *V1Client) handleBasicAuth() error {
	httpTransport := h.baseTransport()
	c := clientv1.New(httpTransport, strfmt.Default)

	params := &clientv1.V1AuthenticateParams{
		Body: &models.V1AuthLogin{
			EmailID:  h.username,
			Password: strfmt.Password(h.password),
		},
	}
	resp, err := c.V1Authenticate(params)
	if err != nil {
		return err
	}
	h.jwt = resp.Payload.Authorization

	return nil
}

func (h *V1Client) baseTransport() *transport.Runtime {
	httpTransport := transport.NewWithClient(h.paletteURI, "", h.schemes, h.httpClient())
	httpTransport.RetryAttempts = h.retryAttempts
	httpTransport.Debug = h.transportDebug
	return httpTransport
}

func (h *V1Client) httpClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: h.insecureSkipVerify, // #nosec G402 - InsecureSkipVerify is enabled via user input
			},
		},
	}
}

// Validate validates the configuration for a project-scoped client.
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

// ValidateTenantAdmin validates the configuration for a tenant-scoped client.
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
