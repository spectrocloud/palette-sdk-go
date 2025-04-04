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

// NewV1SpectroClustersGenericImportParams creates a new V1SpectroClustersGenericImportParams object
// with the default values initialized.
func NewV1SpectroClustersGenericImportParams() *V1SpectroClustersGenericImportParams {
	var ()
	return &V1SpectroClustersGenericImportParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1SpectroClustersGenericImportParamsWithTimeout creates a new V1SpectroClustersGenericImportParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1SpectroClustersGenericImportParamsWithTimeout(timeout time.Duration) *V1SpectroClustersGenericImportParams {
	var ()
	return &V1SpectroClustersGenericImportParams{

		timeout: timeout,
	}
}

// NewV1SpectroClustersGenericImportParamsWithContext creates a new V1SpectroClustersGenericImportParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1SpectroClustersGenericImportParamsWithContext(ctx context.Context) *V1SpectroClustersGenericImportParams {
	var ()
	return &V1SpectroClustersGenericImportParams{

		Context: ctx,
	}
}

// NewV1SpectroClustersGenericImportParamsWithHTTPClient creates a new V1SpectroClustersGenericImportParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1SpectroClustersGenericImportParamsWithHTTPClient(client *http.Client) *V1SpectroClustersGenericImportParams {
	var ()
	return &V1SpectroClustersGenericImportParams{
		HTTPClient: client,
	}
}

/*
V1SpectroClustersGenericImportParams contains all the parameters to send to the API endpoint
for the v1 spectro clusters generic import operation typically these are written to a http.Request
*/
type V1SpectroClustersGenericImportParams struct {

	/*Body*/
	Body *models.V1SpectroGenericClusterImportEntity

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 spectro clusters generic import params
func (o *V1SpectroClustersGenericImportParams) WithTimeout(timeout time.Duration) *V1SpectroClustersGenericImportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 spectro clusters generic import params
func (o *V1SpectroClustersGenericImportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 spectro clusters generic import params
func (o *V1SpectroClustersGenericImportParams) WithContext(ctx context.Context) *V1SpectroClustersGenericImportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 spectro clusters generic import params
func (o *V1SpectroClustersGenericImportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 spectro clusters generic import params
func (o *V1SpectroClustersGenericImportParams) WithHTTPClient(client *http.Client) *V1SpectroClustersGenericImportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 spectro clusters generic import params
func (o *V1SpectroClustersGenericImportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 spectro clusters generic import params
func (o *V1SpectroClustersGenericImportParams) WithBody(body *models.V1SpectroGenericClusterImportEntity) *V1SpectroClustersGenericImportParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 spectro clusters generic import params
func (o *V1SpectroClustersGenericImportParams) SetBody(body *models.V1SpectroGenericClusterImportEntity) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *V1SpectroClustersGenericImportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
