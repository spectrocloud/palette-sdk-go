package diagnostics

import (
	"errors"

	"github.com/spectrocloud/palette-sdk-go/api/models"
)

func (s *service) CreateAlert(body *models.V1Channel, projectUID, component string) (string, error) {
	if body == nil {
		return "", errors.New("alert channel body is required")
	}
	return s.client.CreateAlert(body, projectUID, component)
}

func (s *service) GetAlert(projectUID, component, alertUID string) (*models.V1Channel, error) {
	if alertUID == "" {
		return nil, errors.New("alert UID is required")
	}
	return s.client.GetAlert(projectUID, component, alertUID)
}

func (s *service) UpdateAlert(body *models.V1Channel, projectUID, component, alertUID string) (string, error) {
	if alertUID == "" {
		return "", errors.New("alert UID is required")
	}
	return s.client.UpdateAlert(body, projectUID, component, alertUID)
}

func (s *service) DeleteAlert(projectUID, component, alertUID string) error {
	if alertUID == "" {
		return errors.New("alert UID is required")
	}
	return s.client.DeleteAlert(projectUID, component, alertUID)
}
