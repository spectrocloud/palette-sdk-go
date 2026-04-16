package clusters

import "github.com/spectrocloud/palette-sdk-go/api/models"

func (s *service) GetBackupConfig(uid string) (*models.V1ClusterBackup, error) {
	if uid == "" {
		return nil, errUIDRequired
	}
	return s.client.GetClusterBackupConfig(uid)
}

func (s *service) CreateBackupConfig(uid string, config *models.V1ClusterBackupConfig) error {
	if uid == "" {
		return errUIDRequired
	}
	return s.client.CreateClusterBackupConfig(uid, config)
}

func (s *service) UpdateBackupConfig(uid string, config *models.V1ClusterBackupConfig) error {
	if uid == "" {
		return errUIDRequired
	}
	return s.client.UpdateClusterBackupConfig(uid, config)
}
