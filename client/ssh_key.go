package client

import (
	"fmt"

	clientv1 "github.com/spectrocloud/palette-sdk-go/api/client/version1"
	"github.com/spectrocloud/palette-sdk-go/api/models"
)

// CreateSSHKey creates a new SSH key.
func (h *V1Client) CreateSSHKey(body *models.V1UserAssetSSH) (string, error) {
	sshEntity := &models.V1UserAssetSSHEntity{
		Metadata: &models.V1ObjectMetaInputEntity{
			Annotations: nil,
			Labels:      nil,
			Name:        body.Metadata.Name,
		},
		Spec: &models.V1UserAssetSSHSpec{PublicKey: body.Spec.PublicKey},
	}

	params := clientv1.NewV1UserAssetsSSHCreateParamsWithContext(h.ctx).
		WithBody(sshEntity)
	resp, err := h.Client.V1UserAssetsSSHCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

// GetSSHKeyByName retrieves an existing SSH key by name.
func (h *V1Client) GetSSHKeyByName(name string) (*models.V1UserAssetSSH, error) {
	params := clientv1.NewV1UsersAssetsSSHGetParamsWithContext(h.ctx)
	resp, err := h.Client.V1UsersAssetsSSHGet(params)
	if err != nil {
		return nil, err
	}
	for _, key := range resp.Payload.Items {
		if key.Metadata.Name == name {
			return key, nil
		}
	}
	return nil, fmt.Errorf("ssh key '%s' not found", name)
}

// GetSSHKey retrieves an existing SSH key by UID.
func (h *V1Client) GetSSHKey(uid string) (*models.V1UserAssetSSH, error) {
	params := clientv1.NewV1UsersAssetSSHGetUIDParamsWithContext(h.ctx).
		WithUID(uid)
	resp, err := h.Client.V1UsersAssetSSHGetUID(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// UpdateSSHKey updates an existing SSH key.
func (h *V1Client) UpdateSSHKey(uid string, body *models.V1UserAssetSSH) error {
	params := clientv1.NewV1UsersAssetSSHUpdateParamsWithContext(h.ctx).
		WithUID(uid).
		WithBody(body)
	_, err := h.Client.V1UsersAssetSSHUpdate(params)
	return err
}

// DeleteSSHKey deletes an existing SSH key.
func (h *V1Client) DeleteSSHKey(uid string) error {
	params := clientv1.NewV1UsersAssetSSHDeleteParamsWithContext(h.ctx).
		WithUID(uid)
	_, err := h.Client.V1UsersAssetSSHDelete(params)
	return err
}
