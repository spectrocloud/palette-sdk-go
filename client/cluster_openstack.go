package client

import (
	"github.com/spectrocloud/hapi/models"
	clusterC "github.com/spectrocloud/hapi/spectrocluster/client/v1"

	"github.com/spectrocloud/palette-sdk-go/client/herr"
)

// Cluster

func (h *V1Client) CreateClusterOpenStack(cluster *models.V1SpectroOpenStackClusterEntity, ClusterContext string) (string, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return "", err
	}

	var params *clusterC.V1SpectroClustersOpenStackCreateParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1SpectroClustersOpenStackCreateParamsWithContext(h.Ctx).WithBody(cluster)
	case "tenant":
		params = clusterC.NewV1SpectroClustersOpenStackCreateParams().WithBody(cluster)
	}

	success, err := client.V1SpectroClustersOpenStackCreate(params)
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

// Machine Pool

func (h *V1Client) CreateMachinePoolOpenStack(cloudConfigId string, machinePool *models.V1OpenStackMachinePoolConfigEntity) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil
	}

	params := clusterC.NewV1CloudConfigsOpenStackMachinePoolCreateParamsWithContext(h.Ctx).WithConfigUID(cloudConfigId).WithBody(machinePool)
	_, err = client.V1CloudConfigsOpenStackMachinePoolCreate(params)
	return err
}

func (h *V1Client) UpdateMachinePoolOpenStack(cloudConfigId string, machinePool *models.V1OpenStackMachinePoolConfigEntity) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil
	}

	params := clusterC.NewV1CloudConfigsOpenStackMachinePoolUpdateParamsWithContext(h.Ctx).
		WithConfigUID(cloudConfigId).
		WithMachinePoolName(*machinePool.PoolConfig.Name).
		WithBody(machinePool)
	_, err = client.V1CloudConfigsOpenStackMachinePoolUpdate(params)
	return err
}

func (h *V1Client) DeleteMachinePoolOpenStack(cloudConfigId, machinePoolName string) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil
	}

	params := clusterC.NewV1CloudConfigsOpenStackMachinePoolDeleteParamsWithContext(h.Ctx).WithConfigUID(cloudConfigId).WithMachinePoolName(machinePoolName)
	_, err = client.V1CloudConfigsOpenStackMachinePoolDelete(params)
	return err
}

// Cloud Config

func (h *V1Client) GetCloudConfigOpenStack(configUID string) (*models.V1OpenStackCloudConfig, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}

	params := clusterC.NewV1CloudConfigsOpenStackGetParamsWithContext(h.Ctx).WithConfigUID(configUID)
	success, err := client.V1CloudConfigsOpenStackGet(params)

	if herr.IsNotFound(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return success.Payload, nil
}

// Import

func (h *V1Client) ImportClusterOpenStack(meta *models.V1ObjectMetaInputEntity) (string, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return "", err
	}

	params := clusterC.NewV1SpectroClustersOpenStackImportParamsWithContext(h.Ctx).WithBody(
		&models.V1SpectroOpenStackClusterImportEntity{
			Metadata: meta,
		},
	)
	success, err := client.V1SpectroClustersOpenStackImport(params)
	if err != nil {
		return "", err
	}
	return *success.Payload.UID, nil
}
