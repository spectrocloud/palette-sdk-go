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

// NewV1CloudAccountsGcpDeleteParams creates a new V1CloudAccountsGcpDeleteParams object
// with the default values initialized.
func NewV1CloudAccountsGcpDeleteParams() *V1CloudAccountsGcpDeleteParams {
	var ()
	return &V1CloudAccountsGcpDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudAccountsGcpDeleteParamsWithTimeout creates a new V1CloudAccountsGcpDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudAccountsGcpDeleteParamsWithTimeout(timeout time.Duration) *V1CloudAccountsGcpDeleteParams {
	var ()
	return &V1CloudAccountsGcpDeleteParams{

		timeout: timeout,
	}
}

// NewV1CloudAccountsGcpDeleteParamsWithContext creates a new V1CloudAccountsGcpDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudAccountsGcpDeleteParamsWithContext(ctx context.Context) *V1CloudAccountsGcpDeleteParams {
	var ()
	return &V1CloudAccountsGcpDeleteParams{

		Context: ctx,
	}
}

// NewV1CloudAccountsGcpDeleteParamsWithHTTPClient creates a new V1CloudAccountsGcpDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudAccountsGcpDeleteParamsWithHTTPClient(client *http.Client) *V1CloudAccountsGcpDeleteParams {
	var ()
	return &V1CloudAccountsGcpDeleteParams{
		HTTPClient: client,
	}
}

/*
V1CloudAccountsGcpDeleteParams contains all the parameters to send to the API endpoint
for the v1 cloud accounts gcp delete operation typically these are written to a http.Request
*/
type V1CloudAccountsGcpDeleteParams struct {

	/*UID
	  GCP cloud account uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cloud accounts gcp delete params
func (o *V1CloudAccountsGcpDeleteParams) WithTimeout(timeout time.Duration) *V1CloudAccountsGcpDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud accounts gcp delete params
func (o *V1CloudAccountsGcpDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud accounts gcp delete params
func (o *V1CloudAccountsGcpDeleteParams) WithContext(ctx context.Context) *V1CloudAccountsGcpDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud accounts gcp delete params
func (o *V1CloudAccountsGcpDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud accounts gcp delete params
func (o *V1CloudAccountsGcpDeleteParams) WithHTTPClient(client *http.Client) *V1CloudAccountsGcpDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud accounts gcp delete params
func (o *V1CloudAccountsGcpDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUID adds the uid to the v1 cloud accounts gcp delete params
func (o *V1CloudAccountsGcpDeleteParams) WithUID(uid string) *V1CloudAccountsGcpDeleteParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 cloud accounts gcp delete params
func (o *V1CloudAccountsGcpDeleteParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudAccountsGcpDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
