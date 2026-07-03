package client

import (
	clientv1 "github.com/spectrocloud/palette-sdk-go/api/client/version1"
	"github.com/spectrocloud/palette-sdk-go/api/models"
	"github.com/spectrocloud/palette-sdk-go/client/apiutil"
)

func toCloudWatchConfig(config *models.V1DataSinkConfig) *models.V1CloudWatchConfig {
	if config == nil || config.Spec == nil || len(config.Spec.AuditDataSinks) == 0 {
		return nil
	}
	for _, sink := range config.Spec.AuditDataSinks {
		if sink != nil && sink.CloudWatch != nil {
			return &models.V1CloudWatchConfig{
				Credentials: sink.CloudWatch.Credentials,
				Group:       sink.CloudWatch.Group,
				Region:      sink.CloudWatch.Region,
				Stream:      sink.CloudWatch.Stream,
			}
		}
	}
	return nil
}

// ValidateCloudWatchAuditTrail validates CloudWatch audit trail credentials and config.
func (h *V1Client) ValidateCloudWatchAuditTrail(config *models.V1CloudWatchConfig) error {
	params := clientv1.NewV1CloudsAwsCloudWatchValidateParamsWithContext(h.ctx).
		WithCloudWatchConfig(config)
	_, err := h.Client.V1CloudsAwsCloudWatchValidate(params)
	return err
}

// CreateCloudWatchAuditTrail creates a CloudWatch audit trail for the tenant.
func (h *V1Client) CreateCloudWatchAuditTrail(tenantUID string, config *models.V1DataSinkConfig) (string, error) {
	if cwConfig := toCloudWatchConfig(config); cwConfig != nil {
		if err := h.ValidateCloudWatchAuditTrail(cwConfig); err != nil {
			return "", err
		}
	}
	params := clientv1.NewV1TenantUIDAssetsDataSinksCreateParamsWithContext(h.ctx).
		WithTenantUID(tenantUID).
		WithBody(config)
	resp, err := h.Client.V1TenantUIDAssetsDataSinksCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

// GetCloudWatchAuditTrail retrieves the CloudWatch audit trail for the tenant.
func (h *V1Client) GetCloudWatchAuditTrail(tenantUID string) (*models.V1DataSinkConfig, error) {
	params := clientv1.NewV1TenantUIDAssetsDataSinksGetParamsWithContext(h.ctx).
		WithTenantUID(tenantUID)
	resp, err := h.Client.V1TenantUIDAssetsDataSinksGet(params)
	if apiutil.Is404(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// UpdateCloudWatchAuditTrail updates the CloudWatch audit trail for the tenant.
func (h *V1Client) UpdateCloudWatchAuditTrail(tenantUID string, config *models.V1DataSinkConfig) error {
	if cwConfig := toCloudWatchConfig(config); cwConfig != nil {
		if err := h.ValidateCloudWatchAuditTrail(cwConfig); err != nil {
			return err
		}
	}
	params := clientv1.NewV1TenantUIDAssetsDataSinksUpdateParamsWithContext(h.ctx).
		WithTenantUID(tenantUID).
		WithBody(config)
	_, err := h.Client.V1TenantUIDAssetsDataSinksUpdate(params)
	return err
}

// DeleteCloudWatchAuditTrail deletes the CloudWatch audit trail for the tenant.
func (h *V1Client) DeleteCloudWatchAuditTrail(tenantUID string) error {
	params := clientv1.NewV1TenantUIDAssetsDataSinksDeleteParamsWithContext(h.ctx).
		WithTenantUID(tenantUID)
	_, err := h.Client.V1TenantUIDAssetsDataSinksDelete(params)
	return err
}

// ValidateSplunkAuditTrail validates Splunk HEC audit trail credentials and config.
func (h *V1Client) ValidateSplunkAuditTrail(tenantUID string, spec *models.V1SplunkSinkSpec) error {
	params := clientv1.NewV1TenantUIDDatasinksSplunkValidateParamsWithContext(h.ctx).
		WithTenantUID(tenantUID).
		WithBody(spec)
	_, err := h.Client.V1TenantUIDDatasinksSplunkValidate(params)
	return err
}

// CreateSplunkAuditTrail creates a Splunk audit trail for the tenant.
func (h *V1Client) CreateSplunkAuditTrail(tenantUID string, entity *models.V1SplunkSinkEntity) (string, error) {
	if entity != nil && entity.Spec != nil {
		if err := h.ValidateSplunkAuditTrail(tenantUID, entity.Spec); err != nil {
			return "", err
		}
	}
	params := clientv1.NewV1TenantUIDDatasinksSplunkCreateParamsWithContext(h.ctx).
		WithTenantUID(tenantUID).
		WithBody(entity)
	resp, err := h.Client.V1TenantUIDDatasinksSplunkCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

// GetSplunkAuditTrail retrieves a Splunk audit trail by UID.
func (h *V1Client) GetSplunkAuditTrail(tenantUID, uid string) (*models.V1SplunkSink, error) {
	params := clientv1.NewV1TenantUIDDatasinksSplunkGetParamsWithContext(h.ctx).
		WithTenantUID(tenantUID).
		WithUID(uid)
	resp, err := h.Client.V1TenantUIDDatasinksSplunkGet(params)
	if apiutil.Is404(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// UpdateSplunkAuditTrail updates a Splunk audit trail.
func (h *V1Client) UpdateSplunkAuditTrail(tenantUID, uid string, entity *models.V1SplunkSinkEntity) error {
	params := clientv1.NewV1TenantUIDDatasinksSplunkUpdateParamsWithContext(h.ctx).
		WithTenantUID(tenantUID).
		WithUID(uid).
		WithBody(entity)
	_, err := h.Client.V1TenantUIDDatasinksSplunkUpdate(params)
	return err
}

// DeleteSplunkAuditTrail deletes a Splunk audit trail by UID.
func (h *V1Client) DeleteSplunkAuditTrail(tenantUID, uid string) error {
	params := clientv1.NewV1TenantUIDDatasinksSplunkDeleteParamsWithContext(h.ctx).
		WithTenantUID(tenantUID).
		WithUID(uid)
	_, err := h.Client.V1TenantUIDDatasinksSplunkDelete(params)
	return err
}
