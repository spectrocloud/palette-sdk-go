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

// NewV1CloudAccountsCustomCreateParams creates a new V1CloudAccountsCustomCreateParams object
// with the default values initialized.
func NewV1CloudAccountsCustomCreateParams() *V1CloudAccountsCustomCreateParams {
	var ()
	return &V1CloudAccountsCustomCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudAccountsCustomCreateParamsWithTimeout creates a new V1CloudAccountsCustomCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudAccountsCustomCreateParamsWithTimeout(timeout time.Duration) *V1CloudAccountsCustomCreateParams {
	var ()
	return &V1CloudAccountsCustomCreateParams{

		timeout: timeout,
	}
}

// NewV1CloudAccountsCustomCreateParamsWithContext creates a new V1CloudAccountsCustomCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudAccountsCustomCreateParamsWithContext(ctx context.Context) *V1CloudAccountsCustomCreateParams {
	var ()
	return &V1CloudAccountsCustomCreateParams{

		Context: ctx,
	}
}

// NewV1CloudAccountsCustomCreateParamsWithHTTPClient creates a new V1CloudAccountsCustomCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudAccountsCustomCreateParamsWithHTTPClient(client *http.Client) *V1CloudAccountsCustomCreateParams {
	var ()
	return &V1CloudAccountsCustomCreateParams{
		HTTPClient: client,
	}
}

/*
V1CloudAccountsCustomCreateParams contains all the parameters to send to the API endpoint
for the v1 cloud accounts custom create operation typically these are written to a http.Request
*/
type V1CloudAccountsCustomCreateParams struct {

	/*Body
	  Request payload to validate Custom cloud account

	*/
	Body *models.V1CustomAccountEntity
	/*CloudType
	  Custom cloud type

	*/
	CloudType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cloud accounts custom create params
func (o *V1CloudAccountsCustomCreateParams) WithTimeout(timeout time.Duration) *V1CloudAccountsCustomCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud accounts custom create params
func (o *V1CloudAccountsCustomCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud accounts custom create params
func (o *V1CloudAccountsCustomCreateParams) WithContext(ctx context.Context) *V1CloudAccountsCustomCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud accounts custom create params
func (o *V1CloudAccountsCustomCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud accounts custom create params
func (o *V1CloudAccountsCustomCreateParams) WithHTTPClient(client *http.Client) *V1CloudAccountsCustomCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud accounts custom create params
func (o *V1CloudAccountsCustomCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 cloud accounts custom create params
func (o *V1CloudAccountsCustomCreateParams) WithBody(body *models.V1CustomAccountEntity) *V1CloudAccountsCustomCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 cloud accounts custom create params
func (o *V1CloudAccountsCustomCreateParams) SetBody(body *models.V1CustomAccountEntity) {
	o.Body = body
}

// WithCloudType adds the cloudType to the v1 cloud accounts custom create params
func (o *V1CloudAccountsCustomCreateParams) WithCloudType(cloudType string) *V1CloudAccountsCustomCreateParams {
	o.SetCloudType(cloudType)
	return o
}

// SetCloudType adds the cloudType to the v1 cloud accounts custom create params
func (o *V1CloudAccountsCustomCreateParams) SetCloudType(cloudType string) {
	o.CloudType = cloudType
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudAccountsCustomCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param cloudType
	if err := r.SetPathParam("cloudType", o.CloudType); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}