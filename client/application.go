package client

import (
	"errors"

	"github.com/spectrocloud/palette-api-go/apiutil/transport"
	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
)

func (h *V1Client) GetApplication(uid string) (*models.V1AppDeployment, error) {
	params := clientV1.NewV1AppDeploymentsUIDGetParamsWithContext(h.Ctx).WithUID(uid)
	success, err := h.Client.V1AppDeploymentsUIDGet(params)

	var e *transport.TransportError
	if errors.As(err, &e) && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	// special check if the cluster is marked deleted
	application := success.Payload //success.Payload.Spec.Config.Target.ClusterRef.UID
	return application, nil
}

func (h *V1Client) SearchAppDeploymentSummaries(scope string, filter *models.V1AppDeploymentFilterSpec, sortBy []*models.V1AppDeploymentSortSpec) ([]*models.V1AppDeploymentSummary, error) {
	var params *clientV1.V1DashboardAppDeploymentsParams
	switch scope {
	case "project":
		params = clientV1.NewV1DashboardAppDeploymentsParams().WithContext(h.Ctx)
	case "tenant":
		params = clientV1.NewV1DashboardAppDeploymentsParams()
	}
	params.Body = &models.V1AppDeploymentsFilterSpec{
		Filter: filter,
		Sort:   sortBy,
	}

	resp, err := h.Client.V1DashboardAppDeployments(params)
	var e *transport.TransportError
	if errors.As(err, &e) && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return resp.Payload.AppDeployments, nil
}

func (h *V1Client) CreateApplicationWithNewSandboxCluster(body *models.V1AppDeploymentClusterGroupEntity) (string, error) {
	params := clientV1.NewV1AppDeploymentsClusterGroupCreateParams().WithContext(h.Ctx).WithBody(body)
	success, err := h.Client.V1AppDeploymentsClusterGroupCreate(params)
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

func (h *V1Client) CreateApplicationWithExistingSandboxCluster(body *models.V1AppDeploymentVirtualClusterEntity) (string, error) {
	params := clientV1.NewV1AppDeploymentsVirtualClusterCreateParams().WithContext(h.Ctx).WithBody(body)
	success, err := h.Client.V1AppDeploymentsVirtualClusterCreate(params)
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

func (h *V1Client) DeleteApplication(uid string) error {
	params := clientV1.NewV1AppDeploymentsUIDDeleteParamsWithContext(h.Ctx).WithUID(uid)
	_, err := h.Client.V1AppDeploymentsUIDDelete(params)
	return err
}

func (h *V1Client) ApplyApplicationUpdate(uid, notificationUid string) error {
	params := clientV1.NewV1AppDeploymentsUIDProfileApplyParamsWithContext(h.Ctx).WithUID(uid).WithNotify(&notificationUid)
	_, err := h.Client.V1AppDeploymentsUIDProfileApply(params)
	return err
}
