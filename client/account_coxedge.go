package client

import (
	"errors"

	"github.com/spectrocloud/palette-api-go/apiutil/transport"
	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
)

func (h *V1Client) CreateCloudAccountCoxEdge(account *models.V1CoxEdgeAccount, AccountContext string) (string, error) {
	var params *clientV1.V1CloudAccountsCoxEdgeCreateParams
	switch AccountContext {
	case "project":
		params = clientV1.NewV1CloudAccountsCoxEdgeCreateParamsWithContext(h.Ctx).WithBody(account)
	case "tenant":
		params = clientV1.NewV1CloudAccountsCoxEdgeCreateParams().WithBody(account)
	}
	success, err := h.GetClient().V1CloudAccountsCoxEdgeCreate(params)
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

func (h *V1Client) UpdateCloudAccountCoxEdge(account *models.V1CoxEdgeAccount) error {
	uid := account.Metadata.UID
	params := clientV1.NewV1CloudAccountsCoxEdgeUpdateParamsWithContext(h.Ctx).WithUID(uid).WithBody(account)
	_, err := h.GetClient().V1CloudAccountsCoxEdgeUpdate(params)
	return err
}

func (h *V1Client) DeleteCloudAccountCoxEdge(uid, AccountContext string) error {
	var params *clientV1.V1CloudAccountsCoxEdgeDeleteParams
	switch AccountContext {
	case "project":
		params = clientV1.NewV1CloudAccountsCoxEdgeDeleteParamsWithContext(h.Ctx).WithUID(uid)
	case "tenant":
		params = clientV1.NewV1CloudAccountsCoxEdgeDeleteParams().WithUID(uid)
	}
	_, err := h.GetClient().V1CloudAccountsCoxEdgeDelete(params)
	return err
}

func (h *V1Client) GetCloudAccountCoxEdge(uid, AccountContext string) (*models.V1CoxEdgeAccount, error) {
	var params *clientV1.V1CloudAccountsCoxEdgeGetParams
	switch AccountContext {
	case "project":
		params = clientV1.NewV1CloudAccountsCoxEdgeGetParamsWithContext(h.Ctx).WithUID(uid)
	case "tenant":
		params = clientV1.NewV1CloudAccountsCoxEdgeGetParams().WithUID(uid)
	}

	success, err := h.GetClient().V1CloudAccountsCoxEdgeGet(params)

	var e *transport.TransportError
	if errors.As(err, &e) && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return success.Payload, nil
}

func (h *V1Client) GetCloudAccountsCoxEdge() ([]*models.V1CoxEdgeAccount, error) {
	limit := int64(0)
	params := clientV1.NewV1CloudAccountsCoxEdgeListParamsWithContext(h.Ctx).WithLimit(&limit)
	response, err := h.GetClient().V1CloudAccountsCoxEdgeList(params)
	if err != nil {
		return nil, err
	}

	accounts := make([]*models.V1CoxEdgeAccount, len(response.Payload.Items))
	copy(accounts, response.Payload.Items)

	return accounts, nil
}
