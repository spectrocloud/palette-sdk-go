package client

import (
	"errors"
	"time"

	"github.com/spectrocloud/palette-api-go/apiutil/transport"
	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
)

func (h *V1Client) CreateClusterEks(cluster *models.V1SpectroEksClusterEntity, scope string) (string, error) {
	var params *clientV1.V1SpectroClustersEksCreateParams
	switch scope {
	case "project":
		params = clientV1.NewV1SpectroClustersEksCreateParamsWithContext(h.Ctx).WithBody(cluster)
	case "tenant":
		params = clientV1.NewV1SpectroClustersEksCreateParams().WithBody(cluster)
	}

	success, err := h.Client.V1SpectroClustersEksCreate(params.WithTimeout(90 * time.Second))
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

func (h *V1Client) CreateMachinePoolEks(CloudConfigId, scope string, machinePool *models.V1EksMachinePoolConfigEntity) error {
	var params *clientV1.V1CloudConfigsEksMachinePoolCreateParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsEksMachinePoolCreateParamsWithContext(h.Ctx).WithConfigUID(CloudConfigId).WithBody(machinePool)
	case "tenant":
		params = clientV1.NewV1CloudConfigsEksMachinePoolCreateParams().WithConfigUID(CloudConfigId).WithBody(machinePool)
	}

	_, err := h.Client.V1CloudConfigsEksMachinePoolCreate(params)
	return err
}

func (h *V1Client) UpdateMachinePoolEks(CloudConfigId, scope string, machinePool *models.V1EksMachinePoolConfigEntity) error {
	var params *clientV1.V1CloudConfigsEksMachinePoolUpdateParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsEksMachinePoolUpdateParamsWithContext(h.Ctx).
			WithConfigUID(CloudConfigId).
			WithMachinePoolName(*machinePool.PoolConfig.Name).
			WithBody(machinePool)
	case "tenant":
		params = clientV1.NewV1CloudConfigsEksMachinePoolUpdateParams().
			WithConfigUID(CloudConfigId).
			WithMachinePoolName(*machinePool.PoolConfig.Name).
			WithBody(machinePool)
	}

	_, err := h.Client.V1CloudConfigsEksMachinePoolUpdate(params)
	return err
}

func (h *V1Client) DeleteMachinePoolEks(CloudConfigId, machinePoolName, scope string) error {
	var params *clientV1.V1CloudConfigsEksMachinePoolDeleteParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsEksMachinePoolDeleteParamsWithContext(h.Ctx).WithConfigUID(CloudConfigId).WithMachinePoolName(machinePoolName)
	case "tenant":
		params = clientV1.NewV1CloudConfigsEksMachinePoolDeleteParams().WithConfigUID(CloudConfigId).WithMachinePoolName(machinePoolName)
	}

	_, err := h.Client.V1CloudConfigsEksMachinePoolDelete(params)
	return err
}

func (h *V1Client) UpdateFargateProfilesEks(CloudConfigId, scope string, fargateProfiles *models.V1EksFargateProfiles) error {
	var params *clientV1.V1CloudConfigsEksUIDFargateProfilesUpdateParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsEksUIDFargateProfilesUpdateParamsWithContext(h.Ctx).
			WithConfigUID(CloudConfigId).
			WithBody(fargateProfiles)
	case "tenant":
		params = clientV1.NewV1CloudConfigsEksUIDFargateProfilesUpdateParams().
			WithConfigUID(CloudConfigId).
			WithBody(fargateProfiles)
	}

	_, err := h.Client.V1CloudConfigsEksUIDFargateProfilesUpdate(params)
	return err
}

func (h *V1Client) GetCloudConfigEks(configUID, scope string) (*models.V1EksCloudConfig, error) {
	var params *clientV1.V1CloudConfigsEksGetParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsEksGetParamsWithContext(h.Ctx).WithConfigUID(configUID)
	case "tenant":
		params = clientV1.NewV1CloudConfigsEksGetParams().WithConfigUID(configUID)
	}

	success, err := h.Client.V1CloudConfigsEksGet(params)
	var e *transport.TransportError
	if errors.As(err, &e) && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return success.Payload, nil
}

func (h *V1Client) GetNodeStatusMapEks(configUID, machinePoolName, scope string) (map[string]models.V1CloudMachineStatus, error) {
	var params *clientV1.V1CloudConfigsEksPoolMachinesListParams
	switch scope {
	case "project":
		params = clientV1.NewV1CloudConfigsEksPoolMachinesListParamsWithContext(h.Ctx).WithConfigUID(configUID).WithMachinePoolName(machinePoolName)
	case "tenant":
		params = clientV1.NewV1CloudConfigsEksPoolMachinesListParams().WithConfigUID(configUID).WithMachinePoolName(machinePoolName)
	}

	mpList, err := h.Client.V1CloudConfigsEksPoolMachinesList(params)
	nMap := map[string]models.V1CloudMachineStatus{}
	if len(mpList.Payload.Items) > 0 {
		for _, node := range mpList.Payload.Items {
			nMap[node.Metadata.UID] = *node.Status
		}
	}
	return nMap, err
}
