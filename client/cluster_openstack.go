package client

import (
	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
	"github.com/spectrocloud/palette-sdk-go/client/herr"
)

// Cluster

func (h *V1Client) CreateClusterOpenStack(cluster *models.V1SpectroOpenStackClusterEntity) (string, error) {
	params := clientV1.NewV1SpectroClustersOpenStackCreateParamsWithContext(h.ctx).
		WithBody(cluster)
	resp, err := h.Client.V1SpectroClustersOpenStackCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

// Machine Pool

func (h *V1Client) CreateMachinePoolOpenStack(cloudConfigUID string, machinePool *models.V1OpenStackMachinePoolConfigEntity) error {
	params := clientV1.NewV1CloudConfigsOpenStackMachinePoolCreateParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUID).
		WithBody(machinePool)
	_, err := h.Client.V1CloudConfigsOpenStackMachinePoolCreate(params)
	return err
}

func (h *V1Client) UpdateMachinePoolOpenStack(cloudConfigUID string, machinePool *models.V1OpenStackMachinePoolConfigEntity) error {
	params := clientV1.NewV1CloudConfigsOpenStackMachinePoolUpdateParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUID).
		WithMachinePoolName(*machinePool.PoolConfig.Name).
		WithBody(machinePool)
	_, err := h.Client.V1CloudConfigsOpenStackMachinePoolUpdate(params)
	return err
}

func (h *V1Client) DeleteMachinePoolOpenStack(cloudConfigUID, machinePoolName string) error {
	params := clientV1.NewV1CloudConfigsOpenStackMachinePoolDeleteParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUID).
		WithMachinePoolName(machinePoolName)
	_, err := h.Client.V1CloudConfigsOpenStackMachinePoolDelete(params)
	return err
}

// Cloud Config

func (h *V1Client) GetCloudConfigOpenStack(configUID string) (*models.V1OpenStackCloudConfig, error) {
	params := clientV1.NewV1CloudConfigsOpenStackGetParamsWithContext(h.ctx).
		WithConfigUID(configUID)
	resp, err := h.Client.V1CloudConfigsOpenStackGet(params)
	if herr.IsNotFound(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// Import

func (h *V1Client) ImportClusterOpenStack(meta *models.V1ObjectMetaInputEntity) (string, error) {
	params := clientV1.NewV1SpectroClustersOpenStackImportParamsWithContext(h.ctx).
		WithBody(&models.V1SpectroOpenStackClusterImportEntity{
			Metadata: meta,
		},
		)
	resp, err := h.Client.V1SpectroClustersOpenStackImport(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

func (h *V1Client) GetNodeStatusMapOpenStack(configUID, machinePoolName string) (map[string]models.V1CloudMachineStatus, error) {
	params := clientV1.NewV1CloudConfigsOpenStackPoolMachinesListParamsWithContext(h.ctx).
		WithConfigUID(configUID).
		WithMachinePoolName(machinePoolName)
	mpList, err := h.Client.V1CloudConfigsOpenStackPoolMachinesList(params)
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
