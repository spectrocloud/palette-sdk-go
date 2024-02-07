package client

import (
	"strings"

	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
	"github.com/spectrocloud/palette-sdk-go/client/apiutil"
)

func (h *V1Client) SearchPacks(filter *models.V1PackFilterSpec, sortBy []*models.V1PackSortSpec) ([]*models.V1PackMetadata, error) {
	params := clientV1.NewV1PacksSearchParamsWithContext(h.ctx).
		WithBody(&models.V1PacksFilterSpec{
			Filter: filter,
			Sort:   sortBy,
		})
	resp, err := h.Client.V1PacksSearch(params)
	if err := apiutil.Handle404(err); err != nil {
		return nil, err
	}
	return resp.Payload.Items, nil
}

func (h *V1Client) GetClusterProfileManifestPack(clusterProfileUid, packName string) ([]*models.V1ManifestEntity, error) {
	params := clientV1.NewV1ClusterProfilesUIDPacksUIDManifestsParamsWithContext(h.ctx).
		WithUID(clusterProfileUid).
		WithPackName(packName)
	resp, err := h.Client.V1ClusterProfilesUIDPacksUIDManifests(params)
	if err := apiutil.Handle404(err); err != nil {
		return nil, err
	}
	return resp.Payload.Items, nil
}

func (h *V1Client) GetPacks(filters []string, registryUid string) ([]*models.V1PackSummary, error) {
	params := clientV1.NewV1PacksSummaryListParamsWithContext(h.ctx)
	if filters != nil {
		filterString := apiutil.Ptr(strings.Join(filters, "AND"))
		params = params.WithFilters(filterString)
	}
	resp, err := h.Client.V1PacksSummaryList(params)
	if err != nil {
		return nil, err
	}
	packs := make([]*models.V1PackSummary, 0)
	for _, pack := range resp.Payload.Items {
		if registryUid == "" || pack.Spec.RegistryUID == registryUid {
			packs = append(packs, pack)
		}
	}
	return packs, nil
}

func (h *V1Client) GetPacksByProfile(profileUid string) ([]*models.V1ClusterProfilePacksEntity, error) {
	params := clientV1.NewV1ClusterProfilesUIDPacksGetParamsWithContext(h.ctx).
		WithUID(profileUid)
	resp, err := h.Client.V1ClusterProfilesUIDPacksGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload.Items, nil
}

func (h *V1Client) GetPacksByNameAndRegistry(name, registryUid string) (*models.V1PackTagEntity, error) {
	params := clientV1.NewV1PacksNameRegistryUIDListParamsWithContext(h.ctx).
		WithPackName(name).
		WithRegistryUID(registryUid)
	resp, err := h.Client.V1PacksNameRegistryUIDList(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

func (h *V1Client) GetPack(uid string) (*models.V1PackTagEntity, error) {
	params := clientV1.NewV1PacksUIDParamsWithContext(h.ctx).
		WithUID(uid)
	resp, err := h.Client.V1PacksUID(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

func (h *V1Client) GetPackRegistry(packUid, packType string) string {
	if packUid == "uid" || packType == "manifest" {
		registry, err := h.GetPackRegistryCommonByName("Public Repo")
		if err != nil {
			return ""
		}
		return registry.UID
	}
	packTagEntity, err := h.GetPack(packUid)
	if err != nil {
		return ""
	}
	return packTagEntity.RegistryUID
}
