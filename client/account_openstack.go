package client

import (
	"errors"

	"github.com/spectrocloud/palette-api-go/apiutil/transport"
	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
)

func toV1OverlordsUIDOpenStackAccountValidateBody(account *models.V1OpenStackAccount) clientV1.V1OverlordsUIDOpenStackAccountValidateBody {
	return clientV1.V1OverlordsUIDOpenStackAccountValidateBody{
		Account: account.Spec,
	}
}

func (h *V1Client) CreateCloudAccountOpenStack(account *models.V1OpenStackAccount, AccountContext string) (string, error) {
	// validate account
	if err := validateCloudAccountOpenStack(account, h); err != nil {
		return "", err
	}

	var params *clientV1.V1CloudAccountsOpenStackCreateParams
	switch AccountContext {
	case "project":
		params = clientV1.NewV1CloudAccountsOpenStackCreateParamsWithContext(h.Ctx).WithBody(account)
	case "tenant":
		params = clientV1.NewV1CloudAccountsOpenStackCreateParams().WithBody(account)
	}

	success, err := h.GetClient().V1CloudAccountsOpenStackCreate(params)
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

func validateCloudAccountOpenStack(account *models.V1OpenStackAccount, h *V1Client) error {
	PcgId := account.Metadata.Annotations[OverlordUID]
	// check PCG
	if err := h.CheckPCG(PcgId); err != nil {
		return err
	}

	// validate account
	paramsValidate := clientV1.NewV1OverlordsUIDOpenStackAccountValidateParams().WithUID(PcgId)
	paramsValidate = paramsValidate.WithBody(toV1OverlordsUIDOpenStackAccountValidateBody(account))
	_, err := h.GetClient().V1OverlordsUIDOpenStackAccountValidate(paramsValidate)
	if err != nil {
		return err
	}

	return nil
}

func (h *V1Client) UpdateCloudAccountOpenStack(account *models.V1OpenStackAccount) error {
	if err := validateCloudAccountOpenStack(account, h); err != nil {
		return err
	}

	uid := account.Metadata.UID
	params := clientV1.NewV1CloudAccountsOpenStackUpdateParamsWithContext(h.Ctx).WithUID(uid).WithBody(account)
	_, err := h.GetClient().V1CloudAccountsOpenStackUpdate(params)
	return err
}

func (h *V1Client) DeleteCloudAccountOpenStack(uid, AccountContext string) error {
	var params *clientV1.V1CloudAccountsOpenStackDeleteParams
	switch AccountContext {
	case "project":
		params = clientV1.NewV1CloudAccountsOpenStackDeleteParamsWithContext(h.Ctx).WithUID(uid)
	case "tenant":
		params = clientV1.NewV1CloudAccountsOpenStackDeleteParams().WithUID(uid)
	}

	_, err := h.GetClient().V1CloudAccountsOpenStackDelete(params)
	return err
}

func (h *V1Client) GetCloudAccountOpenStack(uid, AccountContext string) (*models.V1OpenStackAccount, error) {
	var params *clientV1.V1CloudAccountsOpenStackGetParams
	switch AccountContext {
	case "project":
		params = clientV1.NewV1CloudAccountsOpenStackGetParamsWithContext(h.Ctx).WithUID(uid)
	case "tenant":
		params = clientV1.NewV1CloudAccountsOpenStackGetParams().WithUID(uid)
	}

	success, err := h.GetClient().V1CloudAccountsOpenStackGet(params)

	var e *transport.TransportError
	if errors.As(err, &e) && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return success.Payload, nil
}

func (h *V1Client) GetCloudAccountsOpenStack() ([]*models.V1OpenStackAccount, error) {
	limit := int64(0)
	params := clientV1.NewV1CloudAccountsOpenStackListParamsWithContext(h.Ctx).WithLimit(&limit)
	response, err := h.GetClient().V1CloudAccountsOpenStackList(params)
	if err != nil {
		return nil, err
	}

	accounts := make([]*models.V1OpenStackAccount, len(response.Payload.Items))
	copy(accounts, response.Payload.Items)

	return accounts, nil
}
