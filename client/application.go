package client

import (
	clientv1 "github.com/spectrocloud/palette-sdk-go/api/client/version1"
	"github.com/spectrocloud/palette-sdk-go/api/models"
	"github.com/spectrocloud/palette-sdk-go/client/apiutil"
)

// GetApplication retrieves an existing application deployment.
func (h *V1Client) GetApplication(uid string) (*models.V1AppDeployment, error) {
	params := clientv1.NewV1AppDeploymentsUIDGetParamsWithContext(h.ctx).
		WithUID(uid)
	resp, err := h.Client.V1AppDeploymentsUIDGet(params)
	if apiutil.Is404(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// SearchAppDeploymentSummaries retrieves a list of application deployment summaries based on the filter and sort criteria.
func (h *V1Client) SearchAppDeploymentSummaries(filter *models.V1AppDeploymentFilterSpec, sortBy []*models.V1AppDeploymentSortSpec) ([]*models.V1AppDeploymentSummary, error) {
	params := clientv1.NewV1DashboardAppDeploymentsParamsWithContext(h.ctx).
		WithBody(&models.V1AppDeploymentsFilterSpec{
			Filter: filter,
			Sort:   sortBy,
		})
	resp, err := h.Client.V1DashboardAppDeployments(params)
	if apiutil.Is404(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return resp.Payload.AppDeployments, nil
}

// CreateApplicationWithNewSandboxCluster creates a new application deployment inside a new virtual cluster.
func (h *V1Client) CreateApplicationWithNewSandboxCluster(body *models.V1AppDeploymentClusterGroupEntity) (string, error) {
	params := clientv1.NewV1AppDeploymentsClusterGroupCreateParamsWithContext(h.ctx).
		WithBody(body)
	resp, err := h.Client.V1AppDeploymentsClusterGroupCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

// CreateApplicationWithExistingSandboxCluster creates a new application deployment within an existing virtual cluster.
func (h *V1Client) CreateApplicationWithExistingSandboxCluster(body *models.V1AppDeploymentVirtualClusterEntity) (string, error) {
	params := clientv1.NewV1AppDeploymentsVirtualClusterCreateParamsWithContext(h.ctx).
		WithBody(body)
	resp, err := h.Client.V1AppDeploymentsVirtualClusterCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

// DeleteApplication deletes an existing application deployment.
func (h *V1Client) DeleteApplication(uid string) error {
	params := clientv1.NewV1AppDeploymentsUIDDeleteParamsWithContext(h.ctx).
		WithUID(uid)
	_, err := h.Client.V1AppDeploymentsUIDDelete(params)
	return err
}

// ApplyApplicationUpdate applies an update to an existing application deployment.
func (h *V1Client) ApplyApplicationUpdate(uid, notificationUID string) error {
	params := clientv1.NewV1AppDeploymentsUIDProfileApplyParamsWithContext(h.ctx).
		WithUID(uid).
		WithNotify(&notificationUID)
	_, err := h.Client.V1AppDeploymentsUIDProfileApply(params)
	return err
}
