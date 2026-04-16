package platform

import "github.com/spectrocloud/palette-sdk-go/api/models"

func (s *service) GetOIDC(tenantUID string) (*models.V1TenantOidcClientSpec, error) {
	if tenantUID == "" {
		return nil, errTenantUIDRequired
	}
	return s.client.GetOIDC(tenantUID)
}

func (s *service) UpdateOIDC(tenantUID string, body *models.V1TenantOidcClientSpec) error {
	if tenantUID == "" {
		return errTenantUIDRequired
	}
	return s.client.UpdateOIDC(tenantUID, body)
}

func (s *service) GetSAML(tenantUID string) (*models.V1TenantSamlSpec, error) {
	if tenantUID == "" {
		return nil, errTenantUIDRequired
	}
	return s.client.GetSAML(tenantUID)
}

func (s *service) UpdateSAML(tenantUID string, body *models.V1TenantSamlRequestSpec) error {
	if tenantUID == "" {
		return errTenantUIDRequired
	}
	return s.client.UpdateSAML(tenantUID, body)
}

func (s *service) GetDomains(tenantUID string) (*models.V1TenantDomains, error) {
	if tenantUID == "" {
		return nil, errTenantUIDRequired
	}
	return s.client.GetDomains(tenantUID)
}

func (s *service) UpdateDomain(tenantUID string, body *models.V1TenantDomains) error {
	if tenantUID == "" {
		return errTenantUIDRequired
	}
	return s.client.UpdateDomain(tenantUID, body)
}

func (s *service) GetProviders(tenantUID string) (*models.V1TenantSsoAuthProvidersEntity, error) {
	if tenantUID == "" {
		return nil, errTenantUIDRequired
	}
	return s.client.GetProviders(tenantUID)
}

func (s *service) UpdateProviders(tenantUID string, body *models.V1TenantSsoAuthProvidersEntity) error {
	if tenantUID == "" {
		return errTenantUIDRequired
	}
	return s.client.UpdateProviders(tenantUID, body)
}
