package client

import (
	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
	"github.com/spectrocloud/palette-sdk-go/client/apiutil"
)

func (h *V1Client) CreateClusterAzure(cluster *models.V1SpectroAzureClusterEntity) (string, error) {
	params := clientV1.NewV1SpectroClustersAzureCreateParamsWithContext(h.ctx).
		WithBody(cluster)
	resp, err := h.Client.V1SpectroClustersAzureCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

func (h *V1Client) CreateMachinePoolAzure(cloudConfigUID string, machinePool *models.V1AzureMachinePoolConfigEntity) error {
	params := clientV1.NewV1CloudConfigsAzureMachinePoolCreateParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUID).
		WithBody(machinePool)
	_, err := h.Client.V1CloudConfigsAzureMachinePoolCreate(params)
	return err
}

func (h *V1Client) UpdateMachinePoolAzure(cloudConfigUID string, machinePool *models.V1AzureMachinePoolConfigEntity) error {
	params := clientV1.NewV1CloudConfigsAzureMachinePoolUpdateParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUID).
		WithMachinePoolName(*machinePool.PoolConfig.Name).
		WithBody(machinePool)
	_, err := h.Client.V1CloudConfigsAzureMachinePoolUpdate(params)
	return err
}

func (h *V1Client) DeleteMachinePoolAzure(cloudConfigUID, machinePoolName string) error {
	params := clientV1.NewV1CloudConfigsAzureMachinePoolDeleteParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUID).
		WithMachinePoolName(machinePoolName)
	_, err := h.Client.V1CloudConfigsAzureMachinePoolDelete(params)
	return err
}

func (h *V1Client) GetCloudConfigAzure(configUID string) (*models.V1AzureCloudConfig, error) {
	params := clientV1.NewV1CloudConfigsAzureGetParamsWithContext(h.ctx).
		WithConfigUID(configUID)
	resp, err := h.Client.V1CloudConfigsAzureGet(params)
	if apiutil.Is404(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

func (h *V1Client) ImportClusterAzure(meta *models.V1ObjectMetaInputEntity) (string, error) {
	params := clientV1.NewV1SpectroClustersAzureImportParamsWithContext(h.ctx).
		WithBody(&models.V1SpectroAzureClusterImportEntity{
			Metadata: meta,
		},
		)
	resp, err := h.Client.V1SpectroClustersAzureImport(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

func (h *V1Client) GetNodeStatusMapAzure(configUID, machinePoolName string) (map[string]models.V1CloudMachineStatus, error) {
	params := clientV1.NewV1CloudConfigsAzurePoolMachinesListParamsWithContext(h.ctx).
		WithConfigUID(configUID).
		WithMachinePoolName(machinePoolName)
	mpList, err := h.Client.V1CloudConfigsAzurePoolMachinesList(params)
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
