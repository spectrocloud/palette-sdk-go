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

// NewV1SpectroClustersUIDAssetsGetParams creates a new V1SpectroClustersUIDAssetsGetParams object
// with the default values initialized.
func NewV1SpectroClustersUIDAssetsGetParams() *V1SpectroClustersUIDAssetsGetParams {
	var ()
	return &V1SpectroClustersUIDAssetsGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1SpectroClustersUIDAssetsGetParamsWithTimeout creates a new V1SpectroClustersUIDAssetsGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1SpectroClustersUIDAssetsGetParamsWithTimeout(timeout time.Duration) *V1SpectroClustersUIDAssetsGetParams {
	var ()
	return &V1SpectroClustersUIDAssetsGetParams{

		timeout: timeout,
	}
}

// NewV1SpectroClustersUIDAssetsGetParamsWithContext creates a new V1SpectroClustersUIDAssetsGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1SpectroClustersUIDAssetsGetParamsWithContext(ctx context.Context) *V1SpectroClustersUIDAssetsGetParams {
	var ()
	return &V1SpectroClustersUIDAssetsGetParams{

		Context: ctx,
	}
}

// NewV1SpectroClustersUIDAssetsGetParamsWithHTTPClient creates a new V1SpectroClustersUIDAssetsGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1SpectroClustersUIDAssetsGetParamsWithHTTPClient(client *http.Client) *V1SpectroClustersUIDAssetsGetParams {
	var ()
	return &V1SpectroClustersUIDAssetsGetParams{
		HTTPClient: client,
	}
}

/*
V1SpectroClustersUIDAssetsGetParams contains all the parameters to send to the API endpoint
for the v1 spectro clusters Uid assets get operation typically these are written to a http.Request
*/
type V1SpectroClustersUIDAssetsGetParams struct {

	/*UID*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 spectro clusters Uid assets get params
func (o *V1SpectroClustersUIDAssetsGetParams) WithTimeout(timeout time.Duration) *V1SpectroClustersUIDAssetsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 spectro clusters Uid assets get params
func (o *V1SpectroClustersUIDAssetsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 spectro clusters Uid assets get params
func (o *V1SpectroClustersUIDAssetsGetParams) WithContext(ctx context.Context) *V1SpectroClustersUIDAssetsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 spectro clusters Uid assets get params
func (o *V1SpectroClustersUIDAssetsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 spectro clusters Uid assets get params
func (o *V1SpectroClustersUIDAssetsGetParams) WithHTTPClient(client *http.Client) *V1SpectroClustersUIDAssetsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 spectro clusters Uid assets get params
func (o *V1SpectroClustersUIDAssetsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUID adds the uid to the v1 spectro clusters Uid assets get params
func (o *V1SpectroClustersUIDAssetsGetParams) WithUID(uid string) *V1SpectroClustersUIDAssetsGetParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 spectro clusters Uid assets get params
func (o *V1SpectroClustersUIDAssetsGetParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1SpectroClustersUIDAssetsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
