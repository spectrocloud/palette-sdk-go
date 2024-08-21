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

// NewV1SpectroClustersValidatePacksParams creates a new V1SpectroClustersValidatePacksParams object
// with the default values initialized.
func NewV1SpectroClustersValidatePacksParams() *V1SpectroClustersValidatePacksParams {
	var ()
	return &V1SpectroClustersValidatePacksParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1SpectroClustersValidatePacksParamsWithTimeout creates a new V1SpectroClustersValidatePacksParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1SpectroClustersValidatePacksParamsWithTimeout(timeout time.Duration) *V1SpectroClustersValidatePacksParams {
	var ()
	return &V1SpectroClustersValidatePacksParams{

		timeout: timeout,
	}
}

// NewV1SpectroClustersValidatePacksParamsWithContext creates a new V1SpectroClustersValidatePacksParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1SpectroClustersValidatePacksParamsWithContext(ctx context.Context) *V1SpectroClustersValidatePacksParams {
	var ()
	return &V1SpectroClustersValidatePacksParams{

		Context: ctx,
	}
}

// NewV1SpectroClustersValidatePacksParamsWithHTTPClient creates a new V1SpectroClustersValidatePacksParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1SpectroClustersValidatePacksParamsWithHTTPClient(client *http.Client) *V1SpectroClustersValidatePacksParams {
	var ()
	return &V1SpectroClustersValidatePacksParams{
		HTTPClient: client,
	}
}

/*
V1SpectroClustersValidatePacksParams contains all the parameters to send to the API endpoint
for the v1 spectro clusters validate packs operation typically these are written to a http.Request
*/
type V1SpectroClustersValidatePacksParams struct {

	/*Body*/
	Body *models.V1SpectroClusterPacksEntity

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 spectro clusters validate packs params
func (o *V1SpectroClustersValidatePacksParams) WithTimeout(timeout time.Duration) *V1SpectroClustersValidatePacksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 spectro clusters validate packs params
func (o *V1SpectroClustersValidatePacksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 spectro clusters validate packs params
func (o *V1SpectroClustersValidatePacksParams) WithContext(ctx context.Context) *V1SpectroClustersValidatePacksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 spectro clusters validate packs params
func (o *V1SpectroClustersValidatePacksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 spectro clusters validate packs params
func (o *V1SpectroClustersValidatePacksParams) WithHTTPClient(client *http.Client) *V1SpectroClustersValidatePacksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 spectro clusters validate packs params
func (o *V1SpectroClustersValidatePacksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 spectro clusters validate packs params
func (o *V1SpectroClustersValidatePacksParams) WithBody(body *models.V1SpectroClusterPacksEntity) *V1SpectroClustersValidatePacksParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 spectro clusters validate packs params
func (o *V1SpectroClustersValidatePacksParams) SetBody(body *models.V1SpectroClusterPacksEntity) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *V1SpectroClustersValidatePacksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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