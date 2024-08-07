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

// NewV1CustomCloudTypeCloudAccountKeysUpdateParams creates a new V1CustomCloudTypeCloudAccountKeysUpdateParams object
// with the default values initialized.
func NewV1CustomCloudTypeCloudAccountKeysUpdateParams() *V1CustomCloudTypeCloudAccountKeysUpdateParams {
	var ()
	return &V1CustomCloudTypeCloudAccountKeysUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CustomCloudTypeCloudAccountKeysUpdateParamsWithTimeout creates a new V1CustomCloudTypeCloudAccountKeysUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CustomCloudTypeCloudAccountKeysUpdateParamsWithTimeout(timeout time.Duration) *V1CustomCloudTypeCloudAccountKeysUpdateParams {
	var ()
	return &V1CustomCloudTypeCloudAccountKeysUpdateParams{

		timeout: timeout,
	}
}

// NewV1CustomCloudTypeCloudAccountKeysUpdateParamsWithContext creates a new V1CustomCloudTypeCloudAccountKeysUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CustomCloudTypeCloudAccountKeysUpdateParamsWithContext(ctx context.Context) *V1CustomCloudTypeCloudAccountKeysUpdateParams {
	var ()
	return &V1CustomCloudTypeCloudAccountKeysUpdateParams{

		Context: ctx,
	}
}

// NewV1CustomCloudTypeCloudAccountKeysUpdateParamsWithHTTPClient creates a new V1CustomCloudTypeCloudAccountKeysUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CustomCloudTypeCloudAccountKeysUpdateParamsWithHTTPClient(client *http.Client) *V1CustomCloudTypeCloudAccountKeysUpdateParams {
	var ()
	return &V1CustomCloudTypeCloudAccountKeysUpdateParams{
		HTTPClient: client,
	}
}

/*V1CustomCloudTypeCloudAccountKeysUpdateParams contains all the parameters to send to the API endpoint
for the v1 custom cloud type cloud account keys update operation typically these are written to a http.Request
*/
type V1CustomCloudTypeCloudAccountKeysUpdateParams struct {

	/*Body
	  Request payload for custom cloud meta entity

	*/
	Body *models.V1CustomCloudTypeCloudAccountKeys
	/*CloudType
	  Unique cloud type

	*/
	CloudType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 custom cloud type cloud account keys update params
func (o *V1CustomCloudTypeCloudAccountKeysUpdateParams) WithTimeout(timeout time.Duration) *V1CustomCloudTypeCloudAccountKeysUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 custom cloud type cloud account keys update params
func (o *V1CustomCloudTypeCloudAccountKeysUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 custom cloud type cloud account keys update params
func (o *V1CustomCloudTypeCloudAccountKeysUpdateParams) WithContext(ctx context.Context) *V1CustomCloudTypeCloudAccountKeysUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 custom cloud type cloud account keys update params
func (o *V1CustomCloudTypeCloudAccountKeysUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 custom cloud type cloud account keys update params
func (o *V1CustomCloudTypeCloudAccountKeysUpdateParams) WithHTTPClient(client *http.Client) *V1CustomCloudTypeCloudAccountKeysUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 custom cloud type cloud account keys update params
func (o *V1CustomCloudTypeCloudAccountKeysUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 custom cloud type cloud account keys update params
func (o *V1CustomCloudTypeCloudAccountKeysUpdateParams) WithBody(body *models.V1CustomCloudTypeCloudAccountKeys) *V1CustomCloudTypeCloudAccountKeysUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 custom cloud type cloud account keys update params
func (o *V1CustomCloudTypeCloudAccountKeysUpdateParams) SetBody(body *models.V1CustomCloudTypeCloudAccountKeys) {
	o.Body = body
}

// WithCloudType adds the cloudType to the v1 custom cloud type cloud account keys update params
func (o *V1CustomCloudTypeCloudAccountKeysUpdateParams) WithCloudType(cloudType string) *V1CustomCloudTypeCloudAccountKeysUpdateParams {
	o.SetCloudType(cloudType)
	return o
}

// SetCloudType adds the cloudType to the v1 custom cloud type cloud account keys update params
func (o *V1CustomCloudTypeCloudAccountKeysUpdateParams) SetCloudType(cloudType string) {
	o.CloudType = cloudType
}

// WriteToRequest writes these params to a swagger request
func (o *V1CustomCloudTypeCloudAccountKeysUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
