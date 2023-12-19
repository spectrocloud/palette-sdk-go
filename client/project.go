package client

import (
	"errors"
	"fmt"

	"github.com/spectrocloud/hapi/apiutil/transport"
	hashboardC "github.com/spectrocloud/hapi/hashboard/client/v1"
	"github.com/spectrocloud/hapi/models"
	userC "github.com/spectrocloud/hapi/user/client/v1"
)

func (h *V1Client) CreateProject(body *models.V1ProjectEntity) (string, error) {
	params := userC.NewV1ProjectsCreateParams().WithBody(body)
	success, err := h.GetUserClient().V1ProjectsCreate(params)
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

func (h *V1Client) GetProjectUID(projectName string) (string, error) {
	if h.GetProjectUIDFn != nil {
		return h.GetProjectUIDFn(projectName)
	}

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
	params := userC.NewV1ProjectsUIDGetParams().WithUID(uid)
	project, err := h.GetUserClient().V1ProjectsUIDGet(params)
	if err != nil || project == nil {
		return nil, err
	}

	return project.Payload, nil
}

func (h *V1Client) GetProjects() (*models.V1ProjectsMetadata, error) {
	params := hashboardC.NewV1ProjectsMetadataParams()

	projects, err := h.GetHashboardClient().V1ProjectsMetadata(params)
	if err != nil || projects == nil {
		// to support 2.6 projects list
		var e *transport.TransportError
		if errors.As(err, &e) && e.HttpCode == 404 {
			limit := int64(0)
			oldParams := userC.NewV1ProjectsListParams().WithLimit(&limit)
			oldProjects, err := h.GetUserClient().V1ProjectsList(oldParams)
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
	params := userC.NewV1ProjectsUIDUpdateParams().WithBody(body).WithUID(uid)
	_, err := h.GetUserClient().V1ProjectsUIDUpdate(params)
	if err != nil {
		return err
	}

	return nil
}

func (h *V1Client) DeleteProject(uid string) error {
	params := userC.NewV1ProjectsUIDDeleteParams().WithUID(uid)
	_, err := h.GetUserClient().V1ProjectsUIDDelete(params)
	if err != nil {
		return err
	}

	return nil
}
