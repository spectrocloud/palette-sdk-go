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

// NewV1SpectroClustersUIDStatusParams creates a new V1SpectroClustersUIDStatusParams object
// with the default values initialized.
func NewV1SpectroClustersUIDStatusParams() *V1SpectroClustersUIDStatusParams {
	var ()
	return &V1SpectroClustersUIDStatusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1SpectroClustersUIDStatusParamsWithTimeout creates a new V1SpectroClustersUIDStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1SpectroClustersUIDStatusParamsWithTimeout(timeout time.Duration) *V1SpectroClustersUIDStatusParams {
	var ()
	return &V1SpectroClustersUIDStatusParams{

		timeout: timeout,
	}
}

// NewV1SpectroClustersUIDStatusParamsWithContext creates a new V1SpectroClustersUIDStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1SpectroClustersUIDStatusParamsWithContext(ctx context.Context) *V1SpectroClustersUIDStatusParams {
	var ()
	return &V1SpectroClustersUIDStatusParams{

		Context: ctx,
	}
}

// NewV1SpectroClustersUIDStatusParamsWithHTTPClient creates a new V1SpectroClustersUIDStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1SpectroClustersUIDStatusParamsWithHTTPClient(client *http.Client) *V1SpectroClustersUIDStatusParams {
	var ()
	return &V1SpectroClustersUIDStatusParams{
		HTTPClient: client,
	}
}

/*V1SpectroClustersUIDStatusParams contains all the parameters to send to the API endpoint
for the v1 spectro clusters Uid status operation typically these are written to a http.Request
*/
type V1SpectroClustersUIDStatusParams struct {

	/*UID
	  Cluster uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 spectro clusters Uid status params
func (o *V1SpectroClustersUIDStatusParams) WithTimeout(timeout time.Duration) *V1SpectroClustersUIDStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 spectro clusters Uid status params
func (o *V1SpectroClustersUIDStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 spectro clusters Uid status params
func (o *V1SpectroClustersUIDStatusParams) WithContext(ctx context.Context) *V1SpectroClustersUIDStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 spectro clusters Uid status params
func (o *V1SpectroClustersUIDStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 spectro clusters Uid status params
func (o *V1SpectroClustersUIDStatusParams) WithHTTPClient(client *http.Client) *V1SpectroClustersUIDStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 spectro clusters Uid status params
func (o *V1SpectroClustersUIDStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUID adds the uid to the v1 spectro clusters Uid status params
func (o *V1SpectroClustersUIDStatusParams) WithUID(uid string) *V1SpectroClustersUIDStatusParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 spectro clusters Uid status params
func (o *V1SpectroClustersUIDStatusParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1SpectroClustersUIDStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
