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

// NewV1BasicOciRegistriesUIDDeleteParams creates a new V1BasicOciRegistriesUIDDeleteParams object
// with the default values initialized.
func NewV1BasicOciRegistriesUIDDeleteParams() *V1BasicOciRegistriesUIDDeleteParams {
	var ()
	return &V1BasicOciRegistriesUIDDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1BasicOciRegistriesUIDDeleteParamsWithTimeout creates a new V1BasicOciRegistriesUIDDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1BasicOciRegistriesUIDDeleteParamsWithTimeout(timeout time.Duration) *V1BasicOciRegistriesUIDDeleteParams {
	var ()
	return &V1BasicOciRegistriesUIDDeleteParams{

		timeout: timeout,
	}
}

// NewV1BasicOciRegistriesUIDDeleteParamsWithContext creates a new V1BasicOciRegistriesUIDDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1BasicOciRegistriesUIDDeleteParamsWithContext(ctx context.Context) *V1BasicOciRegistriesUIDDeleteParams {
	var ()
	return &V1BasicOciRegistriesUIDDeleteParams{

		Context: ctx,
	}
}

// NewV1BasicOciRegistriesUIDDeleteParamsWithHTTPClient creates a new V1BasicOciRegistriesUIDDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1BasicOciRegistriesUIDDeleteParamsWithHTTPClient(client *http.Client) *V1BasicOciRegistriesUIDDeleteParams {
	var ()
	return &V1BasicOciRegistriesUIDDeleteParams{
		HTTPClient: client,
	}
}

/*
V1BasicOciRegistriesUIDDeleteParams contains all the parameters to send to the API endpoint
for the v1 basic oci registries Uid delete operation typically these are written to a http.Request
*/
type V1BasicOciRegistriesUIDDeleteParams struct {

	/*UID*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 basic oci registries Uid delete params
func (o *V1BasicOciRegistriesUIDDeleteParams) WithTimeout(timeout time.Duration) *V1BasicOciRegistriesUIDDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 basic oci registries Uid delete params
func (o *V1BasicOciRegistriesUIDDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 basic oci registries Uid delete params
func (o *V1BasicOciRegistriesUIDDeleteParams) WithContext(ctx context.Context) *V1BasicOciRegistriesUIDDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 basic oci registries Uid delete params
func (o *V1BasicOciRegistriesUIDDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 basic oci registries Uid delete params
func (o *V1BasicOciRegistriesUIDDeleteParams) WithHTTPClient(client *http.Client) *V1BasicOciRegistriesUIDDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 basic oci registries Uid delete params
func (o *V1BasicOciRegistriesUIDDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUID adds the uid to the v1 basic oci registries Uid delete params
func (o *V1BasicOciRegistriesUIDDeleteParams) WithUID(uid string) *V1BasicOciRegistriesUIDDeleteParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 basic oci registries Uid delete params
func (o *V1BasicOciRegistriesUIDDeleteParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1BasicOciRegistriesUIDDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
