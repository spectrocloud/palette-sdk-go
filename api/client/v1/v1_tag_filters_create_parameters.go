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

// NewV1TagFiltersCreateParams creates a new V1TagFiltersCreateParams object
// with the default values initialized.
func NewV1TagFiltersCreateParams() *V1TagFiltersCreateParams {
	var ()
	return &V1TagFiltersCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1TagFiltersCreateParamsWithTimeout creates a new V1TagFiltersCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1TagFiltersCreateParamsWithTimeout(timeout time.Duration) *V1TagFiltersCreateParams {
	var ()
	return &V1TagFiltersCreateParams{

		timeout: timeout,
	}
}

// NewV1TagFiltersCreateParamsWithContext creates a new V1TagFiltersCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1TagFiltersCreateParamsWithContext(ctx context.Context) *V1TagFiltersCreateParams {
	var ()
	return &V1TagFiltersCreateParams{

		Context: ctx,
	}
}

// NewV1TagFiltersCreateParamsWithHTTPClient creates a new V1TagFiltersCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1TagFiltersCreateParamsWithHTTPClient(client *http.Client) *V1TagFiltersCreateParams {
	var ()
	return &V1TagFiltersCreateParams{
		HTTPClient: client,
	}
}

/*
V1TagFiltersCreateParams contains all the parameters to send to the API endpoint
for the v1 tag filters create operation typically these are written to a http.Request
*/
type V1TagFiltersCreateParams struct {

	/*Body*/
	Body *models.V1TagFilter

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 tag filters create params
func (o *V1TagFiltersCreateParams) WithTimeout(timeout time.Duration) *V1TagFiltersCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 tag filters create params
func (o *V1TagFiltersCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 tag filters create params
func (o *V1TagFiltersCreateParams) WithContext(ctx context.Context) *V1TagFiltersCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 tag filters create params
func (o *V1TagFiltersCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 tag filters create params
func (o *V1TagFiltersCreateParams) WithHTTPClient(client *http.Client) *V1TagFiltersCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 tag filters create params
func (o *V1TagFiltersCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 tag filters create params
func (o *V1TagFiltersCreateParams) WithBody(body *models.V1TagFilter) *V1TagFiltersCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 tag filters create params
func (o *V1TagFiltersCreateParams) SetBody(body *models.V1TagFilter) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *V1TagFiltersCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
