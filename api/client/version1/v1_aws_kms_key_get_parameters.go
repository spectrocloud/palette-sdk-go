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

// NewV1AwsKmsKeyGetParams creates a new V1AwsKmsKeyGetParams object
// with the default values initialized.
func NewV1AwsKmsKeyGetParams() *V1AwsKmsKeyGetParams {
	var ()
	return &V1AwsKmsKeyGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1AwsKmsKeyGetParamsWithTimeout creates a new V1AwsKmsKeyGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1AwsKmsKeyGetParamsWithTimeout(timeout time.Duration) *V1AwsKmsKeyGetParams {
	var ()
	return &V1AwsKmsKeyGetParams{

		timeout: timeout,
	}
}

// NewV1AwsKmsKeyGetParamsWithContext creates a new V1AwsKmsKeyGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1AwsKmsKeyGetParamsWithContext(ctx context.Context) *V1AwsKmsKeyGetParams {
	var ()
	return &V1AwsKmsKeyGetParams{

		Context: ctx,
	}
}

// NewV1AwsKmsKeyGetParamsWithHTTPClient creates a new V1AwsKmsKeyGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1AwsKmsKeyGetParamsWithHTTPClient(client *http.Client) *V1AwsKmsKeyGetParams {
	var ()
	return &V1AwsKmsKeyGetParams{
		HTTPClient: client,
	}
}

/*
V1AwsKmsKeyGetParams contains all the parameters to send to the API endpoint
for the v1 aws kms key get operation typically these are written to a http.Request
*/
type V1AwsKmsKeyGetParams struct {

	/*CloudAccountUID
	  Uid for the specific AWS cloud account

	*/
	CloudAccountUID string
	/*KeyID
	  The globally unique identifier for the KMS key

	*/
	KeyID string
	/*Region
	  Region for which AWS KMS key belongs

	*/
	Region string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 aws kms key get params
func (o *V1AwsKmsKeyGetParams) WithTimeout(timeout time.Duration) *V1AwsKmsKeyGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 aws kms key get params
func (o *V1AwsKmsKeyGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 aws kms key get params
func (o *V1AwsKmsKeyGetParams) WithContext(ctx context.Context) *V1AwsKmsKeyGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 aws kms key get params
func (o *V1AwsKmsKeyGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 aws kms key get params
func (o *V1AwsKmsKeyGetParams) WithHTTPClient(client *http.Client) *V1AwsKmsKeyGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 aws kms key get params
func (o *V1AwsKmsKeyGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudAccountUID adds the cloudAccountUID to the v1 aws kms key get params
func (o *V1AwsKmsKeyGetParams) WithCloudAccountUID(cloudAccountUID string) *V1AwsKmsKeyGetParams {
	o.SetCloudAccountUID(cloudAccountUID)
	return o
}

// SetCloudAccountUID adds the cloudAccountUid to the v1 aws kms key get params
func (o *V1AwsKmsKeyGetParams) SetCloudAccountUID(cloudAccountUID string) {
	o.CloudAccountUID = cloudAccountUID
}

// WithKeyID adds the keyID to the v1 aws kms key get params
func (o *V1AwsKmsKeyGetParams) WithKeyID(keyID string) *V1AwsKmsKeyGetParams {
	o.SetKeyID(keyID)
	return o
}

// SetKeyID adds the keyId to the v1 aws kms key get params
func (o *V1AwsKmsKeyGetParams) SetKeyID(keyID string) {
	o.KeyID = keyID
}

// WithRegion adds the region to the v1 aws kms key get params
func (o *V1AwsKmsKeyGetParams) WithRegion(region string) *V1AwsKmsKeyGetParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the v1 aws kms key get params
func (o *V1AwsKmsKeyGetParams) SetRegion(region string) {
	o.Region = region
}

// WriteToRequest writes these params to a swagger request
func (o *V1AwsKmsKeyGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param keyId
	if err := r.SetPathParam("keyId", o.KeyID); err != nil {
		return err
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
