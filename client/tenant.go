package client

import "errors"

// GetTenantUID retrieves the tenant UID of the authenticated user.
func (h *V1Client) GetTenantUID() (string, error) {
	resp, err := h.GetUsersInfo()
	if err != nil {
		return "", err
	}
	if resp == nil {
		return "", errors.New("empty response received from GetUsersInfo()")
	}
	return resp.TenantUID, nil
}
