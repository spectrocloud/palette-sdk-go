package client

import (
	"errors"

	mgmtC "github.com/spectrocloud/hapi/mgmt/client/v1"
	"github.com/spectrocloud/hapi/models"
)

func (h *V1Client) GetMgmtAppVersion() (*models.V1MgmtAppVersionVersion, error) {
	params := mgmtC.NewV1MgmtAppVersionGetParams().WithContext(h.Ctx)
	resp, err := h.GetMgmtClient().V1MgmtAppVersionGet(params)
	if err != nil || resp == nil {
		return nil, err
	}
	if resp.Payload == nil || (resp.Payload != nil && resp.Payload.Version == nil) {
		return nil, errors.New("failed to detect mgmt app version")
	}
	return resp.Payload.Version, nil
}
