package client

import (
	hapitransport "github.com/spectrocloud/hapi/apiutil/transport"
	"github.com/spectrocloud/hapi/models"
	clusterC "github.com/spectrocloud/hapi/spectrocluster/client/v1"
)

func (h *V1Client) CreateCloudAccountCoxEdge(account *models.V1CoxEdgeAccount, AccountContext string) (string, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return "", err
	}

	var params *clusterC.V1CloudAccountsCoxEdgeCreateParams
	switch AccountContext {
	case "project":
		params = clusterC.NewV1CloudAccountsCoxEdgeCreateParamsWithContext(h.Ctx).WithBody(account)
	case "tenant":
		params = clusterC.NewV1CloudAccountsCoxEdgeCreateParams().WithBody(account)
	}
	success, err := client.V1CloudAccountsCoxEdgeCreate(params)
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

func (h *V1Client) UpdateCloudAccountCoxEdge(account *models.V1CoxEdgeAccount) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return err
	}

	uid := account.Metadata.UID
	var params *clusterC.V1CloudAccountsCoxEdgeUpdateParams
	switch account.Metadata.Annotations["scope"] {
	case "project":
		params = clusterC.NewV1CloudAccountsCoxEdgeUpdateParamsWithContext(h.Ctx).WithUID(uid).WithBody(account)
	case "tenant":
		params = clusterC.NewV1CloudAccountsCoxEdgeUpdateParams().WithBody(account)
	}
	_, err = client.V1CloudAccountsCoxEdgeUpdate(params)
	return err
}

func (h *V1Client) DeleteCloudAccountCoxEdge(uid string) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return err
	}

	account, err := h.GetCloudAccountCoxEdge(uid)
	if err != nil {
		return err
	}

	var params *clusterC.V1CloudAccountsCoxEdgeDeleteParams
	switch account.Metadata.Annotations["scope"] {
	case "project":
		params = clusterC.NewV1CloudAccountsCoxEdgeDeleteParamsWithContext(h.Ctx).WithUID(uid)
	case "tenant":
		params = clusterC.NewV1CloudAccountsCoxEdgeDeleteParams().WithUID(uid)
	}
	_, err = client.V1CloudAccountsCoxEdgeDelete(params)
	return err
}

func (h *V1Client) GetCloudAccountCoxEdge(uid string) (*models.V1CoxEdgeAccount, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}

	params := clusterC.NewV1CloudAccountsCoxEdgeGetParamsWithContext(h.Ctx).WithUID(uid)
	success, err := client.V1CloudAccountsCoxEdgeGet(params)
	if e, ok := err.(*hapitransport.TransportError); ok && e.HttpCode == 404 {

		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return success.Payload, nil
}

func (h *V1Client) GetCloudAccountsCoxEdge() ([]*models.V1CoxEdgeAccount, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}

	limit := int64(0)
	params := clusterC.NewV1CloudAccountsCoxEdgeListParamsWithContext(h.Ctx).WithLimit(&limit)
	response, err := client.V1CloudAccountsCoxEdgeList(params)
	if err != nil {
		return nil, err
	}

	accounts := make([]*models.V1CoxEdgeAccount, len(response.Payload.Items))
	copy(accounts, response.Payload.Items)

	return accounts, nil
}
