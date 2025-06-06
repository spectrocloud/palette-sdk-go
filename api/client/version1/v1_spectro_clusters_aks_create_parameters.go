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

// NewV1SpectroClustersAksCreateParams creates a new V1SpectroClustersAksCreateParams object
// with the default values initialized.
func NewV1SpectroClustersAksCreateParams() *V1SpectroClustersAksCreateParams {
	var ()
	return &V1SpectroClustersAksCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1SpectroClustersAksCreateParamsWithTimeout creates a new V1SpectroClustersAksCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1SpectroClustersAksCreateParamsWithTimeout(timeout time.Duration) *V1SpectroClustersAksCreateParams {
	var ()
	return &V1SpectroClustersAksCreateParams{

		timeout: timeout,
	}
}

// NewV1SpectroClustersAksCreateParamsWithContext creates a new V1SpectroClustersAksCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1SpectroClustersAksCreateParamsWithContext(ctx context.Context) *V1SpectroClustersAksCreateParams {
	var ()
	return &V1SpectroClustersAksCreateParams{

		Context: ctx,
	}
}

// NewV1SpectroClustersAksCreateParamsWithHTTPClient creates a new V1SpectroClustersAksCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1SpectroClustersAksCreateParamsWithHTTPClient(client *http.Client) *V1SpectroClustersAksCreateParams {
	var ()
	return &V1SpectroClustersAksCreateParams{
		HTTPClient: client,
	}
}

/*
V1SpectroClustersAksCreateParams contains all the parameters to send to the API endpoint
for the v1 spectro clusters aks create operation typically these are written to a http.Request
*/
type V1SpectroClustersAksCreateParams struct {

	/*Body*/
	Body *models.V1SpectroAzureClusterEntity

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 spectro clusters aks create params
func (o *V1SpectroClustersAksCreateParams) WithTimeout(timeout time.Duration) *V1SpectroClustersAksCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 spectro clusters aks create params
func (o *V1SpectroClustersAksCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 spectro clusters aks create params
func (o *V1SpectroClustersAksCreateParams) WithContext(ctx context.Context) *V1SpectroClustersAksCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 spectro clusters aks create params
func (o *V1SpectroClustersAksCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 spectro clusters aks create params
func (o *V1SpectroClustersAksCreateParams) WithHTTPClient(client *http.Client) *V1SpectroClustersAksCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 spectro clusters aks create params
func (o *V1SpectroClustersAksCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 spectro clusters aks create params
func (o *V1SpectroClustersAksCreateParams) WithBody(body *models.V1SpectroAzureClusterEntity) *V1SpectroClustersAksCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 spectro clusters aks create params
func (o *V1SpectroClustersAksCreateParams) SetBody(body *models.V1SpectroAzureClusterEntity) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *V1SpectroClustersAksCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
