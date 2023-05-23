package client

import (
	"fmt"

	authC "github.com/spectrocloud/hapi/auth/client/v1"
	"github.com/spectrocloud/palette-sdk-go/client/herr"
)

func (h *V1Client) PrintLoginBanner(orgName string) error {
	params := authC.NewV1AuthOrgLoginBannerGetParams().WithOrgName(orgName)
	loginBanner, err := AuthClient.V1AuthOrgLoginBannerGet(params)
	if err != nil || loginBanner == nil {
		if herr.IsNotFound(err) {
			return fmt.Errorf("invalid Organization: %s", orgName)
		}
		return err
	} else if !loginBanner.Payload.IsEnabled {
		return nil
	}

	fmt.Println(loginBanner.Payload.Message)
	return nil
}
