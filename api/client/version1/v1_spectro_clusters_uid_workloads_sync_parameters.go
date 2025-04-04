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
)

// NewV1SpectroClustersUIDWorkloadsSyncParams creates a new V1SpectroClustersUIDWorkloadsSyncParams object
// with the default values initialized.
func NewV1SpectroClustersUIDWorkloadsSyncParams() *V1SpectroClustersUIDWorkloadsSyncParams {
	var ()
	return &V1SpectroClustersUIDWorkloadsSyncParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1SpectroClustersUIDWorkloadsSyncParamsWithTimeout creates a new V1SpectroClustersUIDWorkloadsSyncParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1SpectroClustersUIDWorkloadsSyncParamsWithTimeout(timeout time.Duration) *V1SpectroClustersUIDWorkloadsSyncParams {
	var ()
	return &V1SpectroClustersUIDWorkloadsSyncParams{

		timeout: timeout,
	}
}

// NewV1SpectroClustersUIDWorkloadsSyncParamsWithContext creates a new V1SpectroClustersUIDWorkloadsSyncParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1SpectroClustersUIDWorkloadsSyncParamsWithContext(ctx context.Context) *V1SpectroClustersUIDWorkloadsSyncParams {
	var ()
	return &V1SpectroClustersUIDWorkloadsSyncParams{

		Context: ctx,
	}
}

// NewV1SpectroClustersUIDWorkloadsSyncParamsWithHTTPClient creates a new V1SpectroClustersUIDWorkloadsSyncParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1SpectroClustersUIDWorkloadsSyncParamsWithHTTPClient(client *http.Client) *V1SpectroClustersUIDWorkloadsSyncParams {
	var ()
	return &V1SpectroClustersUIDWorkloadsSyncParams{
		HTTPClient: client,
	}
}

/*
V1SpectroClustersUIDWorkloadsSyncParams contains all the parameters to send to the API endpoint
for the v1 spectro clusters Uid workloads sync operation typically these are written to a http.Request
*/
type V1SpectroClustersUIDWorkloadsSyncParams struct {

	/*UID*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 spectro clusters Uid workloads sync params
func (o *V1SpectroClustersUIDWorkloadsSyncParams) WithTimeout(timeout time.Duration) *V1SpectroClustersUIDWorkloadsSyncParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 spectro clusters Uid workloads sync params
func (o *V1SpectroClustersUIDWorkloadsSyncParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 spectro clusters Uid workloads sync params
func (o *V1SpectroClustersUIDWorkloadsSyncParams) WithContext(ctx context.Context) *V1SpectroClustersUIDWorkloadsSyncParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 spectro clusters Uid workloads sync params
func (o *V1SpectroClustersUIDWorkloadsSyncParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 spectro clusters Uid workloads sync params
func (o *V1SpectroClustersUIDWorkloadsSyncParams) WithHTTPClient(client *http.Client) *V1SpectroClustersUIDWorkloadsSyncParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 spectro clusters Uid workloads sync params
func (o *V1SpectroClustersUIDWorkloadsSyncParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUID adds the uid to the v1 spectro clusters Uid workloads sync params
func (o *V1SpectroClustersUIDWorkloadsSyncParams) WithUID(uid string) *V1SpectroClustersUIDWorkloadsSyncParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 spectro clusters Uid workloads sync params
func (o *V1SpectroClustersUIDWorkloadsSyncParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1SpectroClustersUIDWorkloadsSyncParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
