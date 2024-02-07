package client

import (
	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
)

func (h *V1Client) UpdateClusterMetadata(uid, scope string, config *models.V1ObjectMetaInputEntitySchema) error {
	var params *clientV1.V1SpectroClustersUIDMetadataUpdateParams
	switch scope {
	case "project":
		params = clientV1.NewV1SpectroClustersUIDMetadataUpdateParams().WithContext(h.Ctx).WithUID(uid).WithBody(config)
	case "tenant":
		params = clientV1.NewV1SpectroClustersUIDMetadataUpdateParams().WithUID(uid).WithBody(config)
	}

	//params := clientV1.NewV1SpectroClustersUIDMetadataUpdateParams().WithContext(h.Ctx).WithUID(uid).WithBody(config)
	_, err := h.Client.V1SpectroClustersUIDMetadataUpdate(params)
	return err
}

func (h *V1Client) UpdateAdditionalClusterMetadata(uid, scope string, additionalMeta *models.V1ClusterMetaAttributeEntity) error {
	var params *clientV1.V1SpectroClustersUIDClusterMetaAttributeUpdateParams
	switch scope {
	case "project":
		params = clientV1.NewV1SpectroClustersUIDClusterMetaAttributeUpdateParams().WithContext(h.Ctx).WithUID(uid).WithBody(additionalMeta)
	case "tenant":
		params = clientV1.NewV1SpectroClustersUIDClusterMetaAttributeUpdateParams().WithUID(uid).WithBody(additionalMeta)
	}

	_, err := h.Client.V1SpectroClustersUIDClusterMetaAttributeUpdate(params)
	return err
}
