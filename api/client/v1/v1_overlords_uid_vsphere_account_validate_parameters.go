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

// NewV1OverlordsUIDVsphereAccountValidateParams creates a new V1OverlordsUIDVsphereAccountValidateParams object
// with the default values initialized.
func NewV1OverlordsUIDVsphereAccountValidateParams() *V1OverlordsUIDVsphereAccountValidateParams {
	var ()
	return &V1OverlordsUIDVsphereAccountValidateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1OverlordsUIDVsphereAccountValidateParamsWithTimeout creates a new V1OverlordsUIDVsphereAccountValidateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1OverlordsUIDVsphereAccountValidateParamsWithTimeout(timeout time.Duration) *V1OverlordsUIDVsphereAccountValidateParams {
	var ()
	return &V1OverlordsUIDVsphereAccountValidateParams{

		timeout: timeout,
	}
}

// NewV1OverlordsUIDVsphereAccountValidateParamsWithContext creates a new V1OverlordsUIDVsphereAccountValidateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1OverlordsUIDVsphereAccountValidateParamsWithContext(ctx context.Context) *V1OverlordsUIDVsphereAccountValidateParams {
	var ()
	return &V1OverlordsUIDVsphereAccountValidateParams{

		Context: ctx,
	}
}

// NewV1OverlordsUIDVsphereAccountValidateParamsWithHTTPClient creates a new V1OverlordsUIDVsphereAccountValidateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1OverlordsUIDVsphereAccountValidateParamsWithHTTPClient(client *http.Client) *V1OverlordsUIDVsphereAccountValidateParams {
	var ()
	return &V1OverlordsUIDVsphereAccountValidateParams{
		HTTPClient: client,
	}
}

/*
V1OverlordsUIDVsphereAccountValidateParams contains all the parameters to send to the API endpoint
for the v1 overlords Uid vsphere account validate operation typically these are written to a http.Request
*/
type V1OverlordsUIDVsphereAccountValidateParams struct {

	/*Body*/
	Body V1OverlordsUIDVsphereAccountValidateBody
	/*UID*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 overlords Uid vsphere account validate params
func (o *V1OverlordsUIDVsphereAccountValidateParams) WithTimeout(timeout time.Duration) *V1OverlordsUIDVsphereAccountValidateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 overlords Uid vsphere account validate params
func (o *V1OverlordsUIDVsphereAccountValidateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 overlords Uid vsphere account validate params
func (o *V1OverlordsUIDVsphereAccountValidateParams) WithContext(ctx context.Context) *V1OverlordsUIDVsphereAccountValidateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 overlords Uid vsphere account validate params
func (o *V1OverlordsUIDVsphereAccountValidateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 overlords Uid vsphere account validate params
func (o *V1OverlordsUIDVsphereAccountValidateParams) WithHTTPClient(client *http.Client) *V1OverlordsUIDVsphereAccountValidateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 overlords Uid vsphere account validate params
func (o *V1OverlordsUIDVsphereAccountValidateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 overlords Uid vsphere account validate params
func (o *V1OverlordsUIDVsphereAccountValidateParams) WithBody(body V1OverlordsUIDVsphereAccountValidateBody) *V1OverlordsUIDVsphereAccountValidateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 overlords Uid vsphere account validate params
func (o *V1OverlordsUIDVsphereAccountValidateParams) SetBody(body V1OverlordsUIDVsphereAccountValidateBody) {
	o.Body = body
}

// WithUID adds the uid to the v1 overlords Uid vsphere account validate params
func (o *V1OverlordsUIDVsphereAccountValidateParams) WithUID(uid string) *V1OverlordsUIDVsphereAccountValidateParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 overlords Uid vsphere account validate params
func (o *V1OverlordsUIDVsphereAccountValidateParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1OverlordsUIDVsphereAccountValidateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
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
