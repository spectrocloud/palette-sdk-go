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

// NewV1SpectroClustersUIDTokenKubeConfigDeleteParams creates a new V1SpectroClustersUIDTokenKubeConfigDeleteParams object
// with the default values initialized.
func NewV1SpectroClustersUIDTokenKubeConfigDeleteParams() *V1SpectroClustersUIDTokenKubeConfigDeleteParams {
	var ()
	return &V1SpectroClustersUIDTokenKubeConfigDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1SpectroClustersUIDTokenKubeConfigDeleteParamsWithTimeout creates a new V1SpectroClustersUIDTokenKubeConfigDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1SpectroClustersUIDTokenKubeConfigDeleteParamsWithTimeout(timeout time.Duration) *V1SpectroClustersUIDTokenKubeConfigDeleteParams {
	var ()
	return &V1SpectroClustersUIDTokenKubeConfigDeleteParams{

		timeout: timeout,
	}
}

// NewV1SpectroClustersUIDTokenKubeConfigDeleteParamsWithContext creates a new V1SpectroClustersUIDTokenKubeConfigDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1SpectroClustersUIDTokenKubeConfigDeleteParamsWithContext(ctx context.Context) *V1SpectroClustersUIDTokenKubeConfigDeleteParams {
	var ()
	return &V1SpectroClustersUIDTokenKubeConfigDeleteParams{

		Context: ctx,
	}
}

// NewV1SpectroClustersUIDTokenKubeConfigDeleteParamsWithHTTPClient creates a new V1SpectroClustersUIDTokenKubeConfigDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1SpectroClustersUIDTokenKubeConfigDeleteParamsWithHTTPClient(client *http.Client) *V1SpectroClustersUIDTokenKubeConfigDeleteParams {
	var ()
	return &V1SpectroClustersUIDTokenKubeConfigDeleteParams{
		HTTPClient: client,
	}
}

/*
V1SpectroClustersUIDTokenKubeConfigDeleteParams contains all the parameters to send to the API endpoint
for the v1 spectro clusters Uid token kube config delete operation typically these are written to a http.Request
*/
type V1SpectroClustersUIDTokenKubeConfigDeleteParams struct {

	/*UID
	  Cluster uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 spectro clusters Uid token kube config delete params
func (o *V1SpectroClustersUIDTokenKubeConfigDeleteParams) WithTimeout(timeout time.Duration) *V1SpectroClustersUIDTokenKubeConfigDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 spectro clusters Uid token kube config delete params
func (o *V1SpectroClustersUIDTokenKubeConfigDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 spectro clusters Uid token kube config delete params
func (o *V1SpectroClustersUIDTokenKubeConfigDeleteParams) WithContext(ctx context.Context) *V1SpectroClustersUIDTokenKubeConfigDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 spectro clusters Uid token kube config delete params
func (o *V1SpectroClustersUIDTokenKubeConfigDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 spectro clusters Uid token kube config delete params
func (o *V1SpectroClustersUIDTokenKubeConfigDeleteParams) WithHTTPClient(client *http.Client) *V1SpectroClustersUIDTokenKubeConfigDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 spectro clusters Uid token kube config delete params
func (o *V1SpectroClustersUIDTokenKubeConfigDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUID adds the uid to the v1 spectro clusters Uid token kube config delete params
func (o *V1SpectroClustersUIDTokenKubeConfigDeleteParams) WithUID(uid string) *V1SpectroClustersUIDTokenKubeConfigDeleteParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 spectro clusters Uid token kube config delete params
func (o *V1SpectroClustersUIDTokenKubeConfigDeleteParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1SpectroClustersUIDTokenKubeConfigDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
