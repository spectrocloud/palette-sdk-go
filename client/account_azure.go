package client

import (
	"errors"

	"github.com/spectrocloud/hapi/apiutil/transport"
	cloudC "github.com/spectrocloud/hapi/cloud/client/v1"
	"github.com/spectrocloud/hapi/models"
	clusterC "github.com/spectrocloud/hapi/spectrocluster/client/v1"
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
	client, err := h.GetClusterClient()
	if err != nil {
		return "", err
	}

	// validate account
	err = validateCloudAccountAzure(account, h)
	if err != nil {
		return "", err
	}

	var params *clusterC.V1CloudAccountsAzureCreateParams
	switch AccountContext {
	case "project":
		params = clusterC.NewV1CloudAccountsAzureCreateParamsWithContext(h.Ctx).WithBody(account)
	case "tenant":
		params = clusterC.NewV1CloudAccountsAzureCreateParams().WithBody(account)
	}
	success, err := client.V1CloudAccountsAzureCreate(params)
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

func validateCloudAccountAzure(account *models.V1AzureAccount, h *V1Client) error {
	client, err := h.GetCloudClient()
	if err != nil {
		return err
	}

	PcgId := account.Metadata.Annotations[OverlordUID]
	// check PCG
	err = h.CheckPCG(PcgId)
	if err != nil {
		return err
	}

	// validate account
	paramsValidate := cloudC.NewV1AzureAccountValidateParams()
	paramsValidate = paramsValidate.WithAzureCloudAccount(toV1AzureCloudAccount(account))
	_, err = client.V1AzureAccountValidate(paramsValidate)
	if err != nil {
		return err
	}

	return nil
}

func (h *V1Client) UpdateCloudAccountAzure(account *models.V1AzureAccount) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil
	}

	// validate account
	err = validateCloudAccountAzure(account, h)
	if err != nil {
		return err
	}

	uid := account.Metadata.UID
	params := clusterC.NewV1CloudAccountsAzureUpdateParamsWithContext(h.Ctx).WithUID(uid).WithBody(account)
	_, err = client.V1CloudAccountsAzureUpdate(params)
	return err
}

func (h *V1Client) DeleteCloudAccountAzure(uid, AccountContext string) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil
	}

	var params *clusterC.V1CloudAccountsAzureDeleteParams
	switch AccountContext {
	case "project":
		params = clusterC.NewV1CloudAccountsAzureDeleteParamsWithContext(h.Ctx).WithUID(uid)
	case "tenant":
		params = clusterC.NewV1CloudAccountsAzureDeleteParams().WithUID(uid)
	}
	_, err = client.V1CloudAccountsAzureDelete(params)
	return err
}

func (h *V1Client) GetCloudAccountAzure(uid, AccountContext string) (*models.V1AzureAccount, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}

	var params *clusterC.V1CloudAccountsAzureGetParams
	switch AccountContext {
	case "project":
		params = clusterC.NewV1CloudAccountsAzureGetParamsWithContext(h.Ctx).WithUID(uid)
	case "tenant":
		params = clusterC.NewV1CloudAccountsAzureGetParams().WithUID(uid)
	}

	success, err := client.V1CloudAccountsAzureGet(params)
	var e *transport.TransportError
	if errors.As(err, &e) && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return success.Payload, nil
}

func (h *V1Client) GetCloudAccountsAzure() ([]*models.V1AzureAccount, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}

	limit := int64(0)
	params := clusterC.NewV1CloudAccountsAzureListParamsWithContext(h.Ctx).WithLimit(&limit)
	response, err := client.V1CloudAccountsAzureList(params)
	if err != nil {
		return nil, err
	}

	accounts := make([]*models.V1AzureAccount, len(response.Payload.Items))
	copy(accounts, response.Payload.Items)

	return accounts, nil
}
