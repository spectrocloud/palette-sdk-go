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

// NewV1SamlCallbackParams creates a new V1SamlCallbackParams object
// with the default values initialized.
func NewV1SamlCallbackParams() *V1SamlCallbackParams {
	var ()
	return &V1SamlCallbackParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1SamlCallbackParamsWithTimeout creates a new V1SamlCallbackParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1SamlCallbackParamsWithTimeout(timeout time.Duration) *V1SamlCallbackParams {
	var ()
	return &V1SamlCallbackParams{

		timeout: timeout,
	}
}

// NewV1SamlCallbackParamsWithContext creates a new V1SamlCallbackParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1SamlCallbackParamsWithContext(ctx context.Context) *V1SamlCallbackParams {
	var ()
	return &V1SamlCallbackParams{

		Context: ctx,
	}
}

// NewV1SamlCallbackParamsWithHTTPClient creates a new V1SamlCallbackParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1SamlCallbackParamsWithHTTPClient(client *http.Client) *V1SamlCallbackParams {
	var ()
	return &V1SamlCallbackParams{
		HTTPClient: client,
	}
}

/*
V1SamlCallbackParams contains all the parameters to send to the API endpoint
for the v1 saml callback operation typically these are written to a http.Request
*/
type V1SamlCallbackParams struct {

	/*RelayState
	  Describes a state to validate and associate request and response

	*/
	RelayState *string
	/*SAMLResponse
	  Describe the SAML compliant response sent by IDP which will be validated by Palette

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

// WithTimeout adds the timeout to the v1 saml callback params
func (o *V1SamlCallbackParams) WithTimeout(timeout time.Duration) *V1SamlCallbackParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 saml callback params
func (o *V1SamlCallbackParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 saml callback params
func (o *V1SamlCallbackParams) WithContext(ctx context.Context) *V1SamlCallbackParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 saml callback params
func (o *V1SamlCallbackParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 saml callback params
func (o *V1SamlCallbackParams) WithHTTPClient(client *http.Client) *V1SamlCallbackParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 saml callback params
func (o *V1SamlCallbackParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRelayState adds the relayState to the v1 saml callback params
func (o *V1SamlCallbackParams) WithRelayState(relayState *string) *V1SamlCallbackParams {
	o.SetRelayState(relayState)
	return o
}

// SetRelayState adds the relayState to the v1 saml callback params
func (o *V1SamlCallbackParams) SetRelayState(relayState *string) {
	o.RelayState = relayState
}

// WithSAMLResponse adds the sAMLResponse to the v1 saml callback params
func (o *V1SamlCallbackParams) WithSAMLResponse(sAMLResponse *string) *V1SamlCallbackParams {
	o.SetSAMLResponse(sAMLResponse)
	return o
}

// SetSAMLResponse adds the sAMLResponse to the v1 saml callback params
func (o *V1SamlCallbackParams) SetSAMLResponse(sAMLResponse *string) {
	o.SAMLResponse = sAMLResponse
}

// WithAuthToken adds the authToken to the v1 saml callback params
func (o *V1SamlCallbackParams) WithAuthToken(authToken *string) *V1SamlCallbackParams {
	o.SetAuthToken(authToken)
	return o
}

// SetAuthToken adds the authToken to the v1 saml callback params
func (o *V1SamlCallbackParams) SetAuthToken(authToken *string) {
	o.AuthToken = authToken
}

// WithOrg adds the org to the v1 saml callback params
func (o *V1SamlCallbackParams) WithOrg(org string) *V1SamlCallbackParams {
	o.SetOrg(org)
	return o
}

// SetOrg adds the org to the v1 saml callback params
func (o *V1SamlCallbackParams) SetOrg(org string) {
	o.Org = org
}

// WriteToRequest writes these params to a swagger request
func (o *V1SamlCallbackParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.RelayState != nil {

		// form param RelayState
		var frRelayState string
		if o.RelayState != nil {
			frRelayState = *o.RelayState
		}
		fRelayState := frRelayState
		if fRelayState != "" {
			if err := r.SetFormParam("RelayState", fRelayState); err != nil {
				return err
			}
		}

	}

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
