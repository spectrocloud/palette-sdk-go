package client

import (
	"errors"
	"fmt"
	"strings"

	"github.com/spectrocloud/palette-api-go/apiutil/transport"
	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
)

func (h *V1Client) SearchPacks(filter *models.V1PackFilterSpec, sortBy []*models.V1PackSortSpec) ([]*models.V1PackMetadata, error) {
	body := &models.V1PacksFilterSpec{
		Filter: filter,
		Sort:   sortBy,
	}
	params := clientV1.NewV1PacksSearchParams().WithContext(h.Ctx).WithBody(body)

	resp, err := h.Client.V1PacksSearch(params)
	var e *transport.TransportError
	if errors.As(err, &e) && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return resp.Payload.Items, nil
}

func (h *V1Client) GetClusterProfileManifestPack(clusterProfileUID, packName string) ([]*models.V1ManifestEntity, error) {
	//params := clientV1.NewV1ClusterProfilesGetParamsWithContext(h.ctx).WithUID(uid)
	params := clientV1.NewV1ClusterProfilesUIDPacksUIDManifestsParamsWithContext(h.Ctx).
		WithUID(clusterProfileUID).WithPackName(packName)

	success, err := h.Client.V1ClusterProfilesUIDPacksUIDManifests(params)
	var e *transport.TransportError
	if errors.As(err, &e) && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return success.Payload.Items, nil
}

func (h *V1Client) GetPacks(filters []string, registryUID string) ([]*models.V1PackSummary, error) {
	params := clientV1.NewV1PacksSummaryListParamsWithContext(h.Ctx)
	if filters != nil {
		filterString := Ptr(strings.Join(filters, "AND"))
		params = params.WithFilters(filterString)
	}

	response, err := h.Client.V1PacksSummaryList(params)
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

func (h *V1Client) GetPacksByProfile(scope, profileUid string) ([]*models.V1ClusterProfilePacksEntity, error) {
	var params *clientV1.V1ClusterProfilesUIDPacksGetParams
	switch scope {
	case "project":
		params = clientV1.NewV1ClusterProfilesUIDPacksGetParamsWithContext(h.Ctx)
	case "tenant":
		params = clientV1.NewV1ClusterProfilesUIDPacksGetParams()
	default:
		return nil, fmt.Errorf("invalid scope: %s", scope)
	}
	params = params.WithUID(profileUid)

	resp, err := h.Client.V1ClusterProfilesUIDPacksGet(params)
	if err != nil {
		return nil, err
	} else if resp.Payload == nil {
		return nil, fmt.Errorf("no packs found for profile %s", profileUid)
	}
	return resp.Payload.Items, nil
}

func (h *V1Client) GetPacksByNameAndRegistry(name, registryUID, packContext string) (*models.V1PackTagEntity, error) {
	var params *clientV1.V1PacksNameRegistryUIDListParams
	switch packContext {
	case "project":
		params = clientV1.NewV1PacksNameRegistryUIDListParamsWithContext(h.Ctx).WithPackName(name).WithRegistryUID(registryUID)
	case "tenant":
		params = clientV1.NewV1PacksNameRegistryUIDListParams().WithPackName(name).WithRegistryUID(registryUID)
	default:
		return nil, fmt.Errorf("invalid pack context %s", packContext)

	}

	response, err := h.Client.V1PacksNameRegistryUIDList(params)
	if err != nil {
		return nil, err
	}

	return response.Payload, nil
}

func (h *V1Client) GetPack(uid string) (*models.V1PackTagEntity, error) {
	params := clientV1.NewV1PacksUIDParamsWithContext(h.Ctx).WithUID(uid)
	response, err := h.Client.V1PacksUID(params)
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
