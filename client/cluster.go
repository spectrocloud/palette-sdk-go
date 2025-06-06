package client

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	clientv1 "github.com/spectrocloud/palette-sdk-go/api/client/version1"
	"github.com/spectrocloud/palette-sdk-go/api/models"
	"github.com/spectrocloud/palette-sdk-go/client/apiutil"
	"github.com/spectrocloud/palette-sdk-go/client/herr"
)

// DeleteCluster deletes a cluster.
func (h *V1Client) DeleteCluster(uid string) error {
	cluster, err := h.GetCluster(uid)
	if err != nil || cluster == nil {
		return err
	}
	params := clientv1.NewV1SpectroClustersDeleteParamsWithContext(h.ctx).
		WithUID(uid)
	_, err = h.Client.V1SpectroClustersDelete(params)
	return err
}

// ForceDeleteCluster force-deletes a cluster, which may orphan cloud resources.
func (h *V1Client) ForceDeleteCluster(uid string, forceDelete bool) error {
	cluster, err := h.GetCluster(uid)
	if err != nil || cluster == nil {
		return err
	}
	params := clientv1.NewV1SpectroClustersDeleteParamsWithContext(h.ctx).
		WithUID(uid).
		WithForceDelete(&forceDelete)
	_, err = h.Client.V1SpectroClustersDelete(params)
	return err
}

// GetCluster retrieves an existing cluster. Returns nil if the cluster's status is nil or the cluster has been deleted.
func (h *V1Client) GetCluster(uid string) (*models.V1SpectroCluster, error) {
	cluster, err := h.GetClusterWithoutStatus(uid)
	if err != nil {
		return nil, err
	}
	if cluster == nil || cluster.Status == nil || cluster.Status.State == "Deleted" {
		return nil, nil
	}
	return cluster, nil
}

// GetClusterWithoutStatus retrieves an existing cluster regardless of its status.
func (h *V1Client) GetClusterWithoutStatus(uid string) (*models.V1SpectroCluster, error) {
	params := clientv1.NewV1SpectroClustersGetParamsWithContext(h.ctx).
		WithUID(uid)
	resp, err := h.Client.V1SpectroClustersGet(params)
	if apiutil.Is404(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// GetClusterByName retrieves an existing cluster by name. Returns nil if the cluster's status is nil or the cluster has been deleted.
func (h *V1Client) GetClusterByName(name string, virtual bool) (*models.V1SpectroCluster, error) {
	filters := []*models.V1SearchFilterItem{clusterNameEqFilter(name)}
	clusterSummaries, err := h.SearchClusterSummaries(getClusterFilter(filters, virtual), nil)
	if err != nil {
		return nil, err
	}
	if len(clusterSummaries) != 1 {
		return nil, fmt.Errorf("expected 1 cluster: got %d", len(clusterSummaries))
	}
	cluster, err := h.GetCluster(clusterSummaries[0].Metadata.UID)
	if err != nil {
		return nil, err
	}
	return cluster, nil
}

// GetClusterSummary retrieves an existing cluster's summary.
func (h *V1Client) GetClusterSummary(clusterID string) (*models.V1SpectroClusterUIDSummary, error) {
	params := clientv1.NewV1SpectroClustersSummaryUIDParamsWithContext(h.ctx).
		WithUID(clusterID)
	resp, err := h.Client.V1SpectroClustersSummaryUID(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// GetClusterOverview retrieves an existing cluster's overview.
func (h *V1Client) GetClusterOverview(clusterID string) (*models.V1SpectroClusterUIDSummary, error) {
	params := clientv1.NewV1SpectroClustersSummaryUIDOverviewParamsWithContext(h.ctx).
		WithUID(clusterID)
	resp, err := h.Client.V1SpectroClustersSummaryUIDOverview(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// SearchClusterSummaries retrieves a list of cluster summaries based on the provided filter and sort criteria.
func (h *V1Client) SearchClusterSummaries(filter *models.V1SearchFilterSpec, sort []*models.V1SearchFilterSortSpec) ([]*models.V1SpectroClusterSummary, error) {
	params := clientv1.NewV1SpectroClustersSearchFilterSummaryParamsWithContext(h.ctx).
		WithBody(&models.V1SearchFilterSummarySpec{
			Filter: filter,
			Sort:   sort,
		})

	var clusters []*models.V1SpectroClusterSummary
	for {
		resp, err := h.Client.V1SpectroClustersSearchFilterSummary(params)
		if apiutil.Is404(err) {
			return nil, nil
		} else if err != nil {
			return nil, err
		}

		clusters = append(clusters, resp.Payload.Items...)
		if resp.Payload.Listmeta.Continue == "" {
			break
		}

		params = params.WithContinue(&resp.Payload.Listmeta.Continue)
	}

	return clusters, nil
}

// GetClusterKubeConfig retrieves a cluster's kubeconfig.
// Deprecated: use GetClusterClientKubeConfig() instead.
func (h *V1Client) GetClusterKubeConfig(uid string) (string, error) {
	builder := &strings.Builder{}
	params := clientv1.NewV1SpectroClustersUIDKubeConfigParamsWithContext(h.ctx).
		WithUID(uid)
	_, err := h.Client.V1SpectroClustersUIDKubeConfig(params, builder)
	if err != nil {
		if herr.IsNotFound(err) {
			return "", nil
		}
		return "", err
	}
	return parseKubeconfig(builder)
}

// GetClusterClientKubeConfig retrieves a cluster's client kubeconfig.
func (h *V1Client) GetClusterClientKubeConfig(uid string) (string, error) {
	builder := &strings.Builder{}
	params := clientv1.NewV1SpectroClustersUIDKubeConfigClientGetParamsWithContext(h.ctx).
		WithUID(uid)
	_, err := h.Client.V1SpectroClustersUIDKubeConfigClientGet(params, builder)
	if err != nil {
		if herr.IsNotFound(err) {
			return "", nil
		}
		return "", err
	}
	return parseKubeconfig(builder)
}

// GetClusterAdminKubeConfig retrieves a cluster's admin kubeconfig.
func (h *V1Client) GetClusterAdminKubeConfig(uid string) (string, error) {
	builder := &strings.Builder{}
	params := clientv1.NewV1SpectroClustersUIDAdminKubeConfigParamsWithContext(h.ctx).
		WithUID(uid)
	_, err := h.Client.V1SpectroClustersUIDAdminKubeConfig(params, builder)
	if err != nil {
		if herr.IsNotFound(err) {
			return "", nil
		}
		return "", err
	}
	return parseKubeconfig(builder)
}

func parseKubeconfig(builder *strings.Builder) (string, error) {
	kubeconfig := builder.String()
	if apiutil.IsBase64(kubeconfig) {
		bs, err := apiutil.Base64DecodeString(kubeconfig)
		if err != nil {
			return "", err
		}
		return string(bs), nil
	}
	return kubeconfig, nil
}

// GetClusterImportManifest retrieves a cluster's import manifest.
func (h *V1Client) GetClusterImportManifest(uid string) (string, error) {
	builder := &strings.Builder{}
	params := clientv1.NewV1SpectroClustersUIDImportManifestParamsWithContext(h.ctx).
		WithUID(uid)
	_, err := h.Client.V1SpectroClustersUIDImportManifest(params, builder)
	if err != nil {
		return "", err
	}
	return builder.String(), nil
}

// UpdateClusterProfileValues updates a cluster's profile values.
func (h *V1Client) UpdateClusterProfileValues(uid string, profiles *models.V1SpectroClusterProfiles) error {
	params := clientv1.NewV1SpectroClustersUpdateProfilesParamsWithContext(h.ctx).
		WithUID(uid).
		WithBody(profiles).
		WithResolveNotification(apiutil.Ptr(true))
	_, err := h.Client.V1SpectroClustersUpdateProfiles(params)
	return err
}

// ImportClusterGeneric imports a cluster using a generic import manifest.
func (h *V1Client) ImportClusterGeneric(meta *models.V1ObjectMetaInputEntity) (string, error) {
	params := clientv1.NewV1SpectroClustersGenericImportParamsWithContext(h.ctx).
		WithBody(&models.V1SpectroGenericClusterImportEntity{
			Metadata: meta,
		})
	resp, err := h.Client.V1SpectroClustersGenericImport(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

// ApproveClusterRepave approves a cluster repave.
func (h *V1Client) ApproveClusterRepave(clusterUID string) error {
	params := clientv1.NewV1SpectroClustersUIDRepaveApproveUpdateParamsWithContext(h.ctx).
		WithUID(clusterUID)
	_, err := h.Client.V1SpectroClustersUIDRepaveApproveUpdate(params)
	return err
}

// ValidateClusterRepave validates if cluster gets repaved for the specified packs.
func (h *V1Client) ValidateClusterRepave(clusterUID string, body *models.V1SpectroClusterPacksEntity) (*models.V1SpectroClusterRepaveValidationResponse, error) {
	params := clientv1.NewV1SpectroClustersUIDValidateRepaveParamsWithContext(h.ctx).
		WithUID(clusterUID).
		WithBody(body)
	resp, err := h.Client.V1SpectroClustersUIDValidateRepave(params)
	return resp.Payload, err
}

// GetRepaveReasons retrieves a cluster's repave reasons.
func (h *V1Client) GetRepaveReasons(clusterUID string) ([]string, error) {
	params := clientv1.NewV1SpectroClustersUIDRepaveGetParamsWithContext(h.ctx).
		WithUID(clusterUID)
	resp, err := h.Client.V1SpectroClustersUIDRepaveGet(params)
	if err != nil {
		return nil, err
	}
	reasons := make([]string, 0, len(resp.Payload.Spec.Reasons))
	for _, r := range resp.Payload.Spec.Reasons {
		reasons = append(reasons, fmt.Sprintf("Repave - Code: %s, Reason: %s", r.Code, r.Message))
	}
	return reasons, nil
}

func getClusterFilter(extraFilters []*models.V1SearchFilterItem, virtual bool) *models.V1SearchFilterSpec {
	filter := &models.V1SearchFilterSpec{
		Conjunction: and(),
		FilterGroups: []*models.V1SearchFilterGroup{
			{
				Conjunction: and(),
				Filters: []*models.V1SearchFilterItem{
					{
						Condition: &models.V1SearchFilterCondition{
							String: &models.V1SearchFilterStringCondition{
								Match: &models.V1SearchFilterStringConditionMatch{
									Conjunction: or(),
									Values:      []string{"nested"},
								},
								Negation: !virtual,
								Operator: models.V1SearchFilterStringOperatorEq,
							},
						},
						Property: "environment",
						Type:     models.V1SearchFilterPropertyTypeString,
					},
					{
						Condition: &models.V1SearchFilterCondition{
							Bool: &models.V1SearchFilterBoolCondition{
								Value: false,
							},
						},
						Property: "isDeleted",
						Type:     models.V1SearchFilterPropertyTypeBool,
					},
				},
			},
		},
	}

	filter.FilterGroups = append(filter.FilterGroups, &models.V1SearchFilterGroup{Conjunction: and(), Filters: extraFilters})

	return filter
}

func clusterNameEqFilter(name string) *models.V1SearchFilterItem {
	return &models.V1SearchFilterItem{
		Condition: &models.V1SearchFilterCondition{
			String: &models.V1SearchFilterStringCondition{
				Match: &models.V1SearchFilterStringConditionMatch{
					Conjunction: or(),
					Values:      []string{name},
				},
				Operator:   models.V1SearchFilterStringOperatorEq,
				Negation:   false,
				IgnoreCase: false,
			},
		},
		Property: "clusterName",
		Type:     models.V1SearchFilterPropertyTypeString,
	}
}

// GetLogFetcherStatus retrieves a cluster's log fetcher status.
func (h *V1Client) GetLogFetcherStatus(uid string, logFetcherUID *string) (*models.V1ClusterLogFetcher, error) {
	params := clientv1.NewV1ClusterFeatureLogFetcherGetParamsWithContext(h.ctx).WithUID(uid).WithRequestID(logFetcherUID)
	resp, err := h.Client.V1ClusterFeatureLogFetcherGet(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// InitiateDownloadOfClusterLogs initiates the download of a cluster's logs.
// Returns a log fetcher uid which is required to perform the log download.
// Before downloading logs, the log fetcher's status should be checked.
func (h *V1Client) InitiateDownloadOfClusterLogs(uid string, req *models.V1ClusterLogFetcherRequest) (*string, error) {
	params := clientv1.NewV1ClusterFeatureLogFetcherCreateParamsWithContext(h.ctx).WithUID(uid).WithBody(req)
	resp, err := h.Client.V1ClusterFeatureLogFetcherCreate(params)
	if err != nil {
		return nil, err
	}
	return resp.GetPayload().UID, nil
}

// DownloadLogs downloads a cluster's logs, given a cluster uid and a log fetcher uid generated via InitiateDownloadOfClusterLogs.
func (h *V1Client) DownloadLogs(uid string, logFetcherUID string) (io.Writer, error) {
	filename := "logs-" + uid + ".zip"
	var buf bytes.Buffer
	writer := io.Writer(&buf)

	params := clientv1.NewV1ClusterFeatureLogFetcherLogDownloadParamsWithContext(h.ctx).WithUID(logFetcherUID).WithFileName(&filename)
	resp, err := h.Client.V1ClusterFeatureLogFetcherLogDownload(params, writer)
	if err != nil {
		return nil, err
	}
	logfile := resp.GetPayload()

	fileLocation := filepath.Join(os.TempDir(), filename)
	fo, err := os.Create(filepath.Clean(fileLocation))
	if err != nil {
		return nil, fmt.Errorf("error while creating a file %v", err)
	}
	_, err = buf.WriteTo(fo)
	if err != nil {
		return nil, fmt.Errorf("error while writing log content to a file %v", err)
	}

	return logfile, nil
}

// UpdatePauseAgentUpgradeSettingCluster updates the upgrade settings of a specific cluster to pause the agent upgrade process.
func (h *V1Client) UpdatePauseAgentUpgradeSettingCluster(upgradeSetting *models.V1ClusterUpgradeSettingsEntity, clusterUID string) error {
	params := clientv1.NewV1SpectroClustersUIDUpgradeSettingsParamsWithContext(h.ctx).WithUID(clusterUID).WithBody(upgradeSetting)
	_, err := h.Client.V1SpectroClustersUIDUpgradeSettings(params)
	if err != nil {
		return err
	}
	return nil
}

// InitiateTheCertRenewal initiates the certificate renewal process for the specified cluster.
func (h *V1Client) InitiateTheCertRenewal(clusterUID string) error {
	params := clientv1.NewV1SpectroClustersCertificatesRenewParamsWithContext(h.ctx).WithUID(clusterUID)
	_, err := h.Client.V1SpectroClustersCertificatesRenew(params)
	if err != nil {
		return fmt.Errorf("error while renewing the cluster certificates: %w", err)
	}
	return nil
}

// GetTheKubernetesCerts retrieves the Kubernetes certificates for the specified cluster.
func (h *V1Client) GetTheKubernetesCerts(clusterUID string) (*models.V1MachineCertificates, error) {
	params := clientv1.NewV1SpectroClustersK8CertificateParamsWithContext(h.ctx).WithUID(clusterUID)
	resp, err := h.Client.V1SpectroClustersK8Certificate(params)
	if err != nil {
		return nil, fmt.Errorf("error while getting the cluster certificates: %w", err)
	}
	certList := resp.GetPayload()
	// Check if the list contains any machine certificates
	if len(certList.MachineCertificates) == 0 {
		fmt.Println("No machine certificates found.")
		return nil, nil
	}

	return certList, nil
}

// UpdateClusterProfileVariableInCluster Update cluster profile variable in running cluster.
func (h *V1Client) UpdateClusterProfileVariableInCluster(clusterUID string, body []*models.V1SpectroClusterVariableUpdateEntity) error {
	params := clientv1.NewV1SpectroClustersUIDVariablesPatchParamsWithContext(h.ctx).WithUID(clusterUID).WithBody(body)
	_, err := h.Client.V1SpectroClustersUIDVariablesPatch(params)
	return err
}
