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

// NewV1GcpPropertiesValidateParams creates a new V1GcpPropertiesValidateParams object
// with the default values initialized.
func NewV1GcpPropertiesValidateParams() *V1GcpPropertiesValidateParams {
	var ()
	return &V1GcpPropertiesValidateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1GcpPropertiesValidateParamsWithTimeout creates a new V1GcpPropertiesValidateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1GcpPropertiesValidateParamsWithTimeout(timeout time.Duration) *V1GcpPropertiesValidateParams {
	var ()
	return &V1GcpPropertiesValidateParams{

		timeout: timeout,
	}
}

// NewV1GcpPropertiesValidateParamsWithContext creates a new V1GcpPropertiesValidateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1GcpPropertiesValidateParamsWithContext(ctx context.Context) *V1GcpPropertiesValidateParams {
	var ()
	return &V1GcpPropertiesValidateParams{

		Context: ctx,
	}
}

// NewV1GcpPropertiesValidateParamsWithHTTPClient creates a new V1GcpPropertiesValidateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1GcpPropertiesValidateParamsWithHTTPClient(client *http.Client) *V1GcpPropertiesValidateParams {
	var ()
	return &V1GcpPropertiesValidateParams{
		HTTPClient: client,
	}
}

/*
V1GcpPropertiesValidateParams contains all the parameters to send to the API endpoint
for the v1 gcp properties validate operation typically these are written to a http.Request
*/
type V1GcpPropertiesValidateParams struct {

	/*Properties
	  Request payload for GCP properties validate spec

	*/
	Properties *models.V1GcpPropertiesValidateSpec

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 gcp properties validate params
func (o *V1GcpPropertiesValidateParams) WithTimeout(timeout time.Duration) *V1GcpPropertiesValidateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 gcp properties validate params
func (o *V1GcpPropertiesValidateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 gcp properties validate params
func (o *V1GcpPropertiesValidateParams) WithContext(ctx context.Context) *V1GcpPropertiesValidateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 gcp properties validate params
func (o *V1GcpPropertiesValidateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 gcp properties validate params
func (o *V1GcpPropertiesValidateParams) WithHTTPClient(client *http.Client) *V1GcpPropertiesValidateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 gcp properties validate params
func (o *V1GcpPropertiesValidateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProperties adds the properties to the v1 gcp properties validate params
func (o *V1GcpPropertiesValidateParams) WithProperties(properties *models.V1GcpPropertiesValidateSpec) *V1GcpPropertiesValidateParams {
	o.SetProperties(properties)
	return o
}

// SetProperties adds the properties to the v1 gcp properties validate params
func (o *V1GcpPropertiesValidateParams) SetProperties(properties *models.V1GcpPropertiesValidateSpec) {
	o.Properties = properties
}

// WriteToRequest writes these params to a swagger request
func (o *V1GcpPropertiesValidateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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