package client

import (
	"errors"
	"fmt"

	clientv1 "github.com/spectrocloud/palette-sdk-go/api/client/version1"
	"github.com/spectrocloud/palette-sdk-go/api/models"
	"github.com/spectrocloud/palette-sdk-go/client/apiutil"
)

// CRUDL operations on edge registration tokens are all tenant scoped.
// See: hubble/services/svccore/perms/edgetoken_acl.go

// GetRegistrationToken retrieves an existing registration token by name.
func (h *V1Client) GetRegistrationToken(tokenName string) (string, error) {
	// ACL scoped to tenant only
	params := clientv1.NewV1EdgeTokensListParams()
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

// GetRegistrationTokenByUID retrieves an existing registration token by token UID.
func (h *V1Client) GetRegistrationTokenByUID(tokenUID string) (*models.V1EdgeToken, error) {
	// ACL scoped to tenant only
	params := clientv1.NewV1EdgeTokensUIDGetParams().WithUID(tokenUID)
	resp, err := h.Client.V1EdgeTokensUIDGet(params)
	if err != nil {
		return nil, err
	}
	token := resp.Payload
	if token == nil {
		return nil, errors.New("failed to list registration tokens")
	}
	return resp.Payload, nil
}

// GetRegistrationTokenByName retrieves an existing registration token by name.
func (h *V1Client) GetRegistrationTokenByName(tokenName string) (*models.V1EdgeToken, error) {
	// ACL scoped to tenant only
	params := clientv1.NewV1EdgeTokensListParams()
	resp, err := h.Client.V1EdgeTokensList(params)
	if err != nil {
		return nil, err
	}
	tokens := resp.GetPayload()
	if tokens == nil {
		return nil, errors.New("failed to list registration tokens")
	}
	for _, token := range tokens.Items {
		if token.Metadata.Name == tokenName {
			return token, nil
		}
	}
	return nil, nil
}

// CreateRegistrationToken creates a new registration token.
func (h *V1Client) CreateRegistrationToken(tokenName string, body *models.V1EdgeTokenEntity) (string, string, error) {
	// ACL scoped to tenant only
	params := clientv1.NewV1EdgeTokensCreateParams().
		WithBody(body)
	res, err := h.Client.V1EdgeTokensCreate(params)
	if err != nil {
		return "", "", err
	}
	token, err := h.GetRegistrationToken(tokenName)

	return *res.Payload.UID, token, err
}

// UpdateRegistrationTokenByUID update an existing registration token by uid.
func (h *V1Client) UpdateRegistrationTokenByUID(tokenUID string, body *models.V1EdgeTokenUpdate) error {
	// ACL scoped to tenant only
	params := clientv1.NewV1EdgeTokensUIDUpdateParamsWithContext(h.ctx).WithUID(tokenUID).
		WithBody(body)
	_, err := h.Client.V1EdgeTokensUIDUpdate(params)
	if err != nil {
		return err
	}
	return nil
}

// DeleteRegistrationToken deletes a registration token by name.
func (h *V1Client) DeleteRegistrationToken(tokenUID string) error {
	// ACL scoped to tenant only
	params := clientv1.NewV1EdgeTokensUIDDeleteParams().WithUID(tokenUID)
	_, err := h.Client.V1EdgeTokensUIDDelete(params)
	if err != nil {
		return fmt.Errorf("failed to delete registration token: %w", err)
	}
	return nil
}

// GetEdgeHost retrieves an existing edge host by UID.
func (h *V1Client) GetEdgeHost(edgeHostID string) (*models.V1EdgeHostDevice, error) {
	params := clientv1.NewV1EdgeHostDevicesUIDGetParamsWithContext(h.ctx).
		WithUID(edgeHostID)
	resp, err := h.Client.V1EdgeHostDevicesUIDGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// UpdateEdgeHostTunnelConfig updates an edge host's SSH tunnel configuration.
func (h *V1Client) UpdateEdgeHostTunnelConfig(edgeHostID, remoteSSH, remoteSSHTempUser string) error {
	tunnelConfig := &models.V1SpectroTunnelConfig{
		RemoteSSH:         apiutil.Ptr(remoteSSH),
		RemoteSSHTempUser: apiutil.Ptr(remoteSSHTempUser),
	}

	params := clientv1.NewV1EdgeHostDevicesUIDTunnelConfigUpdateParamsWithContext(h.ctx).
		WithUID(edgeHostID).
		WithBody(tunnelConfig)

	_, err := h.Client.V1EdgeHostDevicesUIDTunnelConfigUpdate(params)
	return err
}

// ListEdgeHosts returns a list of all edge hosts.
// TODO: expose pagination params
func (h *V1Client) ListEdgeHosts() ([]*models.V1EdgeHostsMetadata, error) {
	continueToken := ""
	var items []*models.V1EdgeHostsMetadata
	for ok := true; ok; ok = (continueToken != "") {
		params := clientv1.NewV1DashboardEdgehostsSearchParamsWithContext(h.ctx)
		resp, err := h.Client.V1DashboardEdgehostsSearch(params)
		if err != nil {
			return nil, err
		}
		continueToken = resp.Payload.Listmeta.Continue
		items = append(items, resp.Payload.Items...)
	}

	return items, nil
}

// GetEdgeHostsByTags returns a list of edge hosts filtered by the provided tags.
// TODO: expose pagination params
func (h *V1Client) GetEdgeHostsByTags(tags map[string]string) ([]*models.V1EdgeHostsMetadata, error) {
	var offset int64
	continueToken := ""
	var items []*models.V1EdgeHostsMetadata
	filter := GetEdgeFilter(nil, tags)
	for ok := true; ok; ok = (continueToken != "") {
		params := clientv1.NewV1DashboardEdgehostsSearchParamsWithContext(h.ctx).
			WithBody(&models.V1SearchFilterSummarySpec{
				Filter: filter,
				Sort:   nil,
			}).
			WithOffset(&offset)
		resp, err := h.Client.V1DashboardEdgehostsSearch(params)
		if err != nil {
			return nil, err
		}
		continueToken = resp.Payload.Listmeta.Continue
		offset = resp.Payload.Listmeta.Offset
		items = append(items, resp.Payload.Items...)
	}

	return items, nil
}

// GetEdgeFilter returns a filter spec for listing edge hosts by tags and extra filters.
func GetEdgeFilter(extraFilters []*models.V1SearchFilterItem, tags map[string]string) *models.V1SearchFilterSpec {
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

	// Append extra filters if provided
	if extraFilters != nil {
		filter.FilterGroups = append(filter.FilterGroups, &models.V1SearchFilterGroup{Conjunction: and(), Filters: extraFilters})
	}

	return filter
}

// CreateClusterEdgeNative creates a new edge native cluster.
func (h *V1Client) CreateClusterEdgeNative(cluster *models.V1SpectroEdgeNativeClusterEntity) (string, error) {
	params := clientv1.NewV1SpectroClustersEdgeNativeCreateParamsWithContext(h.ctx).
		WithBody(cluster)
	resp, err := h.Client.V1SpectroClustersEdgeNativeCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

// CreateMachinePoolEdgeNative creates a new edge native machine pool.
func (h *V1Client) CreateMachinePoolEdgeNative(cloudConfigUID string, machinePool *models.V1EdgeNativeMachinePoolConfigEntity) error {
	params := clientv1.NewV1CloudConfigsEdgeNativeMachinePoolCreateParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUID).
		WithBody(machinePool)
	_, err := h.Client.V1CloudConfigsEdgeNativeMachinePoolCreate(params)
	return err
}

// UpdateMachinePoolEdgeNative updates an existing edge native machine pool.
func (h *V1Client) UpdateMachinePoolEdgeNative(cloudConfigUID string, machinePool *models.V1EdgeNativeMachinePoolConfigEntity) error {
	params := clientv1.NewV1CloudConfigsEdgeNativeMachinePoolUpdateParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUID).
		WithMachinePoolName(*machinePool.PoolConfig.Name).
		WithBody(machinePool)
	_, err := h.Client.V1CloudConfigsEdgeNativeMachinePoolUpdate(params)
	return err
}

// GetMachineEdgeNative retrieves an existing edge native machine by UID.
func (h *V1Client) GetMachineEdgeNative(machineUID, machinePoolName, cloudConfigUID string) (*models.V1EdgeNativeMachine, error) {
	params := clientv1.NewV1CloudConfigsEdgeNativePoolMachinesUIDGetParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUID).
		WithMachinePoolName(machinePoolName).
		WithMachineUID(machineUID)
	resp, err := h.Client.V1CloudConfigsEdgeNativePoolMachinesUIDGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// DeleteMachineEdgeNative deletes an existing edge native machine.
func (h *V1Client) DeleteMachineEdgeNative(edgeHostUID, machinePool, cloudConfigUID string) error {
	params := clientv1.NewV1CloudConfigsEdgeNativePoolMachinesUIDDeleteParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUID).
		WithMachinePoolName(machinePool).
		WithMachineUID(edgeHostUID)
	_, err := h.Client.V1CloudConfigsEdgeNativePoolMachinesUIDDelete(params)
	return err
}

// ListMachinesInPoolEdgeNative returns a list of machines in an edge native machine pool.
func (h *V1Client) ListMachinesInPoolEdgeNative(machinePoolName, cloudConfigUID string) ([]*models.V1EdgeNativeMachine, error) {
	params := clientv1.NewV1CloudConfigsEdgeNativePoolMachinesListParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUID).
		WithMachinePoolName(machinePoolName)
	resp, err := h.Client.V1CloudConfigsEdgeNativePoolMachinesList(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload.Items, nil
}

// DeleteMachinePoolEdgeNative deletes an existing edge native machine pool.
func (h *V1Client) DeleteMachinePoolEdgeNative(cloudConfigUID, machinePoolName string) error {
	params := clientv1.NewV1CloudConfigsEdgeNativeMachinePoolDeleteParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUID).
		WithMachinePoolName(machinePoolName)
	_, err := h.Client.V1CloudConfigsEdgeNativeMachinePoolDelete(params)
	return err
}

// GetCloudConfigEdgeNative retrieves an existing edge native cluster's cloud config.
func (h *V1Client) GetCloudConfigEdgeNative(configUID string) (*models.V1EdgeNativeCloudConfig, error) {
	params := clientv1.NewV1CloudConfigsEdgeNativeGetParamsWithContext(h.ctx).
		WithConfigUID(configUID)
	resp, err := h.Client.V1CloudConfigsEdgeNativeGet(params)
	if apiutil.Is404(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// GetNodeStatusMapEdgeNative retrieves the status of all nodes in an edge native machine pool.
func (h *V1Client) GetNodeStatusMapEdgeNative(configUID, machinePoolName string) (map[string]models.V1CloudMachineStatus, error) {
	params := clientv1.NewV1CloudConfigsEdgeNativePoolMachinesListParamsWithContext(h.ctx).
		WithConfigUID(configUID).
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

// UpdateRegistrationTokenState set registration token state
func (h *V1Client) UpdateRegistrationTokenState(tokenUID string, body *models.V1EdgeTokenActiveState) error {
	params := clientv1.NewV1EdgeTokensUIDStateParams().WithUID(tokenUID).WithBody(body)
	_, err := h.Client.V1EdgeTokensUIDState(params)
	return err
}

// UpdateEdgeHostDeviceTunnelConfig updates the specified edge host device tunnel configuration
func (h *V1Client) UpdateEdgeHostDeviceTunnelConfig(edgeHostID string, remoteSSH, remoteSSHTempUser *string) error {
	tunnelConfig := &models.V1SpectroTunnelConfig{
		RemoteSSH:         remoteSSH,
		RemoteSSHTempUser: remoteSSHTempUser,
	}

	params := clientv1.NewV1EdgeHostDevicesUIDTunnelConfigUpdateParams().
		WithUID(edgeHostID).
		WithBody(tunnelConfig)
	_, err := h.Client.V1EdgeHostDevicesUIDTunnelConfigUpdate(params)
	return err
}

// UpdateEdgeHostDevice updates an existing edge host device.
func (h *V1Client) UpdateEdgeHostDevice(edgeHostID string, body *models.V1EdgeHostDevice) error {
	params := clientv1.NewV1EdgeHostDevicesUIDUpdateParamsWithContext(h.ctx).
		WithUID(edgeHostID).
		WithBody(body)
	_, err := h.Client.V1EdgeHostDevicesUIDUpdate(params)
	return err
}

// DeleteEdgeHostDevice deletes an edge host device by UID.
func (h *V1Client) DeleteEdgeHostDevice(edgeHostID string) error {
	params := clientv1.NewV1EdgeHostDevicesUIDDeleteParamsWithContext(h.ctx).
		WithUID(edgeHostID)
	_, err := h.Client.V1EdgeHostDevicesUIDDelete(params)
	return err
}

// GetEdgeHostByName retrieves an existing edge host by name.
func (h *V1Client) GetEdgeHostByName(edgeHostName string) (*models.V1EdgeHostDevice, error) {
	params := clientv1.NewV1EdgeHostDevicesUIDGetParamsWithContext(h.ctx)
	resp, err := h.Client.V1EdgeHostDevicesUIDGet(params)
	if err != nil {
		return nil, err
	}
	if resp.Payload.Metadata.Name == edgeHostName {
		return resp.Payload, nil
	}
	return nil, fmt.Errorf("edge host with name %s not found", edgeHostName)
}

// UpdateClusterEdgeNative updates an existing edge native cluster.
func (h *V1Client) UpdateClusterEdgeNative(clusterUID string, cluster *models.V1SpectroEdgeNativeClusterEntity) error {
	params := clientv1.NewV1SpectroClustersEdgeNativeCreateParamsWithContext(h.ctx).
		WithBody(cluster)
	_, err := h.Client.V1SpectroClustersEdgeNativeCreate(params)
	if err != nil {
		return fmt.Errorf("failed to update cluster %s: %w", clusterUID, err)
	}
	return nil
}

// DeleteClusterEdgeNative deletes an edge native cluster by UID.
func (h *V1Client) DeleteClusterEdgeNative(clusterUID string) error {
	params := clientv1.NewV1SpectroClustersEdgeNativeCreateParamsWithContext(h.ctx)
	_, err := h.Client.V1SpectroClustersEdgeNativeCreate(params)
	if err != nil {
		return fmt.Errorf("failed to delete cluster %s: %w", clusterUID, err)
	}
	return nil
}

// UpdateCloudConfigEdgeNative updates an existing edge native cluster's cloud config.
func (h *V1Client) UpdateCloudConfigEdgeNative(configUID string, config *models.V1EdgeNativeCloudClusterConfigEntity) error {
	if config == nil || config.ClusterConfig == nil {
		return fmt.Errorf("invalid cloud config: missing cluster config")
	}

	params := clientv1.NewV1CloudConfigsEdgeNativeUIDClusterConfigParamsWithContext(h.ctx).
		WithConfigUID(configUID).
		WithBody(config)
	_, err := h.Client.V1CloudConfigsEdgeNativeUIDClusterConfig(params)
	if err != nil {
		return fmt.Errorf("failed to update cloud config %s: %w", configUID, err)
	}
	return nil
}
