package client

import (
	"fmt"

	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
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
	params := clientV1.NewV1RegistriesMetadataParams().WithScope(Ptr(""))
	registries, err := h.Client.V1RegistriesMetadata(params)
	if err != nil {
		return nil, err
	}
	return registries.Payload.Items, nil
}

func (h *V1Client) GetPackRegistryByName(registryName string) (*models.V1PackRegistry, error) {
	params := clientV1.NewV1RegistriesPackListParams()
	registries, err := h.Client.V1RegistriesPackList(params)
	if err != nil {
		return nil, err
	}

	for _, registry := range registries.Payload.Items {
		if registry.Metadata.Name == registryName {
			return registry, nil
		}
	}

	return nil, fmt.Errorf("registry '%s' not found", registryName)
}

func (h *V1Client) ListHelmRegistries(scope string) ([]*models.V1HelmRegistrySummary, error) {
	params := clientV1.NewV1RegistriesHelmSummaryListParams().WithScope(&scope)
	helmRegistries, err := h.Client.V1RegistriesHelmSummaryList(params)
	if err != nil {
		return nil, err
	}
	return helmRegistries.Payload.Items, nil
}

func (h *V1Client) GetHelmRegistryByName(registryName string) (*models.V1HelmRegistry, error) {
	params := clientV1.NewV1RegistriesHelmListParams()
	registries, err := h.Client.V1RegistriesHelmList(params)
	if err != nil {
		return nil, err
	}

	for _, registry := range registries.Payload.Items {
		if registry.Metadata.Name == registryName {
			return registry, nil
		}
	}

	return nil, fmt.Errorf("registry '%s' not found", registryName)
}

func (h *V1Client) GetHelmRegistry(uid string) (*models.V1HelmRegistry, error) {
	params := clientV1.NewV1RegistriesHelmUIDGetParams().WithUID(uid)
	response, err := h.Client.V1RegistriesHelmUIDGet(params)
	if err != nil {
		return nil, err
	}
	return response.Payload, nil
}

func (h *V1Client) CreateHelmRegistry(registry *models.V1HelmRegistryEntity) (string, error) {
	params := clientV1.NewV1RegistriesHelmCreateParams().WithBody(registry)
	if resp, err := h.Client.V1RegistriesHelmCreate(params); err != nil {
		return "", err
	} else {
		return *resp.Payload.UID, nil
	}
}

func (h *V1Client) UpdateHelmRegistry(uid string, registry *models.V1HelmRegistry) error {
	params := clientV1.NewV1RegistriesHelmUIDUpdateParams().WithBody(registry).WithUID(uid)
	if _, err := h.Client.V1RegistriesHelmUIDUpdate(params); err != nil {
		return err
	} else {
		return nil
	}
}

func (h *V1Client) DeleteHelmRegistry(uid string) error {
	params := clientV1.NewV1RegistriesHelmUIDDeleteParams().WithUID(uid)
	if _, err := h.Client.V1RegistriesHelmUIDDelete(params); err != nil {
		return err
	} else {
		return nil
	}
}

func (h *V1Client) ListOCIRegistries() ([]*models.V1OciRegistry, error) {
	params := clientV1.NewV1OciRegistriesSummaryParams()
	ociRegistries, err := h.Client.V1OciRegistriesSummary(params)
	if err != nil {
		return nil, err
	}
	return ociRegistries.Payload.Items, nil
}

func (h *V1Client) GetOciRegistryByName(registryName string) (*models.V1OciRegistry, error) {
	params := clientV1.NewV1OciRegistriesSummaryParams()
	registries, err := h.Client.V1OciRegistriesSummary(params)
	if err != nil {
		return nil, err
	}

	for _, registry := range registries.Payload.Items {
		if registry.Metadata.Name == registryName {
			return registry, nil
		}
	}

	return nil, fmt.Errorf("registry '%s' not found", registryName)
}

func (h *V1Client) GetOciBasicRegistry(uid string) (*models.V1BasicOciRegistry, error) {
	params := clientV1.NewV1BasicOciRegistriesUIDGetParams().WithUID(uid)
	response, err := h.Client.V1BasicOciRegistriesUIDGet(params)
	if err != nil {
		return nil, err
	}
	return response.Payload, nil
}

func (h *V1Client) CreateOciBasicRegistry(registry *models.V1BasicOciRegistry) (string, error) {
	params := clientV1.NewV1BasicOciRegistriesCreateParams().WithBody(registry)
	if resp, err := h.Client.V1BasicOciRegistriesCreate(params); err != nil {
		return "", err
	} else {
		return *resp.Payload.UID, nil
	}
}

func (h *V1Client) UpdateOciBasicRegistry(uid string, registry *models.V1BasicOciRegistry) error {
	params := clientV1.NewV1BasicOciRegistriesUIDUpdateParams().WithBody(registry).WithUID(uid)
	_, err := h.Client.V1BasicOciRegistriesUIDUpdate(params)
	if err != nil {
		return err
	}
	return nil
}

func (h *V1Client) DeleteOciBasicRegistry(uid string) error {
	params := clientV1.NewV1BasicOciRegistriesUIDDeleteParams().WithUID(uid)
	_, err := h.Client.V1BasicOciRegistriesUIDDelete(params)
	if err != nil {
		return err
	}
	return nil
}

func (h *V1Client) GetOciEcrRegistry(uid string) (*models.V1EcrRegistry, error) {
	params := clientV1.NewV1EcrRegistriesUIDGetParams().WithUID(uid)
	response, err := h.Client.V1EcrRegistriesUIDGet(params)
	if err != nil {
		return nil, err
	}
	return response.Payload, nil
}

func (h *V1Client) CreateOciEcrRegistry(registry *models.V1EcrRegistry) (string, error) {
	params := clientV1.NewV1EcrRegistriesCreateParams().WithBody(registry)
	if resp, err := h.Client.V1EcrRegistriesCreate(params); err != nil {
		return "", err
	} else {
		return *resp.Payload.UID, nil
	}
}

func (h *V1Client) UpdateEcrRegistry(uid string, registry *models.V1EcrRegistry) error {
	params := clientV1.NewV1EcrRegistriesUIDUpdateParams().WithBody(registry).WithUID(uid)
	_, err := h.Client.V1EcrRegistriesUIDUpdate(params)
	if err != nil {
		return err
	}
	return nil
}

func (h *V1Client) DeleteOciEcrRegistry(uid string) error {
	params := clientV1.NewV1EcrRegistriesUIDDeleteParams().WithUID(uid)
	_, err := h.Client.V1EcrRegistriesUIDDelete(params)
	return err
}
