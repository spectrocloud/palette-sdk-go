package client

import (
	clientv1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
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

// SwitchOrganization switches the caller's active organization.
// Returns an authorization token for the target organization.
func (h *V1Client) SwitchOrganization(orgName string) (string, error) {

	// Need to update this fix , currently HAPI restricted `V1AuthOrgSwitch`

	//params:= clientv1.NewV1AuthOrgSwitchParamsWithContext(h.ctx).
	//	WithOrgName(orgName)
	//resp, err := h.Client.V1AuthOrgSwitch(params)
	//if err != nil || resp == nil {
	//	return "", err
	//}
	//return resp.Payload.Authorization, nil
	return "", nil
}
