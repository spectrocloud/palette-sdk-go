package client

import (
	"fmt"

	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
	"github.com/spectrocloud/palette-sdk-go/client/apiutil"
)

func (h *V1Client) SearchPackRegistryCommon() ([]*models.V1RegistryMetadata, error) {
	return h.getRegistryCommon()
}

func (h *V1Client) GetPackRegistryCommonByName(registryName string) (*models.V1RegistryMetadata, error) {
	registries, err := h.getRegistryCommon()
	if err != nil {
		return nil, err
	}
	for _, registry := range registries {
		if registry.Name == registryName {
			return registry, nil
		}
	}
	return nil, fmt.Errorf("registry '%s' not found", registryName)
}

func (h *V1Client) getRegistryCommon() ([]*models.V1RegistryMetadata, error) {
	params := clientV1.NewV1RegistriesMetadataParamsWithContext(h.ctx).
		WithScope(apiutil.Ptr(""))
	resp, err := h.Client.V1RegistriesMetadata(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload.Items, nil
}

func (h *V1Client) GetPackRegistryByName(registryName string) (*models.V1PackRegistry, error) {
	params := clientV1.NewV1RegistriesPackListParamsWithContext(h.ctx)
	resp, err := h.Client.V1RegistriesPackList(params)
	if err != nil {
		return nil, err
	}
	for _, registry := range resp.Payload.Items {
		if registry.Metadata.Name == registryName {
			return registry, nil
		}
	}
	return nil, fmt.Errorf("registry '%s' not found", registryName)
}

func (h *V1Client) ListHelmRegistries(scope string) ([]*models.V1HelmRegistrySummary, error) {
	params := clientV1.NewV1RegistriesHelmSummaryListParamsWithContext(h.ctx).
		WithScope(&scope)
	resp, err := h.Client.V1RegistriesHelmSummaryList(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload.Items, nil
}

func (h *V1Client) GetHelmRegistryByName(registryName string) (*models.V1HelmRegistry, error) {
	params := clientV1.NewV1RegistriesHelmListParamsWithContext(h.ctx)
	resp, err := h.Client.V1RegistriesHelmList(params)
	if err != nil {
		return nil, err
	}
	for _, registry := range resp.Payload.Items {
		if registry.Metadata.Name == registryName {
			return registry, nil
		}
	}
	return nil, fmt.Errorf("registry '%s' not found", registryName)
}

func (h *V1Client) GetHelmRegistry(uid string) (*models.V1HelmRegistry, error) {

	params := clientV1.NewV1RegistriesHelmUIDGetParamsWithContext(h.ctx).
		WithUID(uid)
	resp, err := h.Client.V1RegistriesHelmUIDGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

func (h *V1Client) CreateHelmRegistry(registry *models.V1HelmRegistryEntity) (string, error) {
	params := clientV1.NewV1RegistriesHelmCreateParamsWithContext(h.ctx).
		WithBody(registry)
	resp, err := h.Client.V1RegistriesHelmCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

func (h *V1Client) UpdateHelmRegistry(uid string, registry *models.V1HelmRegistry) error {
	params := clientV1.NewV1RegistriesHelmUIDUpdateParamsWithContext(h.ctx).
		WithUID(uid).
		WithBody(registry)
	_, err := h.Client.V1RegistriesHelmUIDUpdate(params)
	return err
}

func (h *V1Client) DeleteHelmRegistry(uid string) error {
	params := clientV1.NewV1RegistriesHelmUIDDeleteParamsWithContext(h.ctx).
		WithUID(uid)
	_, err := h.Client.V1RegistriesHelmUIDDelete(params)
	return err
}

func (h *V1Client) ListOCIRegistries() ([]*models.V1OciRegistry, error) {
	params := clientV1.NewV1OciRegistriesSummaryParamsWithContext(h.ctx)
	resp, err := h.Client.V1OciRegistriesSummary(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload.Items, nil
}

func (h *V1Client) GetOciRegistryByName(registryName string) (*models.V1OciRegistry, error) {
	params := clientV1.NewV1OciRegistriesSummaryParamsWithContext(h.ctx)
	resp, err := h.Client.V1OciRegistriesSummary(params)
	if err != nil {
		return nil, err
	}
	for _, registry := range resp.Payload.Items {
		if registry.Metadata.Name == registryName {
			return registry, nil
		}
	}
	return nil, fmt.Errorf("registry '%s' not found", registryName)
}

func (h *V1Client) GetOciBasicRegistry(uid string) (*models.V1BasicOciRegistry, error) {
	params := clientV1.NewV1BasicOciRegistriesUIDGetParamsWithContext(h.ctx).
		WithUID(uid)
	resp, err := h.Client.V1BasicOciRegistriesUIDGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

func (h *V1Client) CreateOciBasicRegistry(registry *models.V1BasicOciRegistry) (string, error) {
	params := clientV1.NewV1BasicOciRegistriesCreateParamsWithContext(h.ctx).
		WithBody(registry)
	resp, err := h.Client.V1BasicOciRegistriesCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

func (h *V1Client) UpdateOciBasicRegistry(uid string, registry *models.V1BasicOciRegistry) error {
	params := clientV1.NewV1BasicOciRegistriesUIDUpdateParamsWithContext(h.ctx).
		WithUID(uid).
		WithBody(registry)
	_, err := h.Client.V1BasicOciRegistriesUIDUpdate(params)
	return err
}

func (h *V1Client) DeleteOciBasicRegistry(uid string) error {
	params := clientV1.NewV1BasicOciRegistriesUIDDeleteParamsWithContext(h.ctx).
		WithUID(uid)
	_, err := h.Client.V1BasicOciRegistriesUIDDelete(params)
	return err
}

func (h *V1Client) GetOciEcrRegistry(uid string) (*models.V1EcrRegistry, error) {
	params := clientV1.NewV1EcrRegistriesUIDGetParamsWithContext(h.ctx).
		WithUID(uid)
	resp, err := h.Client.V1EcrRegistriesUIDGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

func (h *V1Client) CreateOciEcrRegistry(registry *models.V1EcrRegistry) (string, error) {
	params := clientV1.NewV1EcrRegistriesCreateParamsWithContext(h.ctx).
		WithBody(registry)
	resp, err := h.Client.V1EcrRegistriesCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

func (h *V1Client) UpdateEcrRegistry(uid string, registry *models.V1EcrRegistry) error {
	params := clientV1.NewV1EcrRegistriesUIDUpdateParamsWithContext(h.ctx).
		WithUID(uid).
		WithBody(registry)
	_, err := h.Client.V1EcrRegistriesUIDUpdate(params)
	return err
}

func (h *V1Client) DeleteOciEcrRegistry(uid string) error {
	params := clientV1.NewV1EcrRegistriesUIDDeleteParamsWithContext(h.ctx).
		WithUID(uid)
	_, err := h.Client.V1EcrRegistriesUIDDelete(params)
	return err
}
