package client

import (
	"fmt"

	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
	"github.com/spectrocloud/palette-sdk-go/client/herr"
)

func (h *V1Client) GetOrganizationByName(name string) (*models.V1LoginResponse, error) {
	params := clientV1.NewV1AuthOrgParams().WithOrgName(&name)
	resp, err := h.GetClient().V1AuthOrg(params)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (h *V1Client) ListOrganizations(scope string) ([]*models.V1Organization, error) {
	var params *clientV1.V1AuthOrgsParams
	switch scope {
	case "project":
		params = clientV1.NewV1AuthOrgsParams().WithContext(h.Ctx)
	case "tenant":
		params = clientV1.NewV1AuthOrgsParams()
	}

	resp, err := h.GetClient().V1AuthOrgs(params)
	if err != nil || resp == nil {
		return nil, err
	} else if len(resp.Payload.Organizations) == 0 {
		return nil, fmt.Errorf("no Organizations found in scope: %s", scope)
	}

	return resp.Payload.Organizations, nil
}

func (h *V1Client) SwitchOrganization(scope, orgName string) (string, error) {
	var params *clientV1.V1AuthOrgSwitchParams
	switch scope {
	case "project":
		params = clientV1.NewV1AuthOrgSwitchParams().WithContext(h.Ctx).WithOrgName(orgName)
	case "tenant":
		params = clientV1.NewV1AuthOrgSwitchParams().WithOrgName(orgName)
	}

	resp, err := h.GetClient().V1AuthOrgSwitch(params)
	if err != nil || resp == nil {
		return "", err
	}

	return resp.Payload.Authorization, nil
}

func (h *V1Client) GetOrgLoginBanner(orgName string) (*models.V1LoginBannerSettings, error) {
	params := clientV1.NewV1AuthOrgLoginBannerGetParams().WithOrgName(orgName)
	resp, err := h.GetClient().V1AuthOrgLoginBannerGet(params)
	if err != nil || resp == nil {
		if herr.IsNotFound(err) {
			return nil, fmt.Errorf("invalid Organization: %s", orgName)
		}
		return nil, err
	} else if !resp.Payload.IsEnabled {
		return nil, nil
	}

	return resp.Payload, nil
}

func (h *V1Client) GetSystemLoginBanner() (*models.V1LoginBannerSettings, error) {
	params := clientV1.NewV1AuthSystemLoginBannerGetParams()
	resp, err := h.GetClient().V1AuthSystemLoginBannerGet(params)
	if err != nil || resp == nil {
		return nil, err
	} else if !resp.Payload.IsEnabled {
		return nil, nil
	}

	return resp.Payload, nil
}
