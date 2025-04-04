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

// NewV1UsersStatusLoginModeParams creates a new V1UsersStatusLoginModeParams object
// with the default values initialized.
func NewV1UsersStatusLoginModeParams() *V1UsersStatusLoginModeParams {
	var ()
	return &V1UsersStatusLoginModeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1UsersStatusLoginModeParamsWithTimeout creates a new V1UsersStatusLoginModeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1UsersStatusLoginModeParamsWithTimeout(timeout time.Duration) *V1UsersStatusLoginModeParams {
	var ()
	return &V1UsersStatusLoginModeParams{

		timeout: timeout,
	}
}

// NewV1UsersStatusLoginModeParamsWithContext creates a new V1UsersStatusLoginModeParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1UsersStatusLoginModeParamsWithContext(ctx context.Context) *V1UsersStatusLoginModeParams {
	var ()
	return &V1UsersStatusLoginModeParams{

		Context: ctx,
	}
}

// NewV1UsersStatusLoginModeParamsWithHTTPClient creates a new V1UsersStatusLoginModeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1UsersStatusLoginModeParamsWithHTTPClient(client *http.Client) *V1UsersStatusLoginModeParams {
	var ()
	return &V1UsersStatusLoginModeParams{
		HTTPClient: client,
	}
}

/*
V1UsersStatusLoginModeParams contains all the parameters to send to the API endpoint
for the v1 users status login mode operation typically these are written to a http.Request
*/
type V1UsersStatusLoginModeParams struct {

	/*Body*/
	Body *models.V1UserStatusLoginMode
	/*UID*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 users status login mode params
func (o *V1UsersStatusLoginModeParams) WithTimeout(timeout time.Duration) *V1UsersStatusLoginModeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 users status login mode params
func (o *V1UsersStatusLoginModeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 users status login mode params
func (o *V1UsersStatusLoginModeParams) WithContext(ctx context.Context) *V1UsersStatusLoginModeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 users status login mode params
func (o *V1UsersStatusLoginModeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 users status login mode params
func (o *V1UsersStatusLoginModeParams) WithHTTPClient(client *http.Client) *V1UsersStatusLoginModeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 users status login mode params
func (o *V1UsersStatusLoginModeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 users status login mode params
func (o *V1UsersStatusLoginModeParams) WithBody(body *models.V1UserStatusLoginMode) *V1UsersStatusLoginModeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 users status login mode params
func (o *V1UsersStatusLoginModeParams) SetBody(body *models.V1UserStatusLoginMode) {
	o.Body = body
}

// WithUID adds the uid to the v1 users status login mode params
func (o *V1UsersStatusLoginModeParams) WithUID(uid string) *V1UsersStatusLoginModeParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 users status login mode params
func (o *V1UsersStatusLoginModeParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1UsersStatusLoginModeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
