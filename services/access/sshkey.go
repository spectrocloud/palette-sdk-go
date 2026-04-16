package access

import (
	"errors"

	"github.com/spectrocloud/palette-sdk-go/api/models"
)

func (s *service) CreateSSHKey(body *models.V1UserAssetSSH) (string, error) {
	if body == nil {
		return "", errors.New("SSH key body is required")
	}
	return s.client.CreateSSHKey(body)
}

func (s *service) GetSSHKeys() ([]*models.V1UserAssetSSH, error) {
	// V1Client doesn't have a list method; use GetSSHKeyByName or GetSSHKey individually.
	// For now, delegate to GetSSHKey pattern - consumers should use GetSSHKey(uid) directly.
	return nil, errors.New("list SSH keys not supported; use GetSSHKey(uid) to retrieve individual keys")
}

func (s *service) GetSSHKey(uid string) (*models.V1UserAssetSSH, error) {
	if uid == "" {
		return nil, errUIDRequired
	}
	return s.client.GetSSHKey(uid)
}

func (s *service) UpdateSSHKey(uid string, body *models.V1UserAssetSSH) error {
	if uid == "" {
		return errUIDRequired
	}
	return s.client.UpdateSSHKey(uid, body)
}

func (s *service) DeleteSSHKey(uid string) error {
	if uid == "" {
		return errUIDRequired
	}
	return s.client.DeleteSSHKey(uid)
}
