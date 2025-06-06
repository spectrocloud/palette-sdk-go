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

	"github.com/spectrocloud/palette-sdk-go/api/models"
)

// NewV1RegistriesPackValidateParams creates a new V1RegistriesPackValidateParams object
// with the default values initialized.
func NewV1RegistriesPackValidateParams() *V1RegistriesPackValidateParams {
	var ()
	return &V1RegistriesPackValidateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1RegistriesPackValidateParamsWithTimeout creates a new V1RegistriesPackValidateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1RegistriesPackValidateParamsWithTimeout(timeout time.Duration) *V1RegistriesPackValidateParams {
	var ()
	return &V1RegistriesPackValidateParams{

		timeout: timeout,
	}
}

// NewV1RegistriesPackValidateParamsWithContext creates a new V1RegistriesPackValidateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1RegistriesPackValidateParamsWithContext(ctx context.Context) *V1RegistriesPackValidateParams {
	var ()
	return &V1RegistriesPackValidateParams{

		Context: ctx,
	}
}

// NewV1RegistriesPackValidateParamsWithHTTPClient creates a new V1RegistriesPackValidateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1RegistriesPackValidateParamsWithHTTPClient(client *http.Client) *V1RegistriesPackValidateParams {
	var ()
	return &V1RegistriesPackValidateParams{
		HTTPClient: client,
	}
}

/*
V1RegistriesPackValidateParams contains all the parameters to send to the API endpoint
for the v1 registries pack validate operation typically these are written to a http.Request
*/
type V1RegistriesPackValidateParams struct {

	/*Body*/
	Body *models.V1PackRegistrySpec

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 registries pack validate params
func (o *V1RegistriesPackValidateParams) WithTimeout(timeout time.Duration) *V1RegistriesPackValidateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 registries pack validate params
func (o *V1RegistriesPackValidateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 registries pack validate params
func (o *V1RegistriesPackValidateParams) WithContext(ctx context.Context) *V1RegistriesPackValidateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 registries pack validate params
func (o *V1RegistriesPackValidateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 registries pack validate params
func (o *V1RegistriesPackValidateParams) WithHTTPClient(client *http.Client) *V1RegistriesPackValidateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 registries pack validate params
func (o *V1RegistriesPackValidateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 registries pack validate params
func (o *V1RegistriesPackValidateParams) WithBody(body *models.V1PackRegistrySpec) *V1RegistriesPackValidateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 registries pack validate params
func (o *V1RegistriesPackValidateParams) SetBody(body *models.V1PackRegistrySpec) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *V1RegistriesPackValidateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
