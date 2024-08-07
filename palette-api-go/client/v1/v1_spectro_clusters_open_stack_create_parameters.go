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

// NewV1SpectroClustersOpenStackCreateParams creates a new V1SpectroClustersOpenStackCreateParams object
// with the default values initialized.
func NewV1SpectroClustersOpenStackCreateParams() *V1SpectroClustersOpenStackCreateParams {
	var ()
	return &V1SpectroClustersOpenStackCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1SpectroClustersOpenStackCreateParamsWithTimeout creates a new V1SpectroClustersOpenStackCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1SpectroClustersOpenStackCreateParamsWithTimeout(timeout time.Duration) *V1SpectroClustersOpenStackCreateParams {
	var ()
	return &V1SpectroClustersOpenStackCreateParams{

		timeout: timeout,
	}
}

// NewV1SpectroClustersOpenStackCreateParamsWithContext creates a new V1SpectroClustersOpenStackCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1SpectroClustersOpenStackCreateParamsWithContext(ctx context.Context) *V1SpectroClustersOpenStackCreateParams {
	var ()
	return &V1SpectroClustersOpenStackCreateParams{

		Context: ctx,
	}
}

// NewV1SpectroClustersOpenStackCreateParamsWithHTTPClient creates a new V1SpectroClustersOpenStackCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1SpectroClustersOpenStackCreateParamsWithHTTPClient(client *http.Client) *V1SpectroClustersOpenStackCreateParams {
	var ()
	return &V1SpectroClustersOpenStackCreateParams{
		HTTPClient: client,
	}
}

/*V1SpectroClustersOpenStackCreateParams contains all the parameters to send to the API endpoint
for the v1 spectro clusters open stack create operation typically these are written to a http.Request
*/
type V1SpectroClustersOpenStackCreateParams struct {

	/*Body*/
	Body *models.V1SpectroOpenStackClusterEntity

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 spectro clusters open stack create params
func (o *V1SpectroClustersOpenStackCreateParams) WithTimeout(timeout time.Duration) *V1SpectroClustersOpenStackCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 spectro clusters open stack create params
func (o *V1SpectroClustersOpenStackCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 spectro clusters open stack create params
func (o *V1SpectroClustersOpenStackCreateParams) WithContext(ctx context.Context) *V1SpectroClustersOpenStackCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 spectro clusters open stack create params
func (o *V1SpectroClustersOpenStackCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 spectro clusters open stack create params
func (o *V1SpectroClustersOpenStackCreateParams) WithHTTPClient(client *http.Client) *V1SpectroClustersOpenStackCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 spectro clusters open stack create params
func (o *V1SpectroClustersOpenStackCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 spectro clusters open stack create params
func (o *V1SpectroClustersOpenStackCreateParams) WithBody(body *models.V1SpectroOpenStackClusterEntity) *V1SpectroClustersOpenStackCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 spectro clusters open stack create params
func (o *V1SpectroClustersOpenStackCreateParams) SetBody(body *models.V1SpectroOpenStackClusterEntity) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *V1SpectroClustersOpenStackCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
