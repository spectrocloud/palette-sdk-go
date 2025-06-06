package client

import (
	clientv1 "github.com/spectrocloud/palette-sdk-go/api/client/version1"
	"github.com/spectrocloud/palette-sdk-go/api/models"
	"github.com/spectrocloud/palette-sdk-go/client/apiutil"
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

// CreateCloudAccountAzure creates a new Azure cloud account.
func (h *V1Client) CreateCloudAccountAzure(account *models.V1AzureAccount) (string, error) {
	// validate account
	if err := h.validateCloudAccountAzure(account); err != nil {
		return "", err
	}

	params := clientv1.NewV1CloudAccountsAzureCreateParamsWithContext(h.ctx).
		WithBody(account)
	resp, err := h.Client.V1CloudAccountsAzureCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

func (h *V1Client) validateCloudAccountAzure(account *models.V1AzureAccount) error {
	// check PCG
	if err := h.CheckPCG(account.Metadata.Annotations[OverlordUID]); err != nil {
		return err
	}

	// validate account
	params := clientv1.NewV1AzureAccountValidateParamsWithContext(h.ctx).
		WithAzureCloudAccount(toV1AzureCloudAccount(account))

	_, err := h.Client.V1AzureAccountValidate(params)
	return err
}

// UpdateCloudAccountAzure updates an existing Azure cloud account.
func (h *V1Client) UpdateCloudAccountAzure(account *models.V1AzureAccount) error {
	// validate account
	if err := h.validateCloudAccountAzure(account); err != nil {
		return err
	}

	params := clientv1.NewV1CloudAccountsAzureUpdateParamsWithContext(h.ctx).
		WithUID(account.Metadata.UID).
		WithBody(account)

	_, err := h.Client.V1CloudAccountsAzureUpdate(params)
	return err
}

// DeleteCloudAccountAzure deletes an existing Azure cloud account.
func (h *V1Client) DeleteCloudAccountAzure(uid string) error {
	params := clientv1.NewV1CloudAccountsAzureDeleteParamsWithContext(h.ctx).
		WithUID(uid)
	_, err := h.Client.V1CloudAccountsAzureDelete(params)
	return err
}

// GetCloudAccountAzure retrieves an existing Azure cloud account.
func (h *V1Client) GetCloudAccountAzure(uid string) (*models.V1AzureAccount, error) {
	params := clientv1.NewV1CloudAccountsAzureGetParamsWithContext(h.ctx).
		WithUID(uid)
	resp, err := h.Client.V1CloudAccountsAzureGet(params)
	if apiutil.Is404(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// GetCloudAccountsAzure retrieves all Azure cloud accounts.
func (h *V1Client) GetCloudAccountsAzure() ([]*models.V1AzureAccount, error) {
	params := clientv1.NewV1CloudAccountsAzureListParamsWithContext(h.ctx).
		WithLimit(apiutil.Ptr(int64(0)))
	resp, err := h.Client.V1CloudAccountsAzureList(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload.Items, nil
}
