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

// NewV1CloudAccountsCoxEdgeCreateParams creates a new V1CloudAccountsCoxEdgeCreateParams object
// with the default values initialized.
func NewV1CloudAccountsCoxEdgeCreateParams() *V1CloudAccountsCoxEdgeCreateParams {
	var ()
	return &V1CloudAccountsCoxEdgeCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudAccountsCoxEdgeCreateParamsWithTimeout creates a new V1CloudAccountsCoxEdgeCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudAccountsCoxEdgeCreateParamsWithTimeout(timeout time.Duration) *V1CloudAccountsCoxEdgeCreateParams {
	var ()
	return &V1CloudAccountsCoxEdgeCreateParams{

		timeout: timeout,
	}
}

// NewV1CloudAccountsCoxEdgeCreateParamsWithContext creates a new V1CloudAccountsCoxEdgeCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudAccountsCoxEdgeCreateParamsWithContext(ctx context.Context) *V1CloudAccountsCoxEdgeCreateParams {
	var ()
	return &V1CloudAccountsCoxEdgeCreateParams{

		Context: ctx,
	}
}

// NewV1CloudAccountsCoxEdgeCreateParamsWithHTTPClient creates a new V1CloudAccountsCoxEdgeCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudAccountsCoxEdgeCreateParamsWithHTTPClient(client *http.Client) *V1CloudAccountsCoxEdgeCreateParams {
	var ()
	return &V1CloudAccountsCoxEdgeCreateParams{
		HTTPClient: client,
	}
}

/*V1CloudAccountsCoxEdgeCreateParams contains all the parameters to send to the API endpoint
for the v1 cloud accounts cox edge create operation typically these are written to a http.Request
*/
type V1CloudAccountsCoxEdgeCreateParams struct {

	/*Body
	  Request payload to validate CoxEdge cloud account

	*/
	Body *models.V1CoxEdgeAccount

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cloud accounts cox edge create params
func (o *V1CloudAccountsCoxEdgeCreateParams) WithTimeout(timeout time.Duration) *V1CloudAccountsCoxEdgeCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud accounts cox edge create params
func (o *V1CloudAccountsCoxEdgeCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud accounts cox edge create params
func (o *V1CloudAccountsCoxEdgeCreateParams) WithContext(ctx context.Context) *V1CloudAccountsCoxEdgeCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud accounts cox edge create params
func (o *V1CloudAccountsCoxEdgeCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud accounts cox edge create params
func (o *V1CloudAccountsCoxEdgeCreateParams) WithHTTPClient(client *http.Client) *V1CloudAccountsCoxEdgeCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud accounts cox edge create params
func (o *V1CloudAccountsCoxEdgeCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 cloud accounts cox edge create params
func (o *V1CloudAccountsCoxEdgeCreateParams) WithBody(body *models.V1CoxEdgeAccount) *V1CloudAccountsCoxEdgeCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 cloud accounts cox edge create params
func (o *V1CloudAccountsCoxEdgeCreateParams) SetBody(body *models.V1CoxEdgeAccount) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudAccountsCoxEdgeCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
