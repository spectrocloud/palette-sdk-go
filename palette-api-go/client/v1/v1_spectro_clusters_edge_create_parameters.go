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

	"github.com/spectrocloud/palette-api-go/models"
)

// NewV1SpectroClustersEdgeCreateParams creates a new V1SpectroClustersEdgeCreateParams object
// with the default values initialized.
func NewV1SpectroClustersEdgeCreateParams() *V1SpectroClustersEdgeCreateParams {
	var ()
	return &V1SpectroClustersEdgeCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1SpectroClustersEdgeCreateParamsWithTimeout creates a new V1SpectroClustersEdgeCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1SpectroClustersEdgeCreateParamsWithTimeout(timeout time.Duration) *V1SpectroClustersEdgeCreateParams {
	var ()
	return &V1SpectroClustersEdgeCreateParams{

		timeout: timeout,
	}
}

// NewV1SpectroClustersEdgeCreateParamsWithContext creates a new V1SpectroClustersEdgeCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1SpectroClustersEdgeCreateParamsWithContext(ctx context.Context) *V1SpectroClustersEdgeCreateParams {
	var ()
	return &V1SpectroClustersEdgeCreateParams{

		Context: ctx,
	}
}

// NewV1SpectroClustersEdgeCreateParamsWithHTTPClient creates a new V1SpectroClustersEdgeCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1SpectroClustersEdgeCreateParamsWithHTTPClient(client *http.Client) *V1SpectroClustersEdgeCreateParams {
	var ()
	return &V1SpectroClustersEdgeCreateParams{
		HTTPClient: client,
	}
}

/*V1SpectroClustersEdgeCreateParams contains all the parameters to send to the API endpoint
for the v1 spectro clusters edge create operation typically these are written to a http.Request
*/
type V1SpectroClustersEdgeCreateParams struct {

	/*Body*/
	Body *models.V1SpectroEdgeClusterEntity

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 spectro clusters edge create params
func (o *V1SpectroClustersEdgeCreateParams) WithTimeout(timeout time.Duration) *V1SpectroClustersEdgeCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 spectro clusters edge create params
func (o *V1SpectroClustersEdgeCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 spectro clusters edge create params
func (o *V1SpectroClustersEdgeCreateParams) WithContext(ctx context.Context) *V1SpectroClustersEdgeCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 spectro clusters edge create params
func (o *V1SpectroClustersEdgeCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 spectro clusters edge create params
func (o *V1SpectroClustersEdgeCreateParams) WithHTTPClient(client *http.Client) *V1SpectroClustersEdgeCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 spectro clusters edge create params
func (o *V1SpectroClustersEdgeCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 spectro clusters edge create params
func (o *V1SpectroClustersEdgeCreateParams) WithBody(body *models.V1SpectroEdgeClusterEntity) *V1SpectroClustersEdgeCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 spectro clusters edge create params
func (o *V1SpectroClustersEdgeCreateParams) SetBody(body *models.V1SpectroEdgeClusterEntity) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *V1SpectroClustersEdgeCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
