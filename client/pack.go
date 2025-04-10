package client

import (
	"strings"

	clientv1 "github.com/spectrocloud/palette-sdk-go/api/client/version1"
	"github.com/spectrocloud/palette-sdk-go/api/models"
	"github.com/spectrocloud/palette-sdk-go/client/apiutil"
)

// SearchPacks retrieves a list of pack metadata based on the filter and sort criteria.
func (h *V1Client) SearchPacks(filter *models.V1PackFilterSpec, sortBy []*models.V1PackSortSpec) ([]*models.V1PackMetadata, error) {
	params := clientv1.NewV1PacksSearchParamsWithContext(h.ctx).
		WithBody(&models.V1PacksFilterSpec{
			Filter: filter,
			Sort:   sortBy,
		})
	resp, err := h.Client.V1PacksSearch(params)
	if apiutil.Is404(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return resp.Payload.Items, nil
}

// GetClusterProfileManifestPack retrieves all manifests for a pack in a cluster profile.
func (h *V1Client) GetClusterProfileManifestPack(clusterProfileUID, packName string) ([]*models.V1ManifestEntity, error) {
	params := clientv1.NewV1ClusterProfilesUIDPacksUIDManifestsParamsWithContext(h.ctx).
		WithUID(clusterProfileUID).
		WithPackName(packName)
	resp, err := h.Client.V1ClusterProfilesUIDPacksUIDManifests(params)
	if apiutil.Is404(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return resp.Payload.Items, nil
}

// GetPacks retrieves a list of pack summaries based on the provided registry UID and filter criteria.
func (h *V1Client) GetPacks(filters []string, registryUID string) ([]*models.V1PackSummary, error) {
	params := clientv1.NewV1PacksSummaryListParamsWithContext(h.ctx)
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
		if registryUID == "" || pack.Spec.RegistryUID == registryUID {
			packs = append(packs, pack)
		}
	}
	return packs, nil
}

// GetPacksByProfile retrieves all packs for a cluster profile.
func (h *V1Client) GetPacksByProfile(profileUID string) ([]*models.V1ClusterProfilePacksEntity, error) {
	params := clientv1.NewV1ClusterProfilesUIDPacksGetParamsWithContext(h.ctx).
		WithUID(profileUID)
	resp, err := h.Client.V1ClusterProfilesUIDPacksGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload.Items, nil
}

// GetPacksByNameAndRegistry retrieves a pack by name and registry UID.
func (h *V1Client) GetPacksByNameAndRegistry(name, registryUID string) (*models.V1PackTagEntity, error) {
	params := clientv1.NewV1PacksNameRegistryUIDListParamsWithContext(h.ctx).
		WithPackName(name).
		WithRegistryUID(registryUID)
	resp, err := h.Client.V1PacksNameRegistryUIDList(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// GetPack retrieves a pack by UID.
func (h *V1Client) GetPack(uid string) (*models.V1PackTagEntity, error) {
	params := clientv1.NewV1PacksUIDParamsWithContext(h.ctx).
		WithUID(uid)
	resp, err := h.Client.V1PacksUID(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// GetPackRegistry retrieves a pack registry by UID and pack type.
// If packUID is "uid" or packType is "manifest", it returns Spectro Cloud's 'Public Repo' registry.
func (h *V1Client) GetPackRegistry(packUID, packType string) string {
	if packUID == "uid" || packType == "manifest" {
		registry, err := h.GetPackRegistryCommonByName("Public Repo")
		if err != nil {
			return ""
		}
		return registry.UID
	}
	packTagEntity, err := h.GetPack(packUID)
	if err != nil {
		return ""
	}
	return packTagEntity.RegistryUID
}
