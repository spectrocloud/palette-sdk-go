package client

import (
	"fmt"
	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
	"github.com/spectrocloud/palette-sdk-go/client/herr"
)

func (h *V1Client) GetClusterBackupConfig(uid string) (*models.V1ClusterBackup, error) {
	params := clientV1.NewV1ClusterFeatureBackupGetParamsWithContext(h.ctx).
		WithUID(uid)
	resp, err := h.Client.V1ClusterFeatureBackupGet(params)
	if err != nil {
		if herr.IsNotFound(err) || herr.IsBackupNotConfigured(err) {
			return nil, nil
		}
		return nil, err
	}
	return resp.Payload, nil
}

func (h *V1Client) CreateClusterBackupConfig(uid string, config *models.V1ClusterBackupConfig) error {
	params := clientV1.NewV1ClusterFeatureBackupCreateParamsWithContext(h.ctx).
		WithUID(uid).
		WithBody(config)
	_, err := h.Client.V1ClusterFeatureBackupCreate(params)
	return err
}

func (h *V1Client) UpdateClusterBackupConfig(uid string, config *models.V1ClusterBackupConfig) error {
	params := clientV1.NewV1ClusterFeatureBackupUpdateParamsWithContext(h.ctx).
		WithUID(uid).
		WithBody(config)
	_, err := h.Client.V1ClusterFeatureBackupUpdate(params)
	return err
}

func (h *V1Client) ApplyClusterBackupConfig(uid string, config *models.V1ClusterBackupConfig) error {
	policy, err := h.GetClusterBackupConfig(uid)
	if err != nil {
		return err
	}
	if policy == nil {
		return h.CreateClusterBackupConfig(uid, config)
	}
	return h.UpdateClusterBackupConfig(uid, config)
}

func (h *V1Client) CreateClusterBackupConfigOnDemand(uid string, config *models.V1ClusterBackupConfig) (*models.V1UID, error) {

	params := clientV1.NewV1ClusterFeatureBackupOnDemandCreateParamsWithContext(h.ctx).
		WithBody(config).WithUID(uid)

	response, errMsg := h.Client.V1ClusterFeatureBackupOnDemandCreate(params)
	if errMsg != nil {
		fmt.Println("Failed to get the Backup %s", errMsg.Error())
		return nil, errMsg
	}

	return response.Payload, nil
}

func (h *V1Client) DeleteClusterBackupConfigOnDemand(uid string, config *models.V1ClusterBackupConfig) error {

	params := clientV1.NewV1ClusterFeatureBackupDeleteParamsWithContext(h.ctx).
		WithUID(uid).WithRequestUID(config.BackupLocationUID).WithBackupName(config.BackupName)

	_, errMsg := h.Client.V1ClusterFeatureBackupDelete(params)
	if errMsg != nil {
		fmt.Println("Failed to get the Backup %s", errMsg.Error())
		return errMsg
	}

	return nil
}
