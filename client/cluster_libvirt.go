package client

import (
	"errors"

	"github.com/spectrocloud/palette-api-go/apiutil/transport"
	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
)

func (h *V1Client) CreateClusterLibvirt(cluster *models.V1SpectroLibvirtClusterEntity, scope string) (string, error) {
	var params *clientV1.V1SpectroClustersLibvirtCreateParams
	switch scope {
	case "project":
		params = clientV1.NewV1SpectroClustersLibvirtCreateParamsWithContext(h.Ctx).WithBody(cluster)
	case "tenant":
		params = clientV1.NewV1SpectroClustersLibvirtCreateParams().WithBody(cluster)
	}

	success, err := h.GetClient().V1SpectroClustersLibvirtCreate(params)
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

func (h *V1Client) CreateMachinePoolLibvirt(CloudConfigId, scope string, machinePool *models.V1LibvirtMachinePoolConfigEntity) error {
	var params *clientV1.V1CloudConfigsLibvirtMachinePoolCreateParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsLibvirtMachinePoolCreateParamsWithContext(h.Ctx).WithConfigUID(CloudConfigId).WithBody(machinePool)
	case "tenant":
		params = clientV1.NewV1CloudConfigsLibvirtMachinePoolCreateParams().WithConfigUID(CloudConfigId).WithBody(machinePool)
	}

	_, err := h.GetClient().V1CloudConfigsLibvirtMachinePoolCreate(params)
	return err
}

func (h *V1Client) UpdateMachinePoolLibvirt(CloudConfigId, scope string, machinePool *models.V1LibvirtMachinePoolConfigEntity) error {
	var params *clientV1.V1CloudConfigsLibvirtMachinePoolUpdateParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsLibvirtMachinePoolUpdateParamsWithContext(h.Ctx).
			WithConfigUID(CloudConfigId).
			WithMachinePoolName(*machinePool.PoolConfig.Name).
			WithBody(machinePool)
	case "tenant":
		params = clientV1.NewV1CloudConfigsLibvirtMachinePoolUpdateParams().
			WithConfigUID(CloudConfigId).
			WithMachinePoolName(*machinePool.PoolConfig.Name).
			WithBody(machinePool)
	}

	_, err := h.GetClient().V1CloudConfigsLibvirtMachinePoolUpdate(params)
	return err
}

func (h *V1Client) DeleteMachinePoolLibvirt(CloudConfigId, machinePoolName, scope string) error {
	var params *clientV1.V1CloudConfigsLibvirtMachinePoolDeleteParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsLibvirtMachinePoolDeleteParamsWithContext(h.Ctx).WithConfigUID(CloudConfigId).WithMachinePoolName(machinePoolName)
	case "tenant":
		params = clientV1.NewV1CloudConfigsLibvirtMachinePoolDeleteParams().WithConfigUID(CloudConfigId).WithMachinePoolName(machinePoolName)
	}

	_, err := h.GetClient().V1CloudConfigsLibvirtMachinePoolDelete(params)
	return err
}

func (h *V1Client) GetCloudConfigLibvirt(configUID, scope string) (*models.V1LibvirtCloudConfig, error) {
	var params *clientV1.V1CloudConfigsLibvirtGetParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsLibvirtGetParamsWithContext(h.Ctx).WithConfigUID(configUID)
	case "tenant":
		params = clientV1.NewV1CloudConfigsLibvirtGetParams().WithConfigUID(configUID)
	}

	success, err := h.GetClient().V1CloudConfigsLibvirtGet(params)

	var e *transport.TransportError
	if errors.As(err, &e) && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return success.Payload, nil
}

func (h *V1Client) GetNodeStatusMapLibvirt(configUID, machinePoolName, scope string) (map[string]models.V1CloudMachineStatus, error) {
	var params *clientV1.V1CloudConfigsLibvirtPoolMachinesListParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsLibvirtPoolMachinesListParamsWithContext(h.Ctx).WithConfigUID(configUID).WithMachinePoolName(machinePoolName)
	case "tenant":
		params = clientV1.NewV1CloudConfigsLibvirtPoolMachinesListParams().WithConfigUID(configUID).WithMachinePoolName(machinePoolName)
	}

	mpList, err := h.GetClient().V1CloudConfigsLibvirtPoolMachinesList(params)
	nMap := map[string]models.V1CloudMachineStatus{}
	if len(mpList.Payload.Items) > 0 {
		for _, node := range mpList.Payload.Items {
			nMap[node.Metadata.UID] = *node.Status
		}
	}
	return nMap, err
}
