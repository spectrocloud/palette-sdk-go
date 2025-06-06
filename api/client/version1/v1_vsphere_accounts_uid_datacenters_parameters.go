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

// NewV1VsphereAccountsUIDDatacentersParams creates a new V1VsphereAccountsUIDDatacentersParams object
// with the default values initialized.
func NewV1VsphereAccountsUIDDatacentersParams() *V1VsphereAccountsUIDDatacentersParams {
	var ()
	return &V1VsphereAccountsUIDDatacentersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1VsphereAccountsUIDDatacentersParamsWithTimeout creates a new V1VsphereAccountsUIDDatacentersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1VsphereAccountsUIDDatacentersParamsWithTimeout(timeout time.Duration) *V1VsphereAccountsUIDDatacentersParams {
	var ()
	return &V1VsphereAccountsUIDDatacentersParams{

		timeout: timeout,
	}
}

// NewV1VsphereAccountsUIDDatacentersParamsWithContext creates a new V1VsphereAccountsUIDDatacentersParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1VsphereAccountsUIDDatacentersParamsWithContext(ctx context.Context) *V1VsphereAccountsUIDDatacentersParams {
	var ()
	return &V1VsphereAccountsUIDDatacentersParams{

		Context: ctx,
	}
}

// NewV1VsphereAccountsUIDDatacentersParamsWithHTTPClient creates a new V1VsphereAccountsUIDDatacentersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1VsphereAccountsUIDDatacentersParamsWithHTTPClient(client *http.Client) *V1VsphereAccountsUIDDatacentersParams {
	var ()
	return &V1VsphereAccountsUIDDatacentersParams{
		HTTPClient: client,
	}
}

/*
V1VsphereAccountsUIDDatacentersParams contains all the parameters to send to the API endpoint
for the v1 vsphere accounts Uid datacenters operation typically these are written to a http.Request
*/
type V1VsphereAccountsUIDDatacentersParams struct {

	/*UID*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 vsphere accounts Uid datacenters params
func (o *V1VsphereAccountsUIDDatacentersParams) WithTimeout(timeout time.Duration) *V1VsphereAccountsUIDDatacentersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 vsphere accounts Uid datacenters params
func (o *V1VsphereAccountsUIDDatacentersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 vsphere accounts Uid datacenters params
func (o *V1VsphereAccountsUIDDatacentersParams) WithContext(ctx context.Context) *V1VsphereAccountsUIDDatacentersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 vsphere accounts Uid datacenters params
func (o *V1VsphereAccountsUIDDatacentersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 vsphere accounts Uid datacenters params
func (o *V1VsphereAccountsUIDDatacentersParams) WithHTTPClient(client *http.Client) *V1VsphereAccountsUIDDatacentersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 vsphere accounts Uid datacenters params
func (o *V1VsphereAccountsUIDDatacentersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUID adds the uid to the v1 vsphere accounts Uid datacenters params
func (o *V1VsphereAccountsUIDDatacentersParams) WithUID(uid string) *V1VsphereAccountsUIDDatacentersParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 vsphere accounts Uid datacenters params
func (o *V1VsphereAccountsUIDDatacentersParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1VsphereAccountsUIDDatacentersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
