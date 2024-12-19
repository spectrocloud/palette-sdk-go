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

// NewV1TenantUIDDomainsUpdateParams creates a new V1TenantUIDDomainsUpdateParams object
// with the default values initialized.
func NewV1TenantUIDDomainsUpdateParams() *V1TenantUIDDomainsUpdateParams {
	var ()
	return &V1TenantUIDDomainsUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1TenantUIDDomainsUpdateParamsWithTimeout creates a new V1TenantUIDDomainsUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1TenantUIDDomainsUpdateParamsWithTimeout(timeout time.Duration) *V1TenantUIDDomainsUpdateParams {
	var ()
	return &V1TenantUIDDomainsUpdateParams{

		timeout: timeout,
	}
}

// NewV1TenantUIDDomainsUpdateParamsWithContext creates a new V1TenantUIDDomainsUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1TenantUIDDomainsUpdateParamsWithContext(ctx context.Context) *V1TenantUIDDomainsUpdateParams {
	var ()
	return &V1TenantUIDDomainsUpdateParams{

		Context: ctx,
	}
}

// NewV1TenantUIDDomainsUpdateParamsWithHTTPClient creates a new V1TenantUIDDomainsUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1TenantUIDDomainsUpdateParamsWithHTTPClient(client *http.Client) *V1TenantUIDDomainsUpdateParams {
	var ()
	return &V1TenantUIDDomainsUpdateParams{
		HTTPClient: client,
	}
}

/*V1TenantUIDDomainsUpdateParams contains all the parameters to send to the API endpoint
for the v1 tenant Uid domains update operation typically these are written to a http.Request
*/
type V1TenantUIDDomainsUpdateParams struct {

	/*Body*/
	Body *models.V1TenantDomains
	/*TenantUID*/
	TenantUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 tenant Uid domains update params
func (o *V1TenantUIDDomainsUpdateParams) WithTimeout(timeout time.Duration) *V1TenantUIDDomainsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 tenant Uid domains update params
func (o *V1TenantUIDDomainsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 tenant Uid domains update params
func (o *V1TenantUIDDomainsUpdateParams) WithContext(ctx context.Context) *V1TenantUIDDomainsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 tenant Uid domains update params
func (o *V1TenantUIDDomainsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 tenant Uid domains update params
func (o *V1TenantUIDDomainsUpdateParams) WithHTTPClient(client *http.Client) *V1TenantUIDDomainsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 tenant Uid domains update params
func (o *V1TenantUIDDomainsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 tenant Uid domains update params
func (o *V1TenantUIDDomainsUpdateParams) WithBody(body *models.V1TenantDomains) *V1TenantUIDDomainsUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 tenant Uid domains update params
func (o *V1TenantUIDDomainsUpdateParams) SetBody(body *models.V1TenantDomains) {
	o.Body = body
}

// WithTenantUID adds the tenantUID to the v1 tenant Uid domains update params
func (o *V1TenantUIDDomainsUpdateParams) WithTenantUID(tenantUID string) *V1TenantUIDDomainsUpdateParams {
	o.SetTenantUID(tenantUID)
	return o
}

// SetTenantUID adds the tenantUid to the v1 tenant Uid domains update params
func (o *V1TenantUIDDomainsUpdateParams) SetTenantUID(tenantUID string) {
	o.TenantUID = tenantUID
}

// WriteToRequest writes these params to a swagger request
func (o *V1TenantUIDDomainsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
