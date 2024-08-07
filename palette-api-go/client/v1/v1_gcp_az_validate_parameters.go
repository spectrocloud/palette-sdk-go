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

	"github.com/spectrocloud/palette-api-go/models"
)

// NewV1GcpAzValidateParams creates a new V1GcpAzValidateParams object
// with the default values initialized.
func NewV1GcpAzValidateParams() *V1GcpAzValidateParams {
	var ()
	return &V1GcpAzValidateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1GcpAzValidateParamsWithTimeout creates a new V1GcpAzValidateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1GcpAzValidateParamsWithTimeout(timeout time.Duration) *V1GcpAzValidateParams {
	var ()
	return &V1GcpAzValidateParams{

		timeout: timeout,
	}
}

// NewV1GcpAzValidateParamsWithContext creates a new V1GcpAzValidateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1GcpAzValidateParamsWithContext(ctx context.Context) *V1GcpAzValidateParams {
	var ()
	return &V1GcpAzValidateParams{

		Context: ctx,
	}
}

// NewV1GcpAzValidateParamsWithHTTPClient creates a new V1GcpAzValidateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1GcpAzValidateParamsWithHTTPClient(client *http.Client) *V1GcpAzValidateParams {
	var ()
	return &V1GcpAzValidateParams{
		HTTPClient: client,
	}
}

/*V1GcpAzValidateParams contains all the parameters to send to the API endpoint
for the v1 gcp az validate operation typically these are written to a http.Request
*/
type V1GcpAzValidateParams struct {

	/*Entity
	  Uid for the specific GCP cloud account

	*/
	Entity *models.V1AzValidateEntity

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 gcp az validate params
func (o *V1GcpAzValidateParams) WithTimeout(timeout time.Duration) *V1GcpAzValidateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 gcp az validate params
func (o *V1GcpAzValidateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 gcp az validate params
func (o *V1GcpAzValidateParams) WithContext(ctx context.Context) *V1GcpAzValidateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 gcp az validate params
func (o *V1GcpAzValidateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 gcp az validate params
func (o *V1GcpAzValidateParams) WithHTTPClient(client *http.Client) *V1GcpAzValidateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 gcp az validate params
func (o *V1GcpAzValidateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEntity adds the entity to the v1 gcp az validate params
func (o *V1GcpAzValidateParams) WithEntity(entity *models.V1AzValidateEntity) *V1GcpAzValidateParams {
	o.SetEntity(entity)
	return o
}

// SetEntity adds the entity to the v1 gcp az validate params
func (o *V1GcpAzValidateParams) SetEntity(entity *models.V1AzValidateEntity) {
	o.Entity = entity
}

// WriteToRequest writes these params to a swagger request
func (o *V1GcpAzValidateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Entity != nil {
		if err := r.SetBodyParam(o.Entity); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
