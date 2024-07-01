package client

import (
	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
)

func (h *V1Client) CreateAlert(body *models.V1Channel, projectUID, component string) (string, error) {
	params := clientV1.NewV1ProjectsUIDAlertCreateParamsWithContext(h.ctx).
		WithBody(body).
		WithUID(projectUID).
		WithComponent(component)
	resp, err := h.Client.V1ProjectsUIDAlertCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

func (h *V1Client) UpdateAlert(body *models.V1Channel, projectUID, component, alertUID string) (string, error) {
	params := clientV1.NewV1ProjectsUIDAlertsUIDUpdateParamsWithContext(h.ctx).
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

func (h *V1Client) ReadAlert(projectUID, component, alertUID string) (*models.V1Channel, error) {
	params := clientV1.NewV1ProjectsUIDAlertsUIDGetParamsWithContext(h.ctx).
		WithUID(projectUID).
		WithComponent(component).
		WithAlertUID(alertUID)
	resp, err := h.Client.V1ProjectsUIDAlertsUIDGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil

}

func (h *V1Client) DeleteAlerts(projectUID, component, alertUID string) error {
	params := clientV1.NewV1ProjectsUIDAlertsUIDDeleteParamsWithContext(h.ctx).
		WithUID(projectUID).
		WithComponent(component).
		WithAlertUID(alertUID)
	_, err := h.Client.V1ProjectsUIDAlertsUIDDelete(params)
	return err
}
