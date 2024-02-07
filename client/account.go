package client

import (
	"errors"
	"fmt"

	"github.com/spectrocloud/palette-api-go/apiutil/transport"
	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
)

func (h *V1Client) ListCloudAccounts(scope string) ([]*models.V1CloudAccountSummary, error) {
	var params *clientV1.V1CloudAccountsListSummaryParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudAccountsListSummaryParams().WithContext(h.Ctx)
	case "tenant":
		params = clientV1.NewV1CloudAccountsListSummaryParams()

	}
	var limit int64 = 0
	params.Limit = &limit
	resp, err := h.Client.V1CloudAccountsListSummary(params)

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
