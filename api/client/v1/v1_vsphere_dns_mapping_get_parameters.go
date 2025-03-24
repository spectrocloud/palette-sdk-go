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

// NewV1VsphereDNSMappingGetParams creates a new V1VsphereDNSMappingGetParams object
// with the default values initialized.
func NewV1VsphereDNSMappingGetParams() *V1VsphereDNSMappingGetParams {
	var ()
	return &V1VsphereDNSMappingGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1VsphereDNSMappingGetParamsWithTimeout creates a new V1VsphereDNSMappingGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1VsphereDNSMappingGetParamsWithTimeout(timeout time.Duration) *V1VsphereDNSMappingGetParams {
	var ()
	return &V1VsphereDNSMappingGetParams{

		timeout: timeout,
	}
}

// NewV1VsphereDNSMappingGetParamsWithContext creates a new V1VsphereDNSMappingGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1VsphereDNSMappingGetParamsWithContext(ctx context.Context) *V1VsphereDNSMappingGetParams {
	var ()
	return &V1VsphereDNSMappingGetParams{

		Context: ctx,
	}
}

// NewV1VsphereDNSMappingGetParamsWithHTTPClient creates a new V1VsphereDNSMappingGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1VsphereDNSMappingGetParamsWithHTTPClient(client *http.Client) *V1VsphereDNSMappingGetParams {
	var ()
	return &V1VsphereDNSMappingGetParams{
		HTTPClient: client,
	}
}

/*V1VsphereDNSMappingGetParams contains all the parameters to send to the API endpoint
for the v1 vsphere Dns mapping get operation typically these are written to a http.Request
*/
type V1VsphereDNSMappingGetParams struct {

	/*UID
	  Specify the vSphere DNS mapping uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 vsphere Dns mapping get params
func (o *V1VsphereDNSMappingGetParams) WithTimeout(timeout time.Duration) *V1VsphereDNSMappingGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 vsphere Dns mapping get params
func (o *V1VsphereDNSMappingGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 vsphere Dns mapping get params
func (o *V1VsphereDNSMappingGetParams) WithContext(ctx context.Context) *V1VsphereDNSMappingGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 vsphere Dns mapping get params
func (o *V1VsphereDNSMappingGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 vsphere Dns mapping get params
func (o *V1VsphereDNSMappingGetParams) WithHTTPClient(client *http.Client) *V1VsphereDNSMappingGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 vsphere Dns mapping get params
func (o *V1VsphereDNSMappingGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUID adds the uid to the v1 vsphere Dns mapping get params
func (o *V1VsphereDNSMappingGetParams) WithUID(uid string) *V1VsphereDNSMappingGetParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 vsphere Dns mapping get params
func (o *V1VsphereDNSMappingGetParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1VsphereDNSMappingGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
