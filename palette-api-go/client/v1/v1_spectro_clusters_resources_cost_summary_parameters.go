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

// NewV1SpectroClustersResourcesCostSummaryParams creates a new V1SpectroClustersResourcesCostSummaryParams object
// with the default values initialized.
func NewV1SpectroClustersResourcesCostSummaryParams() *V1SpectroClustersResourcesCostSummaryParams {
	var ()
	return &V1SpectroClustersResourcesCostSummaryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1SpectroClustersResourcesCostSummaryParamsWithTimeout creates a new V1SpectroClustersResourcesCostSummaryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1SpectroClustersResourcesCostSummaryParamsWithTimeout(timeout time.Duration) *V1SpectroClustersResourcesCostSummaryParams {
	var ()
	return &V1SpectroClustersResourcesCostSummaryParams{

		timeout: timeout,
	}
}

// NewV1SpectroClustersResourcesCostSummaryParamsWithContext creates a new V1SpectroClustersResourcesCostSummaryParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1SpectroClustersResourcesCostSummaryParamsWithContext(ctx context.Context) *V1SpectroClustersResourcesCostSummaryParams {
	var ()
	return &V1SpectroClustersResourcesCostSummaryParams{

		Context: ctx,
	}
}

// NewV1SpectroClustersResourcesCostSummaryParamsWithHTTPClient creates a new V1SpectroClustersResourcesCostSummaryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1SpectroClustersResourcesCostSummaryParamsWithHTTPClient(client *http.Client) *V1SpectroClustersResourcesCostSummaryParams {
	var ()
	return &V1SpectroClustersResourcesCostSummaryParams{
		HTTPClient: client,
	}
}

/*V1SpectroClustersResourcesCostSummaryParams contains all the parameters to send to the API endpoint
for the v1 spectro clusters resources cost summary operation typically these are written to a http.Request
*/
type V1SpectroClustersResourcesCostSummaryParams struct {

	/*Body*/
	Body *models.V1ResourceCostSummarySpec

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 spectro clusters resources cost summary params
func (o *V1SpectroClustersResourcesCostSummaryParams) WithTimeout(timeout time.Duration) *V1SpectroClustersResourcesCostSummaryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 spectro clusters resources cost summary params
func (o *V1SpectroClustersResourcesCostSummaryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 spectro clusters resources cost summary params
func (o *V1SpectroClustersResourcesCostSummaryParams) WithContext(ctx context.Context) *V1SpectroClustersResourcesCostSummaryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 spectro clusters resources cost summary params
func (o *V1SpectroClustersResourcesCostSummaryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 spectro clusters resources cost summary params
func (o *V1SpectroClustersResourcesCostSummaryParams) WithHTTPClient(client *http.Client) *V1SpectroClustersResourcesCostSummaryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 spectro clusters resources cost summary params
func (o *V1SpectroClustersResourcesCostSummaryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 spectro clusters resources cost summary params
func (o *V1SpectroClustersResourcesCostSummaryParams) WithBody(body *models.V1ResourceCostSummarySpec) *V1SpectroClustersResourcesCostSummaryParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 spectro clusters resources cost summary params
func (o *V1SpectroClustersResourcesCostSummaryParams) SetBody(body *models.V1ResourceCostSummarySpec) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *V1SpectroClustersResourcesCostSummaryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
