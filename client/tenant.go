package client

import (
	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
)

func (h *V1Client) GetTenantUID() (string, error) {
	params := clientV1.NewV1UsersMeGetParamsWithContext(h.ctx)
	resp, err := h.Client.V1UsersMeGet(params)
	if err != nil {
		return "", err
	}
	return resp.Payload.Status.Tenant.TenantUID, nil
}
