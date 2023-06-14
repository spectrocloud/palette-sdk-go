package client

import (
	"errors"
	"fmt"
	"github.com/spectrocloud/hapi/models"
	userC "github.com/spectrocloud/hapi/user/client/v1"
)

func (h *V1Client) CreateSSHKey(body *models.V1UserAssetSSH, ProfileContext string) (string, error) {
	client, err := h.GetUserClient()
	SSHEntity := &models.V1UserAssetSSHEntity{
		Metadata: &models.V1ObjectMetaInputEntity{
			Annotations: nil,
			Labels:      nil,
			Name:        body.Metadata.Name,
		},
		Spec: &models.V1UserAssetSSHSpec{PublicKey: body.Spec.PublicKey},
	}
	if err != nil {
		return "", err
	}
	params := userC.NewV1UserAssetsSSHCreateParams()
	switch ProfileContext {
	case "project":
		params = userC.NewV1UserAssetsSSHCreateParamsWithContext(h.Ctx).WithBody(SSHEntity)
	case "tenant":
		params = userC.NewV1UserAssetsSSHCreateParams().WithBody(SSHEntity)
	default:
		return "", errors.New("invalid scope")
	}

	success, err := client.V1UserAssetsSSHCreate(params)
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

func (h *V1Client) GetSSHKeyUID(SSHKeyName string) (string, error) {
	client, err := h.GetUserClient()
	SSHKeys, err := client.V1UsersAssetSSHGet(nil)
	if err != nil {
		return "", err
	}
	print(SSHKeys)
	//for _, key := range SSHKeys.Get {
	//	if key.Metadata.Name == projectName {
	//		return project.Metadata.UID, nil
	//	}
	//}

	return "", fmt.Errorf("project '%s' not found", SSHKeyName)
}

func (h *V1Client) GetSSHKeyByUID(uid string, ProfileContext string) (*models.V1UserAssetSSH, error) {
	client, err := h.GetUserClient()
	if err != nil {
		return nil, err
	}
	params := userC.NewV1UsersAssetSSHGetParams()
	switch ProfileContext {
	case "project":
		params = userC.NewV1UsersAssetSSHGetParamsWithContext(h.Ctx).WithUID(uid)
	case "tenant":
		params = userC.NewV1UsersAssetSSHGetParams().WithUID(uid)
	default:
		return nil, errors.New("invalid scope")
	}
	SSHKey, err := client.V1UsersAssetSSHGet(params)
	if err != nil || SSHKey == nil {
		return nil, err
	}

	return SSHKey.Payload, nil
}

//func (h *V1Client) GetSSHKeys() (*models.V1ProjectsMetadata, error) {
//	client, err := h.GetHashboardClient()
//	if err != nil {
//		return nil, err
//	}
//
//	params := hashboardC.NewV1ProjectsMetadataParams()
//
//	projects, err := client.V1ProjectsMetadata(params)
//	if err != nil || projects == nil {
//		// to support 2.6 projects list
//		if e, ok := err.(*transport.TransportError); ok && e.HttpCode == 404 {
//			limit := int64(0)
//			userClient, err := h.GetUserClient()
//			if err != nil {
//				return nil, err
//			}
//			oldParams := userC.NewV1ProjectsListParams().WithLimit(&limit)
//			oldProjects, err := userClient.V1ProjectsList(oldParams)
//			if err != nil || oldProjects == nil {
//				return nil, err
//			}
//			ret := make([]*models.V1ProjectMetadata, 0)
//			for _, pr := range oldProjects.Payload.Items {
//				ret = append(ret, &models.V1ProjectMetadata{
//					Metadata: &models.V1ObjectEntity{
//						UID:  pr.Metadata.UID,
//						Name: pr.Metadata.Name,
//					},
//				})
//			}
//			return &models.V1ProjectsMetadata{
//				Items: ret,
//			}, nil
//		}
//		return nil, err
//	}
//
//	return projects.Payload, nil
//}

func (h *V1Client) UpdateSSHKey(uid string, body *models.V1UserAssetSSH, ProfileContext string) error {
	client, err := h.GetUserClient()
	if err != nil {
		return err
	}
	params := userC.NewV1UsersAssetSSHUpdateParamsWithContext(h.Ctx).WithUID(uid).WithBody(body)
	switch ProfileContext {
	case "project":
		params = userC.NewV1UsersAssetSSHUpdateParamsWithContext(h.Ctx).WithUID(uid).WithBody(body)
	case "tenant":
		params = userC.NewV1UsersAssetSSHUpdateParams().WithUID(uid).WithBody(body)
	default:
		return errors.New("invalid scope")
	}
	_, err = client.V1UsersAssetSSHUpdate(params)
	if err != nil {
		return err
	}

	return nil
}

func (h *V1Client) DeleteSSHKey(uid string, ProfileContext string) error {
	client, err := h.GetUserClient()
	if err != nil {
		return err
	}
	params := userC.NewV1UsersAssetSSHDeleteParamsWithContext(h.Ctx).WithUID(uid)
	switch ProfileContext {
	case "project":
		params = userC.NewV1UsersAssetSSHDeleteParamsWithContext(h.Ctx).WithUID(uid)
	case "tenant":
		params = userC.NewV1UsersAssetSSHDeleteParams().WithUID(uid)
	default:
		return errors.New("invalid scope")
	}
	_, err = client.V1UsersAssetSSHDelete(params)
	if err != nil {
		return err
	}

	return nil
}
