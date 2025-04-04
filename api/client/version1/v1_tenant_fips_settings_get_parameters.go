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

// NewV1TenantFipsSettingsGetParams creates a new V1TenantFipsSettingsGetParams object
// with the default values initialized.
func NewV1TenantFipsSettingsGetParams() *V1TenantFipsSettingsGetParams {
	var ()
	return &V1TenantFipsSettingsGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1TenantFipsSettingsGetParamsWithTimeout creates a new V1TenantFipsSettingsGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1TenantFipsSettingsGetParamsWithTimeout(timeout time.Duration) *V1TenantFipsSettingsGetParams {
	var ()
	return &V1TenantFipsSettingsGetParams{

		timeout: timeout,
	}
}

// NewV1TenantFipsSettingsGetParamsWithContext creates a new V1TenantFipsSettingsGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1TenantFipsSettingsGetParamsWithContext(ctx context.Context) *V1TenantFipsSettingsGetParams {
	var ()
	return &V1TenantFipsSettingsGetParams{

		Context: ctx,
	}
}

// NewV1TenantFipsSettingsGetParamsWithHTTPClient creates a new V1TenantFipsSettingsGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1TenantFipsSettingsGetParamsWithHTTPClient(client *http.Client) *V1TenantFipsSettingsGetParams {
	var ()
	return &V1TenantFipsSettingsGetParams{
		HTTPClient: client,
	}
}

/*
V1TenantFipsSettingsGetParams contains all the parameters to send to the API endpoint
for the v1 tenant fips settings get operation typically these are written to a http.Request
*/
type V1TenantFipsSettingsGetParams struct {

	/*TenantUID*/
	TenantUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 tenant fips settings get params
func (o *V1TenantFipsSettingsGetParams) WithTimeout(timeout time.Duration) *V1TenantFipsSettingsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 tenant fips settings get params
func (o *V1TenantFipsSettingsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 tenant fips settings get params
func (o *V1TenantFipsSettingsGetParams) WithContext(ctx context.Context) *V1TenantFipsSettingsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 tenant fips settings get params
func (o *V1TenantFipsSettingsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 tenant fips settings get params
func (o *V1TenantFipsSettingsGetParams) WithHTTPClient(client *http.Client) *V1TenantFipsSettingsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 tenant fips settings get params
func (o *V1TenantFipsSettingsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTenantUID adds the tenantUID to the v1 tenant fips settings get params
func (o *V1TenantFipsSettingsGetParams) WithTenantUID(tenantUID string) *V1TenantFipsSettingsGetParams {
	o.SetTenantUID(tenantUID)
	return o
}

// SetTenantUID adds the tenantUid to the v1 tenant fips settings get params
func (o *V1TenantFipsSettingsGetParams) SetTenantUID(tenantUID string) {
	o.TenantUID = tenantUID
}

// WriteToRequest writes these params to a swagger request
func (o *V1TenantFipsSettingsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
