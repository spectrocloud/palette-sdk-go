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

// NewV1UsersUIDResourceRolesUIDDeleteParams creates a new V1UsersUIDResourceRolesUIDDeleteParams object
// with the default values initialized.
func NewV1UsersUIDResourceRolesUIDDeleteParams() *V1UsersUIDResourceRolesUIDDeleteParams {
	var ()
	return &V1UsersUIDResourceRolesUIDDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1UsersUIDResourceRolesUIDDeleteParamsWithTimeout creates a new V1UsersUIDResourceRolesUIDDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1UsersUIDResourceRolesUIDDeleteParamsWithTimeout(timeout time.Duration) *V1UsersUIDResourceRolesUIDDeleteParams {
	var ()
	return &V1UsersUIDResourceRolesUIDDeleteParams{

		timeout: timeout,
	}
}

// NewV1UsersUIDResourceRolesUIDDeleteParamsWithContext creates a new V1UsersUIDResourceRolesUIDDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1UsersUIDResourceRolesUIDDeleteParamsWithContext(ctx context.Context) *V1UsersUIDResourceRolesUIDDeleteParams {
	var ()
	return &V1UsersUIDResourceRolesUIDDeleteParams{

		Context: ctx,
	}
}

// NewV1UsersUIDResourceRolesUIDDeleteParamsWithHTTPClient creates a new V1UsersUIDResourceRolesUIDDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1UsersUIDResourceRolesUIDDeleteParamsWithHTTPClient(client *http.Client) *V1UsersUIDResourceRolesUIDDeleteParams {
	var ()
	return &V1UsersUIDResourceRolesUIDDeleteParams{
		HTTPClient: client,
	}
}

/*
V1UsersUIDResourceRolesUIDDeleteParams contains all the parameters to send to the API endpoint
for the v1 users Uid resource roles Uid delete operation typically these are written to a http.Request
*/
type V1UsersUIDResourceRolesUIDDeleteParams struct {

	/*ResourceRoleUID*/
	ResourceRoleUID string
	/*UID*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 users Uid resource roles Uid delete params
func (o *V1UsersUIDResourceRolesUIDDeleteParams) WithTimeout(timeout time.Duration) *V1UsersUIDResourceRolesUIDDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 users Uid resource roles Uid delete params
func (o *V1UsersUIDResourceRolesUIDDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 users Uid resource roles Uid delete params
func (o *V1UsersUIDResourceRolesUIDDeleteParams) WithContext(ctx context.Context) *V1UsersUIDResourceRolesUIDDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 users Uid resource roles Uid delete params
func (o *V1UsersUIDResourceRolesUIDDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 users Uid resource roles Uid delete params
func (o *V1UsersUIDResourceRolesUIDDeleteParams) WithHTTPClient(client *http.Client) *V1UsersUIDResourceRolesUIDDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 users Uid resource roles Uid delete params
func (o *V1UsersUIDResourceRolesUIDDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithResourceRoleUID adds the resourceRoleUID to the v1 users Uid resource roles Uid delete params
func (o *V1UsersUIDResourceRolesUIDDeleteParams) WithResourceRoleUID(resourceRoleUID string) *V1UsersUIDResourceRolesUIDDeleteParams {
	o.SetResourceRoleUID(resourceRoleUID)
	return o
}

// SetResourceRoleUID adds the resourceRoleUid to the v1 users Uid resource roles Uid delete params
func (o *V1UsersUIDResourceRolesUIDDeleteParams) SetResourceRoleUID(resourceRoleUID string) {
	o.ResourceRoleUID = resourceRoleUID
}

// WithUID adds the uid to the v1 users Uid resource roles Uid delete params
func (o *V1UsersUIDResourceRolesUIDDeleteParams) WithUID(uid string) *V1UsersUIDResourceRolesUIDDeleteParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 users Uid resource roles Uid delete params
func (o *V1UsersUIDResourceRolesUIDDeleteParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1UsersUIDResourceRolesUIDDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param resourceRoleUid
	if err := r.SetPathParam("resourceRoleUid", o.ResourceRoleUID); err != nil {
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
