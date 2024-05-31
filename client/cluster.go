package client

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	clientV1 "github.com/spectrocloud/palette-api-go/client/v1"
	"github.com/spectrocloud/palette-api-go/models"
	"github.com/spectrocloud/palette-sdk-go/client/apiutil"
	"github.com/spectrocloud/palette-sdk-go/client/herr"
)

func (h *V1Client) DeleteCluster(uid string) error {
	cluster, err := h.GetCluster(uid)
	if err != nil || cluster == nil {
		return err
	}
	params := clientV1.NewV1SpectroClustersDeleteParamsWithContext(h.ctx).
		WithUID(uid)
	_, err = h.Client.V1SpectroClustersDelete(params)
	return err
}

func (h *V1Client) ForceDeleteCluster(uid string, forceDelete bool) error {
	cluster, err := h.GetCluster(uid)
	if err != nil || cluster == nil {
		return err
	}
	params := clientV1.NewV1SpectroClustersDeleteParamsWithContext(h.ctx).
		WithUID(uid).
		WithForceDelete(&forceDelete)
	_, err = h.Client.V1SpectroClustersDelete(params)
	return err
}

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

func (h *V1Client) GetClusterWithoutStatus(uid string) (*models.V1SpectroCluster, error) {
	params := clientV1.NewV1SpectroClustersGetParamsWithContext(h.ctx).
		WithUID(uid)
	resp, err := h.Client.V1SpectroClustersGet(params)
	if apiutil.Is404(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

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

func (h *V1Client) GetClusterSummary(clusterId string) (*models.V1SpectroClusterUIDSummary, error) {
	params := clientV1.NewV1SpectroClustersSummaryUIDParamsWithContext(h.ctx).
		WithUID(clusterId)
	resp, err := h.Client.V1SpectroClustersSummaryUID(params)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

func (h *V1Client) SearchClusterSummaries(filter *models.V1SearchFilterSpec, sort []*models.V1SearchFilterSortSpec) ([]*models.V1SpectroClusterSummary, error) {
	params := clientV1.NewV1SpectroClustersSearchFilterSummaryParamsWithContext(h.ctx).
		WithBody(&models.V1SearchFilterSummarySpec{
			Filter: filter,
			Sort:   sort,
		})
	resp, err := h.Client.V1SpectroClustersSearchFilterSummary(params)
	if apiutil.Is404(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return resp.Payload.Items, nil
}

func (h *V1Client) GetClusterKubeConfig(uid string) (string, error) {
	builder := &strings.Builder{}
	params := clientV1.NewV1SpectroClustersUIDKubeConfigParamsWithContext(h.ctx).
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

func (h *V1Client) GetClusterAdminKubeConfig(uid string) (string, error) {
	builder := &strings.Builder{}
	params := clientV1.NewV1SpectroClustersUIDAdminKubeConfigParamsWithContext(h.ctx).
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
		bytes, err := apiutil.Base64DecodeString(kubeconfig)
		if err != nil {
			return "", err
		}
		return string(bytes), nil
	}
	return kubeconfig, nil
}

func (h *V1Client) GetClusterImportManifest(uid string) (string, error) {
	builder := &strings.Builder{}
	params := clientV1.NewV1SpectroClustersUIDImportManifestParamsWithContext(h.ctx).
		WithUID(uid)
	_, err := h.Client.V1SpectroClustersUIDImportManifest(params, builder)
	if err != nil {
		return "", err
	}
	return builder.String(), nil
}

func (h *V1Client) UpdateClusterProfileValues(uid, context string, profiles *models.V1SpectroClusterProfiles) error {
	params := clientV1.NewV1SpectroClustersUpdateProfilesParamsWithContext(h.ctx).
		WithUID(uid).
		WithBody(profiles).
		WithResolveNotification(apiutil.Ptr(true))
	_, err := h.Client.V1SpectroClustersUpdateProfiles(params)
	return err
}

func (h *V1Client) ImportClusterGeneric(meta *models.V1ObjectMetaInputEntity) (string, error) {
	params := clientV1.NewV1SpectroClustersGenericImportParamsWithContext(h.ctx).
		WithBody(&models.V1SpectroGenericClusterImportEntity{
			Metadata: meta,
		})
	resp, err := h.Client.V1SpectroClustersGenericImport(params)
	if err != nil {
		return "", err
	}
	return *resp.Payload.UID, nil
}

func (h *V1Client) ApproveClusterRepave(clusterUid string) error {
	params := clientV1.NewV1SpectroClustersUIDRepaveApproveUpdateParamsWithContext(h.ctx).
		WithUID(clusterUid)
	_, err := h.Client.V1SpectroClustersUIDRepaveApproveUpdate(params)
	return err
}

func (h *V1Client) GetRepaveReasons(clusterUid string) ([]string, error) {
	params := clientV1.NewV1SpectroClustersUIDRepaveGetParamsWithContext(h.ctx).
		WithUID(clusterUid)
	resp, err := h.Client.V1SpectroClustersUIDRepaveGet(params)
	if err != nil {
		return nil, err
	}
	var reasons []string
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

func (h *V1Client) GetLogFetcherStatus(uid string,logFetcherUid *string) (*models.V1ClusterLogFetcher,error) {
	params2 := clientV1.NewV1ClusterFeatureLogFetcherGetParamsWithContext(h.ctx).WithUID(uid).WithRequestID(logFetcherUid)
	resp2, err := h.Client.V1ClusterFeatureLogFetcherGet(params2)
	if err != nil {
		return nil, err
	}
	return resp2.Payload, nil
}

func (h *V1Client) InitiateDownloadOfClusterLogs(uid string,V1ClusterLogFetcherRequestObj *models.V1ClusterLogFetcherRequest) (*string, error) {
	//param := clientV1.V1ClusterFeatureLogFetcherCreateParams
	fmt.Println("new console log test")
	params := clientV1.NewV1ClusterFeatureLogFetcherCreateParamsWithContext(h.ctx).WithUID(uid).WithBody(V1ClusterLogFetcherRequestObj)
	
	resp, err := h.Client.V1ClusterFeatureLogFetcherCreate(params)
	if err != nil {
		return nil,err
	}
	return resp.GetPayload().UID,nil
	
}

func (h *V1Client) DownloadLogs(uid string,logFetcherUid string) (io.Writer,error){

	filename := "logs-"+uid+".zip"
	params1 := clientV1.NewV1ClusterFeatureLogFetcherLogDownloadParamsWithContext(h.ctx).WithUID(logFetcherUid).WithFileName(&filename)
	var buf bytes.Buffer
	writer := io.Writer(&buf)
	resp1,err1 := h.Client.V1ClusterFeatureLogFetcherLogDownload(params1,writer)
	logfile := resp1.GetPayload()
	if err1 != nil {
		return nil,err1
	}
	file_location := "/tmp/"+filename
	fo, _ := os.Create(file_location)
	buf.WriteTo(fo)
	return logfile,nil
}
