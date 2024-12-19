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

// NewV1CloudAccountsVsphereGetParams creates a new V1CloudAccountsVsphereGetParams object
// with the default values initialized.
func NewV1CloudAccountsVsphereGetParams() *V1CloudAccountsVsphereGetParams {
	var ()
	return &V1CloudAccountsVsphereGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudAccountsVsphereGetParamsWithTimeout creates a new V1CloudAccountsVsphereGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudAccountsVsphereGetParamsWithTimeout(timeout time.Duration) *V1CloudAccountsVsphereGetParams {
	var ()
	return &V1CloudAccountsVsphereGetParams{

		timeout: timeout,
	}
}

// NewV1CloudAccountsVsphereGetParamsWithContext creates a new V1CloudAccountsVsphereGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudAccountsVsphereGetParamsWithContext(ctx context.Context) *V1CloudAccountsVsphereGetParams {
	var ()
	return &V1CloudAccountsVsphereGetParams{

		Context: ctx,
	}
}

// NewV1CloudAccountsVsphereGetParamsWithHTTPClient creates a new V1CloudAccountsVsphereGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudAccountsVsphereGetParamsWithHTTPClient(client *http.Client) *V1CloudAccountsVsphereGetParams {
	var ()
	return &V1CloudAccountsVsphereGetParams{
		HTTPClient: client,
	}
}

/*
V1CloudAccountsVsphereGetParams contains all the parameters to send to the API endpoint
for the v1 cloud accounts vsphere get operation typically these are written to a http.Request
*/
type V1CloudAccountsVsphereGetParams struct {

	/*UID
	  VSphere cloud account uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cloud accounts vsphere get params
func (o *V1CloudAccountsVsphereGetParams) WithTimeout(timeout time.Duration) *V1CloudAccountsVsphereGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud accounts vsphere get params
func (o *V1CloudAccountsVsphereGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud accounts vsphere get params
func (o *V1CloudAccountsVsphereGetParams) WithContext(ctx context.Context) *V1CloudAccountsVsphereGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud accounts vsphere get params
func (o *V1CloudAccountsVsphereGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud accounts vsphere get params
func (o *V1CloudAccountsVsphereGetParams) WithHTTPClient(client *http.Client) *V1CloudAccountsVsphereGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud accounts vsphere get params
func (o *V1CloudAccountsVsphereGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUID adds the uid to the v1 cloud accounts vsphere get params
func (o *V1CloudAccountsVsphereGetParams) WithUID(uid string) *V1CloudAccountsVsphereGetParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 cloud accounts vsphere get params
func (o *V1CloudAccountsVsphereGetParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudAccountsVsphereGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
