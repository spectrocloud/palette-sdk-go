package client

import (
	"fmt"

	clientv1 "github.com/spectrocloud/palette-sdk-go/api/client/version1"
	"github.com/spectrocloud/palette-sdk-go/api/models"
	"github.com/spectrocloud/palette-sdk-go/client/apiutil"
)

func toV1OverlordsUIDCloudStackAccountValidateBody(account *models.V1CloudStackAccount) clientv1.V1OverlordsUIDCloudStackAccountValidateBody {
	return clientv1.V1OverlordsUIDCloudStackAccountValidateBody{
		Account: account.Spec,
	}
}

// CreateCloudAccountCloudStack creates a new CloudStack cloud account.
func (h *V1Client) CreateCloudAccountCloudStack(account *models.V1CloudStackAccount) (string, error) {
	// validate account
	if err := h.validateCloudAccountCloudStack(account); err != nil {
		return "", err
	}

	// Convert to input entity type
	inputEntity := &models.V1CloudStackAccountInputEntity{
		Metadata: &models.V1ObjectMetaInputEntity{
			Name:        account.Metadata.Name,
			Labels:      account.Metadata.Labels,
			Annotations: account.Metadata.Annotations,
		},
		Spec: account.Spec,
	}

	params := clientv1.NewV1CloudAccountsCloudStackCreateParamsWithContext(h.ctx).
		WithBody(inputEntity)
	resp, err := h.Client.V1CloudAccountsCloudStackCreate(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

func (h *V1Client) validateCloudAccountCloudStack(account *models.V1CloudStackAccount) error {
	// check PCG
	pcgID := account.Metadata.Annotations[OverlordUID]
	if err := h.CheckPCG(pcgID); err != nil {
		return err
	}

	// validate account
	params := clientv1.NewV1OverlordsUIDCloudStackAccountValidateParamsWithContext(h.ctx).
		WithUID(pcgID).
		WithBody(toV1OverlordsUIDCloudStackAccountValidateBody(account))

	_, err := h.Client.V1OverlordsUIDCloudStackAccountValidate(params)
	return err
}

// UpdateCloudAccountCloudStack updates an existing CloudStack cloud account.
func (h *V1Client) UpdateCloudAccountCloudStack(account *models.V1CloudStackAccount) error {
	// validate account
	if err := h.validateCloudAccountCloudStack(account); err != nil {
		return err
	}

	// Convert to update entity type
	updateEntity := &models.V1CloudStackAccountUpdateEntity{
		Metadata: &models.V1ObjectMetaUpdateEntity{
			Labels:      account.Metadata.Labels,
			Annotations: account.Metadata.Annotations,
		},
		Spec: account.Spec,
	}

	params := clientv1.NewV1CloudAccountsCloudStackUpdateParamsWithContext(h.ctx).
		WithUID(account.Metadata.UID).
		WithBody(updateEntity)

	_, err := h.Client.V1CloudAccountsCloudStackUpdate(params)
	return err
}

// DeleteCloudAccountCloudStack deletes an existing CloudStack cloud account.
func (h *V1Client) DeleteCloudAccountCloudStack(uid string) error {
	params := clientv1.NewV1CloudAccountsCloudStackDeleteParamsWithContext(h.ctx).
		WithUID(uid)
	_, err := h.Client.V1CloudAccountsCloudStackDelete(params)
	return err
}

// GetCloudAccountCloudStack retrieves an existing CloudStack cloud account.
func (h *V1Client) GetCloudAccountCloudStack(uid string) (*models.V1CloudStackAccount, error) {
	params := clientv1.NewV1CloudAccountsCloudStackGetParamsWithContext(h.ctx).
		WithUID(uid)
	resp, err := h.Client.V1CloudAccountsCloudStackGet(params)
	if apiutil.Is404(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// GetCloudAccountsCloudStack retrieves all CloudStack cloud accounts.
func (h *V1Client) GetCloudAccountsCloudStack() ([]*models.V1CloudStackAccount, error) {
	params := clientv1.NewV1CloudAccountsCloudStackListParamsWithContext(h.ctx).
		WithLimit(apiutil.Ptr(int64(0)))
	response, err := h.Client.V1CloudAccountsCloudStackList(params)
	if err != nil {
		return nil, err
	}
	return response.Payload.Items, nil
}

// GetCloudStackAccountNetworks lists CloudStack networks for a cloud account.
// Optional zone, projectID, and vpcID scope the lookup to match Palette UI behavior.
func (h *V1Client) GetCloudStackAccountNetworks(accountUID, zone, projectID, vpcID string) ([]*models.V1CloudStackNetwork, error) {
	params := clientv1.NewV1CloudstackAccountsUIDNetworksParamsWithContext(h.ctx).
		WithUID(accountUID)
	if zone != "" {
		params = params.WithZone(&zone)
	}
	if projectID != "" {
		params = params.WithProjectID(&projectID)
	}
	if vpcID != "" {
		params = params.WithVpcID(&vpcID)
	}
	resp, err := h.Client.V1CloudstackAccountsUIDNetworks(params)
	if err != nil {
		return nil, err
	}
	if resp.Payload == nil {
		return nil, nil
	}
	return resp.Payload.Networks, nil
}

// ResolveCloudStackNetworkID returns the CloudStack network ID for a network name
// within the given account and zone/project/vpc context.
func (h *V1Client) ResolveCloudStackNetworkID(accountUID, networkName, zone, projectID, vpcID string) (string, error) {
	if networkName == "" {
		return "", fmt.Errorf("network name is required to resolve CloudStack network ID")
	}
	networks, err := h.GetCloudStackAccountNetworks(accountUID, zone, projectID, vpcID)
	if err != nil {
		return "", fmt.Errorf("failed to list CloudStack networks for account %q: %w", accountUID, err)
	}
	var matches []string
	for _, network := range networks {
		if network != nil && network.Name == networkName && network.ID != "" {
			matches = append(matches, network.ID)
		}
	}
	switch len(matches) {
	case 0:
		return "", fmt.Errorf("no CloudStack network found with name %q for account %q (zone=%q)", networkName, accountUID, zone)
	case 1:
		return matches[0], nil
	default:
		return "", fmt.Errorf("multiple CloudStack networks found with name %q for account %q (zone=%q); specify network_id explicitly", networkName, accountUID, zone)
	}
}
