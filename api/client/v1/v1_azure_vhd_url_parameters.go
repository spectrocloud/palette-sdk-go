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

// NewV1AzureVhdURLParams creates a new V1AzureVhdURLParams object
// with the default values initialized.
func NewV1AzureVhdURLParams() *V1AzureVhdURLParams {
	var ()
	return &V1AzureVhdURLParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1AzureVhdURLParamsWithTimeout creates a new V1AzureVhdURLParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1AzureVhdURLParamsWithTimeout(timeout time.Duration) *V1AzureVhdURLParams {
	var ()
	return &V1AzureVhdURLParams{

		timeout: timeout,
	}
}

// NewV1AzureVhdURLParamsWithContext creates a new V1AzureVhdURLParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1AzureVhdURLParamsWithContext(ctx context.Context) *V1AzureVhdURLParams {
	var ()
	return &V1AzureVhdURLParams{

		Context: ctx,
	}
}

// NewV1AzureVhdURLParamsWithHTTPClient creates a new V1AzureVhdURLParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1AzureVhdURLParamsWithHTTPClient(client *http.Client) *V1AzureVhdURLParams {
	var ()
	return &V1AzureVhdURLParams{
		HTTPClient: client,
	}
}

/*
V1AzureVhdURLParams contains all the parameters to send to the API endpoint
for the v1 azure vhd Url operation typically these are written to a http.Request
*/
type V1AzureVhdURLParams struct {

	/*Vhd
	  vhd location for which Azure vhd url is requested

	*/
	Vhd string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 azure vhd Url params
func (o *V1AzureVhdURLParams) WithTimeout(timeout time.Duration) *V1AzureVhdURLParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 azure vhd Url params
func (o *V1AzureVhdURLParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 azure vhd Url params
func (o *V1AzureVhdURLParams) WithContext(ctx context.Context) *V1AzureVhdURLParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 azure vhd Url params
func (o *V1AzureVhdURLParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 azure vhd Url params
func (o *V1AzureVhdURLParams) WithHTTPClient(client *http.Client) *V1AzureVhdURLParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 azure vhd Url params
func (o *V1AzureVhdURLParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithVhd adds the vhd to the v1 azure vhd Url params
func (o *V1AzureVhdURLParams) WithVhd(vhd string) *V1AzureVhdURLParams {
	o.SetVhd(vhd)
	return o
}

// SetVhd adds the vhd to the v1 azure vhd Url params
func (o *V1AzureVhdURLParams) SetVhd(vhd string) {
	o.Vhd = vhd
}

// WriteToRequest writes these params to a swagger request
func (o *V1AzureVhdURLParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param vhd
	if err := r.SetPathParam("vhd", o.Vhd); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}