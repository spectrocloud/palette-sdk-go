package clusters

import (
	"errors"

	"github.com/spectrocloud/palette-sdk-go/api/models"
)

func (s *service) CreateClusterGroup(group *models.V1ClusterGroupEntity) (string, error) {
	if group == nil {
		return "", errors.New("cluster group entity is required")
	}
	return s.client.CreateClusterGroup(group)
}

func (s *service) GetClusterGroup(uid string) (*models.V1ClusterGroup, error) {
	if uid == "" {
		return nil, errUIDRequired
	}
	return s.client.GetClusterGroup(uid)
}

func (s *service) DeleteClusterGroup(uid string) error {
	if uid == "" {
		return errUIDRequired
	}
	return s.client.DeleteClusterGroup(uid)
}

func (s *service) UpdateClusterGroup(uid string, group *models.V1ClusterGroupHostClusterEntity) error {
	if uid == "" {
		return errUIDRequired
	}
	return s.client.UpdateClusterGroup(uid, group)
}
