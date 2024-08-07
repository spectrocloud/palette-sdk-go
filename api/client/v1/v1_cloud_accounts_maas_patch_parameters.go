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

	"github.com/spectrocloud/palette-sdk-go/api/models"
)

// NewV1CloudAccountsMaasPatchParams creates a new V1CloudAccountsMaasPatchParams object
// with the default values initialized.
func NewV1CloudAccountsMaasPatchParams() *V1CloudAccountsMaasPatchParams {
	var ()
	return &V1CloudAccountsMaasPatchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudAccountsMaasPatchParamsWithTimeout creates a new V1CloudAccountsMaasPatchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudAccountsMaasPatchParamsWithTimeout(timeout time.Duration) *V1CloudAccountsMaasPatchParams {
	var ()
	return &V1CloudAccountsMaasPatchParams{

		timeout: timeout,
	}
}

// NewV1CloudAccountsMaasPatchParamsWithContext creates a new V1CloudAccountsMaasPatchParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudAccountsMaasPatchParamsWithContext(ctx context.Context) *V1CloudAccountsMaasPatchParams {
	var ()
	return &V1CloudAccountsMaasPatchParams{

		Context: ctx,
	}
}

// NewV1CloudAccountsMaasPatchParamsWithHTTPClient creates a new V1CloudAccountsMaasPatchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudAccountsMaasPatchParamsWithHTTPClient(client *http.Client) *V1CloudAccountsMaasPatchParams {
	var ()
	return &V1CloudAccountsMaasPatchParams{
		HTTPClient: client,
	}
}

/*
V1CloudAccountsMaasPatchParams contains all the parameters to send to the API endpoint
for the v1 cloud accounts maas patch operation typically these are written to a http.Request
*/
type V1CloudAccountsMaasPatchParams struct {

	/*Body*/
	Body models.V1CloudAccountsPatch
	/*UID
	  Maas cloud account uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cloud accounts maas patch params
func (o *V1CloudAccountsMaasPatchParams) WithTimeout(timeout time.Duration) *V1CloudAccountsMaasPatchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud accounts maas patch params
func (o *V1CloudAccountsMaasPatchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud accounts maas patch params
func (o *V1CloudAccountsMaasPatchParams) WithContext(ctx context.Context) *V1CloudAccountsMaasPatchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud accounts maas patch params
func (o *V1CloudAccountsMaasPatchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud accounts maas patch params
func (o *V1CloudAccountsMaasPatchParams) WithHTTPClient(client *http.Client) *V1CloudAccountsMaasPatchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud accounts maas patch params
func (o *V1CloudAccountsMaasPatchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 cloud accounts maas patch params
func (o *V1CloudAccountsMaasPatchParams) WithBody(body models.V1CloudAccountsPatch) *V1CloudAccountsMaasPatchParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 cloud accounts maas patch params
func (o *V1CloudAccountsMaasPatchParams) SetBody(body models.V1CloudAccountsPatch) {
	o.Body = body
}

// WithUID adds the uid to the v1 cloud accounts maas patch params
func (o *V1CloudAccountsMaasPatchParams) WithUID(uid string) *V1CloudAccountsMaasPatchParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 cloud accounts maas patch params
func (o *V1CloudAccountsMaasPatchParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudAccountsMaasPatchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
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
