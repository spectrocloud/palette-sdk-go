package client

import (
	clientv1 "github.com/spectrocloud/palette-sdk-go/api/client/v1"
	"github.com/spectrocloud/palette-sdk-go/api/models"
	"github.com/spectrocloud/palette-sdk-go/client/apiutil"
)

// CreateCloudAccountGcp creates a new GCP cloud account.
func (h *V1Client) CreateCloudAccountGcp(account *models.V1GcpAccountEntity) (string, error) {
	params := clientv1.NewV1CloudAccountsGcpCreateParamsWithContext(h.ctx).
		WithBody(account)
	resp, err := h.Client.V1CloudAccountsGcpCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

// UpdateCloudAccountGcp updates an existing GCP cloud account.
func (h *V1Client) UpdateCloudAccountGcp(account *models.V1GcpAccountEntity) error {
	params := clientv1.NewV1CloudAccountsGcpUpdateParamsWithContext(h.ctx).
		WithUID(account.Metadata.UID).
		WithBody(account)
	_, err := h.Client.V1CloudAccountsGcpUpdate(params)
	return err
}

// DeleteCloudAccountGcp deletes an existing GCP cloud account.
func (h *V1Client) DeleteCloudAccountGcp(uid string) error {
	params := clientv1.NewV1CloudAccountsGcpDeleteParamsWithContext(h.ctx).
		WithUID(uid)
	_, err := h.Client.V1CloudAccountsGcpDelete(params)
	return err
}

// GetCloudAccountGcp retrieves an existing GCP cloud account.
func (h *V1Client) GetCloudAccountGcp(uid string) (*models.V1GcpAccount, error) {
	params := clientv1.NewV1CloudAccountsGcpGetParamsWithContext(h.ctx).
		WithUID(uid)
	resp, err := h.Client.V1CloudAccountsGcpGet(params)
	if apiutil.Is404(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// GetCloudAccountsGcp retrieves all GCP cloud accounts.
func (h *V1Client) GetCloudAccountsGcp() ([]*models.V1GcpAccount, error) {
	params := clientv1.NewV1CloudAccountsGcpListParamsWithContext(h.ctx).
		WithLimit(apiutil.Ptr(int64(0)))
	resp, err := h.Client.V1CloudAccountsGcpList(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload.Items, nil
}
