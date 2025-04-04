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

// NewV1SpectroClustersUIDResetParams creates a new V1SpectroClustersUIDResetParams object
// with the default values initialized.
func NewV1SpectroClustersUIDResetParams() *V1SpectroClustersUIDResetParams {
	var ()
	return &V1SpectroClustersUIDResetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1SpectroClustersUIDResetParamsWithTimeout creates a new V1SpectroClustersUIDResetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1SpectroClustersUIDResetParamsWithTimeout(timeout time.Duration) *V1SpectroClustersUIDResetParams {
	var ()
	return &V1SpectroClustersUIDResetParams{

		timeout: timeout,
	}
}

// NewV1SpectroClustersUIDResetParamsWithContext creates a new V1SpectroClustersUIDResetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1SpectroClustersUIDResetParamsWithContext(ctx context.Context) *V1SpectroClustersUIDResetParams {
	var ()
	return &V1SpectroClustersUIDResetParams{

		Context: ctx,
	}
}

// NewV1SpectroClustersUIDResetParamsWithHTTPClient creates a new V1SpectroClustersUIDResetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1SpectroClustersUIDResetParamsWithHTTPClient(client *http.Client) *V1SpectroClustersUIDResetParams {
	var ()
	return &V1SpectroClustersUIDResetParams{
		HTTPClient: client,
	}
}

/*
V1SpectroClustersUIDResetParams contains all the parameters to send to the API endpoint
for the v1 spectro clusters Uid reset operation typically these are written to a http.Request
*/
type V1SpectroClustersUIDResetParams struct {

	/*UID
	  Cluster uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 spectro clusters Uid reset params
func (o *V1SpectroClustersUIDResetParams) WithTimeout(timeout time.Duration) *V1SpectroClustersUIDResetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 spectro clusters Uid reset params
func (o *V1SpectroClustersUIDResetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 spectro clusters Uid reset params
func (o *V1SpectroClustersUIDResetParams) WithContext(ctx context.Context) *V1SpectroClustersUIDResetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 spectro clusters Uid reset params
func (o *V1SpectroClustersUIDResetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 spectro clusters Uid reset params
func (o *V1SpectroClustersUIDResetParams) WithHTTPClient(client *http.Client) *V1SpectroClustersUIDResetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 spectro clusters Uid reset params
func (o *V1SpectroClustersUIDResetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUID adds the uid to the v1 spectro clusters Uid reset params
func (o *V1SpectroClustersUIDResetParams) WithUID(uid string) *V1SpectroClustersUIDResetParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 spectro clusters Uid reset params
func (o *V1SpectroClustersUIDResetParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1SpectroClustersUIDResetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
