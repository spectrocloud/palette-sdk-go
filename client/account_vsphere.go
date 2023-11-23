package client

import (
	"errors"

	"github.com/spectrocloud/hapi/apiutil/transport"
	"github.com/spectrocloud/hapi/models"
	clusterC "github.com/spectrocloud/hapi/spectrocluster/client/v1"
)

const OverlordUID = "overlordUid"

// convert V1VsphereAccount to V1OverlordVsphereAccountEntity
func toV1OverlordsUIDVsphereAccountValidateBody(account *models.V1VsphereAccount) clusterC.V1OverlordsUIDVsphereAccountValidateBody {
	return clusterC.V1OverlordsUIDVsphereAccountValidateBody{
		Account: account.Spec,
	}
}

func (h *V1Client) CreateCloudAccountVsphere(account *models.V1VsphereAccount, AccountContext string) (string, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return "", err
	}

	// validate account
	err = validateCloudAccountVsphere(account, h)
	if err != nil {
		return "", err
	}

	// create account
	var params *clusterC.V1CloudAccountsVsphereCreateParams
	switch AccountContext {
	case "project":
		params = clusterC.NewV1CloudAccountsVsphereCreateParamsWithContext(h.Ctx)
	case "tenant":
		params = clusterC.NewV1CloudAccountsVsphereCreateParams()
	}

	params = params.WithBody(account)
	success, err := client.V1CloudAccountsVsphereCreate(params)
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

func validateCloudAccountVsphere(account *models.V1VsphereAccount, h *V1Client) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return err
	}

	PcgId := account.Metadata.Annotations[OverlordUID]
	// check PCG
	err = h.CheckPCG(PcgId)
	if err != nil {
		return err
	}

	// validate account
	paramsValidate := clusterC.NewV1OverlordsUIDVsphereAccountValidateParams().WithUID(PcgId)
	paramsValidate = paramsValidate.WithBody(toV1OverlordsUIDVsphereAccountValidateBody(account))
	_, err = client.V1OverlordsUIDVsphereAccountValidate(paramsValidate)
	if err != nil {
		return err
	}

	return nil
}

func (h *V1Client) UpdateCloudAccountVsphere(account *models.V1VsphereAccount, AccountContext string) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return err
	}

	// validate account
	err = validateCloudAccountVsphere(account, h)
	if err != nil {
		return err
	}

	uid := account.Metadata.UID
	params := clusterC.NewV1CloudAccountsVsphereUpdateParamsWithContext(h.Ctx).WithUID(uid).WithBody(account)
	_, err = client.V1CloudAccountsVsphereUpdate(params)
	return err
}

func (h *V1Client) DeleteCloudAccountVsphere(uid, AccountContext string) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return err
	}

	var params *clusterC.V1CloudAccountsVsphereDeleteParams
	switch AccountContext {
	case "project":
		params = clusterC.NewV1CloudAccountsVsphereDeleteParamsWithContext(h.Ctx).WithUID(uid)
	case "tenant":
		params = clusterC.NewV1CloudAccountsVsphereDeleteParams().WithUID(uid)
	}

	_, err = client.V1CloudAccountsVsphereDelete(params)
	return err
}

func (h *V1Client) GetCloudAccountVsphere(uid, AccountContext string) (*models.V1VsphereAccount, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}

	var params *clusterC.V1CloudAccountsVsphereGetParams
	switch AccountContext {
	case "project":
		params = clusterC.NewV1CloudAccountsVsphereGetParamsWithContext(h.Ctx).WithUID(uid)
	case "tenant":
		params = clusterC.NewV1CloudAccountsVsphereGetParams().WithUID(uid)
	}

	success, err := client.V1CloudAccountsVsphereGet(params)

	var e *transport.TransportError
	if errors.As(err, &e) && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return success.Payload, nil
}

func (h *V1Client) GetCloudAccountsVsphere() ([]*models.V1VsphereAccount, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}

	limit := int64(0)
	params := clusterC.NewV1CloudAccountsVsphereListParamsWithContext(h.Ctx).WithLimit(&limit)
	response, err := client.V1CloudAccountsVsphereList(params)
	if err != nil {
		return nil, err
	}

	accounts := make([]*models.V1VsphereAccount, len(response.Payload.Items))
	copy(accounts, response.Payload.Items)

	return accounts, nil
}
