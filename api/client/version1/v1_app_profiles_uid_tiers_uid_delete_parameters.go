// Code generated by go-swagger; DO NOT EDIT.

package version1

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

// NewV1AppProfilesUIDTiersUIDDeleteParams creates a new V1AppProfilesUIDTiersUIDDeleteParams object
// with the default values initialized.
func NewV1AppProfilesUIDTiersUIDDeleteParams() *V1AppProfilesUIDTiersUIDDeleteParams {
	var ()
	return &V1AppProfilesUIDTiersUIDDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1AppProfilesUIDTiersUIDDeleteParamsWithTimeout creates a new V1AppProfilesUIDTiersUIDDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1AppProfilesUIDTiersUIDDeleteParamsWithTimeout(timeout time.Duration) *V1AppProfilesUIDTiersUIDDeleteParams {
	var ()
	return &V1AppProfilesUIDTiersUIDDeleteParams{

		timeout: timeout,
	}
}

// NewV1AppProfilesUIDTiersUIDDeleteParamsWithContext creates a new V1AppProfilesUIDTiersUIDDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1AppProfilesUIDTiersUIDDeleteParamsWithContext(ctx context.Context) *V1AppProfilesUIDTiersUIDDeleteParams {
	var ()
	return &V1AppProfilesUIDTiersUIDDeleteParams{

		Context: ctx,
	}
}

// NewV1AppProfilesUIDTiersUIDDeleteParamsWithHTTPClient creates a new V1AppProfilesUIDTiersUIDDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1AppProfilesUIDTiersUIDDeleteParamsWithHTTPClient(client *http.Client) *V1AppProfilesUIDTiersUIDDeleteParams {
	var ()
	return &V1AppProfilesUIDTiersUIDDeleteParams{
		HTTPClient: client,
	}
}

/*
V1AppProfilesUIDTiersUIDDeleteParams contains all the parameters to send to the API endpoint
for the v1 app profiles Uid tiers Uid delete operation typically these are written to a http.Request
*/
type V1AppProfilesUIDTiersUIDDeleteParams struct {

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

// WithTimeout adds the timeout to the v1 app profiles Uid tiers Uid delete params
func (o *V1AppProfilesUIDTiersUIDDeleteParams) WithTimeout(timeout time.Duration) *V1AppProfilesUIDTiersUIDDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 app profiles Uid tiers Uid delete params
func (o *V1AppProfilesUIDTiersUIDDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 app profiles Uid tiers Uid delete params
func (o *V1AppProfilesUIDTiersUIDDeleteParams) WithContext(ctx context.Context) *V1AppProfilesUIDTiersUIDDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 app profiles Uid tiers Uid delete params
func (o *V1AppProfilesUIDTiersUIDDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 app profiles Uid tiers Uid delete params
func (o *V1AppProfilesUIDTiersUIDDeleteParams) WithHTTPClient(client *http.Client) *V1AppProfilesUIDTiersUIDDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 app profiles Uid tiers Uid delete params
func (o *V1AppProfilesUIDTiersUIDDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTierUID adds the tierUID to the v1 app profiles Uid tiers Uid delete params
func (o *V1AppProfilesUIDTiersUIDDeleteParams) WithTierUID(tierUID string) *V1AppProfilesUIDTiersUIDDeleteParams {
	o.SetTierUID(tierUID)
	return o
}

// SetTierUID adds the tierUid to the v1 app profiles Uid tiers Uid delete params
func (o *V1AppProfilesUIDTiersUIDDeleteParams) SetTierUID(tierUID string) {
	o.TierUID = tierUID
}

// WithUID adds the uid to the v1 app profiles Uid tiers Uid delete params
func (o *V1AppProfilesUIDTiersUIDDeleteParams) WithUID(uid string) *V1AppProfilesUIDTiersUIDDeleteParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 app profiles Uid tiers Uid delete params
func (o *V1AppProfilesUIDTiersUIDDeleteParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1AppProfilesUIDTiersUIDDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
