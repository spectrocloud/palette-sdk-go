package client

import (
	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
	"github.com/spectrocloud/palette-sdk-go/client/apiutil"
)

func toV1OverlordsUIDOpenStackAccountValidateBody(account *models.V1OpenStackAccount) clientV1.V1OverlordsUIDOpenStackAccountValidateBody {
	return clientV1.V1OverlordsUIDOpenStackAccountValidateBody{
		Account: account.Spec,
	}
}

func (h *V1Client) CreateCloudAccountOpenStack(account *models.V1OpenStackAccount) (string, error) {
	// validate account
	if err := h.validateCloudAccountOpenStack(account); err != nil {
		return "", err
	}

	params := clientV1.NewV1CloudAccountsOpenStackCreateParamsWithContext(h.ctx).
		WithBody(account)
	resp, err := h.Client.V1CloudAccountsOpenStackCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

func (h *V1Client) validateCloudAccountOpenStack(account *models.V1OpenStackAccount) error {
	// check PCG
	pcgId := account.Metadata.Annotations[OverlordUID]
	if err := h.CheckPCG(pcgId); err != nil {
		return err
	}

	// validate account
	params := clientV1.NewV1OverlordsUIDOpenStackAccountValidateParamsWithContext(h.ctx).
		WithUID(pcgId).
		WithBody(toV1OverlordsUIDOpenStackAccountValidateBody(account))

	_, err := h.Client.V1OverlordsUIDOpenStackAccountValidate(params)
	return err
}

func (h *V1Client) UpdateCloudAccountOpenStack(account *models.V1OpenStackAccount) error {
	if err := h.validateCloudAccountOpenStack(account); err != nil {
		return err
	}

	params := clientV1.NewV1CloudAccountsOpenStackUpdateParamsWithContext(h.ctx).
		WithUID(account.Metadata.UID).
		WithBody(account)
	_, err := h.Client.V1CloudAccountsOpenStackUpdate(params)
	return err
}

func (h *V1Client) DeleteCloudAccountOpenStack(uid string) error {
	params := clientV1.NewV1CloudAccountsOpenStackDeleteParamsWithContext(h.ctx).
		WithUID(uid)
	_, err := h.Client.V1CloudAccountsOpenStackDelete(params)
	return err
}

func (h *V1Client) GetCloudAccountOpenStack(uid string) (*models.V1OpenStackAccount, error) {
	params := clientV1.NewV1CloudAccountsOpenStackGetParamsWithContext(h.ctx).
		WithUID(uid)
	resp, err := h.Client.V1CloudAccountsOpenStackGet(params)
	if err := apiutil.Handle404(err); err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

func (h *V1Client) GetCloudAccountsOpenStack() ([]*models.V1OpenStackAccount, error) {
	params := clientV1.NewV1CloudAccountsOpenStackListParamsWithContext(h.ctx).
		WithLimit(apiutil.Ptr(int64(0)))
	resp, err := h.Client.V1CloudAccountsOpenStackList(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload.Items, nil
}
