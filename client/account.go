package client

import (
	"fmt"

	clientv1 "github.com/spectrocloud/palette-sdk-go/api/client/version1"
	"github.com/spectrocloud/palette-sdk-go/api/models"
	"github.com/spectrocloud/palette-sdk-go/client/apiutil"
)

// ListCloudAccounts returns a list of all cloud account summaries.
func (h *V1Client) ListCloudAccounts() ([]*models.V1CloudAccountSummary, error) {
	params := clientv1.NewV1CloudAccountsListSummaryParamsWithContext(h.ctx).
		WithLimit(apiutil.Ptr(int64(0)))
	resp, err := h.Client.V1CloudAccountsListSummary(params)
	if apiutil.Is404(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return resp.Payload.Items, nil
}

// GetCloudAccount retrieves an existing cloud account summary.
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
