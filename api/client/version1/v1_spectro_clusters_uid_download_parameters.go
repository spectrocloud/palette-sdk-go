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

// NewV1SpectroClustersUIDDownloadParams creates a new V1SpectroClustersUIDDownloadParams object
// with the default values initialized.
func NewV1SpectroClustersUIDDownloadParams() *V1SpectroClustersUIDDownloadParams {
	var ()
	return &V1SpectroClustersUIDDownloadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1SpectroClustersUIDDownloadParamsWithTimeout creates a new V1SpectroClustersUIDDownloadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1SpectroClustersUIDDownloadParamsWithTimeout(timeout time.Duration) *V1SpectroClustersUIDDownloadParams {
	var ()
	return &V1SpectroClustersUIDDownloadParams{

		timeout: timeout,
	}
}

// NewV1SpectroClustersUIDDownloadParamsWithContext creates a new V1SpectroClustersUIDDownloadParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1SpectroClustersUIDDownloadParamsWithContext(ctx context.Context) *V1SpectroClustersUIDDownloadParams {
	var ()
	return &V1SpectroClustersUIDDownloadParams{

		Context: ctx,
	}
}

// NewV1SpectroClustersUIDDownloadParamsWithHTTPClient creates a new V1SpectroClustersUIDDownloadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1SpectroClustersUIDDownloadParamsWithHTTPClient(client *http.Client) *V1SpectroClustersUIDDownloadParams {
	var ()
	return &V1SpectroClustersUIDDownloadParams{
		HTTPClient: client,
	}
}

/*
V1SpectroClustersUIDDownloadParams contains all the parameters to send to the API endpoint
for the v1 spectro clusters Uid download operation typically these are written to a http.Request
*/
type V1SpectroClustersUIDDownloadParams struct {

	/*UID
	  Cluster uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 spectro clusters Uid download params
func (o *V1SpectroClustersUIDDownloadParams) WithTimeout(timeout time.Duration) *V1SpectroClustersUIDDownloadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 spectro clusters Uid download params
func (o *V1SpectroClustersUIDDownloadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 spectro clusters Uid download params
func (o *V1SpectroClustersUIDDownloadParams) WithContext(ctx context.Context) *V1SpectroClustersUIDDownloadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 spectro clusters Uid download params
func (o *V1SpectroClustersUIDDownloadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 spectro clusters Uid download params
func (o *V1SpectroClustersUIDDownloadParams) WithHTTPClient(client *http.Client) *V1SpectroClustersUIDDownloadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 spectro clusters Uid download params
func (o *V1SpectroClustersUIDDownloadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUID adds the uid to the v1 spectro clusters Uid download params
func (o *V1SpectroClustersUIDDownloadParams) WithUID(uid string) *V1SpectroClustersUIDDownloadParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 spectro clusters Uid download params
func (o *V1SpectroClustersUIDDownloadParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1SpectroClustersUIDDownloadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
