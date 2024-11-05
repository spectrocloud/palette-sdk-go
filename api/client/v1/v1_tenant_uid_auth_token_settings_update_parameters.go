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

// NewV1TenantUIDAuthTokenSettingsUpdateParams creates a new V1TenantUIDAuthTokenSettingsUpdateParams object
// with the default values initialized.
func NewV1TenantUIDAuthTokenSettingsUpdateParams() *V1TenantUIDAuthTokenSettingsUpdateParams {
	var ()
	return &V1TenantUIDAuthTokenSettingsUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1TenantUIDAuthTokenSettingsUpdateParamsWithTimeout creates a new V1TenantUIDAuthTokenSettingsUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1TenantUIDAuthTokenSettingsUpdateParamsWithTimeout(timeout time.Duration) *V1TenantUIDAuthTokenSettingsUpdateParams {
	var ()
	return &V1TenantUIDAuthTokenSettingsUpdateParams{

		timeout: timeout,
	}
}

// NewV1TenantUIDAuthTokenSettingsUpdateParamsWithContext creates a new V1TenantUIDAuthTokenSettingsUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1TenantUIDAuthTokenSettingsUpdateParamsWithContext(ctx context.Context) *V1TenantUIDAuthTokenSettingsUpdateParams {
	var ()
	return &V1TenantUIDAuthTokenSettingsUpdateParams{

		Context: ctx,
	}
}

// NewV1TenantUIDAuthTokenSettingsUpdateParamsWithHTTPClient creates a new V1TenantUIDAuthTokenSettingsUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1TenantUIDAuthTokenSettingsUpdateParamsWithHTTPClient(client *http.Client) *V1TenantUIDAuthTokenSettingsUpdateParams {
	var ()
	return &V1TenantUIDAuthTokenSettingsUpdateParams{
		HTTPClient: client,
	}
}

/*
V1TenantUIDAuthTokenSettingsUpdateParams contains all the parameters to send to the API endpoint
for the v1 tenant Uid auth token settings update operation typically these are written to a http.Request
*/
type V1TenantUIDAuthTokenSettingsUpdateParams struct {

	/*Body*/
	Body *models.V1AuthTokenSettings
	/*TenantUID*/
	TenantUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 tenant Uid auth token settings update params
func (o *V1TenantUIDAuthTokenSettingsUpdateParams) WithTimeout(timeout time.Duration) *V1TenantUIDAuthTokenSettingsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 tenant Uid auth token settings update params
func (o *V1TenantUIDAuthTokenSettingsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 tenant Uid auth token settings update params
func (o *V1TenantUIDAuthTokenSettingsUpdateParams) WithContext(ctx context.Context) *V1TenantUIDAuthTokenSettingsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 tenant Uid auth token settings update params
func (o *V1TenantUIDAuthTokenSettingsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 tenant Uid auth token settings update params
func (o *V1TenantUIDAuthTokenSettingsUpdateParams) WithHTTPClient(client *http.Client) *V1TenantUIDAuthTokenSettingsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 tenant Uid auth token settings update params
func (o *V1TenantUIDAuthTokenSettingsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 tenant Uid auth token settings update params
func (o *V1TenantUIDAuthTokenSettingsUpdateParams) WithBody(body *models.V1AuthTokenSettings) *V1TenantUIDAuthTokenSettingsUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 tenant Uid auth token settings update params
func (o *V1TenantUIDAuthTokenSettingsUpdateParams) SetBody(body *models.V1AuthTokenSettings) {
	o.Body = body
}

// WithTenantUID adds the tenantUID to the v1 tenant Uid auth token settings update params
func (o *V1TenantUIDAuthTokenSettingsUpdateParams) WithTenantUID(tenantUID string) *V1TenantUIDAuthTokenSettingsUpdateParams {
	o.SetTenantUID(tenantUID)
	return o
}

// SetTenantUID adds the tenantUid to the v1 tenant Uid auth token settings update params
func (o *V1TenantUIDAuthTokenSettingsUpdateParams) SetTenantUID(tenantUID string) {
	o.TenantUID = tenantUID
}

// WriteToRequest writes these params to a swagger request
func (o *V1TenantUIDAuthTokenSettingsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
