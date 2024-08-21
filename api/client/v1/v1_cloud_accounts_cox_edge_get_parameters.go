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

// NewV1CloudAccountsCoxEdgeGetParams creates a new V1CloudAccountsCoxEdgeGetParams object
// with the default values initialized.
func NewV1CloudAccountsCoxEdgeGetParams() *V1CloudAccountsCoxEdgeGetParams {
	var ()
	return &V1CloudAccountsCoxEdgeGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudAccountsCoxEdgeGetParamsWithTimeout creates a new V1CloudAccountsCoxEdgeGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudAccountsCoxEdgeGetParamsWithTimeout(timeout time.Duration) *V1CloudAccountsCoxEdgeGetParams {
	var ()
	return &V1CloudAccountsCoxEdgeGetParams{

		timeout: timeout,
	}
}

// NewV1CloudAccountsCoxEdgeGetParamsWithContext creates a new V1CloudAccountsCoxEdgeGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudAccountsCoxEdgeGetParamsWithContext(ctx context.Context) *V1CloudAccountsCoxEdgeGetParams {
	var ()
	return &V1CloudAccountsCoxEdgeGetParams{

		Context: ctx,
	}
}

// NewV1CloudAccountsCoxEdgeGetParamsWithHTTPClient creates a new V1CloudAccountsCoxEdgeGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudAccountsCoxEdgeGetParamsWithHTTPClient(client *http.Client) *V1CloudAccountsCoxEdgeGetParams {
	var ()
	return &V1CloudAccountsCoxEdgeGetParams{
		HTTPClient: client,
	}
}

/*
V1CloudAccountsCoxEdgeGetParams contains all the parameters to send to the API endpoint
for the v1 cloud accounts cox edge get operation typically these are written to a http.Request
*/
type V1CloudAccountsCoxEdgeGetParams struct {

	/*UID
	  CoxEdge cloud account uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cloud accounts cox edge get params
func (o *V1CloudAccountsCoxEdgeGetParams) WithTimeout(timeout time.Duration) *V1CloudAccountsCoxEdgeGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud accounts cox edge get params
func (o *V1CloudAccountsCoxEdgeGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud accounts cox edge get params
func (o *V1CloudAccountsCoxEdgeGetParams) WithContext(ctx context.Context) *V1CloudAccountsCoxEdgeGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud accounts cox edge get params
func (o *V1CloudAccountsCoxEdgeGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud accounts cox edge get params
func (o *V1CloudAccountsCoxEdgeGetParams) WithHTTPClient(client *http.Client) *V1CloudAccountsCoxEdgeGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud accounts cox edge get params
func (o *V1CloudAccountsCoxEdgeGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUID adds the uid to the v1 cloud accounts cox edge get params
func (o *V1CloudAccountsCoxEdgeGetParams) WithUID(uid string) *V1CloudAccountsCoxEdgeGetParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 cloud accounts cox edge get params
func (o *V1CloudAccountsCoxEdgeGetParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudAccountsCoxEdgeGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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