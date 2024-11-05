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

// NewV1TenantClusterSettingsGetParams creates a new V1TenantClusterSettingsGetParams object
// with the default values initialized.
func NewV1TenantClusterSettingsGetParams() *V1TenantClusterSettingsGetParams {
	var ()
	return &V1TenantClusterSettingsGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1TenantClusterSettingsGetParamsWithTimeout creates a new V1TenantClusterSettingsGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1TenantClusterSettingsGetParamsWithTimeout(timeout time.Duration) *V1TenantClusterSettingsGetParams {
	var ()
	return &V1TenantClusterSettingsGetParams{

		timeout: timeout,
	}
}

// NewV1TenantClusterSettingsGetParamsWithContext creates a new V1TenantClusterSettingsGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1TenantClusterSettingsGetParamsWithContext(ctx context.Context) *V1TenantClusterSettingsGetParams {
	var ()
	return &V1TenantClusterSettingsGetParams{

		Context: ctx,
	}
}

// NewV1TenantClusterSettingsGetParamsWithHTTPClient creates a new V1TenantClusterSettingsGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1TenantClusterSettingsGetParamsWithHTTPClient(client *http.Client) *V1TenantClusterSettingsGetParams {
	var ()
	return &V1TenantClusterSettingsGetParams{
		HTTPClient: client,
	}
}

/*
V1TenantClusterSettingsGetParams contains all the parameters to send to the API endpoint
for the v1 tenant cluster settings get operation typically these are written to a http.Request
*/
type V1TenantClusterSettingsGetParams struct {

	/*TenantUID*/
	TenantUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 tenant cluster settings get params
func (o *V1TenantClusterSettingsGetParams) WithTimeout(timeout time.Duration) *V1TenantClusterSettingsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 tenant cluster settings get params
func (o *V1TenantClusterSettingsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 tenant cluster settings get params
func (o *V1TenantClusterSettingsGetParams) WithContext(ctx context.Context) *V1TenantClusterSettingsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 tenant cluster settings get params
func (o *V1TenantClusterSettingsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 tenant cluster settings get params
func (o *V1TenantClusterSettingsGetParams) WithHTTPClient(client *http.Client) *V1TenantClusterSettingsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 tenant cluster settings get params
func (o *V1TenantClusterSettingsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTenantUID adds the tenantUID to the v1 tenant cluster settings get params
func (o *V1TenantClusterSettingsGetParams) WithTenantUID(tenantUID string) *V1TenantClusterSettingsGetParams {
	o.SetTenantUID(tenantUID)
	return o
}

// SetTenantUID adds the tenantUid to the v1 tenant cluster settings get params
func (o *V1TenantClusterSettingsGetParams) SetTenantUID(tenantUID string) {
	o.TenantUID = tenantUID
}

// WriteToRequest writes these params to a swagger request
func (o *V1TenantClusterSettingsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
