package client

import (
	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
)

func (h *V1Client) UpdateClusterHostConfig(uid string, config *models.V1HostClusterConfigEntity) error {
	params := clientV1.NewV1HostClusterConfigUpdateParamsWithContext(h.ctx).
		WithUID(uid).
		WithBody(config)
	_, err := h.Client.V1HostClusterConfigUpdate(params)
	return err
}

func (h *V1Client) ApplyClusterHostConfig(uid string, config *models.V1HostClusterConfigEntity) error {
	policy, err := h.GetClusterScanConfig(uid)
	if err != nil {
		return err
	}
	if policy == nil {
		return h.UpdateClusterHostConfig(uid, config)
	}
	return h.UpdateClusterHostConfig(uid, config)
}
