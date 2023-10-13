package client

import (
	"github.com/spectrocloud/hapi/models"
	clusterC "github.com/spectrocloud/hapi/spectrocluster/client/v1"
)

func (h *V1Client) UpdateClusterMetadata(uid string, clusterContext string, config *models.V1ObjectMetaInputEntitySchema) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return err
	}

	var params *clusterC.V1SpectroClustersUIDMetadataUpdateParams
	switch clusterContext {
	case "project":
		params = clusterC.NewV1SpectroClustersUIDMetadataUpdateParams().WithContext(h.Ctx).WithUID(uid).WithBody(config)
	case "tenant":
		params = clusterC.NewV1SpectroClustersUIDMetadataUpdateParams().WithUID(uid).WithBody(config)
	}

	//params := clusterC.NewV1SpectroClustersUIDMetadataUpdateParams().WithContext(h.Ctx).WithUID(uid).WithBody(config)
	_, err = client.V1SpectroClustersUIDMetadataUpdate(params)
	return err
}

func (h *V1Client) UpdateAdditionalClusterMetadata(uid string, clusterContext string, additionalMeta *models.V1ClusterMetaAttributeEntity) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return err
	}
	var params *clusterC.V1SpectroClustersUIDClusterMetaAttributeUpdateParams
	switch clusterContext {
	case "project":
		params = clusterC.NewV1SpectroClustersUIDClusterMetaAttributeUpdateParams().WithContext(h.Ctx).WithUID(uid).WithBody(additionalMeta)
	case "tenant":
		params = clusterC.NewV1SpectroClustersUIDClusterMetaAttributeUpdateParams().WithUID(uid).WithBody(additionalMeta)
	}
	_, err = client.V1SpectroClustersUIDClusterMetaAttributeUpdate(params)
	return err
}
