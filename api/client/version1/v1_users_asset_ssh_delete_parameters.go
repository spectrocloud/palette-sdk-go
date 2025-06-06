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

// NewV1UsersAssetSSHDeleteParams creates a new V1UsersAssetSSHDeleteParams object
// with the default values initialized.
func NewV1UsersAssetSSHDeleteParams() *V1UsersAssetSSHDeleteParams {
	var ()
	return &V1UsersAssetSSHDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1UsersAssetSSHDeleteParamsWithTimeout creates a new V1UsersAssetSSHDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1UsersAssetSSHDeleteParamsWithTimeout(timeout time.Duration) *V1UsersAssetSSHDeleteParams {
	var ()
	return &V1UsersAssetSSHDeleteParams{

		timeout: timeout,
	}
}

// NewV1UsersAssetSSHDeleteParamsWithContext creates a new V1UsersAssetSSHDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1UsersAssetSSHDeleteParamsWithContext(ctx context.Context) *V1UsersAssetSSHDeleteParams {
	var ()
	return &V1UsersAssetSSHDeleteParams{

		Context: ctx,
	}
}

// NewV1UsersAssetSSHDeleteParamsWithHTTPClient creates a new V1UsersAssetSSHDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1UsersAssetSSHDeleteParamsWithHTTPClient(client *http.Client) *V1UsersAssetSSHDeleteParams {
	var ()
	return &V1UsersAssetSSHDeleteParams{
		HTTPClient: client,
	}
}

/*
V1UsersAssetSSHDeleteParams contains all the parameters to send to the API endpoint
for the v1 users asset Ssh delete operation typically these are written to a http.Request
*/
type V1UsersAssetSSHDeleteParams struct {

	/*UID
	  Specify the SSH key uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 users asset Ssh delete params
func (o *V1UsersAssetSSHDeleteParams) WithTimeout(timeout time.Duration) *V1UsersAssetSSHDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 users asset Ssh delete params
func (o *V1UsersAssetSSHDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 users asset Ssh delete params
func (o *V1UsersAssetSSHDeleteParams) WithContext(ctx context.Context) *V1UsersAssetSSHDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 users asset Ssh delete params
func (o *V1UsersAssetSSHDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 users asset Ssh delete params
func (o *V1UsersAssetSSHDeleteParams) WithHTTPClient(client *http.Client) *V1UsersAssetSSHDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 users asset Ssh delete params
func (o *V1UsersAssetSSHDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUID adds the uid to the v1 users asset Ssh delete params
func (o *V1UsersAssetSSHDeleteParams) WithUID(uid string) *V1UsersAssetSSHDeleteParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 users asset Ssh delete params
func (o *V1UsersAssetSSHDeleteParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1UsersAssetSSHDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
