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

// NewV1OverlordsListParams creates a new V1OverlordsListParams object
// with the default values initialized.
func NewV1OverlordsListParams() *V1OverlordsListParams {
	var ()
	return &V1OverlordsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1OverlordsListParamsWithTimeout creates a new V1OverlordsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1OverlordsListParamsWithTimeout(timeout time.Duration) *V1OverlordsListParams {
	var ()
	return &V1OverlordsListParams{

		timeout: timeout,
	}
}

// NewV1OverlordsListParamsWithContext creates a new V1OverlordsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1OverlordsListParamsWithContext(ctx context.Context) *V1OverlordsListParams {
	var ()
	return &V1OverlordsListParams{

		Context: ctx,
	}
}

// NewV1OverlordsListParamsWithHTTPClient creates a new V1OverlordsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1OverlordsListParamsWithHTTPClient(client *http.Client) *V1OverlordsListParams {
	var ()
	return &V1OverlordsListParams{
		HTTPClient: client,
	}
}

/*V1OverlordsListParams contains all the parameters to send to the API endpoint
for the v1 overlords list operation typically these are written to a http.Request
*/
type V1OverlordsListParams struct {

	/*Name*/
	Name *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 overlords list params
func (o *V1OverlordsListParams) WithTimeout(timeout time.Duration) *V1OverlordsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 overlords list params
func (o *V1OverlordsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 overlords list params
func (o *V1OverlordsListParams) WithContext(ctx context.Context) *V1OverlordsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 overlords list params
func (o *V1OverlordsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 overlords list params
func (o *V1OverlordsListParams) WithHTTPClient(client *http.Client) *V1OverlordsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 overlords list params
func (o *V1OverlordsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the v1 overlords list params
func (o *V1OverlordsListParams) WithName(name *string) *V1OverlordsListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the v1 overlords list params
func (o *V1OverlordsListParams) SetName(name *string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *V1OverlordsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Name != nil {

		// query param name
		var qrName string
		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {
			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
