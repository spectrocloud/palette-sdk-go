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

// NewV1TenantUIDSsoAuthProvidersGetParams creates a new V1TenantUIDSsoAuthProvidersGetParams object
// with the default values initialized.
func NewV1TenantUIDSsoAuthProvidersGetParams() *V1TenantUIDSsoAuthProvidersGetParams {
	var ()
	return &V1TenantUIDSsoAuthProvidersGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1TenantUIDSsoAuthProvidersGetParamsWithTimeout creates a new V1TenantUIDSsoAuthProvidersGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1TenantUIDSsoAuthProvidersGetParamsWithTimeout(timeout time.Duration) *V1TenantUIDSsoAuthProvidersGetParams {
	var ()
	return &V1TenantUIDSsoAuthProvidersGetParams{

		timeout: timeout,
	}
}

// NewV1TenantUIDSsoAuthProvidersGetParamsWithContext creates a new V1TenantUIDSsoAuthProvidersGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1TenantUIDSsoAuthProvidersGetParamsWithContext(ctx context.Context) *V1TenantUIDSsoAuthProvidersGetParams {
	var ()
	return &V1TenantUIDSsoAuthProvidersGetParams{

		Context: ctx,
	}
}

// NewV1TenantUIDSsoAuthProvidersGetParamsWithHTTPClient creates a new V1TenantUIDSsoAuthProvidersGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1TenantUIDSsoAuthProvidersGetParamsWithHTTPClient(client *http.Client) *V1TenantUIDSsoAuthProvidersGetParams {
	var ()
	return &V1TenantUIDSsoAuthProvidersGetParams{
		HTTPClient: client,
	}
}

/*
V1TenantUIDSsoAuthProvidersGetParams contains all the parameters to send to the API endpoint
for the v1 tenant Uid sso auth providers get operation typically these are written to a http.Request
*/
type V1TenantUIDSsoAuthProvidersGetParams struct {

	/*TenantUID*/
	TenantUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 tenant Uid sso auth providers get params
func (o *V1TenantUIDSsoAuthProvidersGetParams) WithTimeout(timeout time.Duration) *V1TenantUIDSsoAuthProvidersGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 tenant Uid sso auth providers get params
func (o *V1TenantUIDSsoAuthProvidersGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 tenant Uid sso auth providers get params
func (o *V1TenantUIDSsoAuthProvidersGetParams) WithContext(ctx context.Context) *V1TenantUIDSsoAuthProvidersGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 tenant Uid sso auth providers get params
func (o *V1TenantUIDSsoAuthProvidersGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 tenant Uid sso auth providers get params
func (o *V1TenantUIDSsoAuthProvidersGetParams) WithHTTPClient(client *http.Client) *V1TenantUIDSsoAuthProvidersGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 tenant Uid sso auth providers get params
func (o *V1TenantUIDSsoAuthProvidersGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTenantUID adds the tenantUID to the v1 tenant Uid sso auth providers get params
func (o *V1TenantUIDSsoAuthProvidersGetParams) WithTenantUID(tenantUID string) *V1TenantUIDSsoAuthProvidersGetParams {
	o.SetTenantUID(tenantUID)
	return o
}

// SetTenantUID adds the tenantUid to the v1 tenant Uid sso auth providers get params
func (o *V1TenantUIDSsoAuthProvidersGetParams) SetTenantUID(tenantUID string) {
	o.TenantUID = tenantUID
}

// WriteToRequest writes these params to a swagger request
func (o *V1TenantUIDSsoAuthProvidersGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
