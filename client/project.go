package client

import (
	"fmt"

	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
)

// CRUDL operations on projects are all tenant scoped.
// See: hubble/services/svccore/perms/user_acl.go

func (h *V1Client) CreateProject(body *models.V1ProjectEntity) (string, error) {
	// ACL scoped to tenant only
	params := clientV1.NewV1ProjectsCreateParams().
		WithBody(body)
	resp, err := h.Client.V1ProjectsCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

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

func (h *V1Client) GetProjectByUID(uid string) (*models.V1Project, error) {
	// ACL scoped to tenant only
	params := clientV1.NewV1ProjectsUIDGetParams().
		WithUID(uid)
	resp, err := h.Client.V1ProjectsUIDGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

func (h *V1Client) GetProjects() (*models.V1ProjectsMetadata, error) {
	// ACL scoped to tenant only
	params := clientV1.NewV1ProjectsMetadataParams()
	resp, err := h.Client.V1ProjectsMetadata(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

func (h *V1Client) UpdateProject(uid string, body *models.V1ProjectEntity) error {
	// ACL scoped to tenant only
	params := clientV1.NewV1ProjectsUIDUpdateParams().
		WithUID(uid).
		WithBody(body)
	_, err := h.Client.V1ProjectsUIDUpdate(params)
	return err
}

func (h *V1Client) DeleteProject(uid string) error {
	// ACL scoped to tenant only
	params := clientV1.NewV1ProjectsUIDDeleteParams().
		WithUID(uid)
	_, err := h.Client.V1ProjectsUIDDelete(params)
	return err
}
