package client

import (
	"github.com/spectrocloud/hapi/models"
	clusterC "github.com/spectrocloud/hapi/spectrocluster/client/v1"
)

func (h *V1Client) UpdateClusterOsPatchConfig(uid, clusterContext string, config *models.V1OsPatchEntity) error {
	var params *clusterC.V1SpectroClustersUIDOsPatchUpdateParams
	switch clusterContext {
	case "project":
		params = clusterC.NewV1SpectroClustersUIDOsPatchUpdateParamsWithContext(h.Ctx).WithUID(uid).WithBody(config)
	case "tenant":
		params = clusterC.NewV1SpectroClustersUIDOsPatchUpdateParams().WithUID(uid).WithBody(config)
	}

	_, err := h.GetClusterClient().V1SpectroClustersUIDOsPatchUpdate(params)
	return err
}
