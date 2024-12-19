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

// NewV1AzureGroupsParams creates a new V1AzureGroupsParams object
// with the default values initialized.
func NewV1AzureGroupsParams() *V1AzureGroupsParams {
	var ()
	return &V1AzureGroupsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1AzureGroupsParamsWithTimeout creates a new V1AzureGroupsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1AzureGroupsParamsWithTimeout(timeout time.Duration) *V1AzureGroupsParams {
	var ()
	return &V1AzureGroupsParams{

		timeout: timeout,
	}
}

// NewV1AzureGroupsParamsWithContext creates a new V1AzureGroupsParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1AzureGroupsParamsWithContext(ctx context.Context) *V1AzureGroupsParams {
	var ()
	return &V1AzureGroupsParams{

		Context: ctx,
	}
}

// NewV1AzureGroupsParamsWithHTTPClient creates a new V1AzureGroupsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1AzureGroupsParamsWithHTTPClient(client *http.Client) *V1AzureGroupsParams {
	var ()
	return &V1AzureGroupsParams{
		HTTPClient: client,
	}
}

/*
V1AzureGroupsParams contains all the parameters to send to the API endpoint
for the v1 azure groups operation typically these are written to a http.Request
*/
type V1AzureGroupsParams struct {

	/*CloudAccountUID
	  Uid for the specific Azure cloud account

	*/
	CloudAccountUID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 azure groups params
func (o *V1AzureGroupsParams) WithTimeout(timeout time.Duration) *V1AzureGroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 azure groups params
func (o *V1AzureGroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 azure groups params
func (o *V1AzureGroupsParams) WithContext(ctx context.Context) *V1AzureGroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 azure groups params
func (o *V1AzureGroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 azure groups params
func (o *V1AzureGroupsParams) WithHTTPClient(client *http.Client) *V1AzureGroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 azure groups params
func (o *V1AzureGroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudAccountUID adds the cloudAccountUID to the v1 azure groups params
func (o *V1AzureGroupsParams) WithCloudAccountUID(cloudAccountUID *string) *V1AzureGroupsParams {
	o.SetCloudAccountUID(cloudAccountUID)
	return o
}

// SetCloudAccountUID adds the cloudAccountUid to the v1 azure groups params
func (o *V1AzureGroupsParams) SetCloudAccountUID(cloudAccountUID *string) {
	o.CloudAccountUID = cloudAccountUID
}

// WriteToRequest writes these params to a swagger request
func (o *V1AzureGroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CloudAccountUID != nil {

		// query param cloudAccountUid
		var qrCloudAccountUID string
		if o.CloudAccountUID != nil {
			qrCloudAccountUID = *o.CloudAccountUID
		}
		qCloudAccountUID := qrCloudAccountUID
		if qCloudAccountUID != "" {
			if err := r.SetQueryParam("cloudAccountUid", qCloudAccountUID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
