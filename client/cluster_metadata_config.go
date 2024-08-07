package client

import (
	clientv1 "github.com/spectrocloud/palette-sdk-go/api/client/v1"
	"github.com/spectrocloud/palette-sdk-go/api/models"
)

// UpdateClusterMetadata updates an existing cluster's metadata: annotations, labels, and name.
func (h *V1Client) UpdateClusterMetadata(uid string, config *models.V1ObjectMetaInputEntitySchema) error {
	params := clientv1.NewV1SpectroClustersUIDMetadataUpdateParamsWithContext(h.ctx).
		WithUID(uid).
		WithBody(config)
	_, err := h.Client.V1SpectroClustersUIDMetadataUpdate(params)
	return err
}

// UpdateAdditionalClusterMetadata updates an existing cluster's additional metadata attributes.
// TODO: what are additional metadata attributes?
func (h *V1Client) UpdateAdditionalClusterMetadata(uid string, additionalMeta *models.V1ClusterMetaAttributeEntity) error {
	params := clientv1.NewV1SpectroClustersUIDClusterMetaAttributeUpdateParamsWithContext(h.ctx).
		WithUID(uid).
		WithBody(additionalMeta)
	_, err := h.Client.V1SpectroClustersUIDClusterMetaAttributeUpdate(params)
	return err
}
