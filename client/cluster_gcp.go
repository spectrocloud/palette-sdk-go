package client

import (
	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
	"github.com/spectrocloud/palette-sdk-go/client/apiutil"
)

func (h *V1Client) CreateClusterGcp(cluster *models.V1SpectroGcpClusterEntity) (string, error) {
	params := clientV1.NewV1SpectroClustersGcpCreateParamsWithContext(h.ctx).
		WithBody(cluster)
	resp, err := h.Client.V1SpectroClustersGcpCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

func (h *V1Client) CreateMachinePoolGcp(cloudConfigUid string, machinePool *models.V1GcpMachinePoolConfigEntity) error {
	params := clientV1.NewV1CloudConfigsGcpMachinePoolCreateParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUid).
		WithBody(machinePool)
	_, err := h.Client.V1CloudConfigsGcpMachinePoolCreate(params)
	return err
}

func (h *V1Client) UpdateMachinePoolGcp(cloudConfigUid string, machinePool *models.V1GcpMachinePoolConfigEntity) error {
	params := clientV1.NewV1CloudConfigsGcpMachinePoolUpdateParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUid).
		WithMachinePoolName(*machinePool.PoolConfig.Name).
		WithBody(machinePool)
	_, err := h.Client.V1CloudConfigsGcpMachinePoolUpdate(params)
	return err
}

func (h *V1Client) DeleteMachinePoolGcp(cloudConfigUid, machinePoolName string) error {
	params := clientV1.NewV1CloudConfigsGcpMachinePoolDeleteParamsWithContext(h.ctx).
		WithConfigUID(cloudConfigUid).
		WithMachinePoolName(machinePoolName)
	_, err := h.Client.V1CloudConfigsGcpMachinePoolDelete(params)
	return err
}

func (h *V1Client) GetCloudConfigGcp(configUid string) (*models.V1GcpCloudConfig, error) {
	params := clientV1.NewV1CloudConfigsGcpGetParamsWithContext(h.ctx).
		WithConfigUID(configUid)
	resp, err := h.Client.V1CloudConfigsGcpGet(params)
	if err := apiutil.Handle404(err); err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

func (h *V1Client) ImportClusterGcp(meta *models.V1ObjectMetaInputEntity) (string, error) {
	params := clientV1.NewV1SpectroClustersGcpImportParamsWithContext(h.ctx).
		WithBody(&models.V1SpectroGcpClusterImportEntity{
			Metadata: meta,
		},
		)
	resp, err := h.Client.V1SpectroClustersGcpImport(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

func (h *V1Client) GetNodeStatusMapGcp(configUid, machinePoolName string) (map[string]models.V1CloudMachineStatus, error) {
	params := clientV1.NewV1CloudConfigsGcpPoolMachinesListParamsWithContext(h.ctx).
		WithConfigUID(configUid).
		WithMachinePoolName(machinePoolName)
	mpList, err := h.Client.V1CloudConfigsGcpPoolMachinesList(params)
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
