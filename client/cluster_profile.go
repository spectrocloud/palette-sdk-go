package client

import (
	"errors"

	"github.com/spectrocloud/hapi/apiutil/transport"
	hashboardC "github.com/spectrocloud/hapi/hashboard/client/v1"
	"github.com/spectrocloud/hapi/models"
	clusterC "github.com/spectrocloud/hapi/spectrocluster/client/v1"
)

func (h *V1Client) DeleteClusterProfile(client clusterC.ClientService, uid string) error {
	profile, err := h.GetClusterProfile(client, uid)
	if err != nil || profile == nil {
		return err
	}

	var params *clusterC.V1ClusterProfilesDeleteParams
	switch profile.Metadata.Annotations["scope"] {
	case "project":
		params = clusterC.NewV1ClusterProfilesDeleteParamsWithContext(h.Ctx).WithUID(uid)
	case "tenant":
		params = clusterC.NewV1ClusterProfilesDeleteParams().WithUID(uid)
	default:
		return errors.New("invalid scope")
	}

	_, err = client.V1ClusterProfilesDelete(params)
	return err
}

func (h *V1Client) GetClusterProfile(client clusterC.ClientService, uid string) (*models.V1ClusterProfile, error) {
	// no need to switch request context here as /v1/clusterprofiles/{uid} works for profile in any scope.
	params := clusterC.NewV1ClusterProfilesGetParamsWithContext(h.Ctx).WithUID(uid)
	success, err := client.V1ClusterProfilesGet(params)
	var e *transport.TransportError
	if errors.As(err, &e) && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return success.Payload, nil
}

func (h *V1Client) GetClusterProfiles(client hashboardC.ClientService) ([]*models.V1ClusterProfileMetadata, error) {
	params := hashboardC.NewV1ClusterProfilesMetadataParamsWithContext(h.Ctx)
	response, err := client.V1ClusterProfilesMetadata(params)
	if err != nil {
		return nil, err
	}
	return response.Payload.Items, nil
}

func (h *V1Client) PatchClusterProfile(client clusterC.ClientService, clusterProfile *models.V1ClusterProfileUpdateEntity, metadata *models.V1ProfileMetaEntity, ProfileContext string) error {
	uid := clusterProfile.Metadata.UID
	var params *clusterC.V1ClusterProfilesUIDMetadataUpdateParams
	switch ProfileContext {
	case "project":
		params = clusterC.NewV1ClusterProfilesUIDMetadataUpdateParamsWithContext(h.Ctx).WithUID(uid).WithBody(metadata)
	case "tenant":
		params = clusterC.NewV1ClusterProfilesUIDMetadataUpdateParams().WithUID(uid).WithBody(metadata)
	default:
		return errors.New("invalid scope")
	}

	_, err := client.V1ClusterProfilesUIDMetadataUpdate(params)
	return err
}

func (h *V1Client) UpdateClusterProfile(client clusterC.ClientService, clusterProfile *models.V1ClusterProfileUpdateEntity, ProfileContext string) error {
	uid := clusterProfile.Metadata.UID
	var params *clusterC.V1ClusterProfilesUpdateParams
	switch ProfileContext {
	case "project":
		params = clusterC.NewV1ClusterProfilesUpdateParamsWithContext(h.Ctx).WithUID(uid).WithBody(clusterProfile)
	case "tenant":
		params = clusterC.NewV1ClusterProfilesUpdateParams().WithUID(uid).WithBody(clusterProfile)
	default:
		return errors.New("invalid scope")
	}

	_, err := client.V1ClusterProfilesUpdate(params)
	return err
}

func (h *V1Client) CreateClusterProfile(client clusterC.ClientService, clusterProfile *models.V1ClusterProfileEntity, ProfileContext string) (string, error) {
	var params *clusterC.V1ClusterProfilesCreateParams
	switch ProfileContext {
	case "project":
		params = clusterC.NewV1ClusterProfilesCreateParamsWithContext(h.Ctx).WithBody(clusterProfile)
	case "tenant":
		params = clusterC.NewV1ClusterProfilesCreateParams().WithBody(clusterProfile)
	default:
		return "", errors.New("invalid scope")
	}

	success, err := client.V1ClusterProfilesCreate(params)
	if err != nil {
		return "", err
	}
	return *success.Payload.UID, nil
}

func (h *V1Client) PublishClusterProfile(client clusterC.ClientService, uid, ProfileContext string) error {
	var params *clusterC.V1ClusterProfilesPublishParams
	switch ProfileContext {
	case "project":
		params = clusterC.NewV1ClusterProfilesPublishParamsWithContext(h.Ctx).WithUID(uid)
	case "tenant":
		params = clusterC.NewV1ClusterProfilesPublishParams().WithUID(uid)
	default:
		return errors.New("invalid scope")
	}

	_, err := client.V1ClusterProfilesPublish(params)
	return err
}

func (h *V1Client) GetProfileVariables(client clusterC.ClientService, uid string) ([]*models.V1Variable, error) {
	params := clusterC.NewV1ClusterProfilesUIDVariablesGetParamsWithContext(h.Ctx).WithUID(uid)
	success, err := client.V1ClusterProfilesUIDVariablesGet(params)
	if err != nil {
		return nil, err
	}
	return success.Payload.Variables, nil
}

func (h *V1Client) UpdateProfileVariables(client clusterC.ClientService, variables *models.V1Variables, uid string) error {
	params := clusterC.NewV1ClusterProfilesUIDVariablesPutParamsWithContext(h.Ctx).WithUID(uid).WithBody(variables)
	_, err := client.V1ClusterProfilesUIDVariablesPut(params)
	if err != nil {
		return err
	}
	return nil
}
