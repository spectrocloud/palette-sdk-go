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

// NewV1SpectroClustersTkeValidateParams creates a new V1SpectroClustersTkeValidateParams object
// with the default values initialized.
func NewV1SpectroClustersTkeValidateParams() *V1SpectroClustersTkeValidateParams {
	var ()
	return &V1SpectroClustersTkeValidateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1SpectroClustersTkeValidateParamsWithTimeout creates a new V1SpectroClustersTkeValidateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1SpectroClustersTkeValidateParamsWithTimeout(timeout time.Duration) *V1SpectroClustersTkeValidateParams {
	var ()
	return &V1SpectroClustersTkeValidateParams{

		timeout: timeout,
	}
}

// NewV1SpectroClustersTkeValidateParamsWithContext creates a new V1SpectroClustersTkeValidateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1SpectroClustersTkeValidateParamsWithContext(ctx context.Context) *V1SpectroClustersTkeValidateParams {
	var ()
	return &V1SpectroClustersTkeValidateParams{

		Context: ctx,
	}
}

// NewV1SpectroClustersTkeValidateParamsWithHTTPClient creates a new V1SpectroClustersTkeValidateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1SpectroClustersTkeValidateParamsWithHTTPClient(client *http.Client) *V1SpectroClustersTkeValidateParams {
	var ()
	return &V1SpectroClustersTkeValidateParams{
		HTTPClient: client,
	}
}

/*
V1SpectroClustersTkeValidateParams contains all the parameters to send to the API endpoint
for the v1 spectro clusters tke validate operation typically these are written to a http.Request
*/
type V1SpectroClustersTkeValidateParams struct {

	/*Body*/
	Body *models.V1SpectroTencentClusterEntity

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 spectro clusters tke validate params
func (o *V1SpectroClustersTkeValidateParams) WithTimeout(timeout time.Duration) *V1SpectroClustersTkeValidateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 spectro clusters tke validate params
func (o *V1SpectroClustersTkeValidateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 spectro clusters tke validate params
func (o *V1SpectroClustersTkeValidateParams) WithContext(ctx context.Context) *V1SpectroClustersTkeValidateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 spectro clusters tke validate params
func (o *V1SpectroClustersTkeValidateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 spectro clusters tke validate params
func (o *V1SpectroClustersTkeValidateParams) WithHTTPClient(client *http.Client) *V1SpectroClustersTkeValidateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 spectro clusters tke validate params
func (o *V1SpectroClustersTkeValidateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 spectro clusters tke validate params
func (o *V1SpectroClustersTkeValidateParams) WithBody(body *models.V1SpectroTencentClusterEntity) *V1SpectroClustersTkeValidateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 spectro clusters tke validate params
func (o *V1SpectroClustersTkeValidateParams) SetBody(body *models.V1SpectroTencentClusterEntity) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *V1SpectroClustersTkeValidateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
