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
	"github.com/go-openapi/swag"

	"github.com/spectrocloud/palette-sdk-go/api/models"
)

// NewV1ClusterProfilesCreateParams creates a new V1ClusterProfilesCreateParams object
// with the default values initialized.
func NewV1ClusterProfilesCreateParams() *V1ClusterProfilesCreateParams {
	var ()
	return &V1ClusterProfilesCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1ClusterProfilesCreateParamsWithTimeout creates a new V1ClusterProfilesCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1ClusterProfilesCreateParamsWithTimeout(timeout time.Duration) *V1ClusterProfilesCreateParams {
	var ()
	return &V1ClusterProfilesCreateParams{

		timeout: timeout,
	}
}

// NewV1ClusterProfilesCreateParamsWithContext creates a new V1ClusterProfilesCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1ClusterProfilesCreateParamsWithContext(ctx context.Context) *V1ClusterProfilesCreateParams {
	var ()
	return &V1ClusterProfilesCreateParams{

		Context: ctx,
	}
}

// NewV1ClusterProfilesCreateParamsWithHTTPClient creates a new V1ClusterProfilesCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1ClusterProfilesCreateParamsWithHTTPClient(client *http.Client) *V1ClusterProfilesCreateParams {
	var ()
	return &V1ClusterProfilesCreateParams{
		HTTPClient: client,
	}
}

/*
V1ClusterProfilesCreateParams contains all the parameters to send to the API endpoint
for the v1 cluster profiles create operation typically these are written to a http.Request
*/
type V1ClusterProfilesCreateParams struct {

	/*Body*/
	Body *models.V1ClusterProfileEntity
	/*Publish
	  If true then cluster profile will be created and published in a single transaction

	*/
	Publish *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cluster profiles create params
func (o *V1ClusterProfilesCreateParams) WithTimeout(timeout time.Duration) *V1ClusterProfilesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cluster profiles create params
func (o *V1ClusterProfilesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cluster profiles create params
func (o *V1ClusterProfilesCreateParams) WithContext(ctx context.Context) *V1ClusterProfilesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cluster profiles create params
func (o *V1ClusterProfilesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cluster profiles create params
func (o *V1ClusterProfilesCreateParams) WithHTTPClient(client *http.Client) *V1ClusterProfilesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cluster profiles create params
func (o *V1ClusterProfilesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 cluster profiles create params
func (o *V1ClusterProfilesCreateParams) WithBody(body *models.V1ClusterProfileEntity) *V1ClusterProfilesCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 cluster profiles create params
func (o *V1ClusterProfilesCreateParams) SetBody(body *models.V1ClusterProfileEntity) {
	o.Body = body
}

// WithPublish adds the publish to the v1 cluster profiles create params
func (o *V1ClusterProfilesCreateParams) WithPublish(publish *bool) *V1ClusterProfilesCreateParams {
	o.SetPublish(publish)
	return o
}

// SetPublish adds the publish to the v1 cluster profiles create params
func (o *V1ClusterProfilesCreateParams) SetPublish(publish *bool) {
	o.Publish = publish
}

// WriteToRequest writes these params to a swagger request
func (o *V1ClusterProfilesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.Publish != nil {

		// query param publish
		var qrPublish bool
		if o.Publish != nil {
			qrPublish = *o.Publish
		}
		qPublish := swag.FormatBool(qrPublish)
		if qPublish != "" {
			if err := r.SetQueryParam("publish", qPublish); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
