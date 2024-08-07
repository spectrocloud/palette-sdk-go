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
	"github.com/go-openapi/swag"
)

// NewV1AuthOrgSwitchParams creates a new V1AuthOrgSwitchParams object
// with the default values initialized.
func NewV1AuthOrgSwitchParams() *V1AuthOrgSwitchParams {
	var (
		setCookieDefault = bool(true)
	)
	return &V1AuthOrgSwitchParams{
		SetCookie: &setCookieDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewV1AuthOrgSwitchParamsWithTimeout creates a new V1AuthOrgSwitchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1AuthOrgSwitchParamsWithTimeout(timeout time.Duration) *V1AuthOrgSwitchParams {
	var (
		setCookieDefault = bool(true)
	)
	return &V1AuthOrgSwitchParams{
		SetCookie: &setCookieDefault,

		timeout: timeout,
	}
}

// NewV1AuthOrgSwitchParamsWithContext creates a new V1AuthOrgSwitchParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1AuthOrgSwitchParamsWithContext(ctx context.Context) *V1AuthOrgSwitchParams {
	var (
		setCookieDefault = bool(true)
	)
	return &V1AuthOrgSwitchParams{
		SetCookie: &setCookieDefault,

		Context: ctx,
	}
}

// NewV1AuthOrgSwitchParamsWithHTTPClient creates a new V1AuthOrgSwitchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1AuthOrgSwitchParamsWithHTTPClient(client *http.Client) *V1AuthOrgSwitchParams {
	var (
		setCookieDefault = bool(true)
	)
	return &V1AuthOrgSwitchParams{
		SetCookie:  &setCookieDefault,
		HTTPClient: client,
	}
}

/*V1AuthOrgSwitchParams contains all the parameters to send to the API endpoint
for the v1 auth org switch operation typically these are written to a http.Request
*/
type V1AuthOrgSwitchParams struct {

	/*OrgName
	  Organization name for which switch request has to be created

	*/
	OrgName string
	/*SetCookie
	  Describes a way to set cookie from backend for switched organization

	*/
	SetCookie *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 auth org switch params
func (o *V1AuthOrgSwitchParams) WithTimeout(timeout time.Duration) *V1AuthOrgSwitchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 auth org switch params
func (o *V1AuthOrgSwitchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 auth org switch params
func (o *V1AuthOrgSwitchParams) WithContext(ctx context.Context) *V1AuthOrgSwitchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 auth org switch params
func (o *V1AuthOrgSwitchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 auth org switch params
func (o *V1AuthOrgSwitchParams) WithHTTPClient(client *http.Client) *V1AuthOrgSwitchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 auth org switch params
func (o *V1AuthOrgSwitchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrgName adds the orgName to the v1 auth org switch params
func (o *V1AuthOrgSwitchParams) WithOrgName(orgName string) *V1AuthOrgSwitchParams {
	o.SetOrgName(orgName)
	return o
}

// SetOrgName adds the orgName to the v1 auth org switch params
func (o *V1AuthOrgSwitchParams) SetOrgName(orgName string) {
	o.OrgName = orgName
}

// WithSetCookie adds the setCookie to the v1 auth org switch params
func (o *V1AuthOrgSwitchParams) WithSetCookie(setCookie *bool) *V1AuthOrgSwitchParams {
	o.SetSetCookie(setCookie)
	return o
}

// SetSetCookie adds the setCookie to the v1 auth org switch params
func (o *V1AuthOrgSwitchParams) SetSetCookie(setCookie *bool) {
	o.SetCookie = setCookie
}

// WriteToRequest writes these params to a swagger request
func (o *V1AuthOrgSwitchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param orgName
	if err := r.SetPathParam("orgName", o.OrgName); err != nil {
		return err
	}

	if o.SetCookie != nil {

		// query param setCookie
		var qrSetCookie bool
		if o.SetCookie != nil {
			qrSetCookie = *o.SetCookie
		}
		qSetCookie := swag.FormatBool(qrSetCookie)
		if qSetCookie != "" {
			if err := r.SetQueryParam("setCookie", qSetCookie); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
