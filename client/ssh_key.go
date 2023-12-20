package client

import (
	"errors"
	"fmt"

	"github.com/spectrocloud/hapi/models"
	userC "github.com/spectrocloud/hapi/user/client/v1"
)

func (h *V1Client) CreateSSHKey(body *models.V1UserAssetSSH, scope string) (string, error) {
	sshEntity := &models.V1UserAssetSSHEntity{
		Metadata: &models.V1ObjectMetaInputEntity{
			Annotations: nil,
			Labels:      nil,
			Name:        body.Metadata.Name,
		},
		Spec: &models.V1UserAssetSSHSpec{PublicKey: body.Spec.PublicKey},
	}

	var params *userC.V1UserAssetsSSHCreateParams
	switch scope {
	case "project":
		params = userC.NewV1UserAssetsSSHCreateParamsWithContext(h.Ctx).WithBody(sshEntity)
	case "tenant":
		params = userC.NewV1UserAssetsSSHCreateParams().WithBody(sshEntity)
	default:
		return "", errors.New("invalid scope")
	}

	success, err := h.GetUserClient().V1UserAssetsSSHCreate(params)
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

func (h *V1Client) GetSSHKeyByName(name, scope string) (*models.V1UserAssetSSH, error) {
	var params *userC.V1UsersAssetsSSHGetParams
	switch scope {
	case "project":
		params = userC.NewV1UsersAssetsSSHGetParamsWithContext(h.Ctx)
	case "tenant":
		params = userC.NewV1UsersAssetsSSHGetParams()
	default:
		return nil, errors.New("invalid scope")
	}

	sshKeys, err := h.GetUserClient().V1UsersAssetsSSHGet(params)
	if err != nil {
		return nil, err
	}
	for _, key := range sshKeys.Payload.Items {
		if key.Metadata.Name == name {
			return key, nil
		}
	}
	return nil, fmt.Errorf("project '%s' not found", name)
}

func (h *V1Client) GetSSHKeyByUID(uid, scope string) (*models.V1UserAssetSSH, error) {
	var params *userC.V1UsersAssetSSHGetUIDParams
	switch scope {
	case "project":
		params = userC.NewV1UsersAssetSSHGetUIDParamsWithContext(h.Ctx).WithUID(uid)
	case "tenant":
		params = userC.NewV1UsersAssetSSHGetUIDParams().WithUID(uid)
	default:
		return nil, errors.New("invalid scope")
	}

	sshKey, err := h.GetUserClient().V1UsersAssetSSHGetUID(params)
	if err != nil || sshKey == nil {
		return nil, err
	}

	return sshKey.Payload, nil
}

func (h *V1Client) UpdateSSHKey(uid string, body *models.V1UserAssetSSH, scope string) error {
	var params *userC.V1UsersAssetSSHUpdateParams
	switch scope {
	case "project":
		params = userC.NewV1UsersAssetSSHUpdateParamsWithContext(h.Ctx).WithUID(uid).WithBody(body)
	case "tenant":
		params = userC.NewV1UsersAssetSSHUpdateParams().WithUID(uid).WithBody(body)
	default:
		return errors.New("invalid scope")
	}

	_, err := h.GetUserClient().V1UsersAssetSSHUpdate(params)
	if err != nil {
		return err
	}

	return nil
}

func (h *V1Client) DeleteSSHKey(uid, scope string) error {
	var params *userC.V1UsersAssetSSHDeleteParams
	switch scope {
	case "project":
		params = userC.NewV1UsersAssetSSHDeleteParamsWithContext(h.Ctx).WithUID(uid)
	case "tenant":
		params = userC.NewV1UsersAssetSSHDeleteParams().WithUID(uid)
	default:
		return errors.New("invalid scope")
	}

	_, err := h.GetUserClient().V1UsersAssetSSHDelete(params)
	if err != nil {
		return err
	}

	return nil
}
