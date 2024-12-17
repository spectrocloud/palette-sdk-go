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
	"github.com/go-openapi/swag"
)

// NewV1SpectroClustersGetParams creates a new V1SpectroClustersGetParams object
// with the default values initialized.
func NewV1SpectroClustersGetParams() *V1SpectroClustersGetParams {
	var (
		includeNonSpectroLabelsDefault = bool(false)
		resolvePackValuesDefault       = bool(false)
	)
	return &V1SpectroClustersGetParams{
		IncludeNonSpectroLabels: &includeNonSpectroLabelsDefault,
		ResolvePackValues:       &resolvePackValuesDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewV1SpectroClustersGetParamsWithTimeout creates a new V1SpectroClustersGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1SpectroClustersGetParamsWithTimeout(timeout time.Duration) *V1SpectroClustersGetParams {
	var (
		includeNonSpectroLabelsDefault = bool(false)
		resolvePackValuesDefault       = bool(false)
	)
	return &V1SpectroClustersGetParams{
		IncludeNonSpectroLabels: &includeNonSpectroLabelsDefault,
		ResolvePackValues:       &resolvePackValuesDefault,

		timeout: timeout,
	}
}

// NewV1SpectroClustersGetParamsWithContext creates a new V1SpectroClustersGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1SpectroClustersGetParamsWithContext(ctx context.Context) *V1SpectroClustersGetParams {
	var (
		includeNonSpectroLabelsDefault = bool(false)
		resolvePackValuesDefault       = bool(false)
	)
	return &V1SpectroClustersGetParams{
		IncludeNonSpectroLabels: &includeNonSpectroLabelsDefault,
		ResolvePackValues:       &resolvePackValuesDefault,

		Context: ctx,
	}
}

// NewV1SpectroClustersGetParamsWithHTTPClient creates a new V1SpectroClustersGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1SpectroClustersGetParamsWithHTTPClient(client *http.Client) *V1SpectroClustersGetParams {
	var (
		includeNonSpectroLabelsDefault = bool(false)
		resolvePackValuesDefault       = bool(false)
	)
	return &V1SpectroClustersGetParams{
		IncludeNonSpectroLabels: &includeNonSpectroLabelsDefault,
		ResolvePackValues:       &resolvePackValuesDefault,
		HTTPClient:              client,
	}
}

/*
V1SpectroClustersGetParams contains all the parameters to send to the API endpoint
for the v1 spectro clusters get operation typically these are written to a http.Request
*/
type V1SpectroClustersGetParams struct {

	/*IncludeNonSpectroLabels
	  Include non spectro labels in the cluster labels if set to true

	*/
	IncludeNonSpectroLabels *bool
	/*IncludePackMeta
	  Includes pack meta such as schema, presets

	*/
	IncludePackMeta *string
	/*IncludeTags
	  Comma separated tags like system,profile

	*/
	IncludeTags *string
	/*ProfileType
	  Filter cluster profile templates by profileType

	*/
	ProfileType *string
	/*ResolvePackValues
	  Resolve pack values if set to true

	*/
	ResolvePackValues *bool
	/*UID
	  Cluster uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 spectro clusters get params
func (o *V1SpectroClustersGetParams) WithTimeout(timeout time.Duration) *V1SpectroClustersGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 spectro clusters get params
func (o *V1SpectroClustersGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 spectro clusters get params
func (o *V1SpectroClustersGetParams) WithContext(ctx context.Context) *V1SpectroClustersGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 spectro clusters get params
func (o *V1SpectroClustersGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 spectro clusters get params
func (o *V1SpectroClustersGetParams) WithHTTPClient(client *http.Client) *V1SpectroClustersGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 spectro clusters get params
func (o *V1SpectroClustersGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIncludeNonSpectroLabels adds the includeNonSpectroLabels to the v1 spectro clusters get params
func (o *V1SpectroClustersGetParams) WithIncludeNonSpectroLabels(includeNonSpectroLabels *bool) *V1SpectroClustersGetParams {
	o.SetIncludeNonSpectroLabels(includeNonSpectroLabels)
	return o
}

// SetIncludeNonSpectroLabels adds the includeNonSpectroLabels to the v1 spectro clusters get params
func (o *V1SpectroClustersGetParams) SetIncludeNonSpectroLabels(includeNonSpectroLabels *bool) {
	o.IncludeNonSpectroLabels = includeNonSpectroLabels
}

// WithIncludePackMeta adds the includePackMeta to the v1 spectro clusters get params
func (o *V1SpectroClustersGetParams) WithIncludePackMeta(includePackMeta *string) *V1SpectroClustersGetParams {
	o.SetIncludePackMeta(includePackMeta)
	return o
}

// SetIncludePackMeta adds the includePackMeta to the v1 spectro clusters get params
func (o *V1SpectroClustersGetParams) SetIncludePackMeta(includePackMeta *string) {
	o.IncludePackMeta = includePackMeta
}

// WithIncludeTags adds the includeTags to the v1 spectro clusters get params
func (o *V1SpectroClustersGetParams) WithIncludeTags(includeTags *string) *V1SpectroClustersGetParams {
	o.SetIncludeTags(includeTags)
	return o
}

// SetIncludeTags adds the includeTags to the v1 spectro clusters get params
func (o *V1SpectroClustersGetParams) SetIncludeTags(includeTags *string) {
	o.IncludeTags = includeTags
}

// WithProfileType adds the profileType to the v1 spectro clusters get params
func (o *V1SpectroClustersGetParams) WithProfileType(profileType *string) *V1SpectroClustersGetParams {
	o.SetProfileType(profileType)
	return o
}

// SetProfileType adds the profileType to the v1 spectro clusters get params
func (o *V1SpectroClustersGetParams) SetProfileType(profileType *string) {
	o.ProfileType = profileType
}

// WithResolvePackValues adds the resolvePackValues to the v1 spectro clusters get params
func (o *V1SpectroClustersGetParams) WithResolvePackValues(resolvePackValues *bool) *V1SpectroClustersGetParams {
	o.SetResolvePackValues(resolvePackValues)
	return o
}

// SetResolvePackValues adds the resolvePackValues to the v1 spectro clusters get params
func (o *V1SpectroClustersGetParams) SetResolvePackValues(resolvePackValues *bool) {
	o.ResolvePackValues = resolvePackValues
}

// WithUID adds the uid to the v1 spectro clusters get params
func (o *V1SpectroClustersGetParams) WithUID(uid string) *V1SpectroClustersGetParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 spectro clusters get params
func (o *V1SpectroClustersGetParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1SpectroClustersGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.IncludeNonSpectroLabels != nil {

		// query param includeNonSpectroLabels
		var qrIncludeNonSpectroLabels bool
		if o.IncludeNonSpectroLabels != nil {
			qrIncludeNonSpectroLabels = *o.IncludeNonSpectroLabels
		}
		qIncludeNonSpectroLabels := swag.FormatBool(qrIncludeNonSpectroLabels)
		if qIncludeNonSpectroLabels != "" {
			if err := r.SetQueryParam("includeNonSpectroLabels", qIncludeNonSpectroLabels); err != nil {
				return err
			}
		}

	}

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

	if o.IncludeTags != nil {

		// query param includeTags
		var qrIncludeTags string
		if o.IncludeTags != nil {
			qrIncludeTags = *o.IncludeTags
		}
		qIncludeTags := qrIncludeTags
		if qIncludeTags != "" {
			if err := r.SetQueryParam("includeTags", qIncludeTags); err != nil {
				return err
			}
		}

	}

	if o.ProfileType != nil {

		// query param profileType
		var qrProfileType string
		if o.ProfileType != nil {
			qrProfileType = *o.ProfileType
		}
		qProfileType := qrProfileType
		if qProfileType != "" {
			if err := r.SetQueryParam("profileType", qProfileType); err != nil {
				return err
			}
		}

	}

	if o.ResolvePackValues != nil {

		// query param resolvePackValues
		var qrResolvePackValues bool
		if o.ResolvePackValues != nil {
			qrResolvePackValues = *o.ResolvePackValues
		}
		qResolvePackValues := swag.FormatBool(qrResolvePackValues)
		if qResolvePackValues != "" {
			if err := r.SetQueryParam("resolvePackValues", qResolvePackValues); err != nil {
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
