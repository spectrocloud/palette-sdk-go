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

// NewV1NotificationsUIDDoneParams creates a new V1NotificationsUIDDoneParams object
// with the default values initialized.
func NewV1NotificationsUIDDoneParams() *V1NotificationsUIDDoneParams {
	var ()
	return &V1NotificationsUIDDoneParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV1NotificationsUIDDoneParamsWithTimeout creates a new V1NotificationsUIDDoneParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV1NotificationsUIDDoneParamsWithTimeout(timeout time.Duration) *V1NotificationsUIDDoneParams {
	var ()
	return &V1NotificationsUIDDoneParams{

		timeout: timeout,
	}
}

// NewV1NotificationsUIDDoneParamsWithContext creates a new V1NotificationsUIDDoneParams object
// with the default values initialized, and the ability to set a context for a request
func NewV1NotificationsUIDDoneParamsWithContext(ctx context.Context) *V1NotificationsUIDDoneParams {
	var ()
	return &V1NotificationsUIDDoneParams{

		Context: ctx,
	}
}

// NewV1NotificationsUIDDoneParamsWithHTTPClient creates a new V1NotificationsUIDDoneParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV1NotificationsUIDDoneParamsWithHTTPClient(client *http.Client) *V1NotificationsUIDDoneParams {
	var ()
	return &V1NotificationsUIDDoneParams{
		HTTPClient: client,
	}
}

/*
V1NotificationsUIDDoneParams contains all the parameters to send to the API endpoint
for the v1 notifications Uid done operation typically these are written to a http.Request
*/
type V1NotificationsUIDDoneParams struct {

	/*UID
	  Describes notification uid

	*/
	UID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v1 notifications Uid done params
func (o *V1NotificationsUIDDoneParams) WithTimeout(timeout time.Duration) *V1NotificationsUIDDoneParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 notifications Uid done params
func (o *V1NotificationsUIDDoneParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 notifications Uid done params
func (o *V1NotificationsUIDDoneParams) WithContext(ctx context.Context) *V1NotificationsUIDDoneParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 notifications Uid done params
func (o *V1NotificationsUIDDoneParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 notifications Uid done params
func (o *V1NotificationsUIDDoneParams) WithHTTPClient(client *http.Client) *V1NotificationsUIDDoneParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 notifications Uid done params
func (o *V1NotificationsUIDDoneParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUID adds the uid to the v1 notifications Uid done params
func (o *V1NotificationsUIDDoneParams) WithUID(uid string) *V1NotificationsUIDDoneParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the v1 notifications Uid done params
func (o *V1NotificationsUIDDoneParams) SetUID(uid string) {
	o.UID = uid
}

// WriteToRequest writes these params to a swagger request
func (o *V1NotificationsUIDDoneParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
