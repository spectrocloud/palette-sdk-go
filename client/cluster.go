package client

import (
	"errors"
	"fmt"
	"strings"

	"github.com/spectrocloud/palette-api-go/apiutil/transport"
	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
	"github.com/spectrocloud/palette-sdk-go/client/apiutil"
	"github.com/spectrocloud/palette-sdk-go/client/herr"
)

func (h *V1Client) DeleteCluster(scope, uid string) error {
	cluster, err := h.GetCluster(scope, uid)
	if err != nil || cluster == nil {
		return err
	}

	var params *clientV1.V1SpectroClustersDeleteParams
	switch scope {
	case "project":
		params = clientV1.NewV1SpectroClustersDeleteParamsWithContext(h.Ctx).WithUID(uid)
	case "tenant":
		params = clientV1.NewV1SpectroClustersDeleteParams().WithUID(uid)
	}

	_, err = h.Client.V1SpectroClustersDelete(params)
	return err
}

func (h *V1Client) ForceDeleteCluster(scope, uid string, forceDelete bool) error {
	cluster, err := h.GetCluster(scope, uid)
	if err != nil || cluster == nil {
		return err
	}

	var params *clientV1.V1SpectroClustersDeleteParams
	switch scope {
	case "project":
		params = clientV1.NewV1SpectroClustersDeleteParamsWithContext(h.Ctx).WithUID(uid).WithForceDelete(&forceDelete)
	case "tenant":
		params = clientV1.NewV1SpectroClustersDeleteParams().WithUID(uid).WithForceDelete(&forceDelete)
	}

	_, err = h.Client.V1SpectroClustersDelete(params)
	return err
}

func (h *V1Client) GetCluster(scope, uid string) (*models.V1SpectroCluster, error) {
	cluster, err := h.GetClusterWithoutStatus(scope, uid)
	if err != nil {
		return nil, err
	}

	if cluster == nil || cluster.Status.State == "Deleted" {
		return nil, nil
	}

	return cluster, nil
}

func (h *V1Client) GetClusterSummary(scope, clusterId string) (*models.V1SpectroClusterUIDSummary, error) {
	var params *clientV1.V1SpectroClustersSummaryUIDParams
	switch scope {
	case "project":
		params = clientV1.NewV1SpectroClustersSummaryUIDParamsWithContext(h.Ctx)
	case "tenant":
		params = clientV1.NewV1SpectroClustersSummaryUIDParams()
	default:
		return nil, fmt.Errorf("invalid scope %s", scope)
	}
	params = params.WithUID(clusterId)

	resp, err := h.Client.V1SpectroClustersSummaryUID(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

func (h *V1Client) SearchClusterSummaries(scope string, filter *models.V1SearchFilterSpec, sort []*models.V1SearchFilterSortSpec) ([]*models.V1SpectroClusterSummary, error) {
	var params *clientV1.V1SpectroClustersSearchFilterSummaryParams
	switch scope {
	case "project":
		params = clientV1.NewV1SpectroClustersSearchFilterSummaryParamsWithContext(h.Ctx)
	case "tenant":
		params = clientV1.NewV1SpectroClustersSearchFilterSummaryParams()
	}
	params.Body = &models.V1SearchFilterSummarySpec{
		Filter: filter,
		Sort:   sort,
	}

	resp, err := h.Client.V1SpectroClustersSearchFilterSummary(params)
	var e *transport.TransportError
	if errors.As(err, &e) && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return resp.Payload.Items, nil
}

func (h *V1Client) GetClusterWithoutStatus(scope, uid string) (*models.V1SpectroCluster, error) {
	var params *clientV1.V1SpectroClustersGetParams
	switch scope {
	case "project":
		params = clientV1.NewV1SpectroClustersGetParamsWithContext(h.Ctx).WithUID(uid)
	case "tenant":
		params = clientV1.NewV1SpectroClustersGetParams().WithUID(uid)
	default:
		return nil, fmt.Errorf("invalid scope %s", scope)
	}

	success, err := h.Client.V1SpectroClustersGet(params)
	if err != nil {
		return nil, err
	}

	// special check if the cluster is marked deleted
	cluster := success.Payload
	return cluster, nil
}

func (h *V1Client) GetClusterKubeConfig(uid, scope string) (string, error) {
	var params *clientV1.V1SpectroClustersUIDKubeConfigParams
	switch scope {
	case "project":
		params = clientV1.NewV1SpectroClustersUIDKubeConfigParamsWithContext(h.Ctx).WithUID(uid)
	case "tenant":
		params = clientV1.NewV1SpectroClustersUIDKubeConfigParams().WithUID(uid)
	default:
		return "", errors.New("invalid cluster scope specified")
	}

	builder := new(strings.Builder)
	_, err := h.Client.V1SpectroClustersUIDKubeConfig(params, builder)
	if err != nil {
		if herr.IsNotFound(err) {
			return "", nil
		}
		return "", err
	}
	kubeconfig := builder.String()

	if apiutil.IsBase64(kubeconfig) {
		if bytes, err := apiutil.Base64DecodeString(kubeconfig); err != nil {
			return "", err
		} else {
			return string(bytes), nil
		}
	} else {
		return kubeconfig, nil
	}
}

func (h *V1Client) GetClusterAdminKubeConfig(uid, scope string) (string, error) {
	var params *clientV1.V1SpectroClustersUIDAdminKubeConfigParams
	switch scope {
	case "project":
		params = clientV1.NewV1SpectroClustersUIDAdminKubeConfigParamsWithContext(h.Ctx).WithUID(uid)
	case "tenant":
		params = clientV1.NewV1SpectroClustersUIDAdminKubeConfigParams().WithUID(uid)
	default:
		return "", errors.New("invalid cluster scope specified")
	}

	builder := new(strings.Builder)
	_, err := h.Client.V1SpectroClustersUIDAdminKubeConfig(params, builder)
	if err != nil {
		if herr.IsNotFound(err) {
			return "", nil
		}
		return "", err
	}

	return builder.String(), nil
}

func (h *V1Client) GetClusterImportManifest(uid, scope string) (string, error) {
	builder := new(strings.Builder)
	var params *clientV1.V1SpectroClustersUIDImportManifestParams
	switch scope {
	case "project":
		params = clientV1.NewV1SpectroClustersUIDImportManifestParamsWithContext(h.Ctx).WithUID(uid)
	case "tenant":
		params = clientV1.NewV1SpectroClustersUIDImportManifestParams().WithUID(uid)
	}

	_, err := h.Client.V1SpectroClustersUIDImportManifest(params, builder)
	if err != nil {
		return "", err
	}

	return builder.String(), nil
}

func (h *V1Client) UpdateClusterProfileValues(uid, context string, profiles *models.V1SpectroClusterProfiles) error {
	var params *clientV1.V1SpectroClustersUpdateProfilesParams
	switch context {
	case "project":
		params = clientV1.NewV1SpectroClustersUpdateProfilesParamsWithContext(h.Ctx).WithUID(uid).
			WithBody(profiles).WithResolveNotification(Ptr(true))
	case "tenant":
		params = clientV1.NewV1SpectroClustersUpdateProfilesParams().WithUID(uid).
			WithBody(profiles).WithResolveNotification(Ptr(true))
	}

	_, err := h.Client.V1SpectroClustersUpdateProfiles(params)
	return err
}

func (h *V1Client) ImportClusterGeneric(meta *models.V1ObjectMetaInputEntity) (string, error) {
	params := clientV1.NewV1SpectroClustersGenericImportParamsWithContext(h.Ctx).WithBody(
		&models.V1SpectroGenericClusterImportEntity{
			Metadata: meta,
		},
	)
	success, err := h.Client.V1SpectroClustersGenericImport(params)
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

func (h *V1Client) ApproveClusterRepave(context, clusterUID string) error {
	var params *clientV1.V1SpectroClustersUIDRepaveApproveUpdateParams
	switch context {
	case "project":
		params = clientV1.NewV1SpectroClustersUIDRepaveApproveUpdateParamsWithContext(h.Ctx).WithUID(clusterUID)
	case "tenant":
		params = clientV1.NewV1SpectroClustersUIDRepaveApproveUpdateParams().WithUID(clusterUID)
	default:
		return fmt.Errorf("invalid context: %s", context)
	}

	_, err := h.Client.V1SpectroClustersUIDRepaveApproveUpdate(params)
	return err
}

func (h *V1Client) GetRepaveReasons(context, clusterUID string) ([]string, error) {
	var params *clientV1.V1SpectroClustersUIDRepaveGetParams
	switch context {
	case "project":
		params = clientV1.NewV1SpectroClustersUIDRepaveGetParamsWithContext(h.Ctx).WithUID(clusterUID)
	case "tenant":
		params = clientV1.NewV1SpectroClustersUIDRepaveGetParams().WithUID(clusterUID)
	default:
		return nil, fmt.Errorf("invalid context: %s", context)
	}

	res, err := h.Client.V1SpectroClustersUIDRepaveGet(params)
	if err != nil {
		return nil, err
	}
	var reasons []string
	for _, r := range res.Payload.Spec.Reasons {
		reasons = append(reasons, fmt.Sprintf("Repave Reason Code -  %s , Reason - %s.", r.Code, r.Message))
	}
	return reasons, err
}
