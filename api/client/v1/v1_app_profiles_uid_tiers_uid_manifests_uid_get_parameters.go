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

// NewV1AppProfilesUIDTiersUIDManifestsUIDGetParams creates a new V1AppProfilesUIDTiersUIDManifestsUIDGetParams object
// with the default values initialized.
func NewV1AppProfilesUIDTiersUIDManifestsUIDGetParams() *V1AppProfilesUIDTiersUIDManifestsUIDGetParams {
	var ()
	return &V1AppProfilesUIDTiersUIDManifestsUIDGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1AppProfilesUIDTiersUIDManifestsUIDGetParamsWithTimeout creates a new V1AppProfilesUIDTiersUIDManifestsUIDGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1AppProfilesUIDTiersUIDManifestsUIDGetParamsWithTimeout(timeout time.Duration) *V1AppProfilesUIDTiersUIDManifestsUIDGetParams {
	var ()
	return &V1AppProfilesUIDTiersUIDManifestsUIDGetParams{

		timeout: timeout,
	}
}

// NewV1AppProfilesUIDTiersUIDManifestsUIDGetParamsWithContext creates a new V1AppProfilesUIDTiersUIDManifestsUIDGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1AppProfilesUIDTiersUIDManifestsUIDGetParamsWithContext(ctx context.Context) *V1AppProfilesUIDTiersUIDManifestsUIDGetParams {
	var ()
	return &V1AppProfilesUIDTiersUIDManifestsUIDGetParams{

		Context: ctx,
	}
}

// NewV1AppProfilesUIDTiersUIDManifestsUIDGetParamsWithHTTPClient creates a new V1AppProfilesUIDTiersUIDManifestsUIDGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1AppProfilesUIDTiersUIDManifestsUIDGetParamsWithHTTPClient(client *http.Client) *V1AppProfilesUIDTiersUIDManifestsUIDGetParams {
	var ()
	return &V1AppProfilesUIDTiersUIDManifestsUIDGetParams{
		HTTPClient: client,
	}
}

/*
V1AppProfilesUIDTiersUIDManifestsUIDGetParams contains all the parameters to send to the API endpoint
for the v1 app profiles Uid tiers Uid manifests Uid get operation typically these are written to a http.Request
*/
type V1AppProfilesUIDTiersUIDManifestsUIDGetParams struct {

	/*ManifestUID
	  Application profile tier manifest uid

	*/
	ManifestUID string
	/*TierUID
	  Application profile tier uid

	*/
	TierUID string
	/*UID
	  Application profile uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 app profiles Uid tiers Uid manifests Uid get params
func (o *V1AppProfilesUIDTiersUIDManifestsUIDGetParams) WithTimeout(timeout time.Duration) *V1AppProfilesUIDTiersUIDManifestsUIDGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 app profiles Uid tiers Uid manifests Uid get params
func (o *V1AppProfilesUIDTiersUIDManifestsUIDGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 app profiles Uid tiers Uid manifests Uid get params
func (o *V1AppProfilesUIDTiersUIDManifestsUIDGetParams) WithContext(ctx context.Context) *V1AppProfilesUIDTiersUIDManifestsUIDGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 app profiles Uid tiers Uid manifests Uid get params
func (o *V1AppProfilesUIDTiersUIDManifestsUIDGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 app profiles Uid tiers Uid manifests Uid get params
func (o *V1AppProfilesUIDTiersUIDManifestsUIDGetParams) WithHTTPClient(client *http.Client) *V1AppProfilesUIDTiersUIDManifestsUIDGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 app profiles Uid tiers Uid manifests Uid get params
func (o *V1AppProfilesUIDTiersUIDManifestsUIDGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithManifestUID adds the manifestUID to the v1 app profiles Uid tiers Uid manifests Uid get params
func (o *V1AppProfilesUIDTiersUIDManifestsUIDGetParams) WithManifestUID(manifestUID string) *V1AppProfilesUIDTiersUIDManifestsUIDGetParams {
	o.SetManifestUID(manifestUID)
	return o
}

// SetManifestUID adds the manifestUid to the v1 app profiles Uid tiers Uid manifests Uid get params
func (o *V1AppProfilesUIDTiersUIDManifestsUIDGetParams) SetManifestUID(manifestUID string) {
	o.ManifestUID = manifestUID
}

// WithTierUID adds the tierUID to the v1 app profiles Uid tiers Uid manifests Uid get params
func (o *V1AppProfilesUIDTiersUIDManifestsUIDGetParams) WithTierUID(tierUID string) *V1AppProfilesUIDTiersUIDManifestsUIDGetParams {
	o.SetTierUID(tierUID)
	return o
}

// SetTierUID adds the tierUid to the v1 app profiles Uid tiers Uid manifests Uid get params
func (o *V1AppProfilesUIDTiersUIDManifestsUIDGetParams) SetTierUID(tierUID string) {
	o.TierUID = tierUID
}

// WithUID adds the uid to the v1 app profiles Uid tiers Uid manifests Uid get params
func (o *V1AppProfilesUIDTiersUIDManifestsUIDGetParams) WithUID(uid string) *V1AppProfilesUIDTiersUIDManifestsUIDGetParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 app profiles Uid tiers Uid manifests Uid get params
func (o *V1AppProfilesUIDTiersUIDManifestsUIDGetParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1AppProfilesUIDTiersUIDManifestsUIDGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
