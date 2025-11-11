package client

import (
	"fmt"

	clientv1 "github.com/spectrocloud/palette-sdk-go/api/client/version1"
	"github.com/spectrocloud/palette-sdk-go/api/models"
	"github.com/spectrocloud/palette-sdk-go/client/apiutil"
)

// CreateClusterConfigPolicy creates a new cluster config policy (maintenance policy).
func (h *V1Client) CreateClusterConfigPolicy(body *models.V1SpcPolicyEntity) (*models.V1UID, error) {
	params := clientv1.NewV1SpcPoliciesMaintenanceCreateParamsWithContext(h.ctx).
		WithBody(body)
	resp, err := h.Client.V1SpcPoliciesMaintenanceCreate(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// UpdateClusterConfigPolicy updates an existing cluster config policy (maintenance policy).
func (h *V1Client) UpdateClusterConfigPolicy(uid string, body *models.V1SpcPolicyEntity) error {
	params := clientv1.NewV1SpcPoliciesMaintenanceUIDUpdateParamsWithContext(h.ctx).
		WithUID(uid).
		WithBody(body)
	_, err := h.Client.V1SpcPoliciesMaintenanceUIDUpdate(params)
	return err
}

// GetClusterConfigPolicy retrieves an existing cluster config policy by UID.
func (h *V1Client) GetClusterConfigPolicy(uid string) (*models.V1SpcPolicyEntity, error) {
	params := clientv1.NewV1SpcPoliciesMaintenanceUIDGetParamsWithContext(h.ctx).
		WithUID(uid)
	resp, err := h.Client.V1SpcPoliciesMaintenanceUIDGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// ListClusterConfigPolicies retrieves all cluster config policies.
func (h *V1Client) ListClusterConfigPolicies() (*models.V1SpcPoliciesSummary, error) {
	params := clientv1.NewV1SpcPoliciesFilterSummaryParamsWithContext(h.ctx).
		WithLimit(apiutil.Ptr(int64(0)))
	resp, err := h.Client.V1SpcPoliciesFilterSummary(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// GetClusterConfigPolicyByName retrieves an existing cluster config policy by name.
func (h *V1Client) GetClusterConfigPolicyByName(name string) (*models.V1SpcPolicySummary, error) {
	policies, err := h.ListClusterConfigPolicies()
	if err != nil {
		return nil, err
	}
	for _, policy := range policies.Items {
		if policy.Metadata.Name == name {
			return policy, nil
		}
	}
	return nil, fmt.Errorf("cluster config policy not found for name %s", name)
}

// DeleteClusterConfigPolicy deletes an existing cluster config policy.
func (h *V1Client) DeleteClusterConfigPolicy(uid string) error {
	params := clientv1.NewV1SpcPoliciesUIDDeleteParamsWithContext(h.ctx).
		WithUID(uid)
	_, err := h.Client.V1SpcPoliciesUIDDelete(params)
	return err
}
