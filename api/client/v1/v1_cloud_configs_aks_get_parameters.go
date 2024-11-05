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

// NewV1CloudConfigsAksGetParams creates a new V1CloudConfigsAksGetParams object
// with the default values initialized.
func NewV1CloudConfigsAksGetParams() *V1CloudConfigsAksGetParams {
	var ()
	return &V1CloudConfigsAksGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudConfigsAksGetParamsWithTimeout creates a new V1CloudConfigsAksGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudConfigsAksGetParamsWithTimeout(timeout time.Duration) *V1CloudConfigsAksGetParams {
	var ()
	return &V1CloudConfigsAksGetParams{

		timeout: timeout,
	}
}

// NewV1CloudConfigsAksGetParamsWithContext creates a new V1CloudConfigsAksGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudConfigsAksGetParamsWithContext(ctx context.Context) *V1CloudConfigsAksGetParams {
	var ()
	return &V1CloudConfigsAksGetParams{

		Context: ctx,
	}
}

// NewV1CloudConfigsAksGetParamsWithHTTPClient creates a new V1CloudConfigsAksGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudConfigsAksGetParamsWithHTTPClient(client *http.Client) *V1CloudConfigsAksGetParams {
	var ()
	return &V1CloudConfigsAksGetParams{
		HTTPClient: client,
	}
}

/*
V1CloudConfigsAksGetParams contains all the parameters to send to the API endpoint
for the v1 cloud configs aks get operation typically these are written to a http.Request
*/
type V1CloudConfigsAksGetParams struct {

	/*ConfigUID
	  Cluster's cloud config uid

	*/
	ConfigUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cloud configs aks get params
func (o *V1CloudConfigsAksGetParams) WithTimeout(timeout time.Duration) *V1CloudConfigsAksGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud configs aks get params
func (o *V1CloudConfigsAksGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud configs aks get params
func (o *V1CloudConfigsAksGetParams) WithContext(ctx context.Context) *V1CloudConfigsAksGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud configs aks get params
func (o *V1CloudConfigsAksGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud configs aks get params
func (o *V1CloudConfigsAksGetParams) WithHTTPClient(client *http.Client) *V1CloudConfigsAksGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud configs aks get params
func (o *V1CloudConfigsAksGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConfigUID adds the configUID to the v1 cloud configs aks get params
func (o *V1CloudConfigsAksGetParams) WithConfigUID(configUID string) *V1CloudConfigsAksGetParams {
	o.SetConfigUID(configUID)
	return o
}

// SetConfigUID adds the configUid to the v1 cloud configs aks get params
func (o *V1CloudConfigsAksGetParams) SetConfigUID(configUID string) {
	o.ConfigUID = configUID
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudConfigsAksGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
