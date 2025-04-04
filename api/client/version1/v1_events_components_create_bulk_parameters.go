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

// NewV1EventsComponentsCreateBulkParams creates a new V1EventsComponentsCreateBulkParams object
// with the default values initialized.
func NewV1EventsComponentsCreateBulkParams() *V1EventsComponentsCreateBulkParams {
	var ()
	return &V1EventsComponentsCreateBulkParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1EventsComponentsCreateBulkParamsWithTimeout creates a new V1EventsComponentsCreateBulkParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1EventsComponentsCreateBulkParamsWithTimeout(timeout time.Duration) *V1EventsComponentsCreateBulkParams {
	var ()
	return &V1EventsComponentsCreateBulkParams{

		timeout: timeout,
	}
}

// NewV1EventsComponentsCreateBulkParamsWithContext creates a new V1EventsComponentsCreateBulkParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1EventsComponentsCreateBulkParamsWithContext(ctx context.Context) *V1EventsComponentsCreateBulkParams {
	var ()
	return &V1EventsComponentsCreateBulkParams{

		Context: ctx,
	}
}

// NewV1EventsComponentsCreateBulkParamsWithHTTPClient creates a new V1EventsComponentsCreateBulkParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1EventsComponentsCreateBulkParamsWithHTTPClient(client *http.Client) *V1EventsComponentsCreateBulkParams {
	var ()
	return &V1EventsComponentsCreateBulkParams{
		HTTPClient: client,
	}
}

/*
V1EventsComponentsCreateBulkParams contains all the parameters to send to the API endpoint
for the v1 events components create bulk operation typically these are written to a http.Request
*/
type V1EventsComponentsCreateBulkParams struct {

	/*Body*/
	Body models.V1BulkEvents

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 events components create bulk params
func (o *V1EventsComponentsCreateBulkParams) WithTimeout(timeout time.Duration) *V1EventsComponentsCreateBulkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 events components create bulk params
func (o *V1EventsComponentsCreateBulkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 events components create bulk params
func (o *V1EventsComponentsCreateBulkParams) WithContext(ctx context.Context) *V1EventsComponentsCreateBulkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 events components create bulk params
func (o *V1EventsComponentsCreateBulkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 events components create bulk params
func (o *V1EventsComponentsCreateBulkParams) WithHTTPClient(client *http.Client) *V1EventsComponentsCreateBulkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 events components create bulk params
func (o *V1EventsComponentsCreateBulkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 events components create bulk params
func (o *V1EventsComponentsCreateBulkParams) WithBody(body models.V1BulkEvents) *V1EventsComponentsCreateBulkParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 events components create bulk params
func (o *V1EventsComponentsCreateBulkParams) SetBody(body models.V1BulkEvents) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *V1EventsComponentsCreateBulkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
