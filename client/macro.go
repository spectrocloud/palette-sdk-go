package client

import (
	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
	"github.com/spectrocloud/palette-sdk-go/client/apiutil"
)

func (h *V1Client) CreateMacro(uid string, macros *models.V1Macros) error {
	if uid != "" {
		params := clientV1.NewV1ProjectsUIDMacrosCreateParamsWithContext(h.ctx).
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
	params := clientV1.NewV1TenantsUIDMacrosCreateParams().
		WithTenantUID(tenantUID).
		WithBody(macros)
	_, err = h.Client.V1TenantsUIDMacrosCreate(params)
	return err
}

func (h *V1Client) GetMacro(name, projectUID string) (*models.V1Macro, error) {
	macros, err := h.GetMacros(projectUID)
	if err != nil {
		return nil, err
	}
	id := h.GetMacroID(projectUID, name)

	for _, macro := range macros {
		if h.GetMacroID(projectUID, macro.Name) == id {
			return macro, nil
		}
	}
	return nil, nil
}

func (h *V1Client) GetMacros(projectUID string) ([]*models.V1Macro, error) {
	var macros []*models.V1Macro

	if projectUID != "" {
		params := clientV1.NewV1ProjectsUIDMacrosListParamsWithContext(h.ctx).
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
		params := clientV1.NewV1TenantsUIDMacrosListParams().
			WithTenantUID(tenantUID)
		resp, err := h.Client.V1TenantsUIDMacrosList(params)
		if err != nil {
			return nil, err
		}
		macros = resp.Payload.Macros
	}

	return macros, nil
}

func (h *V1Client) UpdateMacro(uid string, macros *models.V1Macros) error {
	if uid != "" {
		params := clientV1.NewV1ProjectsUIDMacrosUpdateByMacroNameParamsWithContext(h.ctx).
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
	params := clientV1.NewV1TenantsUIDMacrosUpdateByMacroNameParams().
		WithTenantUID(tenantUID).
		WithBody(macros)
	_, err = h.Client.V1TenantsUIDMacrosUpdateByMacroName(params)
	return err
}

func (h *V1Client) DeleteMacro(uid string, body *models.V1Macros) error {
	if uid != "" {
		params := clientV1.NewV1ProjectsUIDMacrosDeleteByMacroNameParamsWithContext(h.ctx).
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
		params := clientV1.NewV1TenantsUIDMacrosDeleteByMacroNameParams().
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

func (h *V1Client) GetMacroID(uid, name string) string {
	var hash string
	if uid != "" {
		hash = apiutil.StringHash(name + uid)
	} else {
		hash = apiutil.StringHash(name + "%tenant")
	}
	return hash
}
