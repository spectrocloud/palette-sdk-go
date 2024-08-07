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

// NewV1CloudAccountsMaasCreateParams creates a new V1CloudAccountsMaasCreateParams object
// with the default values initialized.
func NewV1CloudAccountsMaasCreateParams() *V1CloudAccountsMaasCreateParams {
	var ()
	return &V1CloudAccountsMaasCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudAccountsMaasCreateParamsWithTimeout creates a new V1CloudAccountsMaasCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudAccountsMaasCreateParamsWithTimeout(timeout time.Duration) *V1CloudAccountsMaasCreateParams {
	var ()
	return &V1CloudAccountsMaasCreateParams{

		timeout: timeout,
	}
}

// NewV1CloudAccountsMaasCreateParamsWithContext creates a new V1CloudAccountsMaasCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudAccountsMaasCreateParamsWithContext(ctx context.Context) *V1CloudAccountsMaasCreateParams {
	var ()
	return &V1CloudAccountsMaasCreateParams{

		Context: ctx,
	}
}

// NewV1CloudAccountsMaasCreateParamsWithHTTPClient creates a new V1CloudAccountsMaasCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudAccountsMaasCreateParamsWithHTTPClient(client *http.Client) *V1CloudAccountsMaasCreateParams {
	var ()
	return &V1CloudAccountsMaasCreateParams{
		HTTPClient: client,
	}
}

/*V1CloudAccountsMaasCreateParams contains all the parameters to send to the API endpoint
for the v1 cloud accounts maas create operation typically these are written to a http.Request
*/
type V1CloudAccountsMaasCreateParams struct {

	/*Body
	  Request payload to validate Maas cloud account

	*/
	Body *models.V1MaasAccount

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cloud accounts maas create params
func (o *V1CloudAccountsMaasCreateParams) WithTimeout(timeout time.Duration) *V1CloudAccountsMaasCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud accounts maas create params
func (o *V1CloudAccountsMaasCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud accounts maas create params
func (o *V1CloudAccountsMaasCreateParams) WithContext(ctx context.Context) *V1CloudAccountsMaasCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud accounts maas create params
func (o *V1CloudAccountsMaasCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud accounts maas create params
func (o *V1CloudAccountsMaasCreateParams) WithHTTPClient(client *http.Client) *V1CloudAccountsMaasCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud accounts maas create params
func (o *V1CloudAccountsMaasCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 cloud accounts maas create params
func (o *V1CloudAccountsMaasCreateParams) WithBody(body *models.V1MaasAccount) *V1CloudAccountsMaasCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 cloud accounts maas create params
func (o *V1CloudAccountsMaasCreateParams) SetBody(body *models.V1MaasAccount) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudAccountsMaasCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
