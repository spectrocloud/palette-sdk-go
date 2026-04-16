package diagnostics

import (
	"errors"

	"github.com/spectrocloud/palette-sdk-go/api/models"
)

func (s *service) CreateTagFilter(body *models.V1TagFilter) (*models.V1UID, error) {
	if body == nil {
		return nil, errors.New("tag filter body is required")
	}
	return s.client.CreateTagFilter(body)
}

func (s *service) GetTagFilter(uid string) (*models.V1TagFilterSummary, error) {
	if uid == "" {
		return nil, errUIDRequired
	}
	return s.client.GetTagFilter(uid)
}

func (s *service) GetTagFilterByName(name string) (*models.V1FilterSummary, error) {
	if name == "" {
		return nil, errors.New("filter name is required")
	}
	return s.client.GetTagFilterByName(name)
}

func (s *service) ListTagFilters() (*models.V1FiltersSummary, error) {
	return s.client.ListTagFilters()
}

func (s *service) UpdateTagFilter(uid string, body *models.V1TagFilter) error {
	if uid == "" {
		return errUIDRequired
	}
	return s.client.UpdateTagFilter(uid, body)
}

func (s *service) DeleteTagFilter(uid string) error {
	if uid == "" {
		return errUIDRequired
	}
	return s.client.DeleteTagFilter(uid)
}
