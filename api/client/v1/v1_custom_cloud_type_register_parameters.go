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

// NewV1CustomCloudTypeRegisterParams creates a new V1CustomCloudTypeRegisterParams object
// with the default values initialized.
func NewV1CustomCloudTypeRegisterParams() *V1CustomCloudTypeRegisterParams {
	var ()
	return &V1CustomCloudTypeRegisterParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CustomCloudTypeRegisterParamsWithTimeout creates a new V1CustomCloudTypeRegisterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CustomCloudTypeRegisterParamsWithTimeout(timeout time.Duration) *V1CustomCloudTypeRegisterParams {
	var ()
	return &V1CustomCloudTypeRegisterParams{

		timeout: timeout,
	}
}

// NewV1CustomCloudTypeRegisterParamsWithContext creates a new V1CustomCloudTypeRegisterParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CustomCloudTypeRegisterParamsWithContext(ctx context.Context) *V1CustomCloudTypeRegisterParams {
	var ()
	return &V1CustomCloudTypeRegisterParams{

		Context: ctx,
	}
}

// NewV1CustomCloudTypeRegisterParamsWithHTTPClient creates a new V1CustomCloudTypeRegisterParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CustomCloudTypeRegisterParamsWithHTTPClient(client *http.Client) *V1CustomCloudTypeRegisterParams {
	var ()
	return &V1CustomCloudTypeRegisterParams{
		HTTPClient: client,
	}
}

/*
V1CustomCloudTypeRegisterParams contains all the parameters to send to the API endpoint
for the v1 custom cloud type register operation typically these are written to a http.Request
*/
type V1CustomCloudTypeRegisterParams struct {

	/*Body
	  Request payload to register custom cloud type

	*/
	Body *models.V1CustomCloudRequestEntity

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 custom cloud type register params
func (o *V1CustomCloudTypeRegisterParams) WithTimeout(timeout time.Duration) *V1CustomCloudTypeRegisterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 custom cloud type register params
func (o *V1CustomCloudTypeRegisterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 custom cloud type register params
func (o *V1CustomCloudTypeRegisterParams) WithContext(ctx context.Context) *V1CustomCloudTypeRegisterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 custom cloud type register params
func (o *V1CustomCloudTypeRegisterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 custom cloud type register params
func (o *V1CustomCloudTypeRegisterParams) WithHTTPClient(client *http.Client) *V1CustomCloudTypeRegisterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 custom cloud type register params
func (o *V1CustomCloudTypeRegisterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 custom cloud type register params
func (o *V1CustomCloudTypeRegisterParams) WithBody(body *models.V1CustomCloudRequestEntity) *V1CustomCloudTypeRegisterParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 custom cloud type register params
func (o *V1CustomCloudTypeRegisterParams) SetBody(body *models.V1CustomCloudRequestEntity) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *V1CustomCloudTypeRegisterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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