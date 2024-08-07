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

// NewV1AppProfilesUIDTiersUIDManifestsUIDDeleteParams creates a new V1AppProfilesUIDTiersUIDManifestsUIDDeleteParams object
// with the default values initialized.
func NewV1AppProfilesUIDTiersUIDManifestsUIDDeleteParams() *V1AppProfilesUIDTiersUIDManifestsUIDDeleteParams {
	var ()
	return &V1AppProfilesUIDTiersUIDManifestsUIDDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1AppProfilesUIDTiersUIDManifestsUIDDeleteParamsWithTimeout creates a new V1AppProfilesUIDTiersUIDManifestsUIDDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1AppProfilesUIDTiersUIDManifestsUIDDeleteParamsWithTimeout(timeout time.Duration) *V1AppProfilesUIDTiersUIDManifestsUIDDeleteParams {
	var ()
	return &V1AppProfilesUIDTiersUIDManifestsUIDDeleteParams{

		timeout: timeout,
	}
}

// NewV1AppProfilesUIDTiersUIDManifestsUIDDeleteParamsWithContext creates a new V1AppProfilesUIDTiersUIDManifestsUIDDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1AppProfilesUIDTiersUIDManifestsUIDDeleteParamsWithContext(ctx context.Context) *V1AppProfilesUIDTiersUIDManifestsUIDDeleteParams {
	var ()
	return &V1AppProfilesUIDTiersUIDManifestsUIDDeleteParams{

		Context: ctx,
	}
}

// NewV1AppProfilesUIDTiersUIDManifestsUIDDeleteParamsWithHTTPClient creates a new V1AppProfilesUIDTiersUIDManifestsUIDDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1AppProfilesUIDTiersUIDManifestsUIDDeleteParamsWithHTTPClient(client *http.Client) *V1AppProfilesUIDTiersUIDManifestsUIDDeleteParams {
	var ()
	return &V1AppProfilesUIDTiersUIDManifestsUIDDeleteParams{
		HTTPClient: client,
	}
}

/*V1AppProfilesUIDTiersUIDManifestsUIDDeleteParams contains all the parameters to send to the API endpoint
for the v1 app profiles Uid tiers Uid manifests Uid delete operation typically these are written to a http.Request
*/
type V1AppProfilesUIDTiersUIDManifestsUIDDeleteParams struct {

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

// WithTimeout adds the timeout to the v1 app profiles Uid tiers Uid manifests Uid delete params
func (o *V1AppProfilesUIDTiersUIDManifestsUIDDeleteParams) WithTimeout(timeout time.Duration) *V1AppProfilesUIDTiersUIDManifestsUIDDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 app profiles Uid tiers Uid manifests Uid delete params
func (o *V1AppProfilesUIDTiersUIDManifestsUIDDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 app profiles Uid tiers Uid manifests Uid delete params
func (o *V1AppProfilesUIDTiersUIDManifestsUIDDeleteParams) WithContext(ctx context.Context) *V1AppProfilesUIDTiersUIDManifestsUIDDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 app profiles Uid tiers Uid manifests Uid delete params
func (o *V1AppProfilesUIDTiersUIDManifestsUIDDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 app profiles Uid tiers Uid manifests Uid delete params
func (o *V1AppProfilesUIDTiersUIDManifestsUIDDeleteParams) WithHTTPClient(client *http.Client) *V1AppProfilesUIDTiersUIDManifestsUIDDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 app profiles Uid tiers Uid manifests Uid delete params
func (o *V1AppProfilesUIDTiersUIDManifestsUIDDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithManifestUID adds the manifestUID to the v1 app profiles Uid tiers Uid manifests Uid delete params
func (o *V1AppProfilesUIDTiersUIDManifestsUIDDeleteParams) WithManifestUID(manifestUID string) *V1AppProfilesUIDTiersUIDManifestsUIDDeleteParams {
	o.SetManifestUID(manifestUID)
	return o
}

// SetManifestUID adds the manifestUid to the v1 app profiles Uid tiers Uid manifests Uid delete params
func (o *V1AppProfilesUIDTiersUIDManifestsUIDDeleteParams) SetManifestUID(manifestUID string) {
	o.ManifestUID = manifestUID
}

// WithTierUID adds the tierUID to the v1 app profiles Uid tiers Uid manifests Uid delete params
func (o *V1AppProfilesUIDTiersUIDManifestsUIDDeleteParams) WithTierUID(tierUID string) *V1AppProfilesUIDTiersUIDManifestsUIDDeleteParams {
	o.SetTierUID(tierUID)
	return o
}

// SetTierUID adds the tierUid to the v1 app profiles Uid tiers Uid manifests Uid delete params
func (o *V1AppProfilesUIDTiersUIDManifestsUIDDeleteParams) SetTierUID(tierUID string) {
	o.TierUID = tierUID
}

// WithUID adds the uid to the v1 app profiles Uid tiers Uid manifests Uid delete params
func (o *V1AppProfilesUIDTiersUIDManifestsUIDDeleteParams) WithUID(uid string) *V1AppProfilesUIDTiersUIDManifestsUIDDeleteParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 app profiles Uid tiers Uid manifests Uid delete params
func (o *V1AppProfilesUIDTiersUIDManifestsUIDDeleteParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1AppProfilesUIDTiersUIDManifestsUIDDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
