package client

import (
	"errors"

	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
)

func (h *V1Client) GetMgmtAppVersion() (*models.V1MgmtAppVersionVersion, error) {
	params := clientV1.NewV1MgmtAppVersionGetParams().WithContext(h.Ctx)
	resp, err := h.GetClient().V1MgmtAppVersionGet(params)
	if err != nil || resp == nil {
		return nil, err
	}
	if resp.Payload == nil || (resp.Payload != nil && resp.Payload.Version == nil) {
		return nil, errors.New("failed to detect mgmt app version")
	}
	return resp.Payload.Version, nil
}
