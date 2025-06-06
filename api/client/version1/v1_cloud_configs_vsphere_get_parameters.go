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

// NewV1CloudConfigsVsphereGetParams creates a new V1CloudConfigsVsphereGetParams object
// with the default values initialized.
func NewV1CloudConfigsVsphereGetParams() *V1CloudConfigsVsphereGetParams {
	var ()
	return &V1CloudConfigsVsphereGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudConfigsVsphereGetParamsWithTimeout creates a new V1CloudConfigsVsphereGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudConfigsVsphereGetParamsWithTimeout(timeout time.Duration) *V1CloudConfigsVsphereGetParams {
	var ()
	return &V1CloudConfigsVsphereGetParams{

		timeout: timeout,
	}
}

// NewV1CloudConfigsVsphereGetParamsWithContext creates a new V1CloudConfigsVsphereGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudConfigsVsphereGetParamsWithContext(ctx context.Context) *V1CloudConfigsVsphereGetParams {
	var ()
	return &V1CloudConfigsVsphereGetParams{

		Context: ctx,
	}
}

// NewV1CloudConfigsVsphereGetParamsWithHTTPClient creates a new V1CloudConfigsVsphereGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudConfigsVsphereGetParamsWithHTTPClient(client *http.Client) *V1CloudConfigsVsphereGetParams {
	var ()
	return &V1CloudConfigsVsphereGetParams{
		HTTPClient: client,
	}
}

/*
V1CloudConfigsVsphereGetParams contains all the parameters to send to the API endpoint
for the v1 cloud configs vsphere get operation typically these are written to a http.Request
*/
type V1CloudConfigsVsphereGetParams struct {

	/*ConfigUID
	  Cluster's cloud config uid

	*/
	ConfigUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cloud configs vsphere get params
func (o *V1CloudConfigsVsphereGetParams) WithTimeout(timeout time.Duration) *V1CloudConfigsVsphereGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud configs vsphere get params
func (o *V1CloudConfigsVsphereGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud configs vsphere get params
func (o *V1CloudConfigsVsphereGetParams) WithContext(ctx context.Context) *V1CloudConfigsVsphereGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud configs vsphere get params
func (o *V1CloudConfigsVsphereGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud configs vsphere get params
func (o *V1CloudConfigsVsphereGetParams) WithHTTPClient(client *http.Client) *V1CloudConfigsVsphereGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud configs vsphere get params
func (o *V1CloudConfigsVsphereGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConfigUID adds the configUID to the v1 cloud configs vsphere get params
func (o *V1CloudConfigsVsphereGetParams) WithConfigUID(configUID string) *V1CloudConfigsVsphereGetParams {
	o.SetConfigUID(configUID)
	return o
}

// SetConfigUID adds the configUid to the v1 cloud configs vsphere get params
func (o *V1CloudConfigsVsphereGetParams) SetConfigUID(configUID string) {
	o.ConfigUID = configUID
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudConfigsVsphereGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
