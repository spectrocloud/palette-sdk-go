package client

import (
	"github.com/spectrocloud/hapi/models"
	v1 "github.com/spectrocloud/hapi/user/client/v1"
)

func (h *V1Client) CreateAlert(body *models.V1Channel, projectUID, component string) (string, error) {
	if h.CreateAlertFn != nil {
		return h.CreateAlertFn(body, projectUID, component)
	}

	params := v1.NewV1ProjectsUIDAlertCreateParams().WithBody(body).WithUID(projectUID).WithComponent(component)
	success, err := h.GetUserClient().V1ProjectsUIDAlertCreate(params)
	if err != nil {
		return "", err
	}
	return *success.Payload.UID, nil
}

func (h *V1Client) UpdateAlert(body *models.V1Channel, projectUID, component, alertUID string) (string, error) {
	if h.UpdateAlertFn != nil {
		return h.UpdateAlertFn(body, projectUID, component, alertUID)
	}

	params := v1.NewV1ProjectsUIDAlertsUIDUpdateParams().WithBody(body).WithUID(projectUID).WithComponent(component).WithAlertUID(alertUID)
	_, err := h.GetUserClient().V1ProjectsUIDAlertsUIDUpdate(params)
	if err != nil {
		return "", err
	}
	return "success", nil

}

func (h *V1Client) ReadAlert(projectUID, component, alertUID string) (*models.V1Channel, error) {
	if h.ReadAlertFn != nil {
		return h.ReadAlertFn(projectUID, component, alertUID)
	}

	params := v1.NewV1ProjectsUIDAlertsUIDGetParams().WithUID(projectUID).WithComponent(component).WithAlertUID(alertUID)
	success, err := h.GetUserClient().V1ProjectsUIDAlertsUIDGet(params)
	if err != nil {
		return nil, err
	}
	return success.Payload, nil

}

func (h *V1Client) DeleteAlerts(projectUID, component, alertUID string) error {
	if h.DeleteAlertsFn != nil {
		return h.DeleteAlertsFn(projectUID, component, alertUID)
	}

	params := v1.NewV1ProjectsUIDAlertsUIDDeleteParams().WithUID(projectUID).WithComponent(component).WithAlertUID(alertUID)
	_, err := h.GetUserClient().V1ProjectsUIDAlertsUIDDelete(params)
	if err != nil {
		return err
	}

	return nil
}
