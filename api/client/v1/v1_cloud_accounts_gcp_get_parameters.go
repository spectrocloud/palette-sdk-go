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

// NewV1CloudAccountsGcpGetParams creates a new V1CloudAccountsGcpGetParams object
// with the default values initialized.
func NewV1CloudAccountsGcpGetParams() *V1CloudAccountsGcpGetParams {
	var ()
	return &V1CloudAccountsGcpGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudAccountsGcpGetParamsWithTimeout creates a new V1CloudAccountsGcpGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudAccountsGcpGetParamsWithTimeout(timeout time.Duration) *V1CloudAccountsGcpGetParams {
	var ()
	return &V1CloudAccountsGcpGetParams{

		timeout: timeout,
	}
}

// NewV1CloudAccountsGcpGetParamsWithContext creates a new V1CloudAccountsGcpGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudAccountsGcpGetParamsWithContext(ctx context.Context) *V1CloudAccountsGcpGetParams {
	var ()
	return &V1CloudAccountsGcpGetParams{

		Context: ctx,
	}
}

// NewV1CloudAccountsGcpGetParamsWithHTTPClient creates a new V1CloudAccountsGcpGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudAccountsGcpGetParamsWithHTTPClient(client *http.Client) *V1CloudAccountsGcpGetParams {
	var ()
	return &V1CloudAccountsGcpGetParams{
		HTTPClient: client,
	}
}

/*
V1CloudAccountsGcpGetParams contains all the parameters to send to the API endpoint
for the v1 cloud accounts gcp get operation typically these are written to a http.Request
*/
type V1CloudAccountsGcpGetParams struct {

	/*UID
	  GCP cloud account uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cloud accounts gcp get params
func (o *V1CloudAccountsGcpGetParams) WithTimeout(timeout time.Duration) *V1CloudAccountsGcpGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud accounts gcp get params
func (o *V1CloudAccountsGcpGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud accounts gcp get params
func (o *V1CloudAccountsGcpGetParams) WithContext(ctx context.Context) *V1CloudAccountsGcpGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud accounts gcp get params
func (o *V1CloudAccountsGcpGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud accounts gcp get params
func (o *V1CloudAccountsGcpGetParams) WithHTTPClient(client *http.Client) *V1CloudAccountsGcpGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud accounts gcp get params
func (o *V1CloudAccountsGcpGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUID adds the uid to the v1 cloud accounts gcp get params
func (o *V1CloudAccountsGcpGetParams) WithUID(uid string) *V1CloudAccountsGcpGetParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 cloud accounts gcp get params
func (o *V1CloudAccountsGcpGetParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudAccountsGcpGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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