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

// NewV1AppDeploymentsUIDProfileVersionsGetParams creates a new V1AppDeploymentsUIDProfileVersionsGetParams object
// with the default values initialized.
func NewV1AppDeploymentsUIDProfileVersionsGetParams() *V1AppDeploymentsUIDProfileVersionsGetParams {
	var ()
	return &V1AppDeploymentsUIDProfileVersionsGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1AppDeploymentsUIDProfileVersionsGetParamsWithTimeout creates a new V1AppDeploymentsUIDProfileVersionsGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1AppDeploymentsUIDProfileVersionsGetParamsWithTimeout(timeout time.Duration) *V1AppDeploymentsUIDProfileVersionsGetParams {
	var ()
	return &V1AppDeploymentsUIDProfileVersionsGetParams{

		timeout: timeout,
	}
}

// NewV1AppDeploymentsUIDProfileVersionsGetParamsWithContext creates a new V1AppDeploymentsUIDProfileVersionsGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1AppDeploymentsUIDProfileVersionsGetParamsWithContext(ctx context.Context) *V1AppDeploymentsUIDProfileVersionsGetParams {
	var ()
	return &V1AppDeploymentsUIDProfileVersionsGetParams{

		Context: ctx,
	}
}

// NewV1AppDeploymentsUIDProfileVersionsGetParamsWithHTTPClient creates a new V1AppDeploymentsUIDProfileVersionsGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1AppDeploymentsUIDProfileVersionsGetParamsWithHTTPClient(client *http.Client) *V1AppDeploymentsUIDProfileVersionsGetParams {
	var ()
	return &V1AppDeploymentsUIDProfileVersionsGetParams{
		HTTPClient: client,
	}
}

/*
V1AppDeploymentsUIDProfileVersionsGetParams contains all the parameters to send to the API endpoint
for the v1 app deployments Uid profile versions get operation typically these are written to a http.Request
*/
type V1AppDeploymentsUIDProfileVersionsGetParams struct {

	/*UID
	  Application deployment uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 app deployments Uid profile versions get params
func (o *V1AppDeploymentsUIDProfileVersionsGetParams) WithTimeout(timeout time.Duration) *V1AppDeploymentsUIDProfileVersionsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 app deployments Uid profile versions get params
func (o *V1AppDeploymentsUIDProfileVersionsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 app deployments Uid profile versions get params
func (o *V1AppDeploymentsUIDProfileVersionsGetParams) WithContext(ctx context.Context) *V1AppDeploymentsUIDProfileVersionsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 app deployments Uid profile versions get params
func (o *V1AppDeploymentsUIDProfileVersionsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 app deployments Uid profile versions get params
func (o *V1AppDeploymentsUIDProfileVersionsGetParams) WithHTTPClient(client *http.Client) *V1AppDeploymentsUIDProfileVersionsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 app deployments Uid profile versions get params
func (o *V1AppDeploymentsUIDProfileVersionsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUID adds the uid to the v1 app deployments Uid profile versions get params
func (o *V1AppDeploymentsUIDProfileVersionsGetParams) WithUID(uid string) *V1AppDeploymentsUIDProfileVersionsGetParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 app deployments Uid profile versions get params
func (o *V1AppDeploymentsUIDProfileVersionsGetParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1AppDeploymentsUIDProfileVersionsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
