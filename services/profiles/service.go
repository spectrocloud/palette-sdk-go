// Package profiles provides operations for cluster profiles, application
// profiles, and pack management.
package profiles

import (
	"errors"

	"github.com/spectrocloud/palette-sdk-go/api/models"
	"github.com/spectrocloud/palette-sdk-go/client"
)

// Service defines operations for profile management.
type Service interface {
	// Cluster profiles
	GetSummary(uid string) (*models.V1ClusterProfileSummary, error)
	Get(uid string) (*models.V1ClusterProfile, error)
	GetUID(profileName, profileVersion string) (string, error)
	Create(profile *models.V1ClusterProfileEntity) (string, error)
	Update(profile *models.V1ClusterProfileUpdateEntity) error
	Delete(uid string) error
	Publish(uid string) error
	Import(content string) (string, error)
	Export(uid, format string) ([]byte, error)
	List() ([]*models.V1ClusterProfileMetadata, error)
	SearchProfiles(filter *models.V1ClusterProfilesFilterSpec) ([]*models.V1ClusterProfileSummary, error)

	// Profile variables
	GetVariables(uid string) ([]*models.V1Variable, error)
	UpdateVariables(variables *models.V1Variables, uid string) error

	// Packs
	GetPacks(filters []string, registryUID string) ([]*models.V1PackSummary, error)
	SearchPacks(filter *models.V1PackFilterSpec, sortBy []*models.V1PackSortSpec) ([]*models.V1PackMetadata, error)
	GetPack(uid string) (*models.V1PackTagEntity, error)

	// Application profiles
	GetApplicationProfile(uid string) (*models.V1AppProfile, error)
	CreateApplicationProfile(profile *models.V1AppProfileEntity) (string, error)
	DeleteApplicationProfile(uid string) error
}

type service struct {
	client *client.V1Client
}

// New creates a new profiles Service.
func New(c *client.V1Client) Service {
	if c == nil {
		panic("profiles: V1Client must not be nil")
	}
	return &service{client: c}
}

var errUIDRequired = errors.New("UID is required")
