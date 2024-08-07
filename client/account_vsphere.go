package client

import (
	clientv1 "github.com/spectrocloud/palette-sdk-go/api/client/v1"
	"github.com/spectrocloud/palette-sdk-go/api/models"
	"github.com/spectrocloud/palette-sdk-go/client/apiutil"
)

// convert V1VsphereAccount to V1OverlordVsphereAccountEntity
func toV1OverlordsUIDVsphereAccountValidateBody(account *models.V1VsphereAccount) clientv1.V1OverlordsUIDVsphereAccountValidateBody {
	return clientv1.V1OverlordsUIDVsphereAccountValidateBody{
		Account: account.Spec,
	}
}

// CreateCloudAccountVsphere creates a new vSphere cloud account.
func (h *V1Client) CreateCloudAccountVsphere(account *models.V1VsphereAccount) (string, error) {
	// validate account
	if err := h.validateCloudAccountVsphere(account); err != nil {
		return "", err
	}

	// create account
	params := clientv1.NewV1CloudAccountsVsphereCreateParamsWithContext(h.ctx).
		WithBody(account)
	resp, err := h.Client.V1CloudAccountsVsphereCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

func (h *V1Client) validateCloudAccountVsphere(account *models.V1VsphereAccount) error {
	// check PCG
	pcgID := account.Metadata.Annotations[OverlordUID]
	if err := h.CheckPCG(pcgID); err != nil {
		return err
	}

	// validate account
	params := clientv1.NewV1OverlordsUIDVsphereAccountValidateParamsWithContext(h.ctx).
		WithUID(pcgID).
		WithBody(toV1OverlordsUIDVsphereAccountValidateBody(account))

	_, err := h.Client.V1OverlordsUIDVsphereAccountValidate(params)
	return err
}

// UpdateCloudAccountVsphere updates an existing vSphere cloud account.
func (h *V1Client) UpdateCloudAccountVsphere(account *models.V1VsphereAccount) error {
	// validate account
	if err := h.validateCloudAccountVsphere(account); err != nil {
		return err
	}

	params := clientv1.NewV1CloudAccountsVsphereUpdateParamsWithContext(h.ctx).
		WithUID(account.Metadata.UID).
		WithBody(account)
	_, err := h.Client.V1CloudAccountsVsphereUpdate(params)
	return err
}

// DeleteCloudAccountVsphere deletes an existing vSphere cloud account.
func (h *V1Client) DeleteCloudAccountVsphere(uid string) error {
	params := clientv1.NewV1CloudAccountsVsphereDeleteParamsWithContext(h.ctx).
		WithUID(uid)
	_, err := h.Client.V1CloudAccountsVsphereDelete(params)
	return err
}

// GetCloudAccountVsphere retrieves an existing vSphere cloud account.
func (h *V1Client) GetCloudAccountVsphere(uid string) (*models.V1VsphereAccount, error) {
	params := clientv1.NewV1CloudAccountsVsphereGetParamsWithContext(h.ctx).
		WithUID(uid)
	resp, err := h.Client.V1CloudAccountsVsphereGet(params)
	if apiutil.Is404(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// GetCloudAccountsVsphere retrieves all vSphere cloud accounts.
func (h *V1Client) GetCloudAccountsVsphere() ([]*models.V1VsphereAccount, error) {
	params := clientv1.NewV1CloudAccountsVsphereListParamsWithContext(h.ctx).
		WithLimit(apiutil.Ptr(int64(0)))
	resp, err := h.Client.V1CloudAccountsVsphereList(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload.Items, nil
}
