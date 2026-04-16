package platform

func (s *service) GetTenantUID() (string, error) {
	return s.client.GetTenantUID()
}
