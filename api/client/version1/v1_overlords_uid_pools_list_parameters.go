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

// NewV1OverlordsUIDPoolsListParams creates a new V1OverlordsUIDPoolsListParams object
// with the default values initialized.
func NewV1OverlordsUIDPoolsListParams() *V1OverlordsUIDPoolsListParams {
	var ()
	return &V1OverlordsUIDPoolsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1OverlordsUIDPoolsListParamsWithTimeout creates a new V1OverlordsUIDPoolsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1OverlordsUIDPoolsListParamsWithTimeout(timeout time.Duration) *V1OverlordsUIDPoolsListParams {
	var ()
	return &V1OverlordsUIDPoolsListParams{

		timeout: timeout,
	}
}

// NewV1OverlordsUIDPoolsListParamsWithContext creates a new V1OverlordsUIDPoolsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1OverlordsUIDPoolsListParamsWithContext(ctx context.Context) *V1OverlordsUIDPoolsListParams {
	var ()
	return &V1OverlordsUIDPoolsListParams{

		Context: ctx,
	}
}

// NewV1OverlordsUIDPoolsListParamsWithHTTPClient creates a new V1OverlordsUIDPoolsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1OverlordsUIDPoolsListParamsWithHTTPClient(client *http.Client) *V1OverlordsUIDPoolsListParams {
	var ()
	return &V1OverlordsUIDPoolsListParams{
		HTTPClient: client,
	}
}

/*
V1OverlordsUIDPoolsListParams contains all the parameters to send to the API endpoint
for the v1 overlords Uid pools list operation typically these are written to a http.Request
*/
type V1OverlordsUIDPoolsListParams struct {

	/*UID*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 overlords Uid pools list params
func (o *V1OverlordsUIDPoolsListParams) WithTimeout(timeout time.Duration) *V1OverlordsUIDPoolsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 overlords Uid pools list params
func (o *V1OverlordsUIDPoolsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 overlords Uid pools list params
func (o *V1OverlordsUIDPoolsListParams) WithContext(ctx context.Context) *V1OverlordsUIDPoolsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 overlords Uid pools list params
func (o *V1OverlordsUIDPoolsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 overlords Uid pools list params
func (o *V1OverlordsUIDPoolsListParams) WithHTTPClient(client *http.Client) *V1OverlordsUIDPoolsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 overlords Uid pools list params
func (o *V1OverlordsUIDPoolsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUID adds the uid to the v1 overlords Uid pools list params
func (o *V1OverlordsUIDPoolsListParams) WithUID(uid string) *V1OverlordsUIDPoolsListParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 overlords Uid pools list params
func (o *V1OverlordsUIDPoolsListParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1OverlordsUIDPoolsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
