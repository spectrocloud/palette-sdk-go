package client

import (
	"errors"

	"github.com/spectrocloud/hapi/apiutil/transport"
	"github.com/spectrocloud/hapi/models"
	clusterC "github.com/spectrocloud/hapi/spectrocluster/client/v1"
)

func (h *V1Client) CreateCloudAccountCoxEdge(account *models.V1CoxEdgeAccount, AccountContext string) (string, error) {
	var params *clusterC.V1CloudAccountsCoxEdgeCreateParams
	switch AccountContext {
	case "project":
		params = clusterC.NewV1CloudAccountsCoxEdgeCreateParamsWithContext(h.Ctx).WithBody(account)
	case "tenant":
		params = clusterC.NewV1CloudAccountsCoxEdgeCreateParams().WithBody(account)
	}
	success, err := h.GetClusterClient().V1CloudAccountsCoxEdgeCreate(params)
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

func (h *V1Client) UpdateCloudAccountCoxEdge(account *models.V1CoxEdgeAccount) error {
	uid := account.Metadata.UID
	params := clusterC.NewV1CloudAccountsCoxEdgeUpdateParamsWithContext(h.Ctx).WithUID(uid).WithBody(account)
	_, err := h.GetClusterClient().V1CloudAccountsCoxEdgeUpdate(params)
	return err
}

func (h *V1Client) DeleteCloudAccountCoxEdge(uid, AccountContext string) error {
	var params *clusterC.V1CloudAccountsCoxEdgeDeleteParams
	switch AccountContext {
	case "project":
		params = clusterC.NewV1CloudAccountsCoxEdgeDeleteParamsWithContext(h.Ctx).WithUID(uid)
	case "tenant":
		params = clusterC.NewV1CloudAccountsCoxEdgeDeleteParams().WithUID(uid)
	}
	_, err := h.GetClusterClient().V1CloudAccountsCoxEdgeDelete(params)
	return err
}

func (h *V1Client) GetCloudAccountCoxEdge(uid, AccountContext string) (*models.V1CoxEdgeAccount, error) {
	var params *clusterC.V1CloudAccountsCoxEdgeGetParams
	switch AccountContext {
	case "project":
		params = clusterC.NewV1CloudAccountsCoxEdgeGetParamsWithContext(h.Ctx).WithUID(uid)
	case "tenant":
		params = clusterC.NewV1CloudAccountsCoxEdgeGetParams().WithUID(uid)
	}

	success, err := h.GetClusterClient().V1CloudAccountsCoxEdgeGet(params)

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
	params := clusterC.NewV1CloudAccountsCoxEdgeListParamsWithContext(h.Ctx).WithLimit(&limit)
	response, err := h.GetClusterClient().V1CloudAccountsCoxEdgeList(params)
	if err != nil {
		return nil, err
	}

	accounts := make([]*models.V1CoxEdgeAccount, len(response.Payload.Items))
	copy(accounts, response.Payload.Items)

	return accounts, nil
}
