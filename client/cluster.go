package client

import (
	"strings"

	"github.com/spectrocloud/hapi/apiutil/transport"
	hashboardC "github.com/spectrocloud/hapi/hashboard/client/v1"
	"github.com/spectrocloud/hapi/models"
	clusterC "github.com/spectrocloud/hapi/spectrocluster/client/v1"
	"github.com/spectrocloud/palette-sdk-go/client/herr"
)

func (h *V1Client) DeleteCluster(uid string) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return err
	}

	cluster, err := h.GetCluster(uid)
	if err != nil || cluster == nil {
		return err
	}

	var params *clusterC.V1SpectroClustersDeleteParams
	switch cluster.Metadata.Annotations["scope"] {
	case "project":
		params = clusterC.NewV1SpectroClustersDeleteParamsWithContext(h.Ctx).WithUID(uid)
	case "tenant":
		params = clusterC.NewV1SpectroClustersDeleteParams().WithUID(uid)
	}

	_, err = client.V1SpectroClustersDelete(params)
	return err
}

func (h *V1Client) GetCluster(uid string) (*models.V1SpectroCluster, error) {
	if h.GetClusterFn != nil {
		return h.GetClusterFn(uid)
	}
	cluster, err := h.GetClusterWithoutStatus(uid)
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

	resp, err := client.V1SpectroClustersList(params)
	if e, ok := err.(*transport.TransportError); ok && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return resp.Payload.Items, nil
}

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

func (h *V1Client) GetClusterWithoutStatus(uid string) (*models.V1SpectroCluster, error) {
	if h.GetClusterWithoutStatusFn != nil {
		return h.GetClusterWithoutStatusFn(uid)
	}
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}

	params := clusterC.NewV1SpectroClustersGetParamsWithContext(h.Ctx).WithUID(uid)
	success, err := client.V1SpectroClustersGet(params)
	// handle tenant context here cluster may be a tenant cluster
	if e, ok := err.(*transport.TransportError); ok && e.HttpCode == 404 {
		params := clusterC.NewV1SpectroClustersGetParams().WithUID(uid)
		success, err = client.V1SpectroClustersGet(params)
		if e, ok := err.(*transport.TransportError); ok && e.HttpCode == 404 {
			return nil, nil
		} else if err != nil {
			return nil, err
		}
		//return nil, nil
	}
	if e, ok := err.(*transport.TransportError); ok && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	// special check if the cluster is marked deleted
	cluster := success.Payload
	return cluster, nil
}

func (h *V1Client) GetClusterByName(name, clusterContext string) (*models.V1SpectroCluster, error) {
	clusters, err := h.listClusters(clusterContext)
	if err != nil {
		return nil, err
	}

	for _, cluster := range clusters {
		if cluster.Metadata.Name == name && cluster.Status.State != "Deleted" {
			return cluster, nil
		}
	}

	return nil, nil
}

func (h *V1Client) GetClusterKubeConfig(uid string) (string, error) {
	if h.GetClusterKubeConfigFn != nil {
		return h.GetClusterKubeConfigFn(uid)
	}
	client, err := h.GetClusterClient()
	if err != nil {
		return "", err
	}

	builder := new(strings.Builder)

	params := clusterC.NewV1SpectroClustersUIDKubeConfigParamsWithContext(h.Ctx).WithUID(uid)
	_, err = client.V1SpectroClustersUIDKubeConfig(params, builder)
	// handle tenant context here cluster may be a tenant cluster
	if e, ok := err.(*transport.TransportError); ok && e.HttpCode == 404 {
		params := clusterC.NewV1SpectroClustersUIDKubeConfigParams().WithUID(uid)
		_, err = client.V1SpectroClustersUIDKubeConfig(params, builder)
		if e, ok := err.(*transport.TransportError); ok && e.HttpCode == 404 {
			return "", nil
		} else if err != nil {
			return "", err
		}
	}
	if err != nil {
		if herr.IsNotFound(err) {
			return "", nil
		}
		return "", err
	}

	return builder.String(), nil
}

func (h *V1Client) GetClusterImportManifest(uid string) (string, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return "", err
	}

	builder := new(strings.Builder)
	params := clusterC.NewV1SpectroClustersUIDImportManifestParamsWithContext(h.Ctx).WithUID(uid)
	_, err = client.V1SpectroClustersUIDImportManifest(params, builder)
	if err != nil {
		return "", err
	}

	return builder.String(), nil
}

func (h *V1Client) UpdateClusterProfileValues(uid string, profiles *models.V1SpectroClusterProfiles) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return err
	}

	resolveNotification := true
	params := clusterC.NewV1SpectroClustersUpdateProfilesParamsWithContext(h.Ctx).WithUID(uid).
		WithBody(profiles).WithResolveNotification(&resolveNotification)
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
