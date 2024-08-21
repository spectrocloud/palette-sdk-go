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

// NewV1TenantUIDLoginBannerUpdateParams creates a new V1TenantUIDLoginBannerUpdateParams object
// with the default values initialized.
func NewV1TenantUIDLoginBannerUpdateParams() *V1TenantUIDLoginBannerUpdateParams {
	var ()
	return &V1TenantUIDLoginBannerUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1TenantUIDLoginBannerUpdateParamsWithTimeout creates a new V1TenantUIDLoginBannerUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1TenantUIDLoginBannerUpdateParamsWithTimeout(timeout time.Duration) *V1TenantUIDLoginBannerUpdateParams {
	var ()
	return &V1TenantUIDLoginBannerUpdateParams{

		timeout: timeout,
	}
}

// NewV1TenantUIDLoginBannerUpdateParamsWithContext creates a new V1TenantUIDLoginBannerUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1TenantUIDLoginBannerUpdateParamsWithContext(ctx context.Context) *V1TenantUIDLoginBannerUpdateParams {
	var ()
	return &V1TenantUIDLoginBannerUpdateParams{

		Context: ctx,
	}
}

// NewV1TenantUIDLoginBannerUpdateParamsWithHTTPClient creates a new V1TenantUIDLoginBannerUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1TenantUIDLoginBannerUpdateParamsWithHTTPClient(client *http.Client) *V1TenantUIDLoginBannerUpdateParams {
	var ()
	return &V1TenantUIDLoginBannerUpdateParams{
		HTTPClient: client,
	}
}

/*
V1TenantUIDLoginBannerUpdateParams contains all the parameters to send to the API endpoint
for the v1 tenant Uid login banner update operation typically these are written to a http.Request
*/
type V1TenantUIDLoginBannerUpdateParams struct {

	/*Body*/
	Body *models.V1LoginBannerSettings
	/*TenantUID*/
	TenantUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 tenant Uid login banner update params
func (o *V1TenantUIDLoginBannerUpdateParams) WithTimeout(timeout time.Duration) *V1TenantUIDLoginBannerUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 tenant Uid login banner update params
func (o *V1TenantUIDLoginBannerUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 tenant Uid login banner update params
func (o *V1TenantUIDLoginBannerUpdateParams) WithContext(ctx context.Context) *V1TenantUIDLoginBannerUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 tenant Uid login banner update params
func (o *V1TenantUIDLoginBannerUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 tenant Uid login banner update params
func (o *V1TenantUIDLoginBannerUpdateParams) WithHTTPClient(client *http.Client) *V1TenantUIDLoginBannerUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 tenant Uid login banner update params
func (o *V1TenantUIDLoginBannerUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 tenant Uid login banner update params
func (o *V1TenantUIDLoginBannerUpdateParams) WithBody(body *models.V1LoginBannerSettings) *V1TenantUIDLoginBannerUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 tenant Uid login banner update params
func (o *V1TenantUIDLoginBannerUpdateParams) SetBody(body *models.V1LoginBannerSettings) {
	o.Body = body
}

// WithTenantUID adds the tenantUID to the v1 tenant Uid login banner update params
func (o *V1TenantUIDLoginBannerUpdateParams) WithTenantUID(tenantUID string) *V1TenantUIDLoginBannerUpdateParams {
	o.SetTenantUID(tenantUID)
	return o
}

// SetTenantUID adds the tenantUid to the v1 tenant Uid login banner update params
func (o *V1TenantUIDLoginBannerUpdateParams) SetTenantUID(tenantUID string) {
	o.TenantUID = tenantUID
}

// WriteToRequest writes these params to a swagger request
func (o *V1TenantUIDLoginBannerUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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