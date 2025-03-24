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

// NewV1CloudConfigsEksMachinePoolCreateParams creates a new V1CloudConfigsEksMachinePoolCreateParams object
// with the default values initialized.
func NewV1CloudConfigsEksMachinePoolCreateParams() *V1CloudConfigsEksMachinePoolCreateParams {
	var ()
	return &V1CloudConfigsEksMachinePoolCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1CloudConfigsEksMachinePoolCreateParamsWithTimeout creates a new V1CloudConfigsEksMachinePoolCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1CloudConfigsEksMachinePoolCreateParamsWithTimeout(timeout time.Duration) *V1CloudConfigsEksMachinePoolCreateParams {
	var ()
	return &V1CloudConfigsEksMachinePoolCreateParams{

		timeout: timeout,
	}
}

// NewV1CloudConfigsEksMachinePoolCreateParamsWithContext creates a new V1CloudConfigsEksMachinePoolCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1CloudConfigsEksMachinePoolCreateParamsWithContext(ctx context.Context) *V1CloudConfigsEksMachinePoolCreateParams {
	var ()
	return &V1CloudConfigsEksMachinePoolCreateParams{

		Context: ctx,
	}
}

// NewV1CloudConfigsEksMachinePoolCreateParamsWithHTTPClient creates a new V1CloudConfigsEksMachinePoolCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1CloudConfigsEksMachinePoolCreateParamsWithHTTPClient(client *http.Client) *V1CloudConfigsEksMachinePoolCreateParams {
	var ()
	return &V1CloudConfigsEksMachinePoolCreateParams{
		HTTPClient: client,
	}
}

/*V1CloudConfigsEksMachinePoolCreateParams contains all the parameters to send to the API endpoint
for the v1 cloud configs eks machine pool create operation typically these are written to a http.Request
*/
type V1CloudConfigsEksMachinePoolCreateParams struct {

	/*Body*/
	Body *models.V1EksMachinePoolConfigEntity
	/*ConfigUID
	  Cluster's cloud config uid

	*/
	ConfigUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 cloud configs eks machine pool create params
func (o *V1CloudConfigsEksMachinePoolCreateParams) WithTimeout(timeout time.Duration) *V1CloudConfigsEksMachinePoolCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 cloud configs eks machine pool create params
func (o *V1CloudConfigsEksMachinePoolCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 cloud configs eks machine pool create params
func (o *V1CloudConfigsEksMachinePoolCreateParams) WithContext(ctx context.Context) *V1CloudConfigsEksMachinePoolCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 cloud configs eks machine pool create params
func (o *V1CloudConfigsEksMachinePoolCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 cloud configs eks machine pool create params
func (o *V1CloudConfigsEksMachinePoolCreateParams) WithHTTPClient(client *http.Client) *V1CloudConfigsEksMachinePoolCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 cloud configs eks machine pool create params
func (o *V1CloudConfigsEksMachinePoolCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 cloud configs eks machine pool create params
func (o *V1CloudConfigsEksMachinePoolCreateParams) WithBody(body *models.V1EksMachinePoolConfigEntity) *V1CloudConfigsEksMachinePoolCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 cloud configs eks machine pool create params
func (o *V1CloudConfigsEksMachinePoolCreateParams) SetBody(body *models.V1EksMachinePoolConfigEntity) {
	o.Body = body
}

// WithConfigUID adds the configUID to the v1 cloud configs eks machine pool create params
func (o *V1CloudConfigsEksMachinePoolCreateParams) WithConfigUID(configUID string) *V1CloudConfigsEksMachinePoolCreateParams {
	o.SetConfigUID(configUID)
	return o
}

// SetConfigUID adds the configUid to the v1 cloud configs eks machine pool create params
func (o *V1CloudConfigsEksMachinePoolCreateParams) SetConfigUID(configUID string) {
	o.ConfigUID = configUID
}

// WriteToRequest writes these params to a swagger request
func (o *V1CloudConfigsEksMachinePoolCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
