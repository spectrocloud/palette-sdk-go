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

func getClusterFilter(extraFilters []*models.V1SearchFilterItem) *models.V1SearchFilterSpec {
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
								Negation: true,
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

func (h *V1Client) SearchClusterSummariesNonVirtual(filter *models.V1SearchFilterSpec, sort []*models.V1SearchFilterSortSpec) ([]*models.V1SpectroClusterSummary, error) {
	return h.SearchClusterSummaries("tenant", filter, sort) // regular clusters search is working only in tenant context.
}

func (h *V1Client) GetClusterByName(name, clusterContext string) (*models.V1SpectroCluster, error) {
	filters := []*models.V1SearchFilterItem{clusterNameEqFilter(name)}
	clusterSummaries, err := h.SearchClusterSummariesNonVirtual(getClusterFilter(filters), nil)
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
