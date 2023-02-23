package client

import (
	"github.com/spectrocloud/hapi/apiutil/transport"
	"github.com/spectrocloud/hapi/models"
	v1 "github.com/spectrocloud/hapi/spectrocluster/client/v1"
)

func (h *V1Client) GetApplication(uid string) (*models.V1AppDeployment, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}

	params := v1.NewV1AppDeploymentsUIDGetParamsWithContext(h.Ctx).WithUID(uid)
	success, err := client.V1AppDeploymentsUIDGet(params)
	if e, ok := err.(*transport.TransportError); ok && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	// special check if the cluster is marked deleted
	application := success.Payload //success.Payload.Spec.Config.Target.ClusterRef.UID
	return application, nil
}

func (h *V1Client) CreateApplicationWithNewSandboxCluster(body *models.V1AppDeploymentClusterGroupEntity) (string, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return "", err
	}

	params := v1.NewV1AppDeploymentsClusterGroupCreateParams().WithContext(h.Ctx).WithBody(body)
	success, err := client.V1AppDeploymentsClusterGroupCreate(params)

	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

func (h *V1Client) CreateApplicationWithExistingSandboxCluster(body *models.V1AppDeploymentVirtualClusterEntity) (string, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return "", err
	}

	params := v1.NewV1AppDeploymentsVirtualClusterCreateParams().WithContext(h.Ctx).WithBody(body)
	success, err := client.V1AppDeploymentsVirtualClusterCreate(params)

	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

func (h *V1Client) DeleteApplication(uid string) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil
	}
	params := v1.NewV1AppDeploymentsUIDDeleteParamsWithContext(h.Ctx).WithUID(uid)
	_, err = client.V1AppDeploymentsUIDDelete(params)
	return err
}
