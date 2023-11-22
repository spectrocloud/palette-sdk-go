package client

import (
	"errors"
	"fmt"
	"strings"

	"github.com/spectrocloud/hapi/apiutil/transport"
	hashboardC "github.com/spectrocloud/hapi/hashboard/client/v1"
	"github.com/spectrocloud/hapi/models"
	clusterC "github.com/spectrocloud/hapi/spectrocluster/client/v1"

	"github.com/spectrocloud/palette-sdk-go/client/herr"
)

func (h *V1Client) DeleteCluster(scope, uid string) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return err
	}

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

	_, err = client.V1SpectroClustersDelete(params)
	return err
}

func (h *V1Client) ForceDeleteCluster(scope, uid string, forceDelete bool) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return err
	}

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

	_, err = client.V1SpectroClustersDelete(params)
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
	client, err := h.GetHashboardClient()
	if err != nil {
		return nil, err
	}

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

	resp, err := client.V1SpectroClustersSearchFilterSummary(params)
	if e, ok := err.(*transport.TransportError); ok && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return resp.Payload.Items, nil
}

func (h *V1Client) listClusters(clusterContext string) ([]*models.V1SpectroCluster, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}

	var params *clusterC.V1SpectroClustersListParams
	switch clusterContext {
	case "project":
		params = clusterC.NewV1SpectroClustersListParamsWithContext(h.Ctx)
	case "tenant":
		params = clusterC.NewV1SpectroClustersListParams()
	}
	var limit int64 = 0
	params.Limit = &limit
	resp, err := client.V1SpectroClustersList(params)
	if e, ok := err.(*transport.TransportError); ok && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return resp.Payload.Items, nil
}

func (h *V1Client) listClustersMetadata(clusterContext string) ([]*models.V1ObjectMeta, error) {
	client, err := h.GetHashboardClient()
	if err != nil {
		return nil, err
	}

	var params *hashboardC.V1SpectroClustersMetadataParams
	switch clusterContext {
	case "project":
		params = hashboardC.NewV1SpectroClustersMetadataParams().WithContext(h.Ctx)
	case "tenant":
		params = hashboardC.NewV1SpectroClustersMetadataParams()
	}
	resp, err := client.V1SpectroClustersMetadata(params)
	if e, ok := err.(*transport.TransportError); ok && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return resp.Payload.Items, nil
}

// ListClusters This API is Deprecated in 3.1 in hubble, so basically it will return only 50 cluster by default.
func (h *V1Client) ListClusters(clusterContext string) ([]*models.V1SpectroCluster, error) {
	allClusters, err := h.listClusters(clusterContext)
	if err != nil {
		return nil, err
	}
	clusters := make([]*models.V1SpectroCluster, 0)
	for _, c := range allClusters {
		if c.Status.State != "Deleted" {
			clusters = append(clusters, c)
		}
	}
	return clusters, nil
}

func (h *V1Client) GetClusterWithoutStatus(scope, uid string) (*models.V1SpectroCluster, error) {
	if h.GetClusterWithoutStatusFn != nil {
		return h.GetClusterWithoutStatusFn(uid)
	}
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
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

	success, err := client.V1SpectroClustersGet(params)
	if err != nil {
		return nil, err
	}

	// special check if the cluster is marked deleted
	cluster := success.Payload
	return cluster, nil
}

func (h *V1Client) GetClusterByName(name, clusterContext string) (*models.V1SpectroCluster, error) {
	clustersMetadataList, err := h.listClustersMetadata(clusterContext)
	if err != nil {
		return nil, err
	}

	for _, clusterMetadata := range clustersMetadataList {
		if clusterMetadata.Name == name {
			cluster, err := h.GetCluster(clusterContext, clusterMetadata.UID)
			if err != nil {
				return nil, err
			}
			if cluster.Status.State != "Deleted" {
				return cluster, nil
			}
		}
	}

	return nil, nil
}

func (h *V1Client) GetClusterKubeConfig(uid, ClusterContext string) (string, error) {
	if h.GetClusterKubeConfigFn != nil {
		return h.GetClusterKubeConfigFn(uid)
	}
	client, err := h.GetClusterClient()
	if err != nil {
		return "", err
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
	_, err = client.V1SpectroClustersUIDKubeConfig(params, builder)
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
	client, err := h.GetClusterClient()
	if err != nil {
		return "", err
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
	_, err = client.V1SpectroClustersUIDAdminKubeConfig(params, builder)
	if err != nil {
		if herr.IsNotFound(err) {
			return "", nil
		}
		return "", err
	}

	return builder.String(), nil
}

func (h *V1Client) GetClusterImportManifest(uid string, clusterContext string) (string, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return "", err
	}

	builder := new(strings.Builder)
	var params *clusterC.V1SpectroClustersUIDImportManifestParams
	switch clusterContext {
	case "project":
		params = clusterC.NewV1SpectroClustersUIDImportManifestParamsWithContext(h.Ctx).WithUID(uid)
	case "tenant":
		params = clusterC.NewV1SpectroClustersUIDImportManifestParams().WithUID(uid)
	}
	_, err = client.V1SpectroClustersUIDImportManifest(params, builder)
	if err != nil {
		return "", err
	}

	return builder.String(), nil
}

func (h *V1Client) UpdateClusterProfileValues(uid string, context string, profiles *models.V1SpectroClusterProfiles) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return err
	}

	resolveNotification := true
	var params *clusterC.V1SpectroClustersUpdateProfilesParams
	switch context {
	case "project":
		params = clusterC.NewV1SpectroClustersUpdateProfilesParamsWithContext(h.Ctx).WithUID(uid).
			WithBody(profiles).WithResolveNotification(&resolveNotification)
	case "tenant":
		params = clusterC.NewV1SpectroClustersUpdateProfilesParams().WithUID(uid).
			WithBody(profiles).WithResolveNotification(&resolveNotification)
	}
	_, err = client.V1SpectroClustersUpdateProfiles(params)
	return err
}

func (h *V1Client) ImportClusterGeneric(meta *models.V1ObjectMetaInputEntity) (string, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return "", err
	}

	params := clusterC.NewV1SpectroClustersGenericImportParamsWithContext(h.Ctx).WithBody(
		&models.V1SpectroGenericClusterImportEntity{
			Metadata: meta,
		},
	)
	success, err := client.V1SpectroClustersGenericImport(params)
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

func (h *V1Client) ApproveClusterRepave(clusterUID string, context string) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return err
	}
	var params *clusterC.V1SpectroClustersUIDRepaveApproveUpdateParams
	switch context {
	case "project":
		params = clusterC.NewV1SpectroClustersUIDRepaveApproveUpdateParamsWithContext(h.Ctx).WithUID(clusterUID)
	case "tenant":
		params = clusterC.NewV1SpectroClustersUIDRepaveApproveUpdateParams().WithUID(clusterUID)
	}
	_, err = client.V1SpectroClustersUIDRepaveApproveUpdate(params)
	return err
}
