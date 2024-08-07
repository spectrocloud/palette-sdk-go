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

// NewV1TeamsUIDResourceRolesUIDDeleteParams creates a new V1TeamsUIDResourceRolesUIDDeleteParams object
// with the default values initialized.
func NewV1TeamsUIDResourceRolesUIDDeleteParams() *V1TeamsUIDResourceRolesUIDDeleteParams {
	var ()
	return &V1TeamsUIDResourceRolesUIDDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1TeamsUIDResourceRolesUIDDeleteParamsWithTimeout creates a new V1TeamsUIDResourceRolesUIDDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1TeamsUIDResourceRolesUIDDeleteParamsWithTimeout(timeout time.Duration) *V1TeamsUIDResourceRolesUIDDeleteParams {
	var ()
	return &V1TeamsUIDResourceRolesUIDDeleteParams{

		timeout: timeout,
	}
}

// NewV1TeamsUIDResourceRolesUIDDeleteParamsWithContext creates a new V1TeamsUIDResourceRolesUIDDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1TeamsUIDResourceRolesUIDDeleteParamsWithContext(ctx context.Context) *V1TeamsUIDResourceRolesUIDDeleteParams {
	var ()
	return &V1TeamsUIDResourceRolesUIDDeleteParams{

		Context: ctx,
	}
}

// NewV1TeamsUIDResourceRolesUIDDeleteParamsWithHTTPClient creates a new V1TeamsUIDResourceRolesUIDDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1TeamsUIDResourceRolesUIDDeleteParamsWithHTTPClient(client *http.Client) *V1TeamsUIDResourceRolesUIDDeleteParams {
	var ()
	return &V1TeamsUIDResourceRolesUIDDeleteParams{
		HTTPClient: client,
	}
}

/*
V1TeamsUIDResourceRolesUIDDeleteParams contains all the parameters to send to the API endpoint
for the v1 teams Uid resource roles Uid delete operation typically these are written to a http.Request
*/
type V1TeamsUIDResourceRolesUIDDeleteParams struct {

	/*ResourceRoleUID*/
	ResourceRoleUID string
	/*UID*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 teams Uid resource roles Uid delete params
func (o *V1TeamsUIDResourceRolesUIDDeleteParams) WithTimeout(timeout time.Duration) *V1TeamsUIDResourceRolesUIDDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 teams Uid resource roles Uid delete params
func (o *V1TeamsUIDResourceRolesUIDDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 teams Uid resource roles Uid delete params
func (o *V1TeamsUIDResourceRolesUIDDeleteParams) WithContext(ctx context.Context) *V1TeamsUIDResourceRolesUIDDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 teams Uid resource roles Uid delete params
func (o *V1TeamsUIDResourceRolesUIDDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 teams Uid resource roles Uid delete params
func (o *V1TeamsUIDResourceRolesUIDDeleteParams) WithHTTPClient(client *http.Client) *V1TeamsUIDResourceRolesUIDDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 teams Uid resource roles Uid delete params
func (o *V1TeamsUIDResourceRolesUIDDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithResourceRoleUID adds the resourceRoleUID to the v1 teams Uid resource roles Uid delete params
func (o *V1TeamsUIDResourceRolesUIDDeleteParams) WithResourceRoleUID(resourceRoleUID string) *V1TeamsUIDResourceRolesUIDDeleteParams {
	o.SetResourceRoleUID(resourceRoleUID)
	return o
}

// SetResourceRoleUID adds the resourceRoleUid to the v1 teams Uid resource roles Uid delete params
func (o *V1TeamsUIDResourceRolesUIDDeleteParams) SetResourceRoleUID(resourceRoleUID string) {
	o.ResourceRoleUID = resourceRoleUID
}

// WithUID adds the uid to the v1 teams Uid resource roles Uid delete params
func (o *V1TeamsUIDResourceRolesUIDDeleteParams) WithUID(uid string) *V1TeamsUIDResourceRolesUIDDeleteParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 teams Uid resource roles Uid delete params
func (o *V1TeamsUIDResourceRolesUIDDeleteParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1TeamsUIDResourceRolesUIDDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
