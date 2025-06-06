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

// NewV1UsersProjectRolesParams creates a new V1UsersProjectRolesParams object
// with the default values initialized.
func NewV1UsersProjectRolesParams() *V1UsersProjectRolesParams {
	var ()
	return &V1UsersProjectRolesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1UsersProjectRolesParamsWithTimeout creates a new V1UsersProjectRolesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1UsersProjectRolesParamsWithTimeout(timeout time.Duration) *V1UsersProjectRolesParams {
	var ()
	return &V1UsersProjectRolesParams{

		timeout: timeout,
	}
}

// NewV1UsersProjectRolesParamsWithContext creates a new V1UsersProjectRolesParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1UsersProjectRolesParamsWithContext(ctx context.Context) *V1UsersProjectRolesParams {
	var ()
	return &V1UsersProjectRolesParams{

		Context: ctx,
	}
}

// NewV1UsersProjectRolesParamsWithHTTPClient creates a new V1UsersProjectRolesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1UsersProjectRolesParamsWithHTTPClient(client *http.Client) *V1UsersProjectRolesParams {
	var ()
	return &V1UsersProjectRolesParams{
		HTTPClient: client,
	}
}

/*
V1UsersProjectRolesParams contains all the parameters to send to the API endpoint
for the v1 users project roles operation typically these are written to a http.Request
*/
type V1UsersProjectRolesParams struct {

	/*UID*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 users project roles params
func (o *V1UsersProjectRolesParams) WithTimeout(timeout time.Duration) *V1UsersProjectRolesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 users project roles params
func (o *V1UsersProjectRolesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 users project roles params
func (o *V1UsersProjectRolesParams) WithContext(ctx context.Context) *V1UsersProjectRolesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 users project roles params
func (o *V1UsersProjectRolesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 users project roles params
func (o *V1UsersProjectRolesParams) WithHTTPClient(client *http.Client) *V1UsersProjectRolesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 users project roles params
func (o *V1UsersProjectRolesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUID adds the uid to the v1 users project roles params
func (o *V1UsersProjectRolesParams) WithUID(uid string) *V1UsersProjectRolesParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 users project roles params
func (o *V1UsersProjectRolesParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1UsersProjectRolesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
