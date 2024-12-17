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

// NewV1ProjectsMetadataParams creates a new V1ProjectsMetadataParams object
// with the default values initialized.
func NewV1ProjectsMetadataParams() *V1ProjectsMetadataParams {
	var ()
	return &V1ProjectsMetadataParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1ProjectsMetadataParamsWithTimeout creates a new V1ProjectsMetadataParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1ProjectsMetadataParamsWithTimeout(timeout time.Duration) *V1ProjectsMetadataParams {
	var ()
	return &V1ProjectsMetadataParams{

		timeout: timeout,
	}
}

// NewV1ProjectsMetadataParamsWithContext creates a new V1ProjectsMetadataParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1ProjectsMetadataParamsWithContext(ctx context.Context) *V1ProjectsMetadataParams {
	var ()
	return &V1ProjectsMetadataParams{

		Context: ctx,
	}
}

// NewV1ProjectsMetadataParamsWithHTTPClient creates a new V1ProjectsMetadataParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1ProjectsMetadataParamsWithHTTPClient(client *http.Client) *V1ProjectsMetadataParams {
	var ()
	return &V1ProjectsMetadataParams{
		HTTPClient: client,
	}
}

/*
V1ProjectsMetadataParams contains all the parameters to send to the API endpoint
for the v1 projects metadata operation typically these are written to a http.Request
*/
type V1ProjectsMetadataParams struct {

	/*Name
	  Name of the project

	*/
	Name *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 projects metadata params
func (o *V1ProjectsMetadataParams) WithTimeout(timeout time.Duration) *V1ProjectsMetadataParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 projects metadata params
func (o *V1ProjectsMetadataParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 projects metadata params
func (o *V1ProjectsMetadataParams) WithContext(ctx context.Context) *V1ProjectsMetadataParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 projects metadata params
func (o *V1ProjectsMetadataParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 projects metadata params
func (o *V1ProjectsMetadataParams) WithHTTPClient(client *http.Client) *V1ProjectsMetadataParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 projects metadata params
func (o *V1ProjectsMetadataParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the v1 projects metadata params
func (o *V1ProjectsMetadataParams) WithName(name *string) *V1ProjectsMetadataParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the v1 projects metadata params
func (o *V1ProjectsMetadataParams) SetName(name *string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *V1ProjectsMetadataParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
