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

// NewV1AppDeploymentsClusterGroupCreateParams creates a new V1AppDeploymentsClusterGroupCreateParams object
// with the default values initialized.
func NewV1AppDeploymentsClusterGroupCreateParams() *V1AppDeploymentsClusterGroupCreateParams {
	var ()
	return &V1AppDeploymentsClusterGroupCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1AppDeploymentsClusterGroupCreateParamsWithTimeout creates a new V1AppDeploymentsClusterGroupCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1AppDeploymentsClusterGroupCreateParamsWithTimeout(timeout time.Duration) *V1AppDeploymentsClusterGroupCreateParams {
	var ()
	return &V1AppDeploymentsClusterGroupCreateParams{

		timeout: timeout,
	}
}

// NewV1AppDeploymentsClusterGroupCreateParamsWithContext creates a new V1AppDeploymentsClusterGroupCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1AppDeploymentsClusterGroupCreateParamsWithContext(ctx context.Context) *V1AppDeploymentsClusterGroupCreateParams {
	var ()
	return &V1AppDeploymentsClusterGroupCreateParams{

		Context: ctx,
	}
}

// NewV1AppDeploymentsClusterGroupCreateParamsWithHTTPClient creates a new V1AppDeploymentsClusterGroupCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1AppDeploymentsClusterGroupCreateParamsWithHTTPClient(client *http.Client) *V1AppDeploymentsClusterGroupCreateParams {
	var ()
	return &V1AppDeploymentsClusterGroupCreateParams{
		HTTPClient: client,
	}
}

/*V1AppDeploymentsClusterGroupCreateParams contains all the parameters to send to the API endpoint
for the v1 app deployments cluster group create operation typically these are written to a http.Request
*/
type V1AppDeploymentsClusterGroupCreateParams struct {

	/*Body*/
	Body *models.V1AppDeploymentClusterGroupEntity

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 app deployments cluster group create params
func (o *V1AppDeploymentsClusterGroupCreateParams) WithTimeout(timeout time.Duration) *V1AppDeploymentsClusterGroupCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 app deployments cluster group create params
func (o *V1AppDeploymentsClusterGroupCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 app deployments cluster group create params
func (o *V1AppDeploymentsClusterGroupCreateParams) WithContext(ctx context.Context) *V1AppDeploymentsClusterGroupCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 app deployments cluster group create params
func (o *V1AppDeploymentsClusterGroupCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 app deployments cluster group create params
func (o *V1AppDeploymentsClusterGroupCreateParams) WithHTTPClient(client *http.Client) *V1AppDeploymentsClusterGroupCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 app deployments cluster group create params
func (o *V1AppDeploymentsClusterGroupCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 app deployments cluster group create params
func (o *V1AppDeploymentsClusterGroupCreateParams) WithBody(body *models.V1AppDeploymentClusterGroupEntity) *V1AppDeploymentsClusterGroupCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 app deployments cluster group create params
func (o *V1AppDeploymentsClusterGroupCreateParams) SetBody(body *models.V1AppDeploymentClusterGroupEntity) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *V1AppDeploymentsClusterGroupCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
