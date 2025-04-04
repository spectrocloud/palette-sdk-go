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

// NewV1SpectroClustersAzureImportParams creates a new V1SpectroClustersAzureImportParams object
// with the default values initialized.
func NewV1SpectroClustersAzureImportParams() *V1SpectroClustersAzureImportParams {
	var ()
	return &V1SpectroClustersAzureImportParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1SpectroClustersAzureImportParamsWithTimeout creates a new V1SpectroClustersAzureImportParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1SpectroClustersAzureImportParamsWithTimeout(timeout time.Duration) *V1SpectroClustersAzureImportParams {
	var ()
	return &V1SpectroClustersAzureImportParams{

		timeout: timeout,
	}
}

// NewV1SpectroClustersAzureImportParamsWithContext creates a new V1SpectroClustersAzureImportParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1SpectroClustersAzureImportParamsWithContext(ctx context.Context) *V1SpectroClustersAzureImportParams {
	var ()
	return &V1SpectroClustersAzureImportParams{

		Context: ctx,
	}
}

// NewV1SpectroClustersAzureImportParamsWithHTTPClient creates a new V1SpectroClustersAzureImportParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1SpectroClustersAzureImportParamsWithHTTPClient(client *http.Client) *V1SpectroClustersAzureImportParams {
	var ()
	return &V1SpectroClustersAzureImportParams{
		HTTPClient: client,
	}
}

/*
V1SpectroClustersAzureImportParams contains all the parameters to send to the API endpoint
for the v1 spectro clusters azure import operation typically these are written to a http.Request
*/
type V1SpectroClustersAzureImportParams struct {

	/*Body*/
	Body *models.V1SpectroAzureClusterImportEntity

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 spectro clusters azure import params
func (o *V1SpectroClustersAzureImportParams) WithTimeout(timeout time.Duration) *V1SpectroClustersAzureImportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 spectro clusters azure import params
func (o *V1SpectroClustersAzureImportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 spectro clusters azure import params
func (o *V1SpectroClustersAzureImportParams) WithContext(ctx context.Context) *V1SpectroClustersAzureImportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 spectro clusters azure import params
func (o *V1SpectroClustersAzureImportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 spectro clusters azure import params
func (o *V1SpectroClustersAzureImportParams) WithHTTPClient(client *http.Client) *V1SpectroClustersAzureImportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 spectro clusters azure import params
func (o *V1SpectroClustersAzureImportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 spectro clusters azure import params
func (o *V1SpectroClustersAzureImportParams) WithBody(body *models.V1SpectroAzureClusterImportEntity) *V1SpectroClustersAzureImportParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 spectro clusters azure import params
func (o *V1SpectroClustersAzureImportParams) SetBody(body *models.V1SpectroAzureClusterImportEntity) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *V1SpectroClustersAzureImportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
