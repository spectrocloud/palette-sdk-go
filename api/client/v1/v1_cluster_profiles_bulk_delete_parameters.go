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

// NewV1ClusterProfilesBulkDeleteParams creates a new V1ClusterProfilesBulkDeleteParams object
// with the default values initialized.
func NewV1ClusterProfilesBulkDeleteParams() *V1ClusterProfilesBulkDeleteParams {
	var ()
	return &V1ClusterProfilesBulkDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1ClusterProfilesBulkDeleteParamsWithTimeout creates a new V1ClusterProfilesBulkDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1ClusterProfilesBulkDeleteParamsWithTimeout(timeout time.Duration) *V1ClusterProfilesBulkDeleteParams {
	var ()
	return &V1ClusterProfilesBulkDeleteParams{

		timeout: timeout,
	}
}

// NewV1ClusterProfilesBulkDeleteParamsWithContext creates a new V1ClusterProfilesBulkDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1ClusterProfilesBulkDeleteParamsWithContext(ctx context.Context) *V1ClusterProfilesBulkDeleteParams {
	var ()
	return &V1ClusterProfilesBulkDeleteParams{

		Context: ctx,
	}
}

// NewV1ClusterProfilesBulkDeleteParamsWithHTTPClient creates a new V1ClusterProfilesBulkDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1ClusterProfilesBulkDeleteParamsWithHTTPClient(client *http.Client) *V1ClusterProfilesBulkDeleteParams {
	var ()
	return &V1ClusterProfilesBulkDeleteParams{
		HTTPClient: client,
	}
}

/*
V1ClusterProfilesBulkDeleteParams contains all the parameters to send to the API endpoint
for the v1 cluster profiles bulk delete operation typically these are written to a http.Request
*/
type V1ClusterProfilesBulkDeleteParams struct {

	/*Body*/
	Body *models.V1BulkDeleteRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cluster profiles bulk delete params
func (o *V1ClusterProfilesBulkDeleteParams) WithTimeout(timeout time.Duration) *V1ClusterProfilesBulkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cluster profiles bulk delete params
func (o *V1ClusterProfilesBulkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cluster profiles bulk delete params
func (o *V1ClusterProfilesBulkDeleteParams) WithContext(ctx context.Context) *V1ClusterProfilesBulkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cluster profiles bulk delete params
func (o *V1ClusterProfilesBulkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cluster profiles bulk delete params
func (o *V1ClusterProfilesBulkDeleteParams) WithHTTPClient(client *http.Client) *V1ClusterProfilesBulkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cluster profiles bulk delete params
func (o *V1ClusterProfilesBulkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 cluster profiles bulk delete params
func (o *V1ClusterProfilesBulkDeleteParams) WithBody(body *models.V1BulkDeleteRequest) *V1ClusterProfilesBulkDeleteParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 cluster profiles bulk delete params
func (o *V1ClusterProfilesBulkDeleteParams) SetBody(body *models.V1BulkDeleteRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *V1ClusterProfilesBulkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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