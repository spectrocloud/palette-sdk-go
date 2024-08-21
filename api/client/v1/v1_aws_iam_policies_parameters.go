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

	"github.com/spectrocloud/palette-sdk-go/api/models"
)

// NewV1AwsIamPoliciesParams creates a new V1AwsIamPoliciesParams object
// with the default values initialized.
func NewV1AwsIamPoliciesParams() *V1AwsIamPoliciesParams {
	var ()
	return &V1AwsIamPoliciesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1AwsIamPoliciesParamsWithTimeout creates a new V1AwsIamPoliciesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1AwsIamPoliciesParamsWithTimeout(timeout time.Duration) *V1AwsIamPoliciesParams {
	var ()
	return &V1AwsIamPoliciesParams{

		timeout: timeout,
	}
}

// NewV1AwsIamPoliciesParamsWithContext creates a new V1AwsIamPoliciesParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1AwsIamPoliciesParamsWithContext(ctx context.Context) *V1AwsIamPoliciesParams {
	var ()
	return &V1AwsIamPoliciesParams{

		Context: ctx,
	}
}

// NewV1AwsIamPoliciesParamsWithHTTPClient creates a new V1AwsIamPoliciesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1AwsIamPoliciesParamsWithHTTPClient(client *http.Client) *V1AwsIamPoliciesParams {
	var ()
	return &V1AwsIamPoliciesParams{
		HTTPClient: client,
	}
}

/*
V1AwsIamPoliciesParams contains all the parameters to send to the API endpoint
for the v1 aws iam policies operation typically these are written to a http.Request
*/
type V1AwsIamPoliciesParams struct {

	/*Account
	  Request payload for AWS Cloud Account

	*/
	Account *models.V1AwsCloudAccount
	/*CloudAccountUID
	  Uid for the specific AWS cloud account

	*/
	CloudAccountUID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 aws iam policies params
func (o *V1AwsIamPoliciesParams) WithTimeout(timeout time.Duration) *V1AwsIamPoliciesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 aws iam policies params
func (o *V1AwsIamPoliciesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 aws iam policies params
func (o *V1AwsIamPoliciesParams) WithContext(ctx context.Context) *V1AwsIamPoliciesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 aws iam policies params
func (o *V1AwsIamPoliciesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 aws iam policies params
func (o *V1AwsIamPoliciesParams) WithHTTPClient(client *http.Client) *V1AwsIamPoliciesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 aws iam policies params
func (o *V1AwsIamPoliciesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccount adds the account to the v1 aws iam policies params
func (o *V1AwsIamPoliciesParams) WithAccount(account *models.V1AwsCloudAccount) *V1AwsIamPoliciesParams {
	o.SetAccount(account)
	return o
}

// SetAccount adds the account to the v1 aws iam policies params
func (o *V1AwsIamPoliciesParams) SetAccount(account *models.V1AwsCloudAccount) {
	o.Account = account
}

// WithCloudAccountUID adds the cloudAccountUID to the v1 aws iam policies params
func (o *V1AwsIamPoliciesParams) WithCloudAccountUID(cloudAccountUID *string) *V1AwsIamPoliciesParams {
	o.SetCloudAccountUID(cloudAccountUID)
	return o
}

// SetCloudAccountUID adds the cloudAccountUid to the v1 aws iam policies params
func (o *V1AwsIamPoliciesParams) SetCloudAccountUID(cloudAccountUID *string) {
	o.CloudAccountUID = cloudAccountUID
}

// WriteToRequest writes these params to a swagger request
func (o *V1AwsIamPoliciesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Account != nil {
		if err := r.SetBodyParam(o.Account); err != nil {
			return err
		}
	}

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