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

// NewV1ClusterGroupsDeveloperCreditUsageGetParams creates a new V1ClusterGroupsDeveloperCreditUsageGetParams object
// with the default values initialized.
func NewV1ClusterGroupsDeveloperCreditUsageGetParams() *V1ClusterGroupsDeveloperCreditUsageGetParams {
	var ()
	return &V1ClusterGroupsDeveloperCreditUsageGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1ClusterGroupsDeveloperCreditUsageGetParamsWithTimeout creates a new V1ClusterGroupsDeveloperCreditUsageGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1ClusterGroupsDeveloperCreditUsageGetParamsWithTimeout(timeout time.Duration) *V1ClusterGroupsDeveloperCreditUsageGetParams {
	var ()
	return &V1ClusterGroupsDeveloperCreditUsageGetParams{

		timeout: timeout,
	}
}

// NewV1ClusterGroupsDeveloperCreditUsageGetParamsWithContext creates a new V1ClusterGroupsDeveloperCreditUsageGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1ClusterGroupsDeveloperCreditUsageGetParamsWithContext(ctx context.Context) *V1ClusterGroupsDeveloperCreditUsageGetParams {
	var ()
	return &V1ClusterGroupsDeveloperCreditUsageGetParams{

		Context: ctx,
	}
}

// NewV1ClusterGroupsDeveloperCreditUsageGetParamsWithHTTPClient creates a new V1ClusterGroupsDeveloperCreditUsageGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1ClusterGroupsDeveloperCreditUsageGetParamsWithHTTPClient(client *http.Client) *V1ClusterGroupsDeveloperCreditUsageGetParams {
	var ()
	return &V1ClusterGroupsDeveloperCreditUsageGetParams{
		HTTPClient: client,
	}
}

/*
V1ClusterGroupsDeveloperCreditUsageGetParams contains all the parameters to send to the API endpoint
for the v1 cluster groups developer credit usage get operation typically these are written to a http.Request
*/
type V1ClusterGroupsDeveloperCreditUsageGetParams struct {

	/*Scope*/
	Scope string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cluster groups developer credit usage get params
func (o *V1ClusterGroupsDeveloperCreditUsageGetParams) WithTimeout(timeout time.Duration) *V1ClusterGroupsDeveloperCreditUsageGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cluster groups developer credit usage get params
func (o *V1ClusterGroupsDeveloperCreditUsageGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cluster groups developer credit usage get params
func (o *V1ClusterGroupsDeveloperCreditUsageGetParams) WithContext(ctx context.Context) *V1ClusterGroupsDeveloperCreditUsageGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cluster groups developer credit usage get params
func (o *V1ClusterGroupsDeveloperCreditUsageGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cluster groups developer credit usage get params
func (o *V1ClusterGroupsDeveloperCreditUsageGetParams) WithHTTPClient(client *http.Client) *V1ClusterGroupsDeveloperCreditUsageGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cluster groups developer credit usage get params
func (o *V1ClusterGroupsDeveloperCreditUsageGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithScope adds the scope to the v1 cluster groups developer credit usage get params
func (o *V1ClusterGroupsDeveloperCreditUsageGetParams) WithScope(scope string) *V1ClusterGroupsDeveloperCreditUsageGetParams {
	o.SetScope(scope)
	return o
}

// SetScope adds the scope to the v1 cluster groups developer credit usage get params
func (o *V1ClusterGroupsDeveloperCreditUsageGetParams) SetScope(scope string) {
	o.Scope = scope
}

// WriteToRequest writes these params to a swagger request
func (o *V1ClusterGroupsDeveloperCreditUsageGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param scope
	if err := r.SetPathParam("scope", o.Scope); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
