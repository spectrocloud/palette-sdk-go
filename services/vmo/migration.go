package vmo

import "errors"

func (s *service) MigrateVM(clusterUID, vmName, vmNamespace string) error {
	if clusterUID == "" {
		return errClusterUIDRequired
	}
	if vmName == "" {
		return errors.New("VM name is required")
	}
	return s.client.MigrateVirtualMachineNodeToNode(clusterUID, vmName, vmNamespace)
}
