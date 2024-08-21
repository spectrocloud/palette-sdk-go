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

// NewV1TenantUIDAssetsDataSinksDeleteParams creates a new V1TenantUIDAssetsDataSinksDeleteParams object
// with the default values initialized.
func NewV1TenantUIDAssetsDataSinksDeleteParams() *V1TenantUIDAssetsDataSinksDeleteParams {
	var ()
	return &V1TenantUIDAssetsDataSinksDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1TenantUIDAssetsDataSinksDeleteParamsWithTimeout creates a new V1TenantUIDAssetsDataSinksDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1TenantUIDAssetsDataSinksDeleteParamsWithTimeout(timeout time.Duration) *V1TenantUIDAssetsDataSinksDeleteParams {
	var ()
	return &V1TenantUIDAssetsDataSinksDeleteParams{

		timeout: timeout,
	}
}

// NewV1TenantUIDAssetsDataSinksDeleteParamsWithContext creates a new V1TenantUIDAssetsDataSinksDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1TenantUIDAssetsDataSinksDeleteParamsWithContext(ctx context.Context) *V1TenantUIDAssetsDataSinksDeleteParams {
	var ()
	return &V1TenantUIDAssetsDataSinksDeleteParams{

		Context: ctx,
	}
}

// NewV1TenantUIDAssetsDataSinksDeleteParamsWithHTTPClient creates a new V1TenantUIDAssetsDataSinksDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1TenantUIDAssetsDataSinksDeleteParamsWithHTTPClient(client *http.Client) *V1TenantUIDAssetsDataSinksDeleteParams {
	var ()
	return &V1TenantUIDAssetsDataSinksDeleteParams{
		HTTPClient: client,
	}
}

/*
V1TenantUIDAssetsDataSinksDeleteParams contains all the parameters to send to the API endpoint
for the v1 tenant Uid assets data sinks delete operation typically these are written to a http.Request
*/
type V1TenantUIDAssetsDataSinksDeleteParams struct {

	/*TenantUID*/
	TenantUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 tenant Uid assets data sinks delete params
func (o *V1TenantUIDAssetsDataSinksDeleteParams) WithTimeout(timeout time.Duration) *V1TenantUIDAssetsDataSinksDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 tenant Uid assets data sinks delete params
func (o *V1TenantUIDAssetsDataSinksDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 tenant Uid assets data sinks delete params
func (o *V1TenantUIDAssetsDataSinksDeleteParams) WithContext(ctx context.Context) *V1TenantUIDAssetsDataSinksDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 tenant Uid assets data sinks delete params
func (o *V1TenantUIDAssetsDataSinksDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 tenant Uid assets data sinks delete params
func (o *V1TenantUIDAssetsDataSinksDeleteParams) WithHTTPClient(client *http.Client) *V1TenantUIDAssetsDataSinksDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 tenant Uid assets data sinks delete params
func (o *V1TenantUIDAssetsDataSinksDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTenantUID adds the tenantUID to the v1 tenant Uid assets data sinks delete params
func (o *V1TenantUIDAssetsDataSinksDeleteParams) WithTenantUID(tenantUID string) *V1TenantUIDAssetsDataSinksDeleteParams {
	o.SetTenantUID(tenantUID)
	return o
}

// SetTenantUID adds the tenantUid to the v1 tenant Uid assets data sinks delete params
func (o *V1TenantUIDAssetsDataSinksDeleteParams) SetTenantUID(tenantUID string) {
	o.TenantUID = tenantUID
}

// WriteToRequest writes these params to a swagger request
func (o *V1TenantUIDAssetsDataSinksDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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