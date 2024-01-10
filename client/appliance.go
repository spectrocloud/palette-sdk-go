package client

import (
	"errors"

	"github.com/spectrocloud/hapi/apiutil/transport"
	hashboardC "github.com/spectrocloud/hapi/hashboard/client/v1"
	"github.com/spectrocloud/hapi/models"
	clusterC "github.com/spectrocloud/hapi/spectrocluster/client/v1"

	"github.com/spectrocloud/palette-sdk-go/client/herr"
)

func (h *V1Client) SearchApplianceSummaries(applianceContext string, filter *models.V1SearchFilterSpec, sort []*models.V1SearchFilterSortSpec) ([]*models.V1EdgeHostsMetadata, error) {
	var params *hashboardC.V1DashboardEdgehostsSearchParams
	switch applianceContext {
	case "project":
		params = hashboardC.NewV1DashboardEdgehostsSearchParamsWithContext(h.Ctx)
	case "tenant":
		params = hashboardC.NewV1DashboardEdgehostsSearchParams()
	}
	params.Body = &models.V1SearchFilterSummarySpec{
		Filter: filter,
		Sort:   sort,
	}

	resp, err := h.GetHashboardClient().V1DashboardEdgehostsSearch(params)
	var e *transport.TransportError
	if errors.As(err, &e) && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
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
	params := clusterC.NewV1EdgeHostDevicesUIDGetParamsWithContext(h.Ctx).WithUID(uid)
	response, err := h.GetClusterClient().V1EdgeHostDevicesUIDGet(params)
	if err != nil {
		if herr.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}

	return response.Payload, nil
}

func (h *V1Client) CreateAppliance(createHostDevice *models.V1EdgeHostDeviceEntity) (string, error) {
	params := clusterC.NewV1EdgeHostDevicesCreateParams().WithContext(h.Ctx).WithBody(createHostDevice)
	if resp, err := h.GetClusterClient().V1EdgeHostDevicesCreate(params); err != nil {
		return "", err
	} else {
		return *resp.Payload.UID, nil
	}
}

func (h *V1Client) UpdateAppliance(uid string, appliance *models.V1EdgeHostDevice) error {
	params := clusterC.NewV1EdgeHostDevicesUIDUpdateParams().WithContext(h.Ctx).WithBody(appliance).WithUID(uid)
	_, err := h.GetClusterClient().V1EdgeHostDevicesUIDUpdate(params)
	if err != nil && !herr.IsEdgeHostDeviceNotRegistered(err) {
		return err
	}

	return nil
}

func (h *V1Client) UpdateApplianceMeta(uid string, appliance *models.V1EdgeHostDeviceMetaUpdateEntity) error {
	params := clusterC.NewV1EdgeHostDevicesUIDMetaUpdateParams().WithContext(h.Ctx).WithBody(appliance).WithUID(uid)
	_, err := h.GetClusterClient().V1EdgeHostDevicesUIDMetaUpdate(params)
	if err != nil {
		return err
	}

	return nil
}

func (h *V1Client) DeleteAppliance(uid string) error {
	params := clusterC.NewV1EdgeHostDevicesUIDDeleteParams().WithContext(h.Ctx).WithUID(uid)
	_, err := h.GetClusterClient().V1EdgeHostDevicesUIDDelete(params)
	if err != nil {
		return err
	}

	return nil
}
