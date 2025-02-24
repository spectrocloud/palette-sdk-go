package client

import (
	clientv1 "github.com/spectrocloud/palette-sdk-go/api/client/v1"
	"github.com/spectrocloud/palette-sdk-go/api/models"
)

// UpdateDeveloperSetting update a Developer Settings
func (h *V1Client) UpdateDeveloperSetting(tenantUID string, body *models.V1DeveloperCredit) error {
	params := clientv1.NewV1TenantDeveloperCreditUpdateParamsWithContext(h.ctx).WithTenantUID(tenantUID).WithBody(body)
	_, err := h.Client.V1TenantDeveloperCreditUpdate(params)
	return err
}

// GetDeveloperSetting retrieves an existing  a Developer Settings by tenant UID
func (h *V1Client) GetDeveloperSetting(tenantUID string) (*models.V1DeveloperCredit, error) {
	params := clientv1.NewV1TenantDeveloperCreditGetParamsWithContext(h.ctx).WithTenantUID(tenantUID)
	resp, err := h.Client.V1TenantDeveloperCreditGet(params)
	return resp.Payload, err
}
