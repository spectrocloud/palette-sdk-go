package client

import clientV1 "github.com/spectrocloud/palette-api-go/client/v1"

func (h *V1Client) GetTenantUID() (string, error) {
	params := clientV1.NewV1UsersMeGetParams()
	me, err := h.GetClient().V1UsersMeGet(params)
	if err != nil || me == nil {
		return "", err
	}
	return me.Payload.Status.Tenant.TenantUID, nil
}
