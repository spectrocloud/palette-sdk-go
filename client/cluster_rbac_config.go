package client

import (
	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
	"github.com/spectrocloud/palette-sdk-go/client/herr"
)

func (h *V1Client) GetClusterRbacConfig(uid, scope string) (*models.V1ClusterRbacs, error) {
	var params *clientV1.V1SpectroClustersUIDConfigRbacsGetParams
	switch scope {
	case "project":
		params = clientV1.NewV1SpectroClustersUIDConfigRbacsGetParamsWithContext(h.Ctx).WithUID(uid)
	case "tenant":
		params = clientV1.NewV1SpectroClustersUIDConfigRbacsGetParams().WithUID(uid)
	}

	success, err := h.GetClient().V1SpectroClustersUIDConfigRbacsGet(params)
	if err != nil {
		if herr.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}

	return success.Payload, nil
}

func (h *V1Client) CreateClusterRbacConfig(uid, scope string, config *models.V1ClusterRbac) error {
	var params *clientV1.V1WorkspacesClusterRbacCreateParams
	switch scope {
	case "project":
		params = clientV1.NewV1WorkspacesClusterRbacCreateParamsWithContext(h.Ctx).WithUID(uid).WithBody(config)
	case "tenant":
		params = clientV1.NewV1WorkspacesClusterRbacCreateParams().WithUID(uid).WithBody(config)
	}

	_, err := h.GetClient().V1WorkspacesClusterRbacCreate(params)
	return err
}

func (h *V1Client) UpdateClusterRbacConfig(uid, scope string, config *models.V1ClusterRbacResourcesUpdateEntity) error {
	var params *clientV1.V1SpectroClustersUIDConfigRbacsUpdateParams
	switch scope {
	case "project":
		params = clientV1.NewV1SpectroClustersUIDConfigRbacsUpdateParamsWithContext(h.Ctx).WithUID(uid).WithBody(config)
	case "tenant":
		params = clientV1.NewV1SpectroClustersUIDConfigRbacsUpdateParams().WithUID(uid).WithBody(config)
	}

	_, err := h.GetClient().V1SpectroClustersUIDConfigRbacsUpdate(params)
	return err
}

func (h *V1Client) ApplyClusterRbacConfig(uid, scope string, config []*models.V1ClusterRbacInputEntity) error {
	if rbac, err := h.GetClusterRbacConfig(uid, scope); err != nil {
		return err
	} else if rbac == nil {
		return h.CreateClusterRbacConfig(uid, scope, toCreateClusterRbac(config))
	} else {
		return h.UpdateClusterRbacConfig(uid, scope, &models.V1ClusterRbacResourcesUpdateEntity{
			Rbacs: config,
		})
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
