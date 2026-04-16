package cloudaccounts

import "github.com/spectrocloud/palette-sdk-go/api/models"

func (s *service) List() ([]*models.V1CloudAccountSummary, error) {
	return s.client.ListCloudAccounts()
}

func (s *service) Get(uid string) (*models.V1CloudAccountSummary, error) {
	if uid == "" {
		return nil, errUIDRequired
	}
	return s.client.GetCloudAccount(uid)
}
