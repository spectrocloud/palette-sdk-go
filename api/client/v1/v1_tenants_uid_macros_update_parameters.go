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

	"github.com/spectrocloud/palette-sdk-go/api/models"
)

// NewV1TenantsUIDMacrosUpdateParams creates a new V1TenantsUIDMacrosUpdateParams object
// with the default values initialized.
func NewV1TenantsUIDMacrosUpdateParams() *V1TenantsUIDMacrosUpdateParams {
	var ()
	return &V1TenantsUIDMacrosUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1TenantsUIDMacrosUpdateParamsWithTimeout creates a new V1TenantsUIDMacrosUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1TenantsUIDMacrosUpdateParamsWithTimeout(timeout time.Duration) *V1TenantsUIDMacrosUpdateParams {
	var ()
	return &V1TenantsUIDMacrosUpdateParams{

		timeout: timeout,
	}
}

// NewV1TenantsUIDMacrosUpdateParamsWithContext creates a new V1TenantsUIDMacrosUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1TenantsUIDMacrosUpdateParamsWithContext(ctx context.Context) *V1TenantsUIDMacrosUpdateParams {
	var ()
	return &V1TenantsUIDMacrosUpdateParams{

		Context: ctx,
	}
}

// NewV1TenantsUIDMacrosUpdateParamsWithHTTPClient creates a new V1TenantsUIDMacrosUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1TenantsUIDMacrosUpdateParamsWithHTTPClient(client *http.Client) *V1TenantsUIDMacrosUpdateParams {
	var ()
	return &V1TenantsUIDMacrosUpdateParams{
		HTTPClient: client,
	}
}

/*
V1TenantsUIDMacrosUpdateParams contains all the parameters to send to the API endpoint
for the v1 tenants Uid macros update operation typically these are written to a http.Request
*/
type V1TenantsUIDMacrosUpdateParams struct {

	/*Body*/
	Body *models.V1Macros
	/*TenantUID*/
	TenantUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 tenants Uid macros update params
func (o *V1TenantsUIDMacrosUpdateParams) WithTimeout(timeout time.Duration) *V1TenantsUIDMacrosUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 tenants Uid macros update params
func (o *V1TenantsUIDMacrosUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 tenants Uid macros update params
func (o *V1TenantsUIDMacrosUpdateParams) WithContext(ctx context.Context) *V1TenantsUIDMacrosUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 tenants Uid macros update params
func (o *V1TenantsUIDMacrosUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 tenants Uid macros update params
func (o *V1TenantsUIDMacrosUpdateParams) WithHTTPClient(client *http.Client) *V1TenantsUIDMacrosUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 tenants Uid macros update params
func (o *V1TenantsUIDMacrosUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 tenants Uid macros update params
func (o *V1TenantsUIDMacrosUpdateParams) WithBody(body *models.V1Macros) *V1TenantsUIDMacrosUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 tenants Uid macros update params
func (o *V1TenantsUIDMacrosUpdateParams) SetBody(body *models.V1Macros) {
	o.Body = body
}

// WithTenantUID adds the tenantUID to the v1 tenants Uid macros update params
func (o *V1TenantsUIDMacrosUpdateParams) WithTenantUID(tenantUID string) *V1TenantsUIDMacrosUpdateParams {
	o.SetTenantUID(tenantUID)
	return o
}

// SetTenantUID adds the tenantUid to the v1 tenants Uid macros update params
func (o *V1TenantsUIDMacrosUpdateParams) SetTenantUID(tenantUID string) {
	o.TenantUID = tenantUID
}

// WriteToRequest writes these params to a swagger request
func (o *V1TenantsUIDMacrosUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param tenantUid
	if err := r.SetPathParam("tenantUid", o.TenantUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}