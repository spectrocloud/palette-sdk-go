package vmo

import (
	"errors"

	"github.com/spectrocloud/palette-sdk-go/api/models"
)

func (s *service) CreateVM(clusterUID string, body *models.V1ClusterVirtualMachine) (*models.V1ClusterVirtualMachine, error) {
	if clusterUID == "" {
		return nil, errClusterUIDRequired
	}
	if body == nil {
		return nil, errors.New("virtual machine body is required")
	}
	return s.client.CreateVirtualMachine(clusterUID, body)
}

func (s *service) GetVM(clusterUID, namespace, name string) (*models.V1ClusterVirtualMachine, error) {
	if clusterUID == "" {
		return nil, errClusterUIDRequired
	}
	if name == "" {
		return nil, errors.New("VM name is required")
	}
	return s.client.GetVirtualMachine(clusterUID, namespace, name)
}

func (s *service) UpdateVM(cluster *models.V1SpectroCluster, vmName string, body *models.V1ClusterVirtualMachine) (*models.V1ClusterVirtualMachine, error) {
	if cluster == nil {
		return nil, errors.New("cluster is required")
	}
	if vmName == "" {
		return nil, errors.New("VM name is required")
	}
	return s.client.UpdateVirtualMachine(cluster, vmName, body)
}

func (s *service) DeleteVM(clusterUID, namespace, name string) error {
	if clusterUID == "" {
		return errClusterUIDRequired
	}
	if name == "" {
		return errors.New("VM name is required")
	}
	return s.client.DeleteVirtualMachine(clusterUID, namespace, name)
}

func (s *service) IsVMExists(clusterUID, name, namespace string) (bool, error) {
	if clusterUID == "" {
		return false, errClusterUIDRequired
	}
	return s.client.IsVMExists(clusterUID, name, namespace)
}

func (s *service) ListVMs(cluster *models.V1SpectroCluster) ([]*models.V1ClusterVirtualMachine, error) {
	if cluster == nil {
		return nil, errors.New("cluster is required")
	}
	return s.client.GetVirtualMachines(cluster)
}
