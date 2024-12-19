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

// NewV1AwsKmsKeysParams creates a new V1AwsKmsKeysParams object
// with the default values initialized.
func NewV1AwsKmsKeysParams() *V1AwsKmsKeysParams {
	var ()
	return &V1AwsKmsKeysParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1AwsKmsKeysParamsWithTimeout creates a new V1AwsKmsKeysParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1AwsKmsKeysParamsWithTimeout(timeout time.Duration) *V1AwsKmsKeysParams {
	var ()
	return &V1AwsKmsKeysParams{

		timeout: timeout,
	}
}

// NewV1AwsKmsKeysParamsWithContext creates a new V1AwsKmsKeysParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1AwsKmsKeysParamsWithContext(ctx context.Context) *V1AwsKmsKeysParams {
	var ()
	return &V1AwsKmsKeysParams{

		Context: ctx,
	}
}

// NewV1AwsKmsKeysParamsWithHTTPClient creates a new V1AwsKmsKeysParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1AwsKmsKeysParamsWithHTTPClient(client *http.Client) *V1AwsKmsKeysParams {
	var ()
	return &V1AwsKmsKeysParams{
		HTTPClient: client,
	}
}

/*
V1AwsKmsKeysParams contains all the parameters to send to the API endpoint
for the v1 aws kms keys operation typically these are written to a http.Request
*/
type V1AwsKmsKeysParams struct {

	/*CloudAccountUID
	  Uid for the specific AWS cloud account

	*/
	CloudAccountUID string
	/*Region
	  Region for which AWS KMS key are requested

	*/
	Region string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 aws kms keys params
func (o *V1AwsKmsKeysParams) WithTimeout(timeout time.Duration) *V1AwsKmsKeysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 aws kms keys params
func (o *V1AwsKmsKeysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 aws kms keys params
func (o *V1AwsKmsKeysParams) WithContext(ctx context.Context) *V1AwsKmsKeysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 aws kms keys params
func (o *V1AwsKmsKeysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 aws kms keys params
func (o *V1AwsKmsKeysParams) WithHTTPClient(client *http.Client) *V1AwsKmsKeysParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 aws kms keys params
func (o *V1AwsKmsKeysParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudAccountUID adds the cloudAccountUID to the v1 aws kms keys params
func (o *V1AwsKmsKeysParams) WithCloudAccountUID(cloudAccountUID string) *V1AwsKmsKeysParams {
	o.SetCloudAccountUID(cloudAccountUID)
	return o
}

// SetCloudAccountUID adds the cloudAccountUid to the v1 aws kms keys params
func (o *V1AwsKmsKeysParams) SetCloudAccountUID(cloudAccountUID string) {
	o.CloudAccountUID = cloudAccountUID
}

// WithRegion adds the region to the v1 aws kms keys params
func (o *V1AwsKmsKeysParams) WithRegion(region string) *V1AwsKmsKeysParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the v1 aws kms keys params
func (o *V1AwsKmsKeysParams) SetRegion(region string) {
	o.Region = region
}

// WriteToRequest writes these params to a swagger request
func (o *V1AwsKmsKeysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param region
	if err := r.SetPathParam("region", o.Region); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
