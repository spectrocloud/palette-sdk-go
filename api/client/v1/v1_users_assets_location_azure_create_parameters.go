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

// NewV1UsersAssetsLocationAzureCreateParams creates a new V1UsersAssetsLocationAzureCreateParams object
// with the default values initialized.
func NewV1UsersAssetsLocationAzureCreateParams() *V1UsersAssetsLocationAzureCreateParams {
	var ()
	return &V1UsersAssetsLocationAzureCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1UsersAssetsLocationAzureCreateParamsWithTimeout creates a new V1UsersAssetsLocationAzureCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1UsersAssetsLocationAzureCreateParamsWithTimeout(timeout time.Duration) *V1UsersAssetsLocationAzureCreateParams {
	var ()
	return &V1UsersAssetsLocationAzureCreateParams{

		timeout: timeout,
	}
}

// NewV1UsersAssetsLocationAzureCreateParamsWithContext creates a new V1UsersAssetsLocationAzureCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1UsersAssetsLocationAzureCreateParamsWithContext(ctx context.Context) *V1UsersAssetsLocationAzureCreateParams {
	var ()
	return &V1UsersAssetsLocationAzureCreateParams{

		Context: ctx,
	}
}

// NewV1UsersAssetsLocationAzureCreateParamsWithHTTPClient creates a new V1UsersAssetsLocationAzureCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1UsersAssetsLocationAzureCreateParamsWithHTTPClient(client *http.Client) *V1UsersAssetsLocationAzureCreateParams {
	var ()
	return &V1UsersAssetsLocationAzureCreateParams{
		HTTPClient: client,
	}
}

/*V1UsersAssetsLocationAzureCreateParams contains all the parameters to send to the API endpoint
for the v1 users assets location azure create operation typically these are written to a http.Request
*/
type V1UsersAssetsLocationAzureCreateParams struct {

	/*Body*/
	Body *models.V1UserAssetsLocationAzure

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 users assets location azure create params
func (o *V1UsersAssetsLocationAzureCreateParams) WithTimeout(timeout time.Duration) *V1UsersAssetsLocationAzureCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 users assets location azure create params
func (o *V1UsersAssetsLocationAzureCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 users assets location azure create params
func (o *V1UsersAssetsLocationAzureCreateParams) WithContext(ctx context.Context) *V1UsersAssetsLocationAzureCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 users assets location azure create params
func (o *V1UsersAssetsLocationAzureCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 users assets location azure create params
func (o *V1UsersAssetsLocationAzureCreateParams) WithHTTPClient(client *http.Client) *V1UsersAssetsLocationAzureCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 users assets location azure create params
func (o *V1UsersAssetsLocationAzureCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 users assets location azure create params
func (o *V1UsersAssetsLocationAzureCreateParams) WithBody(body *models.V1UserAssetsLocationAzure) *V1UsersAssetsLocationAzureCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 users assets location azure create params
func (o *V1UsersAssetsLocationAzureCreateParams) SetBody(body *models.V1UserAssetsLocationAzure) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *V1UsersAssetsLocationAzureCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
