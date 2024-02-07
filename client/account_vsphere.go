package client

import (
	"errors"

	"github.com/spectrocloud/palette-api-go/apiutil/transport"
	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
)

const OverlordUID = "overlordUid"

// convert V1VsphereAccount to V1OverlordVsphereAccountEntity
func toV1OverlordsUIDVsphereAccountValidateBody(account *models.V1VsphereAccount) clientV1.V1OverlordsUIDVsphereAccountValidateBody {
	return clientV1.V1OverlordsUIDVsphereAccountValidateBody{
		Account: account.Spec,
	}
}

func (h *V1Client) CreateCloudAccountVsphere(account *models.V1VsphereAccount, AccountContext string) (string, error) {
	// validate account
	if err := validateCloudAccountVsphere(account, h); err != nil {
		return "", err
	}

	// create account
	var params *clientV1.V1CloudAccountsVsphereCreateParams
	switch AccountContext {
	case "project":
		params = clientV1.NewV1CloudAccountsVsphereCreateParamsWithContext(h.Ctx)
	case "tenant":
		params = clientV1.NewV1CloudAccountsVsphereCreateParams()
	}

	params = params.WithBody(account)
	success, err := h.Client.V1CloudAccountsVsphereCreate(params)
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

func validateCloudAccountVsphere(account *models.V1VsphereAccount, h *V1Client) error {
	PcgId := account.Metadata.Annotations[OverlordUID]
	// check PCG
	if err := h.CheckPCG(PcgId); err != nil {
		return err
	}

	// validate account
	paramsValidate := clientV1.NewV1OverlordsUIDVsphereAccountValidateParams().WithUID(PcgId)
	paramsValidate = paramsValidate.WithBody(toV1OverlordsUIDVsphereAccountValidateBody(account))
	_, err := h.Client.V1OverlordsUIDVsphereAccountValidate(paramsValidate)
	if err != nil {
		return err
	}

	return nil
}

func (h *V1Client) UpdateCloudAccountVsphere(account *models.V1VsphereAccount, AccountContext string) error {
	// validate account
	if err := validateCloudAccountVsphere(account, h); err != nil {
		return err
	}

	uid := account.Metadata.UID
	params := clientV1.NewV1CloudAccountsVsphereUpdateParamsWithContext(h.Ctx).WithUID(uid).WithBody(account)
	_, err := h.Client.V1CloudAccountsVsphereUpdate(params)
	return err
}

func (h *V1Client) DeleteCloudAccountVsphere(uid, AccountContext string) error {
	var params *clientV1.V1CloudAccountsVsphereDeleteParams
	switch AccountContext {
	case "project":
		params = clientV1.NewV1CloudAccountsVsphereDeleteParamsWithContext(h.Ctx).WithUID(uid)
	case "tenant":
		params = clientV1.NewV1CloudAccountsVsphereDeleteParams().WithUID(uid)
	}

	_, err := h.Client.V1CloudAccountsVsphereDelete(params)
	return err
}

func (h *V1Client) GetCloudAccountVsphere(uid, AccountContext string) (*models.V1VsphereAccount, error) {
	var params *clientV1.V1CloudAccountsVsphereGetParams
	switch AccountContext {
	case "project":
		params = clientV1.NewV1CloudAccountsVsphereGetParamsWithContext(h.Ctx).WithUID(uid)
	case "tenant":
		params = clientV1.NewV1CloudAccountsVsphereGetParams().WithUID(uid)
	}

	success, err := h.Client.V1CloudAccountsVsphereGet(params)

	var e *transport.TransportError
	if errors.As(err, &e) && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return success.Payload, nil
}

func (h *V1Client) GetCloudAccountsVsphere() ([]*models.V1VsphereAccount, error) {
	limit := int64(0)
	params := clientV1.NewV1CloudAccountsVsphereListParamsWithContext(h.Ctx).WithLimit(&limit)
	response, err := h.Client.V1CloudAccountsVsphereList(params)
	if err != nil {
		return nil, err
	}

	accounts := make([]*models.V1VsphereAccount, len(response.Payload.Items))
	copy(accounts, response.Payload.Items)

	return accounts, nil
}
