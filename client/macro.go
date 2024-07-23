package client

import (
	clientv1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
	"github.com/spectrocloud/palette-sdk-go/client/apiutil"
)

// CreateMacro creates a new macro.
func (h *V1Client) CreateMacro(uid string, macros *models.V1Macros) error {
	if uid != "" {
		params := clientv1.NewV1ProjectsUIDMacrosCreateParamsWithContext(h.ctx).
			WithUID(uid).
			WithBody(macros)
		_, err := h.Client.V1ProjectsUIDMacrosCreate(params)
		return err
	}
	tenantUID, err := h.GetTenantUID()
	if err != nil {
		return err
	}
	// As discussed with hubble team, we should not set context for tenant macros.
	params := clientv1.NewV1TenantsUIDMacrosCreateParams().
		WithTenantUID(tenantUID).
		WithBody(macros)
	_, err = h.Client.V1TenantsUIDMacrosCreate(params)
	return err
}

// GetMacro retrieves an existing macro by name and project UID.
func (h *V1Client) GetMacro(name, projectUID string) (*models.V1Macro, error) {
	macros, err := h.GetMacros(projectUID)
	if err != nil {
		return nil, err
	}
	id := GetMacroID(projectUID, name)

	for _, macro := range macros {
		if GetMacroID(projectUID, macro.Name) == id {
			return macro, nil
		}
	}
	return nil, nil
}

// GetMacros retrieves all macros in a project or tenant.
// If projectUID is empty, tenant macros are returned.
func (h *V1Client) GetMacros(projectUID string) ([]*models.V1Macro, error) {
	var macros []*models.V1Macro

	if projectUID != "" {
		params := clientv1.NewV1ProjectsUIDMacrosListParamsWithContext(h.ctx).
			WithUID(projectUID)
		resp, err := h.Client.V1ProjectsUIDMacrosList(params)
		if err != nil {
			return nil, err
		}
		macros = resp.Payload.Macros
	} else {
		tenantUID, err := h.GetTenantUID()
		if err != nil || tenantUID == "" {
			return nil, err
		}
		// As discussed with hubble team, we should not set context for tenant macros.
		params := clientv1.NewV1TenantsUIDMacrosListParams().
			WithTenantUID(tenantUID)
		resp, err := h.Client.V1TenantsUIDMacrosList(params)
		if err != nil {
			return nil, err
		}
		macros = resp.Payload.Macros
	}

	return macros, nil
}

// UpdateMacro updates an existing macro.
func (h *V1Client) UpdateMacro(uid string, macros *models.V1Macros) error {
	if uid != "" {
		params := clientv1.NewV1ProjectsUIDMacrosUpdateByMacroNameParamsWithContext(h.ctx).
			WithUID(uid).
			WithBody(macros)
		_, err := h.Client.V1ProjectsUIDMacrosUpdateByMacroName(params)
		return err
	}
	tenantUID, err := h.GetTenantUID()
	if err != nil || tenantUID == "" {
		return err
	}
	// As discussed with hubble team, we should not set context for tenant macros.
	params := clientv1.NewV1TenantsUIDMacrosUpdateByMacroNameParams().
		WithTenantUID(tenantUID).
		WithBody(macros)
	_, err = h.Client.V1TenantsUIDMacrosUpdateByMacroName(params)
	return err
}

// DeleteMacro deletes an existing macro.
func (h *V1Client) DeleteMacro(uid string, body *models.V1Macros) error {
	if uid != "" {
		params := clientv1.NewV1ProjectsUIDMacrosDeleteByMacroNameParamsWithContext(h.ctx).
			WithUID(uid).
			WithBody(body)
		_, err := h.Client.V1ProjectsUIDMacrosDeleteByMacroName(params)
		if err != nil {
			return err
		}
	} else {
		tenantUID, err := h.GetTenantUID()
		if err != nil {
			return err
		}
		// As discussed with hubble team, we should not set context for tenant macros.
		params := clientv1.NewV1TenantsUIDMacrosDeleteByMacroNameParams().
			WithTenantUID(tenantUID).
			WithBody(body)
		_, err = h.Client.V1TenantsUIDMacrosDeleteByMacroName(params)
		if err != nil {
			return err
		}
	}
	_, err := h.GetMacros(uid)
	return err
}

// GetMacroID returns the hash of the macro name and project UID.
// Provide an empty string for project UID to get the hash for tenant macros.
func GetMacroID(uid, name string) string {
	var hash string
	if uid != "" {
		hash = apiutil.StringHash(name + uid)
	} else {
		hash = apiutil.StringHash(name + "%tenant")
	}
	return hash
}
