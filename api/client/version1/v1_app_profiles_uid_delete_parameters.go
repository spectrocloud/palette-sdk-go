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

// NewV1AppProfilesUIDDeleteParams creates a new V1AppProfilesUIDDeleteParams object
// with the default values initialized.
func NewV1AppProfilesUIDDeleteParams() *V1AppProfilesUIDDeleteParams {
	var ()
	return &V1AppProfilesUIDDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1AppProfilesUIDDeleteParamsWithTimeout creates a new V1AppProfilesUIDDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1AppProfilesUIDDeleteParamsWithTimeout(timeout time.Duration) *V1AppProfilesUIDDeleteParams {
	var ()
	return &V1AppProfilesUIDDeleteParams{

		timeout: timeout,
	}
}

// NewV1AppProfilesUIDDeleteParamsWithContext creates a new V1AppProfilesUIDDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1AppProfilesUIDDeleteParamsWithContext(ctx context.Context) *V1AppProfilesUIDDeleteParams {
	var ()
	return &V1AppProfilesUIDDeleteParams{

		Context: ctx,
	}
}

// NewV1AppProfilesUIDDeleteParamsWithHTTPClient creates a new V1AppProfilesUIDDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1AppProfilesUIDDeleteParamsWithHTTPClient(client *http.Client) *V1AppProfilesUIDDeleteParams {
	var ()
	return &V1AppProfilesUIDDeleteParams{
		HTTPClient: client,
	}
}

/*
V1AppProfilesUIDDeleteParams contains all the parameters to send to the API endpoint
for the v1 app profiles Uid delete operation typically these are written to a http.Request
*/
type V1AppProfilesUIDDeleteParams struct {

	/*UID*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 app profiles Uid delete params
func (o *V1AppProfilesUIDDeleteParams) WithTimeout(timeout time.Duration) *V1AppProfilesUIDDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 app profiles Uid delete params
func (o *V1AppProfilesUIDDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 app profiles Uid delete params
func (o *V1AppProfilesUIDDeleteParams) WithContext(ctx context.Context) *V1AppProfilesUIDDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 app profiles Uid delete params
func (o *V1AppProfilesUIDDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 app profiles Uid delete params
func (o *V1AppProfilesUIDDeleteParams) WithHTTPClient(client *http.Client) *V1AppProfilesUIDDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 app profiles Uid delete params
func (o *V1AppProfilesUIDDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUID adds the uid to the v1 app profiles Uid delete params
func (o *V1AppProfilesUIDDeleteParams) WithUID(uid string) *V1AppProfilesUIDDeleteParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 app profiles Uid delete params
func (o *V1AppProfilesUIDDeleteParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1AppProfilesUIDDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
