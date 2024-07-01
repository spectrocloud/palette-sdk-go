package client

import (
	"fmt"

	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
)

func (h *V1Client) CreateMacros(uid string, macros *models.V1Macros) (string, error) {
	if uid != "" {
		params := clientV1.NewV1ProjectsUIDMacrosCreateParamsWithContext(h.ctx).
			WithUID(uid).
			WithBody(macros)
		_, err := h.Client.V1ProjectsUIDMacrosCreate(params)
		if err != nil {
			return "", err
		}
	} else {
		tenantUID, err := h.GetTenantUID()
		if err != nil {
			return "", err
		}
		// As discussed with hubble team, we should not set context for tenant macros.
		params := clientV1.NewV1TenantsUIDMacrosCreateParams().
			WithTenantUID(tenantUID).
			WithBody(macros)
		_, err = h.Client.V1TenantsUIDMacrosCreate(params)
		if err != nil {
			return "", err
		}
	}
	macrosId, err := h.GetMacrosId(uid)
	if err != nil {
		return "", err
	}
	return macrosId, nil
}

func (h *V1Client) GetTFMacrosV2(tfMacrosMap map[string]interface{}, projectUid string) ([]*models.V1Macro, error) {
	allMacros, err := h.GetMacrosV2(projectUid)
	if err != nil {
		return nil, err
	}

	var tfMacros []*models.V1Macro
	for _, v := range allMacros {
		if value, ok := tfMacrosMap[v.Name]; ok && value != nil {
			tfMacros = append(tfMacros, &models.V1Macro{
				Name:  v.Name,
				Value: v.Value,
			})
		}
	}

	return tfMacros, nil
}

func (h *V1Client) GetExistMacros(tfMacrosMap map[string]interface{}, projectUid string) ([]*models.V1Macro, error) {
	allMacros, err := h.GetMacrosV2(projectUid)
	if err != nil {
		return nil, err
	}

	var existMacros []*models.V1Macro
	for _, v := range allMacros {
		if _, ok := tfMacrosMap[v.Name]; !ok {
			existMacros = append(existMacros, &models.V1Macro{
				Name:  v.Name,
				Value: v.Value,
			})
		}
	}

	return existMacros, nil

}

func (h *V1Client) GetMacrosV2(projectUid string) ([]*models.V1Macro, error) {
	var macros []*models.V1Macro

	if projectUid != "" {
		params := clientV1.NewV1ProjectsUIDMacrosListParamsWithContext(h.ctx).
			WithUID(projectUid)
		macrosListOk, err := h.Client.V1ProjectsUIDMacrosList(params)
		if err != nil {
			return nil, err
		}
		macros = macrosListOk.Payload.Macros
	} else {
		tenantUID, err := h.GetTenantUID()
		if err != nil || tenantUID == "" {
			return nil, err
		}
		// As discussed with hubble team, we should not set context for tenant macros.
		params := clientV1.NewV1TenantsUIDMacrosListParams().
			WithTenantUID(tenantUID)
		macrosListOk, err := h.Client.V1TenantsUIDMacrosList(params)
		if err != nil {
			return nil, err
		}
		macros = macrosListOk.Payload.Macros
	}

	return macros, nil
}

func (h *V1Client) UpdateMacros(uid string, macros *models.V1Macros) error {
	if uid != "" {
		params := clientV1.NewV1ProjectsUIDMacrosUpdateParamsWithContext(h.ctx).
			WithUID(uid).
			WithBody(macros)
		_, err := h.Client.V1ProjectsUIDMacrosUpdate(params)
		return err
	} else {
		tenantUID, err := h.GetTenantUID()
		if err != nil || tenantUID == "" {
			return err
		}
		// As discussed with hubble team, we should not set context for tenant macros.
		params := clientV1.NewV1TenantsUIDMacrosUpdateParams().
			WithTenantUID(tenantUID).
			WithBody(macros)
		_, err = h.Client.V1TenantsUIDMacrosUpdate(params)
		return err
	}
}

func (h *V1Client) DeleteMacros(uid string, body *models.V1Macros) error {
	if uid != "" {
		params := clientV1.NewV1ProjectsUIDMacrosDeleteByMacroNameParamsWithContext(h.ctx).
			WithUID(uid).
			WithBody(body)
		_, err := h.Client.V1ProjectsUIDMacrosDeleteByMacroName(params)
		return err
	} else {
		tenantUID, err := h.GetTenantUID()
		if err != nil {
			return err
		}
		// As discussed with hubble team, we should not set context for tenant macros.
		params := clientV1.NewV1TenantsUIDMacrosDeleteByMacroNameParams().
			WithTenantUID(tenantUID).
			WithBody(body)
		_, err = h.Client.V1TenantsUIDMacrosDeleteByMacroName(params)
		return err
	}
}

func (h *V1Client) GetMacrosId(uid string) (string, error) {
	hashId := ""
	if uid != "" {
		hashId = fmt.Sprintf("%s-%s-%s", "project", "macros", uid)
	} else {
		tenantID, err := h.GetTenantUID()
		if err != nil {
			return "", err
		}
		hashId = fmt.Sprintf("%s-%s-%s", "tenant", "macros", tenantID)
	}
	return hashId, nil
}
