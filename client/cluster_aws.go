package client

import (
	"errors"

	"github.com/spectrocloud/palette-api-go/apiutil/transport"
	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
)

func (h *V1Client) CreateClusterAws(cluster *models.V1SpectroAwsClusterEntity, scope string) (string, error) {
	var params *clientV1.V1SpectroClustersAwsCreateParams
	switch scope {
	case "project":
		params = clientV1.NewV1SpectroClustersAwsCreateParamsWithContext(h.Ctx).WithBody(cluster)
	case "tenant":
		params = clientV1.NewV1SpectroClustersAwsCreateParams().WithBody(cluster)
	}

	success, err := h.GetClient().V1SpectroClustersAwsCreate(params)
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

func (h *V1Client) CreateMachinePoolAws(CloudConfigId string, machinePool *models.V1AwsMachinePoolConfigEntity, scope string) error {
	var params *clientV1.V1CloudConfigsAwsMachinePoolCreateParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsAwsMachinePoolCreateParamsWithContext(h.Ctx).WithConfigUID(CloudConfigId).WithBody(machinePool)
	case "tenant":
		params = clientV1.NewV1CloudConfigsAwsMachinePoolCreateParams().WithConfigUID(CloudConfigId).WithBody(machinePool)
	}

	_, err := h.GetClient().V1CloudConfigsAwsMachinePoolCreate(params)
	return err
}

func (h *V1Client) UpdateMachinePoolAws(CloudConfigId string, machinePool *models.V1AwsMachinePoolConfigEntity, scope string) error {
	var params *clientV1.V1CloudConfigsAwsMachinePoolUpdateParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsAwsMachinePoolUpdateParamsWithContext(h.Ctx).WithConfigUID(CloudConfigId).
			WithMachinePoolName(*machinePool.PoolConfig.Name).
			WithBody(machinePool)
	case "tenant":
		params = clientV1.NewV1CloudConfigsAwsMachinePoolUpdateParams().WithConfigUID(CloudConfigId).
			WithMachinePoolName(*machinePool.PoolConfig.Name).
			WithBody(machinePool)
	}

	_, err := h.GetClient().V1CloudConfigsAwsMachinePoolUpdate(params)
	return err
}

func (h *V1Client) DeleteMachinePoolAws(CloudConfigId, machinePoolName, scope string) error {
	var params *clientV1.V1CloudConfigsAwsMachinePoolDeleteParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsAwsMachinePoolDeleteParamsWithContext(h.Ctx).WithConfigUID(CloudConfigId).WithMachinePoolName(machinePoolName)
	case "tenant":
		params = clientV1.NewV1CloudConfigsAwsMachinePoolDeleteParams().WithConfigUID(CloudConfigId).WithMachinePoolName(machinePoolName)
	}

	_, err := h.GetClient().V1CloudConfigsAwsMachinePoolDelete(params)
	return err
}

func (h *V1Client) GetCloudConfigAws(configUID, scope string) (*models.V1AwsCloudConfig, error) {
	var params *clientV1.V1CloudConfigsAwsGetParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsAwsGetParamsWithContext(h.Ctx).WithConfigUID(configUID)
	case "tenant":
		params = clientV1.NewV1CloudConfigsAwsGetParams().WithConfigUID(configUID)
	}

	success, err := h.GetClient().V1CloudConfigsAwsGet(params)
	var e *transport.TransportError
	if errors.As(err, &e) && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return success.Payload, nil
}

func (h *V1Client) ImportClusterAws(meta *models.V1ObjectMetaInputEntity) (string, error) {
	params := clientV1.NewV1SpectroClustersAwsImportParamsWithContext(h.Ctx).WithBody(
		&models.V1SpectroAwsClusterImportEntity{
			Metadata: meta,
		},
	)
	success, err := h.GetClient().V1SpectroClustersAwsImport(params)
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

func (h *V1Client) GetNodeStatusMapAws(configUID, machinePoolName, scope string) (map[string]models.V1CloudMachineStatus, error) {
	var params *clientV1.V1CloudConfigsAwsPoolMachinesListParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsAwsPoolMachinesListParamsWithContext(h.Ctx).WithConfigUID(configUID).WithMachinePoolName(machinePoolName)
	case "tenant":
		params = clientV1.NewV1CloudConfigsAwsPoolMachinesListParams().WithConfigUID(configUID).WithMachinePoolName(machinePoolName)
	}

	mpList, err := h.GetClient().V1CloudConfigsAwsPoolMachinesList(params)
	if err != nil {
		return nil, err
	}
	nMap := map[string]models.V1CloudMachineStatus{}
	if len(mpList.Payload.Items) > 0 {
		for _, node := range mpList.Payload.Items {
			if node.Status.MaintenanceStatus.Action != "" {
				nMap[node.Metadata.UID] = *node.Status
			}
		}
	}
	return nMap, err
}
