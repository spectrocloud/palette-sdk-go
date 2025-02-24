package client

import (
	clientv1 "github.com/spectrocloud/palette-sdk-go/api/client/v1"
	"github.com/spectrocloud/palette-sdk-go/api/models"
)

// UpdateSessionTimeout update session timout for platform
func (h *V1Client) UpdateSessionTimeout(tenantUID string, body *models.V1AuthTokenSettings) error {
	params := clientv1.NewV1TenantUIDAuthTokenSettingsUpdateParamsWithContext(h.ctx).WithTenantUID(tenantUID).WithBody(body)
	_, err := h.Client.V1TenantUIDAuthTokenSettingsUpdate(params)
	return err
}

// GetSessionTimeout get session timout for platform
func (h *V1Client) GetSessionTimeout(tenantUID string) (*models.V1AuthTokenSettings, error) {
	params := clientv1.NewV1TenantUIDAuthTokenSettingsGetParamsWithContext(h.ctx).WithTenantUID(tenantUID)
	resp, err := h.Client.V1TenantUIDAuthTokenSettingsGet(params)
	return resp.Payload, err
}

// UpdatePlatformClusterUpgradeSetting update clusters agent upgrade setting for tenant
func (h *V1Client) UpdatePlatformClusterUpgradeSetting(body *models.V1ClusterUpgradeSettingsEntity) error {
	params := clientv1.NewV1SpectroClustersUpgradeSettingsParamsWithContext(h.ctx).WithBody(body)
	_, err := h.Client.V1SpectroClustersUpgradeSettings(params)
	return err
}

// GetPlatformClustersUpgradeSetting get clusters agent upgrade setting for tenant
func (h *V1Client) GetPlatformClustersUpgradeSetting() (*models.V1ClusterUpgradeSettingsEntity, error) {
	params := clientv1.NewV1SpectroClustersUpgradeSettingsGetParamsWithContext(h.ctx)
	resp, err := h.Client.V1SpectroClustersUpgradeSettingsGet(params)
	return resp.Payload, err
}

// UpdateClusterAutoRemediationForTenant update cluster auto remediation for platform
func (h *V1Client) UpdateClusterAutoRemediationForTenant(tenantUID string, body *models.V1NodesAutoRemediationSettings) error {
	params := clientv1.NewV1TenantClustersNodesAutoRemediationSettingUpdateParamsWithContext(h.ctx).
		WithTenantUID(tenantUID).WithBody(body)
	_, err := h.Client.V1TenantClustersNodesAutoRemediationSettingUpdate(params)
	return err
}

// GetClusterAutoRemediationForTenant get cluster auto remediation for platform
func (h *V1Client) GetClusterAutoRemediationForTenant(tenantUID string) (*models.V1TenantClusterSettings, error) {
	params := clientv1.NewV1TenantClusterSettingsGetParamsWithContext(h.ctx).WithTenantUID(tenantUID)
	resp, err := h.Client.V1TenantClusterSettingsGet(params)
	return resp.Payload, err
}

// UpdateClusterAutoRemediationForProject update cluster auto remediation for project
func (h *V1Client) UpdateClusterAutoRemediationForProject(projectUID string, body *models.V1NodesAutoRemediationSettings) error {
	params := clientv1.NewV1ProjectClustersNodesAutoRemediationSettingUpdateParamsWithContext(h.ctx).
		WithUID(projectUID).WithBody(body)
	_, err := h.Client.V1ProjectClustersNodesAutoRemediationSettingUpdate(params)
	return err
}

// GetClusterAutoRemediationForProject get cluster auto remediation for project
func (h *V1Client) GetClusterAutoRemediationForProject(projectUID string) (*models.V1ProjectClusterSettings, error) {
	params := clientv1.NewV1ProjectClusterSettingsGetParamsWithContext(h.ctx).WithUID(projectUID)
	resp, err := h.Client.V1ProjectClusterSettingsGet(params)
	return resp.Payload, err
}

// UpdateLoginBanner update login banner details for platform
func (h *V1Client) UpdateLoginBanner(tenantUID string, body *models.V1LoginBannerSettings) error {
	params := clientv1.NewV1TenantUIDLoginBannerUpdateParamsWithContext(h.ctx).WithTenantUID(tenantUID).WithBody(body)
	_, err := h.Client.V1TenantUIDLoginBannerUpdate(params)
	return err
}

// GetLoginBanner get login banner details for platform
func (h *V1Client) GetLoginBanner(tenantUID string) (*models.V1LoginBannerSettings, error) {
	params := clientv1.NewV1TenantUIDLoginBannerGetParamsWithContext(h.ctx).WithTenantUID(tenantUID)
	resp, err := h.Client.V1TenantUIDLoginBannerGet(params)
	return resp.Payload, err
}
