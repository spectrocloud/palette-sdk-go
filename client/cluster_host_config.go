package client

import (
	"github.com/spectrocloud/hapi/models"
	clusterC "github.com/spectrocloud/hapi/spectrocluster/client/v1"
)

func (h *V1Client) UpdateClusterHostConfig(uid, clusterContext string, config *models.V1HostClusterConfigEntity) error {
	var params *clusterC.V1HostClusterConfigUpdateParams
	switch clusterContext {
	case "project":
		params = clusterC.NewV1HostClusterConfigUpdateParamsWithContext(h.Ctx).WithUID(uid).WithBody(config)
	case "tenant":
		params = clusterC.NewV1HostClusterConfigUpdateParams().WithUID(uid).WithBody(config)
	}

	_, err := h.GetClusterClient().V1HostClusterConfigUpdate(params)
	return err
}

func (h *V1Client) ApplyClusterHostConfig(uid, clusterContext string, config *models.V1HostClusterConfigEntity) error {
	if policy, err := h.GetClusterScanConfig(uid, clusterContext); err != nil {
		return err
	} else if policy == nil {
		return h.UpdateClusterHostConfig(uid, clusterContext, config)
	} else {
		return h.UpdateClusterHostConfig(uid, clusterContext, config)
	}
}
