package client

import (
	"time"

	clientv1 "github.com/spectrocloud/palette-sdk-go/api/client/version1"
	"github.com/spectrocloud/palette-sdk-go/api/models"
)

// GetEvents retrieves events for a given object.
// TODO: which kinds are supported?
func (h *V1Client) GetEvents(kind, uid string, continueVar, fields, filters, orderBy *string, limit, offset *int64, timeout *time.Duration) ([]*models.V1Event, error) {
	params := clientv1.NewV1EventsComponentsObjTypeUIDListParamsWithContext(h.ctx).
		WithObjectKind(kind).
		WithObjectUID(uid)
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
	resp, err := h.Client.V1EventsComponentsObjTypeUIDList(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload.Items, nil
}

// GetNotifications retrieves notifications for a given object.
// TODO: which kinds are supported?
func (h *V1Client) GetNotifications(kind, uid string, continueVar, fields, filters, isDone, orderBy *string, limit, offset *int64, timeout *time.Duration) ([]*models.V1Notification, error) {
	params := clientv1.NewV1NotificationsObjTypeUIDListParamsWithContext(h.ctx).
		WithObjectKind(kind).
		WithObjectUID(uid)
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
	resp, err := h.Client.V1NotificationsObjTypeUIDList(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload.Items, nil
}
