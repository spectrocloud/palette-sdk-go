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

// NewV1CloudConfigsMaasMachinePoolCreateParams creates a new V1CloudConfigsMaasMachinePoolCreateParams object
// with the default values initialized.
func NewV1CloudConfigsMaasMachinePoolCreateParams() *V1CloudConfigsMaasMachinePoolCreateParams {
	var ()
	return &V1CloudConfigsMaasMachinePoolCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudConfigsMaasMachinePoolCreateParamsWithTimeout creates a new V1CloudConfigsMaasMachinePoolCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudConfigsMaasMachinePoolCreateParamsWithTimeout(timeout time.Duration) *V1CloudConfigsMaasMachinePoolCreateParams {
	var ()
	return &V1CloudConfigsMaasMachinePoolCreateParams{

		timeout: timeout,
	}
}

// NewV1CloudConfigsMaasMachinePoolCreateParamsWithContext creates a new V1CloudConfigsMaasMachinePoolCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudConfigsMaasMachinePoolCreateParamsWithContext(ctx context.Context) *V1CloudConfigsMaasMachinePoolCreateParams {
	var ()
	return &V1CloudConfigsMaasMachinePoolCreateParams{

		Context: ctx,
	}
}

// NewV1CloudConfigsMaasMachinePoolCreateParamsWithHTTPClient creates a new V1CloudConfigsMaasMachinePoolCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudConfigsMaasMachinePoolCreateParamsWithHTTPClient(client *http.Client) *V1CloudConfigsMaasMachinePoolCreateParams {
	var ()
	return &V1CloudConfigsMaasMachinePoolCreateParams{
		HTTPClient: client,
	}
}

/*V1CloudConfigsMaasMachinePoolCreateParams contains all the parameters to send to the API endpoint
for the v1 cloud configs maas machine pool create operation typically these are written to a http.Request
*/
type V1CloudConfigsMaasMachinePoolCreateParams struct {

	/*Body*/
	Body *models.V1MaasMachinePoolConfigEntity
	/*ConfigUID
	  Cluster's cloud config uid

	*/
	ConfigUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cloud configs maas machine pool create params
func (o *V1CloudConfigsMaasMachinePoolCreateParams) WithTimeout(timeout time.Duration) *V1CloudConfigsMaasMachinePoolCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud configs maas machine pool create params
func (o *V1CloudConfigsMaasMachinePoolCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud configs maas machine pool create params
func (o *V1CloudConfigsMaasMachinePoolCreateParams) WithContext(ctx context.Context) *V1CloudConfigsMaasMachinePoolCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud configs maas machine pool create params
func (o *V1CloudConfigsMaasMachinePoolCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud configs maas machine pool create params
func (o *V1CloudConfigsMaasMachinePoolCreateParams) WithHTTPClient(client *http.Client) *V1CloudConfigsMaasMachinePoolCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud configs maas machine pool create params
func (o *V1CloudConfigsMaasMachinePoolCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 cloud configs maas machine pool create params
func (o *V1CloudConfigsMaasMachinePoolCreateParams) WithBody(body *models.V1MaasMachinePoolConfigEntity) *V1CloudConfigsMaasMachinePoolCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 cloud configs maas machine pool create params
func (o *V1CloudConfigsMaasMachinePoolCreateParams) SetBody(body *models.V1MaasMachinePoolConfigEntity) {
	o.Body = body
}

// WithConfigUID adds the configUID to the v1 cloud configs maas machine pool create params
func (o *V1CloudConfigsMaasMachinePoolCreateParams) WithConfigUID(configUID string) *V1CloudConfigsMaasMachinePoolCreateParams {
	o.SetConfigUID(configUID)
	return o
}

// SetConfigUID adds the configUid to the v1 cloud configs maas machine pool create params
func (o *V1CloudConfigsMaasMachinePoolCreateParams) SetConfigUID(configUID string) {
	o.ConfigUID = configUID
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudConfigsMaasMachinePoolCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param configUid
	if err := r.SetPathParam("configUid", o.ConfigUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
