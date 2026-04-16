// Package vmo provides operations for virtual machine orchestration including
// VM lifecycle, migration, cloning, and data volume management.
package vmo

import (
	"errors"

	"github.com/spectrocloud/palette-sdk-go/api/models"
	"github.com/spectrocloud/palette-sdk-go/client"
)

// Service defines operations for virtual machine management.
type Service interface {
	// VM CRUD
	CreateVM(clusterUID string, body *models.V1ClusterVirtualMachine) (*models.V1ClusterVirtualMachine, error)
	GetVM(clusterUID, namespace, name string) (*models.V1ClusterVirtualMachine, error)
	UpdateVM(cluster *models.V1SpectroCluster, vmName string, body *models.V1ClusterVirtualMachine) (*models.V1ClusterVirtualMachine, error)
	DeleteVM(clusterUID, namespace, name string) error
	IsVMExists(clusterUID, name, namespace string) (bool, error)
	ListVMs(cluster *models.V1SpectroCluster) ([]*models.V1ClusterVirtualMachine, error)

	// Lifecycle
	StartVM(clusterUID, vmName, vmNamespace string) error
	StopVM(clusterUID, vmName, vmNamespace string) error
	PauseVM(clusterUID, vmName, vmNamespace string) error
	ResumeVM(clusterUID, vmName, vmNamespace string) error
	RestartVM(clusterUID, vmName, vmNamespace string) error

	// Migration
	MigrateVM(clusterUID, vmName, vmNamespace string) error

	// Clone
	CloneVM(clusterUID, sourceVMName, cloneVMName, namespace string) error

	// Data Volumes
	CreateDataVolume(clusterUID, name string, body *models.V1VMAddVolumeEntity) (string, error)
	DeleteDataVolume(clusterUID, namespace, name string, body *models.V1VMRemoveVolumeEntity) error
}

type service struct {
	client *client.V1Client
}

// New creates a new VMO Service.
func New(c *client.V1Client) Service {
	if c == nil {
		panic("vmo: V1Client must not be nil")
	}
	return &service{client: c}
}

var errClusterUIDRequired = errors.New("cluster UID is required")
