package client

import (
	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
)

func (h *V1Client) UpdateClusterMetadata(uid string, config *models.V1ObjectMetaInputEntitySchema) error {
	params := clientV1.NewV1SpectroClustersUIDMetadataUpdateParamsWithContext(h.ctx).
		WithUID(uid).
		WithBody(config)
	_, err := h.Client.V1SpectroClustersUIDMetadataUpdate(params)
	return err
}

func (h *V1Client) UpdateAdditionalClusterMetadata(uid string, additionalMeta *models.V1ClusterMetaAttributeEntity) error {
	params := clientV1.NewV1SpectroClustersUIDClusterMetaAttributeUpdateParamsWithContext(h.ctx).
		WithUID(uid).
		WithBody(additionalMeta)
	_, err := h.Client.V1SpectroClustersUIDClusterMetaAttributeUpdate(params)
	return err
}
