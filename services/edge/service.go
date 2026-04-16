// Package edge provides operations for edge host and registration token management.
package edge

import (
	"errors"

	"github.com/spectrocloud/palette-sdk-go/api/models"
	"github.com/spectrocloud/palette-sdk-go/client"
)

// Service defines operations for edge host management.
type Service interface {
	// Registration tokens
	GetTokenValue(name string) (string, error)
	GetTokenByUID(uid string) (*models.V1EdgeToken, error)
	GetTokenByName(name string) (*models.V1EdgeToken, error)
	CreateToken(name string, body *models.V1EdgeTokenEntity) (string, string, error) // returns UID, token value
	DeleteToken(uid string) error

	// Edge hosts
	CreateEdgeHost(body *models.V1EdgeHostDeviceEntity) (string, error)
	ListEdgeHosts() ([]*models.V1EdgeHostsMetadata, error)
	GetEdgeHost(uid string) (*models.V1EdgeHostDevice, error)
	GetEdgeHostByName(name string) (*models.V1EdgeHostDevice, error)
	GetEdgeHostsByTags(tags map[string]string) ([]*models.V1EdgeHostsMetadata, error)
	DeleteEdgeHost(uid string) error
}

type service struct {
	client *client.V1Client
}

// New creates a new edge Service.
func New(c *client.V1Client) Service {
	if c == nil {
		panic("edge: V1Client must not be nil")
	}
	return &service{client: c}
}

var errUIDRequired = errors.New("UID is required")
var errNameRequired = errors.New("name is required")
