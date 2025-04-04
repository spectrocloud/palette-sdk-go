package client

import (
	"fmt"

	clientv1 "github.com/spectrocloud/palette-sdk-go/api/client/version1"
	"github.com/spectrocloud/palette-sdk-go/api/models"
	"github.com/spectrocloud/palette-sdk-go/client/apiutil"
	"github.com/spectrocloud/palette-sdk-go/client/herr"
)

// TODO: edgev1 deprecation

// SearchApplianceSummaries retrieves a list of edge host summaries based on the provided filter and sort criteria.
func (h *V1Client) SearchApplianceSummaries(filter *models.V1SearchFilterSpec, sort []*models.V1SearchFilterSortSpec) ([]*models.V1EdgeHostsMetadata, error) {
	params := clientv1.NewV1DashboardEdgehostsSearchParamsWithContext(h.ctx).
		WithBody(&models.V1SearchFilterSummarySpec{
			Filter: filter,
			Sort:   sort,
		})
	isContinue := true
	var hosts []*models.V1EdgeHostsMetadata
	for isContinue {
		resp, err := h.Client.V1DashboardEdgehostsSearch(params)
		if err != nil {
			return nil, err
		}
		hosts = append(hosts, resp.Payload.Items...)
		if resp.Payload.Listmeta.Continue == "" {
			isContinue = false
		}
		params.WithContinue(apiutil.Ptr(resp.Payload.Listmeta.Continue))
	}
	return hosts, nil
}

// GetAppliances returns a list of edge host metadata, filtered by the provided tags, status, health, and architecture.
func (h *V1Client) GetAppliances(tags map[string]string, status, health, architecture string) ([]*models.V1EdgeHostsMetadata, error) {
	filter := getApplianceFilter(nil, tags, status, health, architecture)
	appliances, err := h.SearchApplianceSummaries(filter, nil)
	if err != nil {
		return nil, err
	}
	return appliances, nil
}

// GetAppliance retrieves an existing edge host by UID.
func (h *V1Client) GetAppliance(uid string) (*models.V1EdgeHostDevice, error) {
	params := clientv1.NewV1EdgeHostDevicesUIDGetParamsWithContext(h.ctx).
		WithUID(uid)
	resp, err := h.Client.V1EdgeHostDevicesUIDGet(params)
	if err != nil {
		if herr.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}
	return resp.Payload, nil
}

// GetApplianceByName retrieves an existing edge host by name.
func (h *V1Client) GetApplianceByName(name string, tags map[string]string, status, health, architecture string) (*models.V1EdgeHostDevice, error) {
	filters := []*models.V1SearchFilterItem{applianceNameEqFilter(name)}
	applianceSummaries, err := h.SearchApplianceSummaries(getApplianceFilter(filters, tags, status, health, architecture), nil)
	if err != nil {
		return nil, err
	}
	if len(applianceSummaries) != 1 {
		return nil, fmt.Errorf("expected 1 appliance, got %d", len(applianceSummaries))
	}
	appliance, err := h.GetAppliance(applianceSummaries[0].Metadata.UID)
	if err != nil {
		return nil, err
	}
	return appliance, nil
}

// CreateAppliance creates a new edge host.
func (h *V1Client) CreateAppliance(createHostDevice *models.V1EdgeHostDeviceEntity) (string, error) {
	params := clientv1.NewV1EdgeHostDevicesCreateParamsWithContext(h.ctx).
		WithBody(createHostDevice)
	resp, err := h.Client.V1EdgeHostDevicesCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

// UpdateAppliance updates an existing edge host.
func (h *V1Client) UpdateAppliance(uid string, appliance *models.V1EdgeHostDevice) error {
	params := clientv1.NewV1EdgeHostDevicesUIDUpdateParamsWithContext(h.ctx).
		WithBody(appliance).
		WithUID(uid)
	_, err := h.Client.V1EdgeHostDevicesUIDUpdate(params)
	if err != nil && !herr.IsEdgeHostDeviceNotRegistered(err) {
		return err
	}
	return nil
}

// UpdateApplianceMeta updates the metadata of an existing edge host.
func (h *V1Client) UpdateApplianceMeta(uid string, appliance *models.V1EdgeHostDeviceMetaUpdateEntity) error {
	params := clientv1.NewV1EdgeHostDevicesUIDMetaUpdateParamsWithContext(h.ctx).
		WithBody(appliance).
		WithUID(uid)
	_, err := h.Client.V1EdgeHostDevicesUIDMetaUpdate(params)
	return err
}

// DeleteAppliance deletes an existing edge host.
func (h *V1Client) DeleteAppliance(uid string) error {
	params := clientv1.NewV1EdgeHostDevicesUIDDeleteParamsWithContext(h.ctx).
		WithUID(uid)
	_, err := h.Client.V1EdgeHostDevicesUIDDelete(params)
	return err
}

func getApplianceFilter(extraFilters []*models.V1SearchFilterItem, tags map[string]string, state, healthState, architecture string) *models.V1SearchFilterSpec {
	filter := &models.V1SearchFilterSpec{
		Conjunction: and(),
		FilterGroups: []*models.V1SearchFilterGroup{
			{
				Conjunction: and(),
				Filters:     []*models.V1SearchFilterItem{},
			},
		},
	}

	// Tags filter
	if len(tags) > 0 {
		var tagValues []string
		for key, value := range tags {
			tagValues = append(tagValues, fmt.Sprintf("%s:%s", key, value))
		}
		tagsFilter := &models.V1SearchFilterItem{
			Condition: &models.V1SearchFilterCondition{
				String: &models.V1SearchFilterStringCondition{
					Match: &models.V1SearchFilterStringConditionMatch{
						Conjunction: or(),
						Values:      tagValues,
					},
					Operator:   models.V1SearchFilterStringOperatorEq,
					Negation:   false,
					IgnoreCase: false,
				},
			},
			Property: "tags",
			Type:     models.V1SearchFilterPropertyTypeString,
		}
		filter.FilterGroups[0].Filters = append(filter.FilterGroups[0].Filters, tagsFilter)
	}

	// State filter
	if state != "" {
		stateFilter := &models.V1SearchFilterItem{
			Condition: &models.V1SearchFilterCondition{
				String: &models.V1SearchFilterStringCondition{
					Match: &models.V1SearchFilterStringConditionMatch{
						Conjunction: or(),
						Values:      []string{state},
					},
					Operator: models.V1SearchFilterStringOperatorEq,
				},
			},
			Property: "state",
			Type:     models.V1SearchFilterPropertyTypeString,
		}
		filter.FilterGroups[0].Filters = append(filter.FilterGroups[0].Filters, stateFilter)
	}

	// Health state filter
	if healthState != "" {
		healthStateFilter := &models.V1SearchFilterItem{
			Condition: &models.V1SearchFilterCondition{
				String: &models.V1SearchFilterStringCondition{
					Match: &models.V1SearchFilterStringConditionMatch{
						Conjunction: or(),
						Values:      []string{healthState},
					},
					Operator: models.V1SearchFilterStringOperatorEq,
				},
			},
			Property: "healthState",
			Type:     models.V1SearchFilterPropertyTypeString,
		}
		filter.FilterGroups[0].Filters = append(filter.FilterGroups[0].Filters, healthStateFilter)
	}

	// Architecture filter
	if architecture != "" {
		architectureFilter := &models.V1SearchFilterItem{
			Condition: &models.V1SearchFilterCondition{
				String: &models.V1SearchFilterStringCondition{
					Match: &models.V1SearchFilterStringConditionMatch{
						Conjunction: or(),
						Values:      []string{architecture},
					},
					Operator: models.V1SearchFilterStringOperatorEq,
				},
			},
			Property: "architecture",
			Type:     models.V1SearchFilterPropertyTypeString,
		}
		filter.FilterGroups[0].Filters = append(filter.FilterGroups[0].Filters, architectureFilter)
	}

	// Append extra filters if provided
	if extraFilters != nil {
		filter.FilterGroups = append(filter.FilterGroups, &models.V1SearchFilterGroup{Conjunction: and(), Filters: extraFilters})
	}

	return filter
}

func applianceNameEqFilter(name string) *models.V1SearchFilterItem {
	return &models.V1SearchFilterItem{
		Condition: &models.V1SearchFilterCondition{
			String: &models.V1SearchFilterStringCondition{
				Match: &models.V1SearchFilterStringConditionMatch{
					Conjunction: or(),
					Values:      []string{name},
				},
				Operator:   models.V1SearchFilterStringOperatorEq,
				Negation:   false,
				IgnoreCase: false,
			},
		},
		Property: "name",
		Type:     models.V1SearchFilterPropertyTypeString,
	}
}
