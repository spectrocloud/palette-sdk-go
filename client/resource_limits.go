package client

import (
	clientv1 "github.com/spectrocloud/palette-sdk-go/api/client/v1"
	"github.com/spectrocloud/palette-sdk-go/api/models"
)

// UpdateResourceLimits updates a Resource limits for tenant.
func (h *V1Client) UpdateResourceLimits(tenantUID string, body *models.V1TenantResourceLimitsEntity) error {
	params := clientv1.NewV1TenantResourceLimitsUpdateParamsWithContext(h.ctx).WithTenantUID(tenantUID).WithBody(body)
	_, err := h.Client.V1TenantResourceLimitsUpdate(params)
	return err
}

// GetResourceLimits retrieves an existing resource limits by tenant UID.
func (h *V1Client) GetResourceLimits(tenantUID string) (*models.V1TenantResourceLimits, error) {
	params := clientv1.NewV1TenantResourceLimitsGetParamsWithContext(h.ctx).WithTenantUID(tenantUID)
	resp, err := h.Client.V1TenantResourceLimitsGet(params)
	return resp.Payload, err
}
