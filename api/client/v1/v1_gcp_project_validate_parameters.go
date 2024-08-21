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

// NewV1GcpProjectValidateParams creates a new V1GcpProjectValidateParams object
// with the default values initialized.
func NewV1GcpProjectValidateParams() *V1GcpProjectValidateParams {
	var ()
	return &V1GcpProjectValidateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1GcpProjectValidateParamsWithTimeout creates a new V1GcpProjectValidateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1GcpProjectValidateParamsWithTimeout(timeout time.Duration) *V1GcpProjectValidateParams {
	var ()
	return &V1GcpProjectValidateParams{

		timeout: timeout,
	}
}

// NewV1GcpProjectValidateParamsWithContext creates a new V1GcpProjectValidateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1GcpProjectValidateParamsWithContext(ctx context.Context) *V1GcpProjectValidateParams {
	var ()
	return &V1GcpProjectValidateParams{

		Context: ctx,
	}
}

// NewV1GcpProjectValidateParamsWithHTTPClient creates a new V1GcpProjectValidateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1GcpProjectValidateParamsWithHTTPClient(client *http.Client) *V1GcpProjectValidateParams {
	var ()
	return &V1GcpProjectValidateParams{
		HTTPClient: client,
	}
}

/*
V1GcpProjectValidateParams contains all the parameters to send to the API endpoint
for the v1 gcp project validate operation typically these are written to a http.Request
*/
type V1GcpProjectValidateParams struct {

	/*CloudAccountUID
	  Uid for the specific GCP cloud account

	*/
	CloudAccountUID *models.V1CloudAccountUIDEntity
	/*Project
	  GCP project uid

	*/
	Project string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 gcp project validate params
func (o *V1GcpProjectValidateParams) WithTimeout(timeout time.Duration) *V1GcpProjectValidateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 gcp project validate params
func (o *V1GcpProjectValidateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 gcp project validate params
func (o *V1GcpProjectValidateParams) WithContext(ctx context.Context) *V1GcpProjectValidateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 gcp project validate params
func (o *V1GcpProjectValidateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 gcp project validate params
func (o *V1GcpProjectValidateParams) WithHTTPClient(client *http.Client) *V1GcpProjectValidateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 gcp project validate params
func (o *V1GcpProjectValidateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudAccountUID adds the cloudAccountUID to the v1 gcp project validate params
func (o *V1GcpProjectValidateParams) WithCloudAccountUID(cloudAccountUID *models.V1CloudAccountUIDEntity) *V1GcpProjectValidateParams {
	o.SetCloudAccountUID(cloudAccountUID)
	return o
}

// SetCloudAccountUID adds the cloudAccountUid to the v1 gcp project validate params
func (o *V1GcpProjectValidateParams) SetCloudAccountUID(cloudAccountUID *models.V1CloudAccountUIDEntity) {
	o.CloudAccountUID = cloudAccountUID
}

// WithProject adds the project to the v1 gcp project validate params
func (o *V1GcpProjectValidateParams) WithProject(project string) *V1GcpProjectValidateParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the v1 gcp project validate params
func (o *V1GcpProjectValidateParams) SetProject(project string) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *V1GcpProjectValidateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CloudAccountUID != nil {
		if err := r.SetBodyParam(o.CloudAccountUID); err != nil {
			return err
		}
	}

	// path param project
	if err := r.SetPathParam("project", o.Project); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}