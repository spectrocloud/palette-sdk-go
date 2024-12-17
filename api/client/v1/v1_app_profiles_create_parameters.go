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

// NewV1AppProfilesCreateParams creates a new V1AppProfilesCreateParams object
// with the default values initialized.
func NewV1AppProfilesCreateParams() *V1AppProfilesCreateParams {
	var ()
	return &V1AppProfilesCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1AppProfilesCreateParamsWithTimeout creates a new V1AppProfilesCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1AppProfilesCreateParamsWithTimeout(timeout time.Duration) *V1AppProfilesCreateParams {
	var ()
	return &V1AppProfilesCreateParams{

		timeout: timeout,
	}
}

// NewV1AppProfilesCreateParamsWithContext creates a new V1AppProfilesCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1AppProfilesCreateParamsWithContext(ctx context.Context) *V1AppProfilesCreateParams {
	var ()
	return &V1AppProfilesCreateParams{

		Context: ctx,
	}
}

// NewV1AppProfilesCreateParamsWithHTTPClient creates a new V1AppProfilesCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1AppProfilesCreateParamsWithHTTPClient(client *http.Client) *V1AppProfilesCreateParams {
	var ()
	return &V1AppProfilesCreateParams{
		HTTPClient: client,
	}
}

/*
V1AppProfilesCreateParams contains all the parameters to send to the API endpoint
for the v1 app profiles create operation typically these are written to a http.Request
*/
type V1AppProfilesCreateParams struct {

	/*Body*/
	Body *models.V1AppProfileEntity

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 app profiles create params
func (o *V1AppProfilesCreateParams) WithTimeout(timeout time.Duration) *V1AppProfilesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 app profiles create params
func (o *V1AppProfilesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 app profiles create params
func (o *V1AppProfilesCreateParams) WithContext(ctx context.Context) *V1AppProfilesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 app profiles create params
func (o *V1AppProfilesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 app profiles create params
func (o *V1AppProfilesCreateParams) WithHTTPClient(client *http.Client) *V1AppProfilesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 app profiles create params
func (o *V1AppProfilesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 app profiles create params
func (o *V1AppProfilesCreateParams) WithBody(body *models.V1AppProfileEntity) *V1AppProfilesCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 app profiles create params
func (o *V1AppProfilesCreateParams) SetBody(body *models.V1AppProfileEntity) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *V1AppProfilesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
