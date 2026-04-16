package edge

import (
	"github.com/spectrocloud/palette-sdk-go/api/models"
)

func (s *service) CreateEdgeHost(body *models.V1EdgeHostDeviceEntity) (string, error) {
	if body == nil || body.Metadata == nil || body.Metadata.Name == "" {
		return "", errNameRequired
	}
	return s.client.CreateAppliance(body)
}

func (s *service) ListEdgeHosts() ([]*models.V1EdgeHostsMetadata, error) {
	return s.client.ListEdgeHosts()
}

func (s *service) GetEdgeHost(uid string) (*models.V1EdgeHostDevice, error) {
	if uid == "" {
		return nil, errUIDRequired
	}
	return s.client.GetEdgeHost(uid)
}

func (s *service) GetEdgeHostByName(name string) (*models.V1EdgeHostDevice, error) {
	if name == "" {
		return nil, errNameRequired
	}
	return s.client.GetEdgeHostByName(name)
}

func (s *service) GetEdgeHostsByTags(tags map[string]string) ([]*models.V1EdgeHostsMetadata, error) {
	return s.client.GetEdgeHostsByTags(tags)
}

func (s *service) DeleteEdgeHost(uid string) error {
	if uid == "" {
		return errUIDRequired
	}
	return s.client.DeleteEdgeHostDevice(uid)
}
