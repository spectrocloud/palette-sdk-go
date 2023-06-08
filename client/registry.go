package client

import (
	"fmt"

	"github.com/spectrocloud/gomi/pkg/ptr"
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
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}

	params := clusterC.NewV1RegistriesMetadataParams().WithScope(ptr.StringPtr(""))
	registries, err := client.V1RegistriesMetadata(params)
	if err != nil {
		return nil, err
	}
	return registries.Payload.Items, nil
}

func (h *V1Client) GetPackRegistryByName(registryName string) (*models.V1PackRegistry, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}

	params := clusterC.NewV1RegistriesPackListParams()
	registries, err := client.V1RegistriesPackList(params)
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
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}

	params := clusterC.NewV1RegistriesHelmSummaryListParams().WithContext(h.Ctx).WithScope(&scope)
	helmRegistries, err := client.V1RegistriesHelmSummaryList(params)
	if err != nil {
		return nil, err
	}

	return helmRegistries.Payload.Items, nil
}

func (h *V1Client) GetHelmRegistryByName(registryName string) (*models.V1HelmRegistry, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}

	params := clusterC.NewV1RegistriesHelmListParams().WithContext(h.Ctx)
	registries, err := client.V1RegistriesHelmList(params)
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
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}

	params := clusterC.NewV1RegistriesHelmUIDGetParams().WithContext(h.Ctx).WithUID(uid)
	response, err := client.V1RegistriesHelmUIDGet(params)
	if err != nil {
		return nil, err
	}

	return response.Payload, nil
}

func (h *V1Client) CreateHelmRegistry(registry *models.V1HelmRegistryEntity) (string, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return "", err
	}

	params := clusterC.NewV1RegistriesHelmCreateParams().WithContext(h.Ctx).WithBody(registry)
	if resp, err := client.V1RegistriesHelmCreate(params); err != nil {
		return "", err
	} else {
		return *resp.Payload.UID, nil
	}
}

func (h *V1Client) UpdateHelmRegistry(uid string, registry *models.V1HelmRegistry) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return err
	}

	params := clusterC.NewV1RegistriesHelmUIDUpdateParams().WithContext(h.Ctx).WithBody(registry).WithUID(uid)
	if _, err := client.V1RegistriesHelmUIDUpdate(params); err != nil {
		return err
	} else {
		return nil
	}
}

func (h *V1Client) DeleteHelmRegistry(uid string) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return err
	}

	params := clusterC.NewV1RegistriesHelmUIDDeleteParams().WithContext(h.Ctx).WithUID(uid)
	if _, err := client.V1RegistriesHelmUIDDelete(params); err != nil {
		return err
	} else {
		return nil
	}
}

func (h *V1Client) ListOCIRegistries() ([]*models.V1OciRegistry, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}

	params := clusterC.NewV1OciRegistriesSummaryParams().WithContext(h.Ctx)
	ociRegistries, err := client.V1OciRegistriesSummary(params)
	if err != nil {
		return nil, err
	}

	return ociRegistries.Payload.Items, nil
}

func (h *V1Client) GetOciRegistryByName(registryName string) (*models.V1OciRegistry, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}

	params := clusterC.NewV1OciRegistriesSummaryParams().WithContext(h.Ctx)
	registries, err := client.V1OciRegistriesSummary(params)
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
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}

	params := clusterC.NewV1BasicOciRegistriesUIDGetParams().WithContext(h.Ctx).WithUID(uid)
	response, err := client.V1BasicOciRegistriesUIDGet(params)
	if err != nil {
		return nil, err
	}

	return response.Payload, nil
}

func (h *V1Client) CreateOciBasicRegistry(registry *models.V1BasicOciRegistry) (string, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return "", err
	}

	params := clusterC.NewV1BasicOciRegistriesCreateParams().WithContext(h.Ctx).WithBody(registry)
	if resp, err := client.V1BasicOciRegistriesCreate(params); err != nil {
		return "", err
	} else {
		return *resp.Payload.UID, nil
	}
}

func (h *V1Client) UpdateOciBasicRegistry(uid string, registry *models.V1BasicOciRegistry) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return err
	}

	params := clusterC.NewV1BasicOciRegistriesUIDUpdateParams().WithContext(h.Ctx).WithBody(registry).WithUID(uid)
	_, err = client.V1BasicOciRegistriesUIDUpdate(params)
	if err != nil {
		return err
	}

	return nil
}

func (h *V1Client) DeleteOciBasicRegistry(uid string) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return err
	}

	params := clusterC.NewV1BasicOciRegistriesUIDDeleteParams().WithContext(h.Ctx).WithUID(uid)
	_, err = client.V1BasicOciRegistriesUIDDelete(params)
	if err != nil {
		return err
	}

	return nil
}

func (h *V1Client) GetOciEcrRegistry(uid string) (*models.V1EcrRegistry, error) {
	if h.GetOciRegistryFn != nil {
		return h.GetOciRegistryFn(uid)
	}
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}

	params := clusterC.NewV1EcrRegistriesUIDGetParams().WithContext(h.Ctx).WithUID(uid)
	response, err := client.V1EcrRegistriesUIDGet(params)
	if err != nil {
		return nil, err
	}

	return response.Payload, nil
}

func (h *V1Client) GetOciBasicRegistry(uid string) (*models.V1BasicOciRegistry, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}

	params := clusterC.NewV1BasicOciRegistriesUIDGetParams().WithUID(uid)
	response, err := client.V1BasicOciRegistriesUIDGet(params)
	if err != nil {
		return nil, err
	}

	return response.Payload, nil
}

func (h *V1Client) CreateOciEcrRegistry(registry *models.V1EcrRegistry) (string, error) {
	if h.CreateOciEcrRegistryFn != nil {
		return h.CreateOciEcrRegistryFn(registry)
	}
	client, err := h.GetClusterClient()
	if err != nil {
		return "", err
	}

	params := clusterC.NewV1EcrRegistriesCreateParams().WithContext(h.Ctx).WithBody(registry)
	if resp, err := client.V1EcrRegistriesCreate(params); err != nil {
		return "", err
	} else {
		return *resp.Payload.UID, nil
	}
}

func (h *V1Client) UpdateEcrRegistry(uid string, registry *models.V1EcrRegistry) error {
	if h.UpdateEcrRegistryFn != nil {
		return h.UpdateEcrRegistryFn(uid, registry)
	}
	client, err := h.GetClusterClient()
	if err != nil {
		return err
	}

	params := clusterC.NewV1EcrRegistriesUIDUpdateParams().WithContext(h.Ctx).WithBody(registry).WithUID(uid)
	_, err = client.V1EcrRegistriesUIDUpdate(params)
	if err != nil {
		return err
	}

	return nil
}

func (h *V1Client) DeleteOciEcrRegistry(uid string) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return err
	}

	params := clusterC.NewV1EcrRegistriesUIDDeleteParams().WithContext(h.Ctx).WithUID(uid)
	_, err = client.V1EcrRegistriesUIDDelete(params)
	if err != nil {
		return err
	}

	return nil
}
