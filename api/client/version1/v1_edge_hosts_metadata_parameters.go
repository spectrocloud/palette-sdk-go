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

// NewV1EdgeHostsMetadataParams creates a new V1EdgeHostsMetadataParams object
// with the default values initialized.
func NewV1EdgeHostsMetadataParams() *V1EdgeHostsMetadataParams {
	var ()
	return &V1EdgeHostsMetadataParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1EdgeHostsMetadataParamsWithTimeout creates a new V1EdgeHostsMetadataParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1EdgeHostsMetadataParamsWithTimeout(timeout time.Duration) *V1EdgeHostsMetadataParams {
	var ()
	return &V1EdgeHostsMetadataParams{

		timeout: timeout,
	}
}

// NewV1EdgeHostsMetadataParamsWithContext creates a new V1EdgeHostsMetadataParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1EdgeHostsMetadataParamsWithContext(ctx context.Context) *V1EdgeHostsMetadataParams {
	var ()
	return &V1EdgeHostsMetadataParams{

		Context: ctx,
	}
}

// NewV1EdgeHostsMetadataParamsWithHTTPClient creates a new V1EdgeHostsMetadataParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1EdgeHostsMetadataParamsWithHTTPClient(client *http.Client) *V1EdgeHostsMetadataParams {
	var ()
	return &V1EdgeHostsMetadataParams{
		HTTPClient: client,
	}
}

/*
V1EdgeHostsMetadataParams contains all the parameters to send to the API endpoint
for the v1 edge hosts metadata operation typically these are written to a http.Request
*/
type V1EdgeHostsMetadataParams struct {

	/*Body*/
	Body *models.V1EdgeHostsMetadataFilter

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 edge hosts metadata params
func (o *V1EdgeHostsMetadataParams) WithTimeout(timeout time.Duration) *V1EdgeHostsMetadataParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 edge hosts metadata params
func (o *V1EdgeHostsMetadataParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 edge hosts metadata params
func (o *V1EdgeHostsMetadataParams) WithContext(ctx context.Context) *V1EdgeHostsMetadataParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 edge hosts metadata params
func (o *V1EdgeHostsMetadataParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 edge hosts metadata params
func (o *V1EdgeHostsMetadataParams) WithHTTPClient(client *http.Client) *V1EdgeHostsMetadataParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 edge hosts metadata params
func (o *V1EdgeHostsMetadataParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 edge hosts metadata params
func (o *V1EdgeHostsMetadataParams) WithBody(body *models.V1EdgeHostsMetadataFilter) *V1EdgeHostsMetadataParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 edge hosts metadata params
func (o *V1EdgeHostsMetadataParams) SetBody(body *models.V1EdgeHostsMetadataFilter) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *V1EdgeHostsMetadataParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
