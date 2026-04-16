package deployments

import (
	"errors"

	"github.com/spectrocloud/palette-sdk-go/api/models"
)

func (s *service) GetApplication(uid string) (*models.V1AppDeployment, error) {
	if uid == "" {
		return nil, errUIDRequired
	}
	return s.client.GetApplication(uid)
}

func (s *service) CreateApplicationWithNewSandboxCluster(body *models.V1AppDeploymentClusterGroupEntity) (string, error) {
	if body == nil {
		return "", errors.New("application deployment entity is required")
	}
	return s.client.CreateApplicationWithNewSandboxCluster(body)
}

func (s *service) CreateApplicationWithExistingSandboxCluster(body *models.V1AppDeploymentVirtualClusterEntity) (string, error) {
	if body == nil {
		return "", errors.New("application deployment entity is required")
	}
	return s.client.CreateApplicationWithExistingSandboxCluster(body)
}

func (s *service) DeleteApplication(uid string) error {
	if uid == "" {
		return errUIDRequired
	}
	return s.client.DeleteApplication(uid)
}
