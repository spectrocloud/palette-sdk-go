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

// NewV1TenantUIDLoginBannerGetParams creates a new V1TenantUIDLoginBannerGetParams object
// with the default values initialized.
func NewV1TenantUIDLoginBannerGetParams() *V1TenantUIDLoginBannerGetParams {
	var ()
	return &V1TenantUIDLoginBannerGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1TenantUIDLoginBannerGetParamsWithTimeout creates a new V1TenantUIDLoginBannerGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1TenantUIDLoginBannerGetParamsWithTimeout(timeout time.Duration) *V1TenantUIDLoginBannerGetParams {
	var ()
	return &V1TenantUIDLoginBannerGetParams{

		timeout: timeout,
	}
}

// NewV1TenantUIDLoginBannerGetParamsWithContext creates a new V1TenantUIDLoginBannerGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1TenantUIDLoginBannerGetParamsWithContext(ctx context.Context) *V1TenantUIDLoginBannerGetParams {
	var ()
	return &V1TenantUIDLoginBannerGetParams{

		Context: ctx,
	}
}

// NewV1TenantUIDLoginBannerGetParamsWithHTTPClient creates a new V1TenantUIDLoginBannerGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1TenantUIDLoginBannerGetParamsWithHTTPClient(client *http.Client) *V1TenantUIDLoginBannerGetParams {
	var ()
	return &V1TenantUIDLoginBannerGetParams{
		HTTPClient: client,
	}
}

/*
V1TenantUIDLoginBannerGetParams contains all the parameters to send to the API endpoint
for the v1 tenant Uid login banner get operation typically these are written to a http.Request
*/
type V1TenantUIDLoginBannerGetParams struct {

	/*TenantUID*/
	TenantUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 tenant Uid login banner get params
func (o *V1TenantUIDLoginBannerGetParams) WithTimeout(timeout time.Duration) *V1TenantUIDLoginBannerGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 tenant Uid login banner get params
func (o *V1TenantUIDLoginBannerGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 tenant Uid login banner get params
func (o *V1TenantUIDLoginBannerGetParams) WithContext(ctx context.Context) *V1TenantUIDLoginBannerGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 tenant Uid login banner get params
func (o *V1TenantUIDLoginBannerGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 tenant Uid login banner get params
func (o *V1TenantUIDLoginBannerGetParams) WithHTTPClient(client *http.Client) *V1TenantUIDLoginBannerGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 tenant Uid login banner get params
func (o *V1TenantUIDLoginBannerGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTenantUID adds the tenantUID to the v1 tenant Uid login banner get params
func (o *V1TenantUIDLoginBannerGetParams) WithTenantUID(tenantUID string) *V1TenantUIDLoginBannerGetParams {
	o.SetTenantUID(tenantUID)
	return o
}

// SetTenantUID adds the tenantUid to the v1 tenant Uid login banner get params
func (o *V1TenantUIDLoginBannerGetParams) SetTenantUID(tenantUID string) {
	o.TenantUID = tenantUID
}

// WriteToRequest writes these params to a swagger request
func (o *V1TenantUIDLoginBannerGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
