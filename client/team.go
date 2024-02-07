package client

import (
	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
)

func (h *V1Client) CreateTeam(team *models.V1Team) (string, error) {
	params := clientV1.NewV1TeamsCreateParams().WithBody(team)
	success, err := h.Client.V1TeamsCreate(params)
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

func (h *V1Client) UpdateTeam(uid string, team *models.V1Team) error {
	params := clientV1.NewV1TeamsUIDUpdateParams().WithBody(team).WithUID(uid)
	_, err := h.Client.V1TeamsUIDUpdate(params)
	if err != nil {
		return err
	}

	return nil
}

func (h *V1Client) GetTeam(uid string) (*models.V1Team, error) {
	params := clientV1.NewV1TeamsUIDGetParams().WithUID(uid)
	success, err := h.Client.V1TeamsUIDGet(params)
	if err != nil {
		return nil, err
	}

	return success.Payload, nil
}

func (h *V1Client) DeleteTeam(uid string) error {
	params := clientV1.NewV1TeamsUIDDeleteParams().WithUID(uid)
	_, err := h.Client.V1TeamsUIDDelete(params)
	if err != nil {
		return err
	}

	return nil
}

func (h *V1Client) AssociateTeamProjectRole(uid string, body *models.V1ProjectRolesPatch) error {
	params := clientV1.NewV1TeamsProjectRolesPutParams().WithUID(uid).WithBody(body)
	_, err := h.Client.V1TeamsProjectRolesPut(params)
	if err != nil {
		return err
	}

	return nil
}

func (h *V1Client) GetTeamProjectRoleAssociation(uid string) (*models.V1ProjectRolesEntity, error) {
	params := clientV1.NewV1TeamsProjectRolesParams().WithUID(uid)
	success, err := h.Client.V1TeamsProjectRoles(params)
	if err != nil {
		return nil, err
	}

	return success.Payload, nil
}

func (h *V1Client) AssociateTeamTenantRole(uid string, body *models.V1TeamTenantRolesUpdate) error {
	params := clientV1.NewV1TeamsUIDTenantRolesUpdateParams().WithUID(uid).WithBody(body)
	_, err := h.Client.V1TeamsUIDTenantRolesUpdate(params)
	if err != nil {
		return err
	}

	return nil
}

func (h *V1Client) GetTeamTenantRoleAssociation(uid string) (*models.V1TeamTenantRolesEntity, error) {
	params := clientV1.NewV1TeamsUIDTenantRolesGetParams().WithUID(uid)
	success, err := h.Client.V1TeamsUIDTenantRolesGet(params)
	if err != nil {
		return nil, err
	}

	return success.Payload, nil
}

func (h *V1Client) AssociateTeamWorkspaceRole(uid string, body *models.V1WorkspacesRolesPatch) error {
	params := clientV1.NewV1TeamsWorkspaceRolesPutParams().WithTeamUID(uid).WithBody(body)
	_, err := h.Client.V1TeamsWorkspaceRolesPut(params)
	if err != nil {
		return err
	}

	return nil
}

func (h *V1Client) GetTeamWorkspaceRoleAssociation(uid string) (*models.V1WorkspaceScopeRoles, error) {
	params := clientV1.NewV1TeamsWorkspaceGetRolesParams().WithTeamUID(uid)
	success, err := h.Client.V1TeamsWorkspaceGetRoles(params)
	if err != nil {
		return nil, err
	}

	return success.Payload, nil
}

func (h *V1Client) AssociateTeamResourceRole(uid string, body *models.V1ResourceRolesUpdateEntity) error {
	params := clientV1.NewV1TeamsUIDResourceRolesCreateParams().WithUID(uid).WithBody(body)
	_, err := h.Client.V1TeamsUIDResourceRolesCreate(params)
	if err != nil {
		return err
	}

	return nil
}

func (h *V1Client) UpdateTeamResourceRole(uid string, body *models.V1ResourceRolesUpdateEntity) error {
	params := clientV1.NewV1TeamsResourceRolesUIDUpdateParams().WithUID(uid).WithBody(body)
	_, err := h.Client.V1TeamsResourceRolesUIDUpdate(params)
	if err != nil {
		return err
	}

	return nil
}

func (h *V1Client) GetTeamResourceRole(uid, name string) (*models.V1UIDSummary, error) {
	params := clientV1.NewV1TeamsUIDResourceRolesParams().WithUID(uid)
	success, err := h.Client.V1TeamsUIDResourceRoles(params)
	if err != nil {
		return nil, err
	}

	for _, roles := range success.Payload.ResourceRoles {
		for _, role := range roles.Roles {
			if role.Name == name {
				return role, nil
			}
		}
	}

	return nil, nil
}
