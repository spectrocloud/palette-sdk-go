package clusters

import (
	"errors"

	"github.com/spectrocloud/palette-sdk-go/api/models"
)

func (s *service) CreateAws(cluster *models.V1SpectroAwsClusterEntity) (string, error) {
	if cluster == nil {
		return "", errors.New("AWS cluster entity is required")
	}
	return s.client.CreateClusterAws(cluster)
}

func (s *service) CreateAzure(cluster *models.V1SpectroAzureClusterEntity) (string, error) {
	if cluster == nil {
		return "", errors.New("Azure cluster entity is required")
	}
	return s.client.CreateClusterAzure(cluster)
}

func (s *service) CreateGcp(cluster *models.V1SpectroGcpClusterEntity) (string, error) {
	if cluster == nil {
		return "", errors.New("GCP cluster entity is required")
	}
	return s.client.CreateClusterGcp(cluster)
}

func (s *service) CreateVsphere(cluster *models.V1SpectroVsphereClusterEntity) (string, error) {
	if cluster == nil {
		return "", errors.New("vSphere cluster entity is required")
	}
	return s.client.CreateClusterVsphere(cluster)
}

func (s *service) CreateMaas(cluster *models.V1SpectroMaasClusterEntity) (string, error) {
	if cluster == nil {
		return "", errors.New("MAAS cluster entity is required")
	}
	return s.client.CreateClusterMaas(cluster)
}

func (s *service) CreateCloudStack(cluster *models.V1SpectroCloudStackClusterEntity) (string, error) {
	if cluster == nil {
		return "", errors.New("CloudStack cluster entity is required")
	}
	return s.client.CreateClusterCloudStack(cluster)
}

func (s *service) CreateEks(cluster *models.V1SpectroEksClusterEntity) (string, error) {
	if cluster == nil {
		return "", errors.New("EKS cluster entity is required")
	}
	return s.client.CreateClusterEks(cluster)
}

func (s *service) CreateAks(cluster *models.V1SpectroAzureClusterEntity) (string, error) {
	if cluster == nil {
		return "", errors.New("AKS cluster entity is required")
	}
	return s.client.CreateClusterAks(cluster)
}

func (s *service) CreateGke(cluster *models.V1SpectroGcpClusterEntity) (string, error) {
	if cluster == nil {
		return "", errors.New("GKE cluster entity is required")
	}
	return s.client.CreateClusterGke(cluster)
}

func (s *service) CreateEdgeNative(cluster *models.V1SpectroEdgeNativeClusterEntity) (string, error) {
	if cluster == nil {
		return "", errors.New("Edge Native cluster entity is required")
	}
	return s.client.CreateClusterEdgeNative(cluster)
}

func (s *service) CreateEdgeVsphere(cluster *models.V1SpectroVsphereClusterEntity) (string, error) {
	if cluster == nil {
		return "", errors.New("Edge vSphere cluster entity is required")
	}
	return s.client.CreateClusterEdgeVsphere(cluster)
}

func (s *service) CreateVirtual(cluster *models.V1SpectroVirtualClusterEntity) (string, error) {
	if cluster == nil {
		return "", errors.New("virtual cluster entity is required")
	}
	return s.client.CreateClusterVirtual(cluster)
}

func (s *service) CreateCustomCloud(cluster *models.V1SpectroCustomClusterEntity, cloudType string) (string, error) {
	if cluster == nil {
		return "", errors.New("custom cloud cluster entity is required")
	}
	if cloudType == "" {
		return "", errors.New("cloud type is required")
	}
	return s.client.CreateClusterCustomCloud(cluster, cloudType)
}
