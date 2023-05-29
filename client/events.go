package client

import (
	"time"

	event "github.com/spectrocloud/hapi/event/client/v1"
	"github.com/spectrocloud/hapi/models"
)

func (h *V1Client) GetEvents(kind, uid string, continueVar, fields, filters, orderBy *string, limit, offset *int64, timeout *time.Duration) ([]*models.V1Event, error) {
	client, err := h.GetEventClient()
	if err != nil {
		return nil, err
	}

	params := event.NewV1EventsComponentsObjTypeUIDListParamsWithContext(h.Ctx).WithObjectKind(kind).WithObjectUID(uid)
	if continueVar != nil {
		params = params.WithContinue(continueVar)
	}
	if fields != nil {
		params = params.WithFields(fields)
	}
	if filters != nil {
		params = params.WithFilters(filters)
	}
	if limit != nil {
		params = params.WithLimit(limit)
	}
	if offset != nil {
		params = params.WithOffset(offset)
	}
	if orderBy != nil {
		params = params.WithOrderBy(orderBy)
	}
	if timeout != nil {
		params = params.WithTimeout(*timeout)
	}

	resp, err := client.V1EventsComponentsObjTypeUIDList(params)
	if err != nil {
		return nil, err
	}

	return resp.Payload.Items, nil
}

func (h *V1Client) GetNotifications(kind, uid string, continueVar, fields, filters, isDone, orderBy *string, limit, offset *int64, timeout *time.Duration) ([]*models.V1Notification, error) {
	client, err := h.GetEventClient()
	if err != nil {
		return nil, err
	}

	params := event.NewV1NotificationsObjTypeUIDListParamsWithContext(h.Ctx).WithObjectKind(kind).WithObjectUID(uid)
	if continueVar != nil {
		params = params.WithContinue(continueVar)
	}
	if fields != nil {
		params = params.WithFields(fields)
	}
	if filters != nil {
		params = params.WithFilters(filters)
	}
	if isDone != nil {
		params = params.WithIsDone(isDone)
	}
	if limit != nil {
		params = params.WithLimit(limit)
	}
	if offset != nil {
		params = params.WithOffset(offset)
	}
	if orderBy != nil {
		params = params.WithOrderBy(orderBy)
	}
	if timeout != nil {
		params = params.WithTimeout(*timeout)
	}

	resp, err := client.V1NotificationsObjTypeUIDList(params)
	if err != nil {
		return nil, err
	}

	return resp.Payload.Items, nil
}
