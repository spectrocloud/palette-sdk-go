package client

import (
	"fmt"
	cloudC "github.com/spectrocloud/hapi/cloud/client/v1"
	"github.com/spectrocloud/hapi/models"
	clusterC "github.com/spectrocloud/hapi/spectrocluster/client/v1"
)

func (h *V1Client) CreateCustomCloudAccount(account *models.V1CustomAccountEntity, cloudType string, accountContext string) (string, error) {
	if h.CreateCustomCloudAccountFn != nil {
		return h.CreateCustomCloudAccountFn(account, cloudType, accountContext)
	}
	var params *clusterC.V1CloudAccountsCustomCreateParams

	// check PCG
	PcgId := account.Metadata.Annotations[OverlordUID]
	if err := h.CheckPCG(PcgId); err != nil {
		return "", err
	}
	switch accountContext {
	case "project":
		params = clusterC.NewV1CloudAccountsCustomCreateParamsWithContext(h.Ctx)
	case "tenant":
		params = clusterC.NewV1CloudAccountsCustomCreateParams()
	}
	params = params.WithBody(account).WithCloudType(cloudType)
	success, err := h.GetClusterClient().V1CloudAccountsCustomCreate(params)
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

func (h *V1Client) GetCustomCloudAccount(uid, cloudType string, accountContext string) (*models.V1CustomAccount, error) {
	if h.GetCustomCloudAccountFn != nil {
		return h.GetCustomCloudAccountFn(uid, cloudType, accountContext)
	}
	var params *clusterC.V1CloudAccountsCustomGetParams

	switch accountContext {
	case "project":
		params = clusterC.NewV1CloudAccountsCustomGetParamsWithContext(h.Ctx)
	case "tenant":
		params = clusterC.NewV1CloudAccountsCustomGetParams()
	}
	params = params.WithCloudType(cloudType).WithUID(uid)
	success, err := h.GetClusterClient().V1CloudAccountsCustomGet(params)
	if err != nil {
		return nil, err
	}
	return success.Payload, nil
}

func (h *V1Client) UpdateCustomCloudAccount(uid string, account *models.V1CustomAccountEntity, cloudType string, accountContext string) error {
	if h.UpdateCustomCloudAccountFn != nil {
		return h.UpdateCustomCloudAccountFn(uid, account, cloudType, accountContext)
	}
	var params *clusterC.V1CloudAccountsCustomUpdateParams

	switch accountContext {
	case "project":
		params = clusterC.NewV1CloudAccountsCustomUpdateParamsWithContext(h.Ctx)
	case "tenant":
		params = clusterC.NewV1CloudAccountsCustomUpdateParams()
	}
	params = params.WithBody(account).WithCloudType(cloudType).WithUID(uid)
	_, err := h.GetClusterClient().V1CloudAccountsCustomUpdate(params)
	if err != nil {
		return err
	}

	return nil
}

func (h *V1Client) DeleteCustomCloudAccount(uid, cloudType string, accountContext string) error {
	if h.DeleteCustomCloudAccountFn != nil {
		return h.DeleteCustomCloudAccountFn(uid, cloudType, accountContext)
	}
	var params *clusterC.V1CloudAccountsCustomDeleteParams

	switch accountContext {
	case "project":
		params = clusterC.NewV1CloudAccountsCustomDeleteParamsWithContext(h.Ctx)
	case "tenant":
		params = clusterC.NewV1CloudAccountsCustomDeleteParams()
	}
	params = params.WithCloudType(cloudType).WithUID(uid)
	_, err := h.GetClusterClient().V1CloudAccountsCustomDelete(params)
	return err
}

func (h *V1Client) ValidateCustomCloudType(cloudType string, cloudContext string) error {
	if h.ValidateCustomCloudTypeFn != nil {
		return h.ValidateCustomCloudTypeFn(cloudType, cloudContext)
	}
	var params *cloudC.V1CustomCloudTypesGetParams
	switch cloudContext {
	case "project":
		params = cloudC.NewV1CustomCloudTypesGetParamsWithContext(h.Ctx)
	case "tenant":
		params = cloudC.NewV1CustomCloudTypesGetParams()
	}
	success, err := h.GetCloudClient().V1CustomCloudTypesGet(params)
	if err != nil {
		return err
	}
	customCloudTypes := success.GetPayload().CloudTypes
	for _, c := range customCloudTypes {
		if c.Name == cloudType {
			if c.IsCustom {
				return nil
			} else {
				return fmt.Errorf("cloud - `%s` is not a valid custom cloud", cloudType)
			}
		}
	}
	return fmt.Errorf("cloud - `%s` is not a valid cloud", cloudType)
}
