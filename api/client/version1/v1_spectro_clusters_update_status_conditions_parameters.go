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

// NewV1SpectroClustersUpdateStatusConditionsParams creates a new V1SpectroClustersUpdateStatusConditionsParams object
// with the default values initialized.
func NewV1SpectroClustersUpdateStatusConditionsParams() *V1SpectroClustersUpdateStatusConditionsParams {
	var ()
	return &V1SpectroClustersUpdateStatusConditionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1SpectroClustersUpdateStatusConditionsParamsWithTimeout creates a new V1SpectroClustersUpdateStatusConditionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1SpectroClustersUpdateStatusConditionsParamsWithTimeout(timeout time.Duration) *V1SpectroClustersUpdateStatusConditionsParams {
	var ()
	return &V1SpectroClustersUpdateStatusConditionsParams{

		timeout: timeout,
	}
}

// NewV1SpectroClustersUpdateStatusConditionsParamsWithContext creates a new V1SpectroClustersUpdateStatusConditionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1SpectroClustersUpdateStatusConditionsParamsWithContext(ctx context.Context) *V1SpectroClustersUpdateStatusConditionsParams {
	var ()
	return &V1SpectroClustersUpdateStatusConditionsParams{

		Context: ctx,
	}
}

// NewV1SpectroClustersUpdateStatusConditionsParamsWithHTTPClient creates a new V1SpectroClustersUpdateStatusConditionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1SpectroClustersUpdateStatusConditionsParamsWithHTTPClient(client *http.Client) *V1SpectroClustersUpdateStatusConditionsParams {
	var ()
	return &V1SpectroClustersUpdateStatusConditionsParams{
		HTTPClient: client,
	}
}

/*
V1SpectroClustersUpdateStatusConditionsParams contains all the parameters to send to the API endpoint
for the v1 spectro clusters update status conditions operation typically these are written to a http.Request
*/
type V1SpectroClustersUpdateStatusConditionsParams struct {

	/*Body*/
	Body []*models.V1ClusterCondition
	/*UID*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 spectro clusters update status conditions params
func (o *V1SpectroClustersUpdateStatusConditionsParams) WithTimeout(timeout time.Duration) *V1SpectroClustersUpdateStatusConditionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 spectro clusters update status conditions params
func (o *V1SpectroClustersUpdateStatusConditionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 spectro clusters update status conditions params
func (o *V1SpectroClustersUpdateStatusConditionsParams) WithContext(ctx context.Context) *V1SpectroClustersUpdateStatusConditionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 spectro clusters update status conditions params
func (o *V1SpectroClustersUpdateStatusConditionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 spectro clusters update status conditions params
func (o *V1SpectroClustersUpdateStatusConditionsParams) WithHTTPClient(client *http.Client) *V1SpectroClustersUpdateStatusConditionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 spectro clusters update status conditions params
func (o *V1SpectroClustersUpdateStatusConditionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 spectro clusters update status conditions params
func (o *V1SpectroClustersUpdateStatusConditionsParams) WithBody(body []*models.V1ClusterCondition) *V1SpectroClustersUpdateStatusConditionsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 spectro clusters update status conditions params
func (o *V1SpectroClustersUpdateStatusConditionsParams) SetBody(body []*models.V1ClusterCondition) {
	o.Body = body
}

// WithUID adds the uid to the v1 spectro clusters update status conditions params
func (o *V1SpectroClustersUpdateStatusConditionsParams) WithUID(uid string) *V1SpectroClustersUpdateStatusConditionsParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 spectro clusters update status conditions params
func (o *V1SpectroClustersUpdateStatusConditionsParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1SpectroClustersUpdateStatusConditionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
