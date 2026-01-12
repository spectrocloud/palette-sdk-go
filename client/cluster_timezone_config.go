package client

import (
	clientv1 "github.com/spectrocloud/palette-sdk-go/api/client/version1"
	"github.com/spectrocloud/palette-sdk-go/api/models"
	"github.com/spectrocloud/palette-sdk-go/client/apiutil"
)

// UpdateClusterTimezone updates the timezone configuration for a cluster.
// The timezone must be a valid IANA timezone string (e.g., "America/New_York", "Asia/Kolkata", "Europe/London").
// This is mandatory when the cluster is deployed through a template.
func (h *V1Client) UpdateClusterTimezone(uid string, timezone string) error {
	params := clientv1.NewV1SpectroClustersUIDTimezoneUpdateParamsWithContext(h.ctx).
		WithUID(uid).
		WithBody(&models.V1TimezoneUpdateEntity{
			Timezone: apiutil.Ptr(timezone),
		})
	_, err := h.Client.V1SpectroClustersUIDTimezoneUpdate(params)
	return err
}
