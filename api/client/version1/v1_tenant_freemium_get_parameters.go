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

// NewV1TenantFreemiumGetParams creates a new V1TenantFreemiumGetParams object
// with the default values initialized.
func NewV1TenantFreemiumGetParams() *V1TenantFreemiumGetParams {
	var ()
	return &V1TenantFreemiumGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1TenantFreemiumGetParamsWithTimeout creates a new V1TenantFreemiumGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1TenantFreemiumGetParamsWithTimeout(timeout time.Duration) *V1TenantFreemiumGetParams {
	var ()
	return &V1TenantFreemiumGetParams{

		timeout: timeout,
	}
}

// NewV1TenantFreemiumGetParamsWithContext creates a new V1TenantFreemiumGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1TenantFreemiumGetParamsWithContext(ctx context.Context) *V1TenantFreemiumGetParams {
	var ()
	return &V1TenantFreemiumGetParams{

		Context: ctx,
	}
}

// NewV1TenantFreemiumGetParamsWithHTTPClient creates a new V1TenantFreemiumGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1TenantFreemiumGetParamsWithHTTPClient(client *http.Client) *V1TenantFreemiumGetParams {
	var ()
	return &V1TenantFreemiumGetParams{
		HTTPClient: client,
	}
}

/*
V1TenantFreemiumGetParams contains all the parameters to send to the API endpoint
for the v1 tenant freemium get operation typically these are written to a http.Request
*/
type V1TenantFreemiumGetParams struct {

	/*TenantUID*/
	TenantUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 tenant freemium get params
func (o *V1TenantFreemiumGetParams) WithTimeout(timeout time.Duration) *V1TenantFreemiumGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 tenant freemium get params
func (o *V1TenantFreemiumGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 tenant freemium get params
func (o *V1TenantFreemiumGetParams) WithContext(ctx context.Context) *V1TenantFreemiumGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 tenant freemium get params
func (o *V1TenantFreemiumGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 tenant freemium get params
func (o *V1TenantFreemiumGetParams) WithHTTPClient(client *http.Client) *V1TenantFreemiumGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 tenant freemium get params
func (o *V1TenantFreemiumGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTenantUID adds the tenantUID to the v1 tenant freemium get params
func (o *V1TenantFreemiumGetParams) WithTenantUID(tenantUID string) *V1TenantFreemiumGetParams {
	o.SetTenantUID(tenantUID)
	return o
}

// SetTenantUID adds the tenantUid to the v1 tenant freemium get params
func (o *V1TenantFreemiumGetParams) SetTenantUID(tenantUID string) {
	o.TenantUID = tenantUID
}

// WriteToRequest writes these params to a swagger request
func (o *V1TenantFreemiumGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param tenantUid
	if err := r.SetPathParam("tenantUid", o.TenantUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
