package client

import (
	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
	"github.com/spectrocloud/palette-sdk-go/client/herr"
)

func (h *V1Client) GetClusterScanConfig(uid string) (*models.V1ClusterComplianceScan, error) {
	params := clientV1.NewV1ClusterFeatureComplianceScanGetParamsWithContext(h.ctx).
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

func (h *V1Client) CreateClusterScanConfig(uid string, config *models.V1ClusterComplianceScheduleConfig) error {
	params := clientV1.NewV1ClusterFeatureComplianceScanCreateParamsWithContext(h.ctx).
		WithUID(uid).
		WithBody(config)
	_, err := h.Client.V1ClusterFeatureComplianceScanCreate(params)
	return err
}

func (h *V1Client) UpdateClusterScanConfig(uid string, config *models.V1ClusterComplianceScheduleConfig) error {
	params := clientV1.NewV1ClusterFeatureComplianceScanUpdateParamsWithContext(h.ctx).
		WithUID(uid).
		WithBody(config)
	_, err := h.Client.V1ClusterFeatureComplianceScanUpdate(params)
	return err
}

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
