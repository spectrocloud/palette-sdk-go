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

// NewV1ProjectsCreateParams creates a new V1ProjectsCreateParams object
// with the default values initialized.
func NewV1ProjectsCreateParams() *V1ProjectsCreateParams {
	var ()
	return &V1ProjectsCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1ProjectsCreateParamsWithTimeout creates a new V1ProjectsCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1ProjectsCreateParamsWithTimeout(timeout time.Duration) *V1ProjectsCreateParams {
	var ()
	return &V1ProjectsCreateParams{

		timeout: timeout,
	}
}

// NewV1ProjectsCreateParamsWithContext creates a new V1ProjectsCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1ProjectsCreateParamsWithContext(ctx context.Context) *V1ProjectsCreateParams {
	var ()
	return &V1ProjectsCreateParams{

		Context: ctx,
	}
}

// NewV1ProjectsCreateParamsWithHTTPClient creates a new V1ProjectsCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1ProjectsCreateParamsWithHTTPClient(client *http.Client) *V1ProjectsCreateParams {
	var ()
	return &V1ProjectsCreateParams{
		HTTPClient: client,
	}
}

/*V1ProjectsCreateParams contains all the parameters to send to the API endpoint
for the v1 projects create operation typically these are written to a http.Request
*/
type V1ProjectsCreateParams struct {

	/*Body*/
	Body *models.V1ProjectEntity

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 projects create params
func (o *V1ProjectsCreateParams) WithTimeout(timeout time.Duration) *V1ProjectsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 projects create params
func (o *V1ProjectsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 projects create params
func (o *V1ProjectsCreateParams) WithContext(ctx context.Context) *V1ProjectsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 projects create params
func (o *V1ProjectsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 projects create params
func (o *V1ProjectsCreateParams) WithHTTPClient(client *http.Client) *V1ProjectsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 projects create params
func (o *V1ProjectsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 projects create params
func (o *V1ProjectsCreateParams) WithBody(body *models.V1ProjectEntity) *V1ProjectsCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 projects create params
func (o *V1ProjectsCreateParams) SetBody(body *models.V1ProjectEntity) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *V1ProjectsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
