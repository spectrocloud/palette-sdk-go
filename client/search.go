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

	filter.FilterGroups[0].Filters = append(filter.FilterGroups[0].Filters, extraFilters...)

	return filter
}

func nameFilter(name string) *models.V1SearchFilterItem {
	return &models.V1SearchFilterItem{
		Condition: &models.V1SearchFilterCondition{
			String: &models.V1SearchFilterStringCondition{
				Match: &models.V1SearchFilterStringConditionMatch{
					Conjunction: or(),
					Values:      []string{name},
				},
				Operator: models.V1SearchFilterStringOperatorEq,
			},
		},
		Property: "name",
		Type:     models.V1SearchFilterPropertyTypeString,
	}
}

func (h *V1Client) GetClusterByName(name, clusterContext string) (*models.V1SpectroCluster, error) {
	filters := []*models.V1SearchFilterItem{nameFilter(name)}
	clusterSummaries, err := h.SearchClusterSummaries(
		clusterContext, getClusterFilter(filters), nil,
	)
	if err != nil {
		return nil, err
	}
	if len(clusterSummaries) != 1 {
		return nil, fmt.Errorf("expected 1 cluster: %v in %v, got %d", name, clusterContext, len(clusterSummaries))
	}

	cluster, err := h.GetCluster(clusterContext, clusterSummaries[0].Metadata.UID)
	if err != nil {
		return nil, err
	}
	return cluster, nil
}
