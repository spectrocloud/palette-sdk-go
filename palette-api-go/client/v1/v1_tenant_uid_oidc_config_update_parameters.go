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

	"github.com/spectrocloud/palette-api-go/models"
)

// NewV1TenantUIDOidcConfigUpdateParams creates a new V1TenantUIDOidcConfigUpdateParams object
// with the default values initialized.
func NewV1TenantUIDOidcConfigUpdateParams() *V1TenantUIDOidcConfigUpdateParams {
	var ()
	return &V1TenantUIDOidcConfigUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1TenantUIDOidcConfigUpdateParamsWithTimeout creates a new V1TenantUIDOidcConfigUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1TenantUIDOidcConfigUpdateParamsWithTimeout(timeout time.Duration) *V1TenantUIDOidcConfigUpdateParams {
	var ()
	return &V1TenantUIDOidcConfigUpdateParams{

		timeout: timeout,
	}
}

// NewV1TenantUIDOidcConfigUpdateParamsWithContext creates a new V1TenantUIDOidcConfigUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1TenantUIDOidcConfigUpdateParamsWithContext(ctx context.Context) *V1TenantUIDOidcConfigUpdateParams {
	var ()
	return &V1TenantUIDOidcConfigUpdateParams{

		Context: ctx,
	}
}

// NewV1TenantUIDOidcConfigUpdateParamsWithHTTPClient creates a new V1TenantUIDOidcConfigUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1TenantUIDOidcConfigUpdateParamsWithHTTPClient(client *http.Client) *V1TenantUIDOidcConfigUpdateParams {
	var ()
	return &V1TenantUIDOidcConfigUpdateParams{
		HTTPClient: client,
	}
}

/*V1TenantUIDOidcConfigUpdateParams contains all the parameters to send to the API endpoint
for the v1 tenant Uid oidc config update operation typically these are written to a http.Request
*/
type V1TenantUIDOidcConfigUpdateParams struct {

	/*Body*/
	Body *models.V1TenantOidcClientSpec
	/*TenantUID*/
	TenantUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 tenant Uid oidc config update params
func (o *V1TenantUIDOidcConfigUpdateParams) WithTimeout(timeout time.Duration) *V1TenantUIDOidcConfigUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 tenant Uid oidc config update params
func (o *V1TenantUIDOidcConfigUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 tenant Uid oidc config update params
func (o *V1TenantUIDOidcConfigUpdateParams) WithContext(ctx context.Context) *V1TenantUIDOidcConfigUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 tenant Uid oidc config update params
func (o *V1TenantUIDOidcConfigUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 tenant Uid oidc config update params
func (o *V1TenantUIDOidcConfigUpdateParams) WithHTTPClient(client *http.Client) *V1TenantUIDOidcConfigUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 tenant Uid oidc config update params
func (o *V1TenantUIDOidcConfigUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 tenant Uid oidc config update params
func (o *V1TenantUIDOidcConfigUpdateParams) WithBody(body *models.V1TenantOidcClientSpec) *V1TenantUIDOidcConfigUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 tenant Uid oidc config update params
func (o *V1TenantUIDOidcConfigUpdateParams) SetBody(body *models.V1TenantOidcClientSpec) {
	o.Body = body
}

// WithTenantUID adds the tenantUID to the v1 tenant Uid oidc config update params
func (o *V1TenantUIDOidcConfigUpdateParams) WithTenantUID(tenantUID string) *V1TenantUIDOidcConfigUpdateParams {
	o.SetTenantUID(tenantUID)
	return o
}

// SetTenantUID adds the tenantUid to the v1 tenant Uid oidc config update params
func (o *V1TenantUIDOidcConfigUpdateParams) SetTenantUID(tenantUID string) {
	o.TenantUID = tenantUID
}

// WriteToRequest writes these params to a swagger request
func (o *V1TenantUIDOidcConfigUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
