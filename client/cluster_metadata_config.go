package client

import (
	"github.com/spectrocloud/hapi/apiutil/transport"
	"github.com/spectrocloud/hapi/models"
	clusterC "github.com/spectrocloud/hapi/spectrocluster/client/v1"
)

func (h *V1Client) UpdateClusterMetadata(uid string, config *models.V1ObjectMetaInputEntitySchema) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return err
	}

	params := clusterC.NewV1SpectroClustersUIDMetadataUpdateParams().WithContext(h.Ctx).WithUID(uid).WithBody(config)
	_, err = client.V1SpectroClustersUIDMetadataUpdate(params)
	// handle tenant context here cluster may be a tenant cluster
	if e, ok := err.(*transport.TransportError); ok && e.HttpCode == 404 {
		params = clusterC.NewV1SpectroClustersUIDMetadataUpdateParams().WithUID(uid).WithBody(config)
		_, err = client.V1SpectroClustersUIDMetadataUpdate(params)
		if err != nil {
			return err
		}
	}

	return err
}
