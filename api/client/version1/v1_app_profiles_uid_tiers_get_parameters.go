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

// NewV1AppProfilesUIDTiersGetParams creates a new V1AppProfilesUIDTiersGetParams object
// with the default values initialized.
func NewV1AppProfilesUIDTiersGetParams() *V1AppProfilesUIDTiersGetParams {
	var ()
	return &V1AppProfilesUIDTiersGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1AppProfilesUIDTiersGetParamsWithTimeout creates a new V1AppProfilesUIDTiersGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1AppProfilesUIDTiersGetParamsWithTimeout(timeout time.Duration) *V1AppProfilesUIDTiersGetParams {
	var ()
	return &V1AppProfilesUIDTiersGetParams{

		timeout: timeout,
	}
}

// NewV1AppProfilesUIDTiersGetParamsWithContext creates a new V1AppProfilesUIDTiersGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1AppProfilesUIDTiersGetParamsWithContext(ctx context.Context) *V1AppProfilesUIDTiersGetParams {
	var ()
	return &V1AppProfilesUIDTiersGetParams{

		Context: ctx,
	}
}

// NewV1AppProfilesUIDTiersGetParamsWithHTTPClient creates a new V1AppProfilesUIDTiersGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1AppProfilesUIDTiersGetParamsWithHTTPClient(client *http.Client) *V1AppProfilesUIDTiersGetParams {
	var ()
	return &V1AppProfilesUIDTiersGetParams{
		HTTPClient: client,
	}
}

/*
V1AppProfilesUIDTiersGetParams contains all the parameters to send to the API endpoint
for the v1 app profiles Uid tiers get operation typically these are written to a http.Request
*/
type V1AppProfilesUIDTiersGetParams struct {

	/*UID
	  Application profile uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 app profiles Uid tiers get params
func (o *V1AppProfilesUIDTiersGetParams) WithTimeout(timeout time.Duration) *V1AppProfilesUIDTiersGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 app profiles Uid tiers get params
func (o *V1AppProfilesUIDTiersGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 app profiles Uid tiers get params
func (o *V1AppProfilesUIDTiersGetParams) WithContext(ctx context.Context) *V1AppProfilesUIDTiersGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 app profiles Uid tiers get params
func (o *V1AppProfilesUIDTiersGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 app profiles Uid tiers get params
func (o *V1AppProfilesUIDTiersGetParams) WithHTTPClient(client *http.Client) *V1AppProfilesUIDTiersGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 app profiles Uid tiers get params
func (o *V1AppProfilesUIDTiersGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUID adds the uid to the v1 app profiles Uid tiers get params
func (o *V1AppProfilesUIDTiersGetParams) WithUID(uid string) *V1AppProfilesUIDTiersGetParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 app profiles Uid tiers get params
func (o *V1AppProfilesUIDTiersGetParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1AppProfilesUIDTiersGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param uid
	if err := r.SetPathParam("uid", o.UID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
