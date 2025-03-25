package client

import (
	clientv1 "github.com/spectrocloud/palette-sdk-go/api/client/v1"
	"github.com/spectrocloud/palette-sdk-go/api/models"
)

// UpdateOIDC update OIDC SSO Configurations.
func (h *V1Client) UpdateOIDC(tenantUID string, body *models.V1TenantOidcClientSpec) error {
	params := clientv1.NewV1TenantUIDOidcConfigUpdateParamsWithContext(h.ctx).
		WithBody(body).WithTenantUID(tenantUID)
	_, err := h.Client.V1TenantUIDOidcConfigUpdate(params)
	return err
}

// GetOIDC get OIDC SSO Configurations.
func (h *V1Client) GetOIDC(tenantUID string) (*models.V1TenantOidcClientSpec, error) {
	params := clientv1.NewV1TenantUIDOidcConfigGetParamsWithContext(h.ctx).WithTenantUID(tenantUID)
	resp, err := h.Client.V1TenantUIDOidcConfigGet(params)
	return resp.Payload, err
}

// UpdateSAML update SAML SSO Configurations.
func (h *V1Client) UpdateSAML(tenantUID string, body *models.V1TenantSamlRequestSpec) error {
	params := clientv1.NewV1TenantUIDSamlConfigUpdateParamsWithContext(h.ctx).
		WithBody(body).WithTenantUID(tenantUID)
	_, err := h.Client.V1TenantUIDSamlConfigUpdate(params)
	return err
}

// GetSAML get SAML SSO Configurations.
func (h *V1Client) GetSAML(tenantUID string) (*models.V1TenantSamlSpec, error) {
	params := clientv1.NewV1TenantUIDSamlConfigSpecGetParamsWithContext(h.ctx).WithTenantUID(tenantUID)
	resp, err := h.Client.V1TenantUIDSamlConfigSpecGet(params)
	return resp.Payload, err
}

// UpdateDomain update domains for SSO .
func (h *V1Client) UpdateDomain(tenantUID string, body *models.V1TenantDomains) error {
	params := clientv1.NewV1TenantUIDDomainsUpdateParamsWithContext(h.ctx).
		WithBody(body).WithTenantUID(tenantUID)
	_, err := h.Client.V1TenantUIDDomainsUpdate(params)
	return err
}

// GetDomains get SAML SSO Configurations.
func (h *V1Client) GetDomains(tenantUID string) (*models.V1TenantDomains, error) {
	params := clientv1.NewV1TenantUIDDomainsGetParamsWithContext(h.ctx).WithTenantUID(tenantUID)
	resp, err := h.Client.V1TenantUIDDomainsGet(params)
	return resp.Payload, err
}

// UpdateProviders update providers for SSO .
func (h *V1Client) UpdateProviders(tenantUID string, body *models.V1TenantSsoAuthProvidersEntity) error {
	params := clientv1.NewV1TenantUIDSsoAuthProvidersUpdateParamsWithContext(h.ctx).
		WithBody(body).WithTenantUID(tenantUID)
	_, err := h.Client.V1TenantUIDSsoAuthProvidersUpdate(params)
	return err
}

// GetProviders get SAML SSO Configurations.
func (h *V1Client) GetProviders(tenantUID string) (*models.V1TenantSsoAuthProvidersEntity, error) {
	params := clientv1.NewV1TenantUIDSsoAuthProvidersGetParamsWithContext(h.ctx).WithTenantUID(tenantUID)
	resp, err := h.Client.V1TenantUIDSsoAuthProvidersGet(params)
	return resp.Payload, err
}
