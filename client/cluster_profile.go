package client

import (
	"errors"
	"fmt"
	"os"

	"github.com/spectrocloud/palette-api-go/apiutil/transport"
	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
)

func (h *V1Client) GetClusterProfileSummary(scope, uid string) (*models.V1ClusterProfileSummary, error) {
	var params *clientV1.V1ClusterProfilesUIDSummaryParams
	switch scope {
	case "project":
		params = clientV1.NewV1ClusterProfilesUIDSummaryParamsWithContext(h.Ctx).WithUID(uid)
	case "tenant":
		params = clientV1.NewV1ClusterProfilesUIDSummaryParams().WithUID(uid)
	default:
		return nil, errors.New("invalid scope")
	}

	resp, err := h.Client.V1ClusterProfilesUIDSummary(params)
	if err != nil {
		return nil, err
	}

	return resp.GetPayload(), nil
}

func (h *V1Client) GetClusterProfileUid(scope, profileName, profileVersion string) (string, error) {
	var params *clientV1.V1ClusterProfilesMetadataParams
	switch scope {
	case "project":
		params = clientV1.NewV1ClusterProfilesMetadataParamsWithContext(h.Ctx)
	case "tenant":
		params = clientV1.NewV1ClusterProfilesMetadataParams()
	default:
		return "", errors.New("invalid scope")
	}

	resp, err := h.Client.V1ClusterProfilesMetadata(params)
	if err != nil {
		return "", err
	}
	profiles := resp.GetPayload()
	if profiles == nil {
		return "", fmt.Errorf("no cluster profiles found")
	}
	for _, profile := range profiles.Items {
		if profile.Metadata.Name == profileName && profile.Spec.Version == profileVersion {
			return profile.Metadata.UID, nil
		}
	}
	return "", fmt.Errorf("cluster profile %s not found", profileName)
}

func (h *V1Client) ImportClusterProfile(scope, profileContent string) (string, error) {
	tmpFile, err := os.CreateTemp(os.TempDir(), "profile-import")
	if err != nil {
		return "", err
	}
	defer os.Remove(tmpFile.Name())
	if _, err = tmpFile.Write([]byte(profileContent)); err != nil {
		return "", err
	}
	f, err := os.Open(tmpFile.Name())
	if err != nil {
		return "", err
	}

	var params *clientV1.V1ClusterProfilesImportFileParams
	switch scope {
	case "project":
		params = clientV1.NewV1ClusterProfilesImportFileParamsWithContext(h.Ctx)
	case "tenant":
		params = clientV1.NewV1ClusterProfilesImportFileParams()
	default:
		return "", errors.New("invalid scope")
	}

	publish := true
	params = params.WithPublish(&publish).WithImportFile(f)

	resp, err := h.Client.V1ClusterProfilesImportFile(params)
	if err != nil {
		return "", err
	}
	uid := resp.GetPayload()
	if uid == nil {
		return "", errors.New("import failed")
	}
	return *uid.UID, nil

}

func (h *V1Client) UpgradeClusterProfile(scope, clusterUid string, body *models.V1SpectroClusterProfiles) error {
	var params *clientV1.V1SpectroClustersUpdateProfilesParams
	switch scope {
	case "project":
		params = clientV1.NewV1SpectroClustersUpdateProfilesParamsWithContext(h.Ctx)
	case "tenant":
		params = clientV1.NewV1SpectroClustersUpdateProfilesParams()
	default:
		return fmt.Errorf("invalid scope: %s", scope)
	}
	params = params.WithUID(clusterUid).WithBody(body)

	if _, err := h.Client.V1SpectroClustersUpdateProfiles(params); err != nil {
		return err
	}
	return nil
}

func (h *V1Client) AttachAddonToCluster(scope, clusterUid, existingProfileUid string, profileIdList []string) error {
	// get existing cluster profile uid list on the cluster
	currentClusterInfo, err := h.GetCluster(scope, clusterUid)
	if err != nil {
		return err
	}

	profileList := make([]*models.V1SpectroClusterProfileEntity, 0, len(existingProfileUid))
	packList := make([]*models.V1PackValuesEntity, 0, len(currentClusterInfo.Spec.ClusterProfileTemplates))

	for _, existingClusterProfile := range currentClusterInfo.Spec.ClusterProfileTemplates {
		for _, value := range existingClusterProfile.Packs {
			packList = append(packList, &models.V1PackValuesEntity{
				Manifests: nil,
				Name:      value.Name,
				Tag:       value.Tag,
				//Type:      models.NewV1PackType(models.V1PackType(value.Type)),
				Type:   models.V1PackType(value.Type),
				Values: value.Values,
			})
		}
		profileList = append(profileList, &models.V1SpectroClusterProfileEntity{
			PackValues: packList,
			UID:        existingClusterProfile.UID,
		})
	}

	for _, profileUid := range profileIdList {
		packs, err := h.GetPacksByProfile(scope, profileUid)
		if err != nil {
			return err
		}
		packList = make([]*models.V1PackValuesEntity, 0, len(packs))
		for _, value := range packs {
			packList = append(packList, &models.V1PackValuesEntity{
				Manifests: nil,
				Name:      &value.Spec.Name,
				Tag:       value.Spec.Version,
				Type:      value.Spec.Type,
				Values:    value.Spec.Values,
			})
		}
		profileList = append(profileList, &models.V1SpectroClusterProfileEntity{
			PackValues: packList,
			UID:        profileUid,
		})
	}
	profiles := &models.V1SpectroClusterProfiles{
		Profiles: profileList,
	}
	if err := h.UpgradeClusterProfile(scope, clusterUid, profiles); err != nil {
		return err
	}

	return nil
}

func (h *V1Client) DeleteClusterProfile(client clientV1.ClientService, uid string) error {
	profile, err := h.GetClusterProfile(client, uid)
	if err != nil || profile == nil {
		return err
	}

	var params *clientV1.V1ClusterProfilesDeleteParams
	switch profile.Metadata.Annotations["scope"] {
	case "project":
		params = clientV1.NewV1ClusterProfilesDeleteParamsWithContext(h.Ctx).WithUID(uid)
	case "tenant":
		params = clientV1.NewV1ClusterProfilesDeleteParams().WithUID(uid)
	default:
		return errors.New("invalid scope")
	}

	_, err = client.V1ClusterProfilesDelete(params)
	return err
}

func (h *V1Client) GetClusterProfile(client clientV1.ClientService, uid string) (*models.V1ClusterProfile, error) {
	// no need to switch request context here as /v1/clusterprofiles/{uid} works for profile in any scope.
	params := clientV1.NewV1ClusterProfilesGetParamsWithContext(h.Ctx).WithUID(uid)
	success, err := client.V1ClusterProfilesGet(params)
	var e *transport.TransportError
	if errors.As(err, &e) && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return success.Payload, nil
}

func (h *V1Client) GetClusterProfiles(client clientV1.ClientService) ([]*models.V1ClusterProfileMetadata, error) {
	params := clientV1.NewV1ClusterProfilesMetadataParamsWithContext(h.Ctx)
	response, err := client.V1ClusterProfilesMetadata(params)
	if err != nil {
		return nil, err
	}
	return response.Payload.Items, nil
}

func (h *V1Client) PatchClusterProfile(client clientV1.ClientService, clusterProfile *models.V1ClusterProfileUpdateEntity, metadata *models.V1ProfileMetaEntity, ProfileContext string) error {
	uid := clusterProfile.Metadata.UID
	var params *clientV1.V1ClusterProfilesUIDMetadataUpdateParams
	switch ProfileContext {
	case "project":
		params = clientV1.NewV1ClusterProfilesUIDMetadataUpdateParamsWithContext(h.Ctx).WithUID(uid).WithBody(metadata)
	case "tenant":
		params = clientV1.NewV1ClusterProfilesUIDMetadataUpdateParams().WithUID(uid).WithBody(metadata)
	default:
		return errors.New("invalid scope")
	}

	_, err := client.V1ClusterProfilesUIDMetadataUpdate(params)
	return err
}

func (h *V1Client) UpdateClusterProfile(client clientV1.ClientService, clusterProfile *models.V1ClusterProfileUpdateEntity, ProfileContext string) error {
	uid := clusterProfile.Metadata.UID
	var params *clientV1.V1ClusterProfilesUpdateParams
	switch ProfileContext {
	case "project":
		params = clientV1.NewV1ClusterProfilesUpdateParamsWithContext(h.Ctx).WithUID(uid).WithBody(clusterProfile)
	case "tenant":
		params = clientV1.NewV1ClusterProfilesUpdateParams().WithUID(uid).WithBody(clusterProfile)
	default:
		return errors.New("invalid scope")
	}

	_, err := client.V1ClusterProfilesUpdate(params)
	return err
}

func (h *V1Client) CreateClusterProfile(client clientV1.ClientService, clusterProfile *models.V1ClusterProfileEntity, ProfileContext string) (string, error) {
	var params *clientV1.V1ClusterProfilesCreateParams
	switch ProfileContext {
	case "project":
		params = clientV1.NewV1ClusterProfilesCreateParamsWithContext(h.Ctx).WithBody(clusterProfile)
	case "tenant":
		params = clientV1.NewV1ClusterProfilesCreateParams().WithBody(clusterProfile)
	default:
		return "", errors.New("invalid scope")
	}

	success, err := client.V1ClusterProfilesCreate(params)
	if err != nil {
		return "", err
	}
	return *success.Payload.UID, nil
}

func (h *V1Client) PublishClusterProfile(client clientV1.ClientService, uid, ProfileContext string) error {
	var params *clientV1.V1ClusterProfilesPublishParams
	switch ProfileContext {
	case "project":
		params = clientV1.NewV1ClusterProfilesPublishParamsWithContext(h.Ctx).WithUID(uid)
	case "tenant":
		params = clientV1.NewV1ClusterProfilesPublishParams().WithUID(uid)
	default:
		return errors.New("invalid scope")
	}

	_, err := client.V1ClusterProfilesPublish(params)
	return err
}
