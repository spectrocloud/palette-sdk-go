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

// NewV1AzureStorageTypesParams creates a new V1AzureStorageTypesParams object
// with the default values initialized.
func NewV1AzureStorageTypesParams() *V1AzureStorageTypesParams {
	var ()
	return &V1AzureStorageTypesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1AzureStorageTypesParamsWithTimeout creates a new V1AzureStorageTypesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1AzureStorageTypesParamsWithTimeout(timeout time.Duration) *V1AzureStorageTypesParams {
	var ()
	return &V1AzureStorageTypesParams{

		timeout: timeout,
	}
}

// NewV1AzureStorageTypesParamsWithContext creates a new V1AzureStorageTypesParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1AzureStorageTypesParamsWithContext(ctx context.Context) *V1AzureStorageTypesParams {
	var ()
	return &V1AzureStorageTypesParams{

		Context: ctx,
	}
}

// NewV1AzureStorageTypesParamsWithHTTPClient creates a new V1AzureStorageTypesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1AzureStorageTypesParamsWithHTTPClient(client *http.Client) *V1AzureStorageTypesParams {
	var ()
	return &V1AzureStorageTypesParams{
		HTTPClient: client,
	}
}

/*
V1AzureStorageTypesParams contains all the parameters to send to the API endpoint
for the v1 azure storage types operation typically these are written to a http.Request
*/
type V1AzureStorageTypesParams struct {

	/*Region
	  Region for which Azure storage types are requested

	*/
	Region string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 azure storage types params
func (o *V1AzureStorageTypesParams) WithTimeout(timeout time.Duration) *V1AzureStorageTypesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 azure storage types params
func (o *V1AzureStorageTypesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 azure storage types params
func (o *V1AzureStorageTypesParams) WithContext(ctx context.Context) *V1AzureStorageTypesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 azure storage types params
func (o *V1AzureStorageTypesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 azure storage types params
func (o *V1AzureStorageTypesParams) WithHTTPClient(client *http.Client) *V1AzureStorageTypesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 azure storage types params
func (o *V1AzureStorageTypesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRegion adds the region to the v1 azure storage types params
func (o *V1AzureStorageTypesParams) WithRegion(region string) *V1AzureStorageTypesParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the v1 azure storage types params
func (o *V1AzureStorageTypesParams) SetRegion(region string) {
	o.Region = region
}

// WriteToRequest writes these params to a swagger request
func (o *V1AzureStorageTypesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param region
	if err := r.SetPathParam("region", o.Region); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
