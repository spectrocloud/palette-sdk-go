// Package platform provides operations for tenant management, organization,
// SSO configuration, and platform-level settings.
package platform

import (
	"errors"

	"github.com/spectrocloud/palette-sdk-go/api/models"
	"github.com/spectrocloud/palette-sdk-go/client"
)

// Service defines operations for platform-level resources including tenants,
// organizations, SSO providers, and platform settings.
type Service interface {
	// Tenant
	GetTenantUID() (string, error)

	// Organization
	GetOrganizationByName(name string) (*models.V1LoginResponse, error)
	ListOrganizations() ([]*models.V1Organization, error)

	// SSO
	GetOIDC(tenantUID string) (*models.V1TenantOidcClientSpec, error)
	UpdateOIDC(tenantUID string, body *models.V1TenantOidcClientSpec) error
	GetSAML(tenantUID string) (*models.V1TenantSamlSpec, error)
	UpdateSAML(tenantUID string, body *models.V1TenantSamlRequestSpec) error
	GetDomains(tenantUID string) (*models.V1TenantDomains, error)
	UpdateDomain(tenantUID string, body *models.V1TenantDomains) error
	GetProviders(tenantUID string) (*models.V1TenantSsoAuthProvidersEntity, error)
	UpdateProviders(tenantUID string, body *models.V1TenantSsoAuthProvidersEntity) error

	// Platform Settings
	GetSessionTimeout(tenantUID string) (*models.V1AuthTokenSettings, error)
	UpdateSessionTimeout(tenantUID string, body *models.V1AuthTokenSettings) error
	GetLoginBanner(tenantUID string) (*models.V1LoginBannerSettings, error)
	UpdateLoginBanner(tenantUID string, body *models.V1LoginBannerSettings) error
	GetFIPSPreference(tenantUID string) (*models.V1FipsSettings, error)
	UpdateFIPSPreference(tenantUID string, body *models.V1FipsSettings) error
	GetPasswordPolicy(tenantUID string) (*models.V1TenantPasswordPolicyEntity, error)
	UpdatePasswordPolicy(tenantUID string, body *models.V1TenantPasswordPolicyEntity) error

	// Developer Settings
	GetDeveloperSetting(tenantUID string) (*models.V1DeveloperCredit, error)
	UpdateDeveloperSetting(tenantUID string, body *models.V1DeveloperCredit) error

	// Projects
	CreateProject(body *models.V1ProjectEntity) (string, error)
	GetProject(uid string) (*models.V1Project, error)
	GetProjectUID(name string) (string, error)
	ListProjects() (*models.V1ProjectsMetadata, error)
	UpdateProject(uid string, body *models.V1ProjectEntity) error
	DeleteProject(uid string) error
}

type service struct {
	client *client.V1Client
}

// New creates a new platform Service.
func New(c *client.V1Client) Service {
	if c == nil {
		panic("platform: V1Client must not be nil")
	}
	return &service{client: c}
}

var errTenantUIDRequired = errors.New("tenant UID is required")
