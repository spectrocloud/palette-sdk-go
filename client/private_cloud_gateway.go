package client

import (
	"errors"
	"fmt"

	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
	"github.com/spectrocloud/palette-sdk-go/client/apiutil"
)

// PCG - Generic

func (h *V1Client) GetPCGId(name *string) (string, error) {
	params := clientV1.NewV1OverlordsListParamsWithContext(h.ctx)
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

func (h *V1Client) GetPCGById(uid string) (*models.V1Overlord, error) {
	params := clientV1.NewV1OverlordsUIDGetParamsWithContext(h.ctx).
		WithUID(uid)
	resp, err := h.Client.V1OverlordsUIDGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

func (h *V1Client) GetPCGByName(name string) (*models.V1Overlord, error) {
	params := clientV1.NewV1OverlordsListParamsWithContext(h.ctx).
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

func (h *V1Client) GetPairingCode(cloudType string) (string, error) {
	params := clientV1.NewV1OverlordsPairingCodeParamsWithContext(h.ctx).
		WithCloudType(&cloudType)
	resp, err := h.Client.V1OverlordsPairingCode(params)
	if err != nil {
		return "", err
	}
	return resp.Payload.PairingCode, nil
}

func (h *V1Client) CheckPCG(pcgId string) error {
	if pcgId == "" {
		return nil // no pcg to check
	}
	pcg, err := h.GetPCGById(pcgId)
	if err != nil {
		return err
	}
	if pcg == nil {
		return fmt.Errorf("private cloud gateway not found: %s", pcgId)
	}
	if pcg.Status == nil {
		return fmt.Errorf("private cloud gateway status not found: %s", pcgId)
	}
	if pcg.Status.State != "Running" {
		return fmt.Errorf("private cloud gateway is not running: %s", pcgId)
	}
	return nil // pcg is running
}

// PCG - vSphere

func (h *V1Client) CreatePCGVsphere(uid string, cloudConfig *models.V1OverlordVsphereCloudConfig) (string, error) {
	params := clientV1.NewV1OverlordsUIDVsphereCloudConfigCreateParamsWithContext(h.ctx).
		WithUID(uid).
		WithBody(cloudConfig)
	resp, err := h.Client.V1OverlordsUIDVsphereCloudConfigCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil

}

func (h *V1Client) CreatePCGCloudAccountVsphere(uid string, account *models.V1OverlordVsphereAccountCreate) (string, error) {
	params := clientV1.NewV1OverlordsUIDVsphereAccountCreateParamsWithContext(h.ctx).
		WithUID(uid).
		WithBody(account)
	resp, err := h.Client.V1OverlordsUIDVsphereAccountCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

func (h *V1Client) GetPCGManifestVsphere(pairingCode string) (string, error) {
	params := clientV1.NewV1OverlordsVsphereManifestParamsWithContext(h.ctx).
		WithPairingCode(pairingCode)
	resp, err := h.Client.V1OverlordsVsphereManifest(params)
	if err != nil {
		return "", err
	}
	return resp.Payload.Manifest, nil
}

func (h *V1Client) GetPCGClusterProfileVsphere(uid string) (*models.V1ClusterProfile, error) {
	params := clientV1.NewV1OverlordsUIDVsphereClusterProfileParamsWithContext(h.ctx).
		WithUID(uid)
	resp, err := h.Client.V1OverlordsUIDVsphereClusterProfile(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

func (h *V1Client) CreateDDNSSearchDomainVsphere(vsphereDnsMapping *models.V1VsphereDNSMapping) error {
	params := clientV1.NewV1VsphereDNSMappingCreateParamsWithContext(h.ctx).
		WithBody(vsphereDnsMapping)
	_, err := h.Client.V1VsphereDNSMappingCreate(params)
	return err
}

// PCG - OpenStack

func (h *V1Client) CreatePCGCloudAccountOpenStack(overlordUid string, account *models.V1OverlordOpenStackAccountCreate) (string, error) {
	params := clientV1.NewV1OverlordsUIDOpenStackAccountCreateParamsWithContext(h.ctx).
		WithUID(overlordUid).
		WithBody(account)
	resp, err := h.Client.V1OverlordsUIDOpenStackAccountCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

func (h *V1Client) CreatePCGOpenStack(overlordUid string, cloudConfig *models.V1OverlordOpenStackCloudConfig) (string, error) {
	params := clientV1.NewV1OverlordsUIDOpenStackCloudConfigCreateParamsWithContext(h.ctx).
		WithUID(overlordUid).
		WithBody(cloudConfig)
	resp, err := h.Client.V1OverlordsUIDOpenStackCloudConfigCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

func (h *V1Client) GetPCGManifestOpenStack(pairingCode string) (string, error) {
	params := clientV1.NewV1OverlordsOpenStackManifestParamsWithContext(h.ctx).
		WithPairingCode(pairingCode)
	resp, err := h.Client.V1OverlordsOpenStackManifest(params)
	if err != nil {
		return "", err
	}
	return resp.Payload.Manifest, nil
}

func (h *V1Client) GetPCGClusterProfileOpenStack(uid string) (*models.V1ClusterProfile, error) {
	params := clientV1.NewV1OverlordsUIDOpenStackClusterProfileParamsWithContext(h.ctx).
		WithUID(uid)
	resp, err := h.Client.V1OverlordsUIDOpenStackClusterProfile(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// PCG - MAAS

func (h *V1Client) CreatePCGCloudAccountMaas(overlordUid string, account *models.V1OverlordMaasAccountCreate) (string, error) {
	params := clientV1.NewV1OverlordsUIDMaasAccountCreateParamsWithContext(h.ctx).
		WithUID(overlordUid).
		WithBody(account)
	resp, err := h.Client.V1OverlordsUIDMaasAccountCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

func (h *V1Client) CreatePCGMaas(overlordUid string, cloudConfig *models.V1OverlordMaasCloudConfig) (string, error) {
	params := clientV1.NewV1OverlordsUIDMaasCloudConfigCreateParamsWithContext(h.ctx).
		WithUID(overlordUid).
		WithBody(cloudConfig)
	resp, err := h.Client.V1OverlordsUIDMaasCloudConfigCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

func (h *V1Client) GetPCGManifestMaas(pairingCode string) (string, error) {
	params := clientV1.NewV1OverlordsMaasManifestParamsWithContext(h.ctx).
		WithPairingCode(pairingCode)
	resp, err := h.Client.V1OverlordsMaasManifest(params)
	if err != nil {
		return "", err
	}
	return resp.Payload.Manifest, nil
}

func (h *V1Client) GetPCGClusterProfileMaas(uid string) (*models.V1ClusterProfile, error) {
	params := clientV1.NewV1OverlordsUIDMaasClusterProfileParamsWithContext(h.ctx).
		WithUID(uid)
	resp, err := h.Client.V1OverlordsUIDMaasClusterProfile(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// IP Pool

func (h *V1Client) CreateIpPool(pcgUID string, pool *models.V1IPPoolInputEntity) (string, error) {
	params := clientV1.NewV1OverlordsUIDPoolCreateParamsWithContext(h.ctx).
		WithUID(pcgUID).
		WithBody(pool)
	resp, err := h.Client.V1OverlordsUIDPoolCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

func (h *V1Client) GetIpPool(pcgUid, poolUid string) (*models.V1IPPoolEntity, error) {
	pools, err := h.GetIpPools(pcgUid)
	if err != nil {
		return nil, err
	}
	for _, pool := range pools {
		if pool.Metadata.UID == poolUid {
			return pool, nil
		}
	}
	return nil, nil
}

func (h *V1Client) GetIpPoolByName(pcgUid, poolName string) (*models.V1IPPoolEntity, error) {
	pools, err := h.GetIpPools(pcgUid)
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

func (h *V1Client) GetIpPools(pcgUid string) ([]*models.V1IPPoolEntity, error) {
	params := clientV1.NewV1OverlordsUIDPoolsListParamsWithContext(h.ctx).
		WithUID(pcgUid)
	resp, err := h.Client.V1OverlordsUIDPoolsList(params)
	if err != nil {
		if v1Err := apiutil.ToV1ErrorObj(err); v1Err.Code != "ResourceNotFound" {
			return nil, err
		}
	}
	return resp.Payload.Items, nil
}

func (h *V1Client) UpdateIpPool(pcgUid, poolUid string, pool *models.V1IPPoolInputEntity) error {
	params := clientV1.NewV1OverlordsUIDPoolUpdateParamsWithContext(h.ctx).
		WithUID(pcgUid).
		WithBody(pool).
		WithPoolUID(poolUid)
	_, err := h.Client.V1OverlordsUIDPoolUpdate(params)
	return err
}

func (h *V1Client) DeleteIpPool(pcgUid, poolUid string) error {
	params := clientV1.NewV1OverlordsUIDPoolDeleteParamsWithContext(h.ctx).
		WithUID(pcgUid).
		WithPoolUID(poolUid)
	_, err := h.Client.V1OverlordsUIDPoolDelete(params)
	return err
}
