package client

import (
	"errors"

	"github.com/go-openapi/runtime"

	"github.com/spectrocloud/hapi/apiutil/transport"
	"github.com/spectrocloud/hapi/models"
	clusterC "github.com/spectrocloud/hapi/spectrocluster/client/v1"
)

func (h *V1Client) CreateClusterProfileImport(importFile runtime.NamedReadCloser, ProfileContext string) (string, error) {
	var params *clusterC.V1ClusterProfilesImportFileParams
	switch ProfileContext {
	case "project":
		params = clusterC.NewV1ClusterProfilesImportFileParamsWithContext(h.Ctx)
	case "tenant":
		params = clusterC.NewV1ClusterProfilesImportFileParams()
	}

	params = params.WithPublish(Ptr(true)).WithImportFile(importFile)
	success, err := h.GetClusterClient().V1ClusterProfilesImportFile(params)
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

func (h *V1Client) ClusterProfileExport(uid string) (*models.V1ClusterProfile, error) {
	// no need to switch request context here as /v1/clusterprofiles/{uid} works for profile in any scope.
	params := clusterC.NewV1ClusterProfilesGetParamsWithContext(h.Ctx).WithUID(uid)
	success, err := h.GetClusterClient().V1ClusterProfilesGet(params)
	var e *transport.TransportError
	if errors.As(err, &e) && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return success.Payload, nil
}
