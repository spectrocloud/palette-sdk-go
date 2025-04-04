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

// NewV1OverlordsUIDVsphereDatacentersParams creates a new V1OverlordsUIDVsphereDatacentersParams object
// with the default values initialized.
func NewV1OverlordsUIDVsphereDatacentersParams() *V1OverlordsUIDVsphereDatacentersParams {
	var ()
	return &V1OverlordsUIDVsphereDatacentersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1OverlordsUIDVsphereDatacentersParamsWithTimeout creates a new V1OverlordsUIDVsphereDatacentersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1OverlordsUIDVsphereDatacentersParamsWithTimeout(timeout time.Duration) *V1OverlordsUIDVsphereDatacentersParams {
	var ()
	return &V1OverlordsUIDVsphereDatacentersParams{

		timeout: timeout,
	}
}

// NewV1OverlordsUIDVsphereDatacentersParamsWithContext creates a new V1OverlordsUIDVsphereDatacentersParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1OverlordsUIDVsphereDatacentersParamsWithContext(ctx context.Context) *V1OverlordsUIDVsphereDatacentersParams {
	var ()
	return &V1OverlordsUIDVsphereDatacentersParams{

		Context: ctx,
	}
}

// NewV1OverlordsUIDVsphereDatacentersParamsWithHTTPClient creates a new V1OverlordsUIDVsphereDatacentersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1OverlordsUIDVsphereDatacentersParamsWithHTTPClient(client *http.Client) *V1OverlordsUIDVsphereDatacentersParams {
	var ()
	return &V1OverlordsUIDVsphereDatacentersParams{
		HTTPClient: client,
	}
}

/*
V1OverlordsUIDVsphereDatacentersParams contains all the parameters to send to the API endpoint
for the v1 overlords Uid vsphere datacenters operation typically these are written to a http.Request
*/
type V1OverlordsUIDVsphereDatacentersParams struct {

	/*UID*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 overlords Uid vsphere datacenters params
func (o *V1OverlordsUIDVsphereDatacentersParams) WithTimeout(timeout time.Duration) *V1OverlordsUIDVsphereDatacentersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 overlords Uid vsphere datacenters params
func (o *V1OverlordsUIDVsphereDatacentersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 overlords Uid vsphere datacenters params
func (o *V1OverlordsUIDVsphereDatacentersParams) WithContext(ctx context.Context) *V1OverlordsUIDVsphereDatacentersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 overlords Uid vsphere datacenters params
func (o *V1OverlordsUIDVsphereDatacentersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 overlords Uid vsphere datacenters params
func (o *V1OverlordsUIDVsphereDatacentersParams) WithHTTPClient(client *http.Client) *V1OverlordsUIDVsphereDatacentersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 overlords Uid vsphere datacenters params
func (o *V1OverlordsUIDVsphereDatacentersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUID adds the uid to the v1 overlords Uid vsphere datacenters params
func (o *V1OverlordsUIDVsphereDatacentersParams) WithUID(uid string) *V1OverlordsUIDVsphereDatacentersParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 overlords Uid vsphere datacenters params
func (o *V1OverlordsUIDVsphereDatacentersParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1OverlordsUIDVsphereDatacentersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
