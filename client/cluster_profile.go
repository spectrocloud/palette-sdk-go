package client

import (
	"bytes"
	"fmt"
	"io"
	"os"

	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
	"github.com/spectrocloud/palette-sdk-go/client/apiutil"
)

func (h *V1Client) GetClusterProfileSummary(uid string) (*models.V1ClusterProfileSummary, error) {
	params := clientV1.NewV1ClusterProfilesUIDSummaryParamsWithContext(h.ctx).
		WithUID(uid)
	resp, err := h.Client.V1ClusterProfilesUIDSummary(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

func (h *V1Client) GetClusterProfileUid(profileName, profileVersion string) (string, error) {
	params := clientV1.NewV1ClusterProfilesMetadataParamsWithContext(h.ctx)
	resp, err := h.Client.V1ClusterProfilesMetadata(params)
	if err != nil {
		return "", err
	}
	for _, profile := range resp.Payload.Items {
		if profile.Metadata.Name == profileName && profile.Spec.Version == profileVersion {
			return profile.Metadata.UID, nil
		}
	}
	return "", fmt.Errorf("cluster profile %s not found", profileName)
}

func (h *V1Client) ImportClusterProfile(profileContent string) (string, error) {
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
	params := clientV1.NewV1ClusterProfilesImportFileParamsWithContext(h.ctx).
		WithPublish(apiutil.Ptr(true)).
		WithImportFile(f)
	resp, err := h.Client.V1ClusterProfilesImportFile(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil

}

func (h *V1Client) UpgradeClusterProfile(clusterUid string, body *models.V1SpectroClusterProfiles) error {
	params := clientV1.NewV1SpectroClustersUpdateProfilesParamsWithContext(h.ctx).
		WithUID(clusterUid).
		WithBody(body)
	_, err := h.Client.V1SpectroClustersUpdateProfiles(params)
	return err
}

func (h *V1Client) AttachAddonToCluster(clusterUid, profileUid string, profileUids []string) error {
	// get existing cluster profile uid list on the cluster
	currentClusterInfo, err := h.GetCluster(clusterUid)
	if err != nil {
		return err
	}

	profileList := make([]*models.V1SpectroClusterProfileEntity, 0, len(profileUid))
	packList := make([]*models.V1PackValuesEntity, 0, len(currentClusterInfo.Spec.ClusterProfileTemplates))

	for _, clusterProfile := range currentClusterInfo.Spec.ClusterProfileTemplates {
		for _, value := range clusterProfile.Packs {
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
			UID:        clusterProfile.UID,
		})
	}

	for _, uid := range profileUids {
		packs, err := h.GetPacksByProfile(uid)
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
			UID:        uid,
		})
	}
	profiles := &models.V1SpectroClusterProfiles{
		Profiles: profileList,
	}
	if err := h.UpgradeClusterProfile(clusterUid, profiles); err != nil {
		return err
	}

	return nil
}

func (h *V1Client) DeleteClusterProfile(uid string) error {
	profile, err := h.GetClusterProfile(uid)
	if err != nil {
		return err
	} else if profile == nil {
		return fmt.Errorf("cluster profile %s not found", uid)
	}
	params := clientV1.NewV1ClusterProfilesDeleteParamsWithContext(h.ctx).
		WithUID(uid)
	_, err = h.Client.V1ClusterProfilesDelete(params)
	return err
}

func (h *V1Client) GetClusterProfile(uid string) (*models.V1ClusterProfile, error) {
	params := clientV1.NewV1ClusterProfilesGetParamsWithContext(h.ctx).
		WithUID(uid)
	resp, err := h.Client.V1ClusterProfilesGet(params)
	if apiutil.Is404(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

func (h *V1Client) GetClusterProfiles() ([]*models.V1ClusterProfileMetadata, error) {
	params := clientV1.NewV1ClusterProfilesMetadataParamsWithContext(h.ctx)
	resp, err := h.Client.V1ClusterProfilesMetadata(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload.Items, nil
}

func (h *V1Client) PatchClusterProfile(clusterProfile *models.V1ClusterProfileUpdateEntity, metadata *models.V1ProfileMetaEntity) error {
	params := clientV1.NewV1ClusterProfilesUIDMetadataUpdateParamsWithContext(h.ctx).
		WithUID(clusterProfile.Metadata.UID).
		WithBody(metadata)
	_, err := h.Client.V1ClusterProfilesUIDMetadataUpdate(params)
	return err
}

func (h *V1Client) UpdateClusterProfile(clusterProfile *models.V1ClusterProfileUpdateEntity) error {
	params := clientV1.NewV1ClusterProfilesUpdateParamsWithContext(h.ctx).
		WithUID(clusterProfile.Metadata.UID).
		WithBody(clusterProfile)
	_, err := h.Client.V1ClusterProfilesUpdate(params)
	return err
}

func (h *V1Client) CreateClusterProfile(clusterProfile *models.V1ClusterProfileEntity) (string, error) {
	params := clientV1.NewV1ClusterProfilesCreateParamsWithContext(h.ctx).
		WithBody(clusterProfile)
	resp, err := h.Client.V1ClusterProfilesCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

func (h *V1Client) PublishClusterProfile(uid string) error {
	params := clientV1.NewV1ClusterProfilesPublishParamsWithContext(h.ctx).
		WithUID(uid)
	_, err := h.Client.V1ClusterProfilesPublish(params)
	return err
}

func (h *V1Client) GetProfileVariables(uid string) ([]*models.V1Variable, error) {
	params := clientV1.NewV1ClusterProfilesUIDVariablesGetParamsWithContext(h.ctx).WithUID(uid)
	resp, err := h.Client.V1ClusterProfilesUIDVariablesGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload.Variables, nil
}

func (h *V1Client) UpdateProfileVariables(variables *models.V1Variables, uid string) error {
	params := clientV1.NewV1ClusterProfilesUIDVariablesPutParamsWithContext(h.ctx).
		WithUID(uid).
		WithBody(variables)
	_, err := h.Client.V1ClusterProfilesUIDVariablesPut(params)
	return err
}

func (h *V1Client) ExportClusterProfile(uid, format string) (bytes.Buffer, error) {
	var buf bytes.Buffer
	writer := io.Writer(&buf)
	params := clientV1.NewV1ClusterProfilesUIDExportParamsWithContext(h.ctx).
		WithUID(uid).
		WithFormat(&format)
	_, err := h.Client.V1ClusterProfilesUIDExport(params, writer)
	return buf, err
}

func (h *V1Client) BulkDeleteClusterProfiles(uids []string) (*models.V1BulkDeleteResponse, error) {
	params := clientV1.NewV1ClusterProfilesBulkDeleteParamsWithContext(h.ctx).
		WithBody(&models.V1BulkDeleteRequest{
			Uids: uids,
		})
	resp, err := h.Client.V1ClusterProfilesBulkDelete(params)
	if err != nil {
		return nil, err
	}
	return resp.GetPayload(), nil
}
