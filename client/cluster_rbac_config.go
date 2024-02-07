package client

import (
	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
	"github.com/spectrocloud/palette-sdk-go/client/herr"
)

func (h *V1Client) GetClusterRbacConfig(uid string) (*models.V1ClusterRbacs, error) {
	params := clientV1.NewV1SpectroClustersUIDConfigRbacsGetParamsWithContext(h.ctx).
		WithUID(uid)
	resp, err := h.Client.V1SpectroClustersUIDConfigRbacsGet(params)
	if err != nil {
		if herr.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}
	return resp.Payload, nil
}

func (h *V1Client) CreateClusterRbacConfig(uid string, config *models.V1ClusterRbac) error {
	params := clientV1.NewV1WorkspacesClusterRbacCreateParamsWithContext(h.ctx).
		WithUID(uid).
		WithBody(config)
	_, err := h.Client.V1WorkspacesClusterRbacCreate(params)
	return err
}

func (h *V1Client) UpdateClusterRbacConfig(uid string, config *models.V1ClusterRbacResourcesUpdateEntity) error {
	params := clientV1.NewV1SpectroClustersUIDConfigRbacsUpdateParamsWithContext(h.ctx).
		WithUID(uid).
		WithBody(config)
	_, err := h.Client.V1SpectroClustersUIDConfigRbacsUpdate(params)
	return err
}

func (h *V1Client) ApplyClusterRbacConfig(uid string, config []*models.V1ClusterRbacInputEntity) error {
	rbac, err := h.GetClusterRbacConfig(uid)
	if err != nil {
		return err
	}
	if rbac == nil {
		return h.CreateClusterRbacConfig(uid, toCreateClusterRbac(config))
	}
	return h.UpdateClusterRbacConfig(uid, &models.V1ClusterRbacResourcesUpdateEntity{Rbacs: config})
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
