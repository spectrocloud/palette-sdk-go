package client

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/spectrocloud/hapi/apiutil/transport"
	hashboardC "github.com/spectrocloud/hapi/hashboard/client/v1"
	"github.com/spectrocloud/hapi/models"
	clusterC "github.com/spectrocloud/hapi/spectrocluster/client/v1"

	"github.com/spectrocloud/palette-sdk-go/client/herr"
)

func (h *V1Client) DeleteCluster(scope, uid string) error {
	cluster, err := h.GetCluster(scope, uid)
	if err != nil || cluster == nil {
		return err
	}

	var params *clusterC.V1SpectroClustersDeleteParams
	switch scope {
	case "project":
		params = clusterC.NewV1SpectroClustersDeleteParamsWithContext(h.Ctx).WithUID(uid)
	case "tenant":
		params = clusterC.NewV1SpectroClustersDeleteParams().WithUID(uid)
	}

	_, err = h.GetClusterClient().V1SpectroClustersDelete(params)
	return err
}

func (h *V1Client) ForceDeleteCluster(scope, uid string, forceDelete bool) error {
	cluster, err := h.GetCluster(scope, uid)
	if err != nil || cluster == nil {
		return err
	}

	var params *clusterC.V1SpectroClustersDeleteParams
	switch scope {
	case "project":
		params = clusterC.NewV1SpectroClustersDeleteParamsWithContext(h.Ctx).WithUID(uid).WithForceDelete(&forceDelete)
	case "tenant":
		params = clusterC.NewV1SpectroClustersDeleteParams().WithUID(uid).WithForceDelete(&forceDelete)
	}

	_, err = h.GetClusterClient().V1SpectroClustersDelete(params)
	return err
}

func (h *V1Client) GetCluster(scope, uid string) (*models.V1SpectroCluster, error) {
	if h.GetClusterFn != nil {
		return h.GetClusterFn(scope, uid)
	}
	cluster, err := h.GetClusterWithoutStatus(scope, uid)
	if err != nil {
		return nil, err
	}

	if cluster == nil || cluster.Status.State == "Deleted" {
		return nil, nil
	}

	return cluster, nil
}

func (h *V1Client) SearchClusterSummaries(clusterContext string, filter *models.V1SearchFilterSpec, sort []*models.V1SearchFilterSortSpec) ([]*models.V1SpectroClusterSummary, error) {
	var params *hashboardC.V1SpectroClustersSearchFilterSummaryParams
	switch clusterContext {
	case "project":
		params = hashboardC.NewV1SpectroClustersSearchFilterSummaryParamsWithContext(h.Ctx)
	case "tenant":
		params = hashboardC.NewV1SpectroClustersSearchFilterSummaryParams()
	}
	params.Body = &models.V1SearchFilterSummarySpec{
		Filter: filter,
		Sort:   sort,
	}

	resp, err := h.GetHashboardClient().V1SpectroClustersSearchFilterSummary(params)
	var e *transport.TransportError
	if errors.As(err, &e) && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return resp.Payload.Items, nil
}

func (h *V1Client) GetClusterWithoutStatus(scope, uid string) (*models.V1SpectroCluster, error) {
	if h.GetClusterWithoutStatusFn != nil {
		return h.GetClusterWithoutStatusFn(uid)
	}

	var params *clusterC.V1SpectroClustersGetParams
	switch scope {
	case "project":
		params = clusterC.NewV1SpectroClustersGetParamsWithContext(h.Ctx).WithUID(uid)
	case "tenant":
		params = clusterC.NewV1SpectroClustersGetParams().WithUID(uid)
	default:
		return nil, fmt.Errorf("invalid scope %s", scope)
	}

	success, err := h.GetClusterClient().V1SpectroClustersGet(params)
	var e *transport.TransportError
	if errors.As(err, &e) && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	// special check if the cluster is marked deleted
	cluster := success.Payload
	return cluster, nil
}

func (h *V1Client) GetClusterKubeConfig(uid, ClusterContext string) (string, error) {
	if h.GetClusterKubeConfigFn != nil {
		return h.GetClusterKubeConfigFn(uid)
	}

	var params *clusterC.V1SpectroClustersUIDKubeConfigParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1SpectroClustersUIDKubeConfigParamsWithContext(h.Ctx).WithUID(uid)
	case "tenant":
		params = clusterC.NewV1SpectroClustersUIDKubeConfigParams().WithUID(uid)
	default:
		return "", errors.New("invalid cluster scope specified")
	}

	builder := new(strings.Builder)
	_, err := h.GetClusterClient().V1SpectroClustersUIDKubeConfig(params, builder)
	if err != nil {
		if herr.IsNotFound(err) {
			return "", nil
		}
		return "", err
	}

	return builder.String(), nil
}

func (h *V1Client) GetClusterAdminKubeConfig(uid, ClusterContext string) (string, error) {
	if h.GetClusterAdminConfigFn != nil {
		return h.GetClusterAdminConfigFn(uid)
	}

	var params *clusterC.V1SpectroClustersUIDAdminKubeConfigParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1SpectroClustersUIDAdminKubeConfigParamsWithContext(h.Ctx).WithUID(uid)
	case "tenant":
		params = clusterC.NewV1SpectroClustersUIDAdminKubeConfigParams().WithUID(uid)
	default:
		return "", errors.New("invalid cluster scope specified")
	}

	builder := new(strings.Builder)
	_, err := h.GetClusterClient().V1SpectroClustersUIDAdminKubeConfig(params, builder)
	if err != nil {
		if herr.IsNotFound(err) {
			return "", nil
		}
		return "", err
	}

	return builder.String(), nil
}

func (h *V1Client) GetClusterImportManifest(uid, clusterContext string) (string, error) {
	builder := new(strings.Builder)
	var params *clusterC.V1SpectroClustersUIDImportManifestParams
	switch clusterContext {
	case "project":
		params = clusterC.NewV1SpectroClustersUIDImportManifestParamsWithContext(h.Ctx).WithUID(uid)
	case "tenant":
		params = clusterC.NewV1SpectroClustersUIDImportManifestParams().WithUID(uid)
	}

	_, err := h.GetClusterClient().V1SpectroClustersUIDImportManifest(params, builder)
	if err != nil {
		return "", err
	}

	return builder.String(), nil
}

func (h *V1Client) UpdateClusterProfileValues(uid, context string, profiles *models.V1SpectroClusterProfiles) error {
	var params *clusterC.V1SpectroClustersUpdateProfilesParams
	switch context {
	case "project":
		params = clusterC.NewV1SpectroClustersUpdateProfilesParamsWithContext(h.Ctx).WithUID(uid).
			WithBody(profiles).WithResolveNotification(Ptr(true))
	case "tenant":
		params = clusterC.NewV1SpectroClustersUpdateProfilesParams().WithUID(uid).
			WithBody(profiles).WithResolveNotification(Ptr(true))
	}
	params = params.WithTimeout(60 * time.Second)
	_, err := h.GetClusterClient().V1SpectroClustersUpdateProfiles(params)
	return err
}

func (h *V1Client) ImportClusterGeneric(meta *models.V1ObjectMetaInputEntity) (string, error) {
	params := clusterC.NewV1SpectroClustersGenericImportParamsWithContext(h.Ctx).WithBody(
		&models.V1SpectroGenericClusterImportEntity{
			Metadata: meta,
		},
	)
	success, err := h.GetClusterClient().V1SpectroClustersGenericImport(params)
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

func (h *V1Client) ApproveClusterRepave(context, clusterUID string) error {
	if h.ApproveClusterRepaveFn != nil {
		return h.ApproveClusterRepaveFn(clusterUID, context)
	}

	var params *clusterC.V1SpectroClustersUIDRepaveApproveUpdateParams
	switch context {
	case "project":
		params = clusterC.NewV1SpectroClustersUIDRepaveApproveUpdateParamsWithContext(h.Ctx).WithUID(clusterUID)
	case "tenant":
		params = clusterC.NewV1SpectroClustersUIDRepaveApproveUpdateParams().WithUID(clusterUID)
	default:
		return fmt.Errorf("invalid context: %s", context)
	}

	_, err := h.GetClusterClient().V1SpectroClustersUIDRepaveApproveUpdate(params)
	return err
}

func (h *V1Client) GetRepaveReasons(context, clusterUID string) ([]string, error) {
	if h.GetRepaveReasonsFn != nil {
		return h.GetRepaveReasonsFn(context, clusterUID)
	}

	var params *clusterC.V1SpectroClustersUIDRepaveGetParams
	switch context {
	case "project":
		params = clusterC.NewV1SpectroClustersUIDRepaveGetParamsWithContext(h.Ctx).WithUID(clusterUID)
	case "tenant":
		params = clusterC.NewV1SpectroClustersUIDRepaveGetParams().WithUID(clusterUID)
	default:
		return nil, fmt.Errorf("invalid context: %s", context)
	}

	res, err := h.GetClusterClient().V1SpectroClustersUIDRepaveGet(params)
	if err != nil {
		return nil, err
	}
	var reasons []string
	for _, r := range res.Payload.Spec.Reasons {
		reasons = append(reasons, fmt.Sprintf("Repave Reason Code -  %s , Reason - %s.", r.Code, r.Message))
	}
	return reasons, err
}

func (h *V1Client) UpdatePauseAgentUpgradeSettingCluster(upgradeSetting *models.V1ClusterUpgradeSettingsEntity, clusterUID string, context string) error {
	if h.UpdatePauseAgentUpgradeSettingClusterFn != nil {
		return h.UpdatePauseAgentUpgradeSettingClusterFn(upgradeSetting, clusterUID, context)
	}
	var params *clusterC.V1SpectroClustersUIDUpgradeSettingsParams

	switch context {
	case "project":
		params = clusterC.NewV1SpectroClustersUIDUpgradeSettingsParamsWithContext(h.Ctx)
	case "tenant":
		params = clusterC.NewV1SpectroClustersUIDUpgradeSettingsParams()
	default:
		return fmt.Errorf("invalid context: %s", context)
	}
	params = params.WithUID(clusterUID).WithBody(upgradeSetting)
	_, err := h.GetClusterClient().V1SpectroClustersUIDUpgradeSettings(params)
	if err != nil {
		return err
	}
	return nil
}

func (h *V1Client) UpdatePauseAgentUpgradeSettingContext(upgradeSetting *models.V1ClusterUpgradeSettingsEntity, context string) error {
	if h.UpdatePauseAgentUpgradeSettingContextFn != nil {
		return h.UpdatePauseAgentUpgradeSettingContextFn(upgradeSetting, context)
	}
	var params *clusterC.V1SpectroClustersUpgradeSettingsParams

	switch context {
	case "project":
		params = clusterC.NewV1SpectroClustersUpgradeSettingsParamsWithContext(h.Ctx)
	case "tenant":
		params = clusterC.NewV1SpectroClustersUpgradeSettingsParams()
	default:
		return fmt.Errorf("invalid context: %s", context)
	}
	params = params.WithBody(upgradeSetting)
	_, err := h.GetClusterClient().V1SpectroClustersUpgradeSettings(params)
	if err != nil {
		return err
	}
	return nil
}

func (h *V1Client) GetPauseAgentUpgradeSettingContext(context string) (string, error) {
	if h.GetPauseAgentUpgradeSettingContextFn != nil {
		return h.GetPauseAgentUpgradeSettingContextFn(context)
	}
	var params *clusterC.V1SpectroClustersUpgradeSettingsGetParams

	switch context {
	case "project":
		params = clusterC.NewV1SpectroClustersUpgradeSettingsGetParamsWithContext(h.Ctx)
	case "tenant":
		params = clusterC.NewV1SpectroClustersUpgradeSettingsGetParams()
	default:
		return "", fmt.Errorf("invalid context: %s", context)
	}
	resp, err := h.GetClusterClient().V1SpectroClustersUpgradeSettingsGet(params)
	if err != nil {
		return "", err
	}
	return resp.Payload.SpectroComponents, nil
}
