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

// NewV1EdgeHostDevicesUIDDeleteParams creates a new V1EdgeHostDevicesUIDDeleteParams object
// with the default values initialized.
func NewV1EdgeHostDevicesUIDDeleteParams() *V1EdgeHostDevicesUIDDeleteParams {
	var ()
	return &V1EdgeHostDevicesUIDDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1EdgeHostDevicesUIDDeleteParamsWithTimeout creates a new V1EdgeHostDevicesUIDDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1EdgeHostDevicesUIDDeleteParamsWithTimeout(timeout time.Duration) *V1EdgeHostDevicesUIDDeleteParams {
	var ()
	return &V1EdgeHostDevicesUIDDeleteParams{

		timeout: timeout,
	}
}

// NewV1EdgeHostDevicesUIDDeleteParamsWithContext creates a new V1EdgeHostDevicesUIDDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1EdgeHostDevicesUIDDeleteParamsWithContext(ctx context.Context) *V1EdgeHostDevicesUIDDeleteParams {
	var ()
	return &V1EdgeHostDevicesUIDDeleteParams{

		Context: ctx,
	}
}

// NewV1EdgeHostDevicesUIDDeleteParamsWithHTTPClient creates a new V1EdgeHostDevicesUIDDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1EdgeHostDevicesUIDDeleteParamsWithHTTPClient(client *http.Client) *V1EdgeHostDevicesUIDDeleteParams {
	var ()
	return &V1EdgeHostDevicesUIDDeleteParams{
		HTTPClient: client,
	}
}

/*
V1EdgeHostDevicesUIDDeleteParams contains all the parameters to send to the API endpoint
for the v1 edge host devices Uid delete operation typically these are written to a http.Request
*/
type V1EdgeHostDevicesUIDDeleteParams struct {

	/*UID*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 edge host devices Uid delete params
func (o *V1EdgeHostDevicesUIDDeleteParams) WithTimeout(timeout time.Duration) *V1EdgeHostDevicesUIDDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 edge host devices Uid delete params
func (o *V1EdgeHostDevicesUIDDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 edge host devices Uid delete params
func (o *V1EdgeHostDevicesUIDDeleteParams) WithContext(ctx context.Context) *V1EdgeHostDevicesUIDDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 edge host devices Uid delete params
func (o *V1EdgeHostDevicesUIDDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 edge host devices Uid delete params
func (o *V1EdgeHostDevicesUIDDeleteParams) WithHTTPClient(client *http.Client) *V1EdgeHostDevicesUIDDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 edge host devices Uid delete params
func (o *V1EdgeHostDevicesUIDDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUID adds the uid to the v1 edge host devices Uid delete params
func (o *V1EdgeHostDevicesUIDDeleteParams) WithUID(uid string) *V1EdgeHostDevicesUIDDeleteParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 edge host devices Uid delete params
func (o *V1EdgeHostDevicesUIDDeleteParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1EdgeHostDevicesUIDDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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