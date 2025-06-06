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

// NewV1OverlordsUIDPoolDeleteParams creates a new V1OverlordsUIDPoolDeleteParams object
// with the default values initialized.
func NewV1OverlordsUIDPoolDeleteParams() *V1OverlordsUIDPoolDeleteParams {
	var ()
	return &V1OverlordsUIDPoolDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1OverlordsUIDPoolDeleteParamsWithTimeout creates a new V1OverlordsUIDPoolDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1OverlordsUIDPoolDeleteParamsWithTimeout(timeout time.Duration) *V1OverlordsUIDPoolDeleteParams {
	var ()
	return &V1OverlordsUIDPoolDeleteParams{

		timeout: timeout,
	}
}

// NewV1OverlordsUIDPoolDeleteParamsWithContext creates a new V1OverlordsUIDPoolDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1OverlordsUIDPoolDeleteParamsWithContext(ctx context.Context) *V1OverlordsUIDPoolDeleteParams {
	var ()
	return &V1OverlordsUIDPoolDeleteParams{

		Context: ctx,
	}
}

// NewV1OverlordsUIDPoolDeleteParamsWithHTTPClient creates a new V1OverlordsUIDPoolDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1OverlordsUIDPoolDeleteParamsWithHTTPClient(client *http.Client) *V1OverlordsUIDPoolDeleteParams {
	var ()
	return &V1OverlordsUIDPoolDeleteParams{
		HTTPClient: client,
	}
}

/*
V1OverlordsUIDPoolDeleteParams contains all the parameters to send to the API endpoint
for the v1 overlords Uid pool delete operation typically these are written to a http.Request
*/
type V1OverlordsUIDPoolDeleteParams struct {

	/*PoolUID*/
	PoolUID string
	/*UID*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 overlords Uid pool delete params
func (o *V1OverlordsUIDPoolDeleteParams) WithTimeout(timeout time.Duration) *V1OverlordsUIDPoolDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 overlords Uid pool delete params
func (o *V1OverlordsUIDPoolDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 overlords Uid pool delete params
func (o *V1OverlordsUIDPoolDeleteParams) WithContext(ctx context.Context) *V1OverlordsUIDPoolDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 overlords Uid pool delete params
func (o *V1OverlordsUIDPoolDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 overlords Uid pool delete params
func (o *V1OverlordsUIDPoolDeleteParams) WithHTTPClient(client *http.Client) *V1OverlordsUIDPoolDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 overlords Uid pool delete params
func (o *V1OverlordsUIDPoolDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPoolUID adds the poolUID to the v1 overlords Uid pool delete params
func (o *V1OverlordsUIDPoolDeleteParams) WithPoolUID(poolUID string) *V1OverlordsUIDPoolDeleteParams {
	o.SetPoolUID(poolUID)
	return o
}

// SetPoolUID adds the poolUid to the v1 overlords Uid pool delete params
func (o *V1OverlordsUIDPoolDeleteParams) SetPoolUID(poolUID string) {
	o.PoolUID = poolUID
}

// WithUID adds the uid to the v1 overlords Uid pool delete params
func (o *V1OverlordsUIDPoolDeleteParams) WithUID(uid string) *V1OverlordsUIDPoolDeleteParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 overlords Uid pool delete params
func (o *V1OverlordsUIDPoolDeleteParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1OverlordsUIDPoolDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param poolUid
	if err := r.SetPathParam("poolUid", o.PoolUID); err != nil {
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
