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

// NewV1SystemConfigReverseProxyUpdateParams creates a new V1SystemConfigReverseProxyUpdateParams object
// with the default values initialized.
func NewV1SystemConfigReverseProxyUpdateParams() *V1SystemConfigReverseProxyUpdateParams {
	var ()
	return &V1SystemConfigReverseProxyUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1SystemConfigReverseProxyUpdateParamsWithTimeout creates a new V1SystemConfigReverseProxyUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1SystemConfigReverseProxyUpdateParamsWithTimeout(timeout time.Duration) *V1SystemConfigReverseProxyUpdateParams {
	var ()
	return &V1SystemConfigReverseProxyUpdateParams{

		timeout: timeout,
	}
}

// NewV1SystemConfigReverseProxyUpdateParamsWithContext creates a new V1SystemConfigReverseProxyUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1SystemConfigReverseProxyUpdateParamsWithContext(ctx context.Context) *V1SystemConfigReverseProxyUpdateParams {
	var ()
	return &V1SystemConfigReverseProxyUpdateParams{

		Context: ctx,
	}
}

// NewV1SystemConfigReverseProxyUpdateParamsWithHTTPClient creates a new V1SystemConfigReverseProxyUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1SystemConfigReverseProxyUpdateParamsWithHTTPClient(client *http.Client) *V1SystemConfigReverseProxyUpdateParams {
	var ()
	return &V1SystemConfigReverseProxyUpdateParams{
		HTTPClient: client,
	}
}

/*V1SystemConfigReverseProxyUpdateParams contains all the parameters to send to the API endpoint
for the v1 system config reverse proxy update operation typically these are written to a http.Request
*/
type V1SystemConfigReverseProxyUpdateParams struct {

	/*Body*/
	Body *models.V1SystemReverseProxy

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 system config reverse proxy update params
func (o *V1SystemConfigReverseProxyUpdateParams) WithTimeout(timeout time.Duration) *V1SystemConfigReverseProxyUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 system config reverse proxy update params
func (o *V1SystemConfigReverseProxyUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 system config reverse proxy update params
func (o *V1SystemConfigReverseProxyUpdateParams) WithContext(ctx context.Context) *V1SystemConfigReverseProxyUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 system config reverse proxy update params
func (o *V1SystemConfigReverseProxyUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 system config reverse proxy update params
func (o *V1SystemConfigReverseProxyUpdateParams) WithHTTPClient(client *http.Client) *V1SystemConfigReverseProxyUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 system config reverse proxy update params
func (o *V1SystemConfigReverseProxyUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 system config reverse proxy update params
func (o *V1SystemConfigReverseProxyUpdateParams) WithBody(body *models.V1SystemReverseProxy) *V1SystemConfigReverseProxyUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 system config reverse proxy update params
func (o *V1SystemConfigReverseProxyUpdateParams) SetBody(body *models.V1SystemReverseProxy) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *V1SystemConfigReverseProxyUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
