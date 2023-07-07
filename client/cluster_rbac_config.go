package client

import (
	"github.com/spectrocloud/hapi/models"
	clusterC "github.com/spectrocloud/hapi/spectrocluster/client/v1"

	"github.com/spectrocloud/palette-sdk-go/client/herr"
)

func (h *V1Client) GetClusterRbacConfig(uid, ClusterContext string) (*models.V1ClusterRbacs, error) {
	if h.GetClusterRbacConfigFn != nil {
		return h.GetClusterRbacConfigFn(uid)
	}
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}

	var params *clusterC.V1SpectroClustersUIDConfigRbacsGetParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1SpectroClustersUIDConfigRbacsGetParamsWithContext(h.Ctx).WithUID(uid)
	case "tenant":
		params = clusterC.NewV1SpectroClustersUIDConfigRbacsGetParams().WithUID(uid)
	}

	success, err := client.V1SpectroClustersUIDConfigRbacsGet(params)
	if err != nil {
		if herr.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}

	return success.Payload, nil
}

func (h *V1Client) CreateClusterRbacConfig(uid string, config *models.V1ClusterRbac, ClusterContext string) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return err
	}

	var params *clusterC.V1WorkspacesClusterRbacCreateParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1WorkspacesClusterRbacCreateParamsWithContext(h.Ctx).WithUID(uid).WithBody(config)
	case "tenant":
		params = clusterC.NewV1WorkspacesClusterRbacCreateParams().WithUID(uid).WithBody(config)
	}

	_, err = client.V1WorkspacesClusterRbacCreate(params)
	return err
}

func (h *V1Client) UpdateClusterRbacConfig(uid string, config *models.V1ClusterRbacResourcesUpdateEntity, ClusterContext string) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return err
	}

	var params *clusterC.V1SpectroClustersUIDConfigRbacsUpdateParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1SpectroClustersUIDConfigRbacsUpdateParamsWithContext(h.Ctx).WithUID(uid).WithBody(config)
	case "tenant":
		params = clusterC.NewV1SpectroClustersUIDConfigRbacsUpdateParams().WithUID(uid).WithBody(config)
	}

	_, err = client.V1SpectroClustersUIDConfigRbacsUpdate(params)
	return err
}

func (h *V1Client) ApplyClusterRbacConfig(uid string, config []*models.V1ClusterRbacInputEntity, ClusterContext string) error {
	if rbac, err := h.GetClusterRbacConfig(uid, ClusterContext); err != nil {
		return err
	} else if rbac == nil {
		return h.CreateClusterRbacConfig(uid, toCreateClusterRbac(config), ClusterContext)
	} else {
		return h.UpdateClusterRbacConfig(uid, &models.V1ClusterRbacResourcesUpdateEntity{
			Rbacs: config,
		}, ClusterContext)
	}
}

func toCreateClusterRbac(rbacs []*models.V1ClusterRbacInputEntity) *models.V1ClusterRbac {
	bindings := make([]*models.V1ClusterRbacBinding, 0)

	for _, rbac := range rbacs {
		bindings = append(bindings, rbac.Spec.Bindings...)
	}

	return &models.V1ClusterRbac{
		Spec: &models.V1ClusterRbacSpec{
			Bindings: bindings,
		},
	}
}
