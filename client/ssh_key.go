package client

import (
	"errors"
	"fmt"
	"github.com/spectrocloud/hapi/models"
	userC "github.com/spectrocloud/hapi/user/client/v1"
)

func (h *V1Client) CreateSSHKey(body *models.V1UserAssetSSH, SSHKeyContext string) (string, error) {
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
	switch SSHKeyContext {
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

func (h *V1Client) GetSSHKeyByName(SSHKeyName string, SSHKeyContext string) (*models.V1UserAssetSSH, error) {
	client, err := h.GetUserClient()
	if err != nil {
		return nil, err
	}
	var params *userC.V1UsersAssetsSSHGetParams
	switch SSHKeyContext {
	case "project":
		params = userC.NewV1UsersAssetsSSHGetParamsWithContext(h.Ctx)
	case "tenant":
		params = userC.NewV1UsersAssetsSSHGetParams()
	default:
		return nil, errors.New("invalid scope")
	}
	SSHKeys, err := client.V1UsersAssetsSSHGet(params)
	if err != nil {
		return nil, err
	}
	for _, key := range SSHKeys.Payload.Items {
		if key.Metadata.Name == SSHKeyName {
			return key, nil
		}
	}
	return nil, fmt.Errorf("project '%s' not found", SSHKeyName)
}

func (h *V1Client) GetSSHKeyByUID(uid string, SSHKeyContext string) (*models.V1UserAssetSSH, error) {
	client, err := h.GetUserClient()
	if err != nil {
		return nil, err
	}
	var params *userC.V1UsersAssetSSHGetUIDParams
	switch SSHKeyContext {
	case "project":
		params = userC.NewV1UsersAssetSSHGetUIDParamsWithContext(h.Ctx).WithUID(uid)
	case "tenant":
		params = userC.NewV1UsersAssetSSHGetUIDParams().WithUID(uid)
	default:
		return nil, errors.New("invalid scope")
	}
	SSHKey, err := client.V1UsersAssetSSHGetUID(params)
	if err != nil || SSHKey == nil {
		return nil, err
	}

	return SSHKey.Payload, nil
}

func (h *V1Client) UpdateSSHKey(uid string, body *models.V1UserAssetSSH, SSHKeyContext string) error {
	client, err := h.GetUserClient()
	if err != nil {
		return err
	}
	var params *userC.V1UsersAssetSSHUpdateParams
	switch SSHKeyContext {
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

func (h *V1Client) DeleteSSHKey(uid string, SSHKeyContext string) error {
	client, err := h.GetUserClient()
	if err != nil {
		return err
	}
	var params *userC.V1UsersAssetSSHDeleteParams
	switch SSHKeyContext {
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
