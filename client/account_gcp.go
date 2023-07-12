package client

import (
	"github.com/spectrocloud/hapi/apiutil/transport"
	"github.com/spectrocloud/hapi/models"
	clusterC "github.com/spectrocloud/hapi/spectrocluster/client/v1"
)

func (h *V1Client) CreateCloudAccountGcp(account *models.V1GcpAccountEntity) (string, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return "", err
	}

	params := clusterC.NewV1CloudAccountsGcpCreateParamsWithContext(h.Ctx).WithBody(account)
	success, err := client.V1CloudAccountsGcpCreate(params)
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

func (h *V1Client) UpdateCloudAccountGcp(account *models.V1GcpAccountEntity) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil
	}

	uid := account.Metadata.UID
	params := clusterC.NewV1CloudAccountsGcpUpdateParamsWithContext(h.Ctx).WithUID(uid).WithBody(account)
	_, err = client.V1CloudAccountsGcpUpdate(params)
	return err
}

func (h *V1Client) DeleteCloudAccountGcp(uid string) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil
	}

	params := clusterC.NewV1CloudAccountsGcpDeleteParamsWithContext(h.Ctx).WithUID(uid)
	_, err = client.V1CloudAccountsGcpDelete(params)
	return err
}

func (h *V1Client) GetCloudAccountGcp(uid string) (*models.V1GcpAccount, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}

	params := clusterC.NewV1CloudAccountsGcpGetParamsWithContext(h.Ctx).WithUID(uid)
	success, err := client.V1CloudAccountsGcpGet(params)
	if e, ok := err.(*transport.TransportError); ok && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return success.Payload, nil
}

func (h *V1Client) GetCloudAccountsGcp() ([]*models.V1GcpAccount, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}

	limit := int64(0)
	params := clusterC.NewV1CloudAccountsGcpListParamsWithContext(h.Ctx).WithLimit(&limit)
	response, err := client.V1CloudAccountsGcpList(params)
	if err != nil {
		return nil, err
	}

	accounts := make([]*models.V1GcpAccount, len(response.Payload.Items))
	copy(accounts, response.Payload.Items)

	return accounts, nil
}
