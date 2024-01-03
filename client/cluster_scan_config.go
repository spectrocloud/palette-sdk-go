package client

import (
	"github.com/spectrocloud/hapi/models"
	clusterC "github.com/spectrocloud/hapi/spectrocluster/client/v1"

	"github.com/spectrocloud/palette-sdk-go/client/herr"
)

func (h *V1Client) GetClusterScanConfig(uid, clusterContext string) (*models.V1ClusterComplianceScan, error) {
	if h.GetClusterScanConfigFn != nil {
		return h.GetClusterScanConfigFn(uid)
	}

	var params *clusterC.V1ClusterFeatureComplianceScanGetParams
	switch clusterContext {
	case "project":
		params = clusterC.NewV1ClusterFeatureComplianceScanGetParamsWithContext(h.Ctx).WithUID(uid)
	case "tenant":
		params = clusterC.NewV1ClusterFeatureComplianceScanGetParams().WithUID(uid)
	}

	success, err := h.GetClusterClient().V1ClusterFeatureComplianceScanGet(params)
	if err != nil {
		if herr.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}

	return success.Payload, nil
}

func (h *V1Client) CreateClusterScanConfig(uid string, config *models.V1ClusterComplianceScheduleConfig, ClusterContext string) error {
	var params *clusterC.V1ClusterFeatureComplianceScanCreateParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1ClusterFeatureComplianceScanCreateParamsWithContext(h.Ctx).WithUID(uid).WithBody(config)
	case "tenant":
		params = clusterC.NewV1ClusterFeatureComplianceScanCreateParams().WithUID(uid).WithBody(config)
	}

	_, err := h.GetClusterClient().V1ClusterFeatureComplianceScanCreate(params)
	return err
}

func (h *V1Client) UpdateClusterScanConfig(uid string, config *models.V1ClusterComplianceScheduleConfig, ClusterContext string) error {
	var params *clusterC.V1ClusterFeatureComplianceScanUpdateParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1ClusterFeatureComplianceScanUpdateParamsWithContext(h.Ctx).WithUID(uid).WithBody(config)
	case "tenant":
		params = clusterC.NewV1ClusterFeatureComplianceScanUpdateParams().WithUID(uid).WithBody(config)
	}

	_, err := h.GetClusterClient().V1ClusterFeatureComplianceScanUpdate(params)
	return err
}

func (h *V1Client) ApplyClusterScanConfig(uid string, config *models.V1ClusterComplianceScheduleConfig, ClusterContext string) error {
	if policy, err := h.GetClusterScanConfig(uid, ClusterContext); err != nil {
		return err
	} else if policy == nil {
		return h.CreateClusterScanConfig(uid, config, ClusterContext)
	} else {
		return h.UpdateClusterScanConfig(uid, config, ClusterContext)
	}
}
