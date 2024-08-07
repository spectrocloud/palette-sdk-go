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

// NewV1UsersSummaryGetParams creates a new V1UsersSummaryGetParams object
// with the default values initialized.
func NewV1UsersSummaryGetParams() *V1UsersSummaryGetParams {
	var ()
	return &V1UsersSummaryGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1UsersSummaryGetParamsWithTimeout creates a new V1UsersSummaryGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1UsersSummaryGetParamsWithTimeout(timeout time.Duration) *V1UsersSummaryGetParams {
	var ()
	return &V1UsersSummaryGetParams{

		timeout: timeout,
	}
}

// NewV1UsersSummaryGetParamsWithContext creates a new V1UsersSummaryGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1UsersSummaryGetParamsWithContext(ctx context.Context) *V1UsersSummaryGetParams {
	var ()
	return &V1UsersSummaryGetParams{

		Context: ctx,
	}
}

// NewV1UsersSummaryGetParamsWithHTTPClient creates a new V1UsersSummaryGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1UsersSummaryGetParamsWithHTTPClient(client *http.Client) *V1UsersSummaryGetParams {
	var ()
	return &V1UsersSummaryGetParams{
		HTTPClient: client,
	}
}

/*V1UsersSummaryGetParams contains all the parameters to send to the API endpoint
for the v1 users summary get operation typically these are written to a http.Request
*/
type V1UsersSummaryGetParams struct {

	/*Body*/
	Body *models.V1UsersSummarySpec

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 users summary get params
func (o *V1UsersSummaryGetParams) WithTimeout(timeout time.Duration) *V1UsersSummaryGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 users summary get params
func (o *V1UsersSummaryGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 users summary get params
func (o *V1UsersSummaryGetParams) WithContext(ctx context.Context) *V1UsersSummaryGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 users summary get params
func (o *V1UsersSummaryGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 users summary get params
func (o *V1UsersSummaryGetParams) WithHTTPClient(client *http.Client) *V1UsersSummaryGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 users summary get params
func (o *V1UsersSummaryGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 users summary get params
func (o *V1UsersSummaryGetParams) WithBody(body *models.V1UsersSummarySpec) *V1UsersSummaryGetParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 users summary get params
func (o *V1UsersSummaryGetParams) SetBody(body *models.V1UsersSummarySpec) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *V1UsersSummaryGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
