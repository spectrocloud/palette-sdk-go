package client

import (
	"errors"

	"github.com/spectrocloud/hapi/apiutil/transport"
	"github.com/spectrocloud/hapi/models"
	clusterC "github.com/spectrocloud/hapi/spectrocluster/client/v1"
)

func (h *V1Client) CreateCloudAccountTke(account *models.V1TencentAccount, AccountContext string) (string, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return "", err
	}

	var params *clusterC.V1CloudAccountsTencentCreateParams
	switch AccountContext {
	case "project":
		params = clusterC.NewV1CloudAccountsTencentCreateParamsWithContext(h.Ctx).WithBody(account)
	case "tenant":
		params = clusterC.NewV1CloudAccountsTencentCreateParams().WithBody(account)
	}

	success, err := client.V1CloudAccountsTencentCreate(params)
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

func (h *V1Client) UpdateCloudAccountTencent(account *models.V1TencentAccount) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil
	}

	uid := account.Metadata.UID
	params := clusterC.NewV1CloudAccountsTencentUpdateParamsWithContext(h.Ctx).WithUID(uid).WithBody(account)
	_, err = client.V1CloudAccountsTencentUpdate(params)
	return err
}

func (h *V1Client) DeleteCloudAccountTke(uid, AccountContext string) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil
	}

	var params *clusterC.V1CloudAccountsTencentDeleteParams
	switch AccountContext {
	case "project":
		params = clusterC.NewV1CloudAccountsTencentDeleteParamsWithContext(h.Ctx).WithUID(uid)
	case "tenant":
		params = clusterC.NewV1CloudAccountsTencentDeleteParams().WithUID(uid)
	}

	_, err = client.V1CloudAccountsTencentDelete(params)
	return err
}

func (h *V1Client) GetCloudAccountTke(uid, AccountContext string) (*models.V1TencentAccount, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}

	var params *clusterC.V1CloudAccountsTencentGetParams
	switch AccountContext {
	case "project":
		params = clusterC.NewV1CloudAccountsTencentGetParamsWithContext(h.Ctx).WithUID(uid)
	case "tenant":
		params = clusterC.NewV1CloudAccountsTencentGetParams().WithUID(uid)
	}

	success, err := client.V1CloudAccountsTencentGet(params)

	var e *transport.TransportError
	if errors.As(err, &e) && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return success.Payload, nil
}

func (h *V1Client) GetCloudAccountsTke() ([]*models.V1TencentAccount, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}

	limit := int64(0)
	params := clusterC.NewV1CloudAccountsTencentListParamsWithContext(h.Ctx).WithLimit(&limit)
	response, err := client.V1CloudAccountsTencentList(params)
	if err != nil {
		return nil, err
	}

	accounts := make([]*models.V1TencentAccount, len(response.Payload.Items))
	copy(accounts, response.Payload.Items)

	return accounts, nil
}
