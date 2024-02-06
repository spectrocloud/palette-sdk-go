package client

import (
	"errors"
	"fmt"

	"github.com/spectrocloud/palette-api-go/apiutil/transport"
	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
)

func (h *V1Client) CreateProject(body *models.V1ProjectEntity) (string, error) {
	params := clientV1.NewV1ProjectsCreateParams().WithBody(body)
	success, err := h.GetClient().V1ProjectsCreate(params)
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
	project, err := h.GetClient().V1ProjectsUIDGet(params)
	if err != nil || project == nil {
		return nil, err
	}

	return project.Payload, nil
}

func (h *V1Client) GetProjects() (*models.V1ProjectsMetadata, error) {
	params := clientV1.NewV1ProjectsMetadataParams()

	projects, err := h.GetClient().V1ProjectsMetadata(params)
	if err != nil || projects == nil {
		// to support 2.6 projects list
		var e *transport.TransportError
		if errors.As(err, &e) && e.HttpCode == 404 {
			limit := int64(0)
			oldParams := clientV1.NewV1ProjectsListParams().WithLimit(&limit)
			oldProjects, err := h.GetClient().V1ProjectsList(oldParams)
			if err != nil || oldProjects == nil {
				return nil, err
			}
			ret := make([]*models.V1ProjectMetadata, 0)
			for _, pr := range oldProjects.Payload.Items {
				ret = append(ret, &models.V1ProjectMetadata{
					Metadata: &models.V1ObjectEntity{
						UID:  pr.Metadata.UID,
						Name: pr.Metadata.Name,
					},
				})
			}
			return &models.V1ProjectsMetadata{
				Items: ret,
			}, nil
		}
		return nil, err
	}

	return projects.Payload, nil
}

func (h *V1Client) UpdateProject(uid string, body *models.V1ProjectEntity) error {
	params := clientV1.NewV1ProjectsUIDUpdateParams().WithBody(body).WithUID(uid)
	_, err := h.GetClient().V1ProjectsUIDUpdate(params)
	if err != nil {
		return err
	}

	return nil
}

func (h *V1Client) DeleteProject(uid string) error {
	params := clientV1.NewV1ProjectsUIDDeleteParams().WithUID(uid)
	_, err := h.GetClient().V1ProjectsUIDDelete(params)
	if err != nil {
		return err
	}

	return nil
}
