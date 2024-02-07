package client

import (
	"fmt"

	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
	"github.com/spectrocloud/palette-sdk-go/client/apiutil"
)

func (h *V1Client) ListCloudAccounts() ([]*models.V1CloudAccountSummary, error) {
	params := clientV1.NewV1CloudAccountsListSummaryParamsWithContext(h.ctx).
		WithLimit(apiutil.Ptr(int64(0)))
	resp, err := h.Client.V1CloudAccountsListSummary(params)
	if err := apiutil.Handle404(err); err != nil {
		return nil, err
	}
	return resp.Payload.Items, nil
}

func (h *V1Client) GetCloudAccount(uid string) (*models.V1CloudAccountSummary, error) {
	accounts, err := h.ListCloudAccounts()
	if err != nil {
		return nil, err
	}
	for _, account := range accounts {
		if account.Metadata.UID == uid {
			return account, nil
		}
	}
	return nil, fmt.Errorf("account not found with uid %s", uid)
}
