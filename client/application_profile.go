package client

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/spectrocloud/hapi/apiutil/transport"
	hashboardC "github.com/spectrocloud/hapi/hashboard/client/v1"
	"github.com/spectrocloud/hapi/models"
	clusterC "github.com/spectrocloud/hapi/spectrocluster/client/v1"

	"github.com/spectrocloud/palette-sdk-go/client/herr"
)

func (h *V1Client) GetApplicationProfileByNameAndVersion(profileName, version string) (*models.V1AppProfileSummary, string, string, error) {
	limit := int64(0)
	params := hashboardC.NewV1DashboardAppProfilesParamsWithContext(h.Ctx).WithLimit(&limit)
	profiles, err := h.GetHashboardClient().V1DashboardAppProfiles(params)
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
	// Unit test mock handler
	if h.GetApplicationProfileFn != nil {
		return h.GetApplicationProfileFn(uid)
	}

	params := clusterC.NewV1AppProfilesUIDGetParamsWithContext(h.Ctx).WithUID(uid)
	response, err := h.GetClusterClient().V1AppProfilesUIDGet(params)
	if err != nil {
		if herr.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}

	return response.Payload, nil
}

func (h *V1Client) GetApplicationProfileTiers(applicationProfileUID string) ([]*models.V1AppTier, error) {
	if h.GetApplicationProfileTiersFn != nil {
		return h.GetApplicationProfileTiersFn(applicationProfileUID)
	}

	params := clusterC.NewV1AppProfilesUIDTiersGetParamsWithContext(h.Ctx).WithUID(applicationProfileUID)
	success, err := h.GetClusterClient().V1AppProfilesUIDTiersGet(params)

	var e *transport.TransportError
	if errors.As(err, &e) && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return success.Payload.Spec.AppTiers, nil
}

func (h *V1Client) GetApplicationProfileTierManifestContent(applicationProfileUID, tierUID, manifestUID string) (string, error) {
	if h.GetApplicationProfileTierManifestContentFn != nil {
		return h.GetApplicationProfileTierManifestContentFn(applicationProfileUID, tierUID, manifestUID)
	}

	params := &clusterC.V1AppProfilesUIDTiersUIDManifestsUIDGetParams{
		UID:         applicationProfileUID,
		TierUID:     tierUID,
		ManifestUID: manifestUID,
		Context:     h.Ctx,
	}

	success, err := h.GetClusterClient().V1AppProfilesUIDTiersUIDManifestsUIDGet(params)

	var e *transport.TransportError
	if errors.As(err, &e) && e.HttpCode == 404 {
		return "", nil
	} else if err != nil {
		return "", err
	}

	return success.Payload.Spec.Published.Content, nil
}

func (h *V1Client) SearchAppProfileSummaries(scope string, filter *models.V1AppProfileFilterSpec, sortBy []*models.V1AppProfileSortSpec) ([]*models.V1AppProfileSummary, error) {
	var params *hashboardC.V1DashboardAppProfilesParams
	switch scope {
	case "project":
		params = hashboardC.NewV1DashboardAppProfilesParamsWithContext(h.Ctx)
	case "tenant":
		params = hashboardC.NewV1DashboardAppProfilesParams()
	}
	params.Body = &models.V1AppProfilesFilterSpec{
		Filter: filter,
		Sort:   sortBy,
	}

	var appProfile []*models.V1AppProfileSummary
	var resp *hashboardC.V1DashboardAppProfilesOK
	var err error
	for {
		if resp != nil {
			params.Offset = &resp.Payload.Listmeta.Offset
		}
		resp, err = h.GetHashboardClient().V1DashboardAppProfiles(params)
		var e *transport.TransportError
		if errors.As(err, &e) && e.HttpCode == 404 {
			return nil, nil
		} else if err != nil {
			return nil, err
		}
		if resp != nil && resp.Payload != nil {
			if resp.Payload.AppProfiles != nil {
				appProfile = append(appProfile, resp.Payload.AppProfiles...)
			}
			if resp.Payload.Listmeta == nil || len(resp.Payload.Listmeta.Continue) == 0 {
				break
			}
		} else {
			break
		}
	}
	return appProfile, nil
}

func (h *V1Client) PatchApplicationProfile(appProfileUID string, metadata *models.V1AppProfileMetaEntity, ProfileContext string) error {
	var params *clusterC.V1AppProfilesUIDMetadataUpdateParams
	switch ProfileContext {
	case "project":
		params = clusterC.NewV1AppProfilesUIDMetadataUpdateParamsWithContext(h.Ctx).WithUID(appProfileUID).WithBody(metadata)
	case "tenant":
		params = clusterC.NewV1AppProfilesUIDMetadataUpdateParams().WithUID(appProfileUID).WithBody(metadata)
	}

	_, err := h.GetClusterClient().V1AppProfilesUIDMetadataUpdate(params)
	return err
}

func (h *V1Client) CreateApplicationProfileTiers(appProfileUID string, appTiers []*models.V1AppTierEntity, ProfileContext string) error {
	var err error
	for _, appTier := range appTiers {
		var params *clusterC.V1AppProfilesUIDTiersCreateParams
		switch ProfileContext {
		case "project":
			params = clusterC.NewV1AppProfilesUIDTiersCreateParamsWithContext(h.Ctx).WithUID(appProfileUID).WithBody(appTier)
		case "tenant":
			params = clusterC.NewV1AppProfilesUIDTiersCreateParams().WithUID(appProfileUID).WithBody(appTier)
		}

		_, tmpErr := h.GetClusterClient().V1AppProfilesUIDTiersCreate(params)
		if tmpErr != nil {
			err = errors.Wrap(err, tmpErr.Error())
		}
	}
	return err
}

func (h *V1Client) UpdateApplicationProfileTiers(appProfileUID, tierUID string, appTier *models.V1AppTierUpdateEntity, ProfileContext string) error {
	var params *clusterC.V1AppProfilesUIDTiersUIDUpdateParams
	switch ProfileContext {
	case "project":
		params = clusterC.NewV1AppProfilesUIDTiersUIDUpdateParamsWithContext(h.Ctx).WithUID(appProfileUID).WithTierUID(tierUID).WithBody(appTier)
	case "tenant":
		params = clusterC.NewV1AppProfilesUIDTiersUIDUpdateParams().WithUID(appProfileUID).WithTierUID(tierUID).WithBody(appTier)
	}

	_, err := h.GetClusterClient().V1AppProfilesUIDTiersUIDUpdate(params)
	var e *transport.TransportError
	if errors.As(err, &e) && e.HttpCode == 404 {
		return nil
	} else if err != nil {
		return err
	}

	return err
}

func (h *V1Client) DeleteApplicationProfileTiers(appProfileUID string, appTiers []string, ProfileContext string) error {
	var err error
	for _, appTierUID := range appTiers {
		var params *clusterC.V1AppProfilesUIDTiersUIDDeleteParams
		switch ProfileContext {
		case "project":
			params = clusterC.NewV1AppProfilesUIDTiersUIDDeleteParamsWithContext(h.Ctx).WithUID(appProfileUID).WithTierUID(appTierUID)
		case "tenant":
			params = clusterC.NewV1AppProfilesUIDTiersUIDDeleteParams().WithUID(appProfileUID).WithTierUID(appTierUID)
		}

		_, tmpErr := h.GetClusterClient().V1AppProfilesUIDTiersUIDDelete(params)
		if tmpErr != nil {
			err = errors.Wrap(err, tmpErr.Error())
		}
	}
	return err
}

func (h *V1Client) CreateApplicationProfile(appProfile *models.V1AppProfileEntity, ProfileContext string) (string, error) {
	// Unit test mock handler
	if h.CreateApplicationProfileFn != nil {
		return h.CreateApplicationProfileFn(appProfile, ProfileContext)
	}

	var params *clusterC.V1AppProfilesCreateParams
	switch ProfileContext {
	case "project":
		params = clusterC.NewV1AppProfilesCreateParams().WithContext(h.Ctx).WithBody(appProfile)
	case "tenant":
		params = clusterC.NewV1AppProfilesCreateParams().WithBody(appProfile)
	}

	success, err := h.GetClusterClient().V1AppProfilesCreate(params)
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

func (h *V1Client) DeleteApplicationProfile(uid string) error {
	// Unit test mock handler
	if h.DeleteApplicationProfileFn != nil {
		return h.DeleteApplicationProfileFn(uid)
	}

	profile, err := h.GetApplicationProfile(uid)
	if err != nil {
		return err
	}

	var params *clusterC.V1AppProfilesUIDDeleteParams
	switch profile.Metadata.Annotations["scope"] {
	case "project":
		params = clusterC.NewV1AppProfilesUIDDeleteParamsWithContext(h.Ctx).WithUID(uid)
	case "tenant":
		params = clusterC.NewV1AppProfilesUIDDeleteParams().WithUID(uid)
	}

	_, err = h.GetClusterClient().V1AppProfilesUIDDelete(params)
	return err
}
