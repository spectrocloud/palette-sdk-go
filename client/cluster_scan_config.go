package client

import (
	clientv1 "github.com/spectrocloud/palette-sdk-go/api/client/version1"
	"github.com/spectrocloud/palette-sdk-go/api/models"
	"github.com/spectrocloud/palette-sdk-go/client/herr"
)

// GetClusterScanConfig retrieves the cluster scan configuration for a cluster.
func (h *V1Client) GetClusterScanConfig(uid string) (*models.V1ClusterComplianceScan, error) {
	params := clientv1.NewV1ClusterFeatureComplianceScanGetParamsWithContext(h.ctx).
		WithUID(uid)
	resp, err := h.Client.V1ClusterFeatureComplianceScanGet(params)
	if err != nil {
		if herr.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}
	return resp.Payload, nil
}

// CreateClusterScanConfig creates a new cluster scan configuration for a cluster.
func (h *V1Client) CreateClusterScanConfig(uid string, config *models.V1ClusterComplianceScheduleConfig) error {
	params := clientv1.NewV1ClusterFeatureComplianceScanCreateParamsWithContext(h.ctx).
		WithUID(uid).
		WithBody(config)
	_, err := h.Client.V1ClusterFeatureComplianceScanCreate(params)
	return err
}

// UpdateClusterScanConfig updates an existing cluster scan configuration for a cluster.
func (h *V1Client) UpdateClusterScanConfig(uid string, config *models.V1ClusterComplianceScheduleConfig) error {
	params := clientv1.NewV1ClusterFeatureComplianceScanUpdateParamsWithContext(h.ctx).
		WithUID(uid).
		WithBody(config)
	_, err := h.Client.V1ClusterFeatureComplianceScanUpdate(params)
	return err
}

// ApplyClusterScanConfig creates a new cluster scan configuration for a cluster or updates its cluster scan configuration if one exists.
func (h *V1Client) ApplyClusterScanConfig(uid string, config *models.V1ClusterComplianceScheduleConfig) error {
	policy, err := h.GetClusterScanConfig(uid)
	if err != nil {
		return err
	}
	if policy == nil {
		return h.CreateClusterScanConfig(uid, config)
	}
	return h.UpdateClusterScanConfig(uid, config)
}

// GetComplianceScanOnDemandScanLogs retrieves the on-demand scan logs for a cluster.
func (h *V1Client) GetComplianceScanOnDemandScanLogs(uid string) (*models.V1ClusterComplianceScanLogs, error) {
	params := clientv1.NewV1ClusterFeatureComplianceScanLogsGetParamsWithContext(h.ctx).
		WithUID(uid)
	resp, err := h.Client.V1ClusterFeatureComplianceScanLogsGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// CreateComplianceScanOnDemandCreateClusterScanConfig creates a new on-demand scan configuration for a cluster.
func (h *V1Client) CreateComplianceScanOnDemandCreateClusterScanConfig(uid string, config *models.V1ClusterComplianceOnDemandConfig) error {
	params := clientv1.NewV1ClusterFeatureComplianceScanOnDemandCreateParamsWithContext(h.ctx).
		WithUID(uid).
		WithBody(config)
	_, err := h.Client.V1ClusterFeatureComplianceScanOnDemandCreate(params)
	return err
}
