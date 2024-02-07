package client

import (
	"errors"

	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
	"github.com/spectrocloud/palette-sdk-go/client/herr"
)

func (h *V1Client) GetClusterBackupConfig(uid, scope string) (*models.V1ClusterBackup, error) {
	var params *clientV1.V1ClusterFeatureBackupGetParams
	switch scope {
	case "project":
		params = clientV1.NewV1ClusterFeatureBackupGetParamsWithContext(h.Ctx).WithUID(uid)
	case "tenant":
		params = clientV1.NewV1ClusterFeatureBackupGetParams().WithUID(uid)
	default:
		return nil, errors.New("invalid cluster scope specified")

	}

	success, err := h.Client.V1ClusterFeatureBackupGet(params)
	if err != nil {
		if herr.IsNotFound(err) || herr.IsBackupNotConfigured(err) {
			return nil, nil
		}
		return nil, err
	}

	return success.Payload, nil
}

func (h *V1Client) CreateClusterBackupConfig(uid string, config *models.V1ClusterBackupConfig, scope string) error {
	var params *clientV1.V1ClusterFeatureBackupCreateParams
	switch scope {
	case "project":
		params = clientV1.NewV1ClusterFeatureBackupCreateParamsWithContext(h.Ctx).WithUID(uid).WithBody(config)
	case "tenant":
		params = clientV1.NewV1ClusterFeatureBackupCreateParams().WithUID(uid).WithBody(config)
	}

	_, err := h.Client.V1ClusterFeatureBackupCreate(params)
	return err
}

func (h *V1Client) UpdateClusterBackupConfig(uid string, config *models.V1ClusterBackupConfig, scope string) error {
	var params *clientV1.V1ClusterFeatureBackupUpdateParams
	switch scope {
	case "project":
		params = clientV1.NewV1ClusterFeatureBackupUpdateParamsWithContext(h.Ctx).WithUID(uid).WithBody(config)
	case "tenant":
		params = clientV1.NewV1ClusterFeatureBackupUpdateParams().WithUID(uid).WithBody(config)
	}
	_, err := h.Client.V1ClusterFeatureBackupUpdate(params)
	return err
}

func (h *V1Client) ApplyClusterBackupConfig(uid string, config *models.V1ClusterBackupConfig, scope string) error {
	if policy, err := h.GetClusterBackupConfig(uid, scope); err != nil {
		return err
	} else if policy == nil {
		return h.CreateClusterBackupConfig(uid, config, scope)
	} else {
		return h.UpdateClusterBackupConfig(uid, config, scope)
	}
}
