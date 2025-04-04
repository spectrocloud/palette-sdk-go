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

// NewV1CloudConfigsGkeGetParams creates a new V1CloudConfigsGkeGetParams object
// with the default values initialized.
func NewV1CloudConfigsGkeGetParams() *V1CloudConfigsGkeGetParams {
	var ()
	return &V1CloudConfigsGkeGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudConfigsGkeGetParamsWithTimeout creates a new V1CloudConfigsGkeGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudConfigsGkeGetParamsWithTimeout(timeout time.Duration) *V1CloudConfigsGkeGetParams {
	var ()
	return &V1CloudConfigsGkeGetParams{

		timeout: timeout,
	}
}

// NewV1CloudConfigsGkeGetParamsWithContext creates a new V1CloudConfigsGkeGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudConfigsGkeGetParamsWithContext(ctx context.Context) *V1CloudConfigsGkeGetParams {
	var ()
	return &V1CloudConfigsGkeGetParams{

		Context: ctx,
	}
}

// NewV1CloudConfigsGkeGetParamsWithHTTPClient creates a new V1CloudConfigsGkeGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudConfigsGkeGetParamsWithHTTPClient(client *http.Client) *V1CloudConfigsGkeGetParams {
	var ()
	return &V1CloudConfigsGkeGetParams{
		HTTPClient: client,
	}
}

/*
V1CloudConfigsGkeGetParams contains all the parameters to send to the API endpoint
for the v1 cloud configs gke get operation typically these are written to a http.Request
*/
type V1CloudConfigsGkeGetParams struct {

	/*ConfigUID
	  Cluster's cloud config uid

	*/
	ConfigUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cloud configs gke get params
func (o *V1CloudConfigsGkeGetParams) WithTimeout(timeout time.Duration) *V1CloudConfigsGkeGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud configs gke get params
func (o *V1CloudConfigsGkeGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud configs gke get params
func (o *V1CloudConfigsGkeGetParams) WithContext(ctx context.Context) *V1CloudConfigsGkeGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud configs gke get params
func (o *V1CloudConfigsGkeGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud configs gke get params
func (o *V1CloudConfigsGkeGetParams) WithHTTPClient(client *http.Client) *V1CloudConfigsGkeGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud configs gke get params
func (o *V1CloudConfigsGkeGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConfigUID adds the configUID to the v1 cloud configs gke get params
func (o *V1CloudConfigsGkeGetParams) WithConfigUID(configUID string) *V1CloudConfigsGkeGetParams {
	o.SetConfigUID(configUID)
	return o
}

// SetConfigUID adds the configUid to the v1 cloud configs gke get params
func (o *V1CloudConfigsGkeGetParams) SetConfigUID(configUID string) {
	o.ConfigUID = configUID
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudConfigsGkeGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
