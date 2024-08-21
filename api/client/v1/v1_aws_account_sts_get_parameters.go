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

// NewV1AwsAccountStsGetParams creates a new V1AwsAccountStsGetParams object
// with the default values initialized.
func NewV1AwsAccountStsGetParams() *V1AwsAccountStsGetParams {
	var (
		partitionDefault = string("aws")
	)
	return &V1AwsAccountStsGetParams{
		Partition: &partitionDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewV1AwsAccountStsGetParamsWithTimeout creates a new V1AwsAccountStsGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1AwsAccountStsGetParamsWithTimeout(timeout time.Duration) *V1AwsAccountStsGetParams {
	var (
		partitionDefault = string("aws")
	)
	return &V1AwsAccountStsGetParams{
		Partition: &partitionDefault,

		timeout: timeout,
	}
}

// NewV1AwsAccountStsGetParamsWithContext creates a new V1AwsAccountStsGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1AwsAccountStsGetParamsWithContext(ctx context.Context) *V1AwsAccountStsGetParams {
	var (
		partitionDefault = string("aws")
	)
	return &V1AwsAccountStsGetParams{
		Partition: &partitionDefault,

		Context: ctx,
	}
}

// NewV1AwsAccountStsGetParamsWithHTTPClient creates a new V1AwsAccountStsGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1AwsAccountStsGetParamsWithHTTPClient(client *http.Client) *V1AwsAccountStsGetParams {
	var (
		partitionDefault = string("aws")
	)
	return &V1AwsAccountStsGetParams{
		Partition:  &partitionDefault,
		HTTPClient: client,
	}
}

/*
V1AwsAccountStsGetParams contains all the parameters to send to the API endpoint
for the v1 aws account sts get operation typically these are written to a http.Request
*/
type V1AwsAccountStsGetParams struct {

	/*Partition
	  AWS accounts are scoped to a single partition. Allowed values [aws, aws-us-gov], Default values

	*/
	Partition *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 aws account sts get params
func (o *V1AwsAccountStsGetParams) WithTimeout(timeout time.Duration) *V1AwsAccountStsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 aws account sts get params
func (o *V1AwsAccountStsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 aws account sts get params
func (o *V1AwsAccountStsGetParams) WithContext(ctx context.Context) *V1AwsAccountStsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 aws account sts get params
func (o *V1AwsAccountStsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 aws account sts get params
func (o *V1AwsAccountStsGetParams) WithHTTPClient(client *http.Client) *V1AwsAccountStsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 aws account sts get params
func (o *V1AwsAccountStsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPartition adds the partition to the v1 aws account sts get params
func (o *V1AwsAccountStsGetParams) WithPartition(partition *string) *V1AwsAccountStsGetParams {
	o.SetPartition(partition)
	return o
}

// SetPartition adds the partition to the v1 aws account sts get params
func (o *V1AwsAccountStsGetParams) SetPartition(partition *string) {
	o.Partition = partition
}

// WriteToRequest writes these params to a swagger request
func (o *V1AwsAccountStsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Partition != nil {

		// query param partition
		var qrPartition string
		if o.Partition != nil {
			qrPartition = *o.Partition
		}
		qPartition := qrPartition
		if qPartition != "" {
			if err := r.SetQueryParam("partition", qPartition); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}