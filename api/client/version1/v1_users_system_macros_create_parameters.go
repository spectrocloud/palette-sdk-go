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

// NewV1UsersSystemMacrosCreateParams creates a new V1UsersSystemMacrosCreateParams object
// with the default values initialized.
func NewV1UsersSystemMacrosCreateParams() *V1UsersSystemMacrosCreateParams {
	var ()
	return &V1UsersSystemMacrosCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1UsersSystemMacrosCreateParamsWithTimeout creates a new V1UsersSystemMacrosCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1UsersSystemMacrosCreateParamsWithTimeout(timeout time.Duration) *V1UsersSystemMacrosCreateParams {
	var ()
	return &V1UsersSystemMacrosCreateParams{

		timeout: timeout,
	}
}

// NewV1UsersSystemMacrosCreateParamsWithContext creates a new V1UsersSystemMacrosCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1UsersSystemMacrosCreateParamsWithContext(ctx context.Context) *V1UsersSystemMacrosCreateParams {
	var ()
	return &V1UsersSystemMacrosCreateParams{

		Context: ctx,
	}
}

// NewV1UsersSystemMacrosCreateParamsWithHTTPClient creates a new V1UsersSystemMacrosCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1UsersSystemMacrosCreateParamsWithHTTPClient(client *http.Client) *V1UsersSystemMacrosCreateParams {
	var ()
	return &V1UsersSystemMacrosCreateParams{
		HTTPClient: client,
	}
}

/*
V1UsersSystemMacrosCreateParams contains all the parameters to send to the API endpoint
for the v1 users system macros create operation typically these are written to a http.Request
*/
type V1UsersSystemMacrosCreateParams struct {

	/*Body*/
	Body *models.V1Macros

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 users system macros create params
func (o *V1UsersSystemMacrosCreateParams) WithTimeout(timeout time.Duration) *V1UsersSystemMacrosCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 users system macros create params
func (o *V1UsersSystemMacrosCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 users system macros create params
func (o *V1UsersSystemMacrosCreateParams) WithContext(ctx context.Context) *V1UsersSystemMacrosCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 users system macros create params
func (o *V1UsersSystemMacrosCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 users system macros create params
func (o *V1UsersSystemMacrosCreateParams) WithHTTPClient(client *http.Client) *V1UsersSystemMacrosCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 users system macros create params
func (o *V1UsersSystemMacrosCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 users system macros create params
func (o *V1UsersSystemMacrosCreateParams) WithBody(body *models.V1Macros) *V1UsersSystemMacrosCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 users system macros create params
func (o *V1UsersSystemMacrosCreateParams) SetBody(body *models.V1Macros) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *V1UsersSystemMacrosCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
