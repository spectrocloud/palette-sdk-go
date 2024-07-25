package client

import (
	clientv1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
)

// UpdateClusterHostConfig updates an existing cluster's virtual cluster hosting config.
func (h *V1Client) UpdateClusterHostConfig(uid string, config *models.V1HostClusterConfigEntity) error {
	params := clientv1.NewV1HostClusterConfigUpdateParamsWithContext(h.ctx).
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
