package client

import (
	"fmt"

	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
)

func (h *V1Client) CreateProject(body *models.V1ProjectEntity) (string, error) {
	params := clientV1.NewV1ProjectsCreateParamsWithContext(h.ctx).
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
	params := clientV1.NewV1ProjectsUIDGetParamsWithContext(h.ctx).
		WithUID(uid)
	resp, err := h.Client.V1ProjectsUIDGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

func (h *V1Client) GetProjects() (*models.V1ProjectsMetadata, error) {
	params := clientV1.NewV1ProjectsMetadataParamsWithContext(h.ctx)
	resp, err := h.Client.V1ProjectsMetadata(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

func (h *V1Client) UpdateProject(uid string, body *models.V1ProjectEntity) error {
	params := clientV1.NewV1ProjectsUIDUpdateParamsWithContext(h.ctx).
		WithUID(uid).
		WithBody(body)
	_, err := h.Client.V1ProjectsUIDUpdate(params)
	return err
}

func (h *V1Client) DeleteProject(uid string) error {
	params := clientV1.NewV1ProjectsUIDDeleteParamsWithContext(h.ctx).
		WithUID(uid)
	_, err := h.Client.V1ProjectsUIDDelete(params)
	return err
}
