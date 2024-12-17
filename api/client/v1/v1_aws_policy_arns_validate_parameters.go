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

// NewV1AwsPolicyArnsValidateParams creates a new V1AwsPolicyArnsValidateParams object
// with the default values initialized.
func NewV1AwsPolicyArnsValidateParams() *V1AwsPolicyArnsValidateParams {
	var ()
	return &V1AwsPolicyArnsValidateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1AwsPolicyArnsValidateParamsWithTimeout creates a new V1AwsPolicyArnsValidateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1AwsPolicyArnsValidateParamsWithTimeout(timeout time.Duration) *V1AwsPolicyArnsValidateParams {
	var ()
	return &V1AwsPolicyArnsValidateParams{

		timeout: timeout,
	}
}

// NewV1AwsPolicyArnsValidateParamsWithContext creates a new V1AwsPolicyArnsValidateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1AwsPolicyArnsValidateParamsWithContext(ctx context.Context) *V1AwsPolicyArnsValidateParams {
	var ()
	return &V1AwsPolicyArnsValidateParams{

		Context: ctx,
	}
}

// NewV1AwsPolicyArnsValidateParamsWithHTTPClient creates a new V1AwsPolicyArnsValidateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1AwsPolicyArnsValidateParamsWithHTTPClient(client *http.Client) *V1AwsPolicyArnsValidateParams {
	var ()
	return &V1AwsPolicyArnsValidateParams{
		HTTPClient: client,
	}
}

/*V1AwsPolicyArnsValidateParams contains all the parameters to send to the API endpoint
for the v1 aws policy arns validate operation typically these are written to a http.Request
*/
type V1AwsPolicyArnsValidateParams struct {

	/*CloudAccountUID
	  Uid for the specific AWS cloud account

	*/
	CloudAccountUID *string
	/*Spec
	  Request payload to validate AWS policy ARN

	*/
	Spec *models.V1AwsPolicyArnsSpec

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 aws policy arns validate params
func (o *V1AwsPolicyArnsValidateParams) WithTimeout(timeout time.Duration) *V1AwsPolicyArnsValidateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 aws policy arns validate params
func (o *V1AwsPolicyArnsValidateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 aws policy arns validate params
func (o *V1AwsPolicyArnsValidateParams) WithContext(ctx context.Context) *V1AwsPolicyArnsValidateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 aws policy arns validate params
func (o *V1AwsPolicyArnsValidateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 aws policy arns validate params
func (o *V1AwsPolicyArnsValidateParams) WithHTTPClient(client *http.Client) *V1AwsPolicyArnsValidateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 aws policy arns validate params
func (o *V1AwsPolicyArnsValidateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudAccountUID adds the cloudAccountUID to the v1 aws policy arns validate params
func (o *V1AwsPolicyArnsValidateParams) WithCloudAccountUID(cloudAccountUID *string) *V1AwsPolicyArnsValidateParams {
	o.SetCloudAccountUID(cloudAccountUID)
	return o
}

// SetCloudAccountUID adds the cloudAccountUid to the v1 aws policy arns validate params
func (o *V1AwsPolicyArnsValidateParams) SetCloudAccountUID(cloudAccountUID *string) {
	o.CloudAccountUID = cloudAccountUID
}

// WithSpec adds the spec to the v1 aws policy arns validate params
func (o *V1AwsPolicyArnsValidateParams) WithSpec(spec *models.V1AwsPolicyArnsSpec) *V1AwsPolicyArnsValidateParams {
	o.SetSpec(spec)
	return o
}

// SetSpec adds the spec to the v1 aws policy arns validate params
func (o *V1AwsPolicyArnsValidateParams) SetSpec(spec *models.V1AwsPolicyArnsSpec) {
	o.Spec = spec
}

// WriteToRequest writes these params to a swagger request
func (o *V1AwsPolicyArnsValidateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CloudAccountUID != nil {

		// query param cloudAccountUid
		var qrCloudAccountUID string
		if o.CloudAccountUID != nil {
			qrCloudAccountUID = *o.CloudAccountUID
		}
		qCloudAccountUID := qrCloudAccountUID
		if qCloudAccountUID != "" {
			if err := r.SetQueryParam("cloudAccountUid", qCloudAccountUID); err != nil {
				return err
			}
		}

	}

	if o.Spec != nil {
		if err := r.SetBodyParam(o.Spec); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
