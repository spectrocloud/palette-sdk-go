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

// NewV1SamlLogoutParams creates a new V1SamlLogoutParams object
// with the default values initialized.
func NewV1SamlLogoutParams() *V1SamlLogoutParams {
	var ()
	return &V1SamlLogoutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1SamlLogoutParamsWithTimeout creates a new V1SamlLogoutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1SamlLogoutParamsWithTimeout(timeout time.Duration) *V1SamlLogoutParams {
	var ()
	return &V1SamlLogoutParams{

		timeout: timeout,
	}
}

// NewV1SamlLogoutParamsWithContext creates a new V1SamlLogoutParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1SamlLogoutParamsWithContext(ctx context.Context) *V1SamlLogoutParams {
	var ()
	return &V1SamlLogoutParams{

		Context: ctx,
	}
}

// NewV1SamlLogoutParamsWithHTTPClient creates a new V1SamlLogoutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1SamlLogoutParamsWithHTTPClient(client *http.Client) *V1SamlLogoutParams {
	var ()
	return &V1SamlLogoutParams{
		HTTPClient: client,
	}
}

/*V1SamlLogoutParams contains all the parameters to send to the API endpoint
for the v1 saml logout operation typically these are written to a http.Request
*/
type V1SamlLogoutParams struct {

	/*SAMLResponse
	  Describe the SAML compliant response sent by IDP which will be validated by Palette to perform logout.

	*/
	SAMLResponse *string
	/*AuthToken
	  Deprecated.

	*/
	AuthToken *string
	/*Org
	  Organization name

	*/
	Org string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 saml logout params
func (o *V1SamlLogoutParams) WithTimeout(timeout time.Duration) *V1SamlLogoutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 saml logout params
func (o *V1SamlLogoutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 saml logout params
func (o *V1SamlLogoutParams) WithContext(ctx context.Context) *V1SamlLogoutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 saml logout params
func (o *V1SamlLogoutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 saml logout params
func (o *V1SamlLogoutParams) WithHTTPClient(client *http.Client) *V1SamlLogoutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 saml logout params
func (o *V1SamlLogoutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSAMLResponse adds the sAMLResponse to the v1 saml logout params
func (o *V1SamlLogoutParams) WithSAMLResponse(sAMLResponse *string) *V1SamlLogoutParams {
	o.SetSAMLResponse(sAMLResponse)
	return o
}

// SetSAMLResponse adds the sAMLResponse to the v1 saml logout params
func (o *V1SamlLogoutParams) SetSAMLResponse(sAMLResponse *string) {
	o.SAMLResponse = sAMLResponse
}

// WithAuthToken adds the authToken to the v1 saml logout params
func (o *V1SamlLogoutParams) WithAuthToken(authToken *string) *V1SamlLogoutParams {
	o.SetAuthToken(authToken)
	return o
}

// SetAuthToken adds the authToken to the v1 saml logout params
func (o *V1SamlLogoutParams) SetAuthToken(authToken *string) {
	o.AuthToken = authToken
}

// WithOrg adds the org to the v1 saml logout params
func (o *V1SamlLogoutParams) WithOrg(org string) *V1SamlLogoutParams {
	o.SetOrg(org)
	return o
}

// SetOrg adds the org to the v1 saml logout params
func (o *V1SamlLogoutParams) SetOrg(org string) {
	o.Org = org
}

// WriteToRequest writes these params to a swagger request
func (o *V1SamlLogoutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.SAMLResponse != nil {

		// form param SAMLResponse
		var frSAMLResponse string
		if o.SAMLResponse != nil {
			frSAMLResponse = *o.SAMLResponse
		}
		fSAMLResponse := frSAMLResponse
		if fSAMLResponse != "" {
			if err := r.SetFormParam("SAMLResponse", fSAMLResponse); err != nil {
				return err
			}
		}

	}

	if o.AuthToken != nil {

		// query param authToken
		var qrAuthToken string
		if o.AuthToken != nil {
			qrAuthToken = *o.AuthToken
		}
		qAuthToken := qrAuthToken
		if qAuthToken != "" {
			if err := r.SetQueryParam("authToken", qAuthToken); err != nil {
				return err
			}
		}

	}

	// path param org
	if err := r.SetPathParam("org", o.Org); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
