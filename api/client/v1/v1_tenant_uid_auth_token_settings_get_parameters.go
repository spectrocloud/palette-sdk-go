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

// NewV1TenantUIDAuthTokenSettingsGetParams creates a new V1TenantUIDAuthTokenSettingsGetParams object
// with the default values initialized.
func NewV1TenantUIDAuthTokenSettingsGetParams() *V1TenantUIDAuthTokenSettingsGetParams {
	var ()
	return &V1TenantUIDAuthTokenSettingsGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1TenantUIDAuthTokenSettingsGetParamsWithTimeout creates a new V1TenantUIDAuthTokenSettingsGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1TenantUIDAuthTokenSettingsGetParamsWithTimeout(timeout time.Duration) *V1TenantUIDAuthTokenSettingsGetParams {
	var ()
	return &V1TenantUIDAuthTokenSettingsGetParams{

		timeout: timeout,
	}
}

// NewV1TenantUIDAuthTokenSettingsGetParamsWithContext creates a new V1TenantUIDAuthTokenSettingsGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1TenantUIDAuthTokenSettingsGetParamsWithContext(ctx context.Context) *V1TenantUIDAuthTokenSettingsGetParams {
	var ()
	return &V1TenantUIDAuthTokenSettingsGetParams{

		Context: ctx,
	}
}

// NewV1TenantUIDAuthTokenSettingsGetParamsWithHTTPClient creates a new V1TenantUIDAuthTokenSettingsGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1TenantUIDAuthTokenSettingsGetParamsWithHTTPClient(client *http.Client) *V1TenantUIDAuthTokenSettingsGetParams {
	var ()
	return &V1TenantUIDAuthTokenSettingsGetParams{
		HTTPClient: client,
	}
}

/*V1TenantUIDAuthTokenSettingsGetParams contains all the parameters to send to the API endpoint
for the v1 tenant Uid auth token settings get operation typically these are written to a http.Request
*/
type V1TenantUIDAuthTokenSettingsGetParams struct {

	/*TenantUID*/
	TenantUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 tenant Uid auth token settings get params
func (o *V1TenantUIDAuthTokenSettingsGetParams) WithTimeout(timeout time.Duration) *V1TenantUIDAuthTokenSettingsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 tenant Uid auth token settings get params
func (o *V1TenantUIDAuthTokenSettingsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 tenant Uid auth token settings get params
func (o *V1TenantUIDAuthTokenSettingsGetParams) WithContext(ctx context.Context) *V1TenantUIDAuthTokenSettingsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 tenant Uid auth token settings get params
func (o *V1TenantUIDAuthTokenSettingsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 tenant Uid auth token settings get params
func (o *V1TenantUIDAuthTokenSettingsGetParams) WithHTTPClient(client *http.Client) *V1TenantUIDAuthTokenSettingsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 tenant Uid auth token settings get params
func (o *V1TenantUIDAuthTokenSettingsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTenantUID adds the tenantUID to the v1 tenant Uid auth token settings get params
func (o *V1TenantUIDAuthTokenSettingsGetParams) WithTenantUID(tenantUID string) *V1TenantUIDAuthTokenSettingsGetParams {
	o.SetTenantUID(tenantUID)
	return o
}

// SetTenantUID adds the tenantUid to the v1 tenant Uid auth token settings get params
func (o *V1TenantUIDAuthTokenSettingsGetParams) SetTenantUID(tenantUID string) {
	o.TenantUID = tenantUID
}

// WriteToRequest writes these params to a swagger request
func (o *V1TenantUIDAuthTokenSettingsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
