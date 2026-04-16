// Package cloudaccounts provides operations for managing cloud account
// credentials across all supported cloud providers.
package cloudaccounts

import (
	"errors"

	"github.com/spectrocloud/palette-sdk-go/api/models"
	"github.com/spectrocloud/palette-sdk-go/client"
)

// Service defines operations for cloud account management.
type Service interface {
	// Common
	List() ([]*models.V1CloudAccountSummary, error)
	Get(uid string) (*models.V1CloudAccountSummary, error)

	// AWS
	GetAwsByName(name, scope string) (*models.V1AwsAccount, error)
	CreateAws(account *models.V1AwsAccount) (string, error)
	UpdateAws(account *models.V1AwsAccount) error
	DeleteAws(uid string) error
	GetAws(uid string) (*models.V1AwsAccount, error)

	// Azure
	GetAzureByName(name, scope string) (*models.V1AzureAccount, error)
	CreateAzure(account *models.V1AzureAccount) (string, error)
	UpdateAzure(account *models.V1AzureAccount) error
	DeleteAzure(uid string) error
	GetAzure(uid string) (*models.V1AzureAccount, error)

	// GCP
	GetGcpByName(name, scope string) (*models.V1GcpAccount, error)
	CreateGcp(account *models.V1GcpAccountEntity) (string, error)
	UpdateGcp(account *models.V1GcpAccountEntity) error
	DeleteGcp(uid string) error
	GetGcp(uid string) (*models.V1GcpAccount, error)

	// vSphere
	GetVsphereByName(name, scope string) (*models.V1VsphereAccount, error)
	CreateVsphere(account *models.V1VsphereAccount) (string, error)
	UpdateVsphere(account *models.V1VsphereAccount) error
	DeleteVsphere(uid string) error
	GetVsphere(uid string) (*models.V1VsphereAccount, error)

	// MAAS
	GetMaasByName(name, scope string) (*models.V1MaasAccount, error)
	CreateMaas(account *models.V1MaasAccount) (string, error)
	UpdateMaas(account *models.V1MaasAccount) error
	DeleteMaas(uid string) error
	GetMaas(uid string) (*models.V1MaasAccount, error)

	// CloudStack
	GetCloudStackByName(name, scope string) (*models.V1CloudStackAccount, error)
	CreateCloudStack(account *models.V1CloudStackAccount) (string, error)
	UpdateCloudStack(account *models.V1CloudStackAccount) error
	DeleteCloudStack(uid string) error
	GetCloudStack(uid string) (*models.V1CloudStackAccount, error)

	// Custom Cloud
	GetCustomByName(cloudType, name, scope string) (*models.V1CustomAccount, error)
	CreateCustom(account *models.V1CustomAccountEntity, cloudType string) (string, error)
	UpdateCustom(uid string, account *models.V1CustomAccountEntity, cloudType string) error
	DeleteCustom(uid, cloudType string) error
	GetCustom(uid, cloudType string) (*models.V1CustomAccount, error)
}

type service struct {
	client *client.V1Client
}

// New creates a new cloudaccounts Service.
func New(c *client.V1Client) Service {
	if c == nil {
		panic("cloudaccounts: V1Client must not be nil")
	}
	return &service{client: c}
}

var errUIDRequired = errors.New("UID is required")
