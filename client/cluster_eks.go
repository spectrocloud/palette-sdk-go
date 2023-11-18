package client

import (
	"errors"
	"time"

	"github.com/spectrocloud/hapi/apiutil/transport"
	"github.com/spectrocloud/hapi/models"
	clusterC "github.com/spectrocloud/hapi/spectrocluster/client/v1"
)

func (h *V1Client) CreateClusterEks(cluster *models.V1SpectroEksClusterEntity, ClusterContext string) (string, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return "", err
	}

	var params *clusterC.V1SpectroClustersEksCreateParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1SpectroClustersEksCreateParamsWithContext(h.Ctx).WithBody(cluster)
	case "tenant":
		params = clusterC.NewV1SpectroClustersEksCreateParams().WithBody(cluster)
	}

	success, err := client.V1SpectroClustersEksCreate(params.WithTimeout(90 * time.Second))
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

func (h *V1Client) CreateMachinePoolEks(cloudConfigId, ClusterContext string, machinePool *models.V1EksMachinePoolConfigEntity) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil
	}

	var params *clusterC.V1CloudConfigsEksMachinePoolCreateParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsEksMachinePoolCreateParamsWithContext(h.Ctx).WithConfigUID(cloudConfigId).WithBody(machinePool)
	case "tenant":
		params = clusterC.NewV1CloudConfigsEksMachinePoolCreateParams().WithConfigUID(cloudConfigId).WithBody(machinePool)
	}

	_, err = client.V1CloudConfigsEksMachinePoolCreate(params)
	return err
}

func (h *V1Client) UpdateMachinePoolEks(cloudConfigId, ClusterContext string, machinePool *models.V1EksMachinePoolConfigEntity) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil
	}

	var params *clusterC.V1CloudConfigsEksMachinePoolUpdateParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsEksMachinePoolUpdateParamsWithContext(h.Ctx).
			WithConfigUID(cloudConfigId).
			WithMachinePoolName(*machinePool.PoolConfig.Name).
			WithBody(machinePool)
	case "tenant":
		params = clusterC.NewV1CloudConfigsEksMachinePoolUpdateParams().
			WithConfigUID(cloudConfigId).
			WithMachinePoolName(*machinePool.PoolConfig.Name).
			WithBody(machinePool)
	}

	_, err = client.V1CloudConfigsEksMachinePoolUpdate(params)
	return err
}

func (h *V1Client) DeleteMachinePoolEks(cloudConfigId, machinePoolName, ClusterContext string) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil
	}

	var params *clusterC.V1CloudConfigsEksMachinePoolDeleteParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsEksMachinePoolDeleteParamsWithContext(h.Ctx).WithConfigUID(cloudConfigId).WithMachinePoolName(machinePoolName)
	case "tenant":
		params = clusterC.NewV1CloudConfigsEksMachinePoolDeleteParams().WithConfigUID(cloudConfigId).WithMachinePoolName(machinePoolName)
	}

	_, err = client.V1CloudConfigsEksMachinePoolDelete(params)
	return err
}

func (h *V1Client) UpdateFargateProfilesEks(cloudConfigId, ClusterContext string, fargateProfiles *models.V1EksFargateProfiles) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil
	}

	var params *clusterC.V1CloudConfigsEksUIDFargateProfilesUpdateParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsEksUIDFargateProfilesUpdateParamsWithContext(h.Ctx).
			WithConfigUID(cloudConfigId).
			WithBody(fargateProfiles)
	case "tenant":
		params = clusterC.NewV1CloudConfigsEksUIDFargateProfilesUpdateParams().
			WithConfigUID(cloudConfigId).
			WithBody(fargateProfiles)
	}

	_, err = client.V1CloudConfigsEksUIDFargateProfilesUpdate(params)
	return err
}

func (h *V1Client) GetCloudConfigEks(configUID, ClusterContext string) (*models.V1EksCloudConfig, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}

	var params *clusterC.V1CloudConfigsEksGetParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsEksGetParamsWithContext(h.Ctx).WithConfigUID(configUID)
	case "tenant":
		params = clusterC.NewV1CloudConfigsEksGetParams().WithConfigUID(configUID)
	}

	success, err := client.V1CloudConfigsEksGet(params)
	var e *transport.TransportError
	if errors.As(err, &e) && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return success.Payload, nil
}

func (h *V1Client) GetNodeStatusMapEks(configUID string, machinePoolName string, ClusterContext string) (map[string]models.V1CloudMachineStatus, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}

	var params *clusterC.V1CloudConfigsEksPoolMachinesListParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1CloudConfigsEksPoolMachinesListParamsWithContext(h.Ctx).WithConfigUID(configUID).WithMachinePoolName(machinePoolName)
	case "tenant":
		params = clusterC.NewV1CloudConfigsEksPoolMachinesListParams().WithConfigUID(configUID).WithMachinePoolName(machinePoolName)
	}

	mpList, err := client.V1CloudConfigsEksPoolMachinesList(params)
	nMap := map[string]models.V1CloudMachineStatus{}
	if len(mpList.Payload.Items) > 0 {
		for _, node := range mpList.Payload.Items {
			nMap[node.Metadata.UID] = *node.Status
		}
	}
	return nMap, err
}
