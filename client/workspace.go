package client

import (
	"errors"

	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
	"github.com/spectrocloud/palette-sdk-go/client/apiutil"
)

func (h *V1Client) CreateWorkspace(workspace *models.V1WorkspaceEntity) (string, error) {
	params := clientV1.NewV1WorkspacesCreateParamsWithContext(h.ctx).
		WithBody(workspace)
	resp, err := h.Client.V1WorkspacesCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

func (h *V1Client) GetWorkspace(uid string) (*models.V1Workspace, error) {
	params := clientV1.NewV1WorkspacesUIDGetParamsWithContext(h.ctx).
		WithUID(uid)
	resp, err := h.Client.V1WorkspacesUIDGet(params)
	if err := apiutil.Handle404(err); err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

func (h *V1Client) GetWorkspaceByName(name string) (*models.V1DashboardWorkspace, error) {
	params := clientV1.NewV1DashboardWorkspacesListParamsWithContext(h.ctx)
	resp, err := h.Client.V1DashboardWorkspacesList(params)
	if err := apiutil.Handle404(err); err != nil {
		return nil, err
	}
	for _, workspace := range resp.Payload.Items {
		if workspace.Metadata.Name == name {
			return workspace, nil
		}
	}
	return nil, nil
}

func (h *V1Client) UpdateWorkspaceResourceAllocation(uid string, we *models.V1WorkspaceClusterNamespacesEntity) error {
	params := clientV1.NewV1WorkspacesUIDClusterNamespacesUpdateParamsWithContext(h.ctx).
		WithUID(uid).
		WithBody(we)
	_, err := h.Client.V1WorkspacesUIDClusterNamespacesUpdate(params)
	return err
}

func (h *V1Client) UpdateWorkspaceRBACS(uid, rbacUid string, r *models.V1ClusterRbac) error {
	params := clientV1.NewV1WorkspacesUIDClusterRbacUpdateParamsWithContext(h.ctx).
		WithUID(uid).
		WithClusterRbacUID(rbacUid).
		WithBody(r)
	_, err := h.Client.V1WorkspacesUIDClusterRbacUpdate(params)
	return err
}

func (h *V1Client) UpdateWorkspaceBackupConfig(uid string, config *models.V1WorkspaceBackupConfigEntity) error {
	params := clientV1.NewV1WorkspaceOpsBackupUpdateParamsWithContext(h.ctx).
		WithUID(uid).
		WithBody(config)
	_, err := h.Client.V1WorkspaceOpsBackupUpdate(params)
	return err
}

func (h *V1Client) DeleteWorkspace(uid string) error {
	params := clientV1.NewV1WorkspacesUIDDeleteParamsWithContext(h.ctx).
		WithUID(uid)
	_, err := h.Client.V1WorkspacesUIDDelete(params)
	return err
}

func (h *V1Client) GetWorkspaceBackup(uid string) (*models.V1WorkspaceBackup, error) {
	params := clientV1.NewV1WorkspaceOpsBackupGetParamsWithContext(h.ctx).
		WithUID(uid)
	resp, err := h.Client.V1WorkspaceOpsBackupGet(params)
	if err := apiutil.Handle404(err); err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

func (h *V1Client) WorkspaceBackupDelete() error {
	return errors.New("not implemented")
}
