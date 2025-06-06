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

// NewV1SpectroClustersUIDPackManifestsUIDGetParams creates a new V1SpectroClustersUIDPackManifestsUIDGetParams object
// with the default values initialized.
func NewV1SpectroClustersUIDPackManifestsUIDGetParams() *V1SpectroClustersUIDPackManifestsUIDGetParams {
	var (
		resolveManifestValuesDefault = bool(false)
	)
	return &V1SpectroClustersUIDPackManifestsUIDGetParams{
		ResolveManifestValues: &resolveManifestValuesDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewV1SpectroClustersUIDPackManifestsUIDGetParamsWithTimeout creates a new V1SpectroClustersUIDPackManifestsUIDGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1SpectroClustersUIDPackManifestsUIDGetParamsWithTimeout(timeout time.Duration) *V1SpectroClustersUIDPackManifestsUIDGetParams {
	var (
		resolveManifestValuesDefault = bool(false)
	)
	return &V1SpectroClustersUIDPackManifestsUIDGetParams{
		ResolveManifestValues: &resolveManifestValuesDefault,

		timeout: timeout,
	}
}

// NewV1SpectroClustersUIDPackManifestsUIDGetParamsWithContext creates a new V1SpectroClustersUIDPackManifestsUIDGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1SpectroClustersUIDPackManifestsUIDGetParamsWithContext(ctx context.Context) *V1SpectroClustersUIDPackManifestsUIDGetParams {
	var (
		resolveManifestValuesDefault = bool(false)
	)
	return &V1SpectroClustersUIDPackManifestsUIDGetParams{
		ResolveManifestValues: &resolveManifestValuesDefault,

		Context: ctx,
	}
}

// NewV1SpectroClustersUIDPackManifestsUIDGetParamsWithHTTPClient creates a new V1SpectroClustersUIDPackManifestsUIDGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1SpectroClustersUIDPackManifestsUIDGetParamsWithHTTPClient(client *http.Client) *V1SpectroClustersUIDPackManifestsUIDGetParams {
	var (
		resolveManifestValuesDefault = bool(false)
	)
	return &V1SpectroClustersUIDPackManifestsUIDGetParams{
		ResolveManifestValues: &resolveManifestValuesDefault,
		HTTPClient:            client,
	}
}

/*
V1SpectroClustersUIDPackManifestsUIDGetParams contains all the parameters to send to the API endpoint
for the v1 spectro clusters Uid pack manifests Uid get operation typically these are written to a http.Request
*/
type V1SpectroClustersUIDPackManifestsUIDGetParams struct {

	/*ManifestUID
	  manifest uid which is part of the pack ref

	*/
	ManifestUID string
	/*ResolveManifestValues
	  resolve pack manifest values if set to true

	*/
	ResolveManifestValues *bool
	/*UID
	  cluster uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 spectro clusters Uid pack manifests Uid get params
func (o *V1SpectroClustersUIDPackManifestsUIDGetParams) WithTimeout(timeout time.Duration) *V1SpectroClustersUIDPackManifestsUIDGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 spectro clusters Uid pack manifests Uid get params
func (o *V1SpectroClustersUIDPackManifestsUIDGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 spectro clusters Uid pack manifests Uid get params
func (o *V1SpectroClustersUIDPackManifestsUIDGetParams) WithContext(ctx context.Context) *V1SpectroClustersUIDPackManifestsUIDGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 spectro clusters Uid pack manifests Uid get params
func (o *V1SpectroClustersUIDPackManifestsUIDGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 spectro clusters Uid pack manifests Uid get params
func (o *V1SpectroClustersUIDPackManifestsUIDGetParams) WithHTTPClient(client *http.Client) *V1SpectroClustersUIDPackManifestsUIDGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 spectro clusters Uid pack manifests Uid get params
func (o *V1SpectroClustersUIDPackManifestsUIDGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithManifestUID adds the manifestUID to the v1 spectro clusters Uid pack manifests Uid get params
func (o *V1SpectroClustersUIDPackManifestsUIDGetParams) WithManifestUID(manifestUID string) *V1SpectroClustersUIDPackManifestsUIDGetParams {
	o.SetManifestUID(manifestUID)
	return o
}

// SetManifestUID adds the manifestUid to the v1 spectro clusters Uid pack manifests Uid get params
func (o *V1SpectroClustersUIDPackManifestsUIDGetParams) SetManifestUID(manifestUID string) {
	o.ManifestUID = manifestUID
}

// WithResolveManifestValues adds the resolveManifestValues to the v1 spectro clusters Uid pack manifests Uid get params
func (o *V1SpectroClustersUIDPackManifestsUIDGetParams) WithResolveManifestValues(resolveManifestValues *bool) *V1SpectroClustersUIDPackManifestsUIDGetParams {
	o.SetResolveManifestValues(resolveManifestValues)
	return o
}

// SetResolveManifestValues adds the resolveManifestValues to the v1 spectro clusters Uid pack manifests Uid get params
func (o *V1SpectroClustersUIDPackManifestsUIDGetParams) SetResolveManifestValues(resolveManifestValues *bool) {
	o.ResolveManifestValues = resolveManifestValues
}

// WithUID adds the uid to the v1 spectro clusters Uid pack manifests Uid get params
func (o *V1SpectroClustersUIDPackManifestsUIDGetParams) WithUID(uid string) *V1SpectroClustersUIDPackManifestsUIDGetParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 spectro clusters Uid pack manifests Uid get params
func (o *V1SpectroClustersUIDPackManifestsUIDGetParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1SpectroClustersUIDPackManifestsUIDGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param manifestUid
	if err := r.SetPathParam("manifestUid", o.ManifestUID); err != nil {
		return err
	}

	if o.ResolveManifestValues != nil {

		// query param resolveManifestValues
		var qrResolveManifestValues bool
		if o.ResolveManifestValues != nil {
			qrResolveManifestValues = *o.ResolveManifestValues
		}
		qResolveManifestValues := swag.FormatBool(qrResolveManifestValues)
		if qResolveManifestValues != "" {
			if err := r.SetQueryParam("resolveManifestValues", qResolveManifestValues); err != nil {
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
