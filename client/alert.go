package client

import (
	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
)

func (h *V1Client) CreateAlert(body *models.V1Channel, projectUID, component string) (string, error) {
	params := clientV1.NewV1ProjectsUIDAlertCreateParams().WithBody(body).WithUID(projectUID).WithComponent(component)
	success, err := h.GetClient().V1ProjectsUIDAlertCreate(params)
	if err != nil {
		return "", err
	}
	return *success.Payload.UID, nil
}

func (h *V1Client) UpdateAlert(body *models.V1Channel, projectUID, component, alertUID string) (string, error) {
	params := clientV1.NewV1ProjectsUIDAlertsUIDUpdateParams().WithBody(body).WithUID(projectUID).WithComponent(component).WithAlertUID(alertUID)
	_, err := h.GetClient().V1ProjectsUIDAlertsUIDUpdate(params)
	if err != nil {
		return "", err
	}
	return "success", nil

}

func (h *V1Client) ReadAlert(projectUID, component, alertUID string) (*models.V1Channel, error) {
	params := clientV1.NewV1ProjectsUIDAlertsUIDGetParams().WithUID(projectUID).WithComponent(component).WithAlertUID(alertUID)
	success, err := h.GetClient().V1ProjectsUIDAlertsUIDGet(params)
	if err != nil {
		return nil, err
	}
	return success.Payload, nil

}

func (h *V1Client) DeleteAlerts(projectUID, component, alertUID string) error {
	params := clientV1.NewV1ProjectsUIDAlertsUIDDeleteParams().WithUID(projectUID).WithComponent(component).WithAlertUID(alertUID)
	_, err := h.GetClient().V1ProjectsUIDAlertsUIDDelete(params)
	if err != nil {
		return err
	}
	return nil
}
