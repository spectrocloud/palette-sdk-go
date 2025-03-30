// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewV1AppDeploymentsProfileTiersManifestsUIDGetParams creates a new V1AppDeploymentsProfileTiersManifestsUIDGetParams object
// with the default values initialized.
func NewV1AppDeploymentsProfileTiersManifestsUIDGetParams() *V1AppDeploymentsProfileTiersManifestsUIDGetParams {
	var ()
	return &V1AppDeploymentsProfileTiersManifestsUIDGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1AppDeploymentsProfileTiersManifestsUIDGetParamsWithTimeout creates a new V1AppDeploymentsProfileTiersManifestsUIDGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1AppDeploymentsProfileTiersManifestsUIDGetParamsWithTimeout(timeout time.Duration) *V1AppDeploymentsProfileTiersManifestsUIDGetParams {
	var ()
	return &V1AppDeploymentsProfileTiersManifestsUIDGetParams{

		timeout: timeout,
	}
}

// NewV1AppDeploymentsProfileTiersManifestsUIDGetParamsWithContext creates a new V1AppDeploymentsProfileTiersManifestsUIDGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1AppDeploymentsProfileTiersManifestsUIDGetParamsWithContext(ctx context.Context) *V1AppDeploymentsProfileTiersManifestsUIDGetParams {
	var ()
	return &V1AppDeploymentsProfileTiersManifestsUIDGetParams{

		Context: ctx,
	}
}

// NewV1AppDeploymentsProfileTiersManifestsUIDGetParamsWithHTTPClient creates a new V1AppDeploymentsProfileTiersManifestsUIDGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1AppDeploymentsProfileTiersManifestsUIDGetParamsWithHTTPClient(client *http.Client) *V1AppDeploymentsProfileTiersManifestsUIDGetParams {
	var ()
	return &V1AppDeploymentsProfileTiersManifestsUIDGetParams{
		HTTPClient: client,
	}
}

/*
V1AppDeploymentsProfileTiersManifestsUIDGetParams contains all the parameters to send to the API endpoint
for the v1 app deployments profile tiers manifests Uid get operation typically these are written to a http.Request
*/
type V1AppDeploymentsProfileTiersManifestsUIDGetParams struct {

	/*ManifestUID
	  Application deployment tier manifest uid

	*/
	ManifestUID string
	/*TierUID
	  Application deployment tier uid

	*/
	TierUID string
	/*UID
	  Application deployment uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 app deployments profile tiers manifests Uid get params
func (o *V1AppDeploymentsProfileTiersManifestsUIDGetParams) WithTimeout(timeout time.Duration) *V1AppDeploymentsProfileTiersManifestsUIDGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 app deployments profile tiers manifests Uid get params
func (o *V1AppDeploymentsProfileTiersManifestsUIDGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 app deployments profile tiers manifests Uid get params
func (o *V1AppDeploymentsProfileTiersManifestsUIDGetParams) WithContext(ctx context.Context) *V1AppDeploymentsProfileTiersManifestsUIDGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 app deployments profile tiers manifests Uid get params
func (o *V1AppDeploymentsProfileTiersManifestsUIDGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 app deployments profile tiers manifests Uid get params
func (o *V1AppDeploymentsProfileTiersManifestsUIDGetParams) WithHTTPClient(client *http.Client) *V1AppDeploymentsProfileTiersManifestsUIDGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 app deployments profile tiers manifests Uid get params
func (o *V1AppDeploymentsProfileTiersManifestsUIDGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithManifestUID adds the manifestUID to the v1 app deployments profile tiers manifests Uid get params
func (o *V1AppDeploymentsProfileTiersManifestsUIDGetParams) WithManifestUID(manifestUID string) *V1AppDeploymentsProfileTiersManifestsUIDGetParams {
	o.SetManifestUID(manifestUID)
	return o
}

// SetManifestUID adds the manifestUid to the v1 app deployments profile tiers manifests Uid get params
func (o *V1AppDeploymentsProfileTiersManifestsUIDGetParams) SetManifestUID(manifestUID string) {
	o.ManifestUID = manifestUID
}

// WithTierUID adds the tierUID to the v1 app deployments profile tiers manifests Uid get params
func (o *V1AppDeploymentsProfileTiersManifestsUIDGetParams) WithTierUID(tierUID string) *V1AppDeploymentsProfileTiersManifestsUIDGetParams {
	o.SetTierUID(tierUID)
	return o
}

// SetTierUID adds the tierUid to the v1 app deployments profile tiers manifests Uid get params
func (o *V1AppDeploymentsProfileTiersManifestsUIDGetParams) SetTierUID(tierUID string) {
	o.TierUID = tierUID
}

// WithUID adds the uid to the v1 app deployments profile tiers manifests Uid get params
func (o *V1AppDeploymentsProfileTiersManifestsUIDGetParams) WithUID(uid string) *V1AppDeploymentsProfileTiersManifestsUIDGetParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 app deployments profile tiers manifests Uid get params
func (o *V1AppDeploymentsProfileTiersManifestsUIDGetParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1AppDeploymentsProfileTiersManifestsUIDGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param manifestUid
	if err := r.SetPathParam("manifestUid", o.ManifestUID); err != nil {
		return err
	}

	// path param tierUid
	if err := r.SetPathParam("tierUid", o.TierUID); err != nil {
		return err
	}

	// path param uid
	if err := r.SetPathParam("uid", o.UID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
