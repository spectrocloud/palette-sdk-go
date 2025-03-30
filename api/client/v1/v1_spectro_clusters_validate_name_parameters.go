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

// NewV1SpectroClustersValidateNameParams creates a new V1SpectroClustersValidateNameParams object
// with the default values initialized.
func NewV1SpectroClustersValidateNameParams() *V1SpectroClustersValidateNameParams {
	var ()
	return &V1SpectroClustersValidateNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1SpectroClustersValidateNameParamsWithTimeout creates a new V1SpectroClustersValidateNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1SpectroClustersValidateNameParamsWithTimeout(timeout time.Duration) *V1SpectroClustersValidateNameParams {
	var ()
	return &V1SpectroClustersValidateNameParams{

		timeout: timeout,
	}
}

// NewV1SpectroClustersValidateNameParamsWithContext creates a new V1SpectroClustersValidateNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1SpectroClustersValidateNameParamsWithContext(ctx context.Context) *V1SpectroClustersValidateNameParams {
	var ()
	return &V1SpectroClustersValidateNameParams{

		Context: ctx,
	}
}

// NewV1SpectroClustersValidateNameParamsWithHTTPClient creates a new V1SpectroClustersValidateNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1SpectroClustersValidateNameParamsWithHTTPClient(client *http.Client) *V1SpectroClustersValidateNameParams {
	var ()
	return &V1SpectroClustersValidateNameParams{
		HTTPClient: client,
	}
}

/*
V1SpectroClustersValidateNameParams contains all the parameters to send to the API endpoint
for the v1 spectro clusters validate name operation typically these are written to a http.Request
*/
type V1SpectroClustersValidateNameParams struct {

	/*Name
	  Cluster name

	*/
	Name *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 spectro clusters validate name params
func (o *V1SpectroClustersValidateNameParams) WithTimeout(timeout time.Duration) *V1SpectroClustersValidateNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 spectro clusters validate name params
func (o *V1SpectroClustersValidateNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 spectro clusters validate name params
func (o *V1SpectroClustersValidateNameParams) WithContext(ctx context.Context) *V1SpectroClustersValidateNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 spectro clusters validate name params
func (o *V1SpectroClustersValidateNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 spectro clusters validate name params
func (o *V1SpectroClustersValidateNameParams) WithHTTPClient(client *http.Client) *V1SpectroClustersValidateNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 spectro clusters validate name params
func (o *V1SpectroClustersValidateNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the v1 spectro clusters validate name params
func (o *V1SpectroClustersValidateNameParams) WithName(name *string) *V1SpectroClustersValidateNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the v1 spectro clusters validate name params
func (o *V1SpectroClustersValidateNameParams) SetName(name *string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *V1SpectroClustersValidateNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
