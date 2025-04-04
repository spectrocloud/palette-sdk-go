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

// NewV1TenantPrefClusterGroupUpdateParams creates a new V1TenantPrefClusterGroupUpdateParams object
// with the default values initialized.
func NewV1TenantPrefClusterGroupUpdateParams() *V1TenantPrefClusterGroupUpdateParams {
	var ()
	return &V1TenantPrefClusterGroupUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1TenantPrefClusterGroupUpdateParamsWithTimeout creates a new V1TenantPrefClusterGroupUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1TenantPrefClusterGroupUpdateParamsWithTimeout(timeout time.Duration) *V1TenantPrefClusterGroupUpdateParams {
	var ()
	return &V1TenantPrefClusterGroupUpdateParams{

		timeout: timeout,
	}
}

// NewV1TenantPrefClusterGroupUpdateParamsWithContext creates a new V1TenantPrefClusterGroupUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1TenantPrefClusterGroupUpdateParamsWithContext(ctx context.Context) *V1TenantPrefClusterGroupUpdateParams {
	var ()
	return &V1TenantPrefClusterGroupUpdateParams{

		Context: ctx,
	}
}

// NewV1TenantPrefClusterGroupUpdateParamsWithHTTPClient creates a new V1TenantPrefClusterGroupUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1TenantPrefClusterGroupUpdateParamsWithHTTPClient(client *http.Client) *V1TenantPrefClusterGroupUpdateParams {
	var ()
	return &V1TenantPrefClusterGroupUpdateParams{
		HTTPClient: client,
	}
}

/*
V1TenantPrefClusterGroupUpdateParams contains all the parameters to send to the API endpoint
for the v1 tenant pref cluster group update operation typically these are written to a http.Request
*/
type V1TenantPrefClusterGroupUpdateParams struct {

	/*Body*/
	Body *models.V1TenantEnableClusterGroup
	/*TenantUID*/
	TenantUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 tenant pref cluster group update params
func (o *V1TenantPrefClusterGroupUpdateParams) WithTimeout(timeout time.Duration) *V1TenantPrefClusterGroupUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 tenant pref cluster group update params
func (o *V1TenantPrefClusterGroupUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 tenant pref cluster group update params
func (o *V1TenantPrefClusterGroupUpdateParams) WithContext(ctx context.Context) *V1TenantPrefClusterGroupUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 tenant pref cluster group update params
func (o *V1TenantPrefClusterGroupUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 tenant pref cluster group update params
func (o *V1TenantPrefClusterGroupUpdateParams) WithHTTPClient(client *http.Client) *V1TenantPrefClusterGroupUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 tenant pref cluster group update params
func (o *V1TenantPrefClusterGroupUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 tenant pref cluster group update params
func (o *V1TenantPrefClusterGroupUpdateParams) WithBody(body *models.V1TenantEnableClusterGroup) *V1TenantPrefClusterGroupUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 tenant pref cluster group update params
func (o *V1TenantPrefClusterGroupUpdateParams) SetBody(body *models.V1TenantEnableClusterGroup) {
	o.Body = body
}

// WithTenantUID adds the tenantUID to the v1 tenant pref cluster group update params
func (o *V1TenantPrefClusterGroupUpdateParams) WithTenantUID(tenantUID string) *V1TenantPrefClusterGroupUpdateParams {
	o.SetTenantUID(tenantUID)
	return o
}

// SetTenantUID adds the tenantUid to the v1 tenant pref cluster group update params
func (o *V1TenantPrefClusterGroupUpdateParams) SetTenantUID(tenantUID string) {
	o.TenantUID = tenantUID
}

// WriteToRequest writes these params to a swagger request
func (o *V1TenantPrefClusterGroupUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
