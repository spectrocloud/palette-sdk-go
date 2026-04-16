package profiles

import "github.com/spectrocloud/palette-sdk-go/api/models"

func (s *service) GetPacks(filters []string, registryUID string) ([]*models.V1PackSummary, error) {
	return s.client.GetPacks(filters, registryUID)
}

func (s *service) SearchPacks(filter *models.V1PackFilterSpec, sortBy []*models.V1PackSortSpec) ([]*models.V1PackMetadata, error) {
	return s.client.SearchPacks(filter, sortBy)
}
