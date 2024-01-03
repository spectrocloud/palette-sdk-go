package client

import (
	userC "github.com/spectrocloud/hapi/user/client/v1"
)

func (h *V1Client) GetTenantUID() (string, error) {
	params := userC.NewV1UsersMeGetParams()
	me, err := h.GetUserClient().V1UsersMeGet(params)
	if err != nil || me == nil {
		return "", err
	}
	return me.Payload.Status.Tenant.TenantUID, nil
}
