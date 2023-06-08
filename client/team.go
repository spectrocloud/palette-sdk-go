package client

import (
	"github.com/spectrocloud/hapi/models"
	clusterC "github.com/spectrocloud/hapi/spectrocluster/client/v1"
	userC "github.com/spectrocloud/hapi/user/client/v1"
)

func (h *V1Client) CreateTeam(team *models.V1Team) (string, error) {
	client, err := h.GetUserClient()
	if err != nil {
		return "", err
	}

	params := userC.NewV1TeamsCreateParams().WithBody(team)
	success, err := client.V1TeamsCreate(params)
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

func (h *V1Client) UpdateTeam(uid string, team *models.V1Team) error {
	client, err := h.GetUserClient()
	if err != nil {
		return err
	}

	params := userC.NewV1TeamsUIDUpdateParams().WithBody(team).WithUID(uid)
	_, err = client.V1TeamsUIDUpdate(params)
	if err != nil {
		return err
	}

	return nil
}

func (h *V1Client) GetTeam(uid string) (*models.V1Team, error) {
	client, err := h.GetUserClient()
	if err != nil {
		return nil, err
	}

	params := userC.NewV1TeamsUIDGetParams().WithUID(uid)
	success, err := client.V1TeamsUIDGet(params)
	if err != nil {
		return nil, err
	}

	return success.Payload, nil
}

func (h *V1Client) DeleteTeam(uid string) error {
	client, err := h.GetUserClient()
	if err != nil {
		return err
	}

	params := userC.NewV1TeamsUIDDeleteParams().WithUID(uid)
	_, err = client.V1TeamsUIDDelete(params)
	if err != nil {
		return err
	}

	return nil
}

func (h *V1Client) AssociateTeamProjectRole(uid string, body *models.V1ProjectRolesPatch) error {
	client, err := h.GetUserClient()
	if err != nil {
		return err
	}

	params := userC.NewV1TeamsProjectRolesPutParams().WithUID(uid).WithBody(body)
	_, err = client.V1TeamsProjectRolesPut(params)
	if err != nil {
		return err
	}

	return nil
}

func (h *V1Client) GetTeamProjectRoleAssociation(uid string) (*models.V1ProjectRolesEntity, error) {
	client, err := h.GetUserClient()
	if err != nil {
		return nil, err
	}

	params := userC.NewV1TeamsProjectRolesParams().WithUID(uid)
	success, err := client.V1TeamsProjectRoles(params)
	if err != nil {
		return nil, err
	}

	return success.Payload, nil
}

func (h *V1Client) AssociateTeamTenantRole(uid string, body *models.V1TeamTenantRolesUpdate) error {
	client, err := h.GetUserClient()
	if err != nil {
		return err
	}

	params := userC.NewV1TeamsUIDTenantRolesUpdateParams().WithUID(uid).WithBody(body)
	_, err = client.V1TeamsUIDTenantRolesUpdate(params)
	if err != nil {
		return err
	}

	return nil
}

func (h *V1Client) GetTeamTenantRoleAssociation(uid string) (*models.V1TeamTenantRolesEntity, error) {
	client, err := h.GetUserClient()
	if err != nil {
		return nil, err
	}

	params := userC.NewV1TeamsUIDTenantRolesGetParams().WithUID(uid)
	success, err := client.V1TeamsUIDTenantRolesGet(params)
	if err != nil {
		return nil, err
	}

	return success.Payload, nil
}

func (h *V1Client) AssociateTeamWorkspaceRole(uid string, body *models.V1WorkspacesRolesPatch) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return err
	}

	params := clusterC.NewV1TeamsWorkspaceRolesPutParams().WithTeamUID(uid).WithBody(body)
	_, err = client.V1TeamsWorkspaceRolesPut(params)
	if err != nil {
		return err
	}

	return nil
}

func (h *V1Client) GetTeamWorkspaceRoleAssociation(uid string) (*models.V1WorkspaceScopeRoles, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}

	params := clusterC.NewV1TeamsWorkspaceGetRolesParams().WithTeamUID(uid)
	success, err := client.V1TeamsWorkspaceGetRoles(params)
	if err != nil {
		return nil, err
	}

	return success.Payload, nil
}

func (h *V1Client) AssociateTeamResourceRole(uid string, body *models.V1ResourceRolesUpdateEntity) error {
	client, err := h.GetUserClient()
	if err != nil {
		return err
	}

	params := userC.NewV1TeamsUIDResourceRolesCreateParams().WithUID(uid).WithBody(body)
	_, err = client.V1TeamsUIDResourceRolesCreate(params)
	if err != nil {
		return err
	}

	return nil
}

func (h *V1Client) UpdateTeamResourceRole(uid string, body *models.V1ResourceRolesUpdateEntity) error {
	client, err := h.GetUserClient()
	if err != nil {
		return err
	}

	params := userC.NewV1TeamsResourceRolesUIDUpdateParams().WithUID(uid).WithBody(body)
	_, err = client.V1TeamsResourceRolesUIDUpdate(params)
	if err != nil {
		return err
	}

	return nil
}

func (h *V1Client) GetTeamResourceRole(uid string, name string) (*models.V1UIDSummary, error) {
	client, err := h.GetUserClient()
	if err != nil {
		return nil, err
	}

	params := userC.NewV1TeamsUIDResourceRolesParams().WithUID(uid)
	success, err := client.V1TeamsUIDResourceRoles(params)
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