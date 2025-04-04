package client

import (
	clientv1 "github.com/spectrocloud/palette-sdk-go/api/client/version1"
	"github.com/spectrocloud/palette-sdk-go/api/models"
	"github.com/spectrocloud/palette-sdk-go/client/herr"
)

// GetClusterBackupConfig gets the scheduled backup configuration for a cluster.
func (h *V1Client) GetClusterBackupConfig(uid string) (*models.V1ClusterBackup, error) {
	params := clientv1.NewV1ClusterFeatureBackupGetParamsWithContext(h.ctx).
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

// CreateClusterBackupConfig creates a new scheduled backup configuration for a cluster.
func (h *V1Client) CreateClusterBackupConfig(uid string, config *models.V1ClusterBackupConfig) error {
	params := clientv1.NewV1ClusterFeatureBackupCreateParamsWithContext(h.ctx).
		WithUID(uid).
		WithBody(config)
	_, err := h.Client.V1ClusterFeatureBackupCreate(params)
	return err
}

// UpdateClusterBackupConfig updates an existing scheduled backup configuration for a cluster.
func (h *V1Client) UpdateClusterBackupConfig(uid string, config *models.V1ClusterBackupConfig) error {
	params := clientv1.NewV1ClusterFeatureBackupUpdateParamsWithContext(h.ctx).
		WithUID(uid).
		WithBody(config)
	_, err := h.Client.V1ClusterFeatureBackupUpdate(params)
	return err
}

// ApplyClusterBackupConfig creates a new scheduled backup configuration for a cluster or updates its scheduled backup configuration if one exists.
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

// CreateClusterBackupConfigOnDemand creates a new on-demand backup configuration for a cluster.
func (h *V1Client) CreateClusterBackupConfigOnDemand(uid string, config *models.V1ClusterBackupConfig) (*models.V1UID, error) {
	params := clientv1.NewV1ClusterFeatureBackupOnDemandCreateParamsWithContext(h.ctx).
		WithBody(config).
		WithUID(uid)
	resp, err := h.Client.V1ClusterFeatureBackupOnDemandCreate(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// DeleteClusterBackupConfigOnDemand deletes an existing on-demand backup configuration for a cluster.
func (h *V1Client) DeleteClusterBackupConfigOnDemand(uid string, config *models.V1ClusterBackupConfig) error {
	params := clientv1.NewV1ClusterFeatureBackupDeleteParamsWithContext(h.ctx).
		WithUID(uid).
		WithRequestUID(config.BackupLocationUID).
		WithBackupName(config.BackupName)
	_, err := h.Client.V1ClusterFeatureBackupDelete(params)
	return err
}

// CreateClusterRestoreConfigOnDemand creates a new on-demand restore configuration for a cluster.
func (h *V1Client) CreateClusterRestoreConfigOnDemand(uid string, config *models.V1ClusterRestoreConfig) (*models.V1UID, error) {
	params := clientv1.NewV1ClusterFeatureRestoreOnDemandCreateParamsWithContext(h.ctx).
		WithBody(config).
		WithUID(uid)
	resp, err := h.Client.V1ClusterFeatureRestoreOnDemandCreate(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// GetClusterRestoreConfigOnDemand retrieces an existing on-demand restore configuration for a cluster.
func (h *V1Client) GetClusterRestoreConfigOnDemand(uid string) (*models.V1ClusterRestore, error) {
	params := clientv1.NewV1ClusterFeatureRestoreGetParamsWithContext(h.ctx).
		WithUID(uid)
	resp, err := h.Client.V1ClusterFeatureRestoreGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}
