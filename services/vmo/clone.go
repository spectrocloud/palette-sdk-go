package vmo

import "errors"

func (s *service) CloneVM(clusterUID, sourceVMName, cloneVMName, namespace string) error {
	if clusterUID == "" {
		return errClusterUIDRequired
	}
	if sourceVMName == "" {
		return errors.New("source VM name is required")
	}
	if cloneVMName == "" {
		return errors.New("clone VM name is required")
	}
	return s.client.CloneVirtualMachine(clusterUID, sourceVMName, cloneVMName, namespace)
}
