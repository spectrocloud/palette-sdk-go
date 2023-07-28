package client

import (
	"github.com/spectrocloud/hapi/apiutil/transport"
	"github.com/spectrocloud/hapi/models"
	clusterC "github.com/spectrocloud/hapi/spectrocluster/client/v1"
)

func toV1OverlordsUIDMaasAccountValidateBody(account *models.V1MaasAccount) clusterC.V1OverlordsUIDMaasAccountValidateBody {
	return clusterC.V1OverlordsUIDMaasAccountValidateBody{
		Account: account.Spec,
	}
}

func (h *V1Client) CreateCloudAccountMaas(account *models.V1MaasAccount, AccountContext string) (string, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return "", err
	}

	// validate account
	err = validateCloudAccountMaas(account, h)
	if err != nil {
		return "", err
	}

	var params *clusterC.V1CloudAccountsMaasCreateParams
	switch AccountContext {
	case "project":
		params = clusterC.NewV1CloudAccountsMaasCreateParamsWithContext(h.Ctx).WithBody(account)
	case "tenant":
		params = clusterC.NewV1CloudAccountsMaasCreateParams().WithBody(account)
	}

	success, err := client.V1CloudAccountsMaasCreate(params)
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

func validateCloudAccountMaas(account *models.V1MaasAccount, h *V1Client) error {
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
	paramsValidate := clusterC.NewV1OverlordsUIDMaasAccountValidateParams().WithUID(PcgId)
	paramsValidate = paramsValidate.WithBody(toV1OverlordsUIDMaasAccountValidateBody(account))
	_, err = client.V1OverlordsUIDMaasAccountValidate(paramsValidate)
	if err != nil {
		return err
	}

	return nil
}

func (h *V1Client) UpdateCloudAccountMaas(account *models.V1MaasAccount) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return err
	}

	// validate account
	err = validateCloudAccountMaas(account, h)
	if err != nil {
		return err
	}

	uid := account.Metadata.UID
	params := clusterC.NewV1CloudAccountsMaasUpdateParamsWithContext(h.Ctx).WithUID(uid).WithBody(account)
	_, err = client.V1CloudAccountsMaasUpdate(params)
	return err
}

func (h *V1Client) DeleteCloudAccountMaas(uid, AccountContext string) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil
	}

	var params *clusterC.V1CloudAccountsMaasDeleteParams
	switch AccountContext {
	case "project":
		params = clusterC.NewV1CloudAccountsMaasDeleteParamsWithContext(h.Ctx).WithUID(uid)
	case "tenant":
		params = clusterC.NewV1CloudAccountsMaasDeleteParams().WithUID(uid)
	}

	_, err = client.V1CloudAccountsMaasDelete(params)
	return err
}

func (h *V1Client) GetCloudAccountMaas(uid, AccountContext string) (*models.V1MaasAccount, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}

	var params *clusterC.V1CloudAccountsMaasGetParams
	switch AccountContext {
	case "project":
		params = clusterC.NewV1CloudAccountsMaasGetParamsWithContext(h.Ctx).WithUID(uid)
	case "tenant":
		params = clusterC.NewV1CloudAccountsMaasGetParams().WithUID(uid)
	}

	success, err := client.V1CloudAccountsMaasGet(params)
	if e, ok := err.(*transport.TransportError); ok && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return success.Payload, nil
}

func (h *V1Client) GetCloudAccountsMaas() ([]*models.V1MaasAccount, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}

	limit := int64(0)
	params := clusterC.NewV1CloudAccountsMaasListParamsWithContext(h.Ctx).WithLimit(&limit)
	response, err := client.V1CloudAccountsMaasList(params)
	if err != nil {
		return nil, err
	}

	accounts := make([]*models.V1MaasAccount, len(response.Payload.Items))
	copy(accounts, response.Payload.Items)

	return accounts, nil
}
