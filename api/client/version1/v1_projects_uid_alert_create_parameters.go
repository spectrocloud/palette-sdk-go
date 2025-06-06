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

// NewV1ProjectsUIDAlertCreateParams creates a new V1ProjectsUIDAlertCreateParams object
// with the default values initialized.
func NewV1ProjectsUIDAlertCreateParams() *V1ProjectsUIDAlertCreateParams {
	var ()
	return &V1ProjectsUIDAlertCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1ProjectsUIDAlertCreateParamsWithTimeout creates a new V1ProjectsUIDAlertCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1ProjectsUIDAlertCreateParamsWithTimeout(timeout time.Duration) *V1ProjectsUIDAlertCreateParams {
	var ()
	return &V1ProjectsUIDAlertCreateParams{

		timeout: timeout,
	}
}

// NewV1ProjectsUIDAlertCreateParamsWithContext creates a new V1ProjectsUIDAlertCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1ProjectsUIDAlertCreateParamsWithContext(ctx context.Context) *V1ProjectsUIDAlertCreateParams {
	var ()
	return &V1ProjectsUIDAlertCreateParams{

		Context: ctx,
	}
}

// NewV1ProjectsUIDAlertCreateParamsWithHTTPClient creates a new V1ProjectsUIDAlertCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1ProjectsUIDAlertCreateParamsWithHTTPClient(client *http.Client) *V1ProjectsUIDAlertCreateParams {
	var ()
	return &V1ProjectsUIDAlertCreateParams{
		HTTPClient: client,
	}
}

/*
V1ProjectsUIDAlertCreateParams contains all the parameters to send to the API endpoint
for the v1 projects Uid alert create operation typically these are written to a http.Request
*/
type V1ProjectsUIDAlertCreateParams struct {

	/*Body*/
	Body *models.V1Channel
	/*Component*/
	Component string
	/*UID*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 projects Uid alert create params
func (o *V1ProjectsUIDAlertCreateParams) WithTimeout(timeout time.Duration) *V1ProjectsUIDAlertCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 projects Uid alert create params
func (o *V1ProjectsUIDAlertCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 projects Uid alert create params
func (o *V1ProjectsUIDAlertCreateParams) WithContext(ctx context.Context) *V1ProjectsUIDAlertCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 projects Uid alert create params
func (o *V1ProjectsUIDAlertCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 projects Uid alert create params
func (o *V1ProjectsUIDAlertCreateParams) WithHTTPClient(client *http.Client) *V1ProjectsUIDAlertCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 projects Uid alert create params
func (o *V1ProjectsUIDAlertCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 projects Uid alert create params
func (o *V1ProjectsUIDAlertCreateParams) WithBody(body *models.V1Channel) *V1ProjectsUIDAlertCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 projects Uid alert create params
func (o *V1ProjectsUIDAlertCreateParams) SetBody(body *models.V1Channel) {
	o.Body = body
}

// WithComponent adds the component to the v1 projects Uid alert create params
func (o *V1ProjectsUIDAlertCreateParams) WithComponent(component string) *V1ProjectsUIDAlertCreateParams {
	o.SetComponent(component)
	return o
}

// SetComponent adds the component to the v1 projects Uid alert create params
func (o *V1ProjectsUIDAlertCreateParams) SetComponent(component string) {
	o.Component = component
}

// WithUID adds the uid to the v1 projects Uid alert create params
func (o *V1ProjectsUIDAlertCreateParams) WithUID(uid string) *V1ProjectsUIDAlertCreateParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 projects Uid alert create params
func (o *V1ProjectsUIDAlertCreateParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1ProjectsUIDAlertCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param component
	if err := r.SetPathParam("component", o.Component); err != nil {
		return err
	}

	// path param uid
	if err := r.SetPathParam("uid", o.UID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
