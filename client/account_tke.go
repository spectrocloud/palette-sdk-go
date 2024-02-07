package client

import (
	"errors"

	"github.com/spectrocloud/palette-api-go/apiutil/transport"
	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
)

func (h *V1Client) CreateCloudAccountTke(account *models.V1TencentAccount, AccountContext string) (string, error) {
	var params *clientV1.V1CloudAccountsTencentCreateParams
	switch AccountContext {
	case "project":
		params = clientV1.NewV1CloudAccountsTencentCreateParamsWithContext(h.Ctx).WithBody(account)
	case "tenant":
		params = clientV1.NewV1CloudAccountsTencentCreateParams().WithBody(account)
	}

	success, err := h.Client.V1CloudAccountsTencentCreate(params)
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

func (h *V1Client) UpdateCloudAccountTencent(account *models.V1TencentAccount) error {
	uid := account.Metadata.UID
	params := clientV1.NewV1CloudAccountsTencentUpdateParamsWithContext(h.Ctx).WithUID(uid).WithBody(account)
	_, err := h.Client.V1CloudAccountsTencentUpdate(params)
	return err
}

func (h *V1Client) DeleteCloudAccountTke(uid, AccountContext string) error {
	var params *clientV1.V1CloudAccountsTencentDeleteParams
	switch AccountContext {
	case "project":
		params = clientV1.NewV1CloudAccountsTencentDeleteParamsWithContext(h.Ctx).WithUID(uid)
	case "tenant":
		params = clientV1.NewV1CloudAccountsTencentDeleteParams().WithUID(uid)
	}

	_, err := h.Client.V1CloudAccountsTencentDelete(params)
	return err
}

func (h *V1Client) GetCloudAccountTke(uid, AccountContext string) (*models.V1TencentAccount, error) {
	var params *clientV1.V1CloudAccountsTencentGetParams
	switch AccountContext {
	case "project":
		params = clientV1.NewV1CloudAccountsTencentGetParamsWithContext(h.Ctx).WithUID(uid)
	case "tenant":
		params = clientV1.NewV1CloudAccountsTencentGetParams().WithUID(uid)
	}

	success, err := h.Client.V1CloudAccountsTencentGet(params)

	var e *transport.TransportError
	if errors.As(err, &e) && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return success.Payload, nil
}

func (h *V1Client) GetCloudAccountsTke() ([]*models.V1TencentAccount, error) {
	limit := int64(0)
	params := clientV1.NewV1CloudAccountsTencentListParamsWithContext(h.Ctx).WithLimit(&limit)
	response, err := h.Client.V1CloudAccountsTencentList(params)
	if err != nil {
		return nil, err
	}

	accounts := make([]*models.V1TencentAccount, len(response.Payload.Items))
	copy(accounts, response.Payload.Items)

	return accounts, nil
}
