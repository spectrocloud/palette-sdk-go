package client

import (
	"errors"

	"github.com/spectrocloud/palette-api-go/apiutil/transport"
	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
)

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
