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
)

// NewV1UsersKubectlSessionUIDParams creates a new V1UsersKubectlSessionUIDParams object
// with the default values initialized.
func NewV1UsersKubectlSessionUIDParams() *V1UsersKubectlSessionUIDParams {
	var ()
	return &V1UsersKubectlSessionUIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1UsersKubectlSessionUIDParamsWithTimeout creates a new V1UsersKubectlSessionUIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1UsersKubectlSessionUIDParamsWithTimeout(timeout time.Duration) *V1UsersKubectlSessionUIDParams {
	var ()
	return &V1UsersKubectlSessionUIDParams{

		timeout: timeout,
	}
}

// NewV1UsersKubectlSessionUIDParamsWithContext creates a new V1UsersKubectlSessionUIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1UsersKubectlSessionUIDParamsWithContext(ctx context.Context) *V1UsersKubectlSessionUIDParams {
	var ()
	return &V1UsersKubectlSessionUIDParams{

		Context: ctx,
	}
}

// NewV1UsersKubectlSessionUIDParamsWithHTTPClient creates a new V1UsersKubectlSessionUIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1UsersKubectlSessionUIDParamsWithHTTPClient(client *http.Client) *V1UsersKubectlSessionUIDParams {
	var ()
	return &V1UsersKubectlSessionUIDParams{
		HTTPClient: client,
	}
}

/*V1UsersKubectlSessionUIDParams contains all the parameters to send to the API endpoint
for the v1 users kubectl session Uid operation typically these are written to a http.Request
*/
type V1UsersKubectlSessionUIDParams struct {

	/*SessionUID*/
	SessionUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 users kubectl session Uid params
func (o *V1UsersKubectlSessionUIDParams) WithTimeout(timeout time.Duration) *V1UsersKubectlSessionUIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 users kubectl session Uid params
func (o *V1UsersKubectlSessionUIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 users kubectl session Uid params
func (o *V1UsersKubectlSessionUIDParams) WithContext(ctx context.Context) *V1UsersKubectlSessionUIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 users kubectl session Uid params
func (o *V1UsersKubectlSessionUIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 users kubectl session Uid params
func (o *V1UsersKubectlSessionUIDParams) WithHTTPClient(client *http.Client) *V1UsersKubectlSessionUIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 users kubectl session Uid params
func (o *V1UsersKubectlSessionUIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSessionUID adds the sessionUID to the v1 users kubectl session Uid params
func (o *V1UsersKubectlSessionUIDParams) WithSessionUID(sessionUID string) *V1UsersKubectlSessionUIDParams {
	o.SetSessionUID(sessionUID)
	return o
}

// SetSessionUID adds the sessionUid to the v1 users kubectl session Uid params
func (o *V1UsersKubectlSessionUIDParams) SetSessionUID(sessionUID string) {
	o.SessionUID = sessionUID
}

// WriteToRequest writes these params to a swagger request
func (o *V1UsersKubectlSessionUIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param sessionUid
	if err := r.SetPathParam("sessionUid", o.SessionUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
