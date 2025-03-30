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

// NewV1UsersAssetsLocationGcpCreateParams creates a new V1UsersAssetsLocationGcpCreateParams object
// with the default values initialized.
func NewV1UsersAssetsLocationGcpCreateParams() *V1UsersAssetsLocationGcpCreateParams {
	var ()
	return &V1UsersAssetsLocationGcpCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1UsersAssetsLocationGcpCreateParamsWithTimeout creates a new V1UsersAssetsLocationGcpCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1UsersAssetsLocationGcpCreateParamsWithTimeout(timeout time.Duration) *V1UsersAssetsLocationGcpCreateParams {
	var ()
	return &V1UsersAssetsLocationGcpCreateParams{

		timeout: timeout,
	}
}

// NewV1UsersAssetsLocationGcpCreateParamsWithContext creates a new V1UsersAssetsLocationGcpCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1UsersAssetsLocationGcpCreateParamsWithContext(ctx context.Context) *V1UsersAssetsLocationGcpCreateParams {
	var ()
	return &V1UsersAssetsLocationGcpCreateParams{

		Context: ctx,
	}
}

// NewV1UsersAssetsLocationGcpCreateParamsWithHTTPClient creates a new V1UsersAssetsLocationGcpCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1UsersAssetsLocationGcpCreateParamsWithHTTPClient(client *http.Client) *V1UsersAssetsLocationGcpCreateParams {
	var ()
	return &V1UsersAssetsLocationGcpCreateParams{
		HTTPClient: client,
	}
}

/*
V1UsersAssetsLocationGcpCreateParams contains all the parameters to send to the API endpoint
for the v1 users assets location gcp create operation typically these are written to a http.Request
*/
type V1UsersAssetsLocationGcpCreateParams struct {

	/*Body*/
	Body *models.V1UserAssetsLocationGcp

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 users assets location gcp create params
func (o *V1UsersAssetsLocationGcpCreateParams) WithTimeout(timeout time.Duration) *V1UsersAssetsLocationGcpCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 users assets location gcp create params
func (o *V1UsersAssetsLocationGcpCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 users assets location gcp create params
func (o *V1UsersAssetsLocationGcpCreateParams) WithContext(ctx context.Context) *V1UsersAssetsLocationGcpCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 users assets location gcp create params
func (o *V1UsersAssetsLocationGcpCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 users assets location gcp create params
func (o *V1UsersAssetsLocationGcpCreateParams) WithHTTPClient(client *http.Client) *V1UsersAssetsLocationGcpCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 users assets location gcp create params
func (o *V1UsersAssetsLocationGcpCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 users assets location gcp create params
func (o *V1UsersAssetsLocationGcpCreateParams) WithBody(body *models.V1UserAssetsLocationGcp) *V1UsersAssetsLocationGcpCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 users assets location gcp create params
func (o *V1UsersAssetsLocationGcpCreateParams) SetBody(body *models.V1UserAssetsLocationGcp) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *V1UsersAssetsLocationGcpCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
