package client

import (
	clientv1 "github.com/spectrocloud/palette-sdk-go/api/client/version1"
	"github.com/spectrocloud/palette-sdk-go/api/models"
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

// CreateCloudAccountAws creates a new AWS cloud account.
func (h *V1Client) CreateCloudAccountAws(account *models.V1AwsAccount) (string, error) {
	if err := h.validateCloudAccountAws(account); err != nil {
		return "", err
	}
	params := clientv1.NewV1CloudAccountsAwsCreateParamsWithContext(h.ctx).
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
	params := clientv1.NewV1AwsAccountValidateParamsWithContext(h.ctx).
		WithAwsCloudAccount(toV1AwsCloudAccount(account))

	_, err := h.Client.V1AwsAccountValidate(params)
	return err
}

// UpdateCloudAccountAws updates an existing AWS cloud account.
func (h *V1Client) UpdateCloudAccountAws(account *models.V1AwsAccount) error {
	if err := h.validateCloudAccountAws(account); err != nil {
		return err
	}
	params := clientv1.NewV1CloudAccountsAwsUpdateParamsWithContext(h.ctx).
		WithUID(account.Metadata.UID).
		WithBody(account)
	_, err := h.Client.V1CloudAccountsAwsUpdate(params)
	return err
}

// DeleteCloudAccountAws deletes an existing AWS cloud account.
func (h *V1Client) DeleteCloudAccountAws(uid string) error {
	params := clientv1.NewV1CloudAccountsAwsDeleteParamsWithContext(h.ctx).
		WithUID(uid)
	_, err := h.Client.V1CloudAccountsAwsDelete(params)
	return err
}

// GetCloudAccountAws retrieves an existing AWS cloud account.
func (h *V1Client) GetCloudAccountAws(uid string) (*models.V1AwsAccount, error) {
	params := clientv1.NewV1CloudAccountsAwsGetParamsWithContext(h.ctx).
		WithUID(uid)
	resp, err := h.Client.V1CloudAccountsAwsGet(params)
	if apiutil.Is404(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// GetCloudAccountsAws retrieves all AWS cloud accounts.
func (h *V1Client) GetCloudAccountsAws() ([]*models.V1AwsAccount, error) {
	params := clientv1.NewV1CloudAccountsAwsListParamsWithContext(h.ctx).
		WithLimit(apiutil.Ptr(int64(0)))
	resp, err := h.Client.V1CloudAccountsAwsList(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload.Items, nil
}
