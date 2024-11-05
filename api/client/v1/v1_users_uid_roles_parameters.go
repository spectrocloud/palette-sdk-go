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

// NewV1UsersUIDRolesParams creates a new V1UsersUIDRolesParams object
// with the default values initialized.
func NewV1UsersUIDRolesParams() *V1UsersUIDRolesParams {
	var ()
	return &V1UsersUIDRolesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1UsersUIDRolesParamsWithTimeout creates a new V1UsersUIDRolesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1UsersUIDRolesParamsWithTimeout(timeout time.Duration) *V1UsersUIDRolesParams {
	var ()
	return &V1UsersUIDRolesParams{

		timeout: timeout,
	}
}

// NewV1UsersUIDRolesParamsWithContext creates a new V1UsersUIDRolesParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1UsersUIDRolesParamsWithContext(ctx context.Context) *V1UsersUIDRolesParams {
	var ()
	return &V1UsersUIDRolesParams{

		Context: ctx,
	}
}

// NewV1UsersUIDRolesParamsWithHTTPClient creates a new V1UsersUIDRolesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1UsersUIDRolesParamsWithHTTPClient(client *http.Client) *V1UsersUIDRolesParams {
	var ()
	return &V1UsersUIDRolesParams{
		HTTPClient: client,
	}
}

/*
V1UsersUIDRolesParams contains all the parameters to send to the API endpoint
for the v1 users Uid roles operation typically these are written to a http.Request
*/
type V1UsersUIDRolesParams struct {

	/*UID*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 users Uid roles params
func (o *V1UsersUIDRolesParams) WithTimeout(timeout time.Duration) *V1UsersUIDRolesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 users Uid roles params
func (o *V1UsersUIDRolesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 users Uid roles params
func (o *V1UsersUIDRolesParams) WithContext(ctx context.Context) *V1UsersUIDRolesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 users Uid roles params
func (o *V1UsersUIDRolesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 users Uid roles params
func (o *V1UsersUIDRolesParams) WithHTTPClient(client *http.Client) *V1UsersUIDRolesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 users Uid roles params
func (o *V1UsersUIDRolesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUID adds the uid to the v1 users Uid roles params
func (o *V1UsersUIDRolesParams) WithUID(uid string) *V1UsersUIDRolesParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 users Uid roles params
func (o *V1UsersUIDRolesParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1UsersUIDRolesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
