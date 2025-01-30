package client

import (
	clientv1 "github.com/spectrocloud/palette-sdk-go/api/client/v1"
	"github.com/spectrocloud/palette-sdk-go/api/models"
	"github.com/spectrocloud/palette-sdk-go/client/apiutil"
)

// CreateWorkspace creates a new workspace.
func (h *V1Client) CreateWorkspace(workspace *models.V1WorkspaceEntity) (string, error) {
	params := clientv1.NewV1WorkspacesCreateParamsWithContext(h.ctx).
		WithBody(workspace)
	resp, err := h.Client.V1WorkspacesCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

// GetWorkspace retrieves an existing workspace by UID.
func (h *V1Client) GetWorkspace(uid string) (*models.V1Workspace, error) {
	params := clientv1.NewV1WorkspacesUIDGetParamsWithContext(h.ctx).
		WithUID(uid)
	resp, err := h.Client.V1WorkspacesUIDGet(params)
	if apiutil.Is404(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// GetWorkspaceByName retrieves an existing workspace by name.
func (h *V1Client) GetWorkspaceByName(name string) (*models.V1DashboardWorkspace, error) {
	params := clientv1.NewV1DashboardWorkspacesListParamsWithContext(h.ctx)
	resp, err := h.Client.V1DashboardWorkspacesList(params)
	if apiutil.Is404(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	for _, workspace := range resp.Payload.Items {
		if workspace.Metadata.Name == name {
			return workspace, nil
		}
	}
	return nil, nil
}

// UpdateWorkspaceResourceAllocation updates an existing workspace's resource allocation.
func (h *V1Client) UpdateWorkspaceResourceAllocation(uid string, we *models.V1WorkspaceClusterNamespacesEntity) error {
	params := clientv1.NewV1WorkspacesUIDClusterNamespacesUpdateParamsWithContext(h.ctx).
		WithUID(uid).
		WithBody(we)
	_, err := h.Client.V1WorkspacesUIDClusterNamespacesUpdate(params)
	return err
}

// UpdateWorkspaceRBACS updates an existing workspace's RBAC configuration.
func (h *V1Client) UpdateWorkspaceRBACS(uid, rbacUID string, r *models.V1ClusterRbac) error {
	params := clientv1.NewV1WorkspacesUIDClusterRbacUpdateParamsWithContext(h.ctx).
		WithUID(uid).
		WithClusterRbacUID(rbacUID).
		WithBody(r)
	_, err := h.Client.V1WorkspacesUIDClusterRbacUpdate(params)
	return err
}

// CreateWorkspaceBackupConfig create a backup configuration in existing workspace's .
func (h *V1Client) CreateWorkspaceBackupConfig(uid string, config *models.V1WorkspaceBackupConfigEntity) error {
	params := clientv1.NewV1WorkspaceOpsBackupCreateParamsWithContext(h.ctx).
		WithUID(uid).
		WithBody(config)
	_, err := h.Client.V1WorkspaceOpsBackupCreate(params)
	return err
}

// UpdateWorkspaceBackupConfig updates an existing workspace's backup configuration.
func (h *V1Client) UpdateWorkspaceBackupConfig(uid string, config *models.V1WorkspaceBackupConfigEntity) error {
	params := clientv1.NewV1WorkspaceOpsBackupUpdateParamsWithContext(h.ctx).
		WithUID(uid).
		WithBody(config)
	_, err := h.Client.V1WorkspaceOpsBackupUpdate(params)
	return err
}

// DeleteWorkspace deletes an existing workspace.
func (h *V1Client) DeleteWorkspace(uid string) error {
	params := clientv1.NewV1WorkspacesUIDDeleteParamsWithContext(h.ctx).
		WithUID(uid)
	_, err := h.Client.V1WorkspacesUIDDelete(params)
	return err
}

// GetWorkspaceBackup retrieves an existing workspace backup by UID.
func (h *V1Client) GetWorkspaceBackup(uid string) (*models.V1WorkspaceBackup, error) {
	params := clientv1.NewV1WorkspaceOpsBackupGetParamsWithContext(h.ctx).
		WithUID(uid)
	resp, err := h.Client.V1WorkspaceOpsBackupGet(params)
	if apiutil.Is404(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// func (h *V1Client) WorkspaceBackupDelete() error {
// 	return errors.New("not implemented")
// }
