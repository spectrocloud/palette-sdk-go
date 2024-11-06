package client

import (
	clientv1 "github.com/spectrocloud/palette-sdk-go/api/client/v1"
	"github.com/spectrocloud/palette-sdk-go/api/models"
)

// EnableClusterHybridSettings enables hybrid mode on a cluster.
func (h *V1Client) EnableClusterHybridSettings(uid string) error {
	params := clientv1.NewV1SpectroClustersUIDHybridSettingsParamsWithContext(h.ctx).
		WithBody(&models.V1ClusterHybridSettingsEntity{
			Enable: true,
		}).WithUID(uid)

	_, err := h.Client.V1SpectroClustersUIDHybridSettings(params)
	return err
}
