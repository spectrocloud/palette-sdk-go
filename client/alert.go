package client

import (
	clientv1 "github.com/spectrocloud/palette-sdk-go/api/client/v1"
	"github.com/spectrocloud/palette-sdk-go/api/models"
)

// CreateAlert creates a new alert.
func (h *V1Client) CreateAlert(body *models.V1Channel, projectUID, component string) (string, error) {
	params := clientv1.NewV1ProjectsUIDAlertCreateParamsWithContext(h.ctx).
		WithBody(body).
		WithUID(projectUID).
		WithComponent(component)
	resp, err := h.Client.V1ProjectsUIDAlertCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

// UpdateAlert updates an existing alert.
func (h *V1Client) UpdateAlert(body *models.V1Channel, projectUID, component, alertUID string) (string, error) {
	params := clientv1.NewV1ProjectsUIDAlertsUIDUpdateParamsWithContext(h.ctx).
		WithBody(body).
		WithUID(projectUID).
		WithComponent(component).
		WithAlertUID(alertUID)
	_, err := h.Client.V1ProjectsUIDAlertsUIDUpdate(params)
	if err != nil {
		return "", err
	}
	return "success", nil
}

// UpdateProjectAlerts update alerts for project.
func (h *V1Client) UpdateProjectAlerts(body *models.V1AlertEntity, projectUID, component string) error {
	params := clientv1.NewV1ProjectsUIDAlertUpdateParamsWithContext(h.ctx).WithBody(body).WithUID(projectUID).
		WithComponent(component)
	_, err := h.Client.V1ProjectsUIDAlertUpdate(params)
	return err
}

// GetAlert retrieves an existing alert.
func (h *V1Client) GetAlert(projectUID, component, alertUID string) (*models.V1Channel, error) {
	params := clientv1.NewV1ProjectsUIDAlertsUIDGetParamsWithContext(h.ctx).
		WithUID(projectUID).
		WithComponent(component).
		WithAlertUID(alertUID)
	resp, err := h.Client.V1ProjectsUIDAlertsUIDGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// DeleteAlert deletes an existing alert.
func (h *V1Client) DeleteAlert(projectUID, component, alertUID string) error {
	params := clientv1.NewV1ProjectsUIDAlertsUIDDeleteParamsWithContext(h.ctx).
		WithUID(projectUID).
		WithComponent(component).
		WithAlertUID(alertUID)
	_, err := h.Client.V1ProjectsUIDAlertsUIDDelete(params)
	return err
}
