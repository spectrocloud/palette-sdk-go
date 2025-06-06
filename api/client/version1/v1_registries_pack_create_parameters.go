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
	"github.com/go-openapi/swag"

	"github.com/spectrocloud/palette-sdk-go/api/models"
)

// NewV1RegistriesPackCreateParams creates a new V1RegistriesPackCreateParams object
// with the default values initialized.
func NewV1RegistriesPackCreateParams() *V1RegistriesPackCreateParams {
	var (
		scopeDefault        = string("all")
		skipPackSyncDefault = bool(false)
	)
	return &V1RegistriesPackCreateParams{
		Scope:        &scopeDefault,
		SkipPackSync: &skipPackSyncDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewV1RegistriesPackCreateParamsWithTimeout creates a new V1RegistriesPackCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1RegistriesPackCreateParamsWithTimeout(timeout time.Duration) *V1RegistriesPackCreateParams {
	var (
		scopeDefault        = string("all")
		skipPackSyncDefault = bool(false)
	)
	return &V1RegistriesPackCreateParams{
		Scope:        &scopeDefault,
		SkipPackSync: &skipPackSyncDefault,

		timeout: timeout,
	}
}

// NewV1RegistriesPackCreateParamsWithContext creates a new V1RegistriesPackCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1RegistriesPackCreateParamsWithContext(ctx context.Context) *V1RegistriesPackCreateParams {
	var (
		scopeDefault        = string("all")
		skipPackSyncDefault = bool(false)
	)
	return &V1RegistriesPackCreateParams{
		Scope:        &scopeDefault,
		SkipPackSync: &skipPackSyncDefault,

		Context: ctx,
	}
}

// NewV1RegistriesPackCreateParamsWithHTTPClient creates a new V1RegistriesPackCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1RegistriesPackCreateParamsWithHTTPClient(client *http.Client) *V1RegistriesPackCreateParams {
	var (
		scopeDefault        = string("all")
		skipPackSyncDefault = bool(false)
	)
	return &V1RegistriesPackCreateParams{
		Scope:        &scopeDefault,
		SkipPackSync: &skipPackSyncDefault,
		HTTPClient:   client,
	}
}

/*
V1RegistriesPackCreateParams contains all the parameters to send to the API endpoint
for the v1 registries pack create operation typically these are written to a http.Request
*/
type V1RegistriesPackCreateParams struct {

	/*Body*/
	Body *models.V1PackRegistry
	/*Scope*/
	Scope *string
	/*SkipPackSync*/
	SkipPackSync *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 registries pack create params
func (o *V1RegistriesPackCreateParams) WithTimeout(timeout time.Duration) *V1RegistriesPackCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 registries pack create params
func (o *V1RegistriesPackCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 registries pack create params
func (o *V1RegistriesPackCreateParams) WithContext(ctx context.Context) *V1RegistriesPackCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 registries pack create params
func (o *V1RegistriesPackCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 registries pack create params
func (o *V1RegistriesPackCreateParams) WithHTTPClient(client *http.Client) *V1RegistriesPackCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 registries pack create params
func (o *V1RegistriesPackCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 registries pack create params
func (o *V1RegistriesPackCreateParams) WithBody(body *models.V1PackRegistry) *V1RegistriesPackCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 registries pack create params
func (o *V1RegistriesPackCreateParams) SetBody(body *models.V1PackRegistry) {
	o.Body = body
}

// WithScope adds the scope to the v1 registries pack create params
func (o *V1RegistriesPackCreateParams) WithScope(scope *string) *V1RegistriesPackCreateParams {
	o.SetScope(scope)
	return o
}

// SetScope adds the scope to the v1 registries pack create params
func (o *V1RegistriesPackCreateParams) SetScope(scope *string) {
	o.Scope = scope
}

// WithSkipPackSync adds the skipPackSync to the v1 registries pack create params
func (o *V1RegistriesPackCreateParams) WithSkipPackSync(skipPackSync *bool) *V1RegistriesPackCreateParams {
	o.SetSkipPackSync(skipPackSync)
	return o
}

// SetSkipPackSync adds the skipPackSync to the v1 registries pack create params
func (o *V1RegistriesPackCreateParams) SetSkipPackSync(skipPackSync *bool) {
	o.SkipPackSync = skipPackSync
}

// WriteToRequest writes these params to a swagger request
func (o *V1RegistriesPackCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

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

	if o.SkipPackSync != nil {

		// query param skipPackSync
		var qrSkipPackSync bool
		if o.SkipPackSync != nil {
			qrSkipPackSync = *o.SkipPackSync
		}
		qSkipPackSync := swag.FormatBool(qrSkipPackSync)
		if qSkipPackSync != "" {
			if err := r.SetQueryParam("skipPackSync", qSkipPackSync); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
