package access

import (
	"errors"

	"github.com/spectrocloud/palette-sdk-go/api/models"
)

func (s *service) CreateTeam(team *models.V1Team) (string, error) {
	if team == nil {
		return "", errors.New("team is required")
	}
	return s.client.CreateTeam(team)
}

func (s *service) GetTeam(uid string) (*models.V1Team, error) {
	if uid == "" {
		return nil, errUIDRequired
	}
	return s.client.GetTeam(uid)
}

func (s *service) GetTeamByName(name string) (*models.V1Team, error) {
	if name == "" {
		return nil, errors.New("team name is required")
	}
	return s.client.GetTeamWithName(name)
}

func (s *service) UpdateTeam(uid string, team *models.V1Team) error {
	if uid == "" {
		return errUIDRequired
	}
	return s.client.UpdateTeam(uid, team)
}

func (s *service) DeleteTeam(uid string) error {
	if uid == "" {
		return errUIDRequired
	}
	return s.client.DeleteTeam(uid)
}

func (s *service) AssociateTeamProjectRole(uid string, body *models.V1ProjectRolesPatch) error {
	if uid == "" {
		return errUIDRequired
	}
	return s.client.AssociateTeamProjectRole(uid, body)
}

func (s *service) AssociateTeamTenantRole(uid string, body *models.V1TeamTenantRolesUpdate) error {
	if uid == "" {
		return errUIDRequired
	}
	return s.client.AssociateTeamTenantRole(uid, body)
}

func (s *service) AssociateTeamWorkspaceRole(uid string, body *models.V1WorkspacesRolesPatch) error {
	if uid == "" {
		return errUIDRequired
	}
	return s.client.AssociateTeamWorkspaceRole(uid, body)
}
