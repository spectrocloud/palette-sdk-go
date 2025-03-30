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

// NewV1AzureAccountValidateParams creates a new V1AzureAccountValidateParams object
// with the default values initialized.
func NewV1AzureAccountValidateParams() *V1AzureAccountValidateParams {
	var ()
	return &V1AzureAccountValidateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1AzureAccountValidateParamsWithTimeout creates a new V1AzureAccountValidateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1AzureAccountValidateParamsWithTimeout(timeout time.Duration) *V1AzureAccountValidateParams {
	var ()
	return &V1AzureAccountValidateParams{

		timeout: timeout,
	}
}

// NewV1AzureAccountValidateParamsWithContext creates a new V1AzureAccountValidateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1AzureAccountValidateParamsWithContext(ctx context.Context) *V1AzureAccountValidateParams {
	var ()
	return &V1AzureAccountValidateParams{

		Context: ctx,
	}
}

// NewV1AzureAccountValidateParamsWithHTTPClient creates a new V1AzureAccountValidateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1AzureAccountValidateParamsWithHTTPClient(client *http.Client) *V1AzureAccountValidateParams {
	var ()
	return &V1AzureAccountValidateParams{
		HTTPClient: client,
	}
}

/*
V1AzureAccountValidateParams contains all the parameters to send to the API endpoint
for the v1 azure account validate operation typically these are written to a http.Request
*/
type V1AzureAccountValidateParams struct {

	/*AzureCloudAccount
	  Request payload for Azure cloud account

	*/
	AzureCloudAccount *models.V1AzureCloudAccount

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 azure account validate params
func (o *V1AzureAccountValidateParams) WithTimeout(timeout time.Duration) *V1AzureAccountValidateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 azure account validate params
func (o *V1AzureAccountValidateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 azure account validate params
func (o *V1AzureAccountValidateParams) WithContext(ctx context.Context) *V1AzureAccountValidateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 azure account validate params
func (o *V1AzureAccountValidateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 azure account validate params
func (o *V1AzureAccountValidateParams) WithHTTPClient(client *http.Client) *V1AzureAccountValidateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 azure account validate params
func (o *V1AzureAccountValidateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAzureCloudAccount adds the azureCloudAccount to the v1 azure account validate params
func (o *V1AzureAccountValidateParams) WithAzureCloudAccount(azureCloudAccount *models.V1AzureCloudAccount) *V1AzureAccountValidateParams {
	o.SetAzureCloudAccount(azureCloudAccount)
	return o
}

// SetAzureCloudAccount adds the azureCloudAccount to the v1 azure account validate params
func (o *V1AzureAccountValidateParams) SetAzureCloudAccount(azureCloudAccount *models.V1AzureCloudAccount) {
	o.AzureCloudAccount = azureCloudAccount
}

// WriteToRequest writes these params to a swagger request
func (o *V1AzureAccountValidateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AzureCloudAccount != nil {
		if err := r.SetBodyParam(o.AzureCloudAccount); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
