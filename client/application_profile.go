package client

import (
	"fmt"

	"github.com/pkg/errors"

	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
	"github.com/spectrocloud/palette-sdk-go/client/apiutil"
	"github.com/spectrocloud/palette-sdk-go/client/herr"
)

func (h *V1Client) GetApplicationProfileByNameAndVersion(profileName, version string) (*models.V1AppProfileSummary, string, string, error) {
	params := clientV1.NewV1DashboardAppProfilesParamsWithContext(h.ctx).
		WithLimit(apiutil.Ptr(int64(0)))
	profiles, err := h.Client.V1DashboardAppProfiles(params)
	if err != nil {
		return nil, "", "", err
	}
	for _, profile := range profiles.Payload.AppProfiles {
		if profile.Metadata.Name == profileName {
			for i := range profile.Spec.Versions {
				if profile.Spec.Versions[i].Version == version {
					return profile, profile.Spec.Versions[i].UID, profile.Spec.Versions[i].Version, nil
				}
			}
		}
	}
	return nil, "", "", fmt.Errorf("application profile '%s' not found for version '%s'", profileName, version)
}

func (h *V1Client) GetApplicationProfile(uid string) (*models.V1AppProfile, error) {
	params := clientV1.NewV1AppProfilesUIDGetParamsWithContext(h.ctx).
		WithUID(uid)
	resp, err := h.Client.V1AppProfilesUIDGet(params)
	if err != nil {
		if herr.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}
	return resp.Payload, nil
}

func (h *V1Client) GetApplicationProfileTiers(applicationProfileUID string) ([]*models.V1AppTier, error) {
	params := clientV1.NewV1AppProfilesUIDTiersGetParamsWithContext(h.ctx).
		WithUID(applicationProfileUID)
	resp, err := h.Client.V1AppProfilesUIDTiersGet(params)
	if apiutil.Is404(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return resp.Payload.Spec.AppTiers, nil
}

func (h *V1Client) GetApplicationProfileTierManifestContent(applicationProfileUID, tierUID, manifestUID string) (string, error) {
	params := clientV1.NewV1AppProfilesUIDTiersUIDManifestsUIDGetParamsWithContext(h.ctx).
		WithUID(applicationProfileUID).
		WithTierUID(tierUID).
		WithManifestUID(manifestUID)
	resp, err := h.Client.V1AppProfilesUIDTiersUIDManifestsUIDGet(params)
	if apiutil.Is404(err) {
		return "", nil
	} else if err != nil {
		return "", err
	}
	return resp.Payload.Spec.Published.Content, nil
}

func (h *V1Client) SearchAppProfileSummaries(filter *models.V1AppProfileFilterSpec, sortBy []*models.V1AppProfileSortSpec) ([]*models.V1AppProfileSummary, error) {
	params := clientV1.NewV1DashboardAppProfilesParamsWithContext(h.ctx).
		WithBody(&models.V1AppProfilesFilterSpec{
			Filter: filter,
			Sort:   sortBy,
		})
	var appProfiles []*models.V1AppProfileSummary
	var resp *clientV1.V1DashboardAppProfilesOK
	var err error
	for {
		if resp != nil {
			params.Offset = &resp.Payload.Listmeta.Offset
		}
		resp, err = h.Client.V1DashboardAppProfiles(params)
		if apiutil.Is404(err) {
			return nil, nil
		} else if err != nil {
			return nil, err
		}
		if resp != nil && resp.Payload != nil {
			if resp.Payload.AppProfiles != nil {
				appProfiles = append(appProfiles, resp.Payload.AppProfiles...)
			}
			if resp.Payload.Listmeta == nil || len(resp.Payload.Listmeta.Continue) == 0 {
				break
			}
		} else {
			break
		}
	}
	return appProfiles, nil
}

func (h *V1Client) PatchApplicationProfile(appProfileUID string, metadata *models.V1AppProfileMetaEntity) error {
	params := clientV1.NewV1AppProfilesUIDMetadataUpdateParamsWithContext(h.ctx).
		WithUID(appProfileUID).
		WithBody(metadata)
	_, err := h.Client.V1AppProfilesUIDMetadataUpdate(params)
	return err
}

func (h *V1Client) CreateApplicationProfileTiers(appProfileUID string, appTiers []*models.V1AppTierEntity) error {
	var err error
	for _, appTier := range appTiers {
		params := clientV1.NewV1AppProfilesUIDTiersCreateParamsWithContext(h.ctx).
			WithUID(appProfileUID).
			WithBody(appTier)
		_, tmpErr := h.Client.V1AppProfilesUIDTiersCreate(params)
		if tmpErr != nil {
			err = errors.Wrap(err, tmpErr.Error())
		}
	}
	return err
}

func (h *V1Client) UpdateApplicationProfileTiers(appProfileUID, tierUID string, appTier *models.V1AppTierUpdateEntity) error {
	params := clientV1.NewV1AppProfilesUIDTiersUIDUpdateParamsWithContext(h.ctx).
		WithUID(appProfileUID).
		WithTierUID(tierUID).
		WithBody(appTier)
	_, err := h.Client.V1AppProfilesUIDTiersUIDUpdate(params)
	if apiutil.Is404(err) {
		return nil
	}
	return err
}

func (h *V1Client) DeleteApplicationProfileTiers(appProfileUID string, appTiers []string) error {
	var err error
	for _, appTierUID := range appTiers {
		params := clientV1.NewV1AppProfilesUIDTiersUIDDeleteParamsWithContext(h.ctx).
			WithUID(appProfileUID).
			WithTierUID(appTierUID)
		_, tmpErr := h.Client.V1AppProfilesUIDTiersUIDDelete(params)
		if tmpErr != nil {
			err = errors.Wrap(err, tmpErr.Error())
		}
	}
	return err
}

func (h *V1Client) CreateApplicationProfile(appProfile *models.V1AppProfileEntity) (string, error) {
	params := clientV1.NewV1AppProfilesCreateParamsWithContext(h.ctx).
		WithBody(appProfile)
	resp, err := h.Client.V1AppProfilesCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

func (h *V1Client) DeleteApplicationProfile(uid string) error {
	profile, err := h.GetApplicationProfile(uid)
	if err != nil {
		return err
	}
	params := clientV1.NewV1AppProfilesUIDDeleteParams().
		WithContext(ContextForScope(profile.Metadata.Annotations[Scope], h.projectUID)).
		WithUID(uid)
	_, err = h.Client.V1AppProfilesUIDDelete(params)
	return err
}
