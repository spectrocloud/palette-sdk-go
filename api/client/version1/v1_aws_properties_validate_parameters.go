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

	"github.com/spectrocloud/palette-sdk-go/api/models"
)

// NewV1AwsPropertiesValidateParams creates a new V1AwsPropertiesValidateParams object
// with the default values initialized.
func NewV1AwsPropertiesValidateParams() *V1AwsPropertiesValidateParams {
	var ()
	return &V1AwsPropertiesValidateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1AwsPropertiesValidateParamsWithTimeout creates a new V1AwsPropertiesValidateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1AwsPropertiesValidateParamsWithTimeout(timeout time.Duration) *V1AwsPropertiesValidateParams {
	var ()
	return &V1AwsPropertiesValidateParams{

		timeout: timeout,
	}
}

// NewV1AwsPropertiesValidateParamsWithContext creates a new V1AwsPropertiesValidateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1AwsPropertiesValidateParamsWithContext(ctx context.Context) *V1AwsPropertiesValidateParams {
	var ()
	return &V1AwsPropertiesValidateParams{

		Context: ctx,
	}
}

// NewV1AwsPropertiesValidateParamsWithHTTPClient creates a new V1AwsPropertiesValidateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1AwsPropertiesValidateParamsWithHTTPClient(client *http.Client) *V1AwsPropertiesValidateParams {
	var ()
	return &V1AwsPropertiesValidateParams{
		HTTPClient: client,
	}
}

/*
V1AwsPropertiesValidateParams contains all the parameters to send to the API endpoint
for the v1 aws properties validate operation typically these are written to a http.Request
*/
type V1AwsPropertiesValidateParams struct {

	/*Properties
	  Request payload for AWS properties validate spec

	*/
	Properties *models.V1AwsPropertiesValidateSpec

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 aws properties validate params
func (o *V1AwsPropertiesValidateParams) WithTimeout(timeout time.Duration) *V1AwsPropertiesValidateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 aws properties validate params
func (o *V1AwsPropertiesValidateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 aws properties validate params
func (o *V1AwsPropertiesValidateParams) WithContext(ctx context.Context) *V1AwsPropertiesValidateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 aws properties validate params
func (o *V1AwsPropertiesValidateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 aws properties validate params
func (o *V1AwsPropertiesValidateParams) WithHTTPClient(client *http.Client) *V1AwsPropertiesValidateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 aws properties validate params
func (o *V1AwsPropertiesValidateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProperties adds the properties to the v1 aws properties validate params
func (o *V1AwsPropertiesValidateParams) WithProperties(properties *models.V1AwsPropertiesValidateSpec) *V1AwsPropertiesValidateParams {
	o.SetProperties(properties)
	return o
}

// SetProperties adds the properties to the v1 aws properties validate params
func (o *V1AwsPropertiesValidateParams) SetProperties(properties *models.V1AwsPropertiesValidateSpec) {
	o.Properties = properties
}

// WriteToRequest writes these params to a swagger request
func (o *V1AwsPropertiesValidateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Properties != nil {
		if err := r.SetBodyParam(o.Properties); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
