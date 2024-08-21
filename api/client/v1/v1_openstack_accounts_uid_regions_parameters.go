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

// NewV1OpenstackAccountsUIDRegionsParams creates a new V1OpenstackAccountsUIDRegionsParams object
// with the default values initialized.
func NewV1OpenstackAccountsUIDRegionsParams() *V1OpenstackAccountsUIDRegionsParams {
	var ()
	return &V1OpenstackAccountsUIDRegionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1OpenstackAccountsUIDRegionsParamsWithTimeout creates a new V1OpenstackAccountsUIDRegionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1OpenstackAccountsUIDRegionsParamsWithTimeout(timeout time.Duration) *V1OpenstackAccountsUIDRegionsParams {
	var ()
	return &V1OpenstackAccountsUIDRegionsParams{

		timeout: timeout,
	}
}

// NewV1OpenstackAccountsUIDRegionsParamsWithContext creates a new V1OpenstackAccountsUIDRegionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1OpenstackAccountsUIDRegionsParamsWithContext(ctx context.Context) *V1OpenstackAccountsUIDRegionsParams {
	var ()
	return &V1OpenstackAccountsUIDRegionsParams{

		Context: ctx,
	}
}

// NewV1OpenstackAccountsUIDRegionsParamsWithHTTPClient creates a new V1OpenstackAccountsUIDRegionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1OpenstackAccountsUIDRegionsParamsWithHTTPClient(client *http.Client) *V1OpenstackAccountsUIDRegionsParams {
	var ()
	return &V1OpenstackAccountsUIDRegionsParams{
		HTTPClient: client,
	}
}

/*
V1OpenstackAccountsUIDRegionsParams contains all the parameters to send to the API endpoint
for the v1 openstack accounts Uid regions operation typically these are written to a http.Request
*/
type V1OpenstackAccountsUIDRegionsParams struct {

	/*UID*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 openstack accounts Uid regions params
func (o *V1OpenstackAccountsUIDRegionsParams) WithTimeout(timeout time.Duration) *V1OpenstackAccountsUIDRegionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 openstack accounts Uid regions params
func (o *V1OpenstackAccountsUIDRegionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 openstack accounts Uid regions params
func (o *V1OpenstackAccountsUIDRegionsParams) WithContext(ctx context.Context) *V1OpenstackAccountsUIDRegionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 openstack accounts Uid regions params
func (o *V1OpenstackAccountsUIDRegionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 openstack accounts Uid regions params
func (o *V1OpenstackAccountsUIDRegionsParams) WithHTTPClient(client *http.Client) *V1OpenstackAccountsUIDRegionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 openstack accounts Uid regions params
func (o *V1OpenstackAccountsUIDRegionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUID adds the uid to the v1 openstack accounts Uid regions params
func (o *V1OpenstackAccountsUIDRegionsParams) WithUID(uid string) *V1OpenstackAccountsUIDRegionsParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 openstack accounts Uid regions params
func (o *V1OpenstackAccountsUIDRegionsParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1OpenstackAccountsUIDRegionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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