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

// NewV1OpenStackAccountValidateParams creates a new V1OpenStackAccountValidateParams object
// with the default values initialized.
func NewV1OpenStackAccountValidateParams() *V1OpenStackAccountValidateParams {
	var ()
	return &V1OpenStackAccountValidateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1OpenStackAccountValidateParamsWithTimeout creates a new V1OpenStackAccountValidateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1OpenStackAccountValidateParamsWithTimeout(timeout time.Duration) *V1OpenStackAccountValidateParams {
	var ()
	return &V1OpenStackAccountValidateParams{

		timeout: timeout,
	}
}

// NewV1OpenStackAccountValidateParamsWithContext creates a new V1OpenStackAccountValidateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1OpenStackAccountValidateParamsWithContext(ctx context.Context) *V1OpenStackAccountValidateParams {
	var ()
	return &V1OpenStackAccountValidateParams{

		Context: ctx,
	}
}

// NewV1OpenStackAccountValidateParamsWithHTTPClient creates a new V1OpenStackAccountValidateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1OpenStackAccountValidateParamsWithHTTPClient(client *http.Client) *V1OpenStackAccountValidateParams {
	var ()
	return &V1OpenStackAccountValidateParams{
		HTTPClient: client,
	}
}

/*
V1OpenStackAccountValidateParams contains all the parameters to send to the API endpoint
for the v1 open stack account validate operation typically these are written to a http.Request
*/
type V1OpenStackAccountValidateParams struct {

	/*OpenstackCloudAccount
	  Request payload for OpenStack cloud account

	*/
	OpenstackCloudAccount *models.V1OpenStackCloudAccount

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 open stack account validate params
func (o *V1OpenStackAccountValidateParams) WithTimeout(timeout time.Duration) *V1OpenStackAccountValidateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 open stack account validate params
func (o *V1OpenStackAccountValidateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 open stack account validate params
func (o *V1OpenStackAccountValidateParams) WithContext(ctx context.Context) *V1OpenStackAccountValidateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 open stack account validate params
func (o *V1OpenStackAccountValidateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 open stack account validate params
func (o *V1OpenStackAccountValidateParams) WithHTTPClient(client *http.Client) *V1OpenStackAccountValidateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 open stack account validate params
func (o *V1OpenStackAccountValidateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOpenstackCloudAccount adds the openstackCloudAccount to the v1 open stack account validate params
func (o *V1OpenStackAccountValidateParams) WithOpenstackCloudAccount(openstackCloudAccount *models.V1OpenStackCloudAccount) *V1OpenStackAccountValidateParams {
	o.SetOpenstackCloudAccount(openstackCloudAccount)
	return o
}

// SetOpenstackCloudAccount adds the openstackCloudAccount to the v1 open stack account validate params
func (o *V1OpenStackAccountValidateParams) SetOpenstackCloudAccount(openstackCloudAccount *models.V1OpenStackCloudAccount) {
	o.OpenstackCloudAccount = openstackCloudAccount
}

// WriteToRequest writes these params to a swagger request
func (o *V1OpenStackAccountValidateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.OpenstackCloudAccount != nil {
		if err := r.SetBodyParam(o.OpenstackCloudAccount); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
