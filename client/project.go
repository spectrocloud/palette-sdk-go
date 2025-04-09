package client

import (
	"fmt"

	clientv1 "github.com/spectrocloud/palette-sdk-go/api/client/version1"
	"github.com/spectrocloud/palette-sdk-go/api/models"
)

// CRUDL operations on projects are all tenant scoped.
// See: hubble/services/svccore/perms/user_acl.go

// CreateProject creates a new project.
func (h *V1Client) CreateProject(body *models.V1ProjectEntity) (string, error) {
	// ACL scoped to tenant only
	params := clientv1.NewV1ProjectsCreateParams().
		WithBody(body)
	resp, err := h.Client.V1ProjectsCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

// GetProjectUID retrieves an existing project's UID by name.
func (h *V1Client) GetProjectUID(projectName string) (string, error) {
	projects, err := h.GetProjects()
	if err != nil {
		return "", err
	}

	for _, project := range projects.Items {
		if project.Metadata.Name == projectName {
			return project.Metadata.UID, nil
		}
	}

	return "", fmt.Errorf("project '%s' not found", projectName)
}

// GetProject retrieves an existing project by UID.
func (h *V1Client) GetProject(uid string) (*models.V1Project, error) {
	// ACL scoped to tenant only
	params := clientv1.NewV1ProjectsUIDGetParams().
		WithUID(uid)
	resp, err := h.Client.V1ProjectsUIDGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// GetProjects retrieves all projects' metadata.
func (h *V1Client) GetProjects() (*models.V1ProjectsMetadata, error) {
	// ACL scoped to tenant only
	params := clientv1.NewV1ProjectsMetadataParams()
	resp, err := h.Client.V1ProjectsMetadata(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// UpdateProject updates an existing project.
func (h *V1Client) UpdateProject(uid string, body *models.V1ProjectEntity) error {
	// ACL scoped to tenant only
	params := clientv1.NewV1ProjectsUIDUpdateParams().
		WithUID(uid).
		WithBody(body)
	_, err := h.Client.V1ProjectsUIDUpdate(params)
	return err
}

// DeleteProject deletes an existing project.
func (h *V1Client) DeleteProject(uid string) error {
	// ACL scoped to tenant only
	params := clientv1.NewV1ProjectsUIDDeleteParams().
		WithUID(uid)
	_, err := h.Client.V1ProjectsUIDDelete(params)
	return err
}

// UpdateMetaProject updates an existing project's metadata.
func (h *V1Client) UpdateMetaProject(uid string, body *models.V1ObjectMeta) error {
	// ACL scoped to tenant only
	params := clientv1.NewV1ProjectsUIDMetaUpdateParams().
		WithUID(uid).
		WithBody(body)
	_, err := h.Client.V1ProjectsUIDMetaUpdate(params)
	return err
}
