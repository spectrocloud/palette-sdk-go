package platform

import (
	"errors"

	"github.com/spectrocloud/palette-sdk-go/api/models"
)

func (s *service) CreateProject(body *models.V1ProjectEntity) (string, error) {
	if body == nil {
		return "", errors.New("project entity is required")
	}
	return s.client.CreateProject(body)
}

func (s *service) GetProject(uid string) (*models.V1Project, error) {
	if uid == "" {
		return nil, errors.New("project UID is required")
	}
	return s.client.GetProject(uid)
}

func (s *service) GetProjectUID(name string) (string, error) {
	if name == "" {
		return "", errors.New("project name is required")
	}
	return s.client.GetProjectUID(name)
}

func (s *service) ListProjects() (*models.V1ProjectsMetadata, error) {
	return s.client.GetProjects()
}

func (s *service) UpdateProject(uid string, body *models.V1ProjectEntity) error {
	if uid == "" {
		return errors.New("project UID is required")
	}
	return s.client.UpdateProject(uid, body)
}

func (s *service) DeleteProject(uid string) error {
	if uid == "" {
		return errors.New("project UID is required")
	}
	return s.client.DeleteProject(uid)
}
