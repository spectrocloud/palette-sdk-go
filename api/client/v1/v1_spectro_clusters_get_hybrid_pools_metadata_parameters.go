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

// NewV1SpectroClustersGetHybridPoolsMetadataParams creates a new V1SpectroClustersGetHybridPoolsMetadataParams object
// with the default values initialized.
func NewV1SpectroClustersGetHybridPoolsMetadataParams() *V1SpectroClustersGetHybridPoolsMetadataParams {
	var ()
	return &V1SpectroClustersGetHybridPoolsMetadataParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1SpectroClustersGetHybridPoolsMetadataParamsWithTimeout creates a new V1SpectroClustersGetHybridPoolsMetadataParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1SpectroClustersGetHybridPoolsMetadataParamsWithTimeout(timeout time.Duration) *V1SpectroClustersGetHybridPoolsMetadataParams {
	var ()
	return &V1SpectroClustersGetHybridPoolsMetadataParams{

		timeout: timeout,
	}
}

// NewV1SpectroClustersGetHybridPoolsMetadataParamsWithContext creates a new V1SpectroClustersGetHybridPoolsMetadataParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1SpectroClustersGetHybridPoolsMetadataParamsWithContext(ctx context.Context) *V1SpectroClustersGetHybridPoolsMetadataParams {
	var ()
	return &V1SpectroClustersGetHybridPoolsMetadataParams{

		Context: ctx,
	}
}

// NewV1SpectroClustersGetHybridPoolsMetadataParamsWithHTTPClient creates a new V1SpectroClustersGetHybridPoolsMetadataParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1SpectroClustersGetHybridPoolsMetadataParamsWithHTTPClient(client *http.Client) *V1SpectroClustersGetHybridPoolsMetadataParams {
	var ()
	return &V1SpectroClustersGetHybridPoolsMetadataParams{
		HTTPClient: client,
	}
}

/*
V1SpectroClustersGetHybridPoolsMetadataParams contains all the parameters to send to the API endpoint
for the v1 spectro clusters get hybrid pools metadata operation typically these are written to a http.Request
*/
type V1SpectroClustersGetHybridPoolsMetadataParams struct {

	/*UID
	  Cluster uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 spectro clusters get hybrid pools metadata params
func (o *V1SpectroClustersGetHybridPoolsMetadataParams) WithTimeout(timeout time.Duration) *V1SpectroClustersGetHybridPoolsMetadataParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 spectro clusters get hybrid pools metadata params
func (o *V1SpectroClustersGetHybridPoolsMetadataParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 spectro clusters get hybrid pools metadata params
func (o *V1SpectroClustersGetHybridPoolsMetadataParams) WithContext(ctx context.Context) *V1SpectroClustersGetHybridPoolsMetadataParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 spectro clusters get hybrid pools metadata params
func (o *V1SpectroClustersGetHybridPoolsMetadataParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 spectro clusters get hybrid pools metadata params
func (o *V1SpectroClustersGetHybridPoolsMetadataParams) WithHTTPClient(client *http.Client) *V1SpectroClustersGetHybridPoolsMetadataParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 spectro clusters get hybrid pools metadata params
func (o *V1SpectroClustersGetHybridPoolsMetadataParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUID adds the uid to the v1 spectro clusters get hybrid pools metadata params
func (o *V1SpectroClustersGetHybridPoolsMetadataParams) WithUID(uid string) *V1SpectroClustersGetHybridPoolsMetadataParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 spectro clusters get hybrid pools metadata params
func (o *V1SpectroClustersGetHybridPoolsMetadataParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1SpectroClustersGetHybridPoolsMetadataParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param uid
	if err := r.SetPathParam("uid", o.UID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}