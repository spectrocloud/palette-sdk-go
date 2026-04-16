package clusters

import "github.com/spectrocloud/palette-sdk-go/api/models"

func (s *service) GetKubernetesCerts(uid string) (*models.V1MachineCertificates, error) {
	if uid == "" {
		return nil, errUIDRequired
	}
	return s.client.GetTheKubernetesCerts(uid)
}

func (s *service) InitiateCertRenewal(uid string) error {
	if uid == "" {
		return errUIDRequired
	}
	return s.client.InitiateTheCertRenewal(uid)
}
