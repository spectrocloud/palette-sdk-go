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

// NewV1AwsSecurityGroupsParams creates a new V1AwsSecurityGroupsParams object
// with the default values initialized.
func NewV1AwsSecurityGroupsParams() *V1AwsSecurityGroupsParams {
	var ()
	return &V1AwsSecurityGroupsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1AwsSecurityGroupsParamsWithTimeout creates a new V1AwsSecurityGroupsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1AwsSecurityGroupsParamsWithTimeout(timeout time.Duration) *V1AwsSecurityGroupsParams {
	var ()
	return &V1AwsSecurityGroupsParams{

		timeout: timeout,
	}
}

// NewV1AwsSecurityGroupsParamsWithContext creates a new V1AwsSecurityGroupsParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1AwsSecurityGroupsParamsWithContext(ctx context.Context) *V1AwsSecurityGroupsParams {
	var ()
	return &V1AwsSecurityGroupsParams{

		Context: ctx,
	}
}

// NewV1AwsSecurityGroupsParamsWithHTTPClient creates a new V1AwsSecurityGroupsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1AwsSecurityGroupsParamsWithHTTPClient(client *http.Client) *V1AwsSecurityGroupsParams {
	var ()
	return &V1AwsSecurityGroupsParams{
		HTTPClient: client,
	}
}

/*
V1AwsSecurityGroupsParams contains all the parameters to send to the API endpoint
for the v1 aws security groups operation typically these are written to a http.Request
*/
type V1AwsSecurityGroupsParams struct {

	/*CloudAccountUID
	  Uid for the specific AWS cloud account

	*/
	CloudAccountUID *string
	/*Region
	  Region for which security groups are requested

	*/
	Region string
	/*VpcID
	  Vpc Id for which security groups are requested

	*/
	VpcID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 aws security groups params
func (o *V1AwsSecurityGroupsParams) WithTimeout(timeout time.Duration) *V1AwsSecurityGroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 aws security groups params
func (o *V1AwsSecurityGroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 aws security groups params
func (o *V1AwsSecurityGroupsParams) WithContext(ctx context.Context) *V1AwsSecurityGroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 aws security groups params
func (o *V1AwsSecurityGroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 aws security groups params
func (o *V1AwsSecurityGroupsParams) WithHTTPClient(client *http.Client) *V1AwsSecurityGroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 aws security groups params
func (o *V1AwsSecurityGroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudAccountUID adds the cloudAccountUID to the v1 aws security groups params
func (o *V1AwsSecurityGroupsParams) WithCloudAccountUID(cloudAccountUID *string) *V1AwsSecurityGroupsParams {
	o.SetCloudAccountUID(cloudAccountUID)
	return o
}

// SetCloudAccountUID adds the cloudAccountUid to the v1 aws security groups params
func (o *V1AwsSecurityGroupsParams) SetCloudAccountUID(cloudAccountUID *string) {
	o.CloudAccountUID = cloudAccountUID
}

// WithRegion adds the region to the v1 aws security groups params
func (o *V1AwsSecurityGroupsParams) WithRegion(region string) *V1AwsSecurityGroupsParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the v1 aws security groups params
func (o *V1AwsSecurityGroupsParams) SetRegion(region string) {
	o.Region = region
}

// WithVpcID adds the vpcID to the v1 aws security groups params
func (o *V1AwsSecurityGroupsParams) WithVpcID(vpcID string) *V1AwsSecurityGroupsParams {
	o.SetVpcID(vpcID)
	return o
}

// SetVpcID adds the vpcId to the v1 aws security groups params
func (o *V1AwsSecurityGroupsParams) SetVpcID(vpcID string) {
	o.VpcID = vpcID
}

// WriteToRequest writes these params to a swagger request
func (o *V1AwsSecurityGroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// query param region
	qrRegion := o.Region
	qRegion := qrRegion
	if qRegion != "" {
		if err := r.SetQueryParam("region", qRegion); err != nil {
			return err
		}
	}

	// query param vpcId
	qrVpcID := o.VpcID
	qVpcID := qrVpcID
	if qVpcID != "" {
		if err := r.SetQueryParam("vpcId", qVpcID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
