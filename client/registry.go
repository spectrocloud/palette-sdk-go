package client

import (
	"fmt"

	clientv1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
	"github.com/spectrocloud/palette-sdk-go/client/apiutil"
)

// SearchPackRegistryCommon retrieves a list of all registries.
func (h *V1Client) SearchPackRegistryCommon() ([]*models.V1RegistryMetadata, error) {
	return h.getRegistryCommon()
}

// GetPackRegistryCommonByName retrieves any type of registry by name.
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
	params := clientv1.NewV1RegistriesMetadataParamsWithContext(h.ctx).
		WithScope(apiutil.Ptr(""))
	resp, err := h.Client.V1RegistriesMetadata(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload.Items, nil
}

// GetPackRegistryByName retrieves an existing Pack registry by name.
func (h *V1Client) GetPackRegistryByName(registryName string) (*models.V1PackRegistry, error) {
	params := clientv1.NewV1RegistriesPackListParamsWithContext(h.ctx)
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

// ListHelmRegistries retrieves a list of all Helm registries, filtered by scope.
func (h *V1Client) ListHelmRegistries(scope string) ([]*models.V1HelmRegistrySummary, error) {
	params := clientv1.NewV1RegistriesHelmSummaryListParamsWithContext(h.ctx).
		WithScope(&scope)
	resp, err := h.Client.V1RegistriesHelmSummaryList(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload.Items, nil
}

// GetHelmRegistryByName retrieves an existing Helm registry by name.
func (h *V1Client) GetHelmRegistryByName(registryName string) (*models.V1HelmRegistry, error) {
	params := clientv1.NewV1RegistriesHelmListParamsWithContext(h.ctx)
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

// GetHelmRegistry retrieves an existing Helm registry by UID.
func (h *V1Client) GetHelmRegistry(uid string) (*models.V1HelmRegistry, error) {
	params := clientv1.NewV1RegistriesHelmUIDGetParamsWithContext(h.ctx).
		WithUID(uid)
	resp, err := h.Client.V1RegistriesHelmUIDGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// CreateHelmRegistry creates a new Helm registry.
func (h *V1Client) CreateHelmRegistry(registry *models.V1HelmRegistryEntity) (string, error) {
	params := clientv1.NewV1RegistriesHelmCreateParamsWithContext(h.ctx).
		WithBody(registry)
	resp, err := h.Client.V1RegistriesHelmCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

// UpdateHelmRegistry updates an existing Helm registry.
func (h *V1Client) UpdateHelmRegistry(uid string, registry *models.V1HelmRegistry) error {
	params := clientv1.NewV1RegistriesHelmUIDUpdateParamsWithContext(h.ctx).
		WithUID(uid).
		WithBody(registry)
	_, err := h.Client.V1RegistriesHelmUIDUpdate(params)
	return err
}

// DeleteHelmRegistry deletes an existing Helm registry.
func (h *V1Client) DeleteHelmRegistry(uid string) error {
	params := clientv1.NewV1RegistriesHelmUIDDeleteParamsWithContext(h.ctx).
		WithUID(uid)
	_, err := h.Client.V1RegistriesHelmUIDDelete(params)
	return err
}

// ListOCIRegistries retrieves a list of all OCI registries.
func (h *V1Client) ListOCIRegistries() ([]*models.V1OciRegistry, error) {
	params := clientv1.NewV1OciRegistriesSummaryParamsWithContext(h.ctx)
	resp, err := h.Client.V1OciRegistriesSummary(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload.Items, nil
}

// GetOciRegistryByName retrieves an existing OCI registry by name.
func (h *V1Client) GetOciRegistryByName(registryName string) (*models.V1OciRegistry, error) {
	params := clientv1.NewV1OciRegistriesSummaryParamsWithContext(h.ctx)
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

// GetOciBasicRegistry retrieves an existing standard/basic OCI registry by UID.
func (h *V1Client) GetOciBasicRegistry(uid string) (*models.V1BasicOciRegistry, error) {
	params := clientv1.NewV1BasicOciRegistriesUIDGetParamsWithContext(h.ctx).
		WithUID(uid)
	resp, err := h.Client.V1BasicOciRegistriesUIDGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// CreateOciBasicRegistry creates a new standard/basic OCI registry.
func (h *V1Client) CreateOciBasicRegistry(registry *models.V1BasicOciRegistry) (string, error) {
	params := clientv1.NewV1BasicOciRegistriesCreateParamsWithContext(h.ctx).
		WithBody(registry)
	resp, err := h.Client.V1BasicOciRegistriesCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

// UpdateOciBasicRegistry updates an existing standard/basic OCI registry.
func (h *V1Client) UpdateOciBasicRegistry(uid string, registry *models.V1BasicOciRegistry) error {
	params := clientv1.NewV1BasicOciRegistriesUIDUpdateParamsWithContext(h.ctx).
		WithUID(uid).
		WithBody(registry)
	_, err := h.Client.V1BasicOciRegistriesUIDUpdate(params)
	return err
}

// DeleteOciBasicRegistry deletes an existing standard/basic OCI registry.
func (h *V1Client) DeleteOciBasicRegistry(uid string) error {
	params := clientv1.NewV1BasicOciRegistriesUIDDeleteParamsWithContext(h.ctx).
		WithUID(uid)
	_, err := h.Client.V1BasicOciRegistriesUIDDelete(params)
	return err
}

// GetOciEcrRegistry retrieves an existing ECR OCI registry by UID.
func (h *V1Client) GetOciEcrRegistry(uid string) (*models.V1EcrRegistry, error) {
	params := clientv1.NewV1EcrRegistriesUIDGetParamsWithContext(h.ctx).
		WithUID(uid)
	resp, err := h.Client.V1EcrRegistriesUIDGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// CreateOciEcrRegistry creates a new ECR OCI registry.
func (h *V1Client) CreateOciEcrRegistry(registry *models.V1EcrRegistry) (string, error) {
	params := clientv1.NewV1EcrRegistriesCreateParamsWithContext(h.ctx).
		WithBody(registry)
	resp, err := h.Client.V1EcrRegistriesCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

// UpdateOciEcrRegistry updates an existing ECR OCI registry.
func (h *V1Client) UpdateOciEcrRegistry(uid string, registry *models.V1EcrRegistry) error {
	params := clientv1.NewV1EcrRegistriesUIDUpdateParamsWithContext(h.ctx).
		WithUID(uid).
		WithBody(registry)
	_, err := h.Client.V1EcrRegistriesUIDUpdate(params)
	return err
}

// DeleteOciEcrRegistry deletes an existing ECR OCI registry.
func (h *V1Client) DeleteOciEcrRegistry(uid string) error {
	params := clientv1.NewV1EcrRegistriesUIDDeleteParamsWithContext(h.ctx).
		WithUID(uid)
	_, err := h.Client.V1EcrRegistriesUIDDelete(params)
	return err
}
