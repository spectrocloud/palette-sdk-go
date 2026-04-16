package platform

import "github.com/spectrocloud/palette-sdk-go/api/models"

func (s *service) GetSessionTimeout(tenantUID string) (*models.V1AuthTokenSettings, error) {
	if tenantUID == "" {
		return nil, errTenantUIDRequired
	}
	return s.client.GetSessionTimeout(tenantUID)
}

func (s *service) UpdateSessionTimeout(tenantUID string, body *models.V1AuthTokenSettings) error {
	if tenantUID == "" {
		return errTenantUIDRequired
	}
	return s.client.UpdateSessionTimeout(tenantUID, body)
}

func (s *service) GetLoginBanner(tenantUID string) (*models.V1LoginBannerSettings, error) {
	if tenantUID == "" {
		return nil, errTenantUIDRequired
	}
	return s.client.GetLoginBanner(tenantUID)
}

func (s *service) UpdateLoginBanner(tenantUID string, body *models.V1LoginBannerSettings) error {
	if tenantUID == "" {
		return errTenantUIDRequired
	}
	return s.client.UpdateLoginBanner(tenantUID, body)
}

func (s *service) GetFIPSPreference(tenantUID string) (*models.V1FipsSettings, error) {
	if tenantUID == "" {
		return nil, errTenantUIDRequired
	}
	return s.client.GetFIPSPreference(tenantUID)
}

func (s *service) UpdateFIPSPreference(tenantUID string, body *models.V1FipsSettings) error {
	if tenantUID == "" {
		return errTenantUIDRequired
	}
	return s.client.UpdateFIPSPreference(tenantUID, body)
}

func (s *service) GetPasswordPolicy(tenantUID string) (*models.V1TenantPasswordPolicyEntity, error) {
	if tenantUID == "" {
		return nil, errTenantUIDRequired
	}
	return s.client.GetPasswordPolicy(tenantUID)
}

func (s *service) UpdatePasswordPolicy(tenantUID string, body *models.V1TenantPasswordPolicyEntity) error {
	if tenantUID == "" {
		return errTenantUIDRequired
	}
	return s.client.UpdatePasswordPolicy(tenantUID, body)
}

func (s *service) GetDeveloperSetting(tenantUID string) (*models.V1DeveloperCredit, error) {
	if tenantUID == "" {
		return nil, errTenantUIDRequired
	}
	return s.client.GetDeveloperSetting(tenantUID)
}

func (s *service) UpdateDeveloperSetting(tenantUID string, body *models.V1DeveloperCredit) error {
	if tenantUID == "" {
		return errTenantUIDRequired
	}
	return s.client.UpdateDeveloperSetting(tenantUID, body)
}
