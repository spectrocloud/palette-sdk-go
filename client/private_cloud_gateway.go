package client

import (
	"errors"
	"fmt"

	"github.com/spectrocloud/hapi/apiutil"
	"github.com/spectrocloud/hapi/models"
	clusterC "github.com/spectrocloud/hapi/spectrocluster/client/v1"
	userC "github.com/spectrocloud/hapi/user/client/v1"
)

// PCG - Generic

func (h *V1Client) GetPCGId(name *string) (string, error) {
	params := clusterC.NewV1OverlordsListParams()
	listResp, err := h.GetClusterClient().V1OverlordsList(params)
	if err != nil {
		return "", err
	}

	for _, pcg := range listResp.Payload.Items {
		if pcg.Metadata.Name == *name {
			return pcg.Metadata.UID, nil
		}
	}
	return "", errors.New("Private Cloud Gateway not found: " + *name)
}

func (h *V1Client) GetPCGById(uid string) (*models.V1Overlord, error) {
	params := clusterC.NewV1OverlordsUIDGetParams().WithUID(uid)
	overlord, err := h.GetClusterClient().V1OverlordsUIDGet(params)
	if err != nil {
		return nil, err
	}
	return overlord.Payload, nil
}

func (h *V1Client) GetPCGByName(name *string) (*models.V1Overlord, error) {
	params := clusterC.NewV1OverlordsListParamsWithContext(h.Ctx).WithName(name)
	overlordList, err := h.GetClusterClient().V1OverlordsList(params)
	if err != nil {
		return nil, err
	}

	var overlord *models.V1Overlord
	if len(overlordList.Payload.Items) > 0 {
		overlord = overlordList.Payload.Items[0]
	}
	return overlord, nil
}

func (h *V1Client) GetPairingCode(cloudType string) (string, error) {
	codeParams := clusterC.NewV1OverlordsPairingCodeParams().WithContext(h.Ctx).WithCloudType(&cloudType)
	ret, err := h.GetClusterClient().V1OverlordsPairingCode(codeParams)
	if err != nil {
		return "", err
	}

	return ret.Payload.PairingCode, nil
}

func (h *V1Client) CheckPCG(PcgId string) error {
	if PcgId == "" {
		return nil // no pcg to check
	}
	pcg, err := h.GetPCGById(PcgId)
	if err != nil {
		return err
	}
	if pcg == nil {
		return fmt.Errorf("Private Cloud Gateway not found: %s", PcgId)
	}
	if pcg.Status == nil {
		return fmt.Errorf("Private Cloud Gateway status not found: %s", PcgId)
	}
	if pcg.Status.State != "Running" {
		return fmt.Errorf("Private Cloud Gateway is not running: %s", PcgId)
	} else {
		return nil // pcg is running
	}
}

// PCG - vSphere

func (h *V1Client) CreatePCGVsphere(uid string, cloudConfig *models.V1OverlordVsphereCloudConfig) (string, error) {
	params := clusterC.NewV1OverlordsUIDVsphereCloudConfigCreateParamsWithContext(h.Ctx).WithBody(cloudConfig).WithUID(uid)
	success, err := h.GetClusterClient().V1OverlordsUIDVsphereCloudConfigCreate(params)
	if err != nil {
		return "", err
	}
	return *success.Payload.UID, nil

}

func (h *V1Client) CreatePCGCloudAccountVsphere(uid string, account *models.V1OverlordVsphereAccountCreate) (string, error) {
	params := clusterC.NewV1OverlordsUIDVsphereAccountCreateParamsWithContext(h.Ctx).WithBody(account).WithUID(uid)
	success, err := h.GetClusterClient().V1OverlordsUIDVsphereAccountCreate(params)
	if err != nil {
		return "", err
	}
	return *success.Payload.UID, nil
}

func (h *V1Client) GetPCGManifestVsphere(pairingCode string) (string, error) {
	params := clusterC.NewV1OverlordsVsphereManifestParamsWithContext(h.Ctx).WithPairingCode(pairingCode)
	success, err := h.GetClusterClient().V1OverlordsVsphereManifest(params)
	if err != nil {
		return "", err
	}
	return success.Payload.Manifest, nil
}

func (h *V1Client) GetPCGClusterProfileVsphere(uid string) (*models.V1ClusterProfile, error) {
	params := clusterC.NewV1OverlordsUIDVsphereClusterProfileParams().WithUID(uid)
	resp, err := h.GetClusterClient().V1OverlordsUIDVsphereClusterProfile(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

func (h *V1Client) CreateDDNSSearchDomainVsphere(vsphereDnsMapping *models.V1VsphereDNSMapping) error {
	params := userC.NewV1VsphereDNSMappingCreateParamsWithContext(h.Ctx).WithBody(vsphereDnsMapping)
	if _, err := h.GetUserClient().V1VsphereDNSMappingCreate(params); err != nil {
		return err
	}
	return nil
}

// PCG - OpenStack

func (h *V1Client) CreatePCGCloudAccountOpenStack(overlordUid string, account *models.V1OverlordOpenStackAccountCreate) (string, error) {
	params := clusterC.NewV1OverlordsUIDOpenStackAccountCreateParamsWithContext(h.Ctx).
		WithBody(account).
		WithUID(overlordUid)
	success, err := h.GetClusterClient().V1OverlordsUIDOpenStackAccountCreate(params)
	if err != nil {
		return "", err
	}
	return *success.Payload.UID, nil
}

func (h *V1Client) CreatePCGOpenStack(overlordUid string, cloudConfig *models.V1OverlordOpenStackCloudConfig) (string, error) {
	params := clusterC.NewV1OverlordsUIDOpenStackCloudConfigCreateParamsWithContext(h.Ctx).
		WithBody(cloudConfig).
		WithUID(overlordUid)
	success, err := h.GetClusterClient().V1OverlordsUIDOpenStackCloudConfigCreate(params)
	if err != nil {
		return "", err
	}
	return *success.Payload.UID, nil
}

func (h *V1Client) GetPCGManifestOpenStack(pairingCode string) (string, error) {
	params := clusterC.NewV1OverlordsOpenStackManifestParamsWithContext(h.Ctx).WithPairingCode(pairingCode)
	success, err := h.GetClusterClient().V1OverlordsOpenStackManifest(params)
	if err != nil {
		return "", err
	}
	return success.Payload.Manifest, nil
}

func (h *V1Client) GetPCGClusterProfileOpenStack(uid string) (*models.V1ClusterProfile, error) {
	params := clusterC.NewV1OverlordsUIDOpenStackClusterProfileParams().WithUID(uid)
	resp, err := h.GetClusterClient().V1OverlordsUIDOpenStackClusterProfile(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// PCG - MAAS

func (h *V1Client) CreatePCGCloudAccountMaas(overlordUid string, account *models.V1OverlordMaasAccountCreate) (string, error) {
	params := clusterC.NewV1OverlordsUIDMaasAccountCreateParamsWithContext(h.Ctx).
		WithBody(account).
		WithUID(overlordUid)
	success, err := h.GetClusterClient().V1OverlordsUIDMaasAccountCreate(params)
	if err != nil {
		return "", err
	}
	return *success.Payload.UID, nil
}

func (h *V1Client) CreatePCGMaas(overlordUid string, cloudConfig *models.V1OverlordMaasCloudConfig) (string, error) {
	params := clusterC.NewV1OverlordsUIDMaasCloudConfigCreateParamsWithContext(h.Ctx).
		WithBody(cloudConfig).
		WithUID(overlordUid)
	success, err := h.GetClusterClient().V1OverlordsUIDMaasCloudConfigCreate(params)
	if err != nil {
		return "", err
	}
	return *success.Payload.UID, nil
}

func (h *V1Client) GetPCGManifestMaas(pairingCode string) (string, error) {
	params := clusterC.NewV1OverlordsMaasManifestParamsWithContext(h.Ctx).WithPairingCode(pairingCode)
	success, err := h.GetClusterClient().V1OverlordsMaasManifest(params)
	if err != nil {
		return "", err
	}
	return success.Payload.Manifest, nil
}

func (h *V1Client) GetPCGClusterProfileMaas(uid string) (*models.V1ClusterProfile, error) {
	params := clusterC.NewV1OverlordsUIDMaasClusterProfileParams().WithUID(uid)
	resp, err := h.GetClusterClient().V1OverlordsUIDMaasClusterProfile(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// IP Pool

func (h *V1Client) CreateIpPool(pcgUID string, pool *models.V1IPPoolInputEntity) (string, error) {
	params := clusterC.NewV1OverlordsUIDPoolCreateParams().WithUID(pcgUID).WithBody(pool)
	if resp, err := h.GetClusterClient().V1OverlordsUIDPoolCreate(params); err != nil {
		return "", err
	} else {
		return *resp.Payload.UID, nil
	}
}

func (h *V1Client) GetIpPool(pcgUID, poolUID string) (*models.V1IPPoolEntity, error) {
	pools, err := h.GetIpPools(pcgUID)
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

func (h *V1Client) GetIpPoolByName(pcgUID, poolName string) (*models.V1IPPoolEntity, error) {
	pools, err := h.GetIpPools(pcgUID)
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

func (h *V1Client) GetIpPools(pcgUID string) ([]*models.V1IPPoolEntity, error) {
	params := clusterC.NewV1OverlordsUIDPoolsListParams().WithUID(pcgUID)
	listResp, err := h.GetClusterClient().V1OverlordsUIDPoolsList(params)
	if err != nil {
		if v1Err := apiutil.ToV1ErrorObj(err); v1Err.Code != "ResourceNotFound" {
			return nil, err
		}
	}
	return listResp.Payload.Items, nil
}

func (h *V1Client) UpdateIpPool(pcgUID, poolUID string, pool *models.V1IPPoolInputEntity) error {
	params := clusterC.NewV1OverlordsUIDPoolUpdateParams().
		WithUID(pcgUID).
		WithBody(pool).
		WithPoolUID(poolUID)

	_, err := h.GetClusterClient().V1OverlordsUIDPoolUpdate(params)
	return err
}

func (h *V1Client) DeleteIpPool(pcgUID, poolUID string) error {
	params := clusterC.NewV1OverlordsUIDPoolDeleteParams().WithUID(pcgUID).WithPoolUID(poolUID)
	_, err := h.GetClusterClient().V1OverlordsUIDPoolDelete(params)
	return err
}
