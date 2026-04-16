// Package diagnostics provides operations for cluster scanning, events,
// alerts, and tag filters.
package diagnostics

import (
	"errors"
	"time"

	"github.com/spectrocloud/palette-sdk-go/api/models"
	"github.com/spectrocloud/palette-sdk-go/client"
)

// Service defines operations for cluster diagnostics and observability.
type Service interface {
	// Scan config
	GetClusterScanConfig(uid string) (*models.V1ClusterComplianceScan, error)
	CreateClusterScanConfig(uid string, config *models.V1ClusterComplianceScheduleConfig) error
	UpdateClusterScanConfig(uid string, config *models.V1ClusterComplianceScheduleConfig) error

	// Events
	GetEvents(kind, uid string, continueVar, fields, filters, orderBy *string, limit, offset *int64, timeout *time.Duration) ([]*models.V1Event, error)

	// Alerts
	CreateAlert(body *models.V1Channel, projectUID, component string) (string, error)
	GetAlert(projectUID, component, alertUID string) (*models.V1Channel, error)
	UpdateAlert(body *models.V1Channel, projectUID, component, alertUID string) (string, error)
	DeleteAlert(projectUID, component, alertUID string) error

	// Tag Filters
	CreateTagFilter(body *models.V1TagFilter) (*models.V1UID, error)
	GetTagFilter(uid string) (*models.V1TagFilterSummary, error)
	GetTagFilterByName(name string) (*models.V1FilterSummary, error)
	ListTagFilters() (*models.V1FiltersSummary, error)
	UpdateTagFilter(uid string, body *models.V1TagFilter) error
	DeleteTagFilter(uid string) error
}

type service struct {
	client *client.V1Client
}

// New creates a new diagnostics Service.
func New(c *client.V1Client) Service {
	if c == nil {
		panic("diagnostics: V1Client must not be nil")
	}
	return &service{client: c}
}

var errUIDRequired = errors.New("UID is required")
