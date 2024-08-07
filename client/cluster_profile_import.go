package client

import (
	"github.com/go-openapi/runtime"

	"github.com/spectrocloud/palette-sdk-go/api/models"

	clientv1 "github.com/spectrocloud/palette-sdk-go/api/client/v1"
	"github.com/spectrocloud/palette-sdk-go/client/apiutil"
)

// CreateClusterProfileImport creates a new cluster profile from a cluster profile import file.
func (h *V1Client) CreateClusterProfileImport(importFile runtime.NamedReadCloser) (string, error) {
	params := clientv1.NewV1ClusterProfilesImportFileParamsWithContext(h.ctx).
		WithPublish(apiutil.Ptr(true)).
		WithImportFile(importFile)
	resp, err := h.Client.V1ClusterProfilesImportFile(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

// ClusterProfileExport retrieves and exports a cluster profile by its unique identifier.
func (h *V1Client) ClusterProfileExport(uid string) (*models.V1ClusterProfile, error) {
	// no need to switch request context here as /v1/clusterprofiles/{uid} works for profile in any scope.
	params := clientv1.NewV1ClusterProfilesGetParamsWithContext(h.ctx).WithUID(uid)
	success, err := h.Client.V1ClusterProfilesGet(params)
	if err != nil {
		return nil, err
	}
	return success.Payload, nil

}
