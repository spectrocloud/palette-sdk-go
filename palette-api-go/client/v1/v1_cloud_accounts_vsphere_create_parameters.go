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

// NewV1CloudAccountsVsphereCreateParams creates a new V1CloudAccountsVsphereCreateParams object
// with the default values initialized.
func NewV1CloudAccountsVsphereCreateParams() *V1CloudAccountsVsphereCreateParams {
	var ()
	return &V1CloudAccountsVsphereCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudAccountsVsphereCreateParamsWithTimeout creates a new V1CloudAccountsVsphereCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudAccountsVsphereCreateParamsWithTimeout(timeout time.Duration) *V1CloudAccountsVsphereCreateParams {
	var ()
	return &V1CloudAccountsVsphereCreateParams{

		timeout: timeout,
	}
}

// NewV1CloudAccountsVsphereCreateParamsWithContext creates a new V1CloudAccountsVsphereCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudAccountsVsphereCreateParamsWithContext(ctx context.Context) *V1CloudAccountsVsphereCreateParams {
	var ()
	return &V1CloudAccountsVsphereCreateParams{

		Context: ctx,
	}
}

// NewV1CloudAccountsVsphereCreateParamsWithHTTPClient creates a new V1CloudAccountsVsphereCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudAccountsVsphereCreateParamsWithHTTPClient(client *http.Client) *V1CloudAccountsVsphereCreateParams {
	var ()
	return &V1CloudAccountsVsphereCreateParams{
		HTTPClient: client,
	}
}

/*V1CloudAccountsVsphereCreateParams contains all the parameters to send to the API endpoint
for the v1 cloud accounts vsphere create operation typically these are written to a http.Request
*/
type V1CloudAccountsVsphereCreateParams struct {

	/*Body
	  Request payload to validate VSphere cloud account

	*/
	Body *models.V1VsphereAccount

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cloud accounts vsphere create params
func (o *V1CloudAccountsVsphereCreateParams) WithTimeout(timeout time.Duration) *V1CloudAccountsVsphereCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud accounts vsphere create params
func (o *V1CloudAccountsVsphereCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud accounts vsphere create params
func (o *V1CloudAccountsVsphereCreateParams) WithContext(ctx context.Context) *V1CloudAccountsVsphereCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud accounts vsphere create params
func (o *V1CloudAccountsVsphereCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud accounts vsphere create params
func (o *V1CloudAccountsVsphereCreateParams) WithHTTPClient(client *http.Client) *V1CloudAccountsVsphereCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud accounts vsphere create params
func (o *V1CloudAccountsVsphereCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 cloud accounts vsphere create params
func (o *V1CloudAccountsVsphereCreateParams) WithBody(body *models.V1VsphereAccount) *V1CloudAccountsVsphereCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 cloud accounts vsphere create params
func (o *V1CloudAccountsVsphereCreateParams) SetBody(body *models.V1VsphereAccount) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudAccountsVsphereCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
