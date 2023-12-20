package client

import (
	"github.com/spectrocloud/hapi/models"
	userC "github.com/spectrocloud/hapi/user/client/v1"
)

func (h *V1Client) ListBackupStorageLocation(projectScope bool) ([]*models.V1UserAssetsLocation, error) {
	params := userC.NewV1UsersAssetsLocationGetParams()
	if projectScope {
		params.WithContext(h.Ctx)
	}
	response, err := h.GetUserClient().V1UsersAssetsLocationGet(params)
	if err != nil {
		return nil, err
	}

	bsls := make([]*models.V1UserAssetsLocation, len(response.Payload.Items))
	copy(bsls, response.Payload.Items)

	return bsls, nil
}

func (h *V1Client) GetBackupStorageLocation(uid string) (*models.V1UserAssetsLocation, error) {
	params := userC.NewV1UsersAssetsLocationGetParamsWithContext(h.Ctx)
	response, err := h.GetUserClient().V1UsersAssetsLocationGet(params)
	if err != nil {
		return nil, err
	}

	for _, account := range response.Payload.Items {
		if account.Metadata.UID == uid {
			return account, nil
		}
	}

	return nil, nil
}

func (h *V1Client) GetS3BackupStorageLocation(uid string) (*models.V1UserAssetsLocationS3, error) {
	params := userC.NewV1UsersAssetsLocationS3GetParamsWithContext(h.Ctx).WithUID(uid)
	if response, err := h.GetUserClient().V1UsersAssetsLocationS3Get(params); err != nil {
		return nil, err
	} else {
		return response.Payload, nil
	}
}

func (h *V1Client) CreateS3BackupStorageLocation(bsl *models.V1UserAssetsLocationS3) (string, error) {
	params := userC.NewV1UsersAssetsLocationS3CreateParamsWithContext(h.Ctx).WithBody(bsl)
	if resp, err := h.GetUserClient().V1UsersAssetsLocationS3Create(params); err != nil {
		return "", err
	} else {
		return *resp.Payload.UID, nil
	}
}

func (h *V1Client) UpdateS3BackupStorageLocation(uid string, bsl *models.V1UserAssetsLocationS3) error {
	params := userC.NewV1UsersAssetsLocationS3UpdateParamsWithContext(h.Ctx).WithUID(uid).WithBody(bsl)
	if _, err := h.GetUserClient().V1UsersAssetsLocationS3Update(params); err != nil {
		return err
	}
	return nil
}

func (h *V1Client) DeleteS3BackupStorageLocation(uid string) error {
	params := userC.NewV1UsersAssetsLocationS3DeleteParamsWithContext(h.Ctx).WithUID(uid)
	if _, err := h.GetUserClient().V1UsersAssetsLocationS3Delete(params); err != nil {
		return err
	}
	return nil
}
