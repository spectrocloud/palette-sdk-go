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

// NewV1TenantUIDPasswordPolicyGetParams creates a new V1TenantUIDPasswordPolicyGetParams object
// with the default values initialized.
func NewV1TenantUIDPasswordPolicyGetParams() *V1TenantUIDPasswordPolicyGetParams {
	var ()
	return &V1TenantUIDPasswordPolicyGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1TenantUIDPasswordPolicyGetParamsWithTimeout creates a new V1TenantUIDPasswordPolicyGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1TenantUIDPasswordPolicyGetParamsWithTimeout(timeout time.Duration) *V1TenantUIDPasswordPolicyGetParams {
	var ()
	return &V1TenantUIDPasswordPolicyGetParams{

		timeout: timeout,
	}
}

// NewV1TenantUIDPasswordPolicyGetParamsWithContext creates a new V1TenantUIDPasswordPolicyGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1TenantUIDPasswordPolicyGetParamsWithContext(ctx context.Context) *V1TenantUIDPasswordPolicyGetParams {
	var ()
	return &V1TenantUIDPasswordPolicyGetParams{

		Context: ctx,
	}
}

// NewV1TenantUIDPasswordPolicyGetParamsWithHTTPClient creates a new V1TenantUIDPasswordPolicyGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1TenantUIDPasswordPolicyGetParamsWithHTTPClient(client *http.Client) *V1TenantUIDPasswordPolicyGetParams {
	var ()
	return &V1TenantUIDPasswordPolicyGetParams{
		HTTPClient: client,
	}
}

/*
V1TenantUIDPasswordPolicyGetParams contains all the parameters to send to the API endpoint
for the v1 tenant Uid password policy get operation typically these are written to a http.Request
*/
type V1TenantUIDPasswordPolicyGetParams struct {

	/*TenantUID*/
	TenantUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 tenant Uid password policy get params
func (o *V1TenantUIDPasswordPolicyGetParams) WithTimeout(timeout time.Duration) *V1TenantUIDPasswordPolicyGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 tenant Uid password policy get params
func (o *V1TenantUIDPasswordPolicyGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 tenant Uid password policy get params
func (o *V1TenantUIDPasswordPolicyGetParams) WithContext(ctx context.Context) *V1TenantUIDPasswordPolicyGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 tenant Uid password policy get params
func (o *V1TenantUIDPasswordPolicyGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 tenant Uid password policy get params
func (o *V1TenantUIDPasswordPolicyGetParams) WithHTTPClient(client *http.Client) *V1TenantUIDPasswordPolicyGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 tenant Uid password policy get params
func (o *V1TenantUIDPasswordPolicyGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTenantUID adds the tenantUID to the v1 tenant Uid password policy get params
func (o *V1TenantUIDPasswordPolicyGetParams) WithTenantUID(tenantUID string) *V1TenantUIDPasswordPolicyGetParams {
	o.SetTenantUID(tenantUID)
	return o
}

// SetTenantUID adds the tenantUid to the v1 tenant Uid password policy get params
func (o *V1TenantUIDPasswordPolicyGetParams) SetTenantUID(tenantUID string) {
	o.TenantUID = tenantUID
}

// WriteToRequest writes these params to a swagger request
func (o *V1TenantUIDPasswordPolicyGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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