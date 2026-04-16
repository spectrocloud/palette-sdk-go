// Package access provides operations for user management, teams, roles,
// API keys, and SSH keys.
package access

import (
	"errors"
	"time"

	"github.com/spectrocloud/palette-sdk-go/api/models"
	"github.com/spectrocloud/palette-sdk-go/client"
)

// Service defines operations for access control resources.
type Service interface {
	// Auth
	Authenticate(body *models.V1AuthLogin) (*models.V1UserToken, error)
	AuthRefreshToken(token string) (*models.V1UserToken, error)

	// Users
	CreateUser(user *models.V1UserEntity) (string, error)
	GetUserByID(uid string) (*models.V1User, error)
	GetUserByName(name string) (*models.V1User, error)
	GetUserByEmail(email string) (*models.V1User, error)
	GetUsers() (*models.V1Users, error)
	UpdateUser(uid string, user *models.V1UserUpdateEntity) error
	DeleteUser(uid string) error

	// User role associations
	AssociateUserProjectRole(userUID string, body *models.V1ProjectRolesPatch) error
	AssociateUserTenantRole(userUID string, body *models.V1UserRoleUIDs) error
	AssociateUserWorkspaceRole(userUID string, body *models.V1WorkspacesRolesPatch) error

	// Teams
	CreateTeam(team *models.V1Team) (string, error)
	GetTeam(uid string) (*models.V1Team, error)
	GetTeamByName(name string) (*models.V1Team, error)
	UpdateTeam(uid string, team *models.V1Team) error
	DeleteTeam(uid string) error
	AssociateTeamProjectRole(uid string, body *models.V1ProjectRolesPatch) error
	AssociateTeamTenantRole(uid string, body *models.V1TeamTenantRolesUpdate) error
	AssociateTeamWorkspaceRole(uid string, body *models.V1WorkspacesRolesPatch) error

	// Roles
	CreateRole(role *models.V1Role) (string, error)
	GetRole(name string) (*models.V1Role, error)
	GetRoleByID(uid string) (*models.V1Role, error)
	UpdateRole(role *models.V1Role, uid string) error
	DeleteRole(uid string) error

	// API Keys
	CreateAPIKey(name string, annotations map[string]string, expiry time.Duration) (string, error)
	GetAPIKeys() (*models.V1APIKeys, error)
	DeleteAPIKey(uid string) error

	// SSH Keys
	CreateSSHKey(body *models.V1UserAssetSSH) (string, error)
	GetSSHKeys() ([]*models.V1UserAssetSSH, error)
	GetSSHKey(uid string) (*models.V1UserAssetSSH, error)
	UpdateSSHKey(uid string, body *models.V1UserAssetSSH) error
	DeleteSSHKey(uid string) error
}

type service struct {
	client *client.V1Client
}

// New creates a new access Service.
func New(c *client.V1Client) Service {
	if c == nil {
		panic("access: V1Client must not be nil")
	}
	return &service{client: c}
}

var errUIDRequired = errors.New("UID is required")
