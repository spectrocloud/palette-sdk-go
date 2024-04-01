package client

import (
	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
)

func (h *V1Client) CreateTeam(team *models.V1Team) (string, error) {
	params := clientV1.NewV1TeamsCreateParamsWithContext(h.ctx).
		WithBody(team)
	resp, err := h.Client.V1TeamsCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

func (h *V1Client) UpdateTeam(uid string, team *models.V1Team) error {
	params := clientV1.NewV1TeamsUIDUpdateParamsWithContext(h.ctx).
		WithUID(uid).
		WithBody(team)
	_, err := h.Client.V1TeamsUIDUpdate(params)
	return err
}

func (h *V1Client) GetTeam(uid string) (*models.V1Team, error) {
	params := clientV1.NewV1TeamsUIDGetParamsWithContext(h.ctx).
		WithUID(uid)
	resp, err := h.Client.V1TeamsUIDGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

func (h *V1Client) DeleteTeam(uid string) error {
	params := clientV1.NewV1TeamsUIDDeleteParamsWithContext(h.ctx).
		WithUID(uid)
	_, err := h.Client.V1TeamsUIDDelete(params)
	return err
}

func (h *V1Client) AssociateTeamProjectRole(uid string, body *models.V1ProjectRolesPatch) error {
	params := clientV1.NewV1TeamsProjectRolesPutParamsWithContext(h.ctx).
		WithUID(uid).
		WithBody(body)
	_, err := h.Client.V1TeamsProjectRolesPut(params)
	return err
}

func (h *V1Client) GetTeamProjectRoleAssociation(uid string) (*models.V1ProjectRolesEntity, error) {
	params := clientV1.NewV1TeamsProjectRolesParamsWithContext(h.ctx).
		WithUID(uid)
	resp, err := h.Client.V1TeamsProjectRoles(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

func (h *V1Client) AssociateTeamTenantRole(uid string, body *models.V1TeamTenantRolesUpdate) error {
	params := clientV1.NewV1TeamsUIDTenantRolesUpdateParamsWithContext(h.ctx).
		WithUID(uid).
		WithBody(body)
	_, err := h.Client.V1TeamsUIDTenantRolesUpdate(params)
	return err
}

func (h *V1Client) GetTeamTenantRoleAssociation(uid string) (*models.V1TeamTenantRolesEntity, error) {
	params := clientV1.NewV1TeamsUIDTenantRolesGetParamsWithContext(h.ctx).
		WithUID(uid)
	resp, err := h.Client.V1TeamsUIDTenantRolesGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

func (h *V1Client) AssociateTeamWorkspaceRole(uid string, body *models.V1WorkspacesRolesPatch) error {
	params := clientV1.NewV1TeamsWorkspaceRolesPutParamsWithContext(h.ctx).
		WithTeamUID(uid).
		WithBody(body)
	_, err := h.Client.V1TeamsWorkspaceRolesPut(params)
	return err
}

func (h *V1Client) GetTeamWorkspaceRoleAssociation(uid string) (*models.V1WorkspaceScopeRoles, error) {
	params := clientV1.NewV1TeamsWorkspaceGetRolesParamsWithContext(h.ctx).
		WithTeamUID(uid)
	resp, err := h.Client.V1TeamsWorkspaceGetRoles(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

func (h *V1Client) AssociateTeamResourceRole(uid string, body *models.V1ResourceRolesUpdateEntity) error {
	params := clientV1.NewV1TeamsUIDResourceRolesCreateParamsWithContext(h.ctx).
		WithUID(uid).
		WithBody(body)
	_, err := h.Client.V1TeamsUIDResourceRolesCreate(params)
	return err
}

func (h *V1Client) UpdateTeamResourceRole(uid string, body *models.V1ResourceRolesUpdateEntity) error {
	params := clientV1.NewV1TeamsResourceRolesUIDUpdateParamsWithContext(h.ctx).
		WithUID(uid).
		WithBody(body)
	_, err := h.Client.V1TeamsResourceRolesUIDUpdate(params)
	return err
}

func (h *V1Client) GetTeamResourceRole(uid, name string) (*models.V1UIDSummary, error) {
	params := clientV1.NewV1TeamsUIDResourceRolesParamsWithContext(h.ctx).
		WithUID(uid)
	resp, err := h.Client.V1TeamsUIDResourceRoles(params)
	if err != nil {
		return nil, err
	}
	for _, roles := range resp.Payload.ResourceRoles {
		for _, role := range roles.Roles {
			if role.Name == name {
				return role, nil
			}
		}
	}
	return nil, nil
}
