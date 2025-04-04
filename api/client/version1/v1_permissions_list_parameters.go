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

// NewV1PermissionsListParams creates a new V1PermissionsListParams object
// with the default values initialized.
func NewV1PermissionsListParams() *V1PermissionsListParams {
	var ()
	return &V1PermissionsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1PermissionsListParamsWithTimeout creates a new V1PermissionsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1PermissionsListParamsWithTimeout(timeout time.Duration) *V1PermissionsListParams {
	var ()
	return &V1PermissionsListParams{

		timeout: timeout,
	}
}

// NewV1PermissionsListParamsWithContext creates a new V1PermissionsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1PermissionsListParamsWithContext(ctx context.Context) *V1PermissionsListParams {
	var ()
	return &V1PermissionsListParams{

		Context: ctx,
	}
}

// NewV1PermissionsListParamsWithHTTPClient creates a new V1PermissionsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1PermissionsListParamsWithHTTPClient(client *http.Client) *V1PermissionsListParams {
	var ()
	return &V1PermissionsListParams{
		HTTPClient: client,
	}
}

/*
V1PermissionsListParams contains all the parameters to send to the API endpoint
for the v1 permissions list operation typically these are written to a http.Request
*/
type V1PermissionsListParams struct {

	/*Scope*/
	Scope *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 permissions list params
func (o *V1PermissionsListParams) WithTimeout(timeout time.Duration) *V1PermissionsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 permissions list params
func (o *V1PermissionsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 permissions list params
func (o *V1PermissionsListParams) WithContext(ctx context.Context) *V1PermissionsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 permissions list params
func (o *V1PermissionsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 permissions list params
func (o *V1PermissionsListParams) WithHTTPClient(client *http.Client) *V1PermissionsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 permissions list params
func (o *V1PermissionsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithScope adds the scope to the v1 permissions list params
func (o *V1PermissionsListParams) WithScope(scope *string) *V1PermissionsListParams {
	o.SetScope(scope)
	return o
}

// SetScope adds the scope to the v1 permissions list params
func (o *V1PermissionsListParams) SetScope(scope *string) {
	o.Scope = scope
}

// WriteToRequest writes these params to a swagger request
func (o *V1PermissionsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Scope != nil {

		// query param scope
		var qrScope string
		if o.Scope != nil {
			qrScope = *o.Scope
		}
		qScope := qrScope
		if qScope != "" {
			if err := r.SetQueryParam("scope", qScope); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
