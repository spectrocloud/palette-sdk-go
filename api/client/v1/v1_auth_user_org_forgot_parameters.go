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

// NewV1AuthUserOrgForgotParams creates a new V1AuthUserOrgForgotParams object
// with the default values initialized.
func NewV1AuthUserOrgForgotParams() *V1AuthUserOrgForgotParams {
	var ()
	return &V1AuthUserOrgForgotParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1AuthUserOrgForgotParamsWithTimeout creates a new V1AuthUserOrgForgotParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1AuthUserOrgForgotParamsWithTimeout(timeout time.Duration) *V1AuthUserOrgForgotParams {
	var ()
	return &V1AuthUserOrgForgotParams{

		timeout: timeout,
	}
}

// NewV1AuthUserOrgForgotParamsWithContext creates a new V1AuthUserOrgForgotParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1AuthUserOrgForgotParamsWithContext(ctx context.Context) *V1AuthUserOrgForgotParams {
	var ()
	return &V1AuthUserOrgForgotParams{

		Context: ctx,
	}
}

// NewV1AuthUserOrgForgotParamsWithHTTPClient creates a new V1AuthUserOrgForgotParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1AuthUserOrgForgotParamsWithHTTPClient(client *http.Client) *V1AuthUserOrgForgotParams {
	var ()
	return &V1AuthUserOrgForgotParams{
		HTTPClient: client,
	}
}

/*
V1AuthUserOrgForgotParams contains all the parameters to send to the API endpoint
for the v1 auth user org forgot operation typically these are written to a http.Request
*/
type V1AuthUserOrgForgotParams struct {

	/*EmailID
	  Describes user's email id for sending organzation(s) details via email.

	*/
	EmailID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 auth user org forgot params
func (o *V1AuthUserOrgForgotParams) WithTimeout(timeout time.Duration) *V1AuthUserOrgForgotParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 auth user org forgot params
func (o *V1AuthUserOrgForgotParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 auth user org forgot params
func (o *V1AuthUserOrgForgotParams) WithContext(ctx context.Context) *V1AuthUserOrgForgotParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 auth user org forgot params
func (o *V1AuthUserOrgForgotParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 auth user org forgot params
func (o *V1AuthUserOrgForgotParams) WithHTTPClient(client *http.Client) *V1AuthUserOrgForgotParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 auth user org forgot params
func (o *V1AuthUserOrgForgotParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEmailID adds the emailID to the v1 auth user org forgot params
func (o *V1AuthUserOrgForgotParams) WithEmailID(emailID string) *V1AuthUserOrgForgotParams {
	o.SetEmailID(emailID)
	return o
}

// SetEmailID adds the emailId to the v1 auth user org forgot params
func (o *V1AuthUserOrgForgotParams) SetEmailID(emailID string) {
	o.EmailID = emailID
}

// WriteToRequest writes these params to a swagger request
func (o *V1AuthUserOrgForgotParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param emailId
	qrEmailID := o.EmailID
	qEmailID := qrEmailID
	if qEmailID != "" {
		if err := r.SetQueryParam("emailId", qEmailID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
