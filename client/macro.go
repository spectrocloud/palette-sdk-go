package client

import (
	"hash/fnv"
	"strconv"

	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
)

func (h *V1Client) CreateMacro(uid string, macros *models.V1Macros) error {
	if uid != "" {
		params := clientV1.NewV1ProjectsUIDMacrosCreateParams().WithContext(h.Ctx).WithUID(uid).WithBody(macros)
		_, err := h.GetClient().V1ProjectsUIDMacrosCreate(params)
		if err != nil {
			return err
		}
	} else {
		tenantUID, err := h.GetTenantUID()
		if err != nil {
			return err
		}
		// As discussed with hubble team, we should not set context for tenant macros.
		params := clientV1.NewV1TenantsUIDMacrosCreateParams().WithTenantUID(tenantUID).WithBody(macros)
		_, err = h.GetClient().V1TenantsUIDMacrosCreate(params)
		if err != nil {
			return err
		}
	}

	return nil
}

func (h *V1Client) GetMacro(name, projectUID string) (*models.V1Macro, error) {
	macros, err := h.GetMacros(projectUID)
	if err != nil {
		return nil, err
	}
	id := h.GetMacroId(projectUID, name)

	for _, macro := range macros {
		if h.GetMacroId(projectUID, macro.Name) == id {
			return macro, nil
		}
	}

	return nil, nil
}

func (h *V1Client) GetMacros(projectUID string) ([]*models.V1Macro, error) {
	var macros []*models.V1Macro

	if projectUID != "" {
		params := clientV1.NewV1ProjectsUIDMacrosListParams().WithContext(h.Ctx).WithUID(projectUID)
		macrosListOk, err := h.GetClient().V1ProjectsUIDMacrosList(params)
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
		params := clientV1.NewV1TenantsUIDMacrosListParams().WithTenantUID(tenantUID)
		macrosListOk, err := h.GetClient().V1TenantsUIDMacrosList(params)
		if err != nil {
			return nil, err
		}
		macros = macrosListOk.Payload.Macros

	}

	return macros, nil
}

func (h *V1Client) StringHash(name string) string {
	return strconv.FormatUint(uint64(hash(name)), 10)
}

func hash(s string) uint32 {
	h := fnv.New32a()
	_, _ = h.Write([]byte(s))
	return h.Sum32()
}

func (h *V1Client) UpdateMacro(uid string, macros *models.V1Macros) error {
	if uid != "" {
		params := clientV1.NewV1ProjectsUIDMacrosUpdateByMacroNameParams().WithContext(h.Ctx).WithUID(uid).WithBody(macros)
		_, err := h.GetClient().V1ProjectsUIDMacrosUpdateByMacroName(params)
		return err
	} else {
		tenantUID, err := h.GetTenantUID()
		if err != nil || tenantUID == "" {
			return err
		}
		// As discussed with hubble team, we should not set context for tenant macros.
		params := clientV1.NewV1TenantsUIDMacrosUpdateByMacroNameParams().WithTenantUID(tenantUID).WithBody(macros)
		_, err = h.GetClient().V1TenantsUIDMacrosUpdateByMacroName(params)
		return err
	}
}

func (h *V1Client) DeleteMacro(uid string, body *models.V1Macros) error {
	if uid != "" {
		params := clientV1.NewV1ProjectsUIDMacrosDeleteByMacroNameParams().WithContext(h.Ctx).WithUID(uid).WithBody(body)
		_, err := h.GetClient().V1ProjectsUIDMacrosDeleteByMacroName(params)
		if err != nil {
			return err
		}
	} else {
		tenantUID, err := h.GetTenantUID()
		if err != nil {
			return err
		}
		// As discussed with hubble team, we should not set context for tenant macros.
		params := clientV1.NewV1TenantsUIDMacrosDeleteByMacroNameParams().WithTenantUID(tenantUID).WithBody(body)
		_, err = h.GetClient().V1TenantsUIDMacrosDeleteByMacroName(params)
		if err != nil {
			return err
		}
	}
	_, err := h.GetMacros(uid)
	if err != nil {
		return err
	}
	return nil
}

func (h *V1Client) GetMacroId(uid, name string) string {
	var hash string
	if uid != "" {
		hash = h.StringHash(name + uid)
	} else {
		hash = h.StringHash(name + "%tenant")
	}
	return hash
}
