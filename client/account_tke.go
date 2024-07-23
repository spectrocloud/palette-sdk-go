package client

import (
	clientv1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
	"github.com/spectrocloud/palette-sdk-go/client/apiutil"
)

// CreateCloudAccountTke creates a new TKE cloud account.
func (h *V1Client) CreateCloudAccountTke(account *models.V1TencentAccount) (string, error) {
	params := clientv1.NewV1CloudAccountsTencentCreateParamsWithContext(h.ctx).
		WithBody(account)
	resp, err := h.Client.V1CloudAccountsTencentCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

// UpdateCloudAccountTke updates an existing TKE cloud account.
func (h *V1Client) UpdateCloudAccountTke(account *models.V1TencentAccount) error {
	params := clientv1.NewV1CloudAccountsTencentUpdateParamsWithContext(h.ctx).
		WithUID(account.Metadata.UID).
		WithBody(account)
	_, err := h.Client.V1CloudAccountsTencentUpdate(params)
	return err
}

// DeleteCloudAccountTke deletes an existing TKE cloud account.
func (h *V1Client) DeleteCloudAccountTke(uid string) error {
	params := clientv1.NewV1CloudAccountsTencentDeleteParamsWithContext(h.ctx).
		WithUID(uid)
	_, err := h.Client.V1CloudAccountsTencentDelete(params)
	return err
}

// GetCloudAccountTke retrieves an existing TKE cloud account.
func (h *V1Client) GetCloudAccountTke(uid string) (*models.V1TencentAccount, error) {
	params := clientv1.NewV1CloudAccountsTencentGetParamsWithContext(h.ctx).
		WithUID(uid)
	resp, err := h.Client.V1CloudAccountsTencentGet(params)
	if apiutil.Is404(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// GetCloudAccountsTke retrieves all TKE cloud accounts.
func (h *V1Client) GetCloudAccountsTke() ([]*models.V1TencentAccount, error) {
	params := clientv1.NewV1CloudAccountsTencentListParamsWithContext(h.ctx).
		WithLimit(apiutil.Ptr(int64(0)))
	resp, err := h.Client.V1CloudAccountsTencentList(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload.Items, nil
}
