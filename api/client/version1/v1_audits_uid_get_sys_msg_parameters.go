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
)

// NewV1AuditsUIDGetSysMsgParams creates a new V1AuditsUIDGetSysMsgParams object
// with the default values initialized.
func NewV1AuditsUIDGetSysMsgParams() *V1AuditsUIDGetSysMsgParams {
	var ()
	return &V1AuditsUIDGetSysMsgParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1AuditsUIDGetSysMsgParamsWithTimeout creates a new V1AuditsUIDGetSysMsgParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1AuditsUIDGetSysMsgParamsWithTimeout(timeout time.Duration) *V1AuditsUIDGetSysMsgParams {
	var ()
	return &V1AuditsUIDGetSysMsgParams{

		timeout: timeout,
	}
}

// NewV1AuditsUIDGetSysMsgParamsWithContext creates a new V1AuditsUIDGetSysMsgParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1AuditsUIDGetSysMsgParamsWithContext(ctx context.Context) *V1AuditsUIDGetSysMsgParams {
	var ()
	return &V1AuditsUIDGetSysMsgParams{

		Context: ctx,
	}
}

// NewV1AuditsUIDGetSysMsgParamsWithHTTPClient creates a new V1AuditsUIDGetSysMsgParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1AuditsUIDGetSysMsgParamsWithHTTPClient(client *http.Client) *V1AuditsUIDGetSysMsgParams {
	var ()
	return &V1AuditsUIDGetSysMsgParams{
		HTTPClient: client,
	}
}

/*
V1AuditsUIDGetSysMsgParams contains all the parameters to send to the API endpoint
for the v1 audits Uid get sys msg operation typically these are written to a http.Request
*/
type V1AuditsUIDGetSysMsgParams struct {

	/*UID
	  Specify the audit uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 audits Uid get sys msg params
func (o *V1AuditsUIDGetSysMsgParams) WithTimeout(timeout time.Duration) *V1AuditsUIDGetSysMsgParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 audits Uid get sys msg params
func (o *V1AuditsUIDGetSysMsgParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 audits Uid get sys msg params
func (o *V1AuditsUIDGetSysMsgParams) WithContext(ctx context.Context) *V1AuditsUIDGetSysMsgParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 audits Uid get sys msg params
func (o *V1AuditsUIDGetSysMsgParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 audits Uid get sys msg params
func (o *V1AuditsUIDGetSysMsgParams) WithHTTPClient(client *http.Client) *V1AuditsUIDGetSysMsgParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 audits Uid get sys msg params
func (o *V1AuditsUIDGetSysMsgParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUID adds the uid to the v1 audits Uid get sys msg params
func (o *V1AuditsUIDGetSysMsgParams) WithUID(uid string) *V1AuditsUIDGetSysMsgParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 audits Uid get sys msg params
func (o *V1AuditsUIDGetSysMsgParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1AuditsUIDGetSysMsgParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param uid
	if err := r.SetPathParam("uid", o.UID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
