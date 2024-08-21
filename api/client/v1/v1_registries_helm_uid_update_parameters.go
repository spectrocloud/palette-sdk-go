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

	"github.com/spectrocloud/palette-sdk-go/api/models"
)

// NewV1RegistriesHelmUIDUpdateParams creates a new V1RegistriesHelmUIDUpdateParams object
// with the default values initialized.
func NewV1RegistriesHelmUIDUpdateParams() *V1RegistriesHelmUIDUpdateParams {
	var ()
	return &V1RegistriesHelmUIDUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1RegistriesHelmUIDUpdateParamsWithTimeout creates a new V1RegistriesHelmUIDUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1RegistriesHelmUIDUpdateParamsWithTimeout(timeout time.Duration) *V1RegistriesHelmUIDUpdateParams {
	var ()
	return &V1RegistriesHelmUIDUpdateParams{

		timeout: timeout,
	}
}

// NewV1RegistriesHelmUIDUpdateParamsWithContext creates a new V1RegistriesHelmUIDUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1RegistriesHelmUIDUpdateParamsWithContext(ctx context.Context) *V1RegistriesHelmUIDUpdateParams {
	var ()
	return &V1RegistriesHelmUIDUpdateParams{

		Context: ctx,
	}
}

// NewV1RegistriesHelmUIDUpdateParamsWithHTTPClient creates a new V1RegistriesHelmUIDUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1RegistriesHelmUIDUpdateParamsWithHTTPClient(client *http.Client) *V1RegistriesHelmUIDUpdateParams {
	var ()
	return &V1RegistriesHelmUIDUpdateParams{
		HTTPClient: client,
	}
}

/*
V1RegistriesHelmUIDUpdateParams contains all the parameters to send to the API endpoint
for the v1 registries helm Uid update operation typically these are written to a http.Request
*/
type V1RegistriesHelmUIDUpdateParams struct {

	/*Body*/
	Body *models.V1HelmRegistry
	/*UID*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 registries helm Uid update params
func (o *V1RegistriesHelmUIDUpdateParams) WithTimeout(timeout time.Duration) *V1RegistriesHelmUIDUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 registries helm Uid update params
func (o *V1RegistriesHelmUIDUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 registries helm Uid update params
func (o *V1RegistriesHelmUIDUpdateParams) WithContext(ctx context.Context) *V1RegistriesHelmUIDUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 registries helm Uid update params
func (o *V1RegistriesHelmUIDUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 registries helm Uid update params
func (o *V1RegistriesHelmUIDUpdateParams) WithHTTPClient(client *http.Client) *V1RegistriesHelmUIDUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 registries helm Uid update params
func (o *V1RegistriesHelmUIDUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 registries helm Uid update params
func (o *V1RegistriesHelmUIDUpdateParams) WithBody(body *models.V1HelmRegistry) *V1RegistriesHelmUIDUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 registries helm Uid update params
func (o *V1RegistriesHelmUIDUpdateParams) SetBody(body *models.V1HelmRegistry) {
	o.Body = body
}

// WithUID adds the uid to the v1 registries helm Uid update params
func (o *V1RegistriesHelmUIDUpdateParams) WithUID(uid string) *V1RegistriesHelmUIDUpdateParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 registries helm Uid update params
func (o *V1RegistriesHelmUIDUpdateParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1RegistriesHelmUIDUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param uid
	if err := r.SetPathParam("uid", o.UID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}