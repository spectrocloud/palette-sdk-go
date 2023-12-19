package client

import (
	"errors"
	"fmt"

	"github.com/spectrocloud/hapi/apiutil/transport"
	"github.com/spectrocloud/hapi/models"
	clusterC "github.com/spectrocloud/hapi/spectrocluster/client/v1"
)

func (h *V1Client) ListCloudAccounts(scope string) ([]*models.V1CloudAccountSummary, error) {
	var params *clusterC.V1CloudAccountsListSummaryParams
	switch scope {
	case "project":
		params = clusterC.NewV1CloudAccountsListSummaryParams().WithContext(h.Ctx)
	case "tenant":
		params = clusterC.NewV1CloudAccountsListSummaryParams()

	}
	var limit int64 = 0
	params.Limit = &limit
	resp, err := h.GetClusterClient().V1CloudAccountsListSummary(params)

	var e *transport.TransportError
	if errors.As(err, &e) && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return resp.Payload.Items, nil
}

func (h *V1Client) GetCloudAccount(scope, id string) (*models.V1CloudAccountSummary, error) {
	accounts, err := h.ListCloudAccounts(scope)
	if err != nil {
		return nil, err
	}

	for _, account := range accounts {
		if account.Metadata.UID == id {
			return account, nil
		}
	}

	return nil, fmt.Errorf("account not found with id %s", id)
}
