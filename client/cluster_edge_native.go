package client

import (
	"errors"
	"fmt"

	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
	"github.com/spectrocloud/palette-sdk-go/client/apiutil"
)

// CRUDL operations on edge registration tokens are all tenant scoped.
// See: hubble/services/svccore/perms/edgetoken_acl.go

func (h *V1Client) GetRegistrationToken(tokenName string) (string, error) {
	// ACL scoped to tenant only
	params := clientV1.NewV1EdgeTokensListParams()
	resp, err := h.Client.V1EdgeTokensList(params)
	if err != nil {
		return "", err
	}
	tokens := resp.GetPayload()
	if tokens == nil {
		return "", errors.New("failed to list registration tokens")
	}
	for _, token := range tokens.Items {
		if token.Status.IsActive && token.Metadata.Name == tokenName {
			return token.Spec.Token, nil
		}
	}
	return "", nil
}

func (h *V1Client) CreateRegistrationToken(tokenName string, body *models.V1EdgeTokenEntity) (string, error) {
	// ACL scoped to tenant only
	params := clientV1.NewV1EdgeTokensCreateParams().
		WithBody(body)
	_, err := h.Client.V1EdgeTokensCreate(params)
	if err != nil {
		return "", err
	}
	return h.GetRegistrationToken(tokenName)

}

func (h *V1Client) GetEdgeHost(edgeHostId string) (*models.V1EdgeHostDevice, error) {
	params := clientV1.NewV1EdgeHostDevicesUIDGetParamsWithContext(h.ctx).
		WithUID(edgeHostId)
	resp, err := h.Client.V1EdgeHostDevicesUIDGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

func (h *V1Client) ListEdgeHosts() ([]*models.V1EdgeHostsMetadata, error) {
	continueToken := ""
	var items []*models.V1EdgeHostsMetadata
	for ok := true; ok; ok = (continueToken != "") {
		params := clientV1.NewV1DashboardEdgehostsSearchParamsWithContext(h.ctx)
		resp, err := h.Client.V1DashboardEdgehostsSearch(params)
		if err != nil {
			return nil, err
		}
		continueToken = resp.Payload.Listmeta.Continue
		items = append(items, resp.Payload.Items...)
	}

	return items, nil
}

func (h *V1Client) GetEdgeHostsByTags(tags map[string]string) ([]*models.V1EdgeHostsMetadata, error) {
	continueToken := ""
	var items []*models.V1EdgeHostsMetadata
	filter := getEdgeFilter(nil, tags)
	for ok := true; ok; ok = (continueToken != "") {
		params := clientV1.NewV1DashboardEdgehostsSearchParamsWithContext(h.ctx).
		WithBody(&models.V1SearchFilterSummarySpec{
			Filter: filter,
			Sort:   nil,
		})
		resp, err := h.Client.V1DashboardEdgehostsSearch(params)
		if err != nil {
			return nil, err
		}
		continueToken = resp.Payload.Listmeta.Continue
		items = append(items, resp.Payload.Items...)
	}

	return items, nil
}

func getEdgeFilter(extraFilters []*models.V1SearchFilterItem, tags map[string]string) *models.V1SearchFilterSpec {
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

	// // State filter
	// if state != "" {
	// 	stateFilter := &models.V1SearchFilterItem{
	// 		Condition: &models.V1SearchFilterCondition{
	// 			String: &models.V1SearchFilterStringCondition{
	// 				Match: &models.V1SearchFilterStringConditionMatch{
	// 					Conjunction: or(),
	// 					Values:      []string{state},
	// 				},
	// 				Operator: models.V1SearchFilterStringOperatorEq,
	// 			},
	// 		},
	// 		Property: "state",
	// 		Type:     models.V1SearchFilterPropertyTypeString,
	// 	}
	// 	filter.FilterGroups[0].Filters = append(filter.FilterGroups[0].Filters, stateFilter)
	// }

	// // Health state filter
	// if healthState != "" {
	// 	healthStateFilter := &models.V1SearchFilterItem{
	// 		Condition: &models.V1SearchFilterCondition{
	// 			String: &models.V1SearchFilterStringCondition{
	// 				Match: &models.V1SearchFilterStringConditionMatch{
	// 					Conjunction: or(),
	// 					Values:      []string{healthState},
	// 				},
	// 				Operator: models.V1SearchFilterStringOperatorEq,
	// 			},
	// 		},
	// 		Property: "healthState",
	// 		Type:     models.V1SearchFilterPropertyTypeString,
	// 	}
	// 	filter.FilterGroups[0].Filters = append(filter.FilterGroups[0].Filters, healthStateFilter)
	// }

	// // Architecture filter
	// if architecture != "" {
	// 	architectureFilter := &models.V1SearchFilterItem{
	// 		Condition: &models.V1SearchFilterCondition{
	// 			String: &models.V1SearchFilterStringCondition{
	// 				Match: &models.V1SearchFilterStringConditionMatch{
	// 					Conjunction: or(),
	// 					Values:      []string{architecture},
	// 				},
	// 				Operator: models.V1SearchFilterStringOperatorEq,
	// 			},
	// 		},
	// 		Property: "architecture",
	// 		Type:     models.V1SearchFilterPropertyTypeString,
	// 	}
	// 	filter.FilterGroups[0].Filters = append(filter.FilterGroups[0].Filters, architectureFilter)
	// }

	// Append extra filters if provided
	if extraFilters != nil {
		filter.FilterGroups = append(filter.FilterGroups, &models.V1SearchFilterGroup{Conjunction: and(), Filters: extraFilters})
	}

	return filter
}

func (h *V1Client) CreateClusterEdgeNative(cluster *models.V1SpectroEdgeNativeClusterEntity) (string, error) {
	params := clientV1.NewV1SpectroClustersEdgeNativeCreateParamsWithContext(h.ctx).
		WithBody(cluster)
	resp, err := h.Client.V1SpectroClustersEdgeNativeCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

func (h *V1Client) CreateMachinePoolEdgeNative(cloudConfigUid string, machinePool *models.V1EdgeNativeMachinePoolConfigEntity) error {
	params := clientV1.NewV1CloudConfigsEdgeNativeMachinePoolCreateParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUid).
		WithBody(machinePool)
	_, err := h.Client.V1CloudConfigsEdgeNativeMachinePoolCreate(params)
	return err
}

func (h *V1Client) UpdateMachinePoolEdgeNative(cloudConfigUid string, machinePool *models.V1EdgeNativeMachinePoolConfigEntity) error {
	params := clientV1.NewV1CloudConfigsEdgeNativeMachinePoolUpdateParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUid).
		WithMachinePoolName(*machinePool.PoolConfig.Name).
		WithBody(machinePool)
	_, err := h.Client.V1CloudConfigsEdgeNativeMachinePoolUpdate(params)
	return err
}

func (h *V1Client) GetMachineEdgeNative(machineUid, machinePoolName, cloudConfigUid string) (*models.V1EdgeNativeMachine, error) {
	params := clientV1.NewV1CloudConfigsEdgeNativePoolMachinesUIDGetParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUid).
		WithMachinePoolName(machinePoolName).
		WithMachineUID(machineUid)
	resp, err := h.Client.V1CloudConfigsEdgeNativePoolMachinesUIDGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

func (h *V1Client) DeleteMachineEdgeNative(clusterUid, edgeHostUid, machinePool, cloudConfigUid string) error {
	params := clientV1.NewV1CloudConfigsEdgeNativePoolMachinesUIDDeleteParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUid).
		WithMachinePoolName(machinePool).
		WithMachineUID(edgeHostUid)
	_, err := h.Client.V1CloudConfigsEdgeNativePoolMachinesUIDDelete(params)
	return err
}

func (h *V1Client) ListMachinesInPoolEdgeNative(machinePoolName, cloudConfigUid string) ([]*models.V1EdgeNativeMachine, error) {
	params := clientV1.NewV1CloudConfigsEdgeNativePoolMachinesListParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUid).
		WithMachinePoolName(machinePoolName)
	resp, err := h.Client.V1CloudConfigsEdgeNativePoolMachinesList(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload.Items, nil
}

func (h *V1Client) DeleteMachinePoolEdgeNative(cloudConfigUid, machinePoolName string) error {
	params := clientV1.NewV1CloudConfigsEdgeNativeMachinePoolDeleteParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUid).
		WithMachinePoolName(machinePoolName)
	_, err := h.Client.V1CloudConfigsEdgeNativeMachinePoolDelete(params)
	return err
}

func (h *V1Client) GetCloudConfigEdgeNative(configUid string) (*models.V1EdgeNativeCloudConfig, error) {
	params := clientV1.NewV1CloudConfigsEdgeNativeGetParamsWithContext(h.ctx).
		WithConfigUID(configUid)
	resp, err := h.Client.V1CloudConfigsEdgeNativeGet(params)
	if apiutil.Is404(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

func (h *V1Client) GetNodeStatusMapEdgeNative(configUid, machinePoolName string) (map[string]models.V1CloudMachineStatus, error) {
	params := clientV1.NewV1CloudConfigsEdgeNativePoolMachinesListParamsWithContext(h.ctx).
		WithConfigUID(configUid).
		WithMachinePoolName(machinePoolName)
	mpList, err := h.Client.V1CloudConfigsEdgeNativePoolMachinesList(params)
	if err != nil {
		return nil, err
	}
	nMap := map[string]models.V1CloudMachineStatus{}
	if len(mpList.Payload.Items) > 0 {
		for _, node := range mpList.Payload.Items {
			nMap[node.Metadata.UID] = *node.Status
		}
	}
	return nMap, nil
}
