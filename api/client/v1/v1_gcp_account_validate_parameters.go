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

// NewV1GcpAccountValidateParams creates a new V1GcpAccountValidateParams object
// with the default values initialized.
func NewV1GcpAccountValidateParams() *V1GcpAccountValidateParams {
	var ()
	return &V1GcpAccountValidateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1GcpAccountValidateParamsWithTimeout creates a new V1GcpAccountValidateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1GcpAccountValidateParamsWithTimeout(timeout time.Duration) *V1GcpAccountValidateParams {
	var ()
	return &V1GcpAccountValidateParams{

		timeout: timeout,
	}
}

// NewV1GcpAccountValidateParamsWithContext creates a new V1GcpAccountValidateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1GcpAccountValidateParamsWithContext(ctx context.Context) *V1GcpAccountValidateParams {
	var ()
	return &V1GcpAccountValidateParams{

		Context: ctx,
	}
}

// NewV1GcpAccountValidateParamsWithHTTPClient creates a new V1GcpAccountValidateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1GcpAccountValidateParamsWithHTTPClient(client *http.Client) *V1GcpAccountValidateParams {
	var ()
	return &V1GcpAccountValidateParams{
		HTTPClient: client,
	}
}

/*
V1GcpAccountValidateParams contains all the parameters to send to the API endpoint
for the v1 gcp account validate operation typically these are written to a http.Request
*/
type V1GcpAccountValidateParams struct {

	/*GcpCloudAccount
	  Uid for the specific GCP cloud account

	*/
	GcpCloudAccount *models.V1GcpCloudAccountValidateEntity

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 gcp account validate params
func (o *V1GcpAccountValidateParams) WithTimeout(timeout time.Duration) *V1GcpAccountValidateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 gcp account validate params
func (o *V1GcpAccountValidateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 gcp account validate params
func (o *V1GcpAccountValidateParams) WithContext(ctx context.Context) *V1GcpAccountValidateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 gcp account validate params
func (o *V1GcpAccountValidateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 gcp account validate params
func (o *V1GcpAccountValidateParams) WithHTTPClient(client *http.Client) *V1GcpAccountValidateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 gcp account validate params
func (o *V1GcpAccountValidateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGcpCloudAccount adds the gcpCloudAccount to the v1 gcp account validate params
func (o *V1GcpAccountValidateParams) WithGcpCloudAccount(gcpCloudAccount *models.V1GcpCloudAccountValidateEntity) *V1GcpAccountValidateParams {
	o.SetGcpCloudAccount(gcpCloudAccount)
	return o
}

// SetGcpCloudAccount adds the gcpCloudAccount to the v1 gcp account validate params
func (o *V1GcpAccountValidateParams) SetGcpCloudAccount(gcpCloudAccount *models.V1GcpCloudAccountValidateEntity) {
	o.GcpCloudAccount = gcpCloudAccount
}

// WriteToRequest writes these params to a swagger request
func (o *V1GcpAccountValidateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.GcpCloudAccount != nil {
		if err := r.SetBodyParam(o.GcpCloudAccount); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
