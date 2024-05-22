package client

import (
	"errors"
	"fmt"
	"strings"

	"github.com/spectrocloud/hapi/apiutil/transport"
	"github.com/spectrocloud/hapi/models"
	clusterC "github.com/spectrocloud/hapi/spectrocluster/client/v1"
)

func (h *V1Client) SearchPacks(filter *models.V1PackFilterSpec, sortBy []*models.V1PackSortSpec) ([]*models.V1PackMetadata, error) {
	body := &models.V1PacksFilterSpec{
		Filter: filter,
		Sort:   sortBy,
	}
	params := clusterC.NewV1PacksSearchParams().WithContext(h.Ctx).WithBody(body)

	resp, err := h.GetClusterClient().V1PacksSearch(params)
	var e *transport.TransportError
	if errors.As(err, &e) && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return resp.Payload.Items, nil
}

func (h *V1Client) GetClusterProfileManifestPack(clusterProfileUID, packName string) ([]*models.V1ManifestEntity, error) {
	//params := clusterC.NewV1ClusterProfilesGetParamsWithContext(h.ctx).WithUID(uid)
	params := clusterC.NewV1ClusterProfilesUIDPacksUIDManifestsParamsWithContext(h.Ctx).
		WithUID(clusterProfileUID).WithPackName(packName)

	success, err := h.GetClusterClient().V1ClusterProfilesUIDPacksUIDManifests(params)
	var e *transport.TransportError
	if errors.As(err, &e) && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return success.Payload.Items, nil
}

func (h *V1Client) GetPacks(filters []string, registryUID string) ([]*models.V1PackSummary, error) {
	if h.GetPacksFn != nil {
		return h.GetPacksFn(filters, registryUID)
	}

	params := clusterC.NewV1PacksSummaryListParamsWithContext(h.Ctx)
	if filters != nil {
		filterString := Ptr(strings.Join(filters, "AND"))
		params = params.WithFilters(filterString)
	}

	response, err := h.GetClusterClient().V1PacksSummaryList(params)
	if err != nil {
		return nil, err
	}

	packs := make([]*models.V1PackSummary, 0)
	for _, pack := range response.Payload.Items {
		if registryUID == "" || pack.Spec.RegistryUID == registryUID {
			packs = append(packs, pack)
		}
	}

	return packs, nil
}

func (h *V1Client) GetPacksByNameAndRegistry(name, registryUID, packContext string) (*models.V1PackTagEntity, error) {
	var params *clusterC.V1PacksNameRegistryUIDListParams
	switch packContext {
	case "project":
		params = clusterC.NewV1PacksNameRegistryUIDListParamsWithContext(h.Ctx).WithPackName(name).WithRegistryUID(registryUID)
	case "tenant":
		params = clusterC.NewV1PacksNameRegistryUIDListParams().WithPackName(name).WithRegistryUID(registryUID)
	default:
		return nil, fmt.Errorf("invalid pack context %s", packContext)

	}

	response, err := h.GetClusterClient().V1PacksNameRegistryUIDList(params)
	if err != nil {
		return nil, err
	}

	return response.Payload, nil
}

func (h *V1Client) GetPack(uid string) (*models.V1PackTagEntity, error) {
	params := clusterC.NewV1PacksUIDParamsWithContext(h.Ctx).WithUID(uid)
	response, err := h.GetClusterClient().V1PacksUID(params)
	if err != nil {
		return nil, err
	}

	return response.Payload, nil
}

func (h *V1Client) GetPackRegistry(packUID, packType string) string {
	if packUID == "uid" || packType == "manifest" {
		registry, err := h.GetPackRegistryCommonByName("Public Repo")
		if err != nil {
			return ""
		}
		return registry.UID
	}

	PackTagEntity, err := h.GetPack(packUID)
	if err != nil {
		return ""
	}

	return PackTagEntity.RegistryUID
}
