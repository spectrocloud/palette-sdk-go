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

// NewV1OverlordsUIDGetParams creates a new V1OverlordsUIDGetParams object
// with the default values initialized.
func NewV1OverlordsUIDGetParams() *V1OverlordsUIDGetParams {
	var ()
	return &V1OverlordsUIDGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1OverlordsUIDGetParamsWithTimeout creates a new V1OverlordsUIDGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1OverlordsUIDGetParamsWithTimeout(timeout time.Duration) *V1OverlordsUIDGetParams {
	var ()
	return &V1OverlordsUIDGetParams{

		timeout: timeout,
	}
}

// NewV1OverlordsUIDGetParamsWithContext creates a new V1OverlordsUIDGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1OverlordsUIDGetParamsWithContext(ctx context.Context) *V1OverlordsUIDGetParams {
	var ()
	return &V1OverlordsUIDGetParams{

		Context: ctx,
	}
}

// NewV1OverlordsUIDGetParamsWithHTTPClient creates a new V1OverlordsUIDGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1OverlordsUIDGetParamsWithHTTPClient(client *http.Client) *V1OverlordsUIDGetParams {
	var ()
	return &V1OverlordsUIDGetParams{
		HTTPClient: client,
	}
}

/*
V1OverlordsUIDGetParams contains all the parameters to send to the API endpoint
for the v1 overlords Uid get operation typically these are written to a http.Request
*/
type V1OverlordsUIDGetParams struct {

	/*UID*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 overlords Uid get params
func (o *V1OverlordsUIDGetParams) WithTimeout(timeout time.Duration) *V1OverlordsUIDGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 overlords Uid get params
func (o *V1OverlordsUIDGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 overlords Uid get params
func (o *V1OverlordsUIDGetParams) WithContext(ctx context.Context) *V1OverlordsUIDGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 overlords Uid get params
func (o *V1OverlordsUIDGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 overlords Uid get params
func (o *V1OverlordsUIDGetParams) WithHTTPClient(client *http.Client) *V1OverlordsUIDGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 overlords Uid get params
func (o *V1OverlordsUIDGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUID adds the uid to the v1 overlords Uid get params
func (o *V1OverlordsUIDGetParams) WithUID(uid string) *V1OverlordsUIDGetParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 overlords Uid get params
func (o *V1OverlordsUIDGetParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1OverlordsUIDGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
