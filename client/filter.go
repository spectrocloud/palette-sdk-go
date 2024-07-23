package client

import (
	"fmt"

	clientv1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
	"github.com/spectrocloud/palette-sdk-go/client/apiutil"
)

// CRUDL operations on tag filters are all tenant scoped.
// See: hubble/services/svccore/perms/filter_acl.go

// CreateTagFilter creates a new tag filter.
func (h *V1Client) CreateTagFilter(body *models.V1TagFilter) (*models.V1UID, error) {
	// ACL scoped to tenant only
	params := clientv1.NewV1TagFiltersCreateParams().
		WithBody(body)
	resp, err := h.Client.V1TagFiltersCreate(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// UpdateTagFilter updates an existing tag filter.
func (h *V1Client) UpdateTagFilter(uid string, body *models.V1TagFilter) error {
	// ACL scoped to tenant only
	params := clientv1.NewV1TagFilterUIDUpdateParams().
		WithUID(uid).
		WithBody(body)
	_, err := h.Client.V1TagFilterUIDUpdate(params)
	return err
}

// GetTagFilter retrieves an existing tag filter by UID.
func (h *V1Client) GetTagFilter(uid string) (*models.V1TagFilterSummary, error) {
	// ACL scoped to tenant only
	params := clientv1.NewV1TagFilterUIDGetParams().
		WithUID(uid)
	resp, err := h.Client.V1TagFilterUIDGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// ListTagFilters retrieves all tag filters.
func (h *V1Client) ListTagFilters() (*models.V1FiltersSummary, error) {
	// ACL scoped to tenant only
	params := clientv1.NewV1FiltersListParams().
		WithLimit(apiutil.Ptr(int64(0)))
	resp, err := h.Client.V1FiltersList(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// GetTagFilterByName retrieves an existing tag filter by name.
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

// DeleteTagFilter deletes an existing tag filter.
func (h *V1Client) DeleteTagFilter(uid string) error {
	// ACL scoped to tenant only
	params := clientv1.NewV1TagFilterUIDDeleteParams().
		WithUID(uid)
	_, err := h.Client.V1TagFilterUIDDelete(params)
	return err
}
