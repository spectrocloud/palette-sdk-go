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

// NewV1FeaturesUpdateParams creates a new V1FeaturesUpdateParams object
// with the default values initialized.
func NewV1FeaturesUpdateParams() *V1FeaturesUpdateParams {
	var ()
	return &V1FeaturesUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1FeaturesUpdateParamsWithTimeout creates a new V1FeaturesUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1FeaturesUpdateParamsWithTimeout(timeout time.Duration) *V1FeaturesUpdateParams {
	var ()
	return &V1FeaturesUpdateParams{

		timeout: timeout,
	}
}

// NewV1FeaturesUpdateParamsWithContext creates a new V1FeaturesUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1FeaturesUpdateParamsWithContext(ctx context.Context) *V1FeaturesUpdateParams {
	var ()
	return &V1FeaturesUpdateParams{

		Context: ctx,
	}
}

// NewV1FeaturesUpdateParamsWithHTTPClient creates a new V1FeaturesUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1FeaturesUpdateParamsWithHTTPClient(client *http.Client) *V1FeaturesUpdateParams {
	var ()
	return &V1FeaturesUpdateParams{
		HTTPClient: client,
	}
}

/*
V1FeaturesUpdateParams contains all the parameters to send to the API endpoint
for the v1 features update operation typically these are written to a http.Request
*/
type V1FeaturesUpdateParams struct {

	/*Body*/
	Body *models.V1FeatureUpdate
	/*UID
	  Specify the feature uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 features update params
func (o *V1FeaturesUpdateParams) WithTimeout(timeout time.Duration) *V1FeaturesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 features update params
func (o *V1FeaturesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 features update params
func (o *V1FeaturesUpdateParams) WithContext(ctx context.Context) *V1FeaturesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 features update params
func (o *V1FeaturesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 features update params
func (o *V1FeaturesUpdateParams) WithHTTPClient(client *http.Client) *V1FeaturesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 features update params
func (o *V1FeaturesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 features update params
func (o *V1FeaturesUpdateParams) WithBody(body *models.V1FeatureUpdate) *V1FeaturesUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 features update params
func (o *V1FeaturesUpdateParams) SetBody(body *models.V1FeatureUpdate) {
	o.Body = body
}

// WithUID adds the uid to the v1 features update params
func (o *V1FeaturesUpdateParams) WithUID(uid string) *V1FeaturesUpdateParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 features update params
func (o *V1FeaturesUpdateParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1FeaturesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param uid
	if err := r.SetPathParam("uid", o.UID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
