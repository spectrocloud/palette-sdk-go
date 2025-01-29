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

// NewV1CloudAccountsOpenStackGetParams creates a new V1CloudAccountsOpenStackGetParams object
// with the default values initialized.
func NewV1CloudAccountsOpenStackGetParams() *V1CloudAccountsOpenStackGetParams {
	var ()
	return &V1CloudAccountsOpenStackGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudAccountsOpenStackGetParamsWithTimeout creates a new V1CloudAccountsOpenStackGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudAccountsOpenStackGetParamsWithTimeout(timeout time.Duration) *V1CloudAccountsOpenStackGetParams {
	var ()
	return &V1CloudAccountsOpenStackGetParams{

		timeout: timeout,
	}
}

// NewV1CloudAccountsOpenStackGetParamsWithContext creates a new V1CloudAccountsOpenStackGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudAccountsOpenStackGetParamsWithContext(ctx context.Context) *V1CloudAccountsOpenStackGetParams {
	var ()
	return &V1CloudAccountsOpenStackGetParams{

		Context: ctx,
	}
}

// NewV1CloudAccountsOpenStackGetParamsWithHTTPClient creates a new V1CloudAccountsOpenStackGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudAccountsOpenStackGetParamsWithHTTPClient(client *http.Client) *V1CloudAccountsOpenStackGetParams {
	var ()
	return &V1CloudAccountsOpenStackGetParams{
		HTTPClient: client,
	}
}

/*
V1CloudAccountsOpenStackGetParams contains all the parameters to send to the API endpoint
for the v1 cloud accounts open stack get operation typically these are written to a http.Request
*/
type V1CloudAccountsOpenStackGetParams struct {

	/*UID
	  OpenStack cloud account uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cloud accounts open stack get params
func (o *V1CloudAccountsOpenStackGetParams) WithTimeout(timeout time.Duration) *V1CloudAccountsOpenStackGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud accounts open stack get params
func (o *V1CloudAccountsOpenStackGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud accounts open stack get params
func (o *V1CloudAccountsOpenStackGetParams) WithContext(ctx context.Context) *V1CloudAccountsOpenStackGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud accounts open stack get params
func (o *V1CloudAccountsOpenStackGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud accounts open stack get params
func (o *V1CloudAccountsOpenStackGetParams) WithHTTPClient(client *http.Client) *V1CloudAccountsOpenStackGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud accounts open stack get params
func (o *V1CloudAccountsOpenStackGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUID adds the uid to the v1 cloud accounts open stack get params
func (o *V1CloudAccountsOpenStackGetParams) WithUID(uid string) *V1CloudAccountsOpenStackGetParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 cloud accounts open stack get params
func (o *V1CloudAccountsOpenStackGetParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudAccountsOpenStackGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
