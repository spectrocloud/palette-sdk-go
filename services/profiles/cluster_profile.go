package profiles

import (
	"errors"

	"github.com/spectrocloud/palette-sdk-go/api/models"
)

func (s *service) GetSummary(uid string) (*models.V1ClusterProfileSummary, error) {
	if uid == "" {
		return nil, errUIDRequired
	}
	return s.client.GetClusterProfileSummary(uid)
}

func (s *service) Get(uid string) (*models.V1ClusterProfile, error) {
	if uid == "" {
		return nil, errUIDRequired
	}
	return s.client.GetClusterProfile(uid)
}

func (s *service) GetUID(profileName, profileVersion string) (string, error) {
	if profileName == "" {
		return "", errors.New("profile name is required")
	}
	return s.client.GetClusterProfileUID(profileName, profileVersion)
}

func (s *service) Create(profile *models.V1ClusterProfileEntity) (string, error) {
	if profile == nil {
		return "", errors.New("cluster profile entity is required")
	}
	return s.client.CreateClusterProfile(profile)
}

func (s *service) Update(profile *models.V1ClusterProfileUpdateEntity) error {
	if profile == nil {
		return errors.New("cluster profile update entity is required")
	}
	return s.client.UpdateClusterProfile(profile)
}

func (s *service) Delete(uid string) error {
	if uid == "" {
		return errUIDRequired
	}
	return s.client.DeleteClusterProfile(uid)
}

func (s *service) Publish(uid string) error {
	if uid == "" {
		return errUIDRequired
	}
	return s.client.PublishClusterProfile(uid)
}

func (s *service) Import(content string) (string, error) {
	if content == "" {
		return "", errors.New("profile content is required")
	}
	return s.client.ImportClusterProfile(content)
}

func (s *service) Export(uid, format string) ([]byte, error) {
	if uid == "" {
		return nil, errUIDRequired
	}
	buf, err := s.client.ExportClusterProfile(uid, format)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (s *service) List() ([]*models.V1ClusterProfileMetadata, error) {
	return s.client.GetClusterProfiles()
}

func (s *service) SearchProfiles(filter *models.V1ClusterProfilesFilterSpec) ([]*models.V1ClusterProfileSummary, error) {
	return s.client.SearchClusterProfiles(filter)
}

func (s *service) GetPack(uid string) (*models.V1PackTagEntity, error) {
	if uid == "" {
		return nil, errUIDRequired
	}
	return s.client.GetPack(uid)
}

func (s *service) GetVariables(uid string) ([]*models.V1Variable, error) {
	if uid == "" {
		return nil, errUIDRequired
	}
	return s.client.GetProfileVariables(uid)
}

func (s *service) UpdateVariables(variables *models.V1Variables, uid string) error {
	if uid == "" {
		return errUIDRequired
	}
	return s.client.UpdateProfileVariables(variables, uid)
}
