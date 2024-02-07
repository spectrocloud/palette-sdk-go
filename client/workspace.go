package client

import (
	"errors"

	"github.com/spectrocloud/palette-api-go/apiutil/transport"
	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
)

func (h *V1Client) CreateWorkspace(workspace *models.V1WorkspaceEntity) (string, error) {
	params := clientV1.NewV1WorkspacesCreateParamsWithContext(h.Ctx).WithBody(workspace)
	success, err := h.Client.V1WorkspacesCreate(params)
	if err != nil {
		return "", err
	}

	return *success.Payload.UID, nil
}

func (h *V1Client) GetWorkspace(uid string) (*models.V1Workspace, error) {
	params := clientV1.NewV1WorkspacesUIDGetParamsWithContext(h.Ctx).WithUID(uid)
	success, err := h.Client.V1WorkspacesUIDGet(params)
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
	params := clientV1.NewV1DashboardWorkspacesListParamsWithContext(h.Ctx)
	success, err := h.Client.V1DashboardWorkspacesList(params)
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

func (h *V1Client) UpdateWorkspaceResourceAllocation(uid string, we *models.V1WorkspaceClusterNamespacesEntity) error {
	params := clientV1.NewV1WorkspacesUIDClusterNamespacesUpdateParamsWithContext(h.Ctx).WithUID(uid).WithBody(we)
	if _, err := h.Client.V1WorkspacesUIDClusterNamespacesUpdate(params); err != nil {
		return err
	}
	return nil
}

func (h *V1Client) UpdateWorkspaceRBACS(uid, rbac_uid string, r *models.V1ClusterRbac) error {
	params := clientV1.NewV1WorkspacesUIDClusterRbacUpdateParamsWithContext(h.Ctx).WithUID(uid).WithClusterRbacUID(rbac_uid).WithBody(r)
	if _, err := h.Client.V1WorkspacesUIDClusterRbacUpdate(params); err != nil {
		return err
	}
	return nil
}

func (h *V1Client) UpdateWorkspaceBackupConfig(uid string, config *models.V1WorkspaceBackupConfigEntity) error {
	params := clientV1.NewV1WorkspaceOpsBackupUpdateParams().WithContext(h.Ctx).WithUID(uid).WithBody(config)
	_, err := h.Client.V1WorkspaceOpsBackupUpdate(params)
	return err
}

func (h *V1Client) DeleteWorkspace(uid string) error {
	params := clientV1.NewV1WorkspacesUIDDeleteParamsWithContext(h.Ctx).WithUID(uid)
	if _, err := h.Client.V1WorkspacesUIDDelete(params); err != nil {
		return err
	}
	return nil
}

func (h *V1Client) GetWorkspaceBackup(uid string) (*models.V1WorkspaceBackup, error) {
	params := clientV1.NewV1WorkspaceOpsBackupGetParams().WithContext(h.Ctx).WithUID(uid)
	success, err := h.Client.V1WorkspaceOpsBackupGet(params)
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
