package client

import (
	"errors"
	clientv1 "github.com/spectrocloud/palette-sdk-go/api/client/v1"
	"github.com/spectrocloud/palette-sdk-go/api/models"
)

// CRUDL operations on teams are all tenant scoped.
// See: hubble/services/svccore/perms/user_acl.go

// CreateTeam creates a new team.
func (h *V1Client) CreateTeam(team *models.V1Team) (string, error) {
	// ACL scoped to tenant only
	params := clientv1.NewV1TeamsCreateParams().
		WithBody(team)
	resp, err := h.Client.V1TeamsCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

// UpdateTeam updates an existing team.
func (h *V1Client) UpdateTeam(uid string, team *models.V1Team) error {
	// ACL scoped to tenant only
	params := clientv1.NewV1TeamsUIDUpdateParams().
		WithUID(uid).
		WithBody(team)
	_, err := h.Client.V1TeamsUIDUpdate(params)
	return err
}

// GetTeam retrieves an existing team by UID.
func (h *V1Client) GetTeam(uid string) (*models.V1Team, error) {
	// ACL scoped to tenant only
	params := clientv1.NewV1TeamsUIDGetParams().
		WithUID(uid)
	resp, err := h.Client.V1TeamsUIDGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// GetTeamWithName retrieves an existing team by Team name.
func (h *V1Client) GetTeamWithName(teamName string) (*models.V1Team, error) {
	// ACL scoped to tenant only
	nameFilter := "metadata.name=" + teamName
	params := clientv1.NewV1TeamsListParams().WithFilters(&nameFilter)
	resp, err := h.Client.V1TeamsList(params)
	if err != nil {
		return nil, err
	}
	if resp.Payload.Items != nil {
		if len(resp.Payload.Items) == 1 {
			return resp.Payload.Items[0], nil
		} else {
			return nil, errors.New("More than one team found name: " + teamName)
		}
	} else {
		return nil, errors.New("Team not found for name: " + teamName)
	}
}

// DeleteTeam deletes an existing team by UID.
func (h *V1Client) DeleteTeam(uid string) error {
	// ACL scoped to tenant only
	params := clientv1.NewV1TeamsUIDDeleteParams().
		WithUID(uid)
	_, err := h.Client.V1TeamsUIDDelete(params)
	return err
}

// AssociateTeamProjectRole updates a team's project-role associations.
func (h *V1Client) AssociateTeamProjectRole(uid string, body *models.V1ProjectRolesPatch) error {
	params := clientv1.NewV1TeamsProjectRolesPutParamsWithContext(h.ctx).
		WithUID(uid).
		WithBody(body)
	_, err := h.Client.V1TeamsProjectRolesPut(params)
	return err
}

// GetTeamProjectRoleAssociation retrieves a team's project-role associations.
func (h *V1Client) GetTeamProjectRoleAssociation(uid string) (*models.V1ProjectRolesEntity, error) {
	params := clientv1.NewV1TeamsProjectRolesParamsWithContext(h.ctx).
		WithUID(uid)
	resp, err := h.Client.V1TeamsProjectRoles(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// AssociateTeamTenantRole updates a team's tenant-role associations.
func (h *V1Client) AssociateTeamTenantRole(uid string, body *models.V1TeamTenantRolesUpdate) error {
	params := clientv1.NewV1TeamsUIDTenantRolesUpdateParamsWithContext(h.ctx).
		WithUID(uid).
		WithBody(body)
	_, err := h.Client.V1TeamsUIDTenantRolesUpdate(params)
	return err
}

// GetTeamTenantRoleAssociation retrieves a team's tenant-role associations.
func (h *V1Client) GetTeamTenantRoleAssociation(uid string) (*models.V1TeamTenantRolesEntity, error) {
	params := clientv1.NewV1TeamsUIDTenantRolesGetParamsWithContext(h.ctx).
		WithUID(uid)
	resp, err := h.Client.V1TeamsUIDTenantRolesGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// AssociateTeamWorkspaceRole updates a team's workspace-role associations.
func (h *V1Client) AssociateTeamWorkspaceRole(uid string, body *models.V1WorkspacesRolesPatch) error {
	params := clientv1.NewV1TeamsWorkspaceRolesPutParamsWithContext(h.ctx).
		WithTeamUID(uid).
		WithBody(body)
	_, err := h.Client.V1TeamsWorkspaceRolesPut(params)
	return err
}

// GetTeamWorkspaceRoleAssociation retrieves a team's workspace-role associations.
func (h *V1Client) GetTeamWorkspaceRoleAssociation(uid string) (*models.V1WorkspaceScopeRoles, error) {
	params := clientv1.NewV1TeamsWorkspaceGetRolesParamsWithContext(h.ctx).
		WithTeamUID(uid)
	resp, err := h.Client.V1TeamsWorkspaceGetRoles(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// AssociateTeamResourceRole updates a team's resource-role associations.
func (h *V1Client) AssociateTeamResourceRole(uid string, body *models.V1ResourceRolesUpdateEntity) error {
	params := clientv1.NewV1TeamsUIDResourceRolesCreateParamsWithContext(h.ctx).
		WithUID(uid).
		WithBody(body)
	_, err := h.Client.V1TeamsUIDResourceRolesCreate(params)
	return err
}

// UpdateTeamResourceRole updates a team's resource-role associations.
func (h *V1Client) UpdateTeamResourceRole(uid string, body *models.V1ResourceRolesUpdateEntity) error {
	params := clientv1.NewV1TeamsResourceRolesUIDUpdateParamsWithContext(h.ctx).
		WithUID(uid).
		WithBody(body)
	_, err := h.Client.V1TeamsResourceRolesUIDUpdate(params)
	return err
}

// GetTeamResourceRole retrieves a resource-role by team UID and role name.
func (h *V1Client) GetTeamResourceRole(uid, name string) (*models.V1UIDSummary, error) {
	params := clientv1.NewV1TeamsUIDResourceRolesParamsWithContext(h.ctx).
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
