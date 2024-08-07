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

// NewV1CustomCloudTypeMetaUpdateParams creates a new V1CustomCloudTypeMetaUpdateParams object
// with the default values initialized.
func NewV1CustomCloudTypeMetaUpdateParams() *V1CustomCloudTypeMetaUpdateParams {
	var ()
	return &V1CustomCloudTypeMetaUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CustomCloudTypeMetaUpdateParamsWithTimeout creates a new V1CustomCloudTypeMetaUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CustomCloudTypeMetaUpdateParamsWithTimeout(timeout time.Duration) *V1CustomCloudTypeMetaUpdateParams {
	var ()
	return &V1CustomCloudTypeMetaUpdateParams{

		timeout: timeout,
	}
}

// NewV1CustomCloudTypeMetaUpdateParamsWithContext creates a new V1CustomCloudTypeMetaUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CustomCloudTypeMetaUpdateParamsWithContext(ctx context.Context) *V1CustomCloudTypeMetaUpdateParams {
	var ()
	return &V1CustomCloudTypeMetaUpdateParams{

		Context: ctx,
	}
}

// NewV1CustomCloudTypeMetaUpdateParamsWithHTTPClient creates a new V1CustomCloudTypeMetaUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CustomCloudTypeMetaUpdateParamsWithHTTPClient(client *http.Client) *V1CustomCloudTypeMetaUpdateParams {
	var ()
	return &V1CustomCloudTypeMetaUpdateParams{
		HTTPClient: client,
	}
}

/*V1CustomCloudTypeMetaUpdateParams contains all the parameters to send to the API endpoint
for the v1 custom cloud type meta update operation typically these are written to a http.Request
*/
type V1CustomCloudTypeMetaUpdateParams struct {

	/*Body
	  Request payload for custom cloud meta entity

	*/
	Body *models.V1CustomCloudRequestEntity
	/*CloudType
	  Unique cloud type

	*/
	CloudType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 custom cloud type meta update params
func (o *V1CustomCloudTypeMetaUpdateParams) WithTimeout(timeout time.Duration) *V1CustomCloudTypeMetaUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 custom cloud type meta update params
func (o *V1CustomCloudTypeMetaUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 custom cloud type meta update params
func (o *V1CustomCloudTypeMetaUpdateParams) WithContext(ctx context.Context) *V1CustomCloudTypeMetaUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 custom cloud type meta update params
func (o *V1CustomCloudTypeMetaUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 custom cloud type meta update params
func (o *V1CustomCloudTypeMetaUpdateParams) WithHTTPClient(client *http.Client) *V1CustomCloudTypeMetaUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 custom cloud type meta update params
func (o *V1CustomCloudTypeMetaUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 custom cloud type meta update params
func (o *V1CustomCloudTypeMetaUpdateParams) WithBody(body *models.V1CustomCloudRequestEntity) *V1CustomCloudTypeMetaUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 custom cloud type meta update params
func (o *V1CustomCloudTypeMetaUpdateParams) SetBody(body *models.V1CustomCloudRequestEntity) {
	o.Body = body
}

// WithCloudType adds the cloudType to the v1 custom cloud type meta update params
func (o *V1CustomCloudTypeMetaUpdateParams) WithCloudType(cloudType string) *V1CustomCloudTypeMetaUpdateParams {
	o.SetCloudType(cloudType)
	return o
}

// SetCloudType adds the cloudType to the v1 custom cloud type meta update params
func (o *V1CustomCloudTypeMetaUpdateParams) SetCloudType(cloudType string) {
	o.CloudType = cloudType
}

// WriteToRequest writes these params to a swagger request
func (o *V1CustomCloudTypeMetaUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param cloudType
	if err := r.SetPathParam("cloudType", o.CloudType); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
