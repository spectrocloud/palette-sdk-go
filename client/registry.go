package client

import (
	"fmt"

	"github.com/spectrocloud/hapi/models"
	clusterC "github.com/spectrocloud/hapi/spectrocluster/client/v1"
)

func (h *V1Client) SearchPackRegistryCommon() ([]*models.V1RegistryMetadata, error) {
	return h.getRegistryCommon()
}

func (h *V1Client) GetPackRegistryCommonByName(registryName string) (*models.V1RegistryMetadata, error) {
	if h.GetPackRegistryCommonByNameFn != nil {
		return h.GetPackRegistryCommonByNameFn(registryName)
	}

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
	params := clusterC.NewV1RegistriesMetadataParams().WithScope(Ptr(""))
	registries, err := h.GetClusterClient().V1RegistriesMetadata(params)
	if err != nil {
		return nil, err
	}
	return registries.Payload.Items, nil
}

func (h *V1Client) GetPackRegistryByName(registryName string) (*models.V1PackRegistry, error) {
	params := clusterC.NewV1RegistriesPackListParams()
	registries, err := h.GetClusterClient().V1RegistriesPackList(params)
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
	params := clusterC.NewV1RegistriesHelmSummaryListParams().WithScope(&scope)
	helmRegistries, err := h.GetClusterClient().V1RegistriesHelmSummaryList(params)
	if err != nil {
		return nil, err
	}

	return helmRegistries.Payload.Items, nil
}

func (h *V1Client) GetHelmRegistryByName(registryName string) (*models.V1HelmRegistry, error) {
	params := clusterC.NewV1RegistriesHelmListParams()
	registries, err := h.GetClusterClient().V1RegistriesHelmList(params)
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
	if h.GetHelmRegistryFn != nil {
		return h.GetHelmRegistryFn(uid)
	}

	params := clusterC.NewV1RegistriesHelmUIDGetParams().WithUID(uid)
	response, err := h.GetClusterClient().V1RegistriesHelmUIDGet(params)
	if err != nil {
		return nil, err
	}

	return response.Payload, nil
}

func (h *V1Client) CreateHelmRegistry(registry *models.V1HelmRegistryEntity) (string, error) {
	params := clusterC.NewV1RegistriesHelmCreateParams().WithBody(registry)
	if resp, err := h.GetClusterClient().V1RegistriesHelmCreate(params); err != nil {
		return "", err
	} else {
		return *resp.Payload.UID, nil
	}
}

func (h *V1Client) UpdateHelmRegistry(uid string, registry *models.V1HelmRegistry) error {
	params := clusterC.NewV1RegistriesHelmUIDUpdateParams().WithBody(registry).WithUID(uid)
	if _, err := h.GetClusterClient().V1RegistriesHelmUIDUpdate(params); err != nil {
		return err
	} else {
		return nil
	}
}

func (h *V1Client) DeleteHelmRegistry(uid string) error {
	params := clusterC.NewV1RegistriesHelmUIDDeleteParams().WithUID(uid)
	if _, err := h.GetClusterClient().V1RegistriesHelmUIDDelete(params); err != nil {
		return err
	} else {
		return nil
	}
}

func (h *V1Client) ListOCIRegistries() ([]*models.V1OciRegistry, error) {
	params := clusterC.NewV1OciRegistriesSummaryParams()
	ociRegistries, err := h.GetClusterClient().V1OciRegistriesSummary(params)
	if err != nil {
		return nil, err
	}

	return ociRegistries.Payload.Items, nil
}

func (h *V1Client) GetOciRegistryByName(registryName string) (*models.V1OciRegistry, error) {
	params := clusterC.NewV1OciRegistriesSummaryParams()
	registries, err := h.GetClusterClient().V1OciRegistriesSummary(params)
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
	params := clusterC.NewV1BasicOciRegistriesUIDGetParams().WithUID(uid)
	response, err := h.GetClusterClient().V1BasicOciRegistriesUIDGet(params)
	if err != nil {
		return nil, err
	}

	return response.Payload, nil
}

func (h *V1Client) CreateOciBasicRegistry(registry *models.V1BasicOciRegistry) (string, error) {
	params := clusterC.NewV1BasicOciRegistriesCreateParams().WithBody(registry)
	if resp, err := h.GetClusterClient().V1BasicOciRegistriesCreate(params); err != nil {
		return "", err
	} else {
		return *resp.Payload.UID, nil
	}
}

func (h *V1Client) UpdateOciBasicRegistry(uid string, registry *models.V1BasicOciRegistry) error {
	params := clusterC.NewV1BasicOciRegistriesUIDUpdateParams().WithBody(registry).WithUID(uid)
	_, err := h.GetClusterClient().V1BasicOciRegistriesUIDUpdate(params)
	if err != nil {
		return err
	}

	return nil
}

func (h *V1Client) DeleteOciBasicRegistry(uid string) error {
	params := clusterC.NewV1BasicOciRegistriesUIDDeleteParams().WithUID(uid)
	_, err := h.GetClusterClient().V1BasicOciRegistriesUIDDelete(params)
	if err != nil {
		return err
	}

	return nil
}

func (h *V1Client) GetOciEcrRegistry(uid string) (*models.V1EcrRegistry, error) {
	if h.GetOciRegistryFn != nil {
		return h.GetOciRegistryFn(uid)
	}

	params := clusterC.NewV1EcrRegistriesUIDGetParams().WithUID(uid)
	response, err := h.GetClusterClient().V1EcrRegistriesUIDGet(params)
	if err != nil {
		return nil, err
	}

	return response.Payload, nil
}

func (h *V1Client) CreateOciEcrRegistry(registry *models.V1EcrRegistry) (string, error) {
	if h.CreateOciEcrRegistryFn != nil {
		return h.CreateOciEcrRegistryFn(registry)
	}

	params := clusterC.NewV1EcrRegistriesCreateParams().WithBody(registry)
	if resp, err := h.GetClusterClient().V1EcrRegistriesCreate(params); err != nil {
		return "", err
	} else {
		return *resp.Payload.UID, nil
	}
}

func (h *V1Client) UpdateEcrRegistry(uid string, registry *models.V1EcrRegistry) error {
	if h.UpdateEcrRegistryFn != nil {
		return h.UpdateEcrRegistryFn(uid, registry)
	}
	params := clusterC.NewV1EcrRegistriesUIDUpdateParams().WithBody(registry).WithUID(uid)
	_, err := h.GetClusterClient().V1EcrRegistriesUIDUpdate(params)
	if err != nil {
		return err
	}

	return nil
}

func (h *V1Client) DeleteOciEcrRegistry(uid string) error {
	if h.DeleteOciEcrRegistryFn != nil {
		return h.DeleteOciEcrRegistryFn(uid)
	}

	params := clusterC.NewV1EcrRegistriesUIDDeleteParams().WithUID(uid)
	_, err := h.GetClusterClient().V1EcrRegistriesUIDDelete(params)
	return err
}
