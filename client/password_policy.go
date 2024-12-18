package client

import (
	clientv1 "github.com/spectrocloud/palette-sdk-go/api/client/v1"
	"github.com/spectrocloud/palette-sdk-go/api/models"
)

// UpdatePasswordPolicy updates a Password Policy.
func (h *V1Client) UpdatePasswordPolicy(tenantUID string, body *models.V1TenantPasswordPolicyEntity) error {
	params := clientv1.NewV1TenantUIDPasswordPolicyUpdateParamsWithContext(h.ctx).WithTenantUID(tenantUID).WithBody(body)
	_, err := h.Client.V1TenantUIDPasswordPolicyUpdate(params)
	return err
}
