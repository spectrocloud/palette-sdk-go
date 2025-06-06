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

// NewV1EdgeTokensUIDDeleteParams creates a new V1EdgeTokensUIDDeleteParams object
// with the default values initialized.
func NewV1EdgeTokensUIDDeleteParams() *V1EdgeTokensUIDDeleteParams {
	var ()
	return &V1EdgeTokensUIDDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1EdgeTokensUIDDeleteParamsWithTimeout creates a new V1EdgeTokensUIDDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1EdgeTokensUIDDeleteParamsWithTimeout(timeout time.Duration) *V1EdgeTokensUIDDeleteParams {
	var ()
	return &V1EdgeTokensUIDDeleteParams{

		timeout: timeout,
	}
}

// NewV1EdgeTokensUIDDeleteParamsWithContext creates a new V1EdgeTokensUIDDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1EdgeTokensUIDDeleteParamsWithContext(ctx context.Context) *V1EdgeTokensUIDDeleteParams {
	var ()
	return &V1EdgeTokensUIDDeleteParams{

		Context: ctx,
	}
}

// NewV1EdgeTokensUIDDeleteParamsWithHTTPClient creates a new V1EdgeTokensUIDDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1EdgeTokensUIDDeleteParamsWithHTTPClient(client *http.Client) *V1EdgeTokensUIDDeleteParams {
	var ()
	return &V1EdgeTokensUIDDeleteParams{
		HTTPClient: client,
	}
}

/*
V1EdgeTokensUIDDeleteParams contains all the parameters to send to the API endpoint
for the v1 edge tokens Uid delete operation typically these are written to a http.Request
*/
type V1EdgeTokensUIDDeleteParams struct {

	/*UID
	  Edge token uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 edge tokens Uid delete params
func (o *V1EdgeTokensUIDDeleteParams) WithTimeout(timeout time.Duration) *V1EdgeTokensUIDDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 edge tokens Uid delete params
func (o *V1EdgeTokensUIDDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 edge tokens Uid delete params
func (o *V1EdgeTokensUIDDeleteParams) WithContext(ctx context.Context) *V1EdgeTokensUIDDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 edge tokens Uid delete params
func (o *V1EdgeTokensUIDDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 edge tokens Uid delete params
func (o *V1EdgeTokensUIDDeleteParams) WithHTTPClient(client *http.Client) *V1EdgeTokensUIDDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 edge tokens Uid delete params
func (o *V1EdgeTokensUIDDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUID adds the uid to the v1 edge tokens Uid delete params
func (o *V1EdgeTokensUIDDeleteParams) WithUID(uid string) *V1EdgeTokensUIDDeleteParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 edge tokens Uid delete params
func (o *V1EdgeTokensUIDDeleteParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1EdgeTokensUIDDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
