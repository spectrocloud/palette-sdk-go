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

// NewV1AzurePrivateDNSZonesParams creates a new V1AzurePrivateDNSZonesParams object
// with the default values initialized.
func NewV1AzurePrivateDNSZonesParams() *V1AzurePrivateDNSZonesParams {
	var ()
	return &V1AzurePrivateDNSZonesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1AzurePrivateDNSZonesParamsWithTimeout creates a new V1AzurePrivateDNSZonesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1AzurePrivateDNSZonesParamsWithTimeout(timeout time.Duration) *V1AzurePrivateDNSZonesParams {
	var ()
	return &V1AzurePrivateDNSZonesParams{

		timeout: timeout,
	}
}

// NewV1AzurePrivateDNSZonesParamsWithContext creates a new V1AzurePrivateDNSZonesParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1AzurePrivateDNSZonesParamsWithContext(ctx context.Context) *V1AzurePrivateDNSZonesParams {
	var ()
	return &V1AzurePrivateDNSZonesParams{

		Context: ctx,
	}
}

// NewV1AzurePrivateDNSZonesParamsWithHTTPClient creates a new V1AzurePrivateDNSZonesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1AzurePrivateDNSZonesParamsWithHTTPClient(client *http.Client) *V1AzurePrivateDNSZonesParams {
	var ()
	return &V1AzurePrivateDNSZonesParams{
		HTTPClient: client,
	}
}

/*
V1AzurePrivateDNSZonesParams contains all the parameters to send to the API endpoint
for the v1 azure private Dns zones operation typically these are written to a http.Request
*/
type V1AzurePrivateDNSZonesParams struct {

	/*CloudAccountUID
	  Uid for the specific Azure cloud account

	*/
	CloudAccountUID string
	/*ResourceGroup
	  resourceGroup for which Azure private dns zones are requested

	*/
	ResourceGroup string
	/*SubscriptionID
	  subscriptionId for which Azure private dns zones are requested

	*/
	SubscriptionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 azure private Dns zones params
func (o *V1AzurePrivateDNSZonesParams) WithTimeout(timeout time.Duration) *V1AzurePrivateDNSZonesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 azure private Dns zones params
func (o *V1AzurePrivateDNSZonesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 azure private Dns zones params
func (o *V1AzurePrivateDNSZonesParams) WithContext(ctx context.Context) *V1AzurePrivateDNSZonesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 azure private Dns zones params
func (o *V1AzurePrivateDNSZonesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 azure private Dns zones params
func (o *V1AzurePrivateDNSZonesParams) WithHTTPClient(client *http.Client) *V1AzurePrivateDNSZonesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 azure private Dns zones params
func (o *V1AzurePrivateDNSZonesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudAccountUID adds the cloudAccountUID to the v1 azure private Dns zones params
func (o *V1AzurePrivateDNSZonesParams) WithCloudAccountUID(cloudAccountUID string) *V1AzurePrivateDNSZonesParams {
	o.SetCloudAccountUID(cloudAccountUID)
	return o
}

// SetCloudAccountUID adds the cloudAccountUid to the v1 azure private Dns zones params
func (o *V1AzurePrivateDNSZonesParams) SetCloudAccountUID(cloudAccountUID string) {
	o.CloudAccountUID = cloudAccountUID
}

// WithResourceGroup adds the resourceGroup to the v1 azure private Dns zones params
func (o *V1AzurePrivateDNSZonesParams) WithResourceGroup(resourceGroup string) *V1AzurePrivateDNSZonesParams {
	o.SetResourceGroup(resourceGroup)
	return o
}

// SetResourceGroup adds the resourceGroup to the v1 azure private Dns zones params
func (o *V1AzurePrivateDNSZonesParams) SetResourceGroup(resourceGroup string) {
	o.ResourceGroup = resourceGroup
}

// WithSubscriptionID adds the subscriptionID to the v1 azure private Dns zones params
func (o *V1AzurePrivateDNSZonesParams) WithSubscriptionID(subscriptionID string) *V1AzurePrivateDNSZonesParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the v1 azure private Dns zones params
func (o *V1AzurePrivateDNSZonesParams) SetSubscriptionID(subscriptionID string) {
	o.SubscriptionID = subscriptionID
}

// WriteToRequest writes these params to a swagger request
func (o *V1AzurePrivateDNSZonesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param resourceGroup
	if err := r.SetPathParam("resourceGroup", o.ResourceGroup); err != nil {
		return err
	}

	// query param subscriptionId
	qrSubscriptionID := o.SubscriptionID
	qSubscriptionID := qrSubscriptionID
	if qSubscriptionID != "" {
		if err := r.SetQueryParam("subscriptionId", qSubscriptionID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
