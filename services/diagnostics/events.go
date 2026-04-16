package diagnostics

import (
	"errors"
	"time"

	"github.com/spectrocloud/palette-sdk-go/api/models"
)

func (s *service) GetEvents(kind, uid string, continueVar, fields, filters, orderBy *string, limit, offset *int64, timeout *time.Duration) ([]*models.V1Event, error) {
	if kind == "" {
		return nil, errors.New("event kind is required")
	}
	if uid == "" {
		return nil, errUIDRequired
	}
	return s.client.GetEvents(kind, uid, continueVar, fields, filters, orderBy, limit, offset, timeout)
}
