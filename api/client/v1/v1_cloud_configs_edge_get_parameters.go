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

// NewV1CloudConfigsEdgeGetParams creates a new V1CloudConfigsEdgeGetParams object
// with the default values initialized.
func NewV1CloudConfigsEdgeGetParams() *V1CloudConfigsEdgeGetParams {
	var ()
	return &V1CloudConfigsEdgeGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudConfigsEdgeGetParamsWithTimeout creates a new V1CloudConfigsEdgeGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudConfigsEdgeGetParamsWithTimeout(timeout time.Duration) *V1CloudConfigsEdgeGetParams {
	var ()
	return &V1CloudConfigsEdgeGetParams{

		timeout: timeout,
	}
}

// NewV1CloudConfigsEdgeGetParamsWithContext creates a new V1CloudConfigsEdgeGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudConfigsEdgeGetParamsWithContext(ctx context.Context) *V1CloudConfigsEdgeGetParams {
	var ()
	return &V1CloudConfigsEdgeGetParams{

		Context: ctx,
	}
}

// NewV1CloudConfigsEdgeGetParamsWithHTTPClient creates a new V1CloudConfigsEdgeGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudConfigsEdgeGetParamsWithHTTPClient(client *http.Client) *V1CloudConfigsEdgeGetParams {
	var ()
	return &V1CloudConfigsEdgeGetParams{
		HTTPClient: client,
	}
}

/*
V1CloudConfigsEdgeGetParams contains all the parameters to send to the API endpoint
for the v1 cloud configs edge get operation typically these are written to a http.Request
*/
type V1CloudConfigsEdgeGetParams struct {

	/*ConfigUID
	  Cluster's cloud config uid

	*/
	ConfigUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cloud configs edge get params
func (o *V1CloudConfigsEdgeGetParams) WithTimeout(timeout time.Duration) *V1CloudConfigsEdgeGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud configs edge get params
func (o *V1CloudConfigsEdgeGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud configs edge get params
func (o *V1CloudConfigsEdgeGetParams) WithContext(ctx context.Context) *V1CloudConfigsEdgeGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud configs edge get params
func (o *V1CloudConfigsEdgeGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud configs edge get params
func (o *V1CloudConfigsEdgeGetParams) WithHTTPClient(client *http.Client) *V1CloudConfigsEdgeGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud configs edge get params
func (o *V1CloudConfigsEdgeGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConfigUID adds the configUID to the v1 cloud configs edge get params
func (o *V1CloudConfigsEdgeGetParams) WithConfigUID(configUID string) *V1CloudConfigsEdgeGetParams {
	o.SetConfigUID(configUID)
	return o
}

// SetConfigUID adds the configUid to the v1 cloud configs edge get params
func (o *V1CloudConfigsEdgeGetParams) SetConfigUID(configUID string) {
	o.ConfigUID = configUID
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudConfigsEdgeGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param configUid
	if err := r.SetPathParam("configUid", o.ConfigUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}