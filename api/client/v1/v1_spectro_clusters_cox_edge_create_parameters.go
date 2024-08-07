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

// NewV1SpectroClustersCoxEdgeCreateParams creates a new V1SpectroClustersCoxEdgeCreateParams object
// with the default values initialized.
func NewV1SpectroClustersCoxEdgeCreateParams() *V1SpectroClustersCoxEdgeCreateParams {
	var ()
	return &V1SpectroClustersCoxEdgeCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1SpectroClustersCoxEdgeCreateParamsWithTimeout creates a new V1SpectroClustersCoxEdgeCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1SpectroClustersCoxEdgeCreateParamsWithTimeout(timeout time.Duration) *V1SpectroClustersCoxEdgeCreateParams {
	var ()
	return &V1SpectroClustersCoxEdgeCreateParams{

		timeout: timeout,
	}
}

// NewV1SpectroClustersCoxEdgeCreateParamsWithContext creates a new V1SpectroClustersCoxEdgeCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1SpectroClustersCoxEdgeCreateParamsWithContext(ctx context.Context) *V1SpectroClustersCoxEdgeCreateParams {
	var ()
	return &V1SpectroClustersCoxEdgeCreateParams{

		Context: ctx,
	}
}

// NewV1SpectroClustersCoxEdgeCreateParamsWithHTTPClient creates a new V1SpectroClustersCoxEdgeCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1SpectroClustersCoxEdgeCreateParamsWithHTTPClient(client *http.Client) *V1SpectroClustersCoxEdgeCreateParams {
	var ()
	return &V1SpectroClustersCoxEdgeCreateParams{
		HTTPClient: client,
	}
}

/*
V1SpectroClustersCoxEdgeCreateParams contains all the parameters to send to the API endpoint
for the v1 spectro clusters cox edge create operation typically these are written to a http.Request
*/
type V1SpectroClustersCoxEdgeCreateParams struct {

	/*Body*/
	Body *models.V1SpectroCoxEdgeClusterEntity

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 spectro clusters cox edge create params
func (o *V1SpectroClustersCoxEdgeCreateParams) WithTimeout(timeout time.Duration) *V1SpectroClustersCoxEdgeCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 spectro clusters cox edge create params
func (o *V1SpectroClustersCoxEdgeCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 spectro clusters cox edge create params
func (o *V1SpectroClustersCoxEdgeCreateParams) WithContext(ctx context.Context) *V1SpectroClustersCoxEdgeCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 spectro clusters cox edge create params
func (o *V1SpectroClustersCoxEdgeCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 spectro clusters cox edge create params
func (o *V1SpectroClustersCoxEdgeCreateParams) WithHTTPClient(client *http.Client) *V1SpectroClustersCoxEdgeCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 spectro clusters cox edge create params
func (o *V1SpectroClustersCoxEdgeCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 spectro clusters cox edge create params
func (o *V1SpectroClustersCoxEdgeCreateParams) WithBody(body *models.V1SpectroCoxEdgeClusterEntity) *V1SpectroClustersCoxEdgeCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 spectro clusters cox edge create params
func (o *V1SpectroClustersCoxEdgeCreateParams) SetBody(body *models.V1SpectroCoxEdgeClusterEntity) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *V1SpectroClustersCoxEdgeCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
