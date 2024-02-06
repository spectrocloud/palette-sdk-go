package client

import (
	"errors"

	"github.com/go-openapi/runtime"

	"github.com/spectrocloud/palette-api-go/apiutil/transport"
	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
)

func (h *V1Client) CreateClusterProfileImport(importFile runtime.NamedReadCloser, ProfileContext string) (string, error) {
	var params *clientV1.V1ClusterProfilesImportFileParams
	switch ProfileContext {
	case "project":
		params = clientV1.NewV1ClusterProfilesImportFileParamsWithContext(h.Ctx)
	case "tenant":
		params = clientV1.NewV1ClusterProfilesImportFileParams()
	}

	params = params.WithPublish(Ptr(true)).WithImportFile(importFile)
	success, err := h.GetClient().V1ClusterProfilesImportFile(params)
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

func (h *V1Client) ClusterProfileExport(uid string) (*models.V1ClusterProfile, error) {
	// no need to switch request context here as /v1/clusterprofiles/{uid} works for profile in any scope.
	params := clientV1.NewV1ClusterProfilesGetParamsWithContext(h.Ctx).WithUID(uid)
	success, err := h.GetClient().V1ClusterProfilesGet(params)
	var e *transport.TransportError
	if errors.As(err, &e) && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return success.Payload, nil
}
