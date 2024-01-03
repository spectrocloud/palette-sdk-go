package client

import (
	"errors"

	"github.com/spectrocloud/hapi/models"
	clusterC "github.com/spectrocloud/hapi/spectrocluster/client/v1"

	"github.com/spectrocloud/palette-sdk-go/client/herr"
)

func (h *V1Client) GetClusterBackupConfig(uid, ClusterContext string) (*models.V1ClusterBackup, error) {
	if h.GetClusterBackupConfigFn != nil {
		return h.GetClusterBackupConfigFn(uid)
	}
	var params *clusterC.V1ClusterFeatureBackupGetParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1ClusterFeatureBackupGetParamsWithContext(h.Ctx).WithUID(uid)
	case "tenant":
		params = clusterC.NewV1ClusterFeatureBackupGetParams().WithUID(uid)
	default:
		return nil, errors.New("invalid cluster scope specified")

	}

	success, err := h.GetClusterClient().V1ClusterFeatureBackupGet(params)
	if err != nil {
		if herr.IsNotFound(err) || herr.IsBackupNotConfigured(err) {
			return nil, nil
		}
		return nil, err
	}

	return success.Payload, nil
}

func (h *V1Client) CreateClusterBackupConfig(uid string, config *models.V1ClusterBackupConfig, ClusterContext string) error {
	var params *clusterC.V1ClusterFeatureBackupCreateParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1ClusterFeatureBackupCreateParamsWithContext(h.Ctx).WithUID(uid).WithBody(config)
	case "tenant":
		params = clusterC.NewV1ClusterFeatureBackupCreateParams().WithUID(uid).WithBody(config)
	}

	_, err := h.GetClusterClient().V1ClusterFeatureBackupCreate(params)
	return err
}

func (h *V1Client) UpdateClusterBackupConfig(uid string, config *models.V1ClusterBackupConfig, ClusterContext string) error {
	var params *clusterC.V1ClusterFeatureBackupUpdateParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1ClusterFeatureBackupUpdateParamsWithContext(h.Ctx).WithUID(uid).WithBody(config)
	case "tenant":
		params = clusterC.NewV1ClusterFeatureBackupUpdateParams().WithUID(uid).WithBody(config)
	}
	_, err := h.GetClusterClient().V1ClusterFeatureBackupUpdate(params)
	return err
}

func (h *V1Client) ApplyClusterBackupConfig(uid string, config *models.V1ClusterBackupConfig, ClusterContext string) error {
	if policy, err := h.GetClusterBackupConfig(uid, ClusterContext); err != nil {
		return err
	} else if policy == nil {
		return h.CreateClusterBackupConfig(uid, config, ClusterContext)
	} else {
		return h.UpdateClusterBackupConfig(uid, config, ClusterContext)
	}
}
