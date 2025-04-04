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
)

// NewV1SsoIdpsParams creates a new V1SsoIdpsParams object
// with the default values initialized.
func NewV1SsoIdpsParams() *V1SsoIdpsParams {

	return &V1SsoIdpsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1SsoIdpsParamsWithTimeout creates a new V1SsoIdpsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1SsoIdpsParamsWithTimeout(timeout time.Duration) *V1SsoIdpsParams {

	return &V1SsoIdpsParams{

		timeout: timeout,
	}
}

// NewV1SsoIdpsParamsWithContext creates a new V1SsoIdpsParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1SsoIdpsParamsWithContext(ctx context.Context) *V1SsoIdpsParams {

	return &V1SsoIdpsParams{

		Context: ctx,
	}
}

// NewV1SsoIdpsParamsWithHTTPClient creates a new V1SsoIdpsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1SsoIdpsParamsWithHTTPClient(client *http.Client) *V1SsoIdpsParams {

	return &V1SsoIdpsParams{
		HTTPClient: client,
	}
}

/*
V1SsoIdpsParams contains all the parameters to send to the API endpoint
for the v1 sso idps operation typically these are written to a http.Request
*/
type V1SsoIdpsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 sso idps params
func (o *V1SsoIdpsParams) WithTimeout(timeout time.Duration) *V1SsoIdpsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 sso idps params
func (o *V1SsoIdpsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 sso idps params
func (o *V1SsoIdpsParams) WithContext(ctx context.Context) *V1SsoIdpsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 sso idps params
func (o *V1SsoIdpsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 sso idps params
func (o *V1SsoIdpsParams) WithHTTPClient(client *http.Client) *V1SsoIdpsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 sso idps params
func (o *V1SsoIdpsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *V1SsoIdpsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
