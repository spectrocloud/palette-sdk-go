package client

import (
	clientv1 "github.com/spectrocloud/palette-sdk-go/api/client/version1"
	"github.com/spectrocloud/palette-sdk-go/api/models"
)

// GetRateConfig retrieves the rate config for public and private cloud by tenant UID.
func (h *V1Client) GetRateConfig(tenantUID string) (*models.V1RateConfig, error) {
	params := clientv1.NewV1RateConfigGetParamsWithContext(h.ctx).WithTenantUID(tenantUID)
	resp, err := h.Client.V1RateConfigGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// UpdateRateConfig updates the rate config for public and private cloud for a tenant.
func (h *V1Client) UpdateRateConfig(tenantUID string, body *models.V1RateConfig) error {
	params := clientv1.NewV1RateConfigUpdateParamsWithContext(h.ctx).WithTenantUID(tenantUID).WithBody(body)
	_, err := h.Client.V1RateConfigUpdate(params)
	return err
}
