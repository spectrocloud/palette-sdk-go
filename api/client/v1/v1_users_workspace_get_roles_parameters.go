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

// NewV1UsersWorkspaceGetRolesParams creates a new V1UsersWorkspaceGetRolesParams object
// with the default values initialized.
func NewV1UsersWorkspaceGetRolesParams() *V1UsersWorkspaceGetRolesParams {
	var ()
	return &V1UsersWorkspaceGetRolesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1UsersWorkspaceGetRolesParamsWithTimeout creates a new V1UsersWorkspaceGetRolesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1UsersWorkspaceGetRolesParamsWithTimeout(timeout time.Duration) *V1UsersWorkspaceGetRolesParams {
	var ()
	return &V1UsersWorkspaceGetRolesParams{

		timeout: timeout,
	}
}

// NewV1UsersWorkspaceGetRolesParamsWithContext creates a new V1UsersWorkspaceGetRolesParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1UsersWorkspaceGetRolesParamsWithContext(ctx context.Context) *V1UsersWorkspaceGetRolesParams {
	var ()
	return &V1UsersWorkspaceGetRolesParams{

		Context: ctx,
	}
}

// NewV1UsersWorkspaceGetRolesParamsWithHTTPClient creates a new V1UsersWorkspaceGetRolesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1UsersWorkspaceGetRolesParamsWithHTTPClient(client *http.Client) *V1UsersWorkspaceGetRolesParams {
	var ()
	return &V1UsersWorkspaceGetRolesParams{
		HTTPClient: client,
	}
}

/*
V1UsersWorkspaceGetRolesParams contains all the parameters to send to the API endpoint
for the v1 users workspace get roles operation typically these are written to a http.Request
*/
type V1UsersWorkspaceGetRolesParams struct {

	/*UserUID*/
	UserUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 users workspace get roles params
func (o *V1UsersWorkspaceGetRolesParams) WithTimeout(timeout time.Duration) *V1UsersWorkspaceGetRolesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 users workspace get roles params
func (o *V1UsersWorkspaceGetRolesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 users workspace get roles params
func (o *V1UsersWorkspaceGetRolesParams) WithContext(ctx context.Context) *V1UsersWorkspaceGetRolesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 users workspace get roles params
func (o *V1UsersWorkspaceGetRolesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 users workspace get roles params
func (o *V1UsersWorkspaceGetRolesParams) WithHTTPClient(client *http.Client) *V1UsersWorkspaceGetRolesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 users workspace get roles params
func (o *V1UsersWorkspaceGetRolesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserUID adds the userUID to the v1 users workspace get roles params
func (o *V1UsersWorkspaceGetRolesParams) WithUserUID(userUID string) *V1UsersWorkspaceGetRolesParams {
	o.SetUserUID(userUID)
	return o
}

// SetUserUID adds the userUid to the v1 users workspace get roles params
func (o *V1UsersWorkspaceGetRolesParams) SetUserUID(userUID string) {
	o.UserUID = userUID
}

// WriteToRequest writes these params to a swagger request
func (o *V1UsersWorkspaceGetRolesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param userUid
	if err := r.SetPathParam("userUid", o.UserUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}