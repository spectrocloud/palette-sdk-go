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

// NewV1SpectroClustersUIDKubeConfigClientDeleteParams creates a new V1SpectroClustersUIDKubeConfigClientDeleteParams object
// with the default values initialized.
func NewV1SpectroClustersUIDKubeConfigClientDeleteParams() *V1SpectroClustersUIDKubeConfigClientDeleteParams {
	var ()
	return &V1SpectroClustersUIDKubeConfigClientDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1SpectroClustersUIDKubeConfigClientDeleteParamsWithTimeout creates a new V1SpectroClustersUIDKubeConfigClientDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1SpectroClustersUIDKubeConfigClientDeleteParamsWithTimeout(timeout time.Duration) *V1SpectroClustersUIDKubeConfigClientDeleteParams {
	var ()
	return &V1SpectroClustersUIDKubeConfigClientDeleteParams{

		timeout: timeout,
	}
}

// NewV1SpectroClustersUIDKubeConfigClientDeleteParamsWithContext creates a new V1SpectroClustersUIDKubeConfigClientDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1SpectroClustersUIDKubeConfigClientDeleteParamsWithContext(ctx context.Context) *V1SpectroClustersUIDKubeConfigClientDeleteParams {
	var ()
	return &V1SpectroClustersUIDKubeConfigClientDeleteParams{

		Context: ctx,
	}
}

// NewV1SpectroClustersUIDKubeConfigClientDeleteParamsWithHTTPClient creates a new V1SpectroClustersUIDKubeConfigClientDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1SpectroClustersUIDKubeConfigClientDeleteParamsWithHTTPClient(client *http.Client) *V1SpectroClustersUIDKubeConfigClientDeleteParams {
	var ()
	return &V1SpectroClustersUIDKubeConfigClientDeleteParams{
		HTTPClient: client,
	}
}

/*
V1SpectroClustersUIDKubeConfigClientDeleteParams contains all the parameters to send to the API endpoint
for the v1 spectro clusters Uid kube config client delete operation typically these are written to a http.Request
*/
type V1SpectroClustersUIDKubeConfigClientDeleteParams struct {

	/*UID
	  Cluster uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 spectro clusters Uid kube config client delete params
func (o *V1SpectroClustersUIDKubeConfigClientDeleteParams) WithTimeout(timeout time.Duration) *V1SpectroClustersUIDKubeConfigClientDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 spectro clusters Uid kube config client delete params
func (o *V1SpectroClustersUIDKubeConfigClientDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 spectro clusters Uid kube config client delete params
func (o *V1SpectroClustersUIDKubeConfigClientDeleteParams) WithContext(ctx context.Context) *V1SpectroClustersUIDKubeConfigClientDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 spectro clusters Uid kube config client delete params
func (o *V1SpectroClustersUIDKubeConfigClientDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 spectro clusters Uid kube config client delete params
func (o *V1SpectroClustersUIDKubeConfigClientDeleteParams) WithHTTPClient(client *http.Client) *V1SpectroClustersUIDKubeConfigClientDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 spectro clusters Uid kube config client delete params
func (o *V1SpectroClustersUIDKubeConfigClientDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUID adds the uid to the v1 spectro clusters Uid kube config client delete params
func (o *V1SpectroClustersUIDKubeConfigClientDeleteParams) WithUID(uid string) *V1SpectroClustersUIDKubeConfigClientDeleteParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 spectro clusters Uid kube config client delete params
func (o *V1SpectroClustersUIDKubeConfigClientDeleteParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1SpectroClustersUIDKubeConfigClientDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
