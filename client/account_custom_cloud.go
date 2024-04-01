package client

import (
	"fmt"

	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
	"github.com/spectrocloud/palette-sdk-go/client/apiutil"
)

func (h *V1Client) CreateAccountCustomCloud(account *models.V1CustomAccountEntity, cloudType string) (string, error) {
	pcgId := account.Metadata.Annotations[OverlordUID]
	if err := h.CheckPCG(pcgId); err != nil {
		return "", err
	}
	params := clientV1.NewV1CloudAccountsCustomCreateParamsWithContext(h.ctx).
		WithBody(account).
		WithCloudType(cloudType)
	resp, err := h.Client.V1CloudAccountsCustomCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

func (h *V1Client) GetCustomCloudAccount(uid, cloudType string) (*models.V1CustomAccount, error) {
	params := clientV1.NewV1CloudAccountsCustomGetParamsWithContext(h.ctx).
		WithCloudType(cloudType).
		WithUID(uid)
	resp, err := h.Client.V1CloudAccountsCustomGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

func (h *V1Client) UpdateAccountCustomCloud(uid string, account *models.V1CustomAccountEntity, cloudType string) error {
	params := clientV1.NewV1CloudAccountsCustomUpdateParamsWithContext(h.ctx).
		WithBody(account).
		WithCloudType(cloudType).
		WithUID(uid)
	_, err := h.Client.V1CloudAccountsCustomUpdate(params)
	return err
}

func (h *V1Client) DeleteCloudAccountCustomCloud(uid, cloudType string) error {
	params := clientV1.NewV1CloudAccountsCustomDeleteParamsWithContext(h.ctx).
		WithCloudType(cloudType).
		WithUID(uid)
	_, err := h.Client.V1CloudAccountsCustomDelete(params)
	return err
}

func (h *V1Client) ValidateCustomCloudType(cloudType string) error {
	params := clientV1.NewV1CustomCloudTypesGetParamsWithContext(h.ctx)
	resp, err := h.Client.V1CustomCloudTypesGet(params)
	if err != nil {
		return err
	}
	for _, c := range resp.GetPayload().CloudTypes {
		if c.Name == cloudType {
			if c.IsCustom {
				return nil
			} else {
				return fmt.Errorf("cloud - `%s` is not a valid custom cloud", cloudType)
			}
		}
	}
	return fmt.Errorf("cloud - `%s` is not a valid cloud", cloudType)
}

func (h *V1Client) GetCustomCloudAccountList(cloudType string) ([]*models.V1CustomAccount, error) {
	params := clientV1.NewV1CloudAccountsCustomListParamsWithContext(h.ctx).
		WithLimit(apiutil.Ptr(int64(0))).
		WithCloudType(cloudType)
	resp, err := h.Client.V1CloudAccountsCustomList(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload.Items, nil
}
