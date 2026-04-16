package clusters

import "github.com/spectrocloud/palette-sdk-go/api/models"

func (s *service) GetVariables(uid string) ([]*models.V1SpectroClusterVariables, error) {
	if uid == "" {
		return nil, errUIDRequired
	}
	return s.client.GetClusterVariables(uid)
}

func (s *service) UpdateVariables(uid string, body []*models.V1SpectroClusterVariableUpdateEntity) error {
	if uid == "" {
		return errUIDRequired
	}
	return s.client.UpdateClusterProfileVariableInCluster(uid, body)
}
