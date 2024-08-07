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

// NewV1SpectroClustersUpgradeSettingsGetParams creates a new V1SpectroClustersUpgradeSettingsGetParams object
// with the default values initialized.
func NewV1SpectroClustersUpgradeSettingsGetParams() *V1SpectroClustersUpgradeSettingsGetParams {

	return &V1SpectroClustersUpgradeSettingsGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1SpectroClustersUpgradeSettingsGetParamsWithTimeout creates a new V1SpectroClustersUpgradeSettingsGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1SpectroClustersUpgradeSettingsGetParamsWithTimeout(timeout time.Duration) *V1SpectroClustersUpgradeSettingsGetParams {

	return &V1SpectroClustersUpgradeSettingsGetParams{

		timeout: timeout,
	}
}

// NewV1SpectroClustersUpgradeSettingsGetParamsWithContext creates a new V1SpectroClustersUpgradeSettingsGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1SpectroClustersUpgradeSettingsGetParamsWithContext(ctx context.Context) *V1SpectroClustersUpgradeSettingsGetParams {

	return &V1SpectroClustersUpgradeSettingsGetParams{

		Context: ctx,
	}
}

// NewV1SpectroClustersUpgradeSettingsGetParamsWithHTTPClient creates a new V1SpectroClustersUpgradeSettingsGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1SpectroClustersUpgradeSettingsGetParamsWithHTTPClient(client *http.Client) *V1SpectroClustersUpgradeSettingsGetParams {

	return &V1SpectroClustersUpgradeSettingsGetParams{
		HTTPClient: client,
	}
}

/*V1SpectroClustersUpgradeSettingsGetParams contains all the parameters to send to the API endpoint
for the v1 spectro clusters upgrade settings get operation typically these are written to a http.Request
*/
type V1SpectroClustersUpgradeSettingsGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 spectro clusters upgrade settings get params
func (o *V1SpectroClustersUpgradeSettingsGetParams) WithTimeout(timeout time.Duration) *V1SpectroClustersUpgradeSettingsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 spectro clusters upgrade settings get params
func (o *V1SpectroClustersUpgradeSettingsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 spectro clusters upgrade settings get params
func (o *V1SpectroClustersUpgradeSettingsGetParams) WithContext(ctx context.Context) *V1SpectroClustersUpgradeSettingsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 spectro clusters upgrade settings get params
func (o *V1SpectroClustersUpgradeSettingsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 spectro clusters upgrade settings get params
func (o *V1SpectroClustersUpgradeSettingsGetParams) WithHTTPClient(client *http.Client) *V1SpectroClustersUpgradeSettingsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 spectro clusters upgrade settings get params
func (o *V1SpectroClustersUpgradeSettingsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *V1SpectroClustersUpgradeSettingsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
