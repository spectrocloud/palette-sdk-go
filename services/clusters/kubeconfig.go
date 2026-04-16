package clusters

func (s *service) GetClientKubeConfig(uid string) (string, error) {
	if uid == "" {
		return "", errUIDRequired
	}
	return s.client.GetClusterClientKubeConfig(uid)
}

func (s *service) GetAdminKubeConfig(uid string) (string, error) {
	if uid == "" {
		return "", errUIDRequired
	}
	return s.client.GetClusterAdminKubeConfig(uid)
}

func (s *service) GetImportManifest(uid string) (string, error) {
	if uid == "" {
		return "", errUIDRequired
	}
	return s.client.GetClusterImportManifest(uid)
}
