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

// NewV1CoxEdgeBaseUrlsParams creates a new V1CoxEdgeBaseUrlsParams object
// with the default values initialized.
func NewV1CoxEdgeBaseUrlsParams() *V1CoxEdgeBaseUrlsParams {

	return &V1CoxEdgeBaseUrlsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CoxEdgeBaseUrlsParamsWithTimeout creates a new V1CoxEdgeBaseUrlsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CoxEdgeBaseUrlsParamsWithTimeout(timeout time.Duration) *V1CoxEdgeBaseUrlsParams {

	return &V1CoxEdgeBaseUrlsParams{

		timeout: timeout,
	}
}

// NewV1CoxEdgeBaseUrlsParamsWithContext creates a new V1CoxEdgeBaseUrlsParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CoxEdgeBaseUrlsParamsWithContext(ctx context.Context) *V1CoxEdgeBaseUrlsParams {

	return &V1CoxEdgeBaseUrlsParams{

		Context: ctx,
	}
}

// NewV1CoxEdgeBaseUrlsParamsWithHTTPClient creates a new V1CoxEdgeBaseUrlsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CoxEdgeBaseUrlsParamsWithHTTPClient(client *http.Client) *V1CoxEdgeBaseUrlsParams {

	return &V1CoxEdgeBaseUrlsParams{
		HTTPClient: client,
	}
}

/*
V1CoxEdgeBaseUrlsParams contains all the parameters to send to the API endpoint
for the v1 cox edge base urls operation typically these are written to a http.Request
*/
type V1CoxEdgeBaseUrlsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cox edge base urls params
func (o *V1CoxEdgeBaseUrlsParams) WithTimeout(timeout time.Duration) *V1CoxEdgeBaseUrlsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cox edge base urls params
func (o *V1CoxEdgeBaseUrlsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cox edge base urls params
func (o *V1CoxEdgeBaseUrlsParams) WithContext(ctx context.Context) *V1CoxEdgeBaseUrlsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cox edge base urls params
func (o *V1CoxEdgeBaseUrlsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cox edge base urls params
func (o *V1CoxEdgeBaseUrlsParams) WithHTTPClient(client *http.Client) *V1CoxEdgeBaseUrlsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cox edge base urls params
func (o *V1CoxEdgeBaseUrlsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *V1CoxEdgeBaseUrlsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}