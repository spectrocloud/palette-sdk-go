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

// NewV1GcpContainerImageValidateParams creates a new V1GcpContainerImageValidateParams object
// with the default values initialized.
func NewV1GcpContainerImageValidateParams() *V1GcpContainerImageValidateParams {
	var ()
	return &V1GcpContainerImageValidateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1GcpContainerImageValidateParamsWithTimeout creates a new V1GcpContainerImageValidateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1GcpContainerImageValidateParamsWithTimeout(timeout time.Duration) *V1GcpContainerImageValidateParams {
	var ()
	return &V1GcpContainerImageValidateParams{

		timeout: timeout,
	}
}

// NewV1GcpContainerImageValidateParamsWithContext creates a new V1GcpContainerImageValidateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1GcpContainerImageValidateParamsWithContext(ctx context.Context) *V1GcpContainerImageValidateParams {
	var ()
	return &V1GcpContainerImageValidateParams{

		Context: ctx,
	}
}

// NewV1GcpContainerImageValidateParamsWithHTTPClient creates a new V1GcpContainerImageValidateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1GcpContainerImageValidateParamsWithHTTPClient(client *http.Client) *V1GcpContainerImageValidateParams {
	var ()
	return &V1GcpContainerImageValidateParams{
		HTTPClient: client,
	}
}

/*
V1GcpContainerImageValidateParams contains all the parameters to send to the API endpoint
for the v1 gcp container image validate operation typically these are written to a http.Request
*/
type V1GcpContainerImageValidateParams struct {

	/*ImagePath
	  image path in the container

	*/
	ImagePath string
	/*Tag
	  tag in the GCP container

	*/
	Tag string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 gcp container image validate params
func (o *V1GcpContainerImageValidateParams) WithTimeout(timeout time.Duration) *V1GcpContainerImageValidateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 gcp container image validate params
func (o *V1GcpContainerImageValidateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 gcp container image validate params
func (o *V1GcpContainerImageValidateParams) WithContext(ctx context.Context) *V1GcpContainerImageValidateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 gcp container image validate params
func (o *V1GcpContainerImageValidateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 gcp container image validate params
func (o *V1GcpContainerImageValidateParams) WithHTTPClient(client *http.Client) *V1GcpContainerImageValidateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 gcp container image validate params
func (o *V1GcpContainerImageValidateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImagePath adds the imagePath to the v1 gcp container image validate params
func (o *V1GcpContainerImageValidateParams) WithImagePath(imagePath string) *V1GcpContainerImageValidateParams {
	o.SetImagePath(imagePath)
	return o
}

// SetImagePath adds the imagePath to the v1 gcp container image validate params
func (o *V1GcpContainerImageValidateParams) SetImagePath(imagePath string) {
	o.ImagePath = imagePath
}

// WithTag adds the tag to the v1 gcp container image validate params
func (o *V1GcpContainerImageValidateParams) WithTag(tag string) *V1GcpContainerImageValidateParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the v1 gcp container image validate params
func (o *V1GcpContainerImageValidateParams) SetTag(tag string) {
	o.Tag = tag
}

// WriteToRequest writes these params to a swagger request
func (o *V1GcpContainerImageValidateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param imagePath
	qrImagePath := o.ImagePath
	qImagePath := qrImagePath
	if qImagePath != "" {
		if err := r.SetQueryParam("imagePath", qImagePath); err != nil {
			return err
		}
	}

	// query param tag
	qrTag := o.Tag
	qTag := qrTag
	if qTag != "" {
		if err := r.SetQueryParam("tag", qTag); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
