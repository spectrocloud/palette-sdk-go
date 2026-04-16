package clusters

import "github.com/spectrocloud/palette-sdk-go/api/models"

func (s *service) UpdateProfileValues(uid string, profiles *models.V1SpectroClusterProfiles) error {
	if uid == "" {
		return errUIDRequired
	}
	return s.client.UpdateClusterProfileValues(uid, profiles)
}

func (s *service) PatchProfileValues(uid string, profiles *models.V1SpectroClusterProfiles) error {
	if uid == "" {
		return errUIDRequired
	}
	return s.client.PatchClusterProfileValues(uid, profiles)
}
