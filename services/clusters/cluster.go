package clusters

import (
	"errors"

	"github.com/spectrocloud/palette-sdk-go/api/models"
)

func (s *service) Get(uid string) (*models.V1SpectroCluster, error) {
	if uid == "" {
		return nil, errUIDRequired
	}
	return s.client.GetCluster(uid)
}

func (s *service) GetByName(name string, virtual bool) (*models.V1SpectroCluster, error) {
	if name == "" {
		return nil, errors.New("cluster name is required")
	}
	return s.client.GetClusterByName(name, virtual)
}

func (s *service) GetSummary(uid string) (*models.V1SpectroClusterUIDSummary, error) {
	if uid == "" {
		return nil, errUIDRequired
	}
	return s.client.GetClusterSummary(uid)
}

func (s *service) Delete(uid string) error {
	if uid == "" {
		return errUIDRequired
	}
	return s.client.DeleteCluster(uid)
}

func (s *service) ForceDelete(uid string, forceDelete bool) error {
	if uid == "" {
		return errUIDRequired
	}
	return s.client.ForceDeleteCluster(uid, forceDelete)
}

func (s *service) Search(filter *models.V1SearchFilterSpec, sort []*models.V1SearchFilterSortSpec) ([]*models.V1SpectroClusterSummary, error) {
	return s.client.SearchClusterSummaries(filter, sort)
}

func (s *service) List() ([]*models.V1SpectroClusterSummary, error) {
	return s.client.SearchClusterSummaries(nil, nil)
}

func (s *service) UpdateMetadata(uid string, config *models.V1ObjectMetaInputEntitySchema) error {
	if uid == "" {
		return errUIDRequired
	}
	return s.client.UpdateClusterMetadata(uid, config)
}
