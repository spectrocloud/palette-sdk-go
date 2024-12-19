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

// NewV1PasswordResetParams creates a new V1PasswordResetParams object
// with the default values initialized.
func NewV1PasswordResetParams() *V1PasswordResetParams {
	var ()
	return &V1PasswordResetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1PasswordResetParamsWithTimeout creates a new V1PasswordResetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1PasswordResetParamsWithTimeout(timeout time.Duration) *V1PasswordResetParams {
	var ()
	return &V1PasswordResetParams{

		timeout: timeout,
	}
}

// NewV1PasswordResetParamsWithContext creates a new V1PasswordResetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1PasswordResetParamsWithContext(ctx context.Context) *V1PasswordResetParams {
	var ()
	return &V1PasswordResetParams{

		Context: ctx,
	}
}

// NewV1PasswordResetParamsWithHTTPClient creates a new V1PasswordResetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1PasswordResetParamsWithHTTPClient(client *http.Client) *V1PasswordResetParams {
	var ()
	return &V1PasswordResetParams{
		HTTPClient: client,
	}
}

/*
V1PasswordResetParams contains all the parameters to send to the API endpoint
for the v1 password reset operation typically these are written to a http.Request
*/
type V1PasswordResetParams struct {

	/*Body*/
	Body V1PasswordResetBody
	/*PasswordToken
	  Describes the expirable password token for the user to be used for authentication of user

	*/
	PasswordToken string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 password reset params
func (o *V1PasswordResetParams) WithTimeout(timeout time.Duration) *V1PasswordResetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 password reset params
func (o *V1PasswordResetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 password reset params
func (o *V1PasswordResetParams) WithContext(ctx context.Context) *V1PasswordResetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 password reset params
func (o *V1PasswordResetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 password reset params
func (o *V1PasswordResetParams) WithHTTPClient(client *http.Client) *V1PasswordResetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 password reset params
func (o *V1PasswordResetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 password reset params
func (o *V1PasswordResetParams) WithBody(body V1PasswordResetBody) *V1PasswordResetParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 password reset params
func (o *V1PasswordResetParams) SetBody(body V1PasswordResetBody) {
	o.Body = body
}

// WithPasswordToken adds the passwordToken to the v1 password reset params
func (o *V1PasswordResetParams) WithPasswordToken(passwordToken string) *V1PasswordResetParams {
	o.SetPasswordToken(passwordToken)
	return o
}

// SetPasswordToken adds the passwordToken to the v1 password reset params
func (o *V1PasswordResetParams) SetPasswordToken(passwordToken string) {
	o.PasswordToken = passwordToken
}

// WriteToRequest writes these params to a swagger request
func (o *V1PasswordResetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param passwordToken
	if err := r.SetPathParam("passwordToken", o.PasswordToken); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
