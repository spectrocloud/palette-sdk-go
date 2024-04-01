package client

import (
	"fmt"

	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
	"github.com/spectrocloud/palette-sdk-go/client/apiutil"
)

func (h *V1Client) CreateTagFilter(body *models.V1TagFilter) (*models.V1UID, error) {
	params := clientV1.NewV1TagFiltersCreateParamsWithContext(h.ctx).
		WithBody(body)
	resp, err := h.Client.V1TagFiltersCreate(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

func (h *V1Client) UpdateTagFilter(uid string, body *models.V1TagFilter) error {
	params := clientV1.NewV1TagFilterUIDUpdateParamsWithContext(h.ctx).
		WithUID(uid).
		WithBody(body)
	_, err := h.Client.V1TagFilterUIDUpdate(params)
	return err
}

func (h *V1Client) GetTagFilter(uid string) (*models.V1TagFilterSummary, error) {
	params := clientV1.NewV1TagFilterUIDGetParamsWithContext(h.ctx).
		WithUID(uid)
	resp, err := h.Client.V1TagFilterUIDGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

func (h *V1Client) ListTagFilters() (*models.V1FiltersSummary, error) {
	params := clientV1.NewV1FiltersListParamsWithContext(h.ctx).
		WithLimit(apiutil.Ptr(int64(0)))
	resp, err := h.Client.V1FiltersList(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

func (h *V1Client) GetTagFilterByName(name string) (*models.V1FilterSummary, error) {
	filters, err := h.ListTagFilters()
	if err != nil {
		return nil, err
	}
	for _, filter := range filters.Items {
		if filter.Metadata.Name == name {
			return filter, nil
		}
	}
	return nil, fmt.Errorf("filter not found for name %s", name)
}

func (h *V1Client) DeleteTag(uid string) error {
	params := clientV1.NewV1TagFilterUIDDeleteParamsWithContext(h.ctx).
		WithUID(uid)
	_, err := h.Client.V1TagFilterUIDDelete(params)
	return err
}
