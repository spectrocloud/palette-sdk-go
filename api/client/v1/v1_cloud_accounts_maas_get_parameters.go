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

// NewV1CloudAccountsMaasGetParams creates a new V1CloudAccountsMaasGetParams object
// with the default values initialized.
func NewV1CloudAccountsMaasGetParams() *V1CloudAccountsMaasGetParams {
	var ()
	return &V1CloudAccountsMaasGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudAccountsMaasGetParamsWithTimeout creates a new V1CloudAccountsMaasGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudAccountsMaasGetParamsWithTimeout(timeout time.Duration) *V1CloudAccountsMaasGetParams {
	var ()
	return &V1CloudAccountsMaasGetParams{

		timeout: timeout,
	}
}

// NewV1CloudAccountsMaasGetParamsWithContext creates a new V1CloudAccountsMaasGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudAccountsMaasGetParamsWithContext(ctx context.Context) *V1CloudAccountsMaasGetParams {
	var ()
	return &V1CloudAccountsMaasGetParams{

		Context: ctx,
	}
}

// NewV1CloudAccountsMaasGetParamsWithHTTPClient creates a new V1CloudAccountsMaasGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudAccountsMaasGetParamsWithHTTPClient(client *http.Client) *V1CloudAccountsMaasGetParams {
	var ()
	return &V1CloudAccountsMaasGetParams{
		HTTPClient: client,
	}
}

/*V1CloudAccountsMaasGetParams contains all the parameters to send to the API endpoint
for the v1 cloud accounts maas get operation typically these are written to a http.Request
*/
type V1CloudAccountsMaasGetParams struct {

	/*UID
	  Maas cloud account uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cloud accounts maas get params
func (o *V1CloudAccountsMaasGetParams) WithTimeout(timeout time.Duration) *V1CloudAccountsMaasGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud accounts maas get params
func (o *V1CloudAccountsMaasGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud accounts maas get params
func (o *V1CloudAccountsMaasGetParams) WithContext(ctx context.Context) *V1CloudAccountsMaasGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud accounts maas get params
func (o *V1CloudAccountsMaasGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud accounts maas get params
func (o *V1CloudAccountsMaasGetParams) WithHTTPClient(client *http.Client) *V1CloudAccountsMaasGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud accounts maas get params
func (o *V1CloudAccountsMaasGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUID adds the uid to the v1 cloud accounts maas get params
func (o *V1CloudAccountsMaasGetParams) WithUID(uid string) *V1CloudAccountsMaasGetParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 cloud accounts maas get params
func (o *V1CloudAccountsMaasGetParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudAccountsMaasGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
