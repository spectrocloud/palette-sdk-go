package client

import (
	"time"

	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
	"github.com/spectrocloud/palette-sdk-go/client/apiutil"
)

func (h *V1Client) CreateClusterEks(cluster *models.V1SpectroEksClusterEntity) (string, error) {
	params := clientV1.NewV1SpectroClustersEksCreateParamsWithContext(h.ctx).
		WithBody(cluster).
		WithTimeout(90 * time.Second)
	resp, err := h.Client.V1SpectroClustersEksCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

func (h *V1Client) CreateMachinePoolEks(cloudConfigUID string, machinePool *models.V1EksMachinePoolConfigEntity) error {
	params := clientV1.NewV1CloudConfigsEksMachinePoolCreateParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUID).
		WithBody(machinePool)
	_, err := h.Client.V1CloudConfigsEksMachinePoolCreate(params)
	return err
}

func (h *V1Client) UpdateMachinePoolEks(cloudConfigUID string, machinePool *models.V1EksMachinePoolConfigEntity) error {
	params := clientV1.NewV1CloudConfigsEksMachinePoolUpdateParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUID).
		WithMachinePoolName(*machinePool.PoolConfig.Name).
		WithBody(machinePool)
	_, err := h.Client.V1CloudConfigsEksMachinePoolUpdate(params)
	return err
}

func (h *V1Client) DeleteMachinePoolEks(cloudConfigUID, machinePoolName string) error {
	params := clientV1.NewV1CloudConfigsEksMachinePoolDeleteParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUID).
		WithMachinePoolName(machinePoolName)
	_, err := h.Client.V1CloudConfigsEksMachinePoolDelete(params)
	return err
}

func (h *V1Client) UpdateFargateProfilesEks(cloudConfigUID string, fargateProfiles *models.V1EksFargateProfiles) error {
	params := clientV1.NewV1CloudConfigsEksUIDFargateProfilesUpdateParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUID).
		WithBody(fargateProfiles)
	_, err := h.Client.V1CloudConfigsEksUIDFargateProfilesUpdate(params)
	return err
}

func (h *V1Client) GetCloudConfigEks(configUID string) (*models.V1EksCloudConfig, error) {
	params := clientV1.NewV1CloudConfigsEksGetParamsWithContext(h.ctx).
		WithConfigUID(configUID)
	resp, err := h.Client.V1CloudConfigsEksGet(params)
	if apiutil.Is404(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

func (h *V1Client) GetNodeStatusMapEks(configUID, machinePoolName string) (map[string]models.V1CloudMachineStatus, error) {
	params := clientV1.NewV1CloudConfigsEksPoolMachinesListParamsWithContext(h.ctx).
		WithConfigUID(configUID).
		WithMachinePoolName(machinePoolName)
	mpList, err := h.Client.V1CloudConfigsEksPoolMachinesList(params)
	if err != nil {
		return nil, err
	}
	nMap := map[string]models.V1CloudMachineStatus{}
	if len(mpList.Payload.Items) > 0 {
		for _, node := range mpList.Payload.Items {
			nMap[node.Metadata.UID] = *node.Status
		}
	}
	return nMap, nil
}
