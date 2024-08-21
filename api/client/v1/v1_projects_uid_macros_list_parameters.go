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

// NewV1ProjectsUIDMacrosListParams creates a new V1ProjectsUIDMacrosListParams object
// with the default values initialized.
func NewV1ProjectsUIDMacrosListParams() *V1ProjectsUIDMacrosListParams {
	var ()
	return &V1ProjectsUIDMacrosListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1ProjectsUIDMacrosListParamsWithTimeout creates a new V1ProjectsUIDMacrosListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1ProjectsUIDMacrosListParamsWithTimeout(timeout time.Duration) *V1ProjectsUIDMacrosListParams {
	var ()
	return &V1ProjectsUIDMacrosListParams{

		timeout: timeout,
	}
}

// NewV1ProjectsUIDMacrosListParamsWithContext creates a new V1ProjectsUIDMacrosListParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1ProjectsUIDMacrosListParamsWithContext(ctx context.Context) *V1ProjectsUIDMacrosListParams {
	var ()
	return &V1ProjectsUIDMacrosListParams{

		Context: ctx,
	}
}

// NewV1ProjectsUIDMacrosListParamsWithHTTPClient creates a new V1ProjectsUIDMacrosListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1ProjectsUIDMacrosListParamsWithHTTPClient(client *http.Client) *V1ProjectsUIDMacrosListParams {
	var ()
	return &V1ProjectsUIDMacrosListParams{
		HTTPClient: client,
	}
}

/*
V1ProjectsUIDMacrosListParams contains all the parameters to send to the API endpoint
for the v1 projects Uid macros list operation typically these are written to a http.Request
*/
type V1ProjectsUIDMacrosListParams struct {

	/*UID*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 projects Uid macros list params
func (o *V1ProjectsUIDMacrosListParams) WithTimeout(timeout time.Duration) *V1ProjectsUIDMacrosListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 projects Uid macros list params
func (o *V1ProjectsUIDMacrosListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 projects Uid macros list params
func (o *V1ProjectsUIDMacrosListParams) WithContext(ctx context.Context) *V1ProjectsUIDMacrosListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 projects Uid macros list params
func (o *V1ProjectsUIDMacrosListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 projects Uid macros list params
func (o *V1ProjectsUIDMacrosListParams) WithHTTPClient(client *http.Client) *V1ProjectsUIDMacrosListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 projects Uid macros list params
func (o *V1ProjectsUIDMacrosListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUID adds the uid to the v1 projects Uid macros list params
func (o *V1ProjectsUIDMacrosListParams) WithUID(uid string) *V1ProjectsUIDMacrosListParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 projects Uid macros list params
func (o *V1ProjectsUIDMacrosListParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1ProjectsUIDMacrosListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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