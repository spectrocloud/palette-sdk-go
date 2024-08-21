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

// NewV1WorkspacesCreateParams creates a new V1WorkspacesCreateParams object
// with the default values initialized.
func NewV1WorkspacesCreateParams() *V1WorkspacesCreateParams {
	var ()
	return &V1WorkspacesCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1WorkspacesCreateParamsWithTimeout creates a new V1WorkspacesCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1WorkspacesCreateParamsWithTimeout(timeout time.Duration) *V1WorkspacesCreateParams {
	var ()
	return &V1WorkspacesCreateParams{

		timeout: timeout,
	}
}

// NewV1WorkspacesCreateParamsWithContext creates a new V1WorkspacesCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1WorkspacesCreateParamsWithContext(ctx context.Context) *V1WorkspacesCreateParams {
	var ()
	return &V1WorkspacesCreateParams{

		Context: ctx,
	}
}

// NewV1WorkspacesCreateParamsWithHTTPClient creates a new V1WorkspacesCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1WorkspacesCreateParamsWithHTTPClient(client *http.Client) *V1WorkspacesCreateParams {
	var ()
	return &V1WorkspacesCreateParams{
		HTTPClient: client,
	}
}

/*
V1WorkspacesCreateParams contains all the parameters to send to the API endpoint
for the v1 workspaces create operation typically these are written to a http.Request
*/
type V1WorkspacesCreateParams struct {

	/*Body*/
	Body *models.V1WorkspaceEntity

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 workspaces create params
func (o *V1WorkspacesCreateParams) WithTimeout(timeout time.Duration) *V1WorkspacesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 workspaces create params
func (o *V1WorkspacesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 workspaces create params
func (o *V1WorkspacesCreateParams) WithContext(ctx context.Context) *V1WorkspacesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 workspaces create params
func (o *V1WorkspacesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 workspaces create params
func (o *V1WorkspacesCreateParams) WithHTTPClient(client *http.Client) *V1WorkspacesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 workspaces create params
func (o *V1WorkspacesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 workspaces create params
func (o *V1WorkspacesCreateParams) WithBody(body *models.V1WorkspaceEntity) *V1WorkspacesCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 workspaces create params
func (o *V1WorkspacesCreateParams) SetBody(body *models.V1WorkspaceEntity) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *V1WorkspacesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}