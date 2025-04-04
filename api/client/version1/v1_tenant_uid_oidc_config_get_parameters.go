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

// NewV1TenantUIDOidcConfigGetParams creates a new V1TenantUIDOidcConfigGetParams object
// with the default values initialized.
func NewV1TenantUIDOidcConfigGetParams() *V1TenantUIDOidcConfigGetParams {
	var ()
	return &V1TenantUIDOidcConfigGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1TenantUIDOidcConfigGetParamsWithTimeout creates a new V1TenantUIDOidcConfigGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1TenantUIDOidcConfigGetParamsWithTimeout(timeout time.Duration) *V1TenantUIDOidcConfigGetParams {
	var ()
	return &V1TenantUIDOidcConfigGetParams{

		timeout: timeout,
	}
}

// NewV1TenantUIDOidcConfigGetParamsWithContext creates a new V1TenantUIDOidcConfigGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1TenantUIDOidcConfigGetParamsWithContext(ctx context.Context) *V1TenantUIDOidcConfigGetParams {
	var ()
	return &V1TenantUIDOidcConfigGetParams{

		Context: ctx,
	}
}

// NewV1TenantUIDOidcConfigGetParamsWithHTTPClient creates a new V1TenantUIDOidcConfigGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1TenantUIDOidcConfigGetParamsWithHTTPClient(client *http.Client) *V1TenantUIDOidcConfigGetParams {
	var ()
	return &V1TenantUIDOidcConfigGetParams{
		HTTPClient: client,
	}
}

/*
V1TenantUIDOidcConfigGetParams contains all the parameters to send to the API endpoint
for the v1 tenant Uid oidc config get operation typically these are written to a http.Request
*/
type V1TenantUIDOidcConfigGetParams struct {

	/*TenantUID*/
	TenantUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 tenant Uid oidc config get params
func (o *V1TenantUIDOidcConfigGetParams) WithTimeout(timeout time.Duration) *V1TenantUIDOidcConfigGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 tenant Uid oidc config get params
func (o *V1TenantUIDOidcConfigGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 tenant Uid oidc config get params
func (o *V1TenantUIDOidcConfigGetParams) WithContext(ctx context.Context) *V1TenantUIDOidcConfigGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 tenant Uid oidc config get params
func (o *V1TenantUIDOidcConfigGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 tenant Uid oidc config get params
func (o *V1TenantUIDOidcConfigGetParams) WithHTTPClient(client *http.Client) *V1TenantUIDOidcConfigGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 tenant Uid oidc config get params
func (o *V1TenantUIDOidcConfigGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTenantUID adds the tenantUID to the v1 tenant Uid oidc config get params
func (o *V1TenantUIDOidcConfigGetParams) WithTenantUID(tenantUID string) *V1TenantUIDOidcConfigGetParams {
	o.SetTenantUID(tenantUID)
	return o
}

// SetTenantUID adds the tenantUid to the v1 tenant Uid oidc config get params
func (o *V1TenantUIDOidcConfigGetParams) SetTenantUID(tenantUID string) {
	o.TenantUID = tenantUID
}

// WriteToRequest writes these params to a swagger request
func (o *V1TenantUIDOidcConfigGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
