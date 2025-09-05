package client

import (
	"fmt"

	clientv1 "github.com/spectrocloud/palette-sdk-go/api/client/version1"
	"github.com/spectrocloud/palette-sdk-go/api/models"
	"github.com/spectrocloud/palette-sdk-go/client/apiutil"
)

// GetLatestServiceVersionCLI get service version info for an edge service
func (h *V1Client) GetLatestServiceVersionCLI(serviceName, clusterUID, edgeHostUID string) (string, error) {
	params := clientv1.NewV1ServiceVersionGetParamsWithContext(h.ctx).
		WithServiceName(serviceName)

	if edgeHostUID != "" {
		params.SetEdgeHostUID(&edgeHostUID)
	}

	if clusterUID != "" {
		params.SetClusterUID(&clusterUID)
	}

	versionResp, err := h.Client.V1ServiceVersionGet(params)
	if err != nil {
		return "", fmt.Errorf("failed to fetch latest version for service '%s'. %w", serviceName, err)
	}
	if versionResp.Payload != nil &&
		versionResp.Payload.Spec != nil &&
		versionResp.Payload.Spec.LatestVersion != nil &&
		len(versionResp.Payload.Spec.LatestVersion.Content) > 0 {

		return versionResp.Payload.Spec.LatestVersion.Content, nil
	}
	return "", fmt.Errorf("failed to get latest version for service %s as content is empty", serviceName)
}

// GetManifestsForCLI ..
func (h *V1Client) GetManifestsForCLI(service, svcVersion, resourceFileName, action, edgeHostUID string) (manifests []*models.V1GitRepoFileContent, err error) {
	cP := clientv1.NewV1ServiceManifestGetParamsWithContext(h.ctx).
		WithServiceName(service).
		WithVersion(svcVersion).
		WithAction(action)

	if edgeHostUID != "" {
		cP.SetEdgeHostUID(&edgeHostUID)
	}

	if len(resourceFileName) > 0 {
		cP.WithResourceFilename(&resourceFileName)
	}

	manifestResp, err := h.Client.V1ServiceManifestGet(cP)
	if err != nil {
		if action == "delete" && isDeleteManifestNotExistError(err) {
			return nil, nil
		}
		err = fmt.Errorf("failed to fetch manifest for service '%s' and version %s for action '%s'. %w", service, svcVersion, action, err)
		return []*models.V1GitRepoFileContent{}, err
	}

	if manifestResp.Payload != nil &&
		manifestResp.Payload.Spec != nil &&
		len(manifestResp.Payload.Spec.Manifests) > 0 {
		return manifestResp.Payload.Spec.Manifests, nil
	}

	return []*models.V1GitRepoFileContent{}, fmt.Errorf("failed to get latest manifest for service %s as content is empty", service)
}

func isDeleteManifestNotExistError(err error) bool {
	const (
		RequestError              = "RequestError"
		RepoFileNotFoundError     = "RepoFileNotFound"
		StatusInternalServerError = "StatusInternalServerError"
	)

	return apiutil.ToV1ErrorObj(err).Code == RequestError ||
		apiutil.ToV1ErrorObj(err).Code == RepoFileNotFoundError ||
		apiutil.ToV1ErrorObj(err).Code == StatusInternalServerError
}
