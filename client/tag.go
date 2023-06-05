package client

import (
	hashboardC "github.com/spectrocloud/hapi/hashboard/client/v1"
	"github.com/spectrocloud/hapi/models"
)

func (h *V1Client) CreateTag(body *models.V1TagFilter) (*models.V1UID, error) {
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

func (h *V1Client) UpdateTag(body *models.V1TagFilter) error {
	client, err := h.GetHashboardClient()
	if err != nil {
		return err
	}

	params := hashboardC.NewV1TagFilterUIDUpdateParams().WithBody(body)
	_, err = client.V1TagFilterUIDUpdate(params)
	if err != nil {
		return err
	}

	return nil
}

func (h *V1Client) GetTag(uid string) (*models.V1TagFilterSummary, error) {
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
