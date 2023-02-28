package client

import (
	"context"
	"errors"
	"log"
	"time"

	openapiclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/spectrocloud/hapi/apiutil/transport"
	authC "github.com/spectrocloud/hapi/auth/client/v1"
	cloudC "github.com/spectrocloud/hapi/cloud/client/v1"
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

	AuthClient authC.ClientService

	schemes []string

	authToken   *AuthToken
	tokenExpiry = 10 * time.Minute
)

type AuthToken struct {
	token  *models.V1UserToken
	expiry time.Time
}

type V1Client struct {
	Ctx            context.Context
	email          string
	password       string
	apikey         string
	transportDebug bool
	RetryAttempts  int

	// Cluster client(common)
	GetClusterClientFn func() (clusterC.ClientService, error)

	// Cluster generic
	GetClusterWithoutStatusFn   func(string) (*models.V1SpectroCluster, error)
	GetClusterFn                func(uid string) (*models.V1SpectroCluster, error)
	GetClusterKubeConfigFn      func(uid string) (string, error)
	GetClusterBackupConfigFn    func(uid string) (*models.V1ClusterBackup, error)
	GetClusterScanConfigFn      func(uid string) (*models.V1ClusterComplianceScan, error)
	GetClusterRbacConfigFn      func(uid string) (*models.V1ClusterRbacs, error)
	GetClusterNamespaceConfigFn func(uid string) (*models.V1ClusterNamespaceResources, error)

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

	// Cluster profiles
	ClustersPatchProfilesFn func(*clusterC.V1SpectroClustersPatchProfilesParams) error // used in application deployment
	GetClusterProfileFn     func(string) (*models.V1ClusterProfile, error)
	DeleteClusterProfileFn  func(string) error

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

	//Alert
	CreateAlertFn  func(body *models.V1Channel, projectUID string, component string) (string, error)
	UpdateAlertFn  func(body *models.V1Channel, projectUID string, component string, alertUID string) (string, error)
	ReadAlertFn    func(projectUID string, component string, alertUID string) (*models.V1Channel, error)
	DeleteAlertsFn func(projectUID string, component string, alertUID string) error
}

func New(hubbleHost, email, password, projectUID string, apikey string, transportDebug bool, retryAttempts int) *V1Client {
	ctx := context.Background()
	if projectUID != "" {
		ctx = GetProjectContextWithCtx(ctx, projectUID)
	}

	hubbleUri = hubbleHost
	schemes = []string{"https"}
	authHttpTransport := transport.New(hubbleUri, "", schemes)
	authHttpTransport.RetryAttempts = 0
	//authHttpTransport.Debug = true
	AuthClient = authC.New(authHttpTransport, strfmt.Default)
	return &V1Client{Ctx: ctx, email: email, password: password, apikey: apikey, transportDebug: transportDebug, RetryAttempts: retryAttempts}
}

func (h *V1Client) getNewAuthToken() (*AuthToken, error) {
	//httpClient, err := certs.GetHttpClient()
	//if err != nil {
	//	return nil, err
	//}
	authParam := authC.NewV1AuthenticateParams().
		WithBody(&models.V1AuthLogin{
			EmailID:  h.email,
			Password: strfmt.Password(h.password),
		})
	res, err := AuthClient.V1Authenticate(authParam)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	if len(res.Payload.Authorization) == 0 {
		errMsg := "authorization auth token is empty in hubble payload"
		return nil, errors.New(errMsg)
	}

	authToken = &AuthToken{
		token:  res.Payload,
		expiry: time.Now().Add(tokenExpiry),
	}
	return authToken, nil
}

func GetProjectContextWithCtx(c context.Context, projectUid string) context.Context {
	return context.WithValue(c, transport.CUSTOM_HEADERS, transport.Values{
		HeaderMap: map[string]string{
			"ProjectUid": projectUid,
		}})
}

func (h *V1Client) getTransport() (*transport.Runtime, error) {
	if h.apikey == "" && (authToken == nil || authToken.expiry.Before(time.Now())) {
		if tkn, err := h.getNewAuthToken(); err != nil {
			log.Fatal("Failed to get auth token ", err.Error())
			return nil, err
		} else {
			authToken = tkn
		}
	}

	httpTransport := transport.New(hubbleUri, "", schemes)
	if h.apikey != "" {
		httpTransport.DefaultAuthentication = openapiclient.APIKeyAuth(authApiKey, authTokenInput, h.apikey)
	} else {
		httpTransport.DefaultAuthentication = openapiclient.APIKeyAuth(authTokenKey, authTokenInput, authToken.token.Authorization)
	}
	httpTransport.RetryAttempts = h.RetryAttempts
	httpTransport.Debug = h.transportDebug
	return httpTransport, nil
}

// Clients
func (h *V1Client) GetClusterClient() (clusterC.ClientService, error) {
	if h.GetClusterClientFn != nil {
		return h.GetClusterClientFn()
	}
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

func (h *V1Client) GetHashboard() (hashboardC.ClientService, error) {
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

func (h *V1Client) Validate() error {
	authToken = nil
	_, err := h.getTransport()
	return err
}
