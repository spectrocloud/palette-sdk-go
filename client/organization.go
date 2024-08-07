package client

import (
	clientv1 "github.com/spectrocloud/palette-sdk-go/api/client/v1"
	"github.com/spectrocloud/palette-sdk-go/api/models"
)

// GetOrganizationByName retrieves an organization by name.
func (h *V1Client) GetOrganizationByName(name string) (*models.V1LoginResponse, error) {
	params := clientv1.NewV1AuthOrgParamsWithContext(h.ctx).
		WithOrgName(&name)
	resp, err := h.Client.V1AuthOrg(params)
	if err != nil || resp == nil {
		return nil, err
	}
	return resp.Payload, nil
}

// ListOrganizations retrieves all organizations.
func (h *V1Client) ListOrganizations() ([]*models.V1Organization, error) {
	params := clientv1.NewV1AuthOrgsParamsWithContext(h.ctx)
	resp, err := h.Client.V1AuthOrgs(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload.Organizations, nil
}
