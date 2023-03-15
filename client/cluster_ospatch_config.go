package client

import (
	"github.com/spectrocloud/hapi/apiutil/transport"
	"github.com/spectrocloud/hapi/models"
	clusterC "github.com/spectrocloud/hapi/spectrocluster/client/v1"
)

func (h *V1Client) UpdateClusterOsPatchConfig(uid string, config *models.V1OsPatchEntity) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return err
	}

	params := clusterC.NewV1SpectroClustersUIDOsPatchUpdateParamsWithContext(h.Ctx).WithUID(uid).WithBody(config)
	_, err = client.V1SpectroClustersUIDOsPatchUpdate(params)
	if e, ok := err.(*transport.TransportError); ok && e.HttpCode == 404 {
		params = clusterC.NewV1SpectroClustersUIDOsPatchUpdateParams().WithUID(uid).WithBody(config)
		_, err = client.V1SpectroClustersUIDOsPatchUpdate(params)
	}
	return err
}
