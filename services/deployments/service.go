// Package deployments provides operations for workspace and application
// deployment management.
package deployments

import (
	"errors"

	"github.com/spectrocloud/palette-sdk-go/api/models"
	"github.com/spectrocloud/palette-sdk-go/client"
)

// Service defines operations for deployment management.
type Service interface {
	// Workspaces
	CreateWorkspace(workspace *models.V1WorkspaceEntity) (string, error)
	GetWorkspace(uid string) (*models.V1Workspace, error)
	GetWorkspaceByName(name string) (*models.V1DashboardWorkspace, error)
	DeleteWorkspace(uid string) error
	UpdateWorkspaceMetadata(uid string, meta *models.V1ObjectMeta) error
	GetWorkspaceBackup(uid string) (*models.V1WorkspaceBackup, error)

	// Applications
	GetApplication(uid string) (*models.V1AppDeployment, error)
	CreateApplicationWithNewSandboxCluster(body *models.V1AppDeploymentClusterGroupEntity) (string, error)
	CreateApplicationWithExistingSandboxCluster(body *models.V1AppDeploymentVirtualClusterEntity) (string, error)
	DeleteApplication(uid string) error
}

type service struct {
	client *client.V1Client
}

// New creates a new deployments Service.
func New(c *client.V1Client) Service {
	if c == nil {
		panic("deployments: V1Client must not be nil")
	}
	return &service{client: c}
}

var errUIDRequired = errors.New("UID is required")
