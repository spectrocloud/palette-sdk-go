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

// NewV1SpectroClustersUIDStatusSpcApplyParams creates a new V1SpectroClustersUIDStatusSpcApplyParams object
// with the default values initialized.
func NewV1SpectroClustersUIDStatusSpcApplyParams() *V1SpectroClustersUIDStatusSpcApplyParams {
	var ()
	return &V1SpectroClustersUIDStatusSpcApplyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1SpectroClustersUIDStatusSpcApplyParamsWithTimeout creates a new V1SpectroClustersUIDStatusSpcApplyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1SpectroClustersUIDStatusSpcApplyParamsWithTimeout(timeout time.Duration) *V1SpectroClustersUIDStatusSpcApplyParams {
	var ()
	return &V1SpectroClustersUIDStatusSpcApplyParams{

		timeout: timeout,
	}
}

// NewV1SpectroClustersUIDStatusSpcApplyParamsWithContext creates a new V1SpectroClustersUIDStatusSpcApplyParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1SpectroClustersUIDStatusSpcApplyParamsWithContext(ctx context.Context) *V1SpectroClustersUIDStatusSpcApplyParams {
	var ()
	return &V1SpectroClustersUIDStatusSpcApplyParams{

		Context: ctx,
	}
}

// NewV1SpectroClustersUIDStatusSpcApplyParamsWithHTTPClient creates a new V1SpectroClustersUIDStatusSpcApplyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1SpectroClustersUIDStatusSpcApplyParamsWithHTTPClient(client *http.Client) *V1SpectroClustersUIDStatusSpcApplyParams {
	var ()
	return &V1SpectroClustersUIDStatusSpcApplyParams{
		HTTPClient: client,
	}
}

/*
V1SpectroClustersUIDStatusSpcApplyParams contains all the parameters to send to the API endpoint
for the v1 spectro clusters Uid status spc apply operation typically these are written to a http.Request
*/
type V1SpectroClustersUIDStatusSpcApplyParams struct {

	/*UID*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 spectro clusters Uid status spc apply params
func (o *V1SpectroClustersUIDStatusSpcApplyParams) WithTimeout(timeout time.Duration) *V1SpectroClustersUIDStatusSpcApplyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 spectro clusters Uid status spc apply params
func (o *V1SpectroClustersUIDStatusSpcApplyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 spectro clusters Uid status spc apply params
func (o *V1SpectroClustersUIDStatusSpcApplyParams) WithContext(ctx context.Context) *V1SpectroClustersUIDStatusSpcApplyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 spectro clusters Uid status spc apply params
func (o *V1SpectroClustersUIDStatusSpcApplyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 spectro clusters Uid status spc apply params
func (o *V1SpectroClustersUIDStatusSpcApplyParams) WithHTTPClient(client *http.Client) *V1SpectroClustersUIDStatusSpcApplyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 spectro clusters Uid status spc apply params
func (o *V1SpectroClustersUIDStatusSpcApplyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUID adds the uid to the v1 spectro clusters Uid status spc apply params
func (o *V1SpectroClustersUIDStatusSpcApplyParams) WithUID(uid string) *V1SpectroClustersUIDStatusSpcApplyParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 spectro clusters Uid status spc apply params
func (o *V1SpectroClustersUIDStatusSpcApplyParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1SpectroClustersUIDStatusSpcApplyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
