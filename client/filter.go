package client

import (
	"fmt"

	hashboardC "github.com/spectrocloud/hapi/hashboard/client/v1"
	"github.com/spectrocloud/hapi/models"
)

func (h *V1Client) CreateTagFilter(body *models.V1TagFilter) (*models.V1UID, error) {
	client, err := h.GetHashboardClient()
	if err != nil {
		return nil, err
	}

	params := hashboardC.NewV1TagFiltersCreateParams().WithBody(body)
	tag, err := client.V1TagFiltersCreate(params)
	if err != nil {
		return nil, err
	}

	return tag.Payload, nil
}

func (h *V1Client) UpdateTagFilter(uid string, body *models.V1TagFilter) error {
	client, err := h.GetHashboardClient()
	if err != nil {
		return err
	}

	params := hashboardC.NewV1TagFilterUIDUpdateParams().WithUID(uid).WithBody(body)
	_, err = client.V1TagFilterUIDUpdate(params)
	if err != nil {
		return err
	}

	return nil
}

func (h *V1Client) GetTagFilter(uid string) (*models.V1TagFilterSummary, error) {
	client, err := h.GetHashboardClient()
	if err != nil {
		return nil, err
	}

	params := hashboardC.NewV1TagFilterUIDGetParams().WithUID(uid)
	success, err := client.V1TagFilterUIDGet(params)
	if err != nil {
		return nil, err
	}

	return success.Payload, nil
}

func (h *V1Client) ListTagFilters() (*models.V1FiltersSummary, error) {
	client, err := h.GetHashboardClient()
	if err != nil {
		return nil, err
	}

	limit := int64(0)
	params := hashboardC.NewV1FiltersListParams().WithLimit(&limit)
	success, err := client.V1FiltersList(params)
	if err != nil {
		return nil, err
	}

	return success.Payload, nil
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
	client, err := h.GetHashboardClient()
	if err != nil {
		return err
	}

	params := hashboardC.NewV1TagFilterUIDDeleteParams().WithUID(uid)
	_, err = client.V1TagFilterUIDDelete(params)
	if err != nil {
		return err
	}

	return nil
}
