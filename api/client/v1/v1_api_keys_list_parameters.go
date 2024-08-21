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

// NewV1APIKeysListParams creates a new V1APIKeysListParams object
// with the default values initialized.
func NewV1APIKeysListParams() *V1APIKeysListParams {

	return &V1APIKeysListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1APIKeysListParamsWithTimeout creates a new V1APIKeysListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1APIKeysListParamsWithTimeout(timeout time.Duration) *V1APIKeysListParams {

	return &V1APIKeysListParams{

		timeout: timeout,
	}
}

// NewV1APIKeysListParamsWithContext creates a new V1APIKeysListParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1APIKeysListParamsWithContext(ctx context.Context) *V1APIKeysListParams {

	return &V1APIKeysListParams{

		Context: ctx,
	}
}

// NewV1APIKeysListParamsWithHTTPClient creates a new V1APIKeysListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1APIKeysListParamsWithHTTPClient(client *http.Client) *V1APIKeysListParams {

	return &V1APIKeysListParams{
		HTTPClient: client,
	}
}

/*
V1APIKeysListParams contains all the parameters to send to the API endpoint
for the v1 Api keys list operation typically these are written to a http.Request
*/
type V1APIKeysListParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 Api keys list params
func (o *V1APIKeysListParams) WithTimeout(timeout time.Duration) *V1APIKeysListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 Api keys list params
func (o *V1APIKeysListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 Api keys list params
func (o *V1APIKeysListParams) WithContext(ctx context.Context) *V1APIKeysListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 Api keys list params
func (o *V1APIKeysListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 Api keys list params
func (o *V1APIKeysListParams) WithHTTPClient(client *http.Client) *V1APIKeysListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 Api keys list params
func (o *V1APIKeysListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *V1APIKeysListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}