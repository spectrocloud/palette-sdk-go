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

// NewV1CloudAccountsAwsDeleteParams creates a new V1CloudAccountsAwsDeleteParams object
// with the default values initialized.
func NewV1CloudAccountsAwsDeleteParams() *V1CloudAccountsAwsDeleteParams {
	var ()
	return &V1CloudAccountsAwsDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudAccountsAwsDeleteParamsWithTimeout creates a new V1CloudAccountsAwsDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudAccountsAwsDeleteParamsWithTimeout(timeout time.Duration) *V1CloudAccountsAwsDeleteParams {
	var ()
	return &V1CloudAccountsAwsDeleteParams{

		timeout: timeout,
	}
}

// NewV1CloudAccountsAwsDeleteParamsWithContext creates a new V1CloudAccountsAwsDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudAccountsAwsDeleteParamsWithContext(ctx context.Context) *V1CloudAccountsAwsDeleteParams {
	var ()
	return &V1CloudAccountsAwsDeleteParams{

		Context: ctx,
	}
}

// NewV1CloudAccountsAwsDeleteParamsWithHTTPClient creates a new V1CloudAccountsAwsDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudAccountsAwsDeleteParamsWithHTTPClient(client *http.Client) *V1CloudAccountsAwsDeleteParams {
	var ()
	return &V1CloudAccountsAwsDeleteParams{
		HTTPClient: client,
	}
}

/*
V1CloudAccountsAwsDeleteParams contains all the parameters to send to the API endpoint
for the v1 cloud accounts aws delete operation typically these are written to a http.Request
*/
type V1CloudAccountsAwsDeleteParams struct {

	/*UID
	  AWS cloud account uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cloud accounts aws delete params
func (o *V1CloudAccountsAwsDeleteParams) WithTimeout(timeout time.Duration) *V1CloudAccountsAwsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud accounts aws delete params
func (o *V1CloudAccountsAwsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud accounts aws delete params
func (o *V1CloudAccountsAwsDeleteParams) WithContext(ctx context.Context) *V1CloudAccountsAwsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud accounts aws delete params
func (o *V1CloudAccountsAwsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud accounts aws delete params
func (o *V1CloudAccountsAwsDeleteParams) WithHTTPClient(client *http.Client) *V1CloudAccountsAwsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud accounts aws delete params
func (o *V1CloudAccountsAwsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUID adds the uid to the v1 cloud accounts aws delete params
func (o *V1CloudAccountsAwsDeleteParams) WithUID(uid string) *V1CloudAccountsAwsDeleteParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 cloud accounts aws delete params
func (o *V1CloudAccountsAwsDeleteParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudAccountsAwsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param uid
	if err := r.SetPathParam("uid", o.UID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
