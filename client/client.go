package client

import (
	"context"
	"crypto/tls"
	"net/http"

	openapiclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/spectrocloud/palette-api-go/apiutil/transport"
	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
)

type V1Client struct {
	Client clientV1.ClientService

	ctx                context.Context
	apikey             string
	jwt                string
	username           string
	password           string
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

func WithJWT(jwt string) func(*V1Client) {
	return func(v *V1Client) {
		v.jwt = jwt
	}
}

func WithUsername(username string) func(*V1Client) {
	return func(v *V1Client) {
		v.username = username
	}
}

func WithPassword(password string) func(*V1Client) {
	return func(v *V1Client) {
		v.password = password
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
	opts := []func(*V1Client){
		WithHubbleURI(h.hubbleUri),
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
	if h.projectUid != "" {
		opts = append(opts, WithScopeProject(h.projectUid))
	}
	if h.transportDebug {
		opts = append(opts, WithTransportDebug())
	}
	return New(opts...)
}

func (h *V1Client) getTransport() (t *transport.Runtime) {
	if h.username != "" && h.password != "" {
		if err := h.authenticate(); err != nil {
			return nil
		}
	}
	if h.apikey != "" {
		t = h.apiKeyTransport()
	} else if h.jwt != "" {
		t = h.jwtTransport()
	} else {
		t = h.baseTransport()
	}
	return
}

func (h *V1Client) apiKeyTransport() *transport.Runtime {
	httpTransport := h.baseTransport()
	httpTransport.DefaultAuthentication = openapiclient.APIKeyAuth(authApiKey, authTokenInput, h.apikey)
	return httpTransport
}

func (h *V1Client) jwtTransport() *transport.Runtime {
	httpTransport := h.baseTransport()
	httpTransport.DefaultAuthentication = openapiclient.APIKeyAuth(authJwt, authTokenInput, h.jwt)
	return httpTransport
}

func (h *V1Client) authenticate() error {
	httpTransport := h.baseTransport()
	c := clientV1.New(httpTransport, strfmt.Default)

	params := &clientV1.V1AuthenticateParams{
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
