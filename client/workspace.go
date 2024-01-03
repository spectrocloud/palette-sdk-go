package client

import (
	"errors"

	"github.com/spectrocloud/hapi/apiutil/transport"
	hashboardC "github.com/spectrocloud/hapi/hashboard/client/v1"
	"github.com/spectrocloud/hapi/models"
	clusterC "github.com/spectrocloud/hapi/spectrocluster/client/v1"
)

func (h *V1Client) CreateWorkspace(workspace *models.V1WorkspaceEntity) (string, error) {
	params := clusterC.NewV1WorkspacesCreateParamsWithContext(h.Ctx).WithBody(workspace)
	success, err := h.GetClusterClient().V1WorkspacesCreate(params)
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

func (h *V1Client) GetWorkspace(uid string) (*models.V1Workspace, error) {
	params := clusterC.NewV1WorkspacesUIDGetParamsWithContext(h.Ctx).WithUID(uid)
	success, err := h.GetClusterClient().V1WorkspacesUIDGet(params)
	var e *transport.TransportError
	if errors.As(err, &e) && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	workspace := success.Payload

	return workspace, nil
}

func (h *V1Client) GetWorkspaceByName(name string) (*models.V1DashboardWorkspace, error) {
	params := hashboardC.NewV1DashboardWorkspacesListParamsWithContext(h.Ctx)
	success, err := h.GetHashboardClient().V1DashboardWorkspacesList(params)
	var e *transport.TransportError
	if errors.As(err, &e) && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	for _, workspace := range success.Payload.Items {
		if workspace.Metadata.Name == name {
			return workspace, nil
		}
	}

	return nil, nil
}

func (h *V1Client) UpdateWorkspaceResourceAllocation(uid string, wo *models.V1WorkspaceResourceAllocationsEntity) error {
	params := clusterC.NewV1WorkspacesUIDResourceAllocationsUpdateParamsWithContext(h.Ctx).WithUID(uid).WithBody(wo)
	if _, err := h.GetClusterClient().V1WorkspacesUIDResourceAllocationsUpdate(params); err != nil {
		return err
	}
	return nil
}

func (h *V1Client) UpdateWorkspaceRBACS(uid, rbac_uid string, wo *models.V1ClusterRbac) error {
	params := clusterC.NewV1WorkspacesUIDClusterRbacUpdateParamsWithContext(h.Ctx).WithUID(uid).WithClusterRbacUID(rbac_uid).WithBody(wo)
	if _, err := h.GetClusterClient().V1WorkspacesUIDClusterRbacUpdate(params); err != nil {
		return err
	}
	return nil
}

func (h *V1Client) UpdateWorkspaceBackupConfig(uid string, config *models.V1WorkspaceBackupConfigEntity) error {
	params := clusterC.NewV1WorkspaceOpsBackupUpdateParams().WithContext(h.Ctx).WithUID(uid).WithBody(config)
	_, err := h.GetClusterClient().V1WorkspaceOpsBackupUpdate(params)
	return err
}

func (h *V1Client) DeleteWorkspace(uid string) error {
	params := clusterC.NewV1WorkspacesUIDDeleteParamsWithContext(h.Ctx).WithUID(uid)
	if _, err := h.GetClusterClient().V1WorkspacesUIDDelete(params); err != nil {
		return err
	}
	return nil
}

func (h *V1Client) GetWorkspaceBackup(uid string) (*models.V1WorkspaceBackup, error) {
	params := clusterC.NewV1WorkspaceOpsBackupGetParams().WithContext(h.Ctx).WithUID(uid)
	success, err := h.GetClusterClient().V1WorkspaceOpsBackupGet(params)
	var e *transport.TransportError
	if errors.As(err, &e) && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	workspace := success.Payload

	return workspace, nil
}

func (h *V1Client) WorkspaceBackupDelete() error {
	return errors.New("not implemented")
}
