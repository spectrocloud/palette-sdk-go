package client

import (
	clientv1 "github.com/spectrocloud/palette-sdk-go/api/client/version1"
	"github.com/spectrocloud/palette-sdk-go/api/models"
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

// GetMinioBackupStorageLocation retrieves an existing Minio backup storage location.
func (h *V1Client) GetMinioBackupStorageLocation(uid string) (*models.V1UserAssetsLocationS3, error) {
	params := clientv1.NewV1UsersAssetsLocationMinioGetParamsWithContext(h.ctx).
		WithUID(uid)
	resp, err := h.Client.V1UsersAssetsLocationMinioGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// GetGCPBackupStorageLocation retrieves an existing GCP backup storage location.
func (h *V1Client) GetGCPBackupStorageLocation(uid string) (*models.V1UserAssetsLocationGcp, error) {
	params := clientv1.NewV1UsersAssetsLocationGcpGetParamsWithContext(h.ctx).
		WithUID(uid)
	resp, err := h.Client.V1UsersAssetsLocationGcpGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// GetAzureBackupStorageLocation retrieves an existing Azure backup storage location.
func (h *V1Client) GetAzureBackupStorageLocation(uid string) (*models.V1UserAssetsLocationAzure, error) {
	params := clientv1.NewV1UsersAssetsLocationAzureGetParamsWithContext(h.ctx).
		WithUID(uid)
	resp, err := h.Client.V1UsersAssetsLocationAzureGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// ValidateS3BackupStorageLocation validate a S3 credential for backup storage location.
func (h *V1Client) ValidateS3BackupStorageLocation(awsCred *models.V1AwsS3BucketCredentials) error {
	params := clientv1.NewV1AwsS3ValidateParamsWithContext(h.ctx).WithAwsS3Credential(awsCred)
	_, err := h.Client.V1AwsS3Validate(params)
	return err
}

// ValidateGcpBackupStorageLocation validate a gcp credential for backup storage location.
func (h *V1Client) ValidateGcpBackupStorageLocation(gcpCred *models.V1GcpAccountNameValidateSpec) error {
	params := clientv1.NewV1GcpBucketNameValidateParamsWithContext(h.ctx).WithBody(gcpCred)
	_, err := h.Client.V1GcpBucketNameValidate(params)
	return err
}

// ValidateAzureBackupStorageLocation validate an azure credential for backup storage location.
func (h *V1Client) ValidateAzureBackupStorageLocation(azureCred *models.V1AzureCloudAccount) error {
	params := clientv1.NewV1AzureAccountValidateParamsWithContext(h.ctx).WithAzureCloudAccount(azureCred)
	_, err := h.Client.V1AzureAccountValidate(params)
	return err
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

// CreateMinioBackupStorageLocation creates a new Minio backup storage location.
func (h *V1Client) CreateMinioBackupStorageLocation(bsl *models.V1UserAssetsLocationS3) (string, error) {
	params := clientv1.NewV1UsersAssetsLocationMinioCreateParamsWithContext(h.ctx).
		WithBody(bsl)
	resp, err := h.Client.V1UsersAssetsLocationMinioCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

// CreateGcpBackupStorageLocation creates a new GCP backup storage location.
func (h *V1Client) CreateGcpBackupStorageLocation(bsl *models.V1UserAssetsLocationGcp) (string, error) {
	params := clientv1.NewV1UsersAssetsLocationGcpCreateParamsWithContext(h.ctx).
		WithBody(bsl)
	resp, err := h.Client.V1UsersAssetsLocationGcpCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

// CreateAzureBackupStorageLocation creates a new Azure backup storage location.
func (h *V1Client) CreateAzureBackupStorageLocation(bsl *models.V1UserAssetsLocationAzure) (string, error) {
	params := clientv1.NewV1UsersAssetsLocationAzureCreateParamsWithContext(h.ctx).
		WithBody(bsl)
	resp, err := h.Client.V1UsersAssetsLocationAzureCreate(params)
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

// UpdateMinioBackupStorageLocation updates an existing Minio backup storage location.
func (h *V1Client) UpdateMinioBackupStorageLocation(uid string, bsl *models.V1UserAssetsLocationS3) error {
	params := clientv1.NewV1UsersAssetsLocationMinioUpdateParamsWithContext(h.ctx).
		WithUID(uid).
		WithBody(bsl)
	_, err := h.Client.V1UsersAssetsLocationMinioUpdate(params)
	return err
}

// UpdateGcpBackupStorageLocation updates an existing Gcp backup storage location.
func (h *V1Client) UpdateGcpBackupStorageLocation(uid string, bsl *models.V1UserAssetsLocationGcp) error {
	params := clientv1.NewV1UsersAssetsLocationGcpUpdateParamsWithContext(h.ctx).
		WithUID(uid).
		WithBody(bsl)
	_, err := h.Client.V1UsersAssetsLocationGcpUpdate(params)
	return err
}

// UpdateAzureBackupStorageLocation updates an existing Azure backup storage location.
func (h *V1Client) UpdateAzureBackupStorageLocation(uid string, bsl *models.V1UserAssetsLocationAzure) error {
	params := clientv1.NewV1UsersAssetsLocationAzureUpdateParamsWithContext(h.ctx).
		WithUID(uid).
		WithBody(bsl)
	_, err := h.Client.V1UsersAssetsLocationAzureUpdate(params)
	return err
}

// DeleteS3BackupStorageLocation deletes an existing S3 backup storage location.
func (h *V1Client) DeleteS3BackupStorageLocation(uid string) error {
	params := clientv1.NewV1UsersAssetsLocationS3DeleteParamsWithContext(h.ctx).
		WithUID(uid)
	_, err := h.Client.V1UsersAssetsLocationS3Delete(params)
	return err
}
