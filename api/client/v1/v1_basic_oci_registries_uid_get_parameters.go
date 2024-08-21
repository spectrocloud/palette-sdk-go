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

// NewV1BasicOciRegistriesUIDGetParams creates a new V1BasicOciRegistriesUIDGetParams object
// with the default values initialized.
func NewV1BasicOciRegistriesUIDGetParams() *V1BasicOciRegistriesUIDGetParams {
	var ()
	return &V1BasicOciRegistriesUIDGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1BasicOciRegistriesUIDGetParamsWithTimeout creates a new V1BasicOciRegistriesUIDGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1BasicOciRegistriesUIDGetParamsWithTimeout(timeout time.Duration) *V1BasicOciRegistriesUIDGetParams {
	var ()
	return &V1BasicOciRegistriesUIDGetParams{

		timeout: timeout,
	}
}

// NewV1BasicOciRegistriesUIDGetParamsWithContext creates a new V1BasicOciRegistriesUIDGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1BasicOciRegistriesUIDGetParamsWithContext(ctx context.Context) *V1BasicOciRegistriesUIDGetParams {
	var ()
	return &V1BasicOciRegistriesUIDGetParams{

		Context: ctx,
	}
}

// NewV1BasicOciRegistriesUIDGetParamsWithHTTPClient creates a new V1BasicOciRegistriesUIDGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1BasicOciRegistriesUIDGetParamsWithHTTPClient(client *http.Client) *V1BasicOciRegistriesUIDGetParams {
	var ()
	return &V1BasicOciRegistriesUIDGetParams{
		HTTPClient: client,
	}
}

/*
V1BasicOciRegistriesUIDGetParams contains all the parameters to send to the API endpoint
for the v1 basic oci registries Uid get operation typically these are written to a http.Request
*/
type V1BasicOciRegistriesUIDGetParams struct {

	/*UID*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 basic oci registries Uid get params
func (o *V1BasicOciRegistriesUIDGetParams) WithTimeout(timeout time.Duration) *V1BasicOciRegistriesUIDGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 basic oci registries Uid get params
func (o *V1BasicOciRegistriesUIDGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 basic oci registries Uid get params
func (o *V1BasicOciRegistriesUIDGetParams) WithContext(ctx context.Context) *V1BasicOciRegistriesUIDGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 basic oci registries Uid get params
func (o *V1BasicOciRegistriesUIDGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 basic oci registries Uid get params
func (o *V1BasicOciRegistriesUIDGetParams) WithHTTPClient(client *http.Client) *V1BasicOciRegistriesUIDGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 basic oci registries Uid get params
func (o *V1BasicOciRegistriesUIDGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUID adds the uid to the v1 basic oci registries Uid get params
func (o *V1BasicOciRegistriesUIDGetParams) WithUID(uid string) *V1BasicOciRegistriesUIDGetParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 basic oci registries Uid get params
func (o *V1BasicOciRegistriesUIDGetParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1BasicOciRegistriesUIDGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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