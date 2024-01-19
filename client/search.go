package client

import (
	"fmt"

	"github.com/spectrocloud/hapi/models"
)

func and() *models.V1SearchFilterConjunctionOperator {
	and := models.V1SearchFilterConjunctionOperatorAnd
	return &and
}

func or() *models.V1SearchFilterConjunctionOperator {
	or := models.V1SearchFilterConjunctionOperatorOr
	return &or
}

func getClusterFilter(extraFilters []*models.V1SearchFilterItem, virtual bool) *models.V1SearchFilterSpec {
	filter := &models.V1SearchFilterSpec{
		Conjunction: and(),
		FilterGroups: []*models.V1SearchFilterGroup{
			{
				Conjunction: and(),
				Filters: []*models.V1SearchFilterItem{
					{
						Condition: &models.V1SearchFilterCondition{
							String: &models.V1SearchFilterStringCondition{
								Match: &models.V1SearchFilterStringConditionMatch{
									Conjunction: or(),
									Values:      []string{"nested"},
								},
								Negation: !virtual,
								Operator: models.V1SearchFilterStringOperatorEq,
							},
						},
						Property: "environment",
						Type:     models.V1SearchFilterPropertyTypeString,
					},
					{
						Condition: &models.V1SearchFilterCondition{
							Bool: &models.V1SearchFilterBoolCondition{
								Value: false,
							},
						},
						Property: "isDeleted",
						Type:     models.V1SearchFilterPropertyTypeBool,
					},
				},
			},
		},
	}

	filter.FilterGroups = append(filter.FilterGroups, &models.V1SearchFilterGroup{Conjunction: and(), Filters: extraFilters})

	return filter
}

func clusterNameEqFilter(name string) *models.V1SearchFilterItem {
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
		Property: "clusterName",
		Type:     models.V1SearchFilterPropertyTypeString,
	}
}

func (h *V1Client) GetClusterByName(name, clusterContext string, virtual bool) (*models.V1SpectroCluster, error) {
	filters := []*models.V1SearchFilterItem{clusterNameEqFilter(name)}
	var clusterSummaries []*models.V1SpectroClusterSummary
	var err error
	if virtual {
		clusterSummaries, err = h.SearchClusterSummaries(clusterContext, getClusterFilter(filters, virtual), nil)
	} else {
		clusterSummaries, err = h.SearchClusterSummaries("tenant", getClusterFilter(filters, virtual), nil) // regular clusters search is working only in tenant context.
	}
	if err != nil {
		return nil, err
	}
	if len(clusterSummaries) != 1 {
		return nil, fmt.Errorf("expected 1 cluster: %v in %v, got %d", name, clusterContext, len(clusterSummaries))
	}

	cluster, err := h.GetCluster(clusterSummaries[0].Metadata.Annotations["scope"], clusterSummaries[0].Metadata.UID)
	if err != nil {
		return nil, err
	}
	return cluster, nil
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

func (h *V1Client) GetApplianceByName(name, applianceContext string, tags map[string]string, status, health, architecture string) (*models.V1EdgeHostDevice, error) {
	filters := []*models.V1SearchFilterItem{applianceNameEqFilter(name)}
	applianceSummaries, err := h.SearchApplianceSummaries(applianceContext, getApplianceFilter(filters, tags, status, health, architecture), nil)
	if err != nil {
		return nil, err
	}
	if len(applianceSummaries) != 1 {
		return nil, fmt.Errorf("expected 1 appliance: %v in %v, got %d", name, applianceContext, len(applianceSummaries))
	}

	appliance, err := h.GetAppliance(applianceSummaries[0].Metadata.UID)
	if err != nil {
		return nil, err
	}
	return appliance, nil
}
