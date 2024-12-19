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

// NewV1EventsComponentsCreateParams creates a new V1EventsComponentsCreateParams object
// with the default values initialized.
func NewV1EventsComponentsCreateParams() *V1EventsComponentsCreateParams {
	var ()
	return &V1EventsComponentsCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1EventsComponentsCreateParamsWithTimeout creates a new V1EventsComponentsCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1EventsComponentsCreateParamsWithTimeout(timeout time.Duration) *V1EventsComponentsCreateParams {
	var ()
	return &V1EventsComponentsCreateParams{

		timeout: timeout,
	}
}

// NewV1EventsComponentsCreateParamsWithContext creates a new V1EventsComponentsCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1EventsComponentsCreateParamsWithContext(ctx context.Context) *V1EventsComponentsCreateParams {
	var ()
	return &V1EventsComponentsCreateParams{

		Context: ctx,
	}
}

// NewV1EventsComponentsCreateParamsWithHTTPClient creates a new V1EventsComponentsCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1EventsComponentsCreateParamsWithHTTPClient(client *http.Client) *V1EventsComponentsCreateParams {
	var ()
	return &V1EventsComponentsCreateParams{
		HTTPClient: client,
	}
}

/*
V1EventsComponentsCreateParams contains all the parameters to send to the API endpoint
for the v1 events components create operation typically these are written to a http.Request
*/
type V1EventsComponentsCreateParams struct {

	/*Body*/
	Body *models.V1Event

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 events components create params
func (o *V1EventsComponentsCreateParams) WithTimeout(timeout time.Duration) *V1EventsComponentsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 events components create params
func (o *V1EventsComponentsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 events components create params
func (o *V1EventsComponentsCreateParams) WithContext(ctx context.Context) *V1EventsComponentsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 events components create params
func (o *V1EventsComponentsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 events components create params
func (o *V1EventsComponentsCreateParams) WithHTTPClient(client *http.Client) *V1EventsComponentsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 events components create params
func (o *V1EventsComponentsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 events components create params
func (o *V1EventsComponentsCreateParams) WithBody(body *models.V1Event) *V1EventsComponentsCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 events components create params
func (o *V1EventsComponentsCreateParams) SetBody(body *models.V1Event) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *V1EventsComponentsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
