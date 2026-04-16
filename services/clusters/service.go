// Package clusters provides operations for Spectro Cloud cluster lifecycle
// management across all supported cloud providers.
package clusters

import (
	"errors"

	"github.com/spectrocloud/palette-sdk-go/api/models"
	"github.com/spectrocloud/palette-sdk-go/client"
)

// Service defines operations for cluster management.
type Service interface {
	// Core CRUD
	Get(uid string) (*models.V1SpectroCluster, error)
	GetByName(name string, virtual bool) (*models.V1SpectroCluster, error)
	GetSummary(uid string) (*models.V1SpectroClusterUIDSummary, error)
	Delete(uid string) error
	ForceDelete(uid string, forceDelete bool) error
	Search(filter *models.V1SearchFilterSpec, sort []*models.V1SearchFilterSortSpec) ([]*models.V1SpectroClusterSummary, error)
	List() ([]*models.V1SpectroClusterSummary, error)

	// KubeConfig
	GetClientKubeConfig(uid string) (string, error)
	GetAdminKubeConfig(uid string) (string, error)
	GetImportManifest(uid string) (string, error)

	// Profile operations on cluster
	UpdateProfileValues(uid string, profiles *models.V1SpectroClusterProfiles) error
	PatchProfileValues(uid string, profiles *models.V1SpectroClusterProfiles) error

	// Variables
	GetVariables(uid string) ([]*models.V1SpectroClusterVariables, error)
	UpdateVariables(uid string, body []*models.V1SpectroClusterVariableUpdateEntity) error

	// Repave
	ApproveRepave(uid string) error
	GetRepaveReasons(uid string) ([]string, error)

	// Certificates
	GetKubernetesCerts(uid string) (*models.V1MachineCertificates, error)
	InitiateCertRenewal(uid string) error

	// Logs
	InitiateLogDownload(uid string, req *models.V1ClusterLogFetcherRequest) (*string, error)
	GetLogFetcherStatus(uid string, logFetcherUID *string) (*models.V1ClusterLogFetcher, error)

	// Backup config
	GetBackupConfig(uid string) (*models.V1ClusterBackup, error)
	CreateBackupConfig(uid string, config *models.V1ClusterBackupConfig) error
	UpdateBackupConfig(uid string, config *models.V1ClusterBackupConfig) error

	// Provider-specific create
	CreateAws(cluster *models.V1SpectroAwsClusterEntity) (string, error)
	CreateAzure(cluster *models.V1SpectroAzureClusterEntity) (string, error)
	CreateGcp(cluster *models.V1SpectroGcpClusterEntity) (string, error)
	CreateVsphere(cluster *models.V1SpectroVsphereClusterEntity) (string, error)
	CreateMaas(cluster *models.V1SpectroMaasClusterEntity) (string, error)
	CreateCloudStack(cluster *models.V1SpectroCloudStackClusterEntity) (string, error)
	CreateEks(cluster *models.V1SpectroEksClusterEntity) (string, error)
	CreateAks(cluster *models.V1SpectroAzureClusterEntity) (string, error)
	CreateGke(cluster *models.V1SpectroGcpClusterEntity) (string, error)
	CreateEdgeNative(cluster *models.V1SpectroEdgeNativeClusterEntity) (string, error)
	CreateEdgeVsphere(cluster *models.V1SpectroVsphereClusterEntity) (string, error)
	CreateVirtual(cluster *models.V1SpectroVirtualClusterEntity) (string, error)
	CreateCustomCloud(cluster *models.V1SpectroCustomClusterEntity, cloudType string) (string, error)

	// Cluster group
	CreateClusterGroup(group *models.V1ClusterGroupEntity) (string, error)
	GetClusterGroup(uid string) (*models.V1ClusterGroup, error)
	DeleteClusterGroup(uid string) error
	UpdateClusterGroup(uid string, group *models.V1ClusterGroupHostClusterEntity) error

	// Metadata
	UpdateMetadata(uid string, config *models.V1ObjectMetaInputEntitySchema) error

	// Config sub-operations
	GetNamespaceConfig(uid string) (*models.V1ClusterNamespaceResources, error)
	GetRbacConfig(uid string) (*models.V1ClusterRbacs, error)
	UpdateOsPatchConfig(uid string, config *models.V1OsPatchEntity) error
}

type service struct {
	client *client.V1Client
}

// New creates a new clusters Service.
func New(c *client.V1Client) Service {
	if c == nil {
		panic("clusters: V1Client must not be nil")
	}
	return &service{client: c}
}

var errUIDRequired = errors.New("cluster UID is required")
