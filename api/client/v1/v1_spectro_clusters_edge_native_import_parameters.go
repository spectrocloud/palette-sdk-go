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

// NewV1SpectroClustersEdgeNativeImportParams creates a new V1SpectroClustersEdgeNativeImportParams object
// with the default values initialized.
func NewV1SpectroClustersEdgeNativeImportParams() *V1SpectroClustersEdgeNativeImportParams {
	var ()
	return &V1SpectroClustersEdgeNativeImportParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1SpectroClustersEdgeNativeImportParamsWithTimeout creates a new V1SpectroClustersEdgeNativeImportParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1SpectroClustersEdgeNativeImportParamsWithTimeout(timeout time.Duration) *V1SpectroClustersEdgeNativeImportParams {
	var ()
	return &V1SpectroClustersEdgeNativeImportParams{

		timeout: timeout,
	}
}

// NewV1SpectroClustersEdgeNativeImportParamsWithContext creates a new V1SpectroClustersEdgeNativeImportParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1SpectroClustersEdgeNativeImportParamsWithContext(ctx context.Context) *V1SpectroClustersEdgeNativeImportParams {
	var ()
	return &V1SpectroClustersEdgeNativeImportParams{

		Context: ctx,
	}
}

// NewV1SpectroClustersEdgeNativeImportParamsWithHTTPClient creates a new V1SpectroClustersEdgeNativeImportParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1SpectroClustersEdgeNativeImportParamsWithHTTPClient(client *http.Client) *V1SpectroClustersEdgeNativeImportParams {
	var ()
	return &V1SpectroClustersEdgeNativeImportParams{
		HTTPClient: client,
	}
}

/*V1SpectroClustersEdgeNativeImportParams contains all the parameters to send to the API endpoint
for the v1 spectro clusters edge native import operation typically these are written to a http.Request
*/
type V1SpectroClustersEdgeNativeImportParams struct {

	/*Body*/
	Body *models.V1SpectroEdgeNativeClusterImportEntity

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 spectro clusters edge native import params
func (o *V1SpectroClustersEdgeNativeImportParams) WithTimeout(timeout time.Duration) *V1SpectroClustersEdgeNativeImportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 spectro clusters edge native import params
func (o *V1SpectroClustersEdgeNativeImportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 spectro clusters edge native import params
func (o *V1SpectroClustersEdgeNativeImportParams) WithContext(ctx context.Context) *V1SpectroClustersEdgeNativeImportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 spectro clusters edge native import params
func (o *V1SpectroClustersEdgeNativeImportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 spectro clusters edge native import params
func (o *V1SpectroClustersEdgeNativeImportParams) WithHTTPClient(client *http.Client) *V1SpectroClustersEdgeNativeImportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 spectro clusters edge native import params
func (o *V1SpectroClustersEdgeNativeImportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 spectro clusters edge native import params
func (o *V1SpectroClustersEdgeNativeImportParams) WithBody(body *models.V1SpectroEdgeNativeClusterImportEntity) *V1SpectroClustersEdgeNativeImportParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 spectro clusters edge native import params
func (o *V1SpectroClustersEdgeNativeImportParams) SetBody(body *models.V1SpectroEdgeNativeClusterImportEntity) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *V1SpectroClustersEdgeNativeImportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
