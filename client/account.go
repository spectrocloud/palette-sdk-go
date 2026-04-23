package client

import (
	"fmt"

	clientv1 "github.com/spectrocloud/palette-sdk-go/api/client/version1"
	"github.com/spectrocloud/palette-sdk-go/api/models"
	"github.com/spectrocloud/palette-sdk-go/client/apiutil"
)

// ListCloudAccounts returns a list of all cloud account summaries.
func (h *V1Client) ListCloudAccounts() ([]*models.V1CloudAccountSummary, error) {
	params := clientv1.NewV1CloudAccountsListSummaryParamsWithContext(h.ctx).
		WithLimit(apiutil.Ptr(int64(0)))
	resp, err := h.Client.V1CloudAccountsListSummary(params)
	if apiutil.Is404(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return resp.Payload.Items, nil
}

// GetCloudAccount retrieves an existing cloud account summary.
func (h *V1Client) GetCloudAccount(uid string) (*models.V1CloudAccountSummary, error) {
	accounts, err := h.ListCloudAccounts()
	if err != nil {
		return nil, err
	}
	for _, account := range accounts {
		if account.Metadata.UID == uid {
			return account, nil
		}
	}
	return nil, fmt.Errorf("account not found with uid %s", uid)
}

// GetCloudAccountGcpByName retrieves a GCP cloud account by name (and optional scope) using the list GET.
// scope can be "" (any), "project", or "tenant".
func (h *V1Client) GetCloudAccountGcpByName(name, scope string) (*models.V1GcpAccount, error) {
	params := clientv1.NewV1CloudAccountsGcpListParamsWithContext(h.ctx).
		WithLimit(apiutil.Ptr(int64(0)))
	resp, err := h.Client.V1CloudAccountsGcpList(params)
	if err != nil {
		return nil, err
	}
	matchScope := func(annotations map[string]string) bool {
		return scope == "" || (annotations != nil && annotations["scope"] == scope)
	}
	for _, account := range resp.Payload.Items {
		if account.Metadata != nil && account.Metadata.Name == name && matchScope(account.Metadata.Annotations) {
			return account, nil
		}
	}
	return nil, fmt.Errorf("GCP cloud account with name %q not found", name)
}

// GetCloudAccountAwsByName retrieves an AWS cloud account by name (and optional scope) using the list GET.
// scope can be "" (any), "project", or "tenant".
func (h *V1Client) GetCloudAccountAwsByName(name, scope string) (*models.V1AwsAccount, error) {
	params := clientv1.NewV1CloudAccountsAwsListParamsWithContext(h.ctx).
		WithLimit(apiutil.Ptr(int64(0)))
	resp, err := h.Client.V1CloudAccountsAwsList(params)
	if err != nil {
		return nil, err
	}

	matchScope := func(annotations map[string]string) bool {
		return scope == "" || (annotations != nil && annotations["scope"] == scope)
	}
	for _, account := range resp.Payload.Items {
		if account.Metadata != nil && account.Metadata.Name == name && matchScope(account.Metadata.Annotations) {
			return account, nil
		}
	}
	return nil, fmt.Errorf("AWS cloud account with name %q not found", name)
}

// GetCloudAccountAzureByName retrieves an Azure cloud account by name (and optional scope) using the list GET.
// scope can be "" (any), "project", or "tenant".
func (h *V1Client) GetCloudAccountAzureByName(name, scope string) (*models.V1AzureAccount, error) {
	params := clientv1.NewV1CloudAccountsAzureListParamsWithContext(h.ctx).
		WithLimit(apiutil.Ptr(int64(0)))
	resp, err := h.Client.V1CloudAccountsAzureList(params)
	if err != nil {
		return nil, err
	}
	matchScope := func(annotations map[string]string) bool {
		return scope == "" || (annotations != nil && annotations["scope"] == scope)
	}
	for _, account := range resp.Payload.Items {
		if account.Metadata != nil && account.Metadata.Name == name && matchScope(account.Metadata.Annotations) {
			return account, nil
		}
	}
	return nil, fmt.Errorf("Azure cloud account with name %q not found", name)
}

// GetCloudAccountApacheCloudStackByName retrieves an Apache CloudStack cloud account by name (and optional scope) using the list GET.
// scope can be "" (any), "project", or "tenant".
func (h *V1Client) GetCloudAccountApacheCloudStackByName(name, scope string) (*models.V1CloudStackAccount, error) {
	params := clientv1.NewV1CloudAccountsCloudStackListParamsWithContext(h.ctx).
		WithLimit(apiutil.Ptr(int64(0)))
	resp, err := h.Client.V1CloudAccountsCloudStackList(params)
	if err != nil {
		return nil, err
	}
	matchScope := func(annotations map[string]string) bool {
		return scope == "" || (annotations != nil && annotations["scope"] == scope)
	}
	for _, account := range resp.Payload.Items {
		if account.Metadata != nil && account.Metadata.Name == name && matchScope(account.Metadata.Annotations) {
			return account, nil
		}
	}
	return nil, fmt.Errorf("Apache CloudStack cloud account with name %q not found", name)
}

// GetCloudAccountMaasByName retrieves a MAAS cloud account by name (and optional scope) using the list GET.
// scope can be "" (any), "project", or "tenant".
func (h *V1Client) GetCloudAccountMaasByName(name, scope string) (*models.V1MaasAccount, error) {
	params := clientv1.NewV1CloudAccountsMaasListParamsWithContext(h.ctx).
		WithLimit(apiutil.Ptr(int64(0)))
	resp, err := h.Client.V1CloudAccountsMaasList(params)
	if err != nil {
		return nil, err
	}
	matchScope := func(annotations map[string]string) bool {
		return scope == "" || (annotations != nil && annotations["scope"] == scope)
	}
	for _, account := range resp.Payload.Items {
		if account.Metadata != nil && account.Metadata.Name == name && matchScope(account.Metadata.Annotations) {
			return account, nil
		}
	}
	return nil, fmt.Errorf("MAAS cloud account with name %q not found", name)
}

// GetCloudAccountVsphereByName retrieves a vSphere cloud account by name (and optional scope) using the list GET.
// scope can be "" (any), "project", or "tenant".
func (h *V1Client) GetCloudAccountVsphereByName(name, scope string) (*models.V1VsphereAccount, error) {
	params := clientv1.NewV1CloudAccountsVsphereListParamsWithContext(h.ctx).
		WithLimit(apiutil.Ptr(int64(0)))
	resp, err := h.Client.V1CloudAccountsVsphereList(params)
	if err != nil {
		return nil, err
	}
	matchScope := func(annotations map[string]string) bool {
		return scope == "" || (annotations != nil && annotations["scope"] == scope)
	}
	for _, account := range resp.Payload.Items {
		if account.Metadata != nil && account.Metadata.Name == name && matchScope(account.Metadata.Annotations) {
			return account, nil
		}
	}
	return nil, fmt.Errorf("vSphere cloud account with name %q not found", name)
}

// GetCloudAccountCustomByName retrieves a custom cloud account by name (and optional scope) using the list GET.
func (h *V1Client) GetCloudAccountCustomByName(cloudType, name, scope string) (*models.V1CustomAccount, error) {
	params := clientv1.NewV1CloudAccountsCustomListParamsWithContext(h.ctx).
		WithLimit(apiutil.Ptr(int64(0))).
		WithCloudType(cloudType)
	resp, err := h.Client.V1CloudAccountsCustomList(params)
	if err != nil {
		return nil, err
	}
	matchScope := func(annotations map[string]string) bool {
		return scope == "" || (annotations != nil && annotations["scope"] == scope)
	}
	for _, account := range resp.Payload.Items {
		if account.Metadata != nil && account.Metadata.Name == name && matchScope(account.Metadata.Annotations) {
			return account, nil
		}
	}
	return nil, fmt.Errorf("custom cloud account with name %q for cloud %q not found", name, cloudType)
}
