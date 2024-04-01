package client

import (
	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
	"github.com/spectrocloud/palette-sdk-go/client/apiutil"
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

func (h *V1Client) CreateCloudAccountAws(account *models.V1AwsAccount) (string, error) {
	if err := h.validateCloudAccountAws(account); err != nil {
		return "", err
	}
	params := clientV1.NewV1CloudAccountsAwsCreateParamsWithContext(h.ctx).
		WithBody(account)

	resp, err := h.Client.V1CloudAccountsAwsCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

func (h *V1Client) validateCloudAccountAws(account *models.V1AwsAccount) error {
	// check PCG
	if err := h.CheckPCG(account.Metadata.Annotations[OverlordUID]); err != nil {
		return err
	}

	// validate account
	params := clientV1.NewV1AwsAccountValidateParamsWithContext(h.ctx).
		WithAwsCloudAccount(toV1AwsCloudAccount(account))

	_, err := h.Client.V1AwsAccountValidate(params)
	return err
}

func (h *V1Client) UpdateCloudAccountAws(account *models.V1AwsAccount) error {
	if err := h.validateCloudAccountAws(account); err != nil {
		return err
	}
	params := clientV1.NewV1CloudAccountsAwsUpdateParamsWithContext(h.ctx).
		WithUID(account.Metadata.UID).
		WithBody(account)
	_, err := h.Client.V1CloudAccountsAwsUpdate(params)
	return err
}

func (h *V1Client) DeleteCloudAccountAws(uid string) error {
	params := clientV1.NewV1CloudAccountsAwsDeleteParamsWithContext(h.ctx).
		WithUID(uid)
	_, err := h.Client.V1CloudAccountsAwsDelete(params)
	return err
}

func (h *V1Client) GetCloudAccountAws(uid string) (*models.V1AwsAccount, error) {
	params := clientV1.NewV1CloudAccountsAwsGetParamsWithContext(h.ctx).
		WithUID(uid)
	resp, err := h.Client.V1CloudAccountsAwsGet(params)
	if err := apiutil.Handle404(err); err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

func (h *V1Client) GetCloudAccountsAws() ([]*models.V1AwsAccount, error) {
	params := clientV1.NewV1CloudAccountsAwsListParamsWithContext(h.ctx).
		WithLimit(apiutil.Ptr(int64(0)))
	resp, err := h.Client.V1CloudAccountsAwsList(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload.Items, nil
}
