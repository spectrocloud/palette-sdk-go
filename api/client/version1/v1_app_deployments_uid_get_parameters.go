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

// NewV1AppDeploymentsUIDGetParams creates a new V1AppDeploymentsUIDGetParams object
// with the default values initialized.
func NewV1AppDeploymentsUIDGetParams() *V1AppDeploymentsUIDGetParams {
	var ()
	return &V1AppDeploymentsUIDGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1AppDeploymentsUIDGetParamsWithTimeout creates a new V1AppDeploymentsUIDGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1AppDeploymentsUIDGetParamsWithTimeout(timeout time.Duration) *V1AppDeploymentsUIDGetParams {
	var ()
	return &V1AppDeploymentsUIDGetParams{

		timeout: timeout,
	}
}

// NewV1AppDeploymentsUIDGetParamsWithContext creates a new V1AppDeploymentsUIDGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1AppDeploymentsUIDGetParamsWithContext(ctx context.Context) *V1AppDeploymentsUIDGetParams {
	var ()
	return &V1AppDeploymentsUIDGetParams{

		Context: ctx,
	}
}

// NewV1AppDeploymentsUIDGetParamsWithHTTPClient creates a new V1AppDeploymentsUIDGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1AppDeploymentsUIDGetParamsWithHTTPClient(client *http.Client) *V1AppDeploymentsUIDGetParams {
	var ()
	return &V1AppDeploymentsUIDGetParams{
		HTTPClient: client,
	}
}

/*
V1AppDeploymentsUIDGetParams contains all the parameters to send to the API endpoint
for the v1 app deployments Uid get operation typically these are written to a http.Request
*/
type V1AppDeploymentsUIDGetParams struct {

	/*UID
	  Application deployment uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 app deployments Uid get params
func (o *V1AppDeploymentsUIDGetParams) WithTimeout(timeout time.Duration) *V1AppDeploymentsUIDGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 app deployments Uid get params
func (o *V1AppDeploymentsUIDGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 app deployments Uid get params
func (o *V1AppDeploymentsUIDGetParams) WithContext(ctx context.Context) *V1AppDeploymentsUIDGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 app deployments Uid get params
func (o *V1AppDeploymentsUIDGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 app deployments Uid get params
func (o *V1AppDeploymentsUIDGetParams) WithHTTPClient(client *http.Client) *V1AppDeploymentsUIDGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 app deployments Uid get params
func (o *V1AppDeploymentsUIDGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUID adds the uid to the v1 app deployments Uid get params
func (o *V1AppDeploymentsUIDGetParams) WithUID(uid string) *V1AppDeploymentsUIDGetParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 app deployments Uid get params
func (o *V1AppDeploymentsUIDGetParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1AppDeploymentsUIDGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
