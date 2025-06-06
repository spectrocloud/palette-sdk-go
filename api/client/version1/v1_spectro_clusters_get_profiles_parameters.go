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

// NewV1SpectroClustersGetProfilesParams creates a new V1SpectroClustersGetProfilesParams object
// with the default values initialized.
func NewV1SpectroClustersGetProfilesParams() *V1SpectroClustersGetProfilesParams {
	var ()
	return &V1SpectroClustersGetProfilesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1SpectroClustersGetProfilesParamsWithTimeout creates a new V1SpectroClustersGetProfilesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1SpectroClustersGetProfilesParamsWithTimeout(timeout time.Duration) *V1SpectroClustersGetProfilesParams {
	var ()
	return &V1SpectroClustersGetProfilesParams{

		timeout: timeout,
	}
}

// NewV1SpectroClustersGetProfilesParamsWithContext creates a new V1SpectroClustersGetProfilesParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1SpectroClustersGetProfilesParamsWithContext(ctx context.Context) *V1SpectroClustersGetProfilesParams {
	var ()
	return &V1SpectroClustersGetProfilesParams{

		Context: ctx,
	}
}

// NewV1SpectroClustersGetProfilesParamsWithHTTPClient creates a new V1SpectroClustersGetProfilesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1SpectroClustersGetProfilesParamsWithHTTPClient(client *http.Client) *V1SpectroClustersGetProfilesParams {
	var ()
	return &V1SpectroClustersGetProfilesParams{
		HTTPClient: client,
	}
}

/*
V1SpectroClustersGetProfilesParams contains all the parameters to send to the API endpoint
for the v1 spectro clusters get profiles operation typically these are written to a http.Request
*/
type V1SpectroClustersGetProfilesParams struct {

	/*IncludePackMeta
	  includes pack meta such as schema, presets

	*/
	IncludePackMeta *string
	/*UID
	  Cluster uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 spectro clusters get profiles params
func (o *V1SpectroClustersGetProfilesParams) WithTimeout(timeout time.Duration) *V1SpectroClustersGetProfilesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 spectro clusters get profiles params
func (o *V1SpectroClustersGetProfilesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 spectro clusters get profiles params
func (o *V1SpectroClustersGetProfilesParams) WithContext(ctx context.Context) *V1SpectroClustersGetProfilesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 spectro clusters get profiles params
func (o *V1SpectroClustersGetProfilesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 spectro clusters get profiles params
func (o *V1SpectroClustersGetProfilesParams) WithHTTPClient(client *http.Client) *V1SpectroClustersGetProfilesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 spectro clusters get profiles params
func (o *V1SpectroClustersGetProfilesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIncludePackMeta adds the includePackMeta to the v1 spectro clusters get profiles params
func (o *V1SpectroClustersGetProfilesParams) WithIncludePackMeta(includePackMeta *string) *V1SpectroClustersGetProfilesParams {
	o.SetIncludePackMeta(includePackMeta)
	return o
}

// SetIncludePackMeta adds the includePackMeta to the v1 spectro clusters get profiles params
func (o *V1SpectroClustersGetProfilesParams) SetIncludePackMeta(includePackMeta *string) {
	o.IncludePackMeta = includePackMeta
}

// WithUID adds the uid to the v1 spectro clusters get profiles params
func (o *V1SpectroClustersGetProfilesParams) WithUID(uid string) *V1SpectroClustersGetProfilesParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 spectro clusters get profiles params
func (o *V1SpectroClustersGetProfilesParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1SpectroClustersGetProfilesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.IncludePackMeta != nil {

		// query param includePackMeta
		var qrIncludePackMeta string
		if o.IncludePackMeta != nil {
			qrIncludePackMeta = *o.IncludePackMeta
		}
		qIncludePackMeta := qrIncludePackMeta
		if qIncludePackMeta != "" {
			if err := r.SetQueryParam("includePackMeta", qIncludePackMeta); err != nil {
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
