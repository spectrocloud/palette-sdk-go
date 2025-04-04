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

// NewV1CloudAccountsCustomGetParams creates a new V1CloudAccountsCustomGetParams object
// with the default values initialized.
func NewV1CloudAccountsCustomGetParams() *V1CloudAccountsCustomGetParams {
	var ()
	return &V1CloudAccountsCustomGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudAccountsCustomGetParamsWithTimeout creates a new V1CloudAccountsCustomGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudAccountsCustomGetParamsWithTimeout(timeout time.Duration) *V1CloudAccountsCustomGetParams {
	var ()
	return &V1CloudAccountsCustomGetParams{

		timeout: timeout,
	}
}

// NewV1CloudAccountsCustomGetParamsWithContext creates a new V1CloudAccountsCustomGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudAccountsCustomGetParamsWithContext(ctx context.Context) *V1CloudAccountsCustomGetParams {
	var ()
	return &V1CloudAccountsCustomGetParams{

		Context: ctx,
	}
}

// NewV1CloudAccountsCustomGetParamsWithHTTPClient creates a new V1CloudAccountsCustomGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudAccountsCustomGetParamsWithHTTPClient(client *http.Client) *V1CloudAccountsCustomGetParams {
	var ()
	return &V1CloudAccountsCustomGetParams{
		HTTPClient: client,
	}
}

/*
V1CloudAccountsCustomGetParams contains all the parameters to send to the API endpoint
for the v1 cloud accounts custom get operation typically these are written to a http.Request
*/
type V1CloudAccountsCustomGetParams struct {

	/*CloudType
	  Custom cloud type

	*/
	CloudType string
	/*UID
	  Custom cloud account uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cloud accounts custom get params
func (o *V1CloudAccountsCustomGetParams) WithTimeout(timeout time.Duration) *V1CloudAccountsCustomGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud accounts custom get params
func (o *V1CloudAccountsCustomGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud accounts custom get params
func (o *V1CloudAccountsCustomGetParams) WithContext(ctx context.Context) *V1CloudAccountsCustomGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud accounts custom get params
func (o *V1CloudAccountsCustomGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud accounts custom get params
func (o *V1CloudAccountsCustomGetParams) WithHTTPClient(client *http.Client) *V1CloudAccountsCustomGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud accounts custom get params
func (o *V1CloudAccountsCustomGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudType adds the cloudType to the v1 cloud accounts custom get params
func (o *V1CloudAccountsCustomGetParams) WithCloudType(cloudType string) *V1CloudAccountsCustomGetParams {
	o.SetCloudType(cloudType)
	return o
}

// SetCloudType adds the cloudType to the v1 cloud accounts custom get params
func (o *V1CloudAccountsCustomGetParams) SetCloudType(cloudType string) {
	o.CloudType = cloudType
}

// WithUID adds the uid to the v1 cloud accounts custom get params
func (o *V1CloudAccountsCustomGetParams) WithUID(uid string) *V1CloudAccountsCustomGetParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 cloud accounts custom get params
func (o *V1CloudAccountsCustomGetParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudAccountsCustomGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cloudType
	if err := r.SetPathParam("cloudType", o.CloudType); err != nil {
		return err
	}

	// path param uid
	if err := r.SetPathParam("uid", o.UID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
