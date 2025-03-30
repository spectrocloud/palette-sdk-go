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

// NewV1OciRegistriesSummaryParams creates a new V1OciRegistriesSummaryParams object
// with the default values initialized.
func NewV1OciRegistriesSummaryParams() *V1OciRegistriesSummaryParams {

	return &V1OciRegistriesSummaryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1OciRegistriesSummaryParamsWithTimeout creates a new V1OciRegistriesSummaryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1OciRegistriesSummaryParamsWithTimeout(timeout time.Duration) *V1OciRegistriesSummaryParams {

	return &V1OciRegistriesSummaryParams{

		timeout: timeout,
	}
}

// NewV1OciRegistriesSummaryParamsWithContext creates a new V1OciRegistriesSummaryParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1OciRegistriesSummaryParamsWithContext(ctx context.Context) *V1OciRegistriesSummaryParams {

	return &V1OciRegistriesSummaryParams{

		Context: ctx,
	}
}

// NewV1OciRegistriesSummaryParamsWithHTTPClient creates a new V1OciRegistriesSummaryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1OciRegistriesSummaryParamsWithHTTPClient(client *http.Client) *V1OciRegistriesSummaryParams {

	return &V1OciRegistriesSummaryParams{
		HTTPClient: client,
	}
}

/*
V1OciRegistriesSummaryParams contains all the parameters to send to the API endpoint
for the v1 oci registries summary operation typically these are written to a http.Request
*/
type V1OciRegistriesSummaryParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 oci registries summary params
func (o *V1OciRegistriesSummaryParams) WithTimeout(timeout time.Duration) *V1OciRegistriesSummaryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 oci registries summary params
func (o *V1OciRegistriesSummaryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 oci registries summary params
func (o *V1OciRegistriesSummaryParams) WithContext(ctx context.Context) *V1OciRegistriesSummaryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 oci registries summary params
func (o *V1OciRegistriesSummaryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 oci registries summary params
func (o *V1OciRegistriesSummaryParams) WithHTTPClient(client *http.Client) *V1OciRegistriesSummaryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 oci registries summary params
func (o *V1OciRegistriesSummaryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *V1OciRegistriesSummaryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
