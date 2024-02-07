package client

import (
	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
	"github.com/spectrocloud/palette-sdk-go/client/herr"
)

func (h *V1Client) SearchApplianceSummaries(applianceContext string, filter *models.V1SearchFilterSpec, sort []*models.V1SearchFilterSortSpec) ([]*models.V1EdgeHostsMetadata, error) {
	var params *clientV1.V1DashboardEdgehostsSearchParams
	switch applianceContext {
	case "project":
		params = clientV1.NewV1DashboardEdgehostsSearchParamsWithContext(h.Ctx)
	case "tenant":
		params = clientV1.NewV1DashboardEdgehostsSearchParams()
	}
	params.Body = &models.V1SearchFilterSummarySpec{
		Filter: filter,
		Sort:   sort,
	}

	resp, err := h.Client.V1DashboardEdgehostsSearch(params)
	if err != nil {
		return nil, err
	}

	return resp.Payload.Items, nil
}

func (h *V1Client) GetAppliances(applianceContext string, tags map[string]string, status, health, architecture string) ([]*models.V1EdgeHostsMetadata, error) {
	filter := getApplianceFilter(nil, tags, status, health, architecture)

	appliances, err := h.SearchApplianceSummaries(applianceContext, filter, nil)
	if err != nil {
		return nil, err
	}

	return appliances, nil
}

func (h *V1Client) GetAppliance(uid string) (*models.V1EdgeHostDevice, error) {
	params := clientV1.NewV1EdgeHostDevicesUIDGetParamsWithContext(h.Ctx).WithUID(uid)
	response, err := h.Client.V1EdgeHostDevicesUIDGet(params)
	if err != nil {
		if herr.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}

	return response.Payload, nil
}

func (h *V1Client) CreateAppliance(createHostDevice *models.V1EdgeHostDeviceEntity) (string, error) {
	params := clientV1.NewV1EdgeHostDevicesCreateParams().WithContext(h.Ctx).WithBody(createHostDevice)
	if resp, err := h.Client.V1EdgeHostDevicesCreate(params); err != nil {
		return "", err
	} else {
		return *resp.Payload.UID, nil
	}
}

func (h *V1Client) UpdateAppliance(uid string, appliance *models.V1EdgeHostDevice) error {
	params := clientV1.NewV1EdgeHostDevicesUIDUpdateParams().WithContext(h.Ctx).WithBody(appliance).WithUID(uid)
	_, err := h.Client.V1EdgeHostDevicesUIDUpdate(params)
	if err != nil && !herr.IsEdgeHostDeviceNotRegistered(err) {
		return err
	}

	return nil
}

func (h *V1Client) UpdateApplianceMeta(uid string, appliance *models.V1EdgeHostDeviceMetaUpdateEntity) error {
	params := clientV1.NewV1EdgeHostDevicesUIDMetaUpdateParams().WithContext(h.Ctx).WithBody(appliance).WithUID(uid)
	_, err := h.Client.V1EdgeHostDevicesUIDMetaUpdate(params)
	if err != nil {
		return err
	}

	return nil
}

func (h *V1Client) DeleteAppliance(uid string) error {
	params := clientV1.NewV1EdgeHostDevicesUIDDeleteParams().WithContext(h.Ctx).WithUID(uid)
	_, err := h.Client.V1EdgeHostDevicesUIDDelete(params)
	if err != nil {
		return err
	}

	return nil
}
