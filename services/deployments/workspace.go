package deployments

import (
	"errors"

	"github.com/spectrocloud/palette-sdk-go/api/models"
)

func (s *service) CreateWorkspace(workspace *models.V1WorkspaceEntity) (string, error) {
	if workspace == nil {
		return "", errors.New("workspace entity is required")
	}
	return s.client.CreateWorkspace(workspace)
}

func (s *service) GetWorkspace(uid string) (*models.V1Workspace, error) {
	if uid == "" {
		return nil, errUIDRequired
	}
	return s.client.GetWorkspace(uid)
}

func (s *service) GetWorkspaceByName(name string) (*models.V1DashboardWorkspace, error) {
	if name == "" {
		return nil, errors.New("workspace name is required")
	}
	return s.client.GetWorkspaceByName(name)
}

func (s *service) DeleteWorkspace(uid string) error {
	if uid == "" {
		return errUIDRequired
	}
	return s.client.DeleteWorkspace(uid)
}

func (s *service) UpdateWorkspaceMetadata(uid string, meta *models.V1ObjectMeta) error {
	if uid == "" {
		return errUIDRequired
	}
	return s.client.UpdateWorkspaceMetadata(uid, meta)
}

func (s *service) GetWorkspaceBackup(uid string) (*models.V1WorkspaceBackup, error) {
	if uid == "" {
		return nil, errUIDRequired
	}
	return s.client.GetWorkspaceBackup(uid)
}
