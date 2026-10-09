package client

import (
	clientv1 "github.com/spectrocloud/palette-sdk-go/api/client/version1"
	"github.com/spectrocloud/palette-sdk-go/api/models"
)

// UpdateImagePullSecretClusterExclusion sets or clears the hardened (DHI) image pull secret
// propagation exclude flag on the given clusters for the authenticated tenant.
func (h *V1Client) UpdateImagePullSecretClusterExclusion(clusterUIDs []string, exclude bool) error {
	body := &models.V1ImagePullSecretPropagationClustersUpdate{
		ClusterUids: clusterUIDs,
		Exclude:     &exclude,
	}
	params := clientv1.NewV1SpectroClustersImagePullSecretClustersUpdateParamsWithContext(h.ctx).WithBody(body)
	_, err := h.Client.V1SpectroClustersImagePullSecretClustersUpdate(params)
	return err
}

// GetImagePullSecretStatus retrieves a single page of hardened (DHI) image pull secret
// propagation status for the authenticated tenant. Empty clusterName or projectUID values are
// omitted from the request, and a limit of zero leaves the API default (50) in place. Callers
// that need the full result set are expected to page by advancing offset.
func (h *V1Client) GetImagePullSecretStatus(clusterName, projectUID string, limit, offset int64) (*models.V1ImagePullSecretTenantPropagationStatus, error) {
	params := clientv1.NewV1SpectroClustersImagePullSecretStatusGetParamsWithContext(h.ctx)
	if clusterName != "" {
		params = params.WithClusterName(&clusterName)
	}
	if projectUID != "" {
		params = params.WithProjectUID(&projectUID)
	}
	if limit > 0 {
		params = params.WithLimit(&limit)
	}
	if offset > 0 {
		params = params.WithOffset(&offset)
	}

	resp, err := h.Client.V1SpectroClustersImagePullSecretStatusGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}
