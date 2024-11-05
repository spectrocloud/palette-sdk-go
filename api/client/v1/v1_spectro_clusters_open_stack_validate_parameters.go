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

// NewV1SpectroClustersOpenStackValidateParams creates a new V1SpectroClustersOpenStackValidateParams object
// with the default values initialized.
func NewV1SpectroClustersOpenStackValidateParams() *V1SpectroClustersOpenStackValidateParams {
	var ()
	return &V1SpectroClustersOpenStackValidateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1SpectroClustersOpenStackValidateParamsWithTimeout creates a new V1SpectroClustersOpenStackValidateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1SpectroClustersOpenStackValidateParamsWithTimeout(timeout time.Duration) *V1SpectroClustersOpenStackValidateParams {
	var ()
	return &V1SpectroClustersOpenStackValidateParams{

		timeout: timeout,
	}
}

// NewV1SpectroClustersOpenStackValidateParamsWithContext creates a new V1SpectroClustersOpenStackValidateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1SpectroClustersOpenStackValidateParamsWithContext(ctx context.Context) *V1SpectroClustersOpenStackValidateParams {
	var ()
	return &V1SpectroClustersOpenStackValidateParams{

		Context: ctx,
	}
}

// NewV1SpectroClustersOpenStackValidateParamsWithHTTPClient creates a new V1SpectroClustersOpenStackValidateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1SpectroClustersOpenStackValidateParamsWithHTTPClient(client *http.Client) *V1SpectroClustersOpenStackValidateParams {
	var ()
	return &V1SpectroClustersOpenStackValidateParams{
		HTTPClient: client,
	}
}

/*
V1SpectroClustersOpenStackValidateParams contains all the parameters to send to the API endpoint
for the v1 spectro clusters open stack validate operation typically these are written to a http.Request
*/
type V1SpectroClustersOpenStackValidateParams struct {

	/*Body*/
	Body *models.V1SpectroOpenStackClusterEntity

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 spectro clusters open stack validate params
func (o *V1SpectroClustersOpenStackValidateParams) WithTimeout(timeout time.Duration) *V1SpectroClustersOpenStackValidateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 spectro clusters open stack validate params
func (o *V1SpectroClustersOpenStackValidateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 spectro clusters open stack validate params
func (o *V1SpectroClustersOpenStackValidateParams) WithContext(ctx context.Context) *V1SpectroClustersOpenStackValidateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 spectro clusters open stack validate params
func (o *V1SpectroClustersOpenStackValidateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 spectro clusters open stack validate params
func (o *V1SpectroClustersOpenStackValidateParams) WithHTTPClient(client *http.Client) *V1SpectroClustersOpenStackValidateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 spectro clusters open stack validate params
func (o *V1SpectroClustersOpenStackValidateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 spectro clusters open stack validate params
func (o *V1SpectroClustersOpenStackValidateParams) WithBody(body *models.V1SpectroOpenStackClusterEntity) *V1SpectroClustersOpenStackValidateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 spectro clusters open stack validate params
func (o *V1SpectroClustersOpenStackValidateParams) SetBody(body *models.V1SpectroOpenStackClusterEntity) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *V1SpectroClustersOpenStackValidateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
