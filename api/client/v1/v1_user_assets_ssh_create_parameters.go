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

// NewV1UserAssetsSSHCreateParams creates a new V1UserAssetsSSHCreateParams object
// with the default values initialized.
func NewV1UserAssetsSSHCreateParams() *V1UserAssetsSSHCreateParams {
	var ()
	return &V1UserAssetsSSHCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1UserAssetsSSHCreateParamsWithTimeout creates a new V1UserAssetsSSHCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1UserAssetsSSHCreateParamsWithTimeout(timeout time.Duration) *V1UserAssetsSSHCreateParams {
	var ()
	return &V1UserAssetsSSHCreateParams{

		timeout: timeout,
	}
}

// NewV1UserAssetsSSHCreateParamsWithContext creates a new V1UserAssetsSSHCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1UserAssetsSSHCreateParamsWithContext(ctx context.Context) *V1UserAssetsSSHCreateParams {
	var ()
	return &V1UserAssetsSSHCreateParams{

		Context: ctx,
	}
}

// NewV1UserAssetsSSHCreateParamsWithHTTPClient creates a new V1UserAssetsSSHCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1UserAssetsSSHCreateParamsWithHTTPClient(client *http.Client) *V1UserAssetsSSHCreateParams {
	var ()
	return &V1UserAssetsSSHCreateParams{
		HTTPClient: client,
	}
}

/*
V1UserAssetsSSHCreateParams contains all the parameters to send to the API endpoint
for the v1 user assets Ssh create operation typically these are written to a http.Request
*/
type V1UserAssetsSSHCreateParams struct {

	/*Body*/
	Body *models.V1UserAssetSSHEntity

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 user assets Ssh create params
func (o *V1UserAssetsSSHCreateParams) WithTimeout(timeout time.Duration) *V1UserAssetsSSHCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 user assets Ssh create params
func (o *V1UserAssetsSSHCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 user assets Ssh create params
func (o *V1UserAssetsSSHCreateParams) WithContext(ctx context.Context) *V1UserAssetsSSHCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 user assets Ssh create params
func (o *V1UserAssetsSSHCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 user assets Ssh create params
func (o *V1UserAssetsSSHCreateParams) WithHTTPClient(client *http.Client) *V1UserAssetsSSHCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 user assets Ssh create params
func (o *V1UserAssetsSSHCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 user assets Ssh create params
func (o *V1UserAssetsSSHCreateParams) WithBody(body *models.V1UserAssetSSHEntity) *V1UserAssetsSSHCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 user assets Ssh create params
func (o *V1UserAssetsSSHCreateParams) SetBody(body *models.V1UserAssetSSHEntity) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *V1UserAssetsSSHCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
