package diagnostics

import "github.com/spectrocloud/palette-sdk-go/api/models"

func (s *service) GetClusterScanConfig(uid string) (*models.V1ClusterComplianceScan, error) {
	if uid == "" {
		return nil, errUIDRequired
	}
	return s.client.GetClusterScanConfig(uid)
}

func (s *service) CreateClusterScanConfig(uid string, config *models.V1ClusterComplianceScheduleConfig) error {
	if uid == "" {
		return errUIDRequired
	}
	return s.client.CreateClusterScanConfig(uid, config)
}

func (s *service) UpdateClusterScanConfig(uid string, config *models.V1ClusterComplianceScheduleConfig) error {
	if uid == "" {
		return errUIDRequired
	}
	return s.client.UpdateClusterScanConfig(uid, config)
}
