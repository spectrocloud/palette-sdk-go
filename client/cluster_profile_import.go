package client

import (
	"github.com/go-openapi/runtime"

	clientv1 "github.com/spectrocloud/palette-api-go/client/v1"
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
