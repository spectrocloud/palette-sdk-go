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

// NewV1CloudAccountsTencentDeleteParams creates a new V1CloudAccountsTencentDeleteParams object
// with the default values initialized.
func NewV1CloudAccountsTencentDeleteParams() *V1CloudAccountsTencentDeleteParams {
	var ()
	return &V1CloudAccountsTencentDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudAccountsTencentDeleteParamsWithTimeout creates a new V1CloudAccountsTencentDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudAccountsTencentDeleteParamsWithTimeout(timeout time.Duration) *V1CloudAccountsTencentDeleteParams {
	var ()
	return &V1CloudAccountsTencentDeleteParams{

		timeout: timeout,
	}
}

// NewV1CloudAccountsTencentDeleteParamsWithContext creates a new V1CloudAccountsTencentDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudAccountsTencentDeleteParamsWithContext(ctx context.Context) *V1CloudAccountsTencentDeleteParams {
	var ()
	return &V1CloudAccountsTencentDeleteParams{

		Context: ctx,
	}
}

// NewV1CloudAccountsTencentDeleteParamsWithHTTPClient creates a new V1CloudAccountsTencentDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudAccountsTencentDeleteParamsWithHTTPClient(client *http.Client) *V1CloudAccountsTencentDeleteParams {
	var ()
	return &V1CloudAccountsTencentDeleteParams{
		HTTPClient: client,
	}
}

/*
V1CloudAccountsTencentDeleteParams contains all the parameters to send to the API endpoint
for the v1 cloud accounts tencent delete operation typically these are written to a http.Request
*/
type V1CloudAccountsTencentDeleteParams struct {

	/*UID
	  Tencent cloud account uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cloud accounts tencent delete params
func (o *V1CloudAccountsTencentDeleteParams) WithTimeout(timeout time.Duration) *V1CloudAccountsTencentDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud accounts tencent delete params
func (o *V1CloudAccountsTencentDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud accounts tencent delete params
func (o *V1CloudAccountsTencentDeleteParams) WithContext(ctx context.Context) *V1CloudAccountsTencentDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud accounts tencent delete params
func (o *V1CloudAccountsTencentDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud accounts tencent delete params
func (o *V1CloudAccountsTencentDeleteParams) WithHTTPClient(client *http.Client) *V1CloudAccountsTencentDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud accounts tencent delete params
func (o *V1CloudAccountsTencentDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUID adds the uid to the v1 cloud accounts tencent delete params
func (o *V1CloudAccountsTencentDeleteParams) WithUID(uid string) *V1CloudAccountsTencentDeleteParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 cloud accounts tencent delete params
func (o *V1CloudAccountsTencentDeleteParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudAccountsTencentDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
