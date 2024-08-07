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
)

// NewV1SystemConfigReverseProxyGetParams creates a new V1SystemConfigReverseProxyGetParams object
// with the default values initialized.
func NewV1SystemConfigReverseProxyGetParams() *V1SystemConfigReverseProxyGetParams {

	return &V1SystemConfigReverseProxyGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1SystemConfigReverseProxyGetParamsWithTimeout creates a new V1SystemConfigReverseProxyGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1SystemConfigReverseProxyGetParamsWithTimeout(timeout time.Duration) *V1SystemConfigReverseProxyGetParams {

	return &V1SystemConfigReverseProxyGetParams{

		timeout: timeout,
	}
}

// NewV1SystemConfigReverseProxyGetParamsWithContext creates a new V1SystemConfigReverseProxyGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1SystemConfigReverseProxyGetParamsWithContext(ctx context.Context) *V1SystemConfigReverseProxyGetParams {

	return &V1SystemConfigReverseProxyGetParams{

		Context: ctx,
	}
}

// NewV1SystemConfigReverseProxyGetParamsWithHTTPClient creates a new V1SystemConfigReverseProxyGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1SystemConfigReverseProxyGetParamsWithHTTPClient(client *http.Client) *V1SystemConfigReverseProxyGetParams {

	return &V1SystemConfigReverseProxyGetParams{
		HTTPClient: client,
	}
}

/*V1SystemConfigReverseProxyGetParams contains all the parameters to send to the API endpoint
for the v1 system config reverse proxy get operation typically these are written to a http.Request
*/
type V1SystemConfigReverseProxyGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 system config reverse proxy get params
func (o *V1SystemConfigReverseProxyGetParams) WithTimeout(timeout time.Duration) *V1SystemConfigReverseProxyGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 system config reverse proxy get params
func (o *V1SystemConfigReverseProxyGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 system config reverse proxy get params
func (o *V1SystemConfigReverseProxyGetParams) WithContext(ctx context.Context) *V1SystemConfigReverseProxyGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 system config reverse proxy get params
func (o *V1SystemConfigReverseProxyGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 system config reverse proxy get params
func (o *V1SystemConfigReverseProxyGetParams) WithHTTPClient(client *http.Client) *V1SystemConfigReverseProxyGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 system config reverse proxy get params
func (o *V1SystemConfigReverseProxyGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *V1SystemConfigReverseProxyGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
