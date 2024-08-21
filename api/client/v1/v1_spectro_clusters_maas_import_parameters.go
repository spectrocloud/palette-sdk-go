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

// NewV1SpectroClustersMaasImportParams creates a new V1SpectroClustersMaasImportParams object
// with the default values initialized.
func NewV1SpectroClustersMaasImportParams() *V1SpectroClustersMaasImportParams {
	var ()
	return &V1SpectroClustersMaasImportParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1SpectroClustersMaasImportParamsWithTimeout creates a new V1SpectroClustersMaasImportParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1SpectroClustersMaasImportParamsWithTimeout(timeout time.Duration) *V1SpectroClustersMaasImportParams {
	var ()
	return &V1SpectroClustersMaasImportParams{

		timeout: timeout,
	}
}

// NewV1SpectroClustersMaasImportParamsWithContext creates a new V1SpectroClustersMaasImportParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1SpectroClustersMaasImportParamsWithContext(ctx context.Context) *V1SpectroClustersMaasImportParams {
	var ()
	return &V1SpectroClustersMaasImportParams{

		Context: ctx,
	}
}

// NewV1SpectroClustersMaasImportParamsWithHTTPClient creates a new V1SpectroClustersMaasImportParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1SpectroClustersMaasImportParamsWithHTTPClient(client *http.Client) *V1SpectroClustersMaasImportParams {
	var ()
	return &V1SpectroClustersMaasImportParams{
		HTTPClient: client,
	}
}

/*
V1SpectroClustersMaasImportParams contains all the parameters to send to the API endpoint
for the v1 spectro clusters maas import operation typically these are written to a http.Request
*/
type V1SpectroClustersMaasImportParams struct {

	/*Body*/
	Body *models.V1SpectroMaasClusterImportEntity

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 spectro clusters maas import params
func (o *V1SpectroClustersMaasImportParams) WithTimeout(timeout time.Duration) *V1SpectroClustersMaasImportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 spectro clusters maas import params
func (o *V1SpectroClustersMaasImportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 spectro clusters maas import params
func (o *V1SpectroClustersMaasImportParams) WithContext(ctx context.Context) *V1SpectroClustersMaasImportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 spectro clusters maas import params
func (o *V1SpectroClustersMaasImportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 spectro clusters maas import params
func (o *V1SpectroClustersMaasImportParams) WithHTTPClient(client *http.Client) *V1SpectroClustersMaasImportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 spectro clusters maas import params
func (o *V1SpectroClustersMaasImportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 spectro clusters maas import params
func (o *V1SpectroClustersMaasImportParams) WithBody(body *models.V1SpectroMaasClusterImportEntity) *V1SpectroClustersMaasImportParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 spectro clusters maas import params
func (o *V1SpectroClustersMaasImportParams) SetBody(body *models.V1SpectroMaasClusterImportEntity) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *V1SpectroClustersMaasImportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}