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

// NewV1ClusterGroupsCreateParams creates a new V1ClusterGroupsCreateParams object
// with the default values initialized.
func NewV1ClusterGroupsCreateParams() *V1ClusterGroupsCreateParams {
	var ()
	return &V1ClusterGroupsCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1ClusterGroupsCreateParamsWithTimeout creates a new V1ClusterGroupsCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1ClusterGroupsCreateParamsWithTimeout(timeout time.Duration) *V1ClusterGroupsCreateParams {
	var ()
	return &V1ClusterGroupsCreateParams{

		timeout: timeout,
	}
}

// NewV1ClusterGroupsCreateParamsWithContext creates a new V1ClusterGroupsCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1ClusterGroupsCreateParamsWithContext(ctx context.Context) *V1ClusterGroupsCreateParams {
	var ()
	return &V1ClusterGroupsCreateParams{

		Context: ctx,
	}
}

// NewV1ClusterGroupsCreateParamsWithHTTPClient creates a new V1ClusterGroupsCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1ClusterGroupsCreateParamsWithHTTPClient(client *http.Client) *V1ClusterGroupsCreateParams {
	var ()
	return &V1ClusterGroupsCreateParams{
		HTTPClient: client,
	}
}

/*
V1ClusterGroupsCreateParams contains all the parameters to send to the API endpoint
for the v1 cluster groups create operation typically these are written to a http.Request
*/
type V1ClusterGroupsCreateParams struct {

	/*Body*/
	Body *models.V1ClusterGroupEntity

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cluster groups create params
func (o *V1ClusterGroupsCreateParams) WithTimeout(timeout time.Duration) *V1ClusterGroupsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cluster groups create params
func (o *V1ClusterGroupsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cluster groups create params
func (o *V1ClusterGroupsCreateParams) WithContext(ctx context.Context) *V1ClusterGroupsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cluster groups create params
func (o *V1ClusterGroupsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cluster groups create params
func (o *V1ClusterGroupsCreateParams) WithHTTPClient(client *http.Client) *V1ClusterGroupsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cluster groups create params
func (o *V1ClusterGroupsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 cluster groups create params
func (o *V1ClusterGroupsCreateParams) WithBody(body *models.V1ClusterGroupEntity) *V1ClusterGroupsCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 cluster groups create params
func (o *V1ClusterGroupsCreateParams) SetBody(body *models.V1ClusterGroupEntity) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *V1ClusterGroupsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
