package client

import (
	clientv1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
)

// ListBackupStorageLocation returns a list of all backup storage locations.
func (h *V1Client) ListBackupStorageLocation() ([]*models.V1UserAssetsLocation, error) {
	params := clientv1.NewV1UsersAssetsLocationGetParamsWithContext(h.ctx)
	resp, err := h.Client.V1UsersAssetsLocationGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload.Items, nil
}

// GetBackupStorageLocation retrieves an existing backup storage location.
func (h *V1Client) GetBackupStorageLocation(uid string) (*models.V1UserAssetsLocation, error) {
	params := clientv1.NewV1UsersAssetsLocationGetParamsWithContext(h.ctx)
	resp, err := h.Client.V1UsersAssetsLocationGet(params)
	if err != nil {
		return nil, err
	}
	for _, account := range resp.Payload.Items {
		if account.Metadata.UID == uid {
			return account, nil
		}
	}
	return nil, nil
}

// GetS3BackupStorageLocation retrieves an existing S3 backup storage location.
func (h *V1Client) GetS3BackupStorageLocation(uid string) (*models.V1UserAssetsLocationS3, error) {
	params := clientv1.NewV1UsersAssetsLocationS3GetParamsWithContext(h.ctx).
		WithUID(uid)
	resp, err := h.Client.V1UsersAssetsLocationS3Get(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// CreateS3BackupStorageLocation creates a new S3 backup storage location.
func (h *V1Client) CreateS3BackupStorageLocation(bsl *models.V1UserAssetsLocationS3) (string, error) {
	params := clientv1.NewV1UsersAssetsLocationS3CreateParamsWithContext(h.ctx).
		WithBody(bsl)
	resp, err := h.Client.V1UsersAssetsLocationS3Create(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

// UpdateS3BackupStorageLocation updates an existing S3 backup storage location.
func (h *V1Client) UpdateS3BackupStorageLocation(uid string, bsl *models.V1UserAssetsLocationS3) error {
	params := clientv1.NewV1UsersAssetsLocationS3UpdateParamsWithContext(h.ctx).
		WithUID(uid).
		WithBody(bsl)
	_, err := h.Client.V1UsersAssetsLocationS3Update(params)
	return err
}

// DeleteS3BackupStorageLocation deletes an existing S3 backup storage location.
func (h *V1Client) DeleteS3BackupStorageLocation(uid string) error {
	params := clientv1.NewV1UsersAssetsLocationS3DeleteParamsWithContext(h.ctx).
		WithUID(uid)
	_, err := h.Client.V1UsersAssetsLocationS3Delete(params)
	return err
}
