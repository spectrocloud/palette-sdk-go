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

// NewV1AzureRegionsParams creates a new V1AzureRegionsParams object
// with the default values initialized.
func NewV1AzureRegionsParams() *V1AzureRegionsParams {
	var ()
	return &V1AzureRegionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1AzureRegionsParamsWithTimeout creates a new V1AzureRegionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1AzureRegionsParamsWithTimeout(timeout time.Duration) *V1AzureRegionsParams {
	var ()
	return &V1AzureRegionsParams{

		timeout: timeout,
	}
}

// NewV1AzureRegionsParamsWithContext creates a new V1AzureRegionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1AzureRegionsParamsWithContext(ctx context.Context) *V1AzureRegionsParams {
	var ()
	return &V1AzureRegionsParams{

		Context: ctx,
	}
}

// NewV1AzureRegionsParamsWithHTTPClient creates a new V1AzureRegionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1AzureRegionsParamsWithHTTPClient(client *http.Client) *V1AzureRegionsParams {
	var ()
	return &V1AzureRegionsParams{
		HTTPClient: client,
	}
}

/*V1AzureRegionsParams contains all the parameters to send to the API endpoint
for the v1 azure regions operation typically these are written to a http.Request
*/
type V1AzureRegionsParams struct {

	/*CloudAccountUID
	  Uid for the specific Azure cloud account

	*/
	CloudAccountUID *string
	/*SubscriptionID
	  SubscriptionId for which resources is requested

	*/
	SubscriptionID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 azure regions params
func (o *V1AzureRegionsParams) WithTimeout(timeout time.Duration) *V1AzureRegionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 azure regions params
func (o *V1AzureRegionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 azure regions params
func (o *V1AzureRegionsParams) WithContext(ctx context.Context) *V1AzureRegionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 azure regions params
func (o *V1AzureRegionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 azure regions params
func (o *V1AzureRegionsParams) WithHTTPClient(client *http.Client) *V1AzureRegionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 azure regions params
func (o *V1AzureRegionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudAccountUID adds the cloudAccountUID to the v1 azure regions params
func (o *V1AzureRegionsParams) WithCloudAccountUID(cloudAccountUID *string) *V1AzureRegionsParams {
	o.SetCloudAccountUID(cloudAccountUID)
	return o
}

// SetCloudAccountUID adds the cloudAccountUid to the v1 azure regions params
func (o *V1AzureRegionsParams) SetCloudAccountUID(cloudAccountUID *string) {
	o.CloudAccountUID = cloudAccountUID
}

// WithSubscriptionID adds the subscriptionID to the v1 azure regions params
func (o *V1AzureRegionsParams) WithSubscriptionID(subscriptionID *string) *V1AzureRegionsParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the v1 azure regions params
func (o *V1AzureRegionsParams) SetSubscriptionID(subscriptionID *string) {
	o.SubscriptionID = subscriptionID
}

// WriteToRequest writes these params to a swagger request
func (o *V1AzureRegionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.SubscriptionID != nil {

		// query param subscriptionId
		var qrSubscriptionID string
		if o.SubscriptionID != nil {
			qrSubscriptionID = *o.SubscriptionID
		}
		qSubscriptionID := qrSubscriptionID
		if qSubscriptionID != "" {
			if err := r.SetQueryParam("subscriptionId", qSubscriptionID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
