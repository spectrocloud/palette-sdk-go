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

// NewV1CloudConfigsMachinePoolsMachineUidsGetParams creates a new V1CloudConfigsMachinePoolsMachineUidsGetParams object
// with the default values initialized.
func NewV1CloudConfigsMachinePoolsMachineUidsGetParams() *V1CloudConfigsMachinePoolsMachineUidsGetParams {
	var ()
	return &V1CloudConfigsMachinePoolsMachineUidsGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudConfigsMachinePoolsMachineUidsGetParamsWithTimeout creates a new V1CloudConfigsMachinePoolsMachineUidsGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudConfigsMachinePoolsMachineUidsGetParamsWithTimeout(timeout time.Duration) *V1CloudConfigsMachinePoolsMachineUidsGetParams {
	var ()
	return &V1CloudConfigsMachinePoolsMachineUidsGetParams{

		timeout: timeout,
	}
}

// NewV1CloudConfigsMachinePoolsMachineUidsGetParamsWithContext creates a new V1CloudConfigsMachinePoolsMachineUidsGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudConfigsMachinePoolsMachineUidsGetParamsWithContext(ctx context.Context) *V1CloudConfigsMachinePoolsMachineUidsGetParams {
	var ()
	return &V1CloudConfigsMachinePoolsMachineUidsGetParams{

		Context: ctx,
	}
}

// NewV1CloudConfigsMachinePoolsMachineUidsGetParamsWithHTTPClient creates a new V1CloudConfigsMachinePoolsMachineUidsGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudConfigsMachinePoolsMachineUidsGetParamsWithHTTPClient(client *http.Client) *V1CloudConfigsMachinePoolsMachineUidsGetParams {
	var ()
	return &V1CloudConfigsMachinePoolsMachineUidsGetParams{
		HTTPClient: client,
	}
}

/*
V1CloudConfigsMachinePoolsMachineUidsGetParams contains all the parameters to send to the API endpoint
for the v1 cloud configs machine pools machine uids get operation typically these are written to a http.Request
*/
type V1CloudConfigsMachinePoolsMachineUidsGetParams struct {

	/*ConfigUID
	  Cluster's cloud config uid

	*/
	ConfigUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cloud configs machine pools machine uids get params
func (o *V1CloudConfigsMachinePoolsMachineUidsGetParams) WithTimeout(timeout time.Duration) *V1CloudConfigsMachinePoolsMachineUidsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud configs machine pools machine uids get params
func (o *V1CloudConfigsMachinePoolsMachineUidsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud configs machine pools machine uids get params
func (o *V1CloudConfigsMachinePoolsMachineUidsGetParams) WithContext(ctx context.Context) *V1CloudConfigsMachinePoolsMachineUidsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud configs machine pools machine uids get params
func (o *V1CloudConfigsMachinePoolsMachineUidsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud configs machine pools machine uids get params
func (o *V1CloudConfigsMachinePoolsMachineUidsGetParams) WithHTTPClient(client *http.Client) *V1CloudConfigsMachinePoolsMachineUidsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud configs machine pools machine uids get params
func (o *V1CloudConfigsMachinePoolsMachineUidsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConfigUID adds the configUID to the v1 cloud configs machine pools machine uids get params
func (o *V1CloudConfigsMachinePoolsMachineUidsGetParams) WithConfigUID(configUID string) *V1CloudConfigsMachinePoolsMachineUidsGetParams {
	o.SetConfigUID(configUID)
	return o
}

// SetConfigUID adds the configUid to the v1 cloud configs machine pools machine uids get params
func (o *V1CloudConfigsMachinePoolsMachineUidsGetParams) SetConfigUID(configUID string) {
	o.ConfigUID = configUID
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudConfigsMachinePoolsMachineUidsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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