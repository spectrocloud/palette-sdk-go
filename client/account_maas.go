package client

import (
	clientv1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
	"github.com/spectrocloud/palette-sdk-go/client/apiutil"
)

func toV1OverlordsUIDMaasAccountValidateBody(account *models.V1MaasAccount) clientv1.V1OverlordsUIDMaasAccountValidateBody {
	return clientv1.V1OverlordsUIDMaasAccountValidateBody{
		Account: account.Spec,
	}
}

// CreateCloudAccountMaas creates a new MAAS cloud account.
func (h *V1Client) CreateCloudAccountMaas(account *models.V1MaasAccount) (string, error) {
	// validate account
	if err := h.validateCloudAccountMaas(account); err != nil {
		return "", err
	}

	params := clientv1.NewV1CloudAccountsMaasCreateParamsWithContext(h.ctx).
		WithBody(account)
	resp, err := h.Client.V1CloudAccountsMaasCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

func (h *V1Client) validateCloudAccountMaas(account *models.V1MaasAccount) error {
	// check PCG
	pcgID := account.Metadata.Annotations[OverlordUID]
	if err := h.CheckPCG(pcgID); err != nil {
		return err
	}

	// validate account
	params := clientv1.NewV1OverlordsUIDMaasAccountValidateParamsWithContext(h.ctx).
		WithUID(pcgID).
		WithBody(toV1OverlordsUIDMaasAccountValidateBody(account))

	_, err := h.Client.V1OverlordsUIDMaasAccountValidate(params)
	return err
}

// UpdateCloudAccountMaas updates an existing MAAS cloud account.
func (h *V1Client) UpdateCloudAccountMaas(account *models.V1MaasAccount) error {
	// validate account
	if err := h.validateCloudAccountMaas(account); err != nil {
		return err
	}

	params := clientv1.NewV1CloudAccountsMaasUpdateParamsWithContext(h.ctx).
		WithUID(account.Metadata.UID).
		WithBody(account)

	_, err := h.Client.V1CloudAccountsMaasUpdate(params)
	return err
}

// DeleteCloudAccountMaas deletes an existing MAAS cloud account.
func (h *V1Client) DeleteCloudAccountMaas(uid string) error {
	params := clientv1.NewV1CloudAccountsMaasDeleteParamsWithContext(h.ctx).
		WithUID(uid)
	_, err := h.Client.V1CloudAccountsMaasDelete(params)
	return err
}

// GetCloudAccountMaas retrieves an existing MAAS cloud account.
func (h *V1Client) GetCloudAccountMaas(uid string) (*models.V1MaasAccount, error) {
	params := clientv1.NewV1CloudAccountsMaasGetParamsWithContext(h.ctx).
		WithUID(uid)
	resp, err := h.Client.V1CloudAccountsMaasGet(params)
	if apiutil.Is404(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// GetCloudAccountsMaas retrieves all MAAS cloud accounts.
func (h *V1Client) GetCloudAccountsMaas() ([]*models.V1MaasAccount, error) {
	params := clientv1.NewV1CloudAccountsMaasListParamsWithContext(h.ctx).
		WithLimit(apiutil.Ptr(int64(0)))
	response, err := h.Client.V1CloudAccountsMaasList(params)
	if err != nil {
		return nil, err
	}
	return response.Payload.Items, nil
}
