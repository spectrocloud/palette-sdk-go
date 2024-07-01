package client

import (
	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
	"github.com/spectrocloud/palette-sdk-go/client/apiutil"
)

// convert V1VsphereAccount to V1OverlordVsphereAccountEntity
func toV1OverlordsUIDVsphereAccountValidateBody(account *models.V1VsphereAccount) clientV1.V1OverlordsUIDVsphereAccountValidateBody {
	return clientV1.V1OverlordsUIDVsphereAccountValidateBody{
		Account: account.Spec,
	}
}

func (h *V1Client) CreateCloudAccountVsphere(account *models.V1VsphereAccount) (string, error) {
	// validate account
	if err := h.validateCloudAccountVsphere(account); err != nil {
		return "", err
	}

	// create account
	params := clientV1.NewV1CloudAccountsVsphereCreateParamsWithContext(h.ctx).
		WithBody(account)
	resp, err := h.Client.V1CloudAccountsVsphereCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

func (h *V1Client) validateCloudAccountVsphere(account *models.V1VsphereAccount) error {
	// check PCG
	pcgId := account.Metadata.Annotations[OverlordUID]
	if err := h.CheckPCG(pcgId); err != nil {
		return err
	}

	// validate account
	params := clientV1.NewV1OverlordsUIDVsphereAccountValidateParamsWithContext(h.ctx).
		WithUID(pcgId).
		WithBody(toV1OverlordsUIDVsphereAccountValidateBody(account))

	_, err := h.Client.V1OverlordsUIDVsphereAccountValidate(params)
	return err
}

func (h *V1Client) UpdateCloudAccountVsphere(account *models.V1VsphereAccount) error {
	// validate account
	if err := h.validateCloudAccountVsphere(account); err != nil {
		return err
	}

	params := clientV1.NewV1CloudAccountsVsphereUpdateParamsWithContext(h.ctx).
		WithUID(account.Metadata.UID).
		WithBody(account)
	_, err := h.Client.V1CloudAccountsVsphereUpdate(params)
	return err
}

func (h *V1Client) DeleteCloudAccountVsphere(uid string) error {
	params := clientV1.NewV1CloudAccountsVsphereDeleteParamsWithContext(h.ctx).
		WithUID(uid)
	_, err := h.Client.V1CloudAccountsVsphereDelete(params)
	return err
}

func (h *V1Client) GetCloudAccountVsphere(uid string) (*models.V1VsphereAccount, error) {
	params := clientV1.NewV1CloudAccountsVsphereGetParamsWithContext(h.ctx).
		WithUID(uid)
	resp, err := h.Client.V1CloudAccountsVsphereGet(params)
	if apiutil.Is404(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

func (h *V1Client) GetCloudAccountsVsphere() ([]*models.V1VsphereAccount, error) {
	params := clientV1.NewV1CloudAccountsVsphereListParamsWithContext(h.ctx).
		WithLimit(apiutil.Ptr(int64(0)))
	resp, err := h.Client.V1CloudAccountsVsphereList(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload.Items, nil
}
