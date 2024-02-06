package client

import (
	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
)

func (h *V1Client) UpdateClusterHostConfig(uid, scope string, config *models.V1HostClusterConfigEntity) error {
	var params *clientV1.V1HostClusterConfigUpdateParams
	switch scope {
	case "project":
		params = clientV1.NewV1HostClusterConfigUpdateParamsWithContext(h.Ctx).WithUID(uid).WithBody(config)
	case "tenant":
		params = clientV1.NewV1HostClusterConfigUpdateParams().WithUID(uid).WithBody(config)
	}

	_, err := h.GetClient().V1HostClusterConfigUpdate(params)
	return err
}

func (h *V1Client) ApplyClusterHostConfig(uid, scope string, config *models.V1HostClusterConfigEntity) error {
	if policy, err := h.GetClusterScanConfig(uid, scope); err != nil {
		return err
	} else if policy == nil {
		return h.UpdateClusterHostConfig(uid, scope, config)
	} else {
		return h.UpdateClusterHostConfig(uid, scope, config)
	}
}
