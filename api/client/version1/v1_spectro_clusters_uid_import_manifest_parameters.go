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

// NewV1SpectroClustersUIDImportManifestParams creates a new V1SpectroClustersUIDImportManifestParams object
// with the default values initialized.
func NewV1SpectroClustersUIDImportManifestParams() *V1SpectroClustersUIDImportManifestParams {
	var ()
	return &V1SpectroClustersUIDImportManifestParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1SpectroClustersUIDImportManifestParamsWithTimeout creates a new V1SpectroClustersUIDImportManifestParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1SpectroClustersUIDImportManifestParamsWithTimeout(timeout time.Duration) *V1SpectroClustersUIDImportManifestParams {
	var ()
	return &V1SpectroClustersUIDImportManifestParams{

		timeout: timeout,
	}
}

// NewV1SpectroClustersUIDImportManifestParamsWithContext creates a new V1SpectroClustersUIDImportManifestParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1SpectroClustersUIDImportManifestParamsWithContext(ctx context.Context) *V1SpectroClustersUIDImportManifestParams {
	var ()
	return &V1SpectroClustersUIDImportManifestParams{

		Context: ctx,
	}
}

// NewV1SpectroClustersUIDImportManifestParamsWithHTTPClient creates a new V1SpectroClustersUIDImportManifestParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1SpectroClustersUIDImportManifestParamsWithHTTPClient(client *http.Client) *V1SpectroClustersUIDImportManifestParams {
	var ()
	return &V1SpectroClustersUIDImportManifestParams{
		HTTPClient: client,
	}
}

/*
V1SpectroClustersUIDImportManifestParams contains all the parameters to send to the API endpoint
for the v1 spectro clusters Uid import manifest operation typically these are written to a http.Request
*/
type V1SpectroClustersUIDImportManifestParams struct {

	/*UID*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 spectro clusters Uid import manifest params
func (o *V1SpectroClustersUIDImportManifestParams) WithTimeout(timeout time.Duration) *V1SpectroClustersUIDImportManifestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 spectro clusters Uid import manifest params
func (o *V1SpectroClustersUIDImportManifestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 spectro clusters Uid import manifest params
func (o *V1SpectroClustersUIDImportManifestParams) WithContext(ctx context.Context) *V1SpectroClustersUIDImportManifestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 spectro clusters Uid import manifest params
func (o *V1SpectroClustersUIDImportManifestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 spectro clusters Uid import manifest params
func (o *V1SpectroClustersUIDImportManifestParams) WithHTTPClient(client *http.Client) *V1SpectroClustersUIDImportManifestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 spectro clusters Uid import manifest params
func (o *V1SpectroClustersUIDImportManifestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUID adds the uid to the v1 spectro clusters Uid import manifest params
func (o *V1SpectroClustersUIDImportManifestParams) WithUID(uid string) *V1SpectroClustersUIDImportManifestParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 spectro clusters Uid import manifest params
func (o *V1SpectroClustersUIDImportManifestParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1SpectroClustersUIDImportManifestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
