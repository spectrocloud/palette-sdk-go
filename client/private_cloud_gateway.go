package client

import (
	"errors"
	"fmt"

	clientv1 "github.com/spectrocloud/palette-sdk-go/api/client/v1"
	"github.com/spectrocloud/palette-sdk-go/api/models"
	"github.com/spectrocloud/palette-sdk-go/client/apiutil"
)

// PCG - Generic

// GetPCGId retrieves the UID of a Private Cloud Gateway by name.
func (h *V1Client) GetPCGId(name *string) (string, error) {
	params := clientv1.NewV1OverlordsListParamsWithContext(h.ctx)
	resp, err := h.Client.V1OverlordsList(params)
	if err != nil {
		return "", err
	}
	for _, pcg := range resp.Payload.Items {
		if pcg.Metadata.Name == *name {
			return pcg.Metadata.UID, nil
		}
	}
	return "", errors.New("Private Cloud Gateway not found: " + *name)
}

// GetPCGByID retrieves a Private Cloud Gateway by UID.
func (h *V1Client) GetPCGByID(uid string) (*models.V1Overlord, error) {
	params := clientv1.NewV1OverlordsUIDGetParamsWithContext(h.ctx).
		WithUID(uid)
	resp, err := h.Client.V1OverlordsUIDGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// GetPCGByName retrieves a Private Cloud Gateway by name.
func (h *V1Client) GetPCGByName(name string) (*models.V1Overlord, error) {
	params := clientv1.NewV1OverlordsListParamsWithContext(h.ctx).
		WithName(&name)
	resp, err := h.Client.V1OverlordsList(params)
	if err != nil {
		return nil, err
	}
	if len(resp.Payload.Items) > 0 {
		return resp.Payload.Items[0], nil
	}
	return nil, nil
}

// GetPairingCode retrieves a pairing code for a Private Cloud Gateway by cloud type.
func (h *V1Client) GetPairingCode(cloudType string) (string, error) {
	params := clientv1.NewV1OverlordsPairingCodeParamsWithContext(h.ctx).
		WithCloudType(&cloudType)
	resp, err := h.Client.V1OverlordsPairingCode(params)
	if err != nil {
		return "", err
	}
	return resp.Payload.PairingCode, nil
}

// CheckPCG checks if a Private Cloud Gateway is running.
// Returns an error if the PCG is not found or is not running.
func (h *V1Client) CheckPCG(pcgID string) error {
	if pcgID == "" {
		return nil // no pcg to check
	}
	pcg, err := h.GetPCGByID(pcgID)
	if err != nil {
		return err
	}
	if pcg == nil {
		return fmt.Errorf("private cloud gateway not found: %s", pcgID)
	}
	if pcg.Status == nil {
		return fmt.Errorf("private cloud gateway status not found: %s", pcgID)
	}
	if pcg.Status.State != "Running" {
		return fmt.Errorf("private cloud gateway is not running: %s", pcgID)
	}
	return nil // pcg is running
}

// PCG - vSphere

// CreatePCGVsphere creates a new vSphere Private Cloud Gateway.
func (h *V1Client) CreatePCGVsphere(uid string, cloudConfig *models.V1OverlordVsphereCloudConfig) (string, error) {
	params := clientv1.NewV1OverlordsUIDVsphereCloudConfigCreateParamsWithContext(h.ctx).
		WithUID(uid).
		WithBody(cloudConfig)
	resp, err := h.Client.V1OverlordsUIDVsphereCloudConfigCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

// UpdatePCGVsphere updates a vSphere Private Cloud Gateway VsphereCloudConfig.
func (h *V1Client) UpdatePCGVsphere(uid string, cloudConfig *models.V1OverlordVsphereCloudConfig) error {
	params := clientv1.NewV1OverlordsUIDVsphereCloudConfigUpdateParamsWithContext(h.ctx).
		WithUID(uid).
		WithBody(cloudConfig)
	_, err := h.Client.V1OverlordsUIDVsphereCloudConfigUpdate(params)
	return err
}

// CreatePCGCloudAccountVsphere creates a new vSphere PCG cloud account.
func (h *V1Client) CreatePCGCloudAccountVsphere(uid string, account *models.V1OverlordVsphereAccountCreate) (string, error) {
	params := clientv1.NewV1OverlordsUIDVsphereAccountCreateParamsWithContext(h.ctx).
		WithUID(uid).
		WithBody(account)
	resp, err := h.Client.V1OverlordsUIDVsphereAccountCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

// GetPCGManifestVsphere retrieves a vSphere PCG manifest by pairing code.
func (h *V1Client) GetPCGManifestVsphere(pairingCode string) (string, error) {
	params := clientv1.NewV1OverlordsVsphereManifestParamsWithContext(h.ctx).
		WithPairingCode(pairingCode)
	resp, err := h.Client.V1OverlordsVsphereManifest(params)
	if err != nil {
		return "", err
	}
	return resp.Payload.Manifest, nil
}

// GetPCGClusterProfileVsphere retrieves a vSphere PCG cluster profile by PCG UID.
func (h *V1Client) GetPCGClusterProfileVsphere(uid string) (*models.V1ClusterProfile, error) {
	params := clientv1.NewV1OverlordsUIDVsphereClusterProfileParamsWithContext(h.ctx).
		WithUID(uid)
	resp, err := h.Client.V1OverlordsUIDVsphereClusterProfile(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// CreateDDNSSearchDomainVsphere creates a new vSphere DDNS search domain.
func (h *V1Client) CreateDDNSSearchDomainVsphere(vsphereDNSMapping *models.V1VsphereDNSMapping) error {
	params := clientv1.NewV1VsphereDNSMappingCreateParamsWithContext(h.ctx).
		WithBody(vsphereDNSMapping)
	_, err := h.Client.V1VsphereDNSMappingCreate(params)
	return err
}

// PCG - OpenStack

// CreatePCGCloudAccountOpenStack creates a new OpenStack PCG cloud account.
func (h *V1Client) CreatePCGCloudAccountOpenStack(overlordUID string, account *models.V1OverlordOpenStackAccountCreate) (string, error) {
	params := clientv1.NewV1OverlordsUIDOpenStackAccountCreateParamsWithContext(h.ctx).
		WithUID(overlordUID).
		WithBody(account)
	resp, err := h.Client.V1OverlordsUIDOpenStackAccountCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

// CreatePCGOpenStack creates a new OpenStack Private Cloud Gateway.
func (h *V1Client) CreatePCGOpenStack(overlordUID string, cloudConfig *models.V1OverlordOpenStackCloudConfig) (string, error) {
	params := clientv1.NewV1OverlordsUIDOpenStackCloudConfigCreateParamsWithContext(h.ctx).
		WithUID(overlordUID).
		WithBody(cloudConfig)
	resp, err := h.Client.V1OverlordsUIDOpenStackCloudConfigCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

// GetPCGManifestOpenStack retrieves an OpenStack PCG manifest by pairing code.
func (h *V1Client) GetPCGManifestOpenStack(pairingCode string) (string, error) {
	params := clientv1.NewV1OverlordsOpenStackManifestParamsWithContext(h.ctx).
		WithPairingCode(pairingCode)
	resp, err := h.Client.V1OverlordsOpenStackManifest(params)
	if err != nil {
		return "", err
	}
	return resp.Payload.Manifest, nil
}

// GetPCGClusterProfileOpenStack retrieves an OpenStack PCG cluster profile by PCG UID.
func (h *V1Client) GetPCGClusterProfileOpenStack(uid string) (*models.V1ClusterProfile, error) {
	params := clientv1.NewV1OverlordsUIDOpenStackClusterProfileParamsWithContext(h.ctx).
		WithUID(uid)
	resp, err := h.Client.V1OverlordsUIDOpenStackClusterProfile(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// PCG - MAAS

// CreatePCGCloudAccountMaas creates a new MAAS PCG cloud account.
func (h *V1Client) CreatePCGCloudAccountMaas(overlordUID string, account *models.V1OverlordMaasAccountCreate) (string, error) {
	params := clientv1.NewV1OverlordsUIDMaasAccountCreateParamsWithContext(h.ctx).
		WithUID(overlordUID).
		WithBody(account)
	resp, err := h.Client.V1OverlordsUIDMaasAccountCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

// CreatePCGMaas creates a new MAAS Private Cloud Gateway.
func (h *V1Client) CreatePCGMaas(overlordUID string, cloudConfig *models.V1OverlordMaasCloudConfig) (string, error) {
	params := clientv1.NewV1OverlordsUIDMaasCloudConfigCreateParamsWithContext(h.ctx).
		WithUID(overlordUID).
		WithBody(cloudConfig)
	resp, err := h.Client.V1OverlordsUIDMaasCloudConfigCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

// GetPCGManifestMaas retrieves a MAAS PCG manifest by pairing code.
func (h *V1Client) GetPCGManifestMaas(pairingCode string) (string, error) {
	params := clientv1.NewV1OverlordsMaasManifestParamsWithContext(h.ctx).
		WithPairingCode(pairingCode)
	resp, err := h.Client.V1OverlordsMaasManifest(params)
	if err != nil {
		return "", err
	}
	return resp.Payload.Manifest, nil
}

// GetPCGClusterProfileMaas retrieves a MAAS PCG cluster profile by PCG UID.
func (h *V1Client) GetPCGClusterProfileMaas(uid string) (*models.V1ClusterProfile, error) {
	params := clientv1.NewV1OverlordsUIDMaasClusterProfileParamsWithContext(h.ctx).
		WithUID(uid)
	resp, err := h.Client.V1OverlordsUIDMaasClusterProfile(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// IP Pool

// CreateIPPool creates a new IP pool for a Private Cloud Gateway.
func (h *V1Client) CreateIPPool(pcgUID string, pool *models.V1IPPoolInputEntity) (string, error) {
	params := clientv1.NewV1OverlordsUIDPoolCreateParamsWithContext(h.ctx).
		WithUID(pcgUID).
		WithBody(pool)
	resp, err := h.Client.V1OverlordsUIDPoolCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

// GetIPPool retrieves a PCG's IP pool by pool UID.
func (h *V1Client) GetIPPool(pcgUID, poolUID string) (*models.V1IPPoolEntity, error) {
	pools, err := h.GetIPPools(pcgUID)
	if err != nil {
		return nil, err
	}
	for _, pool := range pools {
		if pool.Metadata.UID == poolUID {
			return pool, nil
		}
	}
	return nil, nil
}

// GetIPPoolByName retrieves a PCG's IP pool by pool name.
func (h *V1Client) GetIPPoolByName(pcgUID, poolName string) (*models.V1IPPoolEntity, error) {
	pools, err := h.GetIPPools(pcgUID)
	if err != nil {
		return nil, err
	}
	for _, pool := range pools {
		if pool.Metadata.Name == poolName {
			return pool, nil
		}
	}
	return nil, errors.New("ip pool not found: " + poolName)
}

// GetIPPools retrieves all IP pools for a Private Cloud Gateway.
func (h *V1Client) GetIPPools(pcgUID string) ([]*models.V1IPPoolEntity, error) {
	params := clientv1.NewV1OverlordsUIDPoolsListParamsWithContext(h.ctx).
		WithUID(pcgUID)
	resp, err := h.Client.V1OverlordsUIDPoolsList(params)
	if err != nil {
		if v1Err := apiutil.ToV1ErrorObj(err); v1Err.Code != "ResourceNotFound" {
			return nil, err
		}
	}
	return resp.Payload.Items, nil
}

// UpdateIPPool updates an existing Private Cloud Gateway IP pool.
func (h *V1Client) UpdateIPPool(pcgUID, poolUID string, pool *models.V1IPPoolInputEntity) error {
	params := clientv1.NewV1OverlordsUIDPoolUpdateParamsWithContext(h.ctx).
		WithUID(pcgUID).
		WithBody(pool).
		WithPoolUID(poolUID)
	_, err := h.Client.V1OverlordsUIDPoolUpdate(params)
	return err
}

// DeleteIPPool deletes an existing Private Cloud Gateway IP pool.
func (h *V1Client) DeleteIPPool(pcgUID, poolUID string) error {
	params := clientv1.NewV1OverlordsUIDPoolDeleteParamsWithContext(h.ctx).
		WithUID(pcgUID).
		WithPoolUID(poolUID)
	_, err := h.Client.V1OverlordsUIDPoolDelete(params)
	return err
}

// CreateVsphereDNSMap creates a new DNS Mapping for a Private Cloud Gateway.
func (h *V1Client) CreateVsphereDNSMap(dnsMapBody *models.V1VsphereDNSMapping) (string, error) {
	params := clientv1.NewV1VsphereDNSMappingCreateParamsWithContext(h.ctx).WithBody(dnsMapBody)
	resp, err := h.Client.V1VsphereDNSMappingCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

// UpdateVsphereDNSMap update an existing DNS Mapping for a Private Cloud Gateway
func (h *V1Client) UpdateVsphereDNSMap(dnsMapID string, dnsMapBody *models.V1VsphereDNSMapping) error {
	params := clientv1.NewV1VsphereDNSMappingUpdateParamsWithContext(h.ctx).WithUID(dnsMapID).WithBody(dnsMapBody)
	_, err := h.Client.V1VsphereDNSMappingUpdate(params)
	return err
}

// DeleteVsphereDNSMap delete an existing DNS Mapping for a Private Cloud Gateway
func (h *V1Client) DeleteVsphereDNSMap(dnsMapID string) error {
	params := clientv1.NewV1VsphereDNSMappingDeleteParamsWithContext(h.ctx).WithUID(dnsMapID)
	_, err := h.Client.V1VsphereDNSMappingDelete(params)
	if err != nil {
		return err
	}
	return nil
}

// GetVsphereDNSMap get an existing DNS Mapping for a Private Cloud Gateway
func (h *V1Client) GetVsphereDNSMap(dnsMapID string) (*models.V1VsphereDNSMapping, error) {
	params := clientv1.NewV1VsphereDNSMappingGetParamsWithContext(h.ctx).WithUID(dnsMapID)
	resp, err := h.Client.V1VsphereDNSMappingGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// GetVsphereDNSMappingsByPCGId get an existing DNS Mappings for a Private Cloud Gateway with PCG-ID
func (h *V1Client) GetVsphereDNSMappingsByPCGId(PCGId string) (*models.V1VsphereDNSMappings, error) {
	filter := "spec.privateGatewayUid=" + PCGId
	params := clientv1.NewV1VsphereDNSMappingsGetParamsWithContext(h.ctx).WithFilters(&filter)
	resp, err := h.Client.V1VsphereDNSMappingsGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}
