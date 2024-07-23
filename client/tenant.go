package client

// GetTenantUID retrieves the tenant UID of the authenticated user.
func (h *V1Client) GetTenantUID() (string, error) {
	resp, err := h.GetMe()
	if err != nil {
		return "", err
	}
	return resp.Status.Tenant.TenantUID, nil
}
