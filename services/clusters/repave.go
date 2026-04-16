package clusters

func (s *service) ApproveRepave(uid string) error {
	if uid == "" {
		return errUIDRequired
	}
	return s.client.ApproveClusterRepave(uid)
}

func (s *service) GetRepaveReasons(uid string) ([]string, error) {
	if uid == "" {
		return nil, errUIDRequired
	}
	return s.client.GetRepaveReasons(uid)
}
