package client

import (
	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
)

func (h *V1Client) UpdateClusterOsPatchConfig(uid, scope string, config *models.V1OsPatchEntity) error {
	var params *clientV1.V1SpectroClustersUIDOsPatchUpdateParams
	switch scope {
	case "project":
		params = clientV1.NewV1SpectroClustersUIDOsPatchUpdateParamsWithContext(h.Ctx).WithUID(uid).WithBody(config)
	case "tenant":
		params = clientV1.NewV1SpectroClustersUIDOsPatchUpdateParams().WithUID(uid).WithBody(config)
	}

	_, err := h.GetClient().V1SpectroClustersUIDOsPatchUpdate(params)
	return err
}
