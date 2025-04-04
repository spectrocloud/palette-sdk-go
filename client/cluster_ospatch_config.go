package client

import (
	clientv1 "github.com/spectrocloud/palette-sdk-go/api/client/version1"
	"github.com/spectrocloud/palette-sdk-go/api/models"
)

// UpdateClusterOsPatchConfig updates an existing cluster's OS patching configuration.
func (h *V1Client) UpdateClusterOsPatchConfig(uid string, config *models.V1OsPatchEntity) error {
	params := clientv1.NewV1SpectroClustersUIDOsPatchUpdateParamsWithContext(h.ctx).
		WithUID(uid).
		WithBody(config)
	_, err := h.Client.V1SpectroClustersUIDOsPatchUpdate(params)
	return err
}
