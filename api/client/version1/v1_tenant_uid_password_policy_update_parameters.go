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

	"github.com/spectrocloud/palette-sdk-go/api/models"
)

// NewV1TenantUIDPasswordPolicyUpdateParams creates a new V1TenantUIDPasswordPolicyUpdateParams object
// with the default values initialized.
func NewV1TenantUIDPasswordPolicyUpdateParams() *V1TenantUIDPasswordPolicyUpdateParams {
	var ()
	return &V1TenantUIDPasswordPolicyUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1TenantUIDPasswordPolicyUpdateParamsWithTimeout creates a new V1TenantUIDPasswordPolicyUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1TenantUIDPasswordPolicyUpdateParamsWithTimeout(timeout time.Duration) *V1TenantUIDPasswordPolicyUpdateParams {
	var ()
	return &V1TenantUIDPasswordPolicyUpdateParams{

		timeout: timeout,
	}
}

// NewV1TenantUIDPasswordPolicyUpdateParamsWithContext creates a new V1TenantUIDPasswordPolicyUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1TenantUIDPasswordPolicyUpdateParamsWithContext(ctx context.Context) *V1TenantUIDPasswordPolicyUpdateParams {
	var ()
	return &V1TenantUIDPasswordPolicyUpdateParams{

		Context: ctx,
	}
}

// NewV1TenantUIDPasswordPolicyUpdateParamsWithHTTPClient creates a new V1TenantUIDPasswordPolicyUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1TenantUIDPasswordPolicyUpdateParamsWithHTTPClient(client *http.Client) *V1TenantUIDPasswordPolicyUpdateParams {
	var ()
	return &V1TenantUIDPasswordPolicyUpdateParams{
		HTTPClient: client,
	}
}

/*
V1TenantUIDPasswordPolicyUpdateParams contains all the parameters to send to the API endpoint
for the v1 tenant Uid password policy update operation typically these are written to a http.Request
*/
type V1TenantUIDPasswordPolicyUpdateParams struct {

	/*Body*/
	Body *models.V1TenantPasswordPolicyEntity
	/*TenantUID*/
	TenantUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 tenant Uid password policy update params
func (o *V1TenantUIDPasswordPolicyUpdateParams) WithTimeout(timeout time.Duration) *V1TenantUIDPasswordPolicyUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 tenant Uid password policy update params
func (o *V1TenantUIDPasswordPolicyUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 tenant Uid password policy update params
func (o *V1TenantUIDPasswordPolicyUpdateParams) WithContext(ctx context.Context) *V1TenantUIDPasswordPolicyUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 tenant Uid password policy update params
func (o *V1TenantUIDPasswordPolicyUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 tenant Uid password policy update params
func (o *V1TenantUIDPasswordPolicyUpdateParams) WithHTTPClient(client *http.Client) *V1TenantUIDPasswordPolicyUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 tenant Uid password policy update params
func (o *V1TenantUIDPasswordPolicyUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 tenant Uid password policy update params
func (o *V1TenantUIDPasswordPolicyUpdateParams) WithBody(body *models.V1TenantPasswordPolicyEntity) *V1TenantUIDPasswordPolicyUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 tenant Uid password policy update params
func (o *V1TenantUIDPasswordPolicyUpdateParams) SetBody(body *models.V1TenantPasswordPolicyEntity) {
	o.Body = body
}

// WithTenantUID adds the tenantUID to the v1 tenant Uid password policy update params
func (o *V1TenantUIDPasswordPolicyUpdateParams) WithTenantUID(tenantUID string) *V1TenantUIDPasswordPolicyUpdateParams {
	o.SetTenantUID(tenantUID)
	return o
}

// SetTenantUID adds the tenantUid to the v1 tenant Uid password policy update params
func (o *V1TenantUIDPasswordPolicyUpdateParams) SetTenantUID(tenantUID string) {
	o.TenantUID = tenantUID
}

// WriteToRequest writes these params to a swagger request
func (o *V1TenantUIDPasswordPolicyUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
