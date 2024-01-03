package client

import (
	"errors"

	"github.com/spectrocloud/hapi/apiutil/transport"
	"github.com/spectrocloud/hapi/models"
	clusterC "github.com/spectrocloud/hapi/spectrocluster/client/v1"
)

func toV1OverlordsUIDOpenStackAccountValidateBody(account *models.V1OpenStackAccount) clusterC.V1OverlordsUIDOpenStackAccountValidateBody {
	return clusterC.V1OverlordsUIDOpenStackAccountValidateBody{
		Account: account.Spec,
	}
}

func (h *V1Client) CreateCloudAccountOpenStack(account *models.V1OpenStackAccount, AccountContext string) (string, error) {
	// validate account
	if err := validateCloudAccountOpenStack(account, h); err != nil {
		return "", err
	}

	var params *clusterC.V1CloudAccountsOpenStackCreateParams
	switch AccountContext {
	case "project":
		params = clusterC.NewV1CloudAccountsOpenStackCreateParamsWithContext(h.Ctx).WithBody(account)
	case "tenant":
		params = clusterC.NewV1CloudAccountsOpenStackCreateParams().WithBody(account)
	}

	success, err := h.GetClusterClient().V1CloudAccountsOpenStackCreate(params)
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

func validateCloudAccountOpenStack(account *models.V1OpenStackAccount, h *V1Client) error {
	PcgId := account.Metadata.Annotations[OverlordUID]
	// check PCG
	if err := h.CheckPCG(PcgId); err != nil {
		return err
	}

	// validate account
	paramsValidate := clusterC.NewV1OverlordsUIDOpenStackAccountValidateParams().WithUID(PcgId)
	paramsValidate = paramsValidate.WithBody(toV1OverlordsUIDOpenStackAccountValidateBody(account))
	_, err := h.GetClusterClient().V1OverlordsUIDOpenStackAccountValidate(paramsValidate)
	if err != nil {
		return err
	}

	return nil
}

func (h *V1Client) UpdateCloudAccountOpenStack(account *models.V1OpenStackAccount) error {
	if err := validateCloudAccountOpenStack(account, h); err != nil {
		return err
	}

	uid := account.Metadata.UID
	params := clusterC.NewV1CloudAccountsOpenStackUpdateParamsWithContext(h.Ctx).WithUID(uid).WithBody(account)
	_, err := h.GetClusterClient().V1CloudAccountsOpenStackUpdate(params)
	return err
}

func (h *V1Client) DeleteCloudAccountOpenStack(uid, AccountContext string) error {
	var params *clusterC.V1CloudAccountsOpenStackDeleteParams
	switch AccountContext {
	case "project":
		params = clusterC.NewV1CloudAccountsOpenStackDeleteParamsWithContext(h.Ctx).WithUID(uid)
	case "tenant":
		params = clusterC.NewV1CloudAccountsOpenStackDeleteParams().WithUID(uid)
	}

	_, err := h.GetClusterClient().V1CloudAccountsOpenStackDelete(params)
	return err
}

func (h *V1Client) GetCloudAccountOpenStack(uid, AccountContext string) (*models.V1OpenStackAccount, error) {
	var params *clusterC.V1CloudAccountsOpenStackGetParams
	switch AccountContext {
	case "project":
		params = clusterC.NewV1CloudAccountsOpenStackGetParamsWithContext(h.Ctx).WithUID(uid)
	case "tenant":
		params = clusterC.NewV1CloudAccountsOpenStackGetParams().WithUID(uid)
	}

	success, err := h.GetClusterClient().V1CloudAccountsOpenStackGet(params)

	var e *transport.TransportError
	if errors.As(err, &e) && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return success.Payload, nil
}

func (h *V1Client) GetCloudAccountsOpenStack() ([]*models.V1OpenStackAccount, error) {
	limit := int64(0)
	params := clusterC.NewV1CloudAccountsOpenStackListParamsWithContext(h.Ctx).WithLimit(&limit)
	response, err := h.GetClusterClient().V1CloudAccountsOpenStackList(params)
	if err != nil {
		return nil, err
	}

	accounts := make([]*models.V1OpenStackAccount, len(response.Payload.Items))
	copy(accounts, response.Payload.Items)

	return accounts, nil
}
