package client

import (
	"errors"

	"github.com/spectrocloud/palette-api-go/apiutil/transport"
	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
)

func toV1OverlordsUIDMaasAccountValidateBody(account *models.V1MaasAccount) clientV1.V1OverlordsUIDMaasAccountValidateBody {
	return clientV1.V1OverlordsUIDMaasAccountValidateBody{
		Account: account.Spec,
	}
}

func (h *V1Client) CreateCloudAccountMaas(account *models.V1MaasAccount, AccountContext string) (string, error) {
	// validate account
	if err := validateCloudAccountMaas(account, h); err != nil {
		return "", err
	}

	var params *clientV1.V1CloudAccountsMaasCreateParams
	switch AccountContext {
	case "project":
		params = clientV1.NewV1CloudAccountsMaasCreateParamsWithContext(h.Ctx).WithBody(account)
	case "tenant":
		params = clientV1.NewV1CloudAccountsMaasCreateParams().WithBody(account)
	}

	success, err := h.Client.V1CloudAccountsMaasCreate(params)
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

func validateCloudAccountMaas(account *models.V1MaasAccount, h *V1Client) error {
	PcgId := account.Metadata.Annotations[OverlordUID]
	// check PCG
	if err := h.CheckPCG(PcgId); err != nil {
		return err
	}

	// validate account
	paramsValidate := clientV1.NewV1OverlordsUIDMaasAccountValidateParams().WithUID(PcgId)
	paramsValidate = paramsValidate.WithBody(toV1OverlordsUIDMaasAccountValidateBody(account))
	_, err := h.Client.V1OverlordsUIDMaasAccountValidate(paramsValidate)
	if err != nil {
		return err
	}

	return nil
}

func (h *V1Client) UpdateCloudAccountMaas(account *models.V1MaasAccount) error {
	// validate account
	if err := validateCloudAccountMaas(account, h); err != nil {
		return err
	}

	uid := account.Metadata.UID
	params := clientV1.NewV1CloudAccountsMaasUpdateParamsWithContext(h.Ctx).WithUID(uid).WithBody(account)
	_, err := h.Client.V1CloudAccountsMaasUpdate(params)
	return err
}

func (h *V1Client) DeleteCloudAccountMaas(uid, AccountContext string) error {
	var params *clientV1.V1CloudAccountsMaasDeleteParams
	switch AccountContext {
	case "project":
		params = clientV1.NewV1CloudAccountsMaasDeleteParamsWithContext(h.Ctx).WithUID(uid)
	case "tenant":
		params = clientV1.NewV1CloudAccountsMaasDeleteParams().WithUID(uid)
	}

	_, err := h.Client.V1CloudAccountsMaasDelete(params)
	return err
}

func (h *V1Client) GetCloudAccountMaas(uid, AccountContext string) (*models.V1MaasAccount, error) {
	var params *clientV1.V1CloudAccountsMaasGetParams
	switch AccountContext {
	case "project":
		params = clientV1.NewV1CloudAccountsMaasGetParamsWithContext(h.Ctx).WithUID(uid)
	case "tenant":
		params = clientV1.NewV1CloudAccountsMaasGetParams().WithUID(uid)
	}

	success, err := h.Client.V1CloudAccountsMaasGet(params)

	var e *transport.TransportError
	if errors.As(err, &e) && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return success.Payload, nil
}

func (h *V1Client) GetCloudAccountsMaas() ([]*models.V1MaasAccount, error) {
	limit := int64(0)
	params := clientV1.NewV1CloudAccountsMaasListParamsWithContext(h.Ctx).WithLimit(&limit)
	response, err := h.Client.V1CloudAccountsMaasList(params)
	if err != nil {
		return nil, err
	}

	accounts := make([]*models.V1MaasAccount, len(response.Payload.Items))
	copy(accounts, response.Payload.Items)

	return accounts, nil
}
