// Package registries provides operations for managing pack, Helm, and OCI
// container registries.
package registries

import (
	"errors"

	"github.com/spectrocloud/palette-sdk-go/api/models"
	"github.com/spectrocloud/palette-sdk-go/client"
)

// Service defines operations for registry management.
type Service interface {
	// Pack registries
	SearchPackRegistries() ([]*models.V1RegistryMetadata, error)
	GetPackRegistryByName(name string) (*models.V1PackRegistry, error)
	ListPackRegistries() ([]*models.V1PackRegistrySummary, error)
	GetPackRegistrySyncStatus(uid string) (*models.V1RegistrySyncStatus, error)

	// Helm registries
	ListHelmRegistries(scope string) ([]*models.V1HelmRegistrySummary, error)
	GetHelmRegistry(uid string) (*models.V1HelmRegistry, error)
	GetHelmRegistryByName(name string) (*models.V1HelmRegistry, error)
	CreateHelmRegistry(registry *models.V1HelmRegistryEntity) (string, error)
	UpdateHelmRegistry(uid string, registry *models.V1HelmRegistry) error
	DeleteHelmRegistry(uid string) error
	GetHelmRegistrySyncStatus(uid string) (*models.V1RegistrySyncStatus, error)

	// OCI basic registries
	ListOCIRegistries() ([]*models.V1OciRegistry, error)
	GetOciBasicRegistry(uid string) (*models.V1BasicOciRegistry, error)
	CreateOciBasicRegistry(registry *models.V1BasicOciRegistry) (string, error)
	UpdateOciBasicRegistry(uid string, registry *models.V1BasicOciRegistry) error
	DeleteOciBasicRegistry(uid string) error
	ValidateOciBasicRegistry(body *models.V1BasicOciRegistrySpec) error

	// OCI ECR registries
	GetOciEcrRegistry(uid string) (*models.V1EcrRegistry, error)
	CreateOciEcrRegistry(registry *models.V1EcrRegistry) (string, error)
	UpdateOciEcrRegistry(uid string, registry *models.V1EcrRegistry) error
	DeleteOciEcrRegistry(uid string) error
	ValidateOciEcrRegistry(body *models.V1EcrRegistrySpec) error
}

type service struct {
	client *client.V1Client
}

// New creates a new registries Service.
func New(c *client.V1Client) Service {
	if c == nil {
		panic("registries: V1Client must not be nil")
	}
	return &service{client: c}
}

var errUIDRequired = errors.New("UID is required")
