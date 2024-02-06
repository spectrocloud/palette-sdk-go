package client

import (
	"context"

	openapiclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/spectrocloud/palette-api-go/apiutil/transport"
	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
)

const (
	authTokenInput string = "header"
	authApiKey     string = "ApiKey"
)

type V1Client struct {
	Ctx            context.Context
	apikey         string
	hubbleUri      string
	schemes        []string
	transportDebug bool
	RetryAttempts  int
}

func New(options ...func(*V1Client)) *V1Client {
	client := &V1Client{
		Ctx:           context.Background(),
		RetryAttempts: 0,
		schemes:       []string{"https"},
	}
	for _, o := range options {
		o(client)
	}
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

func WithProjectUID(projectUid string) func(*V1Client) {
	return func(v *V1Client) {
		v.Ctx = ContextWithProject(v.Ctx, projectUid)
	}
}

func WithRetries(retries int) func(*V1Client) {
	return func(v *V1Client) {
		v.RetryAttempts = retries
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
	httpTransport := transport.New(h.hubbleUri, "", h.schemes)
	httpTransport.RetryAttempts = h.RetryAttempts
	httpTransport.Debug = h.transportDebug
	return httpTransport
}

func (h *V1Client) GetClient() clientV1.ClientService {
	return clientV1.New(h.getTransport(), strfmt.Default)
}

func (h *V1Client) Validate() error {
	// API key can only be validated by making an API call
	if h.apikey != "" {
		c := h.GetClient()
		_, err := c.V1PacksSummaryList(nil)
		if err != nil {
			return err
		}
	}
	return nil
}

func (h *V1Client) ValidateTenantAdmin() error {
	// API key can only be validated by making an API call
	if h.apikey != "" {
		c := h.GetClient()
		_, err := c.V1UsersList(nil)
		if err != nil {
			return err
		}
	}
	return nil
}

func ContextWithProject(c context.Context, projectUid string) context.Context {
	return context.WithValue(c, transport.CUSTOM_HEADERS, transport.Values{
		HeaderMap: map[string]string{
			"ProjectUid": projectUid,
		}})
}
