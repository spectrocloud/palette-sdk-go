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

// NewV1SpectroClustersMetadataParams creates a new V1SpectroClustersMetadataParams object
// with the default values initialized.
func NewV1SpectroClustersMetadataParams() *V1SpectroClustersMetadataParams {
	var ()
	return &V1SpectroClustersMetadataParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1SpectroClustersMetadataParamsWithTimeout creates a new V1SpectroClustersMetadataParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1SpectroClustersMetadataParamsWithTimeout(timeout time.Duration) *V1SpectroClustersMetadataParams {
	var ()
	return &V1SpectroClustersMetadataParams{

		timeout: timeout,
	}
}

// NewV1SpectroClustersMetadataParamsWithContext creates a new V1SpectroClustersMetadataParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1SpectroClustersMetadataParamsWithContext(ctx context.Context) *V1SpectroClustersMetadataParams {
	var ()
	return &V1SpectroClustersMetadataParams{

		Context: ctx,
	}
}

// NewV1SpectroClustersMetadataParamsWithHTTPClient creates a new V1SpectroClustersMetadataParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1SpectroClustersMetadataParamsWithHTTPClient(client *http.Client) *V1SpectroClustersMetadataParams {
	var ()
	return &V1SpectroClustersMetadataParams{
		HTTPClient: client,
	}
}

/*
V1SpectroClustersMetadataParams contains all the parameters to send to the API endpoint
for the v1 spectro clusters metadata operation typically these are written to a http.Request
*/
type V1SpectroClustersMetadataParams struct {

	/*Body*/
	Body *models.V1SpectroClusterMetadataSpec

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 spectro clusters metadata params
func (o *V1SpectroClustersMetadataParams) WithTimeout(timeout time.Duration) *V1SpectroClustersMetadataParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 spectro clusters metadata params
func (o *V1SpectroClustersMetadataParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 spectro clusters metadata params
func (o *V1SpectroClustersMetadataParams) WithContext(ctx context.Context) *V1SpectroClustersMetadataParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 spectro clusters metadata params
func (o *V1SpectroClustersMetadataParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 spectro clusters metadata params
func (o *V1SpectroClustersMetadataParams) WithHTTPClient(client *http.Client) *V1SpectroClustersMetadataParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 spectro clusters metadata params
func (o *V1SpectroClustersMetadataParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 spectro clusters metadata params
func (o *V1SpectroClustersMetadataParams) WithBody(body *models.V1SpectroClusterMetadataSpec) *V1SpectroClustersMetadataParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 spectro clusters metadata params
func (o *V1SpectroClustersMetadataParams) SetBody(body *models.V1SpectroClusterMetadataSpec) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *V1SpectroClustersMetadataParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
