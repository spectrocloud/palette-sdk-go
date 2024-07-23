package client

import (
	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
	"github.com/spectrocloud/palette-sdk-go/client/apiutil"
)

// Cluster

func (h *V1Client) CreateClusterMaas(cluster *models.V1SpectroMaasClusterEntity) (string, error) {
	params := clientV1.NewV1SpectroClustersMaasCreateParamsWithContext(h.ctx).
		WithBody(cluster)
	resp, err := h.Client.V1SpectroClustersMaasCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

// Machine Pool

func (h *V1Client) CreateMachinePoolMaas(cloudConfigUID string, machinePool *models.V1MaasMachinePoolConfigEntity) error {
	params := clientV1.NewV1CloudConfigsMaasMachinePoolCreateParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUID).
		WithBody(machinePool)
	_, err := h.Client.V1CloudConfigsMaasMachinePoolCreate(params)
	return err
}

func (h *V1Client) UpdateMachinePoolMaas(cloudConfigUID string, machinePool *models.V1MaasMachinePoolConfigEntity) error {
	params := clientV1.NewV1CloudConfigsMaasMachinePoolUpdateParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUID).
		WithMachinePoolName(*machinePool.PoolConfig.Name).
		WithBody(machinePool)
	_, err := h.Client.V1CloudConfigsMaasMachinePoolUpdate(params)
	return err
}

func (h *V1Client) DeleteMachinePoolMaas(cloudConfigUID, machinePoolName string) error {
	params := clientV1.NewV1CloudConfigsMaasMachinePoolDeleteParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUID).
		WithMachinePoolName(machinePoolName)
	_, err := h.Client.V1CloudConfigsMaasMachinePoolDelete(params)
	return err
}

// Cloud Config

func (h *V1Client) GetCloudConfigMaas(configUID string) (*models.V1MaasCloudConfig, error) {
	params := clientV1.NewV1CloudConfigsMaasGetParamsWithContext(h.ctx).
		WithConfigUID(configUID)
	resp, err := h.Client.V1CloudConfigsMaasGet(params)
	if apiutil.Is404(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// Import

func (h *V1Client) ImportClusterMaas(meta *models.V1ObjectMetaInputEntity) (string, error) {
	params := clientV1.NewV1SpectroClustersMaasImportParamsWithContext(h.ctx).
		WithBody(&models.V1SpectroMaasClusterImportEntity{
			Metadata: meta,
		},
		)
	resp, err := h.Client.V1SpectroClustersMaasImport(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

func (h *V1Client) GetNodeStatusMapMaas(configUID, machinePoolName string) (map[string]models.V1CloudMachineStatus, error) {
	params := clientV1.NewV1CloudConfigsMaasPoolMachinesListParamsWithContext(h.ctx).
		WithConfigUID(configUID).
		WithMachinePoolName(machinePoolName)
	mpList, err := h.Client.V1CloudConfigsMaasPoolMachinesList(params)
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
