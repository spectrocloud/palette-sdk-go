package client

import (
	"fmt"

	clientv1 "github.com/spectrocloud/palette-sdk-go/api/client/version1"
	"github.com/spectrocloud/palette-sdk-go/api/models"
	"github.com/spectrocloud/palette-sdk-go/client/apiutil"
)

// CreateClusterConfigTemplate creates a new cluster config template.
func (h *V1Client) CreateClusterConfigTemplate(body *models.V1ClusterTemplateEntity) (*models.V1UID, error) {
	params := clientv1.NewV1ClusterTemplatesCreateParamsWithContext(h.ctx).
		WithBody(body)
	resp, err := h.Client.V1ClusterTemplatesCreate(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// UpdateClusterConfigTemplate updates an existing cluster config template metadata.
func (h *V1Client) UpdateClusterConfigTemplate(uid string, body *models.V1ObjectMetaInputEntitySchema) error {
	params := clientv1.NewV1ClusterTemplatesUIDMetadataUpdateParamsWithContext(h.ctx).
		WithUID(uid).
		WithBody(body)
	_, err := h.Client.V1ClusterTemplatesUIDMetadataUpdate(params)
	return err
}

// GetClusterConfigTemplate retrieves an existing cluster config template by UID.
func (h *V1Client) GetClusterConfigTemplate(uid string) (*models.V1ClusterTemplate, error) {
	params := clientv1.NewV1ClusterTemplatesUIDGetParamsWithContext(h.ctx).
		WithUID(uid)
	resp, err := h.Client.V1ClusterTemplatesUIDGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// ListClusterConfigTemplates retrieves all cluster config templates.
func (h *V1Client) ListClusterConfigTemplates() (*models.V1ClusterTemplatesSummary, error) {
	params := clientv1.NewV1ClusterTemplatesFilterSummaryParamsWithContext(h.ctx).
		WithLimit(apiutil.Ptr(int64(0)))
	resp, err := h.Client.V1ClusterTemplatesFilterSummary(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// GetClusterConfigTemplateByName retrieves an existing cluster config template by name.
func (h *V1Client) GetClusterConfigTemplateByName(name string) (*models.V1ClusterTemplateSummary, error) {
	templates, err := h.ListClusterConfigTemplates()
	if err != nil {
		return nil, err
	}
	for _, template := range templates.Items {
		if template.Metadata.Name == name {
			return template, nil
		}
	}
	return nil, fmt.Errorf("cluster config template not found for name %s", name)
}

// DeleteClusterConfigTemplate deletes an existing cluster config template.
func (h *V1Client) DeleteClusterConfigTemplate(uid string) error {
	params := clientv1.NewV1ClusterTemplatesUIDDeleteParamsWithContext(h.ctx).
		WithUID(uid)
	_, err := h.Client.V1ClusterTemplatesUIDDelete(params)
	return err
}

// UpdateClusterConfigTemplatePolicies updates the policies for a cluster config template.
func (h *V1Client) UpdateClusterConfigTemplatePolicies(uid string, body *models.V1ClusterTemplatePoliciesUpdateEntity) error {
	params := clientv1.NewV1ClusterTemplatesUIDPoliciesUpdateParamsWithContext(h.ctx).
		WithUID(uid).
		WithBody(body)
	_, err := h.Client.V1ClusterTemplatesUIDPoliciesUpdate(params)
	return err
}

// UpgradeClusterConfigTemplateClusters triggers an upgrade for all clusters launched from this template.
func (h *V1Client) UpgradeClusterConfigTemplateClusters(uid string) error {
	params := clientv1.NewV1SpectroClustersTemplatesUIDClustersUpgradeParamsWithContext(h.ctx).
		WithUID(uid)
	_, err := h.Client.V1SpectroClustersTemplatesUIDClustersUpgrade(params)
	return err
}

// UpdateClusterConfigTemplateProfiles updates the profiles list for a cluster config template.
// Use this when adding/removing profiles or changing profile UIDs.
func (h *V1Client) UpdateClusterConfigTemplateProfiles(uid string, body *models.V1ClusterTemplateProfilesUpdateEntity) error {
	params := clientv1.NewV1ClusterTemplatesUIDProfilesUpdateParamsWithContext(h.ctx).
		WithUID(uid).
		WithBody(body)
	_, err := h.Client.V1ClusterTemplatesUIDProfilesUpdate(params)
	return err
}

// UpdateClusterConfigTemplateProfilesVariables updates profile variables for a cluster config template.
// Use this when only updating variable values within existing profiles.
func (h *V1Client) UpdateClusterConfigTemplateProfilesVariables(uid string, body *models.V1ClusterTemplateProfilesVariablesBatchEntity) error {
	params := clientv1.NewV1ClusterTemplatesUIDProfilesVariablesPatchParamsWithContext(h.ctx).
		WithUID(uid).
		WithBody(body)
	_, err := h.Client.V1ClusterTemplatesUIDProfilesVariablesPatch(params)
	return err
}
