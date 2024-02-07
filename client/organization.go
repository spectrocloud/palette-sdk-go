package client

import (
	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
)

func (h *V1Client) GetOrganizationByName(name string) (*models.V1LoginResponse, error) {
	params := clientV1.NewV1AuthOrgParamsWithContext(h.ctx).
		WithOrgName(&name)
	resp, err := h.Client.V1AuthOrg(params)
	if err != nil || resp == nil {
		return nil, err
	}
	return resp.Payload, nil
}

func (h *V1Client) ListOrganizations() ([]*models.V1Organization, error) {
	params := clientV1.NewV1AuthOrgsParamsWithContext(h.ctx)
	resp, err := h.Client.V1AuthOrgs(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload.Organizations, nil
}

func (h *V1Client) SwitchOrganization(scope, orgName string) (string, error) {
	params := clientV1.NewV1AuthOrgSwitchParamsWithContext(h.ctx).
		WithOrgName(orgName)
	resp, err := h.Client.V1AuthOrgSwitch(params)
	if err != nil || resp == nil {
		return "", err
	}
	return resp.Payload.Authorization, nil
}
