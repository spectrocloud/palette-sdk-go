package client

import (
	"context"
	"errors"

	openapiclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/spectrocloud/hapi/apiutil/transport"
	authC "github.com/spectrocloud/hapi/auth/client/v1"
	cloudC "github.com/spectrocloud/hapi/cloud/client/v1"
	eventC "github.com/spectrocloud/hapi/event/client/v1"
	hashboardC "github.com/spectrocloud/hapi/hashboard/client/v1"
	"github.com/spectrocloud/hapi/models"
	clusterC "github.com/spectrocloud/hapi/spectrocluster/client/v1"
	userC "github.com/spectrocloud/hapi/user/client/v1"
)

const (
	authTokenInput string = "header"
	authTokenKey   string = "Authorization"
	authApiKey     string = "ApiKey"
)

var (
	hubbleUri string

	schemes []string
)

type V1Client struct {
	Ctx            context.Context
	apikey         string
	transportDebug bool
	RetryAttempts  int

	ClusterC clusterC.ClientService

	// Cluster generic
	GetClusterWithoutStatusFn   func(string) (*models.V1SpectroCluster, error)
	GetClusterFn                func(scope, uid string) (*models.V1SpectroCluster, error)
	GetClusterAdminConfigFn     func(uid string) (string, error)
	GetClusterKubeConfigFn      func(uid string) (string, error)
	GetClusterBackupConfigFn    func(uid string) (*models.V1ClusterBackup, error)
	GetClusterScanConfigFn      func(uid string) (*models.V1ClusterComplianceScan, error)
	GetClusterRbacConfigFn      func(uid string) (*models.V1ClusterRbacs, error)
	GetClusterNamespaceConfigFn func(uid string) (*models.V1ClusterNamespaceResources, error)
	ApproveClusterRepaveFn      func(context, clusterUID string) error
	GetRepaveReasonsFn          func(context, clusterUID string) ([]string, error)

	// Cluster Groups
	CreateClusterGroupFn func(*models.V1ClusterGroupEntity) (string, error)
	GetClusterGroupFn    func(string) (*models.V1ClusterGroup, error)
	UpdateClusterGroupFn func(string, *models.V1ClusterGroupHostClusterEntity) error
	DeleteClusterGroupFn func(string) error

	// Application
	GetApplicationFn func(string) (*models.V1AppDeployment, error)

	// Application Profile
	GetApplicationProfileTiersFn               func(string) ([]*models.V1AppTier, error)
	CreateApplicationProfileFn                 func(*models.V1AppProfileEntity, string) (string, error)
	GetApplicationProfileTierManifestContentFn func(string, string, string) (string, error)
	GetApplicationProfileFn                    func(string) (*models.V1AppProfile, error)
	DeleteApplicationProfileFn                 func(string) error

	// special function for virtual cluster mock
	V1ClusterProfilesDeleteFn            func(params *clusterC.V1ClusterProfilesDeleteParams) (*clusterC.V1ClusterProfilesDeleteNoContent, error)
	V1ClusterProfilesUIDMetadataUpdateFn func(params *clusterC.V1ClusterProfilesUIDMetadataUpdateParams) (*clusterC.V1ClusterProfilesUIDMetadataUpdateNoContent, error)
	V1ClusterProfilesUpdateFn            func(params *clusterC.V1ClusterProfilesUpdateParams) (*clusterC.V1ClusterProfilesUpdateNoContent, error)
	V1ClusterProfilesCreateFn            func(params *clusterC.V1ClusterProfilesCreateParams) (*clusterC.V1ClusterProfilesCreateCreated, error)
	V1ClusterProfilesPublishFn           func(params *clusterC.V1ClusterProfilesPublishParams) (*models.V1ClusterProfile, error)

	// Registry
	GetPackRegistryCommonByNameFn func(string) (*models.V1RegistryMetadata, error)

	// VSphere Cluster
	CreateClusterVsphereFn        func(*models.V1SpectroVsphereClusterEntity) (string, error)
	GetCloudConfigVsphereFn       func(cloudConfigUid string) (*models.V1VsphereCloudConfig, error)
	GetCloudConfigVsphereValuesFn func(uid string) (*models.V1VsphereCloudConfig, error)

	// Project
	GetProjectUIDFn func(projectName string) (string, error)

	// Alert
	CreateAlertFn  func(body *models.V1Channel, projectUID, component string) (string, error)
	UpdateAlertFn  func(body *models.V1Channel, projectUID, component, alertUID string) (string, error)
	ReadAlertFn    func(projectUID, component, alertUID string) (*models.V1Channel, error)
	DeleteAlertsFn func(projectUID, component, alertUID string) error

	// Virtual Machines
	GetVirtualMachineWithoutStatusFn func(string) (*models.V1ClusterVirtualMachine, error)
	GetVirtualMachineFn              func(uid string) (*models.V1ClusterVirtualMachine, error)

	// Data Volumes
	CreateDataVolumeFn func(uid, name string, body *models.V1VMAddVolumeEntity) (string, error)
	DeleteDataVolumeFn func(uid, namespace, name string, body *models.V1VMRemoveVolumeEntity) error

	// Registry
	CreateOciEcrRegistryFn func(registry *models.V1EcrRegistry) (string, error)
	UpdateEcrRegistryFn    func(uid string, registry *models.V1EcrRegistry) error
	GetOciRegistryFn       func(uid string) (*models.V1EcrRegistry, error)
	DeleteOciEcrRegistryFn func(uid string) error

	// Edge Native
	GetCloudConfigEdgeNativeFn func(uid, clusterContext string) (*models.V1EdgeNativeCloudConfig, error)
}

func New(hubbleHost, projectUID, apikey string, transportDebug bool, retryAttempts int) *V1Client {
	ctx := context.Background()
	if projectUID != "" {
		ctx = GetProjectContextWithCtx(ctx, projectUID)
	}

	hubbleUri = hubbleHost
	schemes = []string{"https"}
	authHttpTransport := transport.New(hubbleUri, "", schemes)
	authHttpTransport.RetryAttempts = 0
	//authHttpTransport.Debug = true
	return &V1Client{
		Ctx:            ctx,
		apikey:         apikey,
		transportDebug: transportDebug,
		RetryAttempts:  retryAttempts,
	}
}

func (h *V1Client) getTransport() (*transport.Runtime, error) {
	httpTransport := h.baseTransport()
	if h.apikey != "" {
		httpTransport.DefaultAuthentication = openapiclient.APIKeyAuth(authApiKey, authTokenInput, h.apikey)
	} else {
		return nil, errors.New("api_key is required")
	}
	return httpTransport, nil
}

func (h *V1Client) baseTransport() *transport.Runtime {
	httpTransport := transport.New(hubbleUri, "", schemes)
	httpTransport.RetryAttempts = h.RetryAttempts
	httpTransport.Debug = h.transportDebug
	return httpTransport
}

// Clients
func (h *V1Client) GetAuthClient() (authC.ClientService, error) {
	httpTransport, err := h.getTransport()
	if err != nil {
		return nil, err
	}

	return authC.New(httpTransport, strfmt.Default), nil
}

func (h *V1Client) GetNoAuthClient() (authC.ClientService, error) {
	httpTransport := h.baseTransport()
	return authC.New(httpTransport, strfmt.Default), nil
}

func (h *V1Client) GetClusterClient() (clusterC.ClientService, error) {
	httpTransport, err := h.getTransport()
	if err != nil {
		return nil, err
	}

	return clusterC.New(httpTransport, strfmt.Default), nil
}

func (h *V1Client) GetUserClient() (userC.ClientService, error) {
	httpTransport, err := h.getTransport()
	if err != nil {
		return nil, err
	}

	return userC.New(httpTransport, strfmt.Default), nil
}

func (h *V1Client) GetHashboardClient() (hashboardC.ClientService, error) {
	httpTransport, err := h.getTransport()
	if err != nil {
		return nil, err
	}

	return hashboardC.New(httpTransport, strfmt.Default), nil
}

func (h *V1Client) GetCloudClient() (cloudC.ClientService, error) {
	httpTransport, err := h.getTransport()
	if err != nil {
		return nil, err
	}

	return cloudC.New(httpTransport, strfmt.Default), nil
}

func (h *V1Client) GetEventClient() (eventC.ClientService, error) {
	httpTransport, err := h.getTransport()
	if err != nil {
		return nil, err
	}

	return eventC.New(httpTransport, strfmt.Default), nil
}

func (h *V1Client) Validate() error {
	_, err := h.getTransport()

	// API key can only be validated by making an API call
	if h.apikey != "" {
		c, _ := h.GetUserClient()
		_, err := c.V1ProjectsList(nil)
		if err != nil {
			return err
		}
	}

	return err
}

func (h *V1Client) ValidateTenantAdmin() error {
	_, err := h.getTransport()

	// API key can only be validated by making an API call
	if h.apikey != "" {
		c, _ := h.GetUserClient()
		_, err := c.V1UsersList(nil)
		if err != nil {
			return err
		}
	}

	return err
}

func GetProjectContextWithCtx(c context.Context, projectUid string) context.Context {
	return context.WithValue(c, transport.CUSTOM_HEADERS, transport.Values{
		HeaderMap: map[string]string{
			"ProjectUid": projectUid,
		}})
}
