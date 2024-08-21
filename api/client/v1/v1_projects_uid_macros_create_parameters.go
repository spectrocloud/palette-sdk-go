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

// NewV1ProjectsUIDMacrosCreateParams creates a new V1ProjectsUIDMacrosCreateParams object
// with the default values initialized.
func NewV1ProjectsUIDMacrosCreateParams() *V1ProjectsUIDMacrosCreateParams {
	var ()
	return &V1ProjectsUIDMacrosCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1ProjectsUIDMacrosCreateParamsWithTimeout creates a new V1ProjectsUIDMacrosCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1ProjectsUIDMacrosCreateParamsWithTimeout(timeout time.Duration) *V1ProjectsUIDMacrosCreateParams {
	var ()
	return &V1ProjectsUIDMacrosCreateParams{

		timeout: timeout,
	}
}

// NewV1ProjectsUIDMacrosCreateParamsWithContext creates a new V1ProjectsUIDMacrosCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1ProjectsUIDMacrosCreateParamsWithContext(ctx context.Context) *V1ProjectsUIDMacrosCreateParams {
	var ()
	return &V1ProjectsUIDMacrosCreateParams{

		Context: ctx,
	}
}

// NewV1ProjectsUIDMacrosCreateParamsWithHTTPClient creates a new V1ProjectsUIDMacrosCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1ProjectsUIDMacrosCreateParamsWithHTTPClient(client *http.Client) *V1ProjectsUIDMacrosCreateParams {
	var ()
	return &V1ProjectsUIDMacrosCreateParams{
		HTTPClient: client,
	}
}

/*
V1ProjectsUIDMacrosCreateParams contains all the parameters to send to the API endpoint
for the v1 projects Uid macros create operation typically these are written to a http.Request
*/
type V1ProjectsUIDMacrosCreateParams struct {

	/*Body*/
	Body *models.V1Macros
	/*UID*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 projects Uid macros create params
func (o *V1ProjectsUIDMacrosCreateParams) WithTimeout(timeout time.Duration) *V1ProjectsUIDMacrosCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 projects Uid macros create params
func (o *V1ProjectsUIDMacrosCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 projects Uid macros create params
func (o *V1ProjectsUIDMacrosCreateParams) WithContext(ctx context.Context) *V1ProjectsUIDMacrosCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 projects Uid macros create params
func (o *V1ProjectsUIDMacrosCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 projects Uid macros create params
func (o *V1ProjectsUIDMacrosCreateParams) WithHTTPClient(client *http.Client) *V1ProjectsUIDMacrosCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 projects Uid macros create params
func (o *V1ProjectsUIDMacrosCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 projects Uid macros create params
func (o *V1ProjectsUIDMacrosCreateParams) WithBody(body *models.V1Macros) *V1ProjectsUIDMacrosCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 projects Uid macros create params
func (o *V1ProjectsUIDMacrosCreateParams) SetBody(body *models.V1Macros) {
	o.Body = body
}

// WithUID adds the uid to the v1 projects Uid macros create params
func (o *V1ProjectsUIDMacrosCreateParams) WithUID(uid string) *V1ProjectsUIDMacrosCreateParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 projects Uid macros create params
func (o *V1ProjectsUIDMacrosCreateParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1ProjectsUIDMacrosCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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