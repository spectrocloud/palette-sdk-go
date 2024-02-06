package client

import (
	"errors"

	"github.com/spectrocloud/palette-api-go/apiutil/transport"
	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
)

func (h *V1Client) CreateCloudAccountGcp(account *models.V1GcpAccountEntity, AccountContext string) (string, error) {
	var params *clientV1.V1CloudAccountsGcpCreateParams
	switch AccountContext {
	case "project":
		params = clientV1.NewV1CloudAccountsGcpCreateParamsWithContext(h.Ctx).WithBody(account)
	case "tenant":
		params = clientV1.NewV1CloudAccountsGcpCreateParams().WithBody(account)
	}

	success, err := h.GetClient().V1CloudAccountsGcpCreate(params)
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

func (h *V1Client) UpdateCloudAccountGcp(account *models.V1GcpAccountEntity) error {
	uid := account.Metadata.UID
	params := clientV1.NewV1CloudAccountsGcpUpdateParamsWithContext(h.Ctx).WithUID(uid).WithBody(account)
	_, err := h.GetClient().V1CloudAccountsGcpUpdate(params)
	return err
}

func (h *V1Client) DeleteCloudAccountGcp(uid, AccountContext string) error {
	var params *clientV1.V1CloudAccountsGcpDeleteParams
	switch AccountContext {
	case "project":
		params = clientV1.NewV1CloudAccountsGcpDeleteParamsWithContext(h.Ctx).WithUID(uid)
	case "tenant":
		params = clientV1.NewV1CloudAccountsGcpDeleteParams().WithUID(uid)
	}

	_, err := h.GetClient().V1CloudAccountsGcpDelete(params)
	return err
}

func (h *V1Client) GetCloudAccountGcp(uid, AccountContext string) (*models.V1GcpAccount, error) {
	var params *clientV1.V1CloudAccountsGcpGetParams
	switch AccountContext {
	case "project":
		params = clientV1.NewV1CloudAccountsGcpGetParamsWithContext(h.Ctx).WithUID(uid)
	case "tenant":
		params = clientV1.NewV1CloudAccountsGcpGetParams().WithUID(uid)
	}

	success, err := h.GetClient().V1CloudAccountsGcpGet(params)

	var e *transport.TransportError
	if errors.As(err, &e) && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return success.Payload, nil
}

func (h *V1Client) GetCloudAccountsGcp() ([]*models.V1GcpAccount, error) {
	limit := int64(0)
	params := clientV1.NewV1CloudAccountsGcpListParamsWithContext(h.Ctx).WithLimit(&limit)
	response, err := h.GetClient().V1CloudAccountsGcpList(params)
	if err != nil {
		return nil, err
	}

	accounts := make([]*models.V1GcpAccount, len(response.Payload.Items))
	copy(accounts, response.Payload.Items)

	return accounts, nil
}
