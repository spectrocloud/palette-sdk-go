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
	"github.com/go-openapi/swag"
)

// NewV1SpectroClustersUIDKubeConfigParams creates a new V1SpectroClustersUIDKubeConfigParams object
// with the default values initialized.
func NewV1SpectroClustersUIDKubeConfigParams() *V1SpectroClustersUIDKubeConfigParams {
	var (
		frpDefault = bool(true)
	)
	return &V1SpectroClustersUIDKubeConfigParams{
		Frp: &frpDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewV1SpectroClustersUIDKubeConfigParamsWithTimeout creates a new V1SpectroClustersUIDKubeConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1SpectroClustersUIDKubeConfigParamsWithTimeout(timeout time.Duration) *V1SpectroClustersUIDKubeConfigParams {
	var (
		frpDefault = bool(true)
	)
	return &V1SpectroClustersUIDKubeConfigParams{
		Frp: &frpDefault,

		timeout: timeout,
	}
}

// NewV1SpectroClustersUIDKubeConfigParamsWithContext creates a new V1SpectroClustersUIDKubeConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1SpectroClustersUIDKubeConfigParamsWithContext(ctx context.Context) *V1SpectroClustersUIDKubeConfigParams {
	var (
		frpDefault = bool(true)
	)
	return &V1SpectroClustersUIDKubeConfigParams{
		Frp: &frpDefault,

		Context: ctx,
	}
}

// NewV1SpectroClustersUIDKubeConfigParamsWithHTTPClient creates a new V1SpectroClustersUIDKubeConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1SpectroClustersUIDKubeConfigParamsWithHTTPClient(client *http.Client) *V1SpectroClustersUIDKubeConfigParams {
	var (
		frpDefault = bool(true)
	)
	return &V1SpectroClustersUIDKubeConfigParams{
		Frp:        &frpDefault,
		HTTPClient: client,
	}
}

/*
V1SpectroClustersUIDKubeConfigParams contains all the parameters to send to the API endpoint
for the v1 spectro clusters Uid kube config operation typically these are written to a http.Request
*/
type V1SpectroClustersUIDKubeConfigParams struct {

	/*Frp
	  FRP (reverse-proxy) based kube config will be returned if available

	*/
	Frp *bool
	/*UID
	  Cluster uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 spectro clusters Uid kube config params
func (o *V1SpectroClustersUIDKubeConfigParams) WithTimeout(timeout time.Duration) *V1SpectroClustersUIDKubeConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 spectro clusters Uid kube config params
func (o *V1SpectroClustersUIDKubeConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 spectro clusters Uid kube config params
func (o *V1SpectroClustersUIDKubeConfigParams) WithContext(ctx context.Context) *V1SpectroClustersUIDKubeConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 spectro clusters Uid kube config params
func (o *V1SpectroClustersUIDKubeConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 spectro clusters Uid kube config params
func (o *V1SpectroClustersUIDKubeConfigParams) WithHTTPClient(client *http.Client) *V1SpectroClustersUIDKubeConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 spectro clusters Uid kube config params
func (o *V1SpectroClustersUIDKubeConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFrp adds the frp to the v1 spectro clusters Uid kube config params
func (o *V1SpectroClustersUIDKubeConfigParams) WithFrp(frp *bool) *V1SpectroClustersUIDKubeConfigParams {
	o.SetFrp(frp)
	return o
}

// SetFrp adds the frp to the v1 spectro clusters Uid kube config params
func (o *V1SpectroClustersUIDKubeConfigParams) SetFrp(frp *bool) {
	o.Frp = frp
}

// WithUID adds the uid to the v1 spectro clusters Uid kube config params
func (o *V1SpectroClustersUIDKubeConfigParams) WithUID(uid string) *V1SpectroClustersUIDKubeConfigParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 spectro clusters Uid kube config params
func (o *V1SpectroClustersUIDKubeConfigParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1SpectroClustersUIDKubeConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Frp != nil {

		// query param frp
		var qrFrp bool
		if o.Frp != nil {
			qrFrp = *o.Frp
		}
		qFrp := swag.FormatBool(qrFrp)
		if qFrp != "" {
			if err := r.SetQueryParam("frp", qFrp); err != nil {
				return err
			}
		}

	}

	// path param uid
	if err := r.SetPathParam("uid", o.UID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
