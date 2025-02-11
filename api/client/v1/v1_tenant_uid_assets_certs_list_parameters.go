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

// NewV1TenantUIDAssetsCertsListParams creates a new V1TenantUIDAssetsCertsListParams object
// with the default values initialized.
func NewV1TenantUIDAssetsCertsListParams() *V1TenantUIDAssetsCertsListParams {
	var ()
	return &V1TenantUIDAssetsCertsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1TenantUIDAssetsCertsListParamsWithTimeout creates a new V1TenantUIDAssetsCertsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1TenantUIDAssetsCertsListParamsWithTimeout(timeout time.Duration) *V1TenantUIDAssetsCertsListParams {
	var ()
	return &V1TenantUIDAssetsCertsListParams{

		timeout: timeout,
	}
}

// NewV1TenantUIDAssetsCertsListParamsWithContext creates a new V1TenantUIDAssetsCertsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1TenantUIDAssetsCertsListParamsWithContext(ctx context.Context) *V1TenantUIDAssetsCertsListParams {
	var ()
	return &V1TenantUIDAssetsCertsListParams{

		Context: ctx,
	}
}

// NewV1TenantUIDAssetsCertsListParamsWithHTTPClient creates a new V1TenantUIDAssetsCertsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1TenantUIDAssetsCertsListParamsWithHTTPClient(client *http.Client) *V1TenantUIDAssetsCertsListParams {
	var ()
	return &V1TenantUIDAssetsCertsListParams{
		HTTPClient: client,
	}
}

/*
V1TenantUIDAssetsCertsListParams contains all the parameters to send to the API endpoint
for the v1 tenant u Id assets certs list operation typically these are written to a http.Request
*/
type V1TenantUIDAssetsCertsListParams struct {

	/*TenantUID*/
	TenantUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 tenant u Id assets certs list params
func (o *V1TenantUIDAssetsCertsListParams) WithTimeout(timeout time.Duration) *V1TenantUIDAssetsCertsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 tenant u Id assets certs list params
func (o *V1TenantUIDAssetsCertsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 tenant u Id assets certs list params
func (o *V1TenantUIDAssetsCertsListParams) WithContext(ctx context.Context) *V1TenantUIDAssetsCertsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 tenant u Id assets certs list params
func (o *V1TenantUIDAssetsCertsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 tenant u Id assets certs list params
func (o *V1TenantUIDAssetsCertsListParams) WithHTTPClient(client *http.Client) *V1TenantUIDAssetsCertsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 tenant u Id assets certs list params
func (o *V1TenantUIDAssetsCertsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTenantUID adds the tenantUID to the v1 tenant u Id assets certs list params
func (o *V1TenantUIDAssetsCertsListParams) WithTenantUID(tenantUID string) *V1TenantUIDAssetsCertsListParams {
	o.SetTenantUID(tenantUID)
	return o
}

// SetTenantUID adds the tenantUid to the v1 tenant u Id assets certs list params
func (o *V1TenantUIDAssetsCertsListParams) SetTenantUID(tenantUID string) {
	o.TenantUID = tenantUID
}

// WriteToRequest writes these params to a swagger request
func (o *V1TenantUIDAssetsCertsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
