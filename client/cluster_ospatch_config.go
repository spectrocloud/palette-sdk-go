package client

import (
	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
)

func (h *V1Client) UpdateClusterOsPatchConfig(uid string, config *models.V1OsPatchEntity) error {
	params := clientV1.NewV1SpectroClustersUIDOsPatchUpdateParamsWithContext(h.ctx).
		WithUID(uid).
		WithBody(config)
	_, err := h.Client.V1SpectroClustersUIDOsPatchUpdate(params)
	return err
}
