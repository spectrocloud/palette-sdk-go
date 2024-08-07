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

// NewV1TenantUIDAssetsCertsUIDUpdateParams creates a new V1TenantUIDAssetsCertsUIDUpdateParams object
// with the default values initialized.
func NewV1TenantUIDAssetsCertsUIDUpdateParams() *V1TenantUIDAssetsCertsUIDUpdateParams {
	var ()
	return &V1TenantUIDAssetsCertsUIDUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1TenantUIDAssetsCertsUIDUpdateParamsWithTimeout creates a new V1TenantUIDAssetsCertsUIDUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1TenantUIDAssetsCertsUIDUpdateParamsWithTimeout(timeout time.Duration) *V1TenantUIDAssetsCertsUIDUpdateParams {
	var ()
	return &V1TenantUIDAssetsCertsUIDUpdateParams{

		timeout: timeout,
	}
}

// NewV1TenantUIDAssetsCertsUIDUpdateParamsWithContext creates a new V1TenantUIDAssetsCertsUIDUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1TenantUIDAssetsCertsUIDUpdateParamsWithContext(ctx context.Context) *V1TenantUIDAssetsCertsUIDUpdateParams {
	var ()
	return &V1TenantUIDAssetsCertsUIDUpdateParams{

		Context: ctx,
	}
}

// NewV1TenantUIDAssetsCertsUIDUpdateParamsWithHTTPClient creates a new V1TenantUIDAssetsCertsUIDUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1TenantUIDAssetsCertsUIDUpdateParamsWithHTTPClient(client *http.Client) *V1TenantUIDAssetsCertsUIDUpdateParams {
	var ()
	return &V1TenantUIDAssetsCertsUIDUpdateParams{
		HTTPClient: client,
	}
}

/*
V1TenantUIDAssetsCertsUIDUpdateParams contains all the parameters to send to the API endpoint
for the v1 tenant Uid assets certs Uid update operation typically these are written to a http.Request
*/
type V1TenantUIDAssetsCertsUIDUpdateParams struct {

	/*Body*/
	Body *models.V1TenantAssetCert
	/*CertificateUID*/
	CertificateUID string
	/*TenantUID*/
	TenantUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 tenant Uid assets certs Uid update params
func (o *V1TenantUIDAssetsCertsUIDUpdateParams) WithTimeout(timeout time.Duration) *V1TenantUIDAssetsCertsUIDUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 tenant Uid assets certs Uid update params
func (o *V1TenantUIDAssetsCertsUIDUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 tenant Uid assets certs Uid update params
func (o *V1TenantUIDAssetsCertsUIDUpdateParams) WithContext(ctx context.Context) *V1TenantUIDAssetsCertsUIDUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 tenant Uid assets certs Uid update params
func (o *V1TenantUIDAssetsCertsUIDUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 tenant Uid assets certs Uid update params
func (o *V1TenantUIDAssetsCertsUIDUpdateParams) WithHTTPClient(client *http.Client) *V1TenantUIDAssetsCertsUIDUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 tenant Uid assets certs Uid update params
func (o *V1TenantUIDAssetsCertsUIDUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 tenant Uid assets certs Uid update params
func (o *V1TenantUIDAssetsCertsUIDUpdateParams) WithBody(body *models.V1TenantAssetCert) *V1TenantUIDAssetsCertsUIDUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 tenant Uid assets certs Uid update params
func (o *V1TenantUIDAssetsCertsUIDUpdateParams) SetBody(body *models.V1TenantAssetCert) {
	o.Body = body
}

// WithCertificateUID adds the certificateUID to the v1 tenant Uid assets certs Uid update params
func (o *V1TenantUIDAssetsCertsUIDUpdateParams) WithCertificateUID(certificateUID string) *V1TenantUIDAssetsCertsUIDUpdateParams {
	o.SetCertificateUID(certificateUID)
	return o
}

// SetCertificateUID adds the certificateUid to the v1 tenant Uid assets certs Uid update params
func (o *V1TenantUIDAssetsCertsUIDUpdateParams) SetCertificateUID(certificateUID string) {
	o.CertificateUID = certificateUID
}

// WithTenantUID adds the tenantUID to the v1 tenant Uid assets certs Uid update params
func (o *V1TenantUIDAssetsCertsUIDUpdateParams) WithTenantUID(tenantUID string) *V1TenantUIDAssetsCertsUIDUpdateParams {
	o.SetTenantUID(tenantUID)
	return o
}

// SetTenantUID adds the tenantUid to the v1 tenant Uid assets certs Uid update params
func (o *V1TenantUIDAssetsCertsUIDUpdateParams) SetTenantUID(tenantUID string) {
	o.TenantUID = tenantUID
}

// WriteToRequest writes these params to a swagger request
func (o *V1TenantUIDAssetsCertsUIDUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param certificateUid
	if err := r.SetPathParam("certificateUid", o.CertificateUID); err != nil {
		return err
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
