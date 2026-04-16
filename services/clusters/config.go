package clusters

import "github.com/spectrocloud/palette-sdk-go/api/models"

func (s *service) GetNamespaceConfig(uid string) (*models.V1ClusterNamespaceResources, error) {
	if uid == "" {
		return nil, errUIDRequired
	}
	return s.client.GetClusterNamespaceConfig(uid)
}

func (s *service) GetRbacConfig(uid string) (*models.V1ClusterRbacs, error) {
	if uid == "" {
		return nil, errUIDRequired
	}
	return s.client.GetClusterRbacConfig(uid)
}

func (s *service) UpdateOsPatchConfig(uid string, config *models.V1OsPatchEntity) error {
	if uid == "" {
		return errUIDRequired
	}
	return s.client.UpdateClusterOsPatchConfig(uid, config)
}
