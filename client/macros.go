package client

import (
	"fmt"

	clientv1 "github.com/spectrocloud/palette-sdk-go/api/client/v1"
	"github.com/spectrocloud/palette-sdk-go/api/models"
)

// CreateMacros creates a new macro.
func (h *V1Client) CreateMacros(uid string, macros *models.V1Macros) (string, error) {
	if uid != "" {
		params := clientv1.NewV1ProjectsUIDMacrosCreateParamsWithContext(h.ctx).
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
		params := clientv1.NewV1TenantsUIDMacrosCreateParams().
			WithTenantUID(tenantUID).
			WithBody(macros)
		_, err = h.Client.V1TenantsUIDMacrosCreate(params)
		if err != nil {
			return "", err
		}
	}
	macrosID, err := h.GetMacrosID(uid)
	if err != nil {
		return "", err
	}
	return macrosID, nil
}

// GetTFMacrosV2 retrieves all macros in a project or tenant, filtered by the given map.
// Macros whose names are in the map and values are non-nil are returned.
func (h *V1Client) GetTFMacrosV2(tfMacrosMap map[string]any, projectUID string) ([]*models.V1Macro, error) {
	allMacros, err := h.GetMacrosV2(projectUID)
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

// GetExistMacros retrieves all macros in a project or tenant, filtered by the given map.
// Macros whose names are not in the map are returned.
func (h *V1Client) GetExistMacros(tfMacrosMap map[string]any, projectUID string) ([]*models.V1Macro, error) {
	allMacros, err := h.GetMacrosV2(projectUID)
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

// GetMacrosV2 retrieves all existing macros at either project or tenant scope.
func (h *V1Client) GetMacrosV2(projectUID string) ([]*models.V1Macro, error) {
	if projectUID != "" {
		params := clientv1.NewV1ProjectsUIDMacrosListParamsWithContext(h.ctx).
			WithUID(projectUID)
		macrosListOk, err := h.Client.V1ProjectsUIDMacrosList(params)
		if err != nil {
			return nil, err
		}
		return macrosListOk.Payload.Macros, nil
	}
	tenantUID, err := h.GetTenantUID()
	if err != nil || tenantUID == "" {
		return nil, err
	}
	// As discussed with hubble team, we should not set context for tenant macros.
	params := clientv1.NewV1TenantsUIDMacrosListParams().
		WithTenantUID(tenantUID)
	macrosListOk, err := h.Client.V1TenantsUIDMacrosList(params)
	if err != nil {
		return nil, err
	}
	return macrosListOk.Payload.Macros, nil
}

// UpdateMacros updates an existing macro.
func (h *V1Client) UpdateMacros(uid string, macros *models.V1Macros) error {
	if uid != "" {
		params := clientv1.NewV1ProjectsUIDMacrosUpdateParamsWithContext(h.ctx).
			WithUID(uid).
			WithBody(macros)
		_, err := h.Client.V1ProjectsUIDMacrosUpdate(params)
		return err
	}
	tenantUID, err := h.GetTenantUID()
	if err != nil || tenantUID == "" {
		return err
	}
	// As discussed with hubble team, we should not set context for tenant macros.
	params := clientv1.NewV1TenantsUIDMacrosUpdateParams().
		WithTenantUID(tenantUID).
		WithBody(macros)
	_, err = h.Client.V1TenantsUIDMacrosUpdate(params)
	return err
}

// DeleteMacros deletes an existing macro.
func (h *V1Client) DeleteMacros(uid string, body *models.V1Macros) error {
	if uid != "" {
		params := clientv1.NewV1ProjectsUIDMacrosDeleteByMacroNameParamsWithContext(h.ctx).
			WithUID(uid).
			WithBody(body)
		_, err := h.Client.V1ProjectsUIDMacrosDeleteByMacroName(params)
		return err
	}
	tenantUID, err := h.GetTenantUID()
	if err != nil {
		return err
	}
	// As discussed with hubble team, we should not set context for tenant macros.
	params := clientv1.NewV1TenantsUIDMacrosDeleteByMacroNameParams().
		WithTenantUID(tenantUID).
		WithBody(body)
	_, err = h.Client.V1TenantsUIDMacrosDeleteByMacroName(params)
	return err
}

// GetMacrosID returns the hash ID for a macro, given a project UID.
// If the project UID is empty, the tenant UID is used.
func (h *V1Client) GetMacrosID(uid string) (string, error) {
	hashID := ""
	if uid != "" {
		hashID = fmt.Sprintf("%s-%s-%s", "project", "macros", uid)
	} else {
		tenantUID, err := h.GetTenantUID()
		if err != nil {
			return "", err
		}
		hashID = fmt.Sprintf("%s-%s-%s", "tenant", "macros", tenantUID)
	}
	return hashID, nil
}
