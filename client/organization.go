package client

import (
	"fmt"

	authC "github.com/spectrocloud/hapi/auth/client/v1"
	"github.com/spectrocloud/hapi/models"
	"github.com/spectrocloud/palette-sdk-go/client/herr"
)

func (h *V1Client) GetOrganizationByName(name string) (*models.V1LoginResponse, error) {
	client, err := h.GetAuthClient()
	if err != nil {
		return nil, err
	}

	params := authC.NewV1AuthOrgParams().WithOrgName(&name)
	resp, err := client.V1AuthOrg(params)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (h *V1Client) ListOrganizations(scope string) ([]*models.V1Organization, error) {
	client, err := h.GetAuthClient()
	if err != nil {
		return nil, err
	}

	var params *authC.V1AuthOrgsParams
	switch scope {
	case "project":
		params = authC.NewV1AuthOrgsParams().WithContext(h.Ctx)
	case "tenant":
		params = authC.NewV1AuthOrgsParams()
	}

	resp, err := client.V1AuthOrgs(params)
	if err != nil || resp == nil {
		return nil, err
	} else if len(resp.Payload.Organizations) == 0 {
		return nil, fmt.Errorf("no Organizations found in scope: %s", scope)
	}

	return resp.Payload.Organizations, nil
}

func (h *V1Client) SwitchOrganization(scope, orgName string) (string, error) {
	client, err := h.GetAuthClient()
	if err != nil {
		return "", err
	}

	var params *authC.V1AuthOrgSwitchParams
	switch scope {
	case "project":
		params = authC.NewV1AuthOrgSwitchParams().WithContext(h.Ctx).WithOrgName(orgName)
	case "tenant":
		params = authC.NewV1AuthOrgSwitchParams().WithOrgName(orgName)
	}

	resp, err := client.V1AuthOrgSwitch(params)
	if err != nil || resp == nil {
		return "", err
	}

	return resp.Payload.Authorization, nil
}

func (h *V1Client) PrintLoginBanner(orgName string) (bool, error) {
	client, err := h.GetNoAuthClient()
	if err != nil {
		return false, err
	}

	params := authC.NewV1AuthOrgLoginBannerGetParams().WithOrgName(orgName)
	resp, err := client.V1AuthOrgLoginBannerGet(params)
	if err != nil || resp == nil {
		if herr.IsNotFound(err) {
			return false, fmt.Errorf("invalid Organization: %s", orgName)
		}
		return false, err
	} else if !resp.Payload.IsEnabled {
		return false, nil
	}

	fmt.Println(resp.Payload.Message)
	return true, nil
}
