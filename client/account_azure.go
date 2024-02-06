package client

import (
	"errors"

	"github.com/spectrocloud/palette-api-go/apiutil/transport"
	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
)

func toV1AzureCloudAccount(account *models.V1AzureAccount) *models.V1AzureCloudAccount {
	return &models.V1AzureCloudAccount{
		AzureEnvironment: account.Spec.AzureEnvironment,
		ClientID:         account.Spec.ClientID,
		ClientSecret:     account.Spec.ClientSecret,
		Settings:         account.Spec.Settings,
		TenantID:         account.Spec.TenantID,
		TenantName:       account.Spec.TenantName,
	}
}

func (h *V1Client) CreateCloudAccountAzure(account *models.V1AzureAccount, AccountContext string) (string, error) {
	// validate account
	if err := validateCloudAccountAzure(account, h); err != nil {
		return "", err
	}

	var params *clientV1.V1CloudAccountsAzureCreateParams
	switch AccountContext {
	case "project":
		params = clientV1.NewV1CloudAccountsAzureCreateParamsWithContext(h.Ctx).WithBody(account)
	case "tenant":
		params = clientV1.NewV1CloudAccountsAzureCreateParams().WithBody(account)
	}
	success, err := h.GetClient().V1CloudAccountsAzureCreate(params)
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

func validateCloudAccountAzure(account *models.V1AzureAccount, h *V1Client) error {
	// check PCG
	if err := h.CheckPCG(account.Metadata.Annotations[OverlordUID]); err != nil {
		return err
	}

	// validate account
	paramsValidate := clientV1.NewV1AzureAccountValidateParams()
	paramsValidate = paramsValidate.WithAzureCloudAccount(toV1AzureCloudAccount(account))

	_, err := h.GetClient().V1AzureAccountValidate(paramsValidate)
	if err != nil {
		return err
	}

	return nil
}

func (h *V1Client) UpdateCloudAccountAzure(account *models.V1AzureAccount) error {
	// validate account
	if err := validateCloudAccountAzure(account, h); err != nil {
		return err
	}

	uid := account.Metadata.UID
	params := clientV1.NewV1CloudAccountsAzureUpdateParamsWithContext(h.Ctx).WithUID(uid).WithBody(account)
	_, err := h.GetClient().V1CloudAccountsAzureUpdate(params)
	return err
}

func (h *V1Client) DeleteCloudAccountAzure(uid, AccountContext string) error {
	var params *clientV1.V1CloudAccountsAzureDeleteParams
	switch AccountContext {
	case "project":
		params = clientV1.NewV1CloudAccountsAzureDeleteParamsWithContext(h.Ctx).WithUID(uid)
	case "tenant":
		params = clientV1.NewV1CloudAccountsAzureDeleteParams().WithUID(uid)
	}
	_, err := h.GetClient().V1CloudAccountsAzureDelete(params)
	return err
}

func (h *V1Client) GetCloudAccountAzure(uid, AccountContext string) (*models.V1AzureAccount, error) {
	var params *clientV1.V1CloudAccountsAzureGetParams
	switch AccountContext {
	case "project":
		params = clientV1.NewV1CloudAccountsAzureGetParamsWithContext(h.Ctx).WithUID(uid)
	case "tenant":
		params = clientV1.NewV1CloudAccountsAzureGetParams().WithUID(uid)
	}

	success, err := h.GetClient().V1CloudAccountsAzureGet(params)

	var e *transport.TransportError
	if errors.As(err, &e) && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return success.Payload, nil
}

func (h *V1Client) GetCloudAccountsAzure() ([]*models.V1AzureAccount, error) {
	limit := int64(0)
	params := clientV1.NewV1CloudAccountsAzureListParamsWithContext(h.Ctx).WithLimit(&limit)
	response, err := h.GetClient().V1CloudAccountsAzureList(params)
	if err != nil {
		return nil, err
	}

	accounts := make([]*models.V1AzureAccount, len(response.Payload.Items))
	copy(accounts, response.Payload.Items)

	return accounts, nil
}
