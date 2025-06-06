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

// NewV1AzureSubscriptionListParams creates a new V1AzureSubscriptionListParams object
// with the default values initialized.
func NewV1AzureSubscriptionListParams() *V1AzureSubscriptionListParams {
	var ()
	return &V1AzureSubscriptionListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1AzureSubscriptionListParamsWithTimeout creates a new V1AzureSubscriptionListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1AzureSubscriptionListParamsWithTimeout(timeout time.Duration) *V1AzureSubscriptionListParams {
	var ()
	return &V1AzureSubscriptionListParams{

		timeout: timeout,
	}
}

// NewV1AzureSubscriptionListParamsWithContext creates a new V1AzureSubscriptionListParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1AzureSubscriptionListParamsWithContext(ctx context.Context) *V1AzureSubscriptionListParams {
	var ()
	return &V1AzureSubscriptionListParams{

		Context: ctx,
	}
}

// NewV1AzureSubscriptionListParamsWithHTTPClient creates a new V1AzureSubscriptionListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1AzureSubscriptionListParamsWithHTTPClient(client *http.Client) *V1AzureSubscriptionListParams {
	var ()
	return &V1AzureSubscriptionListParams{
		HTTPClient: client,
	}
}

/*
V1AzureSubscriptionListParams contains all the parameters to send to the API endpoint
for the v1 azure subscription list operation typically these are written to a http.Request
*/
type V1AzureSubscriptionListParams struct {

	/*CloudAccountUID
	  Uid for the specific Azure cloud account

	*/
	CloudAccountUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 azure subscription list params
func (o *V1AzureSubscriptionListParams) WithTimeout(timeout time.Duration) *V1AzureSubscriptionListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 azure subscription list params
func (o *V1AzureSubscriptionListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 azure subscription list params
func (o *V1AzureSubscriptionListParams) WithContext(ctx context.Context) *V1AzureSubscriptionListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 azure subscription list params
func (o *V1AzureSubscriptionListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 azure subscription list params
func (o *V1AzureSubscriptionListParams) WithHTTPClient(client *http.Client) *V1AzureSubscriptionListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 azure subscription list params
func (o *V1AzureSubscriptionListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudAccountUID adds the cloudAccountUID to the v1 azure subscription list params
func (o *V1AzureSubscriptionListParams) WithCloudAccountUID(cloudAccountUID string) *V1AzureSubscriptionListParams {
	o.SetCloudAccountUID(cloudAccountUID)
	return o
}

// SetCloudAccountUID adds the cloudAccountUid to the v1 azure subscription list params
func (o *V1AzureSubscriptionListParams) SetCloudAccountUID(cloudAccountUID string) {
	o.CloudAccountUID = cloudAccountUID
}

// WriteToRequest writes these params to a swagger request
func (o *V1AzureSubscriptionListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param cloudAccountUid
	qrCloudAccountUID := o.CloudAccountUID
	qCloudAccountUID := qrCloudAccountUID
	if qCloudAccountUID != "" {
		if err := r.SetQueryParam("cloudAccountUid", qCloudAccountUID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
