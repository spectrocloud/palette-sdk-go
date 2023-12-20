package client

import (
	"errors"

	"github.com/spectrocloud/hapi/apiutil/transport"
	cloudC "github.com/spectrocloud/hapi/cloud/client/v1"
	"github.com/spectrocloud/hapi/models"
	clusterC "github.com/spectrocloud/hapi/spectrocluster/client/v1"
)

func toV1AwsCloudAccount(account *models.V1AwsAccount) *models.V1AwsCloudAccount {
	return &models.V1AwsCloudAccount{
		AccessKey:      account.Spec.AccessKey,
		CredentialType: account.Spec.CredentialType,
		Partition:      account.Spec.Partition,
		PolicyARNs:     account.Spec.PolicyARNs,
		SecretKey:      account.Spec.SecretKey,
		Sts:            account.Spec.Sts,
	}
}

func (h *V1Client) CreateCloudAccountAws(account *models.V1AwsAccount, AccountContext string) (string, error) {
	// validate account
	if err := validateCloudAccountAws(account, h); err != nil {
		return "", err
	}

	var params *clusterC.V1CloudAccountsAwsCreateParams
	switch AccountContext {
	case "project":
		params = clusterC.NewV1CloudAccountsAwsCreateParamsWithContext(h.Ctx).WithBody(account)
	case "tenant":
		params = clusterC.NewV1CloudAccountsAwsCreateParams().WithBody(account)
	}

	success, err := h.GetClusterClient().V1CloudAccountsAwsCreate(params)
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

func validateCloudAccountAws(account *models.V1AwsAccount, h *V1Client) error {
	// check PCG
	if err := h.CheckPCG(account.Metadata.Annotations[OverlordUID]); err != nil {
		return err
	}

	// validate account
	paramsValidate := cloudC.NewV1AwsAccountValidateParams()
	paramsValidate = paramsValidate.WithAwsCloudAccount(toV1AwsCloudAccount(account))
	_, err := h.GetCloudClient().V1AwsAccountValidate(paramsValidate)
	if err != nil {
		return err
	}

	return nil
}

func (h *V1Client) UpdateCloudAccountAws(account *models.V1AwsAccount) error {
	// validate account
	if err := validateCloudAccountAws(account, h); err != nil {
		return err
	}

	uid := account.Metadata.UID
	params := clusterC.NewV1CloudAccountsAwsUpdateParamsWithContext(h.Ctx).WithUID(uid).WithBody(account)
	_, err := h.GetClusterClient().V1CloudAccountsAwsUpdate(params)
	return err
}

func (h *V1Client) DeleteCloudAccountAws(uid, AccountContext string) error {
	var params *clusterC.V1CloudAccountsAwsDeleteParams
	switch AccountContext {
	case "project":
		params = clusterC.NewV1CloudAccountsAwsDeleteParamsWithContext(h.Ctx).WithUID(uid)
	case "tenant":
		params = clusterC.NewV1CloudAccountsAwsDeleteParams().WithUID(uid)
	}
	_, err := h.GetClusterClient().V1CloudAccountsAwsDelete(params)
	return err
}

func (h *V1Client) GetCloudAccountAws(uid, AccountContext string) (*models.V1AwsAccount, error) {
	var params *clusterC.V1CloudAccountsAwsGetParams
	switch AccountContext {
	case "project":
		params = clusterC.NewV1CloudAccountsAwsGetParamsWithContext(h.Ctx).WithUID(uid)
	case "tenant":
		params = clusterC.NewV1CloudAccountsAwsGetParams().WithUID(uid)
	}
	success, err := h.GetClusterClient().V1CloudAccountsAwsGet(params)

	var e *transport.TransportError
	if errors.As(err, &e) && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return success.Payload, nil
}

func (h *V1Client) GetCloudAccountsAws() ([]*models.V1AwsAccount, error) {
	limit := int64(0)
	params := clusterC.NewV1CloudAccountsAwsListParamsWithContext(h.Ctx).WithLimit(&limit)
	response, err := h.GetClusterClient().V1CloudAccountsAwsList(params)
	if err != nil {
		return nil, err
	}

	accounts := make([]*models.V1AwsAccount, len(response.Payload.Items))
	copy(accounts, response.Payload.Items)

	return accounts, nil
}
