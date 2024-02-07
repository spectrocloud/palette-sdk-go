package client

import (
	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
	"github.com/spectrocloud/palette-sdk-go/client/herr"
)

func (h *V1Client) GetClusterScanConfig(uid, scope string) (*models.V1ClusterComplianceScan, error) {
	var params *clientV1.V1ClusterFeatureComplianceScanGetParams
	switch scope {
	case "project":
		params = clientV1.NewV1ClusterFeatureComplianceScanGetParamsWithContext(h.Ctx).WithUID(uid)
	case "tenant":
		params = clientV1.NewV1ClusterFeatureComplianceScanGetParams().WithUID(uid)
	}

	success, err := h.Client.V1ClusterFeatureComplianceScanGet(params)
	if err != nil {
		if herr.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}

	return success.Payload, nil
}

func (h *V1Client) CreateClusterScanConfig(uid string, config *models.V1ClusterComplianceScheduleConfig, scope string) error {
	var params *clientV1.V1ClusterFeatureComplianceScanCreateParams
	switch scope {
	case "project":
		params = clientV1.NewV1ClusterFeatureComplianceScanCreateParamsWithContext(h.Ctx).WithUID(uid).WithBody(config)
	case "tenant":
		params = clientV1.NewV1ClusterFeatureComplianceScanCreateParams().WithUID(uid).WithBody(config)
	}

	_, err := h.Client.V1ClusterFeatureComplianceScanCreate(params)
	return err
}

func (h *V1Client) UpdateClusterScanConfig(uid string, config *models.V1ClusterComplianceScheduleConfig, scope string) error {
	var params *clientV1.V1ClusterFeatureComplianceScanUpdateParams
	switch scope {
	case "project":
		params = clientV1.NewV1ClusterFeatureComplianceScanUpdateParamsWithContext(h.Ctx).WithUID(uid).WithBody(config)
	case "tenant":
		params = clientV1.NewV1ClusterFeatureComplianceScanUpdateParams().WithUID(uid).WithBody(config)
	}

	_, err := h.Client.V1ClusterFeatureComplianceScanUpdate(params)
	return err
}

func (h *V1Client) ApplyClusterScanConfig(uid string, config *models.V1ClusterComplianceScheduleConfig, scope string) error {
	if policy, err := h.GetClusterScanConfig(uid, scope); err != nil {
		return err
	} else if policy == nil {
		return h.CreateClusterScanConfig(uid, config, scope)
	} else {
		return h.UpdateClusterScanConfig(uid, config, scope)
	}
}
