package client

import (
	"errors"
	"fmt"

	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
)

func (h *V1Client) CreateProject(body *models.V1ProjectEntity) (string, error) {
	params := clientV1.NewV1ProjectsCreateParams().WithBody(body)
	success, err := h.Client.V1ProjectsCreate(params)
	if err != nil {
		return "", err
	}
	return *success.Payload.UID, nil
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
	params := clientV1.NewV1ProjectsUIDGetParams().WithUID(uid)
	project, err := h.Client.V1ProjectsUIDGet(params)
	if err != nil {
		return nil, err
	} else if project == nil {
		return nil, errors.New("project not found")
	}
	return project.Payload, nil
}

func (h *V1Client) GetProjects() (*models.V1ProjectsMetadata, error) {
	params := clientV1.NewV1ProjectsMetadataParams()
	projects, err := h.Client.V1ProjectsMetadata(params)
	if err != nil {
		return nil, err
	} else if projects == nil {
		return nil, errors.New("projects not found")
	}
	return projects.Payload, nil
}

func (h *V1Client) UpdateProject(uid string, body *models.V1ProjectEntity) error {
	params := clientV1.NewV1ProjectsUIDUpdateParams().WithBody(body).WithUID(uid)
	_, err := h.Client.V1ProjectsUIDUpdate(params)
	if err != nil {
		return err
	}
	return nil
}

func (h *V1Client) DeleteProject(uid string) error {
	params := clientV1.NewV1ProjectsUIDDeleteParams().WithUID(uid)
	_, err := h.Client.V1ProjectsUIDDelete(params)
	if err != nil {
		return err
	}
	return nil
}
