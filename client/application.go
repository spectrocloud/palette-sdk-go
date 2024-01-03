package client

import (
	"errors"

	"github.com/spectrocloud/hapi/apiutil/transport"
	hashboardC "github.com/spectrocloud/hapi/hashboard/client/v1"
	"github.com/spectrocloud/hapi/models"
	v1 "github.com/spectrocloud/hapi/spectrocluster/client/v1"
)

func (h *V1Client) GetApplication(uid string) (*models.V1AppDeployment, error) {
	if h.GetApplicationFn != nil {
		return h.GetApplicationFn(uid)
	}

	params := v1.NewV1AppDeploymentsUIDGetParamsWithContext(h.Ctx).WithUID(uid)
	success, err := h.GetClusterClient().V1AppDeploymentsUIDGet(params)

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
	var params *hashboardC.V1DashboardAppDeploymentsParams
	switch scope {
	case "project":
		params = hashboardC.NewV1DashboardAppDeploymentsParams().WithContext(h.Ctx)
	case "tenant":
		params = hashboardC.NewV1DashboardAppDeploymentsParams()
	}
	params.Body = &models.V1AppDeploymentsFilterSpec{
		Filter: filter,
		Sort:   sortBy,
	}

	resp, err := h.GetHashboardClient().V1DashboardAppDeployments(params)
	var e *transport.TransportError
	if errors.As(err, &e) && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return resp.Payload.AppDeployments, nil
}

func (h *V1Client) CreateApplicationWithNewSandboxCluster(body *models.V1AppDeploymentClusterGroupEntity) (string, error) {
	params := v1.NewV1AppDeploymentsClusterGroupCreateParams().WithContext(h.Ctx).WithBody(body)
	success, err := h.GetClusterClient().V1AppDeploymentsClusterGroupCreate(params)
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

func (h *V1Client) CreateApplicationWithExistingSandboxCluster(body *models.V1AppDeploymentVirtualClusterEntity) (string, error) {
	params := v1.NewV1AppDeploymentsVirtualClusterCreateParams().WithContext(h.Ctx).WithBody(body)
	success, err := h.GetClusterClient().V1AppDeploymentsVirtualClusterCreate(params)
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

func (h *V1Client) DeleteApplication(uid string) error {
	params := v1.NewV1AppDeploymentsUIDDeleteParamsWithContext(h.Ctx).WithUID(uid)
	_, err := h.GetClusterClient().V1AppDeploymentsUIDDelete(params)
	return err
}

func (h *V1Client) ApplyApplicationUpdate(uid, notificationUid string) error {
	params := v1.NewV1AppDeploymentsUIDProfileApplyParamsWithContext(h.Ctx).WithUID(uid).WithNotify(&notificationUid)
	_, err := h.GetClusterClient().V1AppDeploymentsUIDProfileApply(params)
	return err
}
