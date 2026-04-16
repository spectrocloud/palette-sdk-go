package vmo

import "errors"

func (s *service) StartVM(clusterUID, vmName, vmNamespace string) error {
	if clusterUID == "" {
		return errClusterUIDRequired
	}
	if vmName == "" {
		return errors.New("VM name is required")
	}
	return s.client.StartVirtualMachine(clusterUID, vmName, vmNamespace)
}

func (s *service) StopVM(clusterUID, vmName, vmNamespace string) error {
	if clusterUID == "" {
		return errClusterUIDRequired
	}
	if vmName == "" {
		return errors.New("VM name is required")
	}
	return s.client.StopVirtualMachine(clusterUID, vmName, vmNamespace)
}

func (s *service) PauseVM(clusterUID, vmName, vmNamespace string) error {
	if clusterUID == "" {
		return errClusterUIDRequired
	}
	if vmName == "" {
		return errors.New("VM name is required")
	}
	return s.client.PauseVirtualMachine(clusterUID, vmName, vmNamespace)
}

func (s *service) ResumeVM(clusterUID, vmName, vmNamespace string) error {
	if clusterUID == "" {
		return errClusterUIDRequired
	}
	if vmName == "" {
		return errors.New("VM name is required")
	}
	return s.client.ResumeVirtualMachine(clusterUID, vmName, vmNamespace)
}

func (s *service) RestartVM(clusterUID, vmName, vmNamespace string) error {
	if clusterUID == "" {
		return errClusterUIDRequired
	}
	if vmName == "" {
		return errors.New("VM name is required")
	}
	return s.client.RestartVirtualMachine(clusterUID, vmName, vmNamespace)
}
